package main

import("os"
	"fmt"
	"unsafe"
	"strings"
	"strconv"
	"flag"
	"hkey"
	"os/exec")

var installdir string

//check os architecture is 64bit or 32bit. If is 64bit then return 1,else return 0
func get_os() int {
	var x uintptr
	os := unsafe.Sizeof(x)
	if os==8 {
		return 1
	}else{
		return 0
	}
}

func set_env() {
	new_path:=installdir+"\\android-sdk-windows;"+installdir+"\\android-ndk-r10c;"+installdir+"\\apache-ant-1.8.2\\bin;"+installdir+"\\android-sdk-windows\\platform-tools;"+installdir+"\\android-sdk-windows\\tools;"+installdir+"\\jdk1.6.0_45\\bin;"+os.Getenv("PATH")
	os.Setenv("NDK_ROOT",installdir+"\\android-ndk-r10c")
	os.Setenv("NDKROOT",installdir+"\\android-ndk-r10c")
	os.Setenv("NVPACK_ROOT",installdir)
	os.Setenv("ANT_HOME",installdir+"\\apache-ant-1.8.2")
	os.Setenv("ANDROID_HOME",installdir+"\\android-sdk-windows")
	os.Setenv("NVPACK_NDK_VERSION","android-ndk-r10c")
	os.Setenv("PATH",new_path)
	os.Setenv("JAVA_HOME",installdir+"\\jdk1.6.0_45")
	os.Setenv("CYGWIN_HOME",installdir+"\\cygwin")	
}

// deploy and install the samples' apk to device
func deploy() {
	//os.Chdir(args[1])
	//curr_dir,_ := os.Getwd()
    //fmt.Println(curr_dir)
    out,_ :=exec.Command("where","-R",".","*-debug.apk").Output()
    //fmt.Printf("%s,error=%s\n",out,erro)

    apks := strings.Fields(string(out[:]))
	//fmt.Println(apks)
    for i:=0;i<len(apks);i++ {
        deploy := exec.Command("adb","install",apks[i])
		Redirector(deploy)
		if err := deploy.Run(); err != nil {
        }

    }
    exec.Command("adb","kill-server")
    Redirector(exec.Command("taskkill","/F","/IM","adb.exe"))
	fmt.Println("##########finish install apk")
}

// check whether file exists
func Exists(name string) bool {
	if _,err := os.Stat(name); err != nil {
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

	tdksamples := installdir + "/Samples/TDK_Samples/android_samples.sln"
	gwsamples := installdir + "/Samples/GameWorks_Samples/samples/build/vs2010android/AllSamples.sln"


	var wow string
	wow = ""
	if get_os()==1 {
		wow = "\\WOW6432Node"
	}else{
		wow = ""
	}
	reg:="Software"+wow+"\\Microsoft\\VisualStudio\\"
	//fmt.Println(reg)
	vs :=[]float64{10.0,11.0,12.0}
	//var devenvcmd string
	for i:=0;i<3;i++ {
		//fmt.Printf("VisualStudio Vesrion is :%s\n",vs[i])
		vspath := reg + strconv.FormatFloat(vs[i],'f',1,64)
		installdir := hkey.Get_Hkey(vspath)
		verifypath := installdir + "devenv.com"
		//fmt.Println(installdir)
		//fmt.Println(verifypath)
		if Exists(verifypath){
			//installdir := hkey.Get_Hkey(verifypath)
				//fmt.Println(installdir)
				os.Chdir(installdir)
				break
		}	 
	}
	if Exists(gwsamples) {
	    fmt.Println("############## Compiling GameWorks_Samples ##############")
        os.Chdir(gwsamples)
		//fmt.Println(gwsamples)	
		cmd := exec.Command("devenv.com","/rebuild","debug",gwsamples)
		Redirector(cmd)

	}
	if Exists(tdksamples) {
		fmt.Println("############## Compiling TDK_Samples ##############")
        os.Chdir(tdksamples)
		//fmt.Println(tdksamples)	
		cmd := exec.Command("devenv.com","/rebuild","debug",tdksamples)
		Redirector(cmd)
	}
}

type CompileCommand struct {
}

func (*CompileCommand) Run(args ...string) {

	fset := flag.NewFlagSet(args[0], flag.ExitOnError)
	err := fset.Parse(args[1:])
	if nil == err {
		if len(args) < 3 {
			panic("Please supply workdirectory and deploy/compile!")
		}
		installdir = args[1]
		root := args[1]+"/Samples"
		os.Chdir(root)
    	set_env()
    	//fmt.Println(os.Getenv("Path"))
		if args[2]=="deploy" {
			deploy()
		}else if args[2]=="compile" {
			compile()
		}else {
			fmt.Println("Erro parameter!")
		}
	} else {
		fmt.Println(err)
	}
}