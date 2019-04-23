package cmd

import(
    "github.com/victoryang/distributed_tasks_framework/config"
    "github.com/victoryang/distributed_tasks_framework/db"
    //Log "github.com/victoryang/distributed_tasks_framework/log"
)

func SetUpDatabase(c *config.Database) {
    db.SetupDB(c.FileName)
}