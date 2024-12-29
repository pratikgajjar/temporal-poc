package main

import (
	"context"
	"github.com/google/uuid"
	flow "github.com/pratikgajjar/tempoc/workflow"
	"go.temporal.io/sdk/client"
	"log"
)

func main() {
	// The client is a heavyweight object that should be created once per process.
	c, err := client.Dial(client.Options{
		HostPort: client.DefaultHostPort,
	})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	// This workflow ID can be user business logic identifier as well.

	var we client.WorkflowRun
	var workflowID string
	workflowOptions := client.StartWorkflowOptions{
		ID:        workflowID,
		TaskQueue: "cron",
		// CronSchedule: "5 * * * *",
	}

	for _ = range 1000 {
		workflowID = "cron_" + uuid.New().String()
		workflowOptions.ID = workflowID
		we, err = c.ExecuteWorkflow(context.Background(), workflowOptions, flow.SampleCronWorkflow)
		if err != nil {
			log.Fatalln("Unable to execute workflow", err)
		}
		log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())
	}
}
