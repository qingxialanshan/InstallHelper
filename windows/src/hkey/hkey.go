package hkey

//#include "hkey.h"
import "C"

func Get_Hkey(regname string)string{
	rname := C.CString(regname)
	kvalue := C.GoString(C.get_hkey(rname))
	
	return kvalue
}
