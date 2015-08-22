// WMCALOADER.h : main header file for the WMCALOADER application
//

#if !defined(AFX_WMCALOADER_H__D006676A_B29A_4E20_A08D_A99129820CAF__INCLUDED_)
#define AFX_WMCALOADER_H__D006676A_B29A_4E20_A08D_A99129820CAF__INCLUDED_

#if _MSC_VER > 1000
#pragma once
#endif // _MSC_VER > 1000

#ifndef __AFXWIN_H__
	#error include 'stdafx.h' before including this file for PCH
#endif

#include "resource.h"		// main symbols

/////////////////////////////////////////////////////////////////////////////
// CWMCALOADERApp:
// See WMCALOADER.cpp for the implementation of this class
//

class CWMCALOADERApp : public CWinApp
{
public:
	CWMCALOADERApp();

// Overrides
	// ClassWizard generated virtual function overrides
	//{{AFX_VIRTUAL(CWMCALOADERApp)
	public:
	virtual BOOL InitInstance();
	//}}AFX_VIRTUAL

// Implementation

	//{{AFX_MSG(CWMCALOADERApp)
	//}}AFX_MSG
	DECLARE_MESSAGE_MAP()
};


/////////////////////////////////////////////////////////////////////////////

//{{AFX_INSERT_LOCATION}}
// Microsoft Visual C++ will insert additional declarations immediately before the previous line.

#endif // !defined(AFX_WMCALOADER_H__D006676A_B29A_4E20_A08D_A99129820CAF__INCLUDED_)
