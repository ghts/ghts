#include "KiwoomEvents.hpp"
#include "Func.hpp"
#include <QDebug>

void KiwoomEvents::OnEventConnectHandler(int nErrCode) {
    qDebug() << "Login" << OK_ERR(nErrCode == 0);

    QLibrary *kiwoom_Go = GetKiwoom_Go();
    if (kiwoom_Go == NULL) {
        qDebug()<<"Abort OnEventConnectHandler().";
        return;
    }

    typedef void (*OnEventConnect)(bool);
    OnEventConnect notifiy = (OnEventConnect)kiwoom_Go->resolve("OnEventConnect");

    if (!notifiy) {
        qDebug()<<"Abort OnEventConnectHandler().";
        return;
    }

    notifiy(nErrCode == 0);
    qDebug()<<"OnEventConnect("<<(nErrCode == 0)<<") OK.";
}

void KiwoomEvents::OnReceiveChejanDataHandler(QString sGubun, int nItemCnt, QString sFIdList) {
    qDebug("TODO : OnReceiveChejanDataHandler()");
}


void KiwoomEvents::OnReceiveConditionVerHandler(int lRet, QString sMsg) {
    qDebug("TODO : OnReceiveConditionVerHandler()");
}


void KiwoomEvents::OnReceiveInvestRealDataHandler(QString sRealKey) {
    qDebug("TODO : OnReceiveInvestRealDataHandler()");
}


void KiwoomEvents::OnReceiveMsgHandler(QString sScrNo, QString sRQName, QString sTrCode, QString sMsg) {
    qDebug("TODO : OnReceiveMsgHandler()");
}


void KiwoomEvents::OnReceiveRealConditionHandler(QString sTrCode, QString strType, QString strConditionName, QString strConditionIndex) {
    qDebug("TODO : OnReceiveRealConditionHandler()");
}


void KiwoomEvents::OnReceiveRealDataHandler(QString sRealKey, QString sRealType, QString sRealData) {
    qDebug("TODO : OnReceiveRealDataHandler()");
}


void KiwoomEvents::OnReceiveTrConditionHandler(QString sScrNo, QString strCodeList, QString strConditionName, int nIndex, int nNext) {
    qDebug("TODO : OnReceiveTrConditionHandler()");
}


void KiwoomEvents::OnReceiveTrDataHandler(QString sScrNo, QString sRQName, QString sTrCode, QString sRecordName, QString sPrevNext, int nDataLength, QString sErrorCode, QString sMessage, QString sSplmMsg) {
    qDebug("TODO : OnReceiveTrDataHandler()");
}

void KiwoomEvents::exceptionHandler(int code, QString source, QString disc, QString help) {
    qDebug("TODO : exceptionHandler()");
}

void KiwoomEvents::propertyChangedHandler(QString name) {
    qDebug("TODO : propertyChangedHandler()");
}

void KiwoomEvents::signalHandler(QString name, int argc, void* argv) {
    qDebug("TODO : signalHandler()");
}
