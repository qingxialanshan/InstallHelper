// Created by cgo - DO NOT EDIT

//line C:\Users\amyl\Perforce\amyl_AMYL-LT_8950\sw\devtools\TADP\External_Scripts\windows\src\get_uninstallstr\get_uninstallstr.go:1
package get_uninstallstr
//line C:\Users\amyl\Perforce\amyl_AMYL-LT_8950\sw\devtools\TADP\External_Scripts\windows\src\get_uninstallstr\get_uninstallstr.go:8

//line C:\Users\amyl\Perforce\amyl_AMYL-LT_8950\sw\devtools\TADP\External_Scripts\windows\src\get_uninstallstr\get_uninstallstr.go:7
func Get_Uninstallstr(regname string, ins_name string) string {
																	rname := _Cfunc_CString(regname)
																	iname := _Cfunc_CString(ins_name)
																	uninstall_str := _Cfunc_GoString(_Cfunc_Get_Uninstallstr(rname, iname))
//line C:\Users\amyl\Perforce\amyl_AMYL-LT_8950\sw\devtools\TADP\External_Scripts\windows\src\get_uninstallstr\get_uninstallstr.go:13

//line C:\Users\amyl\Perforce\amyl_AMYL-LT_8950\sw\devtools\TADP\External_Scripts\windows\src\get_uninstallstr\get_uninstallstr.go:12
	return uninstall_str
}
