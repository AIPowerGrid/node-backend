package redis

import (
	"backend/models"
	"errors"
	"fmt"

	json "github.com/goccy/go-json"
)

/*
keys:

nodes:model:zscore - nodes zscore of all nodeids paticular to a model ....

activeNodes - set of all nodes registered

nodeData:id - json of the node, containing owner information and similar info

owner:$ID:nodes
owner:$ID:machines

pendingJobs:nodeID - list of all jobs currently on a node id
globalPendingJobs: - list of all current jobs


job:key - job data containing prompt, seed, type (image or text gen) - either string or hash
*/

var (
	globalJobsKey  = "globalPendingJobs"
	activeNodesKey = "activeNodes"
)

func RegisterMachine(m models.Machine) error {
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}
	log.Debugf("Register machine called for %v", m)
	pipe := client.Pipeline()
	machineKey := fmt.Sprintf("machine:%s", m.MachineID)
	pipe.Set(ctx, machineKey, string(b), 0)
	pipe.SAdd(ctx, fmt.Sprintf("owner:%s:machines", m.OwnerID), m.MachineID)
	for _, node := range m.Nodes {
		nodeBytes, err := json.Marshal(node)
		if err != nil {
			return err
		}
		pipe.SAdd(ctx, activeNodesKey, node.ID)
		nodeKey := fmt.Sprintf("nodeData:%s", node.ID)
		pipe.Set(ctx, nodeKey, string(nodeBytes), 0)
		pipe.SAdd(ctx, fmt.Sprintf("owner:%s:nodes", m.OwnerID), node.ID)
		m := _mem(node.ID, 0)
		modelZkey := zscoreKey(node.Model)
		pipe.ZAdd(ctx, modelZkey, m)
	}
	_, err = pipe.Exec(ctx)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}
func zscoreKey(model string) string {
	return fmt.Sprintf("nodes:%s:zscore", model)

}
func AddNodeZ(nodeID, model string) {
	k := zscoreKey(model)
	m := _mem(nodeID, 0)
	client.ZAdd(ctx, k, m)
}
func QueueJob(job models.Job) (string, error) {
	/*
		first have to find node with least amount of tasks .
		then have to add that task ID to the node list, as well as add the key to the allJobs list
		we are using lists, not sets, to see jobs so they are in order
	*/
	var finalID string
	jobKey := fmt.Sprintf("job:%s", job.ID)
	zkey := zscoreKey(job.Model)
	log.Debugf("Getting model key %s", zkey)
	zresult, err := client.ZRangeWithScores(ctx, zkey, 0, 0).Result()
	if err != nil {
		log.Error(err)
		return finalID, err
	}
	if len(zresult) == 0 {
		return finalID, errors.New("no elements in sorted set")
	}
	score := int(zresult[0].Score)
	nodeID := zresult[0].Member.(string)
	job.NodeID = nodeID

	nodeJobsKey := fmt.Sprintf("pendingJobs:%s", nodeID)
	log.Debugf("Found Node to distrubute for model %s to %s with pending tasks %d", job.Model, nodeID, score)

	jobString, err := _jobString(job)
	if err != nil {
		log.Error(err)
		return finalID, err
	}
	/* we have the node, now we have to have a pipeline to add all the keys
	first operation: incr score
	second operation: add job to node jobs list
	third operation: add job key to global jobs list
	fourth operation: add job key json
	*/

	pipe := client.Pipeline()
	pipe.ZIncrBy(ctx, zkey, 1, nodeID)
	// pipe.ZIncr(ctx, zkey, &redis.Z{Member: nodeID, Score: 1})
	pipe.RPush(ctx, nodeJobsKey, job.ID)
	pipe.RPush(ctx, globalJobsKey, job.ID)
	pipe.Set(ctx, jobKey, jobString, 0)
	_, err = pipe.Exec(ctx)
	if err != nil {
		log.Error(err)
		return finalID, err
	}

	log.Infof("Queued Job %s for %s, pending count:%d", job.ID, nodeID, score)
	finalID = nodeID
	return finalID, nil

}
func _jobString(job models.Job) (string, error) {
	b, err := json.Marshal(job)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
