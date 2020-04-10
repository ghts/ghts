#include "WinMsgHandler.hpp"
#include "WinMsg.hpp"
#include "Func.hpp"

#include <qt_windows.h>
#include <QDebug>

bool WinMsgHandler::nativeEventFilter(const QByteArray &, void *message, long *) {
    MSG *msg = static_cast<MSG*>(message);
    unsigned int uMsg = (unsigned int)(msg->message);

    // WPARAM에는 일련번호, LPARAM에는 데이터 포인터를 지정하기로 한다.
    WPARAM serialNo = msg->wParam;
    LPARAM ptrData = msg->lParam;

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

        QString sTag = QString((const char*)ptrData);
        QString result = kiwoom->GetLoginInfo(sTag);
        qDebug()<<"GetLoginInfo("<<sTag<<") : '"<<result<<"'";

        Confirm(serialNo, result);
        qDebug()<<"Confirm GetLoginInfo().";

        return true;
    }

    if (uMsg >= KM_INIT && uMsg <= (KM_INIT+100)) {
        qDebug()<<uMsg<<" : Not implemented Yet.";
    }

    return false;
}
