package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type CompileCommand struct {
}

var installdir string

func set_env() {
	new_path := os.Getenv("PATH") + ":" + installdir + "/android-sdk-macosx:" + installdir + "/android-ndk-r10c:" + installdir + "/apache-ant-1.8.2/bin:" + installdir + "/android-sdk-macosx/platform-tools:" + installdir + "/android-sdk-macosx/tools"
	os.Setenv("NDK_ROOT", installdir+"/android-ndk-r10c")
	os.Setenv("NDKROOT", installdir+"/android-ndk-r10c")
	os.Setenv("NVPACK_ROOT", installdir)
	os.Setenv("ANT_HOME", installdir+"/apache-ant-1.8.2")
	os.Setenv("ANDROID_HOME", installdir+"/android-sdk-macosx")
	os.Setenv("NVPACK_NDK_VERSION", "android-ndk-r10c")
	os.Setenv("PATH", new_path)
}

// deploy and install the samples' apk to device
func deploy() {
	out, _ := exec.Command("bash", "-c", `find -L . -iname "*-debug.apk"`).Output()

	apks := strings.Fields(string(out[:]))

	for i := 0; i < len(apks); i++ {
		deploy := exec.Command("adb", "install", apks[i])
		Redirector(deploy)
		if err := deploy.Run(); err != nil {
		}
		if strings.Contains(filepath.Dir(apks[i]), "oclConvolution") {
			os.Chdir(filepath.Dir(apks[i]))
			copy := exec.Command("bash", "-c", `. ../copy_assets.sh`)
			Redirector(copy)
		}
	}
	fmt.Println("##########finish install apk")
}

// check whether file exists
func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func Redirector(cmd *exec.Cmd) {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

// compile samples
func compile() {
	tdksamples := installdir + "/Samples/TDK_Samples"
	cudasamples := installdir + "/Samples/CUDA_Samples"
	oclsamples := installdir + "/Samples/OpenCL_Samples/oclConvolutionSeparable"
	gwsamples := installdir + "/Samples/GameWorks_Samples/samples/build/makeandroid"

	if Exists(tdksamples) {
		fmt.Println("######################## Compiling TDK_Samples ########################")
		os.Chdir(tdksamples)
		out, _ := exec.Command("bash", "-c", `find . -iname "*.vcxproj"|grep -v "libs"`).Output()
		tdk_applist := strings.Fields(string(out[:]))
		for i := 0; i < len(tdk_applist); i++ {
			tdk_appDir := filepath.Dir(tdk_applist[i])
			os.Chdir(tdk_appDir)
			fmt.Println(tdk_appDir)
			cmd1 := exec.Command("bash", "-c", `android update project -p . --target android-15 && ndk-build -C jni`)
			Redirector(cmd1)

			cmd2 := exec.Command("bash", "-c", `ant debug`)
			Redirector(cmd2)
			os.Chdir(tdksamples)
		}
	}

	if Exists(cudasamples) {
		fmt.Println("######################## Compiling CUDA_Samples ########################")
		os.Chdir(cudasamples)
		out, _ := exec.Command("bash", "-c", `ls -d ./*`).Output()
		cuda_applist := strings.Fields(string(out[:]))
		fmt.Println(cuda_applist)
		for i := 0; i < len(cuda_applist); i++ {
			os.Chdir(cuda_applist[i])
			os.Chdir("cuda")
			make := exec.Command("bash", "-c", `make`)
			Redirector(make)
			os.Chdir(cudasamples)
			os.Chdir(cuda_applist[i])
			cmd1 := exec.Command("bash", "-c", `android update project -p . --target android-15 && ndk-build -C jni`)
			cmd2 := exec.Command("bash", "-c", `ant debug`)
			Redirector(cmd1)
			Redirector(cmd2)
			os.Chdir(cudasamples)
		}
	}

	if Exists(oclsamples) {
		fmt.Println("######################## Compiling OCL_Samples ########################")
		os.Chdir(oclsamples)
		os.Chdir("opencl")
		make := exec.Command("bash", "-c", `make`)
		Redirector(make)
		os.Chdir(oclsamples)
		cmd1 := exec.Command("bash", "-c", `android update project -p . --target android-15 && ndk-build -C jni`)
		cmd2 := exec.Command("bash", "-c", `ant debug`)
		Redirector(cmd1)
		Redirector(cmd2)
	}

	if Exists(gwsamples) {
		fmt.Println("######################## Compiling GameWorks_Samples ########################")
		os.Chdir(gwsamples)
		fmt.Println(gwsamples)
		cmd := exec.Command("make")
		Redirector(cmd)

		if err := cmd.Run(); err != nil {
		}
	}
}
func (*CompileCommand) Run(args ...string) {
	fset := flag.NewFlagSet(args[0], flag.ExitOnError)

	err := fset.Parse(args[1:])
	if nil == err {
		if len(args) < 3 {
			panic("Please supply workdirectory and deploy/compile command!")
		}
		installdir = args[1]
		root := args[1] + "/Samples"
		os.Chdir(root)
		set_env()
		if args[2] == "deploy" {
			deploy()
		} else if args[2] == "compile" {
			compile()
		} else {
			fmt.Println("Erro parameter!")
		}
	} else {
		fmt.Println(err)
	}
}
