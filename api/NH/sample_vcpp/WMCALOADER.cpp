// WMCALOADER.cpp : Defines the class behaviors for the application.
//

#include "stdafx.h"
#include "WMCALOADER.h"
#include "WMCALOADERDlg.h"

#ifdef _DEBUG
#define new DEBUG_NEW
#undef THIS_FILE
static char THIS_FILE[] = __FILE__;
#endif

/////////////////////////////////////////////////////////////////////////////
// CWMCALOADERApp

BEGIN_MESSAGE_MAP(CWMCALOADERApp, CWinApp)
	//{{AFX_MSG_MAP(CWMCALOADERApp)
	//}}AFX_MSG
	ON_COMMAND(ID_HELP, CWinApp::OnHelp)
END_MESSAGE_MAP()

/////////////////////////////////////////////////////////////////////////////
// CWMCALOADERApp construction

CWMCALOADERApp::CWMCALOADERApp()
{
}

/////////////////////////////////////////////////////////////////////////////
// The one and only CWMCALOADERApp object

CWMCALOADERApp theApp;

/////////////////////////////////////////////////////////////////////////////
// CWMCALOADERApp initialization

BOOL CWMCALOADERApp::InitInstance()
{
	AfxEnableControlContainer();

	// Standard initialization

#ifdef _AFXDLL
#if _MSC_VER <= 1200     
	Enable3dControls();			// Call this when using MFC in a shared DLL
#endif
#else
	Enable3dControlsStatic();	// Call this when linking to MFC statically
#endif

	CWMCALOADERDlg dlg;
	m_pMainWnd = &dlg;
	int nResponse = dlg.DoModal();
	if (nResponse == IDOK)
	{
	}
	else if (nResponse == IDCANCEL)
	{
	}

	// Since the dialog has been closed, return FALSE so that we exit the
	//  application, rather than start the application's message pump.
	return FALSE;
}
