#include "WinMsgHandler.hpp"
#include "Func.hpp"
#include "Const.hpp"

#include <qt_windows.h>
#include <QDebug>
#include <QString>

void KM_CONNECT_Handler(WPARAM, KHOpenAPILib::KHOpenAPI*);
void KM_LOGIN_INFO_Handler(WPARAM, LPARAM, KHOpenAPILib::KHOpenAPI*);

bool WinMsgHandler::nativeEventFilter(const QByteArray &, void *message, long *) {
    MSG *msg = static_cast<MSG*>(message);
    unsigned int uMsg = (unsigned int)(msg->message);

    // WPARAM에는 일련번호, LPARAM에는 추가적인 데이터 (혹은 포인터)를 지정하기로 한다.
    WPARAM serialNo = msg->wParam;
    LPARAM lParam = msg->lParam;

    if (KM_CONNECT == uMsg) {
        KM_CONNECT_Handler(serialNo, kiwoom);
    } else if (KM_LOGIN_INFO == uMsg) {
        KM_LOGIN_INFO_Handler(serialNo, lParam, kiwoom);
    } else if (KM_PRINT_DEBUG_MSG == uMsg) {
        qDebug()<<QString::fromUtf8((const char*)lParam);
    } else {
        if (uMsg >= KM_INIT && uMsg <= (KM_INIT+100)) {
            qDebug()<<uMsg<<" : Not implemented Yet.";
        }

        return false;
    }

    return true;
}

void KM_CONNECT_Handler(WPARAM serialNo, KHOpenAPILib::KHOpenAPI *kiwoom) {
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

    QString result = kiwoom->GetLoginInfo (tag);
    qDebug()<<"C++ GetLoginInfo("<<tag<<") Result : '"<<result<<"'";

    Confirm(serialNo, result);
}
