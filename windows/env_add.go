package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

type EnvaddCommand struct {
}

func pre_action() {
	cmd1 := exec.Command("adb", "wait-for-device")
	cmd2 := exec.Command("adb", "shell svc power stayon true")
	cmd3 := exec.Command("adb", "root")
	cmd4 := exec.Command("adb", "wait-for-devices")
	Redirector(cmd1)
	Redirector(cmd2)
	Redirector(cmd3)
	Redirector(cmd4)
}
func set_adb_env() {
	curr_path, _ := os.Getwd()
	path := os.Getenv("PATH")
	new_path := path + ";" + curr_path + "\\android-sdk-windows\\platform-tools"
	os.Setenv("PATH", new_path)
}
func (*EnvaddCommand) Run(args ...string) {
	fset := flag.NewFlagSet(args[0], flag.ExitOnError)
	err := fset.Parse(args[1:])
	if nil == err {
		if len(args) < 2 {
			panic("Please supply excute script file!")
		}
		set_adb_env()
		dir := filepath.Dir(args[0])
		os.Chdir(dir)

		cmd := exec.Command(args[1])
		Redirector(cmd)
		fmt.Println("end")

	} else {
		fmt.Println(err)
	}
}
