package nats

import (
	"backend/core"
	"backend/db/redis"
	"backend/models"
	"fmt"
	"os"
	"time"

	json "github.com/goccy/go-json"

	"github.com/nats-io/nats.go"
)

var (
	log = core.GetLogger()
)

var NatsConnection *nats.Conn

func Start() {
	url := os.Getenv("NATS_URL")

	conn, err := nats.Connect(url, nats.UserInfo("admin", "meamadmin321"))
	if err != nil {
		log.Fatal(err)
	}
	NatsConnection = conn

	log.Info("[nats] connected to " + url)

	BackendAPI()

}
func BackendAPI() {
	NatsConnection.Subscribe("comfy.request", ComfyRequest)
	NatsConnection.Subscribe("textgen.request", TextGenRequest)
	NatsConnection.Subscribe("requestModel.image", requestImageModel)
	NatsConnection.Subscribe("requestModel.text", requestTextMode)
	NatsConnection.Subscribe("registerMachine", registerMachine)

}
func requestTextMode(m *nats.Msg) {
	var node models.Node
	err := json.Unmarshal(m.Data, &node)
	if err != nil {
		_returnErr(m, err, "error decoding message", true)
		return
	}
	log.Info("sending model to run for node", node)
	model := "TheBloke_Phind-CodeLlama-34B-v2-GPTQ"
	resp := modelResponse{Model: model, Type: "text"}
	b, err := json.Marshal(resp)
	if err != nil {
		_returnErr(m, err, "unknown error", true)
		return
	}
	m.Respond(b)
}

func requestImageModel(m *nats.Msg) {
	var node models.Node
	err := json.Unmarshal(m.Data, &node)
	if err != nil {
		_returnErr(m, err, "error decoding message", true)
		return
	}
	log.Info("sending model to run for node", node)
	resp := modelResponse{Model: "sdxl", Type: "image"}
	b, err := json.Marshal(resp)
	if err != nil {
		_returnErr(m, err, "unknown error", true)
		return
	}
	m.Respond(b)
}

func addOwner(machine models.Machine) []models.Node {
	nodes := machine.Nodes
	var nn []models.Node
	for i := 0; i < len(nodes); i++ {
		n := nodes[i]
		n.OwnerID = machine.OwnerID
		n.MachineID = machine.MachineID
		nn = append(nn, n)
	}
	return nn

}
func registerMachine(m *nats.Msg) {
	var machine models.Machine
	err := json.Unmarshal(m.Data, &machine)
	if err != nil {
		_returnErr(m, err, "error decoding json", true)
		return
	}
	nodesAppeneded := addOwner(machine)
	machine.Nodes = nodesAppeneded
	err = redis.RegisterMachine(machine)
	if err != nil {
		_returnErr(m, err, "error registering machine", true)
		return
	}
	m.Respond([]byte("OK"))
}

func ComfyRequest(m *nats.Msg) {
	var job models.Job
	err := json.Unmarshal(m.Data, &job)
	if err != nil {
		log.Error(err)
		_returnErr(m, err, "invalid payload", false)
		return
	}
	log.Infof("received comfy request %v ...", job)
	nodeID, err := redis.QueueJob(job)

	if err != nil {
		log.Error(err)
		_returnErr(m, err, "error queueing job", false)
		return
	}
	log.Info(nodeID)
	// c := fmt.Sprintf("")
	// err := NatsConnection.Request(c,)
}
func TextGenRequest(m *nats.Msg) {
	var job models.Job
	err := json.Unmarshal(m.Data, &job)
	if err != nil {
		log.Error(err)
		_returnErr(m, err, "invalid payload", false)
		return
	}
	log.Infof("received text request %v ...", job)
	nodeID, err := redis.QueueJob(job)

	if err != nil {
		log.Error(err)
		_returnErr(m, err, "error queueing job", false)
		return
	}
	c := fmt.Sprintf("node.textgenrequest.%s", nodeID)
	log.Info(nodeID)
	log.Info(c)
	reply, err := NatsConnection.Request(c, m.Data, time.Second*4)
	if err != nil {
		log.Error(err)
		_returnErr(m, err, "error queueing job", false)
		return
	}
	m.Respond(reply.Data)
	// log.Debug(string(reply.Data))

}
