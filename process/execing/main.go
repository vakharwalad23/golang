package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {

	bin, err := exec.LookPath("ls")
	if err != nil {
		panic(err)
	}

	args := []string{"ls", "-a", "-l", "-h"}

	env := os.Environ()

	execErr := syscall.Exec(bin, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
