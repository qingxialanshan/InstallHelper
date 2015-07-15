#!/bin/bash

if [ $# -lt 2 ];then
	echo "You MUST input two parameters"
	cat<<HELP
Usage: build_script [goroot] [platform]
	   goroot : The path for go installed
	   platform: 
			all -- build for macosx and linux
			linux -- only build for linux
			macosx -- only build for macosx
HELP
	exit 0
fi

platform=$2

goroot=$1
if [ -f $goroot ];
then
	echo "File exists"
else
	mkdir $goroot
	echo "$goroot exists"
fi

function gen_go {
	# generate the go bin for cross build
	curr_dir=`pwd`
	cd $goroot
	tar xvf $curr_dir/../installer/go_src/go1.3.src.tar.gz 
	cd go/src

	CGO_ENABLED=0 GOOS=linux GOARCH=386 ./make.bash     -----linux-x86
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 ./make.bash   -----linux-x64
	CGO_ENABLED=0 GOOS=windows GOARCH=386 ./make.bash   -----windows-x86
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 ./make.bash -----windows-x64
	CGO_ENABLE=0 GOOS=darwin GOARCH=amd64 ./make.bash   -----macosx-x64
	
	cd $curr_dir
}

gen_go
export GOROOT=$goroot/go
export PATH=$goroot/go/bin:$PATH
#echo $PATH
go tool dist install cmd/8l
go tool dist install cmd/8a
go tool dist install cmd/8c
go tool dist install cmd/8g
if [ $platform == "all" ] ;then
    cd ../linux
    CGO_ENABLE=0 GOOS=linux GOARCH=amd64 $goroot/go/bin/go build InstallHelper.go env_add.go compile.go kill-process.go
    mv InstallHelper bin/linux-x64/
    CGO_ENABLE=0 GOOS=linux GOARCH=386 $goroot/go/bin/go build InstallHelper.go env_add.go compile.go kill-process.go
    mv InstallHelper bin/linux-x86/
	echo "build for linux"
    cd ../macosx
    echo `pwd`
    echo "build for macosx"
    CGO_ENABLE=0 GOOS=darwin GOARCH=amd64 $goroot/go/bin/go build InstallHelper.go env_add.go compile.go kill-process.go
    mv InstallHelper bin/
elif [ $platform == 'macosx' ];then
	cd ../macosx
	echo `pwd`
	echo "build for macosx"
    CGO_ENABLE=0 GOOS=darwin GOARCH=amd64 $goroot/go/bin/go build InstallHelper.go env_add.go compile.go kill-process.go
	mv InstallHelper bin/
elif [ $platform == 'linux' ];then
    cd ../linux
    CGO_ENABLE=0 GOOS=linux GOARCH=amd64 $goroot/go/bin/go build InstallHelper.go env_add.go compile.go kill-process.go
    mv InstallHelper bin/linux-x64/
    CGO_ENABLE=0 GOOS=linux GOARCH=386 $goroot/go/bin/go build InstallHelper.go env_add.go compile.go kill-process.go
    mv InstallHelper bin/linux-x86/
    echo "build for linux"
elif [ $platform == 'windows' ];then
	echo "please use the build_script.bat to build for windows"
else
	echo "wrong platform name"
fi

source ~/.bashrc
