package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"syscall"
)

func main() {
	args := os.Args
	n, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}

	for i := 0; i != n; i++ {
		c := exec.Command(args[2], args[3:]...)
		c.Stdout = os.Stdout
		c.Stdin = os.Stdin
		c.Stderr = os.Stderr

		RunAndExitOnFail(c)
	}
}

func RunAndExitOnFail(c *exec.Cmd) {
	err := c.Run()
	if err == nil {
		return
	}

	if e2, ok := err.(*exec.ExitError); ok {
		if s, ok := e2.Sys().(syscall.WaitStatus); ok {
			if s.ExitStatus() != 0 {
				os.Exit(s.ExitStatus())
			}
		} else {
			panic(errors.New("Unimplemented for system where exec.ExitError.Sys() is not syscall.WaitStatus."))
		}
	}
}
