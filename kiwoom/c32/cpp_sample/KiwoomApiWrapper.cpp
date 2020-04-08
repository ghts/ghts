#include <KiwoomApiWrapper.hpp>

void KiwoomApiWrapper::connectEventSignalSlot() {
    connect(kiwoom, SIGNAL(OnEventConnect(int)), this,  SLOT(OnEventConnectHandler(int)));
    connect(kiwoom, SIGNAL(OnReceiveChejanData(QString, int, QString)), this,  SLOT(OnReceiveChejanDataHandler(QString, int, QString)));
    connect(kiwoom, SIGNAL(OnReceiveConditionVer(int, QString)), this,  SLOT(OnReceiveConditionVerHandler(int, QString)));
    connect(kiwoom, SIGNAL(OnReceiveInvestRealData(QString)), this,  SLOT(OnReceiveInvestRealDataHandler(QString)));
    connect(kiwoom, SIGNAL(OnReceiveMsg(QString, QString, QString, QString)), this,  SLOT(OnReceiveMsgHandler(QString, QString, QString, QString)));
    connect(kiwoom, SIGNAL(OnReceiveRealCondition(QString, QString, QString, QString)), this,  SLOT(OnReceiveRealConditionHandler(QString, QString, QString, QString)));
    connect(kiwoom, SIGNAL(OnReceiveRealData(QString, QString, QString)), this,  SLOT(OnReceiveRealDataHandler(QString, QString, QString)));
    connect(kiwoom, SIGNAL(OnReceiveTrCondition(QString, QString, QString, int, int)), this,  SLOT(OnReceiveTrConditionHandler(QString, QString, QString, int, int)));
    connect(kiwoom, SIGNAL(OnReceiveTrData(QString, QString, QString, QString, QString, int, QString, QString, QString)), this,  SLOT(OnReceiveTrDataHandler(QString, QString, QString, QString, QString, int, QString, QString, QString)));
    connect(kiwoom, SIGNAL(exception(int, QString, QString, QString)), this,  SLOT(exceptionHandler(int, QString, QString, QString)));
    connect(kiwoom, SIGNAL(propertyChanged(QString)), this,  SLOT(propertyChangedHandler(QString)));
    connect(kiwoom, SIGNAL(signal(QString, int, void*)), this,  SLOT(signalHandler(QString, int, void*)));
}

int KiwoomApiWrapper::CommConnect() {
    int result = kiwoom->CommConnect();

    if (result == 0) {
        qDebug("CommConnect() OK");
        emit postDebugMessage("CommConnect() OK");
    } else if (result < 0) {
        qDebug("CommConnect() Error");
        emit postDebugMessage("CommConnect() Error");
    } else {
        qDebug("CommConnect() Error. Unexpected value.");
        emit postDebugMessage("CommConnect() Error. Unexpected value.");
    }

    return result;
}

void KiwoomApiWrapper::OnEventConnectHandler(int nErrCode) {
    if (nErrCode == 0) {
        qDebug("Login OK.");
        emit postDebugMessage("Login OK.");
    } else if (nErrCode < 0) {
        qDebug("Login Error.");
        emit postDebugMessage("Login Error.");
    } else if (nErrCode > 0) {
        qDebug("Login Error. Unexpected plus value.");
        emit postDebugMessage("Login Error. Unexpected plus value.");
    }
}

void KiwoomApiWrapper::OnReceiveChejanDataHandler(QString sGubun, int nItemCnt, QString sFIdList) {
    emit postDebugMessage("TODO : OnReceiveChejanDataHandler()");
}


void KiwoomApiWrapper::OnReceiveConditionVerHandler(int lRet, QString sMsg) {
    emit postDebugMessage("TODO : OnReceiveConditionVerHandler()");
}


void KiwoomApiWrapper::OnReceiveInvestRealDataHandler(QString sRealKey) {
    emit postDebugMessage("TODO : OnReceiveInvestRealDataHandler()");
}


void KiwoomApiWrapper::OnReceiveMsgHandler(QString sScrNo, QString sRQName, QString sTrCode, QString sMsg) {
    emit postDebugMessage("TODO : OnReceiveMsgHandler()");
}


void KiwoomApiWrapper::OnReceiveRealConditionHandler(QString sTrCode, QString strType, QString strConditionName, QString strConditionIndex) {
    emit postDebugMessage("TODO : OnReceiveRealConditionHandler()");
}


void KiwoomApiWrapper::OnReceiveRealDataHandler(QString sRealKey, QString sRealType, QString sRealData) {
    emit postDebugMessage("TODO : OnReceiveRealDataHandler()");
}


void KiwoomApiWrapper::OnReceiveTrConditionHandler(QString sScrNo, QString strCodeList, QString strConditionName, int nIndex, int nNext) {
    emit postDebugMessage("TODO : OnReceiveTrConditionHandler()");
}


void KiwoomApiWrapper::OnReceiveTrDataHandler(QString sScrNo, QString sRQName, QString sTrCode, QString sRecordName, QString sPrevNext, int nDataLength, QString sErrorCode, QString sMessage, QString sSplmMsg) {
    emit postDebugMessage("TODO : OnReceiveTrDataHandler()");
}

void KiwoomApiWrapper::exceptionHandler(int code, QString source, QString disc, QString help) {
    emit postDebugMessage("TODO : exceptionHandler()");
}

void KiwoomApiWrapper::propertyChangedHandler(QString name) {
    emit postDebugMessage("TODO : propertyChangedHandler()");
}

void KiwoomApiWrapper::signalHandler(QString name, int argc, void* argv) {
    emit postDebugMessage("TODO : signalHandler()");
}
