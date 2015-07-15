package get_uninstallstr


//char * Get_Uninstallstr(char * regname,char * installer_name);
import "C"

func Get_Uninstallstr(regname string,ins_name string)string{
	rname := C.CString(regname)
	iname := C.CString(ins_name)
	uninstall_str := C.GoString(C.Get_Uninstallstr(rname,iname))

	return uninstall_str
}