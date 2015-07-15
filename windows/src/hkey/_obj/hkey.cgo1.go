// Created by cgo - DO NOT EDIT

//line C:\Users\amyl\Perforce\amyl_AMYL-LT_8950\sw\devtools\TADP\External_Scripts\windows\src\hkey\hkey.go:1
package hkey
//line C:\Users\amyl\Perforce\amyl_AMYL-LT_8950\sw\devtools\TADP\External_Scripts\windows\src\hkey\hkey.go:7

//line C:\Users\amyl\Perforce\amyl_AMYL-LT_8950\sw\devtools\TADP\External_Scripts\windows\src\hkey\hkey.go:6
func Get_Hkey(regname string) string {
														rname := _Cfunc_CString(regname)
														kvalue := _Cfunc_GoString(_Cfunc_get_hkey(rname))
//line C:\Users\amyl\Perforce\amyl_AMYL-LT_8950\sw\devtools\TADP\External_Scripts\windows\src\hkey\hkey.go:11

//line C:\Users\amyl\Perforce\amyl_AMYL-LT_8950\sw\devtools\TADP\External_Scripts\windows\src\hkey\hkey.go:10
	return kvalue
}
