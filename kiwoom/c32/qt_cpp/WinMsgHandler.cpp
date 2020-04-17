#include "WinMsgHandler.hpp"
#include "Func.hpp"
#include "Const.hpp"

#include <qt_windows.h>
#include <QDebug>
#include <QString>

typedef void (*MsgHandler)(WPARAM, LPARAM, KHOpenAPILib::KHOpenAPI*);


void KM_CONNECT_Handler(WPARAM, LPARAM, KHOpenAPILib::KHOpenAPI*);
void KM_LOGIN_INFO_Handler(WPARAM, LPARAM, KHOpenAPILib::KHOpenAPI*);
void KM_CODE_LIST_Handler(WPARAM, LPARAM, KHOpenAPILib::KHOpenAPI*);
void KM_CONNECT_STATE_Handler(WPARAM, LPARAM, KHOpenAPILib::KHOpenAPI*);
void dummy(WPARAM, LPARAM, KHOpenAPILib::KHOpenAPI*);

bool WinMsgHandler::nativeEventFilter(const QByteArray &, void *message, long *) {
    MSG *msg = static_cast<MSG*>(message);
    unsigned int uMsg = (unsigned int)(msg->message);
    WPARAM serialNo = msg->wParam;  // 일련번호
    LPARAM lParam = msg->lParam;    // 추가적인 데이터
    MsgHandler msgHandler = dummy;

    if (KM_CONNECT == uMsg) {
        msgHandler = KM_CONNECT_Handler;
    } else if (KM_LOGIN_INFO == uMsg) {
        msgHandler = KM_LOGIN_INFO_Handler;
    } else if (KM_CODE_LIST == uMsg) {
        msgHandler = KM_CODE_LIST_Handler;
    } else if (KM_CONNECT_STATE == uMsg) {
        msgHandler = KM_CONNECT_STATE_Handler;
    } else if (KM_PRINT_DEBUG_MSG == uMsg) {
        qDebug()<<QString::fromUtf8((const char*)lParam);
    } else if (uMsg >= KM_INIT && uMsg <= (KM_INIT+100)) {
        qDebug()<<uMsg<<" : Not implemented Yet.";
    }

    if (msgHandler == dummy) {
        return false;
    }

    msgHandler(serialNo, lParam, kiwoom);

    return true;
}

void KM_CONNECT_Handler(WPARAM serialNo, LPARAM, KHOpenAPILib::KHOpenAPI *kiwoom) {
    int result = kiwoom->CommConnect();

    QString resultString = QString::number(result);
    Confirm(serialNo, resultString);
}

void KM_LOGIN_INFO_Handler(WPARAM serialNo, LPARAM lParam, KHOpenAPILib::KHOpenAPI *kiwoom) {
    QString tag;

    if (ACCOUNT_CNT == lParam) {
        tag = QString("ACCOUNT_CNT");
        qDebug()<<"C++ KM_LOGIN_INFO recevied : ACCOUNT_CNT "<<lParam;
    } else if (ACCNO == lParam) {
        tag = QString("ACCNO");
        qDebug()<<"C++ KM_LOGIN_INFO recevied : ACCNO "<<lParam;
    } else if (USER_ID == lParam) {
        tag = QString("USER_ID");
        qDebug()<<"C++ KM_LOGIN_INFO recevied : USER_ID "<<lParam;
    } else if (USER_NAME == lParam) {
        tag = QString("USER_NAME");
        qDebug()<<"C++ KM_LOGIN_INFO recevied : USER_NAME"<<lParam;
    } else if (KEY_BSECGB == lParam) {
        tag = QString("KEY_BSECGB");
        qDebug()<<"C++ KM_LOGIN_INFO recevied : KEY_BSECGB"<<lParam;
    } else if (FIREW_SECGB == lParam) {
        tag = QString("FIREW_SECGB");
        qDebug()<<"C++ KM_LOGIN_INFO recevied : FIREW_SECGB"<<lParam;
    } else {
        qDebug()<<"C++ KM_LOGIN_INFO 예상하지 못한 값 : "<<lParam;
        return;
    }

    QString result = kiwoom->GetLoginInfo(tag);
    qDebug()<<"C++ GetLoginInfo("<<tag<<") Result : '"<<result<<"'";

    Confirm(serialNo, result);
}

void KM_CODE_LIST_Handler(WPARAM serialNo, LPARAM lParam, KHOpenAPILib::KHOpenAPI *kiwoom) {
    qDebug()<<"KM_CODE_LIST_Handler Start.";

    QString market = NULL;

    if (CODE_LIST_ALL == lParam) {
         qDebug()<<"KM_CODE_LIST_Handler : CODE_LIST_ALL."<<market;
        // market = NULL;
    } else if (CODE_LIST_KOSPI == lParam      ||
               CODE_LIST_KOSDAQ == lParam     ||
               CODE_LIST_ELW == lParam        ||
               CODE_LIST_ETF == lParam        ||
               CODE_LIST_KONEX == lParam      ||
               CODE_LIST_FUND == lParam       ||
               CODE_LIST_WARRANT == lParam    ||
               CODE_LIST_REITS == lParam      ||
               CODE_LIST_HIGH_YIELD == lParam ||
               CODE_LIST_K_OTC == lParam) {
        market = QString::number(lParam);

        qDebug()<<"KM_CODE_LIST_Handler : "<<market;
    } else {
        qDebug()<<"C++ GetCodeListByMarket() unexpeted marekt type : "<<lParam;
    }

    qDebug()<<"C++ GetCodeListByMarket("<<market<<") Call.";
    QString result = kiwoom->GetCodeListByMarket(market);
    qDebug()<<"C++ GetCodeListByMarket("<<market<<") Result : '"<<result<<"'";

    Confirm(serialNo, result);
     qDebug()<<"KM_CODE_LIST_Handler Confirm :"<< result;
}

void KM_CONNECT_STATE_Handler(WPARAM serialNo, LPARAM, KHOpenAPILib::KHOpenAPI *kiwoom) {
    int result = kiwoom->GetConnectState();

    QString resultString = QString::number(result);
    qDebug()<<"C++ GetConnectState() Result : '"<<result<<"'";

    Confirm(serialNo, resultString);
}

void dummy(WPARAM, LPARAM, KHOpenAPILib::KHOpenAPI*) {}
