package main

import (
	"backend/core"
	"backend/db"
	"backend/db/redis"
	"backend/fiber"
	"backend/nats"
)

var (
	log = core.GetLogger()
)

func main() {
	log.Info("getting comfy config..")
	defer core.HandlePanic()
	log.Info("Connecting to Mongo")
	db.Connect()
	log.Info("Connecting to Redis")
	redis.Start()
	redis.WaitActive()
	log.Info("starting nats server..")
	nats.Start()
	log.Info("nats setup done")
	log.Info("starting server")
	fiber.Start()
}
