package cmd

import(
    "github.com/victoryang/api-gateway/config"
    Log "github.com/victoryang/api-gateway/log"
)

var configFile string

func LoadConfig() *config.GlobalConfiguration{
    cfg, err := config.LoadFile(configFile) 
    if err!=nil {
        Log.Error("Parse configure file error: ", err)
        Log.Error("set to default configuration")
        cfg = config.DefaultGlobalConfiguration
    }

    return cfg
}