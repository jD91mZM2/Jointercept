package main;

import (
	"os/exec"
)

func makecmd(name string, arg string) *exec.Cmd{
	return exec.Command("cmd", "/c", name, arg);
}
