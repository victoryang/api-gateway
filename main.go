package main

import (
    "github.com/victoryang/distributed_tasks_demo/cmd"
)

func main() {
    cmd.NewRootCmd().Execute()
}