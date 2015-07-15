// Created by cgo - DO NOT EDIT

//line C:\Users\amyl\Desktop\golang\windows\test\test.go:1
package main
//line C:\Users\amyl\Desktop\golang\windows\test\test.go:4

//line C:\Users\amyl\Desktop\golang\windows\test\test.go:3
import "hkey"
//line C:\Users\amyl\Desktop\golang\windows\test\test.go:7

//line C:\Users\amyl\Desktop\golang\windows\test\test.go:6
func main() {
	reg := "Software\\WOW6432Node\\Microsoft\\VisualStudio\\10.0"
	key := ""
	hkey.Get_Hkey(key, reg)
}
