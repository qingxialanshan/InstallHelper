package main

import (
	"C"
	"flag"
	"fmt"
	"get_uninstallstr"
	"os/exec"
)

type UninstallCommand struct {
}

func (*UninstallCommand) Run(args ...string) {
	fset := flag.NewFlagSet(args[0], flag.ExitOnError)
	err := fset.Parse(args[1:])
	if nil == err {
		if len(args) < 2 {
			fmt.Println("Need to add installer's name, eg pentak/battle/quadd")
			return
		}
		var uni_name string = args[1]
		var uninstall_str string
		var reg string = "SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Installer\\UserData\\"
		if uni_name == "quadd" {
			uninstall_str = get_uninstallstr.Get_Uninstallstr(reg, "Tegra System")

		} else if uni_name == "pentak" {
			uninstall_str = get_uninstallstr.Get_Uninstallstr(reg, "Nsight Tegra")
		} else if uni_name == "battle" {
			uninstall_str = get_uninstallstr.Get_Uninstallstr(reg, "NVIDIA Tegra Graphics Debugger")
		}else {
			fmt.Println("Wrong installer name.")
		}
		//fmt.Println(uninstall_str)
		if uninstall_str == "cannot find" {
			fmt.Println("Not find the " + uni_name)
			return
		}

		uninstall_str = uninstall_str[14:]
		uninstall_cmd := exec.Command("msiexec", "/uninstall", uninstall_str, "/passive")
		Redirector(uninstall_cmd)

	} else {
		fmt.Println(err)
	}
}
