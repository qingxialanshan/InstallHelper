package main

import (
	"flag"
	"fmt"
	"os/exec"
	"strings"
)

type KillprocessCommand struct {
}

func kill() {
	cmd := exec.Command("/bin/sh", "-c", `ps -e|grep "DownloadHelper"|awk '{print $1}'`)
	pid, err := cmd.Output()
	if err != nil {
		fmt.Println("failed")
	}

	pids := strings.Fields(string(pid[:]))

	if len(pids) > 0 {
		for i := 0; i < len(pids); i++ {
			exec.Command("kill", pids[i]).Start()
			fmt.Println("Success")
		}
	} else {
		fmt.Println("cannot find DowloadHelper")
	}
}

func (*KillprocessCommand) Run(args ...string) {
	fset := flag.NewFlagSet(args[0], flag.ExitOnError)
	err := fset.Parse(args[1:])
	if nil == err {
		kill()

	} else {
		fmt.Println(err)
	}
}
