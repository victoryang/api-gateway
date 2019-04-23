package cmd

import (
    "github.com/victoryang/api-gateway/config"
    Log "github.com/victoryang/api-gateway/log"
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