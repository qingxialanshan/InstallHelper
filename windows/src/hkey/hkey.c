#include <stdio.h>
#include <Windows.h>
#include "hkey.h"

char * get_hkey(const char * regname)
{  
 HKEY hKey;
 
 char *sqzpath=malloc(sizeof(char) * 500);
 DWORD dwSize = sizeof(char)*500;

 if (RegOpenKey(HKEY_LOCAL_MACHINE,regname,&hKey) != ERROR_SUCCESS)
 {
 	printf ("Error: Regedit cannot be opened!");
 	return NULL;
 }
 else
 {
 	if (RegQueryValueEx(hKey,"InstallDir",NULL,NULL,(LPBYTE)sqzpath,&dwSize)==ERROR_SUCCESS)
 	{
 		return sqzpath;
 	}//get the value of register
 	
 }
 free(sqzpath);
 sqzpath=NULL;
 return NULL;
}

