package cmd

import (
    "github.com/victoryang/distributed_tasks_framework/config"
    Log "github.com/victoryang/distributed_tasks_framework/log"
)

func ConfigServerLog(cfg *config.GlobalConfiguration) error {
    dir := path.Dir(cfg.ServerLogsFile)
    if _, err := os.Stat(dir); os.IsNotExist(err) {
        os.MkdirAll(dir, os.ModeDir|os.ModePerm)
    }
    if err := Log.OpenFile(cfg.ServerLogsFile); err!=nil {
        Log.Error("Configure server log error: ", err)
        return err
    }

    Log.SetOwnFormatter("text")
    return nil
}