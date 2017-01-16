//+build !windows

package main;

import (
	"os/exec"
	"strconv"
)

func makecmd(name string, arg string) *exec.Cmd{
	return exec.Command("sh", "-c", "./" + strconv.Quote(name) + " " + strconv.Quote(arg));
}
