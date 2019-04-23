package cmd

import (
    "os"
    "os/signal"
    "syscall"

    "github.com/victoryang/api-gateway/api"
    "github.com/victoryang/api-gateway/config"
    Log "github.com/victoryang/api-gateway/log"

    "github.com/spf13/cobra"
)

func runDaemon(cmd *cobra.Command, args []string) error {
    cfg := LoadConfig()

    if err := ConfigServerLog(cfg); err!=nil {
        return returnError(ERR_OPEN_LOG_FILE_FAIL)
    }

    startAdminServer(cfg.Admin)

    SetUpDatabase(cfg.Databases)
}