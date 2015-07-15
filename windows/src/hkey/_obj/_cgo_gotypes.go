// Created by cgo - DO NOT EDIT

package hkey

import "unsafe"

import "syscall"

import _ "runtime/cgo"

type _ unsafe.Pointer

func _Cerrno(dst *error, x int32) { *dst = syscall.Errno(x) }
type _Ctype_char int8

type _Ctype_void [0]byte

func _Cfunc_CString(string) *_Ctype_char
func _Cfunc_GoString(*_Ctype_char) string
func _Cfunc_get_hkey(*_Ctype_char) *_Ctype_char
