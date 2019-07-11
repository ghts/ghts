/* Copyright (C) 2015-2019 김운하(UnHa Kim)  unha.kim.ghts@gmail.com

이 파일은 GHTS의 일부입니다.

이 프로그램은 자유 소프트웨어입니다.
소프트웨어의 피양도자는 자유 소프트웨어 재단이 공표한 GNU LGPL 2.1판
규정에 따라 프로그램을 개작하거나 재배포할 수 있습니다.

이 프로그램은 유용하게 사용될 수 있으리라는 희망에서 배포되고 있지만,
특정한 목적에 적합하다거나, 이익을 안겨줄 수 있다는 묵시적인 보증을 포함한
어떠한 형태의 보증도 제공하지 않습니다.
보다 자세한 사항에 대해서는 GNU LGPL 2.1판을 참고하시기 바랍니다.
GNU LGPL 2.1판은 이 프로그램과 함께 제공됩니다.
만약, 이 문서가 누락되어 있다면 자유 소프트웨어 재단으로 문의하시기 바랍니다.
(자유 소프트웨어 재단 : Free Software Foundation, Inc.,
59 Temple Place - Suite 330, Boston, MA 02111-1307, USA)

Copyright (C) 2015-2019년 UnHa Kim (unha.kim.ghts@gmail.com)

This file is part of GHTS.

GHTS is free software: you can redistribute it and/or modify
it under the terms of the GNU Lesser General Public License as published by
the Free Software Foundation, version 2.1 of the License.

GHTS is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Lesser General Public License for more details.

You should have received a copy of the GNU Lesser General Public License
along with GHTS.  If not, see <http://www.gnu.org/licenses/>. */


#include "event_handler.hpp"

// Exported to Go
void *NewPair() {
    PairObject *p = new PairObject();
    p->pGiExpert = NewDispatch(CLSID_GiExpert, IID_GiExpert);
    p->pGiExpertEventHandler = new EventHandler();
    Advise(p->pGiExpert, p->pGiExpertEventHandler);

    return p;
}

// Exported to Go
void FreeResources(void *pPairObject) {
    PairObject *p = (PairObject*)pPairObject;
    p->pGiExpertEventHandler->Unadvise(p->pGiExpert);
    p->pGiExpert->Release();
//    p->pGiExpertEventHandler->Release();    // 원인을 알 수 없는 에러가 발생함.
    free(p);
}

// Exported to Go
HRESULT IDispatchInvoke(void *pPairObject, int DispId, void *pDispParams, void *pResult, void *pExcepInfo) {
    PairObject *p = (PairObject*)pPairObject;

    return p->pGiExpert->Invoke((DISPID)DispId, IID_GiExpert, LOCALE_USER_DEFAULT, DISPATCH_METHOD,
        (DISPPARAMS*)pDispParams, (VARIANT*)pResult, (EXCEPINFO*)pExcepInfo, NULL);
}

// Not Exported to Go
IDispatch * NewDispatch(REFCLSID rclsid, REFIID riid) {
	IUnknown *pUnknown;
	IDispatch *pDispatch;
	HRESULT hResult;

	hResult = CoCreateInstance(rclsid, NULL, CLSCTX_INPROC_SERVER, riid, (void **) &pUnknown);
	if (hResult != S_OK) {
		printf("Unexpected HRESULT at NewDispatch() 1 : '%ld'\n", hResult);
		exit(-1);
	}

	hResult = pUnknown->QueryInterface(IID_IDispatch, (void **) &pDispatch);
	if (hResult != S_OK) {
		printf("Unexpected HRESULT at NewDispatch() 2 : '%ld'\n", hResult);
		exit(-1);
	}

	return pDispatch;
}

void Advise(IDispatch *eventSrc, EventHandler *eventHandler) {
	IConnectionPointContainer *pContainer;
	IConnectionPoint *pConnectionPoint;
	HRESULT hResult;
	DWORD cookie;

	hResult = eventSrc->QueryInterface(IID_IConnectionPointContainer, (void **) (&pContainer));
	if (hResult != S_OK) {
		printf("Unexpected HRESULT at Advise() 1 : '%ld'\n", hResult);
		exit(-1);
	}

	hResult = pContainer->FindConnectionPoint(eventHandler->iid, &pConnectionPoint);
	if (hResult != S_OK) {
		printf("Unexpected HRESULT at Advise() 2 : '%ld'\n", hResult);
		exit(-1);
	}

	hResult = pConnectionPoint->Advise((IUnknown*) eventHandler, &cookie);
	if (hResult != S_OK) {
		printf("Unexpected HRESULT at Advise() 3 : '%ld'\n", hResult);
		exit(-1);
	}

	eventHandler->dwCookie = cookie;
}

// EventHandler Methods
EventHandler::~EventHandler() {
    this->Release();
}

HRESULT STDMETHODCALLTYPE EventHandler::Unadvise(IDispatch *eventSrc) {
    IConnectionPointContainer *pContainer;
	IConnectionPoint *pConnectionPoint;
	HRESULT hResult;

	hResult = eventSrc->QueryInterface(IID_IConnectionPointContainer, (void **) (&pContainer));
	if (hResult != S_OK) {
		printf("Unexpected HRESULT at Unadvise() 1 : '%ld'\n", hResult);
		exit(-1);
	}

	hResult = pContainer->FindConnectionPoint(this->iid, &pConnectionPoint);
	if (hResult != S_OK) {
		printf("Unexpected HRESULT at Unadvise() 2 : '%ld'\n", hResult);
		exit(-1);
	}

	hResult = pConnectionPoint->Unadvise(this->dwCookie);
	if (hResult != S_OK) {
		printf("Unexpected HRESULT at Unadvise() 3 : '%ld'\n", hResult);
		exit(-1);
	}

    pConnectionPoint->Release();
    //pContainer->Release();  // Release 해도 괜찮은가??

	return hResult;
}

HRESULT STDMETHODCALLTYPE EventHandler::QueryInterface(REFIID riid, void **ppv) {
	if (riid != IID_IUnknown &&
		riid != IID_IDispatch &&
		riid != this->iid) {
		printf("QueryInterface : E_NOINTERFACE '%s'\n", GuidToString(riid).c_str());
		*ppv = 0;
		return(E_NOINTERFACE);
	}

	*ppv = this;
	this->AddRef();

	return(S_OK);
}

ULONG STDMETHODCALLTYPE EventHandler::AddRef() {
	return ++this->count;
}

ULONG STDMETHODCALLTYPE EventHandler::Release() {
	if (--this->count == 0) {
        delete this;
		return(0);
	}

	return this->count;
}

HRESULT STDMETHODCALLTYPE EventHandler::GetTypeInfoCount(UINT *pctinfo) {
	printf("GetTypeInfoCount()\n");

	if (pctinfo != NULL) {
		*pctinfo = 1;
	}

	return S_OK;
}

HRESULT STDMETHODCALLTYPE EventHandler::GetTypeInfo(UINT iTInfo, LCID lcid, ITypeInfo **ppTInfo) {
	printf("GetTypeInfo()\n");

	return E_NOTIMPL;
}

HRESULT STDMETHODCALLTYPE EventHandler::GetIDsOfNames(REFIID riid, LPOLESTR *rgszNames, UINT cNames, LCID lcid, DISPID *rgDispId) {
	printf("GetIDsOfNames()\n");

	return E_NOTIMPL;
}

HRESULT STDMETHODCALLTYPE EventHandler::Invoke(DISPID dispId, REFIID riid, LCID lcid, WORD wFlags,
    DISPPARAMS *pDispParams, VARIANT *pVarResult, EXCEPINFO *pExcepInfo, UINT *puArgErr) {

    printf("EventHandler::Invoke() : '%ld'\n", dispId);
    printf("EventHandler::Invoke : '%s' '%s'\n", GuidToString(riid).c_str(), GuidToString(this->iid).c_str());
    printf("'%u', '%u',\n", pDispParams->cArgs, pDispParams->cNamedArgs);

	if (riid != IID_NULL) {
	    printf("EventHandler::Invoke() : NOT IID_NULL. Unexpected. '%s'\n", GuidToString(riid).c_str());
		exit(-1);
		return(DISP_E_UNKNOWNINTERFACE);
	}

//    switch (dispId) {
//    case IdReceiveData:
//        ReceiveData_Go(this->TrCodeSn, pVarResult)
//        printf("EventHandler ReceiveData\n");
//    case IdReceiveRTData:
//        printf("EventHandler IdReceiveRTData\n");
//    case IdReceiveSysMsg:
//        printf("EventHandler IdReceiveSysMsg\n");
//    default:
//        printf("EventHandler::Invoke() : unexpectetd dispId '%ld'\n", dispId);
//    }
//
//	Invoke_Go(dispId, pDispParams->cArgs, pDispParams->rgvarg, pVarResult);

	return S_OK;
}

// helper function for debugging message
std::string GuidToString(GUID guid) {
	char guid_cstr[39];
	snprintf(guid_cstr, sizeof(guid_cstr),
	         "{%08lx-%04x-%04x-%02x%02x-%02x%02x%02x%02x%02x%02x}",
	         guid.Data1, guid.Data2, guid.Data3,
	         guid.Data4[0], guid.Data4[1], guid.Data4[2], guid.Data4[3],
	         guid.Data4[4], guid.Data4[5], guid.Data4[6], guid.Data4[7]);

	return std::string(guid_cstr);
}