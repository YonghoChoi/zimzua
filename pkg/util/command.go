package util

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

const timeout = 3600 * time.Second // 기본 1시간

func RunCmd(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stdout
	cmd.Start()

	done := make(chan error)
	defer close(done)
	go func() { done <- cmd.Wait() }()

	select {
	case <-time.After(timeout):
		cmd.Process.Kill()
		msg := TransformToKorean(stdout.String())
		return msg, fmt.Errorf("command timed out(%ds)", int(timeout.Seconds()))
	case err := <-done:
		msg := TransformToKorean(stdout.String())
		if err != nil {
			return msg, fmt.Errorf("output : %s, error : %s", msg, err.Error())
		}
		return strings.TrimSpace(msg), nil
	}
}

func RunCmdAsync(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	if err := cmd.Start(); err != nil {
		return err
	}

	return nil
}

func RunPowershell(args string) (string, error) {
	return RunCmd("powershell", args)
}

func RunPowershellAsync(args string) error {
	return RunCmdAsync("powershell", args)
}

func ExecChown(path string, owner string) error {
	if path == "" || owner == "" {
		return fmt.Errorf("invalid Argument. path : %s, owner : %s", path, owner)
	}

	cmd := "chown -R " + owner + " " + path
	RunCmd("bash", "-c", cmd)
	return nil
}
