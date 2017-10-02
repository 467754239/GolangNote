package main

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os/exec"
	"syscall"
)

type CommandResult struct {
	Stdout   string
	Stderr   string
	ExitCode int
}

func NewError(v ...interface{}) error {
	logMsg := fmt.Sprint(v...)
	return errors.New(logMsg)
}

func NewErrorf(format string, v ...interface{}) error {
	logMsg := fmt.Sprintf(format, v...)
	return errors.New(logMsg)
}

func Execute(cmdName string, args ...string) (*CommandResult, error) {
	var err error

	cmd := exec.Command(cmdName, args...)
	if cmd.Stdout != nil {
		return nil, NewError("exec: Stdout already set")
	}
	if cmd.Stderr != nil {
		return nil, NewError("exec: Stderr already set")
	}

	result := &CommandResult{ExitCode: 0}
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err = cmd.Run()
	result.Stdout = string(stdout.Bytes())
	result.Stderr = string(stderr.Bytes())

	// exit code not 0
	if exiterr, ok := err.(*exec.ExitError); ok {
		if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
			// get exit code
			result.ExitCode = status.ExitStatus()
		}
	} else if err != nil {
		return nil, err
	}

	return result, nil
}

func main() {
	result, err := Execute("df", "-h", "|", "grep tmp")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.Stdout)
}
