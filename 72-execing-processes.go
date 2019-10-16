package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	binary, lookErr := exec.LookPath("ls")
	fmt.Println("ls:", binary)
	if lookErr != nil {
		panic(lookErr)
	}
	args := []string{"ls", "-a", "-l", "-h"}
	env := os.Environ()
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
