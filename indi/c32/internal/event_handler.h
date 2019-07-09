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

#include <initguid.h>
#include <OCIdl.h>
#include <objbase.h>
#include <stdio.h>
#include <cguid.h>

// GUID문자열을 구조체에 맞게 나누고, 앞에 0x만 붙인 것임. 별도의 변환 필요없음.
DEFINE_GUID(CLSID_GiExpert, 0x961db208, 0x0f2f, 0x41f0, 0x95, 0xc8, 0x72, 0x36, 0x33, 0x85, 0x78, 0x44); // {961DB208-0F2F-41F0-95C8-723633857844}
DEFINE_GUID(IID_GiExpert, 0xb4a0d92f, 0x6683, 0x4824, 0x9c, 0x63, 0x53, 0xf7, 0xda, 0xd7, 0xf5, 0x7c);   // {B4A0D92F-6683-4824-9C63-53F7DAD7F57C}
DEFINE_GUID(IID_GiExpertEvents, 0xf52e833f, 0x1d17, 0x4184, 0xbe, 0x0d, 0x6b, 0x9f, 0x59, 0xde, 0x7d, 0xac); // {F52E833F-1D17-4184-BE0D-6B9F59DE7DAC}

#ifdef __cplusplus
extern "C" {
#endif

void *NewPair();
void FreeResources(void*);
HRESULT IDispatchInvoke(void*, int, void*, void*, void*);

#ifdef __cplusplus
}
#endif

// C Only. No C++.
#ifndef __cplusplus

// Interface of EventHandler
#undef  INTERFACE
#define INTERFACE IEventHandler
DECLARE_INTERFACE_ (INTERFACE, IDispatch) {
    BEGIN_INTERFACE

	// IUnknown functions
	STDMETHOD  (QueryInterface)		(THIS_ REFIID, void **) PURE;
	STDMETHOD_ (ULONG, AddRef)		(THIS) PURE;
	STDMETHOD_ (ULONG, Release)		(THIS) PURE;
	// IDispatch functions
	STDMETHOD (GetTypeInfoCount)    (THIS_ UINT *) PURE;
	STDMETHOD (GetTypeInfo)		    (THIS_ UINT, LCID, ITypeInfo **) PURE;
	STDMETHOD (GetIDsOfNames)	    (THIS_ REFIID, LPOLESTR *, UINT, LCID, DISPID *) PURE;
	STDMETHOD (Invoke)			    (THIS_ DISPID, REFIID, LCID, WORD, DISPPARAMS *, VARIANT *, EXCEPINFO *, UINT *) PURE;

	END_INTERFACE
};

HRESULT STDMETHODCALLTYPE QueryInterface(IEventHandler*, REFIID, void**);
ULONG STDMETHODCALLTYPE AddRef(IEventHandler*);
ULONG STDMETHODCALLTYPE Release(IEventHandler*);
HRESULT STDMETHODCALLTYPE GetTypeInfoCount(IEventHandler*, UINT*);
HRESULT STDMETHODCALLTYPE GetTypeInfo(IEventHandler*, UINT, LCID, ITypeInfo**);
HRESULT STDMETHODCALLTYPE GetIDsOfNames(IEventHandler*, REFIID, LPOLESTR*, UINT, LCID, DISPID*);
HRESULT STDMETHODCALLTYPE Invoke(IEventHandler*, DISPID, REFIID, LCID, WORD, DISPPARAMS*, VARIANT*, EXCEPINFO*, UINT*);

typedef struct {
	IEventHandlerVtbl	*lpVtbl;
	DWORD				count;
	REFIID				iid;
	DWORD				dwCookie;
} IEventHandlerImpl;

IEventHandler * NewEventHandler();
HRESULT Advise(IDispatch*, REFIID, IEventHandler*);

#endif
