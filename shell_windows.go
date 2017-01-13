package main;

import (
	"os/exec"
)

func execcmd(name string, arg string) error{
	cmd := exec.Command("cmd", "/c", name, arg);
	return cmd.Start();
}
