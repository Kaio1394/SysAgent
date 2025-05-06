package utils

import (
	"os/exec"
	"runtime"
)

func ExecuteScript(command string) (string, error) {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", command)
	} else {
		cmd = exec.Command("sh", "-c", command)
	}

	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}
