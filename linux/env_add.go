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
	cmd := exec.Command("bash", "-c", `adb wait-for-device && adb shell svc power stayon true && adb root && adb wait-for-devices`)
	Redirector(cmd)
}

func set_adb_env() {
	curr_path, _ := os.Getwd()
	path := os.Getenv("PATH")
	new_path := path + ":" + curr_path + "/android-sdk-linux/platform-tools/"
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
		fmt.Println(os.Getenv("PATH"))

		pre_action()

		cmd := exec.Command("sh", args[1])
		Redirector(cmd)
		fmt.Println("end")

	} else {
		fmt.Println(err)
	}
}
