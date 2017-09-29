package main

import (
	"io"
	"os"
	"os/exec"
)

func main() {
	r, w := io.Pipe()
	cmd1 := exec.Command("cat", "/etc/passwd")
	cmd2 := exec.Command("grep", "root")
	cmd1.Stdin = os.Stdin
	cmd1.Stdout = w
	cmd2.Stdin = r
	cmd2.Stdout = os.Stdout

	cmd1.Start()
	cmd2.Start()

	cmd1.Wait()
}
