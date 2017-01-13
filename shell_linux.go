package main;

import (
	"os/exec"
	"strconv"
)

func execcmd(name string, arg string) error{
	cmd := exec.Command("sh", "-c", strconv.Quote(name) + " " + strconv.Quote(arg));
	return cmd.Start();
}
