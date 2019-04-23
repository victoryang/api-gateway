package cmd

import (
    "fmt"
)

const (
    ERR_NONE = iota
    ERR_OPEN_LOG_FILE_FAIL
    ERR_START_APISERVER
)

func returnError(errno int) error {
    return fmt.Errorf("Fail to start server due to: -" + string(errno))
}