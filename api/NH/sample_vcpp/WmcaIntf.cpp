// WmcaIntf.cpp: implementation of the CWmcaIntf class.
//
//////////////////////////////////////////////////////////////////////

#include "stdafx.h"
#include "WmcaIntf.h"

//////////////////////////////////////////////////////////////////////
// Construction/Destruction
//////////////////////////////////////////////////////////////////////

CWmcaIntf::CWmcaIntf()	:
	m_hDll(NULL),
	m_pLoad(NULL),	m_pFree(NULL),	m_pSetServer(NULL),	m_pSetPort(NULL),	m_pIsConnected(NULL),	
	m_pConnect(NULL),	m_pDisconnect(NULL),	m_pTransact(NULL),	m_pQuery(NULL),	m_pRequest(NULL),	m_pAttach(NULL),	
	m_pDetach(NULL),	m_pDetachWindow(NULL),	m_pDetachAll(NULL), m_pSetOption(NULL), 
	m_pSetAccountIndexPwd(NULL),	m_pSetOrderPwd(NULL),	m_pSetHashPwd(NULL),	m_pSetAccountNoPwd(NULL), m_pSetAccountNoByIndex(NULL)
{
	if(!(m_hDll	=LoadLibrary("wmca.dll")))
	{
		AfxMessageBox(	"WMCA.DLL 파일이 없습니다.\n"
						"이 파일이 경로에 있는지 확인하시기 바랍니다.");

		return;
	}

	m_pLoad					=(TLoad*)					GetProcAddress(m_hDll,"wmcaLoad");
	m_pFree					=(TFree*)					GetProcAddress(m_hDll,"wmcaFree");
	m_pSetServer			=(TSetServer*)				GetProcAddress(m_hDll,"wmcaSetServer");
	m_pSetPort				=(TSetPort*)				GetProcAddress(m_hDll,"wmcaSetPort");
	m_pIsConnected			=(TIsConnected*)			GetProcAddress(m_hDll,"wmcaIsConnected");
	m_pConnect				=(TConnect*)				GetProcAddress(m_hDll,"wmcaConnect");
	m_pDisconnect			=(TDisconnect*)				GetProcAddress(m_hDll,"wmcaDisconnect");
	m_pTransact				=(TTransact*)				GetProcAddress(m_hDll,"wmcaTransact");
	m_pQuery				=(TQuery*)					GetProcAddress(m_hDll,"wmcaQuery");
	m_pRequest				=(TRequest*)				GetProcAddress(m_hDll,"wmcaRequest");
	m_pAttach				=(TAttach*)					GetProcAddress(m_hDll,"wmcaAttach");
	m_pDetach				=(TDetach*)					GetProcAddress(m_hDll,"wmcaDetach");
	m_pDetachWindow			=(TDetachWindow*)			GetProcAddress(m_hDll,"wmcaDetachWindow");
	m_pDetachAll			=(TDetachAll*)				GetProcAddress(m_hDll,"wmcaDetachAll");
	m_pSetOption			=(TSetOption*)				GetProcAddress(m_hDll,"wmcaSetOption");
	m_pSetAccountIndexPwd	=(TSetAccountIndexPwd*)		GetProcAddress(m_hDll,"wmcaSetAccountIndexPwd");
	m_pSetOrderPwd			=(TSetOrderPwd*)			GetProcAddress(m_hDll,"wmcaSetOrderPwd");
	m_pSetHashPwd			=(TSetHashPwd*)				GetProcAddress(m_hDll,"wmcaSetHashPwd");
	m_pSetAccountNoPwd		=(TSetAccountNoPwd*)		GetProcAddress(m_hDll,"wmcaSetAccountNoPwd");
	m_pSetAccountNoByIndex	=(TSetAccountNoByIndex*)	GetProcAddress(m_hDll,"wmcaSetAccountNoByIndex");



	if(!(m_pLoad && m_pFree && m_pSetServer && m_pSetPort && m_pIsConnected && m_pConnect && m_pDisconnect && m_pTransact && m_pQuery
		&& m_pRequest && m_pAttach && m_pDetach && m_pDetachWindow && m_pDetachAll && m_pSetOption 
		&& m_pSetAccountIndexPwd && m_pSetOrderPwd && m_pSetHashPwd && m_pSetAccountNoPwd) && m_pSetAccountNoByIndex)
	{
		AfxMessageBox(	"WMCA.DLL 파일을 인식할 수 없습니다.\n"
						"최신 버전의 정상적인 파일인지 확인하시기 바랍니다.");

		return;
	}

}

CWmcaIntf::~CWmcaIntf()
{
	if(m_hDll)
	{
		if(m_pFree)
			m_pFree();

		FreeLibrary(m_hDll);

		m_hDll	=0;
	}
}

BOOL	CWmcaIntf::Load()
{
	if(m_pLoad)
		return	m_pLoad();

	return false;
}

BOOL	CWmcaIntf::Free()
{
	if(m_pFree)
		return	m_pFree();

	return false;
}

BOOL	CWmcaIntf::Connect(HWND hWnd, DWORD msg, char MediaType,char UserType,const char* szID,const char* szPW, const char* szCertPW)
{
	if(m_pConnect)
		return	m_pConnect(hWnd, msg, MediaType,UserType,szID,szPW, szCertPW);

	return false;
}

BOOL	CWmcaIntf::Disconnect()
{
	if(m_pDisconnect)
		return m_pDisconnect();

	return false;
}

BOOL	CWmcaIntf::Transact(HWND hWnd, int nTRID, const char* szTRCode, const char* szInput, int nInputLen, int nHeadType, int nAccountIndex)
{
	if(m_pTransact)
		return m_pTransact(hWnd, nTRID, szTRCode, szInput, nInputLen, nHeadType, nAccountIndex);
	
	return false;
}

BOOL	CWmcaIntf::Query(HWND hWnd, int nTRID, const char* szTRCode, const char* szInput, int nInputLen, int nAccountIndex)
{
	if(m_pQuery)
		return m_pQuery(hWnd, nTRID, szTRCode, szInput, nInputLen, nAccountIndex);
	
	return false;
}

BOOL	CWmcaIntf::Request(HWND hWnd, int nTRID, const char* szTRCode, const char* szInput, int nInputLen, const char* szOpenBranchCode)
{
	if(m_pRequest)
		return m_pRequest(hWnd, nTRID, szTRCode, szInput, nInputLen, szOpenBranchCode);
	
	return false;
}

BOOL	CWmcaIntf::Attach(HWND hWnd, const char* szBCType, const char* szInput, int nCodeLen, int nInputLen)
{
	if(m_pAttach)
		return m_pAttach(hWnd, szBCType, szInput, nCodeLen, nInputLen);

	return false;
}

BOOL	CWmcaIntf::Detach(HWND hWnd, const char* szBCType, const char* szInput, int nCodeLen, int nInputLen)
{
	if(m_pDetach)
		return m_pDetach(hWnd, szBCType, szInput, nCodeLen, nInputLen);

	return false;
}

BOOL	CWmcaIntf::DetachWindow(HWND hWnd)
{
	if(m_pDetachWindow)
		return m_pDetachWindow(hWnd);
	
	return false;
}

BOOL	CWmcaIntf::DetachAll()
{
	if(m_pDetachAll)
		return m_pDetachAll();
	
	return false;
}

BOOL	CWmcaIntf::SetOption(const char* szKey,const char* szValue)
{
	if(m_pSetOption)
		return m_pSetOption(szKey,szValue);
	
	return false;
}

BOOL CWmcaIntf::SetServer(const char* szServer)
{
	if(m_pSetServer)
		return m_pSetServer(szServer);
	
	return false;
}

BOOL CWmcaIntf::SetPort(int port)
{
	if(m_pSetPort)
		return m_pSetPort(port);
	
	return false;
}

BOOL CWmcaIntf::IsConnected()
{
	if(m_pIsConnected)
		return m_pIsConnected();
	
	return false;
}

BOOL CWmcaIntf::SetAccountIndexPwd		(const char* pszHashOut,int nAccountIndex,const char* pszPassword)
{
	if(m_pSetAccountIndexPwd)
		return m_pSetAccountIndexPwd(pszHashOut,nAccountIndex,pszPassword);
	
	return false;
}

BOOL CWmcaIntf::SetOrderPwd		(const char* pszHashOut,const char* pszPassword)
{
	if(m_pSetOrderPwd)
		return m_pSetOrderPwd(pszHashOut,pszPassword);
	
	return false;
}

BOOL CWmcaIntf::SetHashPwd	(const char* pszHashOut,const char* pszKey,const char* pszPassword)
{
	if(m_pSetHashPwd)
		return m_pSetHashPwd(pszHashOut,pszKey,pszPassword);
	
	return false;
}

BOOL CWmcaIntf::SetAccountNoPwd	(const char* pszHashOut,const char* pszAccountNo,const char* pszPassword)
{
	if(m_pSetAccountNoPwd)
		return m_pSetAccountNoPwd(pszHashOut,pszAccountNo,pszPassword);
	
	return false;
}

BOOL CWmcaIntf::SetAccountNoByIndex	(const char* pszHashOut,int nAccountIndex)
{
	if(m_pSetAccountNoByIndex)
		return m_pSetAccountNoByIndex(pszHashOut,nAccountIndex);
	
	return false;
}
