#include <iostream>
#include "IXingAPI.h"
#include "IXingAPI.hpp"
#include "_cgo_export.h"

IXingAPI g_iXingAPI;


void CheckAccountFunctions() {
    if (g_iXingAPI.Init("C:\\eBest\\xingAPI") == FALSE) {
        char msg0[] = "Init() Failed.";
        Check(msg0, -1);
        return;
    }

    char msg00[] = "Init() Success.";
    Check(msg00, -1);

    //std::cout<<"CheckAccountFunctions()\n";
    char msg1[] = "CheckAccountFunctions";
    Check(msg1, -1);

    char szAccount[20];
    char retStr[41];

    int nCount = g_iXingAPI.GetAccountListCount();
    //std::cout<<"Got Account Quantity : "<<nCount<<"\n";
    char msg2[] = "Got Account Quantity";
    Check(msg2, nCount);

    for (int i=0; i<nCount; i++) {
        //std::cout<<i<<"\n";
        char msg3[] = "i :";
        Check(msg3, i);

        for (int k=0; k<sizeof(szAccount); k++) { szAccount[k] = ' '; }
        for (int k=0; k<sizeof(retStr); k++) { retStr[k] = ' '; }

        if (g_iXingAPI.GetAccountList(i, szAccount, sizeof(szAccount)) == FALSE) {
            //std::cout<<"Failed to find account no at index : "<<i<<"\n";
            char msg4[] ="Failed to find account no at index";
            Check(msg4, i);
            return;
        }

        //std::cout<<szAccount<<"\n";
        Check(szAccount, i);

        if (g_iXingAPI.GetAcctDetailName(szAccount, retStr, sizeof(retStr)) == FALSE) {
            //std::cout<<"Failed to find account detail name at index : "<<i<<"\n";
            char msg5[] ="Failed to find account detail name at index";
            Check(msg5, i);
            return;
        }

        //std::cout<<retStr<<"\n";
        Check(retStr, i);
    }
}

void *GetSafeHandle() {
    if (g_iXingAPI.Init("C:\\eBest\\xingAPI") == FALSE) {
        char msg0[] = "Init() Failed.";
        Check(msg0, -1);
        return 0;
    }

    return (void *)(g_iXingAPI.GetSafeHandle());
}

//void SetSafeHandle(void *handle) {
//    g_iXingAPI.SetSafeHandle(HMODULE(handle));
//}
