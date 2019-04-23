package cmd

import(
    "github.com/victoryang/api-gateway/config"
    "github.com/victoryang/api-gateway/db"
    //Log "github.com/victoryang/api-gateway/log"
)

func SetUpDatabase(c *config.Database) {
    db.SetupDB(c.FileName)
}