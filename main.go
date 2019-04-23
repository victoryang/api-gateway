package main

import (
    "github.com/victoryang/distributed_tasks_framework/cmd"
)

func main() {
    cmd.NewRootCmd().Execute()
}