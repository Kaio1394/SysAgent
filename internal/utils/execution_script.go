package utils

import (
	"fmt"
	"github.com/google/uuid"
	"os/exec"
	"runtime"
	"sync"
)

var (
	statusMap = make(map[string]string)
	outputMap = make(map[string]string)
	mutex     = sync.RWMutex{}
)

func ExecuteScript(command string) (string, error) {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", command)
	} else {
		cmd = exec.Command("sh", "-c", command)
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func ExecuteScriptSync(command string) string {
	taskID := uuid.New().String()

	mutex.Lock()
	statusMap[taskID] = "iniciado"
	outputMap[taskID] = ""
	mutex.Unlock()

	go func(id string) {
		mutex.Lock()
		statusMap[id] = "in progress"
		mutex.Unlock()

		var cmd *exec.Cmd
		if runtime.GOOS == "windows" {
			cmd = exec.Command("cmd", "/C", command)
		} else {
			cmd = exec.Command("sh", "-c", command)
		}

		output, err := cmd.CombinedOutput()

		mutex.Lock()
		if err != nil {
			statusMap[id] = "error"
			outputMap[id] = fmt.Sprintf("error: %v\noutput:\n%s", err, string(output))
		} else {
			statusMap[id] = "done"
			outputMap[id] = string(output)
		}
		mutex.Unlock()
	}(taskID)

	return taskID
}
func GetResult(taskID string) (string, string) {
	mutex.RLock()
	status, statusExists := statusMap[taskID]
	output, outputExists := outputMap[taskID]
	mutex.RUnlock()

	if !statusExists || !outputExists {
		return "", ""
	}
	return status, output
}
