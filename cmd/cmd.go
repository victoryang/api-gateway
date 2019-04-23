package cmd

import (
    "fmt"
    "os"
    "github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
    rootCmd := &cobra.Command {
        Use:    "distributed_tasks [option]",
        Short:  "An api server for distributed_tasks",
        Long:   `An api server for distributed_tasks written in go and node`,
        RunE:   runDaemon,
    }

    flags := rootCmd.Flags()
    flags.StringVarP(&configFile, "configure", "C", "/etc/distributed_tasks/config.yaml", "configuration file for api server")
    return rootCmd
}


func Execute(cmd *cobra.Command) {
    if err := cmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}