#include "WinMsgHandler.hpp"
#include "WinMsg.hpp"
#include "Func.hpp"
#include "Const.hpp"

#include <qt_windows.h>
#include <QDebug>
#include <QString>

bool WinMsgHandler::nativeEventFilter(const QByteArray &, void *message, long *) {
    MSG *msg = static_cast<MSG*>(message);
    unsigned int uMsg = (unsigned int)(msg->message);

    // WPARAM에는 일련번호, LPARAM에는 추가적인 데이터 (혹은 포인터)를 지정하기로 한다.
    WPARAM serialNo = msg->wParam;
    LPARAM lParam = msg->lParam;

    if (KM_CONNECT == uMsg) {
        qDebug("KM_CONNECT recevied.");

        int result = kiwoom->CommConnect();
        qDebug()<<"CommConnect()"<<OK_ERR(result == 0);

        QString resultString = QString::number(result);

        Confirm(serialNo, resultString);
        qDebug()<<"Confirm CommConnect(). "<<resultString;

        return true;
    } else if (KM_LOGIN_INFO == uMsg) {
        qDebug("KM_LOGIN_INFO recevied.");

        QString tag;

        switch (lParam) {
        case ACCOUNT_CNT:
            tag = QString("ACCOUNT_CNT");
        case ACCNO:
            tag = QString("ACCNO");
        case USER_ID:
            tag = QString("USER_ID");
        case USER_NAME:
            tag = QString("USER_NAME");
        case KEY_BSECGB:
            tag = QString("KEY_BSECGB");
        case FIREW_SECGB:
            tag = QString("FIREW_SECGB");
        }

        QString result = kiwoom->GetLoginInfo(tag);
        qDebug()<<"GetLoginInfo("<<tag<<") : '"<<result<<"'";

        Confirm(serialNo, result);
        qDebug()<<"Confirm GetLoginInfo() : '"<<result<<"'";

        return true;
    } else if (KM_PRINT_DEBUG_MSG == uMsg) {
        qDebug()<<QString::fromUtf8((const char*)lParam);
    }

    if (uMsg >= KM_INIT && uMsg <= (KM_INIT+100)) {
        qDebug()<<uMsg<<" : Not implemented Yet.";
    }

    return false;
}
