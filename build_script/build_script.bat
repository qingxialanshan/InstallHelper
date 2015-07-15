:: set environment path GOPATH and GOROOT
set GOROOT=C:\Users\amyl\Downloads\go1.3.windows-amd64\go
set GOPATH=C:\Users\amyl\Perforce\amyl_AMYL-LT_8950\sw\devtools\TADP\External_Scripts\windows

CD ..
CD /D windows
go build InstallHelper.go env_add.go uninstall.go compile.go