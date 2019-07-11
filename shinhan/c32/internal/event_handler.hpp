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

#include <iostream>
#include "event_handler.h"
#include "_cgo_export.h"

class EventHandler: public IDispatch {

public:
    DWORD				count;
    REFIID				iid = IID_GiExpertEvents;
    DWORD				dwCookie;

    virtual ~EventHandler();

    HRESULT STDMETHODCALLTYPE Unadvise(IDispatch *eventSrc);

    // Implements IDispatch
	HRESULT STDMETHODCALLTYPE QueryInterface(REFIID riid, void **ppvc);
	ULONG STDMETHODCALLTYPE AddRef();
	ULONG STDMETHODCALLTYPE Release();
	HRESULT STDMETHODCALLTYPE GetTypeInfoCount(UINT *pctinfo);
	HRESULT STDMETHODCALLTYPE GetTypeInfo(UINT iTInfo, LCID lcid, ITypeInfo **ppTInfo);
	HRESULT STDMETHODCALLTYPE GetIDsOfNames(REFIID riid, LPOLESTR *rgszNames, UINT cNames, LCID lcid, DISPID *rgDispId);
	HRESULT STDMETHODCALLTYPE Invoke(DISPID, REFIID, LCID, WORD, DISPPARAMS*, VARIANT*, EXCEPINFO*, UINT*);
};

IDispatch * NewDispatch(REFCLSID, REFIID);
void Advise(IDispatch*, EventHandler*);
std::string GuidToString(GUID);

typedef struct {
    IDispatch *pGiExpert;
    EventHandler *pGiExpertEventHandler;
} PairObject;

const int IdSetSingleData      = 0x01;
const int IdSetMultiData       = 0x02;
const int IdSetQueryName       = 0x03;
const int IdGetQueryName       = 0x04;
const int IdRequestData        = 0x05;
const int IdRequestRTReg       = 0x06;
const int IdUnRequestRTReg     = 0x07;
const int IdGetSingleData      = 0x08;
const int IdGetMultiData       = 0x09;
const int IdGetSingleBlockData = 0x0a;
const int IdGetMultiBlockData  = 0x0b;
const int IdGetSingleRowCount  = 0x0c;
const int IdGetMultiRowCount   = 0x0d;
const int IdGetErrorState      = 0x0e;
const int IdGetErrorCode       = 0x0f;
const int IdGetErrorMessage    = 0x10;
const int IdGetCommState       = 0x11;
const int IdUnRequestRTRegAll  = 0x12;
const int IdSetRQCount         = 0x13;
const int IdClearReceiveBuffer = 0x14;
const int IdSelfMemFree        = 0x15;
const int IdSetID              = 0x16;
const int IdGetCodeByName      = 0x17;
const int IdSetSingleEncData   = 0x18;
const int IdStartIndi          = 0x19;
const int IdCloseIndi          = 0x1a;
const int IdGetInputSingleData = 0x1b;
const int IdGetInputMultiData  = 0x1c;
const int IdGetInputTRName     = 0x1d;

const int IdReceiveData        = 0x01;
const int IdReceiveRTData      = 0x02;
const int IdReceiveSysMsg      = 0x03;


// For Test Only
bool OnLogInEventReceived = false;
HRESULT Invoke(IDispatch*, DISPID, DISPPARAMS*, VARIANT*);
HRESULT IsApiLoaded(IDispatch*);
HRESULT ConnectServer(IDispatch*);
HRESULT LogIn(IDispatch*);
void PumpWaitingMessages();
