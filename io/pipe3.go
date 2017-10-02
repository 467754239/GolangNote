package main

import (
	"os"
	"os/exec"
)

func main() {
	cmd1 := exec.Command("cat", "/etc/passwd")
	cmd2 := exec.Command("grep", "root")
	cmd3 := exec.Command("grep", "/bin/bash")

	cmd2.Stdin, _ = cmd1.StdoutPipe()
	cmd3.Stdin, _ = cmd2.StdoutPipe()

	cmd3.Stdout = os.Stdout
	cmd3.Stderr = os.Stderr

	cmd3.Start()
	cmd2.Start()
	cmd1.Run()

	cmd3.Wait()
}
