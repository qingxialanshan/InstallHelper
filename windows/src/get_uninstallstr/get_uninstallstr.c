// QueryKey - Enumerates the subkeys of key and its associated values.

//     hKey - Key whose subkeys and values are to be enumerated.



#include <windows.h>

#include <stdio.h>

#include <tchar.h>

#include <string.h>





#define MAX_KEY_LENGTH 255

#define MAX_VALUE_NAME 16383

#define SUB_KEY_NUM 1000



void ListFree(char ***KeyList)

{

    char **List = *KeyList;

    int i = 0;

    for(i = 0; i < SUB_KEY_NUM; i++)

    {

        free(List[i]);

    }

    free(List);

}



DWORD QueryKey(HKEY hKey, char ***KeyList) 

{ 

    *KeyList = (char **)malloc(SUB_KEY_NUM * sizeof(char *));

    char **List = *KeyList; 

    int n = 0;

    for(n = 0; n < SUB_KEY_NUM; n++)

    {

        List[n] = (char *)malloc(MAX_KEY_LENGTH * sizeof(char));

    }

    TCHAR    achKey[MAX_KEY_LENGTH];   // buffer for subkey name

    DWORD    cbName;                   // size of name string 

    TCHAR    achClass[MAX_PATH] = TEXT("");  // buffer for class name 

    DWORD    cchClassName = MAX_PATH;  // size of class string 

    DWORD    cSubKeys=0;               // number of subkeys 

    DWORD    cbMaxSubKey;              // longest subkey size 

    DWORD    cchMaxClass;              // longest class string 

    DWORD    cValues;              // number of values for key 

    DWORD    cchMaxValue;          // longest value name 

    DWORD    cbMaxValueData;       // longest value data 

    DWORD    cbSecurityDescriptor; // size of security descriptor 

    FILETIME ftLastWriteTime;      // last write time 

 

    DWORD i, retCode; 

 

    TCHAR  achValue[MAX_VALUE_NAME]; 

    DWORD cchValue = MAX_VALUE_NAME; 

 

    // Get the class name and the value count. 

    retCode = RegQueryInfoKey(

        hKey,                    // key handle 

        achClass,                // buffer for class name 

        &cchClassName,           // size of class string 

        NULL,                    // reserved 

        &cSubKeys,               // number of subkeys 

        &cbMaxSubKey,            // longest subkey size 

        &cchMaxClass,            // longest class string 

        &cValues,                // number of values for this key 

        &cchMaxValue,            // longest value name 

        &cbMaxValueData,         // longest value data 

        &cbSecurityDescriptor,   // security descriptor 

        &ftLastWriteTime);       // last write time 

 

    // Enumerate the subkeys, until RegEnumKeyEx fails.

    

    if (cSubKeys)

    {

        //printf( "\nNumber of subkeys: %d\n", cSubKeys);



        for (i=0; i<cSubKeys; i++) 

        { 

            cbName = MAX_KEY_LENGTH;

            retCode = RegEnumKeyEx(hKey, i,

                     achKey, 

                     &cbName, 

                     NULL, 

                     NULL, 

                     NULL, 

                     &ftLastWriteTime); 

            if (retCode == ERROR_SUCCESS) 

            {

                //_tprintf(TEXT("(%d) %s\n"), i+1, achKey);

                strcpy(List[i], achKey);    

            }

            

        }

    } 

 

    // Enumerate the key values. 



    if (cValues) 

    {

        printf( "\nNumber of values: %d\n", cValues);



        for (i=0, retCode=ERROR_SUCCESS; i<cValues; i++) 

        { 

            cchValue = MAX_VALUE_NAME; 

            achValue[0] = '\0'; 

            retCode = RegEnumValue(hKey, i, 

                achValue, 

                &cchValue, 

                NULL, 

                NULL,

                NULL,

                NULL);

 

            if (retCode == ERROR_SUCCESS ) 

            { 

                _tprintf(TEXT("(%d) %s\n"), i+1, achValue); 

            } 

        }

    }
    return cSubKeys;

}



char * Get_Uninstallstr(char * regname,char * installer_name)

{

   HKEY hTestKey,tmp_subtestkey,subtestkey;

   char **List,**sub_List1;

   char *uninstall_str=(char *)malloc(500*sizeof(char));

   DWORD subkey_num,subkey_num1;

   if( RegOpenKeyEx( HKEY_LOCAL_MACHINE,

        TEXT("SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Installer\\UserData"),

        0,

        KEY_READ,

        &hTestKey) == ERROR_SUCCESS

      )

   {

      subkey_num = QueryKey(hTestKey, &List);

   }

   

   // search from the List

   int i,j;

   char *reg=(char *)malloc(500*sizeof(char));

   char *display_name=(char*)malloc(sizeof(char) * 1024);

   //TCHAR uninstall_str[500];

   DWORD dwSize = sizeof(char)*1024;

   DWORD dwtype = REG_EXPAND_SZ;



   for(i=0;i<subkey_num;i++)

   {

        memset(reg,0,500);

        //src ="SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Installer\\UserData\\";

        strcpy(reg,regname);

        strcat(reg,List[i]);

        strcat(reg,"\\Products");



        //printf("1. %s\n",reg );

        if(RegOpenKeyEx(HKEY_LOCAL_MACHINE,TEXT(reg),0,KEY_READ,&tmp_subtestkey)==ERROR_SUCCESS)

        {

            

            subkey_num1=QueryKey(tmp_subtestkey,&sub_List1);

            for (j=0;j<subkey_num1;j++)

            {

              char *reg1=(char *)malloc(500*sizeof(char));

              memset(reg1,0,500);

              strcpy(reg1,reg);

              strcat(reg1,"\\");

              strcat(reg1,sub_List1[j]);

              strcat(reg1,"\\InstallProperties");

              //printf("2. %s\t",reg1 );



              if(RegOpenKeyEx(HKEY_LOCAL_MACHINE,TEXT(reg1),0,KEY_READ,&subtestkey)==ERROR_SUCCESS)

              {


                dwSize=1024;


                RegQueryValueEx(subtestkey,"DisplayName",NULL,NULL,(LPBYTE)display_name,&dwSize);

                //printf("\t%s\n",display_name);

                if(strstr(display_name,installer_name)!=NULL)

                {

                  printf("The regid is %s\n",reg1);

                  dwSize=1024;

                  RegQueryValueExA(subtestkey,_T("UninstallString"),NULL,&dwtype,(LPBYTE)uninstall_str,&dwSize);

                  printf("The uninstall_str is %s\n",uninstall_str);

                  break;

                }

                // else if(strstr(display_name,"Nsight Tegra") != NULL)

                // {

                //   dwSize=500;

                //   printf("Nsight Tegra's regid is %s\n",reg1);

                //   RegQueryValueEx(subtestkey,"UninstallString",NULL,NULL,(LPBYTE)uninstall_str,&dwSize);

                //   printf("NT uninstall_str is %s\n",uninstall_str);

                // }

                else

                {

                  continue;

                }

              }

             

            }

            if(j==subkey_num1)

            {

              char * err_str="cannot find";

              return err_str;

            }

        }

        return uninstall_str;

        

   }

    ListFree(&List);

    ListFree(&sub_List1);

    free(display_name);

    //free(uninstall_str);



   RegCloseKey(hTestKey);

   RegCloseKey(tmp_subtestkey);

   RegCloseKey(subtestkey);



}



 int main(void)

 {

   char *src ="SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Installer\\UserData\\";

   char *pentk = "Nsight Tegra";

   char *quadd = "Tegra System Profiler";

   TCHAR uninstall_quadd[500];

   TCHAR uninstall_pentk[500];

   char *uninstall_pentk,*uninstall_quadd;

   char *uninstall_pentk=(char *)malloc(500*sizeof(char));

   char *uninstall_quadd=(char *)malloc(sizeof(char) * 500);

   memset(uninstall_quadd,0,500);

   memset(uninstall_pentk,0,500);



   uninstall_quadd=Get_Uninstallstr(src,quadd);

   if(uninstall_quadd=="")

   {

     printf("cannot find quadd\n" );

   }

   else

   {

     printf("quadd\t%s\n",uninstall_quadd);

   }



   uninstall_pentk=Get_Uninstallstr(src,pentk);

   if((uninstall_pentk)=="")

   {

     printf("cannot find pentk\n" );

   }

   else

   {

     printf("pentk\t%s\n",uninstall_pentk);

   }

   free(uninstall_quadd);

   free(uninstall_pentk);

    return 0;

 }

