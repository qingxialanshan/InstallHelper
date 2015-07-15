This is a readme about how to build the cross-compile environment for go.

1. Download the src file for go from the website http://golang.org/dl/go1.3.src.tar.gz. Or get from External_Scripts/installer/go_src
2. tar xvf go1.3.src.tar.gz
3. cd go/src
4. if want to build 32bit on 64bit system you should first manually get the tools first as followings.
	a. go tool dist install cmd/8l
	b. go tool dist install cmd/8a
	c. go tool dist install cmd/8c
	d. go tool dist install cmd/8g
5. CGO_ENABLED=0 GOOS=linux GOARCH=386 ./make.bash     -----linux-x86
  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 ./make.bash   -----linux-x64
  CGO_ENABLED=0 GOOS=windows GOARCH=386 ./make.bash   -----windows-x86
  CGO_ENABLED=0 GOOS=windows GOARCH=amd64 ./make.bash -----windows-x64
  CGO_ENABLE=0 GOOS=darwin GOARCH=amd64 ./make.bash   -----macosx-x64
  CGO_ENABLE=0 GOOS=darwin GOARCH=386   ./make.bash   -----macosx-x32
6. set env
  $GOROOT

ISSUES
This cross-compile is not support CGO

For windows CGO
1. set GOPATH=External_Scripts/windows/
2. Build CGO packages
	a. hkey.a -- used for get the hkey value for the register
	   cd ${GOPATH}/src/hkey
	   go tool cgo hkey.go
	   go install
	b. get_uninstallstr.a --- get the uninstall string for the installer that defined
	   cd ${GOPATH}/src/get_uninstallstr
	   go tool cgo get_uninstallstr
	   go install	   
3. go build compile.go/env_add.go/uninstall.go
