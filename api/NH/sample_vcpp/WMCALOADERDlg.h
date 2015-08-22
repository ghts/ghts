// WMCALOADERDlg.h : header file
//

#if !defined(AFX_WMCALOADERDLG_H__934EBB58_9AE3_4983_B2C6_E5BAB237D35E__INCLUDED_)
#define AFX_WMCALOADERDLG_H__934EBB58_9AE3_4983_B2C6_E5BAB237D35E__INCLUDED_

#if _MSC_VER > 1000
#pragma once
#endif // _MSC_VER > 1000

#include	"wmcaintf.h"

/////////////////////////////////////////////////////////////////////////////
// CWMCALOADERDlg dialog

class CWMCALOADERDlg : public CDialog
{
// Construction
public:
	CWMCALOADERDlg(CWnd* pParent = NULL);	// standard constructor

// Dialog Data
	//{{AFX_DATA(CWMCALOADERDlg)
	enum { IDD = IDD_WMCALOADER_DIALOG };
	CListCtrl	m_listBalance;
	CListCtrl	m_listPrice;
	CListBox	m_listboxTrace;
	CComboBox	m_comboAccountList;
	CStatic	m_staticLoginTime;
	CEdit	m_editSignPassword;
	CEdit	m_editPassword;
	CEdit	m_editID;
	CButton	m_buttonConnect;
	CButton m_buttonDisconnect;
	CButton	m_buttonBalance;
	//}}AFX_DATA

	// ClassWizard generated virtual function overrides
	//{{AFX_VIRTUAL(CWMCALOADERDlg)
	protected:
	virtual void DoDataExchange(CDataExchange* pDX);	// DDX/DDV support
	//}}AFX_VIRTUAL

// Implementation
protected:
	HICON m_hIcon;
	CWmcaIntf	m_wmca;

	void OnWmConnected( LOGINBLOCK* pLogin );
	void OnWmDisconnected();
	void OnWmSocketerror(int socket_error_code);
	void OnWmReceivedata( OUTDATABLOCK* pOutData );
	void OnWmReceivesise( OUTDATABLOCK* pSiseData );
	void OnWmReceivemessage( OUTDATABLOCK* pMessage );
	void OnWmReceivecomplete( OUTDATABLOCK* pOutData );
	void OnWmReceiveerror( OUTDATABLOCK* pError );

	// Generated message map functions
	//{{AFX_MSG(CWMCALOADERDlg)
	virtual BOOL OnInitDialog();
	afx_msg void OnPaint();
	afx_msg HCURSOR OnQueryDragIcon();
	virtual void OnOK();
	virtual void OnCancel();
	afx_msg void OnConnect();
	afx_msg void OnDisconnect();
	afx_msg void OnBalance();
	afx_msg void OnButtonCurrent();
	afx_msg void OnButtonClear();
	afx_msg void OnButtonOrder();
	//}}AFX_MSG
	afx_msg LRESULT OnWmcaEvent(WPARAM wParam, LPARAM lParam);
	afx_msg void	ScrollDown();
	DECLARE_MESSAGE_MAP()
public:
};

//{{AFX_INSERT_LOCATION}}
// Microsoft Visual C++ will insert additional declarations immediately before the previous line.

#endif // !defined(AFX_WMCALOADERDLG_H__934EBB58_9AE3_4983_B2C6_E5BAB237D35E__INCLUDED_)
