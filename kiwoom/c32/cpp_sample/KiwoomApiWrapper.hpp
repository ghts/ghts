#include <QObject>
#include <QString>
#include <QVariant>
#include <QAxWidget>

class KiwoomApiWrapper : public QObject {
    Q_OBJECT
public:
    KiwoomApiWrapper(QObject *parent, QAxWidget *axWidget) : QObject(parent) {
        this->kiwoom = axWidget;
    }

    //~KiwoomApiWrapper(); // Base destructor is automatically called.

    inline int CommConnect();
    inline QString CommGetData(const QString& sJongmokCode, const QString& sRealType, const QString& sFieldName, int nIndex, const QString& sInnerFieldName);
    inline int CommInvestRqData(const QString& sMarketGb, const QString& sRQName, const QString& sScreenNo);
    inline int CommKwRqData(const QString& sArrCode, int bNext, int nCodeCount, int nTypeFlag, const QString& sRQName, const QString& sScreenNo);
    inline int CommRqData(const QString& sRQName, const QString& sTrCode, int nPrevNext, const QString& sScreenNo);
    inline void CommTerminate();
    inline void DisconnectRealData(const QString& sScnNo);
    inline QString GetAPIModulePath();
    inline QString GetActPriceList();
    inline QString GetBranchCodeName();
    inline QString GetChejanData(int nFid);
    inline QString GetCodeListByMarket(const QString& sMarket);
    inline QString GetCommData(const QString& strTrCode, const QString& strRecordName, int nIndex, const QString& strItemName);
    inline QVariant GetCommDataEx(const QString& strTrCode, const QString& strRecordName);
    inline QString GetCommRealData(const QString& sTrCode, int nFid);
    inline int GetConditionLoad();
    inline QString GetConditionNameList();
    inline int GetConnectState();
    inline int GetDataCount(const QString& strRecordName);
    inline QString GetFutureCodeByIndex(int nIndex);
    inline QString GetFutureList();
    inline QString GetLoginInfo(const QString& sTag);
    inline int GetMarketType(const QString& sTrCode);
    inline QString GetMasterCodeName(const QString& sTrCode);
    inline QString GetMasterConstruction(const QString& sTrCode);
    inline QString GetMasterLastPrice(const QString& sTrCode);
    inline int GetMasterListedStockCnt(const QString& sTrCode);
    inline QString GetMasterListedStockDate(const QString& sTrCode);
    inline QString GetMasterStockState(const QString& sTrCode);
    inline QString GetMonthList();
    inline QString GetOptionATM();
    inline QString GetOptionCode(const QString& strActPrice, int nCp, const QString& strMonth);
    inline QString GetOptionCodeByActPrice(const QString& sTrCode, int nCp, int nTick);
    inline QString GetOptionCodeByMonth(const QString& sTrCode, int nCp, const QString& strMonth);
    inline QString GetOutputValue(const QString& strRecordName, int nRepeatIdx, int nItemIdx);
    inline int GetRepeatCnt(const QString& sTrCode, const QString& sRecordName);
    inline QString GetSActPriceList(const QString& strBaseAssetGb);
    inline QString GetSFOBasisAssetList();
    inline QString GetSFutureCodeByIndex(const QString& strBaseAssetCode, int nIndex);
    inline QString GetSFutureList(const QString& strBaseAssetCode);
    inline QString GetSMonthList(const QString& strBaseAssetGb);
    inline QString GetSOptionATM(const QString& strBaseAssetGb);
    inline QString GetSOptionCode(const QString& strBaseAssetGb, const QString& strActPrice, int nCp, const QString& strMonth);
    inline QString GetSOptionCodeByActPrice(const QString& strBaseAssetGb, const QString& sTrCode, int nCp, int nTick);
    inline QString GetSOptionCodeByMonth(const QString& strBaseAssetGb, const QString& sTrCode, int nCp, const QString& strMonth);
    inline QString GetThemeGroupCode(const QString& strThemeCode);
    inline QString GetThemeGroupList(int nType);
    inline QString KOA_Functions(const QString& sFunctionName, const QString& sParam);
    inline int SendCondition(const QString& strScrNo, const QString& strConditionName, int nIndex, int nSearch);
    inline void SendConditionStop(const QString& strScrNo, const QString& strConditionName, int nIndex);
    inline int SendOrder(const QString& sRQName, const QString& sScreenNo, const QString& sAccNo, int nOrderType, const QString& sCode, int nQty, int nPrice, const QString& sHogaGb, const QString& sOrgOrderNo);
    inline int SendOrderCredit(const QString& sRQName, const QString& sScreenNo, const QString& sAccNo, int nOrderType, const QString& sCode, int nQty, int nPrice, const QString& sHogaGb, const QString& sCreditGb, const QString& sLoanDate, const QString& sOrgOrderNo);
    inline int SendOrderFO(const QString& sRQName, const QString& sScreenNo, const QString& sAccNo, const QString& sCode, int lOrdKind, const QString& sSlbyTp, const QString& sOrdTp, int lQty, const QString& sPrice, const QString& sOrgOrdNo);
    inline int SetInfoData(const QString& sInfoData);
    inline void SetInputValue(const QString& sID, const QString& sValue);
    inline int SetOutputFID(const QString& sID);
    inline int SetRealReg(const QString& strScreenNo, const QString& strCodeList, const QString& strFidList, const QString& strOptType);
    inline void SetRealRemove(const QString& strScrNo, const QString& strDelCode);

public slots:
    inline void OnEventConnect(int nErrCode);
    inline void OnReceiveChejanData(const QString &sGubun, int nItemCnt, const QString &sFIdList);
    inline void OnReceiveConditionVer(int lRet, const QString &sMsg);
    inline void OnReceiveInvestRealData(const QString &sRealKey);
    inline void OnReceiveMsg(const QString &sScrNo, const QString &sRQName, const QString &sTrCode, const QString &sMsg);
    inline void OnReceiveRealCondition(const QString &sTrCode, const QString &strType, const QString &strConditionName, const QString &strConditionIndex);
    inline void OnReceiveRealData(const QString &sRealKey, const QString &sRealType, const QString &sRealData);
    inline void OnReceiveTrCondition(const QString &sScrNo, const QString &strCodeList, const QString &strConditionName, int nIndex, int nNext);
    inline void OnReceiveTrData(const QString &sScrNo, const QString &sRQName, const QString &sTrCode, const QString &sRecordName, const QString &sPrevNext, int nDataLength, const QString &sErrorCode, const QString &sMessage, const QString &sSplmMsg);

signals:
    void sendMessage(QString message);

private:
    QAxWidget *kiwoom;
};

inline int KiwoomApiWrapper::CommConnect() {
    QVariant result = this->kiwoom->dynamicCall("CommConnect()");
    int resultInt = result.toInt();

    if (resultInt == 0) {
        sendMessage("CommConnecet() OK");
    } else if (resultInt < 0) {
        sendMessage("CommConnect() Error");
    } else {
        sendMessage("CommConnect() Unexpected value.");
    }

    return resultInt;
}

inline QString KiwoomApiWrapper::CommGetData(const QString& sJongmokCode, const QString& sRealType, const QString& sFieldName, int nIndex, const QString& sInnerFieldName) {
    QString qax_result;
    void *_a[] = {(void*)&qax_result, (void*)&sJongmokCode, (void*)&sRealType, (void*)&sFieldName, (void*)&nIndex, (void*)&sInnerFieldName};
    qt_metacall(QMetaObject::InvokeMetaMethod, 45, _a);
    return qax_result;
}

inline int KiwoomApiWrapper::CommInvestRqData(const QString& sMarketGb, const QString& sRQName, const QString& sScreenNo) {
    int qax_result;
    void *_a[] = {(void*)&qax_result, (void*)&sMarketGb, (void*)&sRQName, (void*)&sScreenNo};
    qt_metacall(QMetaObject::InvokeMetaMethod, 46, _a);
    return qax_result;
}

inline int KiwoomApiWrapper::CommKwRqData(const QString& sArrCode, int bNext, int nCodeCount, int nTypeFlag, const QString& sRQName, const QString& sScreenNo) {
    int qax_result;
    void *_a[] = {(void*)&qax_result, (void*)&sArrCode, (void*)&bNext, (void*)&nCodeCount, (void*)&nTypeFlag, (void*)&sRQName, (void*)&sScreenNo};
    qt_metacall(QMetaObject::InvokeMetaMethod, 11, _a);
    return qax_result;
}

inline int KiwoomApiWrapper::CommRqData(const QString& sRQName, const QString& sTrCode, int nPrevNext, const QString& sScreenNo) {
    int qax_result;
    void *_a[] = {(void*)&qax_result, (void*)&sRQName, (void*)&sTrCode, (void*)&nPrevNext, (void*)&sScreenNo};
    qt_metacall(QMetaObject::InvokeMetaMethod, 12, _a);
    return qax_result;
}

inline void KiwoomApiWrapper::CommTerminate() {
    void *_a[] = {0};
    qt_metacall(QMetaObject::InvokeMetaMethod, 49, _a);
}

inline void KiwoomApiWrapper::DisconnectRealData(const QString& sScnNo) {
    void *_a[] = {0, (void*)&sScnNo};
    qt_metacall(QMetaObject::InvokeMetaMethod, 50, _a);
}

inline QString KiwoomApiWrapper::GetAPIModulePath() {
    QString qax_result;
    void *_a[] = {(void*)&qax_result};
    qt_metacall(QMetaObject::InvokeMetaMethod, 51, _a);
    return qax_result;
}

inline QString KiwoomApiWrapper::GetActPriceList() {
    QString qax_result;
    void *_a[] = {(void*)&qax_result};
    qt_metacall(QMetaObject::InvokeMetaMethod, 52, _a);
    return qax_result;
}

inline QString KiwoomApiWrapper::GetBranchCodeName() {
    QString qax_result;
    void *_a[] = {(void*)&qax_result};
    qt_metacall(QMetaObject::InvokeMetaMethod, 53, _a);
    return qax_result;
}

inline QString KiwoomApiWrapper::GetChejanData(int nFid) {
    QString qax_result;
    void *_a[] = {(void*)&qax_result, (void*)&nFid};
    qt_metacall(QMetaObject::InvokeMetaMethod, 54, _a);
    return qax_result;
}

inline QString KiwoomApiWrapper::GetCodeListByMarket(const QString& sMarket) {
    QString qax_result;
    void *_a[] = {(void*)&qax_result, (void*)&sMarket};
    qt_metacall(QMetaObject::InvokeMetaMethod, 55, _a);
    return qax_result;
}

inline QString KiwoomApiWrapper::GetCommData(const QString& strTrCode, const QString& strRecordName, int nIndex, const QString& strItemName) {
    QString qax_result;
    void *_a[] = {(void*)&qax_result, (void*)&strTrCode, (void*)&strRecordName, (void*)&nIndex, (void*)&strItemName};
    qt_metacall(QMetaObject::InvokeMetaMethod, 56, _a);
    return qax_result;
}

inline QVariant KiwoomApiWrapper::GetCommDataEx(const QString& strTrCode, const QString& strRecordName) {
    QVariant qax_result;
    void *_a[] = {(void*)&qax_result, (void*)&strTrCode, (void*)&strRecordName};
    qt_metacall(QMetaObject::InvokeMetaMethod, 21, _a);
    return qax_result;
}

inline QString KiwoomApiWrapper::GetCommRealData(const QString& sTrCode, int nFid) {
    QString qax_result;
    void *_a[] = {(void*)&qax_result, (void*)&sTrCode, (void*)&nFid};
    qt_metacall(QMetaObject::InvokeMetaMethod, 58, _a);
    return qax_result;
}

inline int KiwoomApiWrapper::GetConditionLoad() {
    int qax_result;
    void *_a[] = {(void*)&qax_result};
    qt_metacall(QMetaObject::InvokeMetaMethod, 59, _a);
    return qax_result;
}

inline QString KiwoomApiWrapper::GetConditionNameList() {
    QString qax_result;
    void *_a[] = {(void*)&qax_result};
    qt_metacall(QMetaObject::InvokeMetaMethod, 60, _a);
    return qax_result;
}

inline int KiwoomApiWrapper::GetConnectState() {
    int qax_result;
    void *_a[] = {(void*)&qax_result};
    qt_metacall(QMetaObject::InvokeMetaMethod, 61, _a);
    return qax_result;
}

inline int KiwoomApiWrapper::GetDataCount(const QString& strRecordName) {
    int qax_result;
    void *_a[] = {(void*)&qax_result, (void*)&strRecordName};
    qt_metacall(QMetaObject::InvokeMetaMethod, 62, _a);
    return qax_result;
}

inline QString KiwoomApiWrapper::GetFutureCodeByIndex(int nIndex) {
    QString qax_result;
    void *_a[] = {(void*)&qax_result, (void*)&nIndex};
    qt_metacall(QMetaObject::InvokeMetaMethod, 63, _a);
    return qax_result;
}

inline QString KiwoomApiWrapper::GetFutureList() {
    QString qax_result;
    void *_a[] = {(void*)&qax_result};
    qt_metacall(QMetaObject::InvokeMetaMethod, 64, _a);
    return qax_result;
}

inline QString KiwoomApiWrapper::GetLoginInfo(const QString& sTag) {
    QString qax_result;
    void *_a[] = {(void*)&qax_result, (void*)&sTag};
    qt_metacall(QMetaObject::InvokeMetaMethod, 65, _a);
    return qax_result;
}

inline int KiwoomApiWrapper::GetMarketType(const QString& sTrCode) {
    int qax_result;
    void *_a[] = {(void*)&qax_result, (void*)&sTrCode};
    qt_metacall(QMetaObject::InvokeMetaMethod, 66, _a);
    return qax_result;
}

inline QString KiwoomApiWrapper::GetMasterCodeName(const QString& sTrCode) {
    QString qax_result;
    void *_a[] = {(void*)&qax_result, (void*)&sTrCode};
    qt_metacall(QMetaObject::InvokeMetaMethod, 67, _a);
    return qax_result;
}

inline QString KiwoomApiWrapper::GetMasterConstruction(const QString& sTrCode) {
    QString qax_result;
    void *_a[] = {(void*)&qax_result, (void*)&sTrCode};
    qt_metacall(QMetaObject::InvokeMetaMethod, 68, _a);
    return qax_result;
}

inline QString KiwoomApiWrapper::GetMasterLastPrice(const QString& sTrCode) {
    QString qax_result;
    void *_a[] = {(void*)&qax_result, (void*)&sTrCode};
    qt_metacall(QMetaObject::InvokeMetaMethod, 69, _a);
    return qax_result;
}

inline int KiwoomApiWrapper::GetMasterListedStockCnt(const QString& sTrCode) {
    int qax_result;
    void *_a[] = {(void*)&qax_result, (void*)&sTrCode};
    qt_metacall(QMetaObject::InvokeMetaMethod, 70, _a);
    return qax_result;
}

inline QString KiwoomApiWrapper::GetMasterListedStockDate(const QString& sTrCode) {
    QString qax_result;
    void *_a[] = {(void*)&qax_result, (void*)&sTrCode};
    qt_metacall(QMetaObject::InvokeMetaMethod, 71, _a);
    return qax_result;
}

inline QString KiwoomApiWrapper::GetMasterStockState(const QString& sTrCode) {
    QString qax_result;
    void *_a[] = {(void*)&qax_result, (void*)&sTrCode};
    qt_metacall(QMetaObject::InvokeMetaMethod, 72, _a);
    return qax_result;
}

inline QString KiwoomApiWrapper::GetMonthList() {
    QString qax_result;
    void *_a[] = {(void*)&qax_result};
    qt_metacall(QMetaObject::InvokeMetaMethod, 73, _a);
    return qax_result;
}

inline QString KiwoomApiWrapper::GetOptionATM() {
    QString qax_result;
    void *_a[] = {(void*)&qax_result};
    qt_metacall(QMetaObject::InvokeMetaMethod, 74, _a);
    return qax_result;
}

inline QString KiwoomApiWrapper::GetOptionCode(const QString& strActPrice, int nCp, const QString& strMonth) {
    QString qax_result;
    void *_a[] = {(void*)&qax_result, (void*)&strActPrice, (void*)&nCp, (void*)&strMonth};
    qt_metacall(QMetaObject::InvokeMetaMethod, 75, _a);
    return qax_result;
}

inline QString KiwoomApiWrapper::GetOptionCodeByActPrice(const QString& sTrCode, int nCp, int nTick) {
    QString qax_result;
    void *_a[] = {(void*)&qax_result, (void*)&sTrCode, (void*)&nCp, (void*)&nTick};
    qt_metacall(QMetaObject::InvokeMetaMethod, 76, _a);
    return qax_result;
}

inline QString KiwoomApiWrapper::GetOptionCodeByMonth(const QString& sTrCode, int nCp, const QString& strMonth)
{
    QString qax_result;
    void *_a[] = {(void*)&qax_result, (void*)&sTrCode, (void*)&nCp, (void*)&strMonth};
    qt_metacall(QMetaObject::InvokeMetaMethod, 77, _a);
    return qax_result;
}

inline QString KiwoomApiWrapper::GetOutputValue(const QString& strRecordName, int nRepeatIdx, int nItemIdx)
{
    QString qax_result;
    void *_a[] = {(void*)&qax_result, (void*)&strRecordName, (void*)&nRepeatIdx, (void*)&nItemIdx};
    qt_metacall(QMetaObject::InvokeMetaMethod, 78, _a);
    return qax_result;
}

inline int KiwoomApiWrapper::GetRepeatCnt(const QString& sTrCode, const QString& sRecordName)
{
    int qax_result;
    void *_a[] = {(void*)&qax_result, (void*)&sTrCode, (void*)&sRecordName};
    qt_metacall(QMetaObject::InvokeMetaMethod, 79, _a);
    return qax_result;
}

inline QString KiwoomApiWrapper::GetSActPriceList(const QString& strBaseAssetGb) {
    QString qax_result;
    void *_a[] = {(void*)&qax_result, (void*)&strBaseAssetGb};
    qt_metacall(QMetaObject::InvokeMetaMethod, 80, _a);
    return qax_result;
}

inline QString KiwoomApiWrapper::GetSFOBasisAssetList() {
    QString qax_result;
    void *_a[] = {(void*)&qax_result};
    qt_metacall(QMetaObject::InvokeMetaMethod, 81, _a);
    return qax_result;
}

inline QString KiwoomApiWrapper::GetSFutureCodeByIndex(const QString& strBaseAssetCode, int nIndex) {
    QString qax_result;
    void *_a[] = {(void*)&qax_result, (void*)&strBaseAssetCode, (void*)&nIndex};
    qt_metacall(QMetaObject::InvokeMetaMethod, 82, _a);
    return qax_result;
}

inline QString KiwoomApiWrapper::GetSFutureList(const QString& strBaseAssetCode) {
    QString qax_result;
    void *_a[] = {(void*)&qax_result, (void*)&strBaseAssetCode};
    qt_metacall(QMetaObject::InvokeMetaMethod, 83, _a);
    return qax_result;
}

inline QString KiwoomApiWrapper::GetSMonthList(const QString& strBaseAssetGb) {
    QString qax_result;
    void *_a[] = {(void*)&qax_result, (void*)&strBaseAssetGb};
    qt_metacall(QMetaObject::InvokeMetaMethod, 84, _a);
    return qax_result;
}

inline QString KiwoomApiWrapper::GetSOptionATM(const QString& strBaseAssetGb) {
    QString qax_result;
    void *_a[] = {(void*)&qax_result, (void*)&strBaseAssetGb};
    qt_metacall(QMetaObject::InvokeMetaMethod, 85, _a);
    return qax_result;
}

inline QString KiwoomApiWrapper::GetSOptionCode(const QString& strBaseAssetGb, const QString& strActPrice, int nCp, const QString& strMonth) {
    QString qax_result;
    void *_a[] = {(void*)&qax_result, (void*)&strBaseAssetGb, (void*)&strActPrice, (void*)&nCp, (void*)&strMonth};
    qt_metacall(QMetaObject::InvokeMetaMethod, 86, _a);
    return qax_result;
}

inline QString KiwoomApiWrapper::GetSOptionCodeByActPrice(const QString& strBaseAssetGb, const QString& sTrCode, int nCp, int nTick) {
    QString qax_result;
    void *_a[] = {(void*)&qax_result, (void*)&strBaseAssetGb, (void*)&sTrCode, (void*)&nCp, (void*)&nTick};
    qt_metacall(QMetaObject::InvokeMetaMethod, 87, _a);
    return qax_result;
}

inline QString KiwoomApiWrapper::GetSOptionCodeByMonth(const QString& strBaseAssetGb, const QString& sTrCode, int nCp, const QString& strMonth) {
    QString qax_result;
    void *_a[] = {(void*)&qax_result, (void*)&strBaseAssetGb, (void*)&sTrCode, (void*)&nCp, (void*)&strMonth};
    qt_metacall(QMetaObject::InvokeMetaMethod, 88, _a);
    return qax_result;
}

inline QString KiwoomApiWrapper::GetThemeGroupCode(const QString& strThemeCode) {
    QString qax_result;
    void *_a[] = {(void*)&qax_result, (void*)&strThemeCode};
    qt_metacall(QMetaObject::InvokeMetaMethod, 89, _a);
    return qax_result;
}

inline QString KiwoomApiWrapper::GetThemeGroupList(int nType) {
    QString qax_result;
    void *_a[] = {(void*)&qax_result, (void*)&nType};
    qt_metacall(QMetaObject::InvokeMetaMethod, 90, _a);
    return qax_result;
}

inline QString KiwoomApiWrapper::KOA_Functions(const QString& sFunctionName, const QString& sParam) {
    QString qax_result;
    void *_a[] = {(void*)&qax_result, (void*)&sFunctionName, (void*)&sParam};
    qt_metacall(QMetaObject::InvokeMetaMethod, 91, _a);
    return qax_result;
}

inline int KiwoomApiWrapper::SendCondition(const QString& strScrNo, const QString& strConditionName, int nIndex, int nSearch) {
    int qax_result;
    void *_a[] = {(void*)&qax_result, (void*)&strScrNo, (void*)&strConditionName, (void*)&nIndex, (void*)&nSearch};
    qt_metacall(QMetaObject::InvokeMetaMethod, 92, _a);
    return qax_result;
}

inline void KiwoomApiWrapper::SendConditionStop(const QString& strScrNo, const QString& strConditionName, int nIndex) {
    void *_a[] = {0, (void*)&strScrNo, (void*)&strConditionName, (void*)&nIndex};
    qt_metacall(QMetaObject::InvokeMetaMethod, 93, _a);
}

inline int KiwoomApiWrapper::SendOrder(const QString& sRQName, const QString& sScreenNo, const QString& sAccNo, int nOrderType, const QString& sCode, int nQty, int nPrice, const QString& sHogaGb, const QString& sOrgOrderNo) {
    int qax_result;
    void *_a[] = {(void*)&qax_result, (void*)&sRQName, (void*)&sScreenNo, (void*)&sAccNo, (void*)&nOrderType, (void*)&sCode, (void*)&nQty, (void*)&nPrice, (void*)&sHogaGb, (void*)&sOrgOrderNo};
    qt_metacall(QMetaObject::InvokeMetaMethod, 94, _a);
    return qax_result;
}

inline int KiwoomApiWrapper::SendOrderCredit(const QString& sRQName, const QString& sScreenNo, const QString& sAccNo, int nOrderType, const QString& sCode, int nQty, int nPrice, const QString& sHogaGb, const QString& sCreditGb, const QString& sLoanDate, const QString& sOrgOrderNo) {
    int qax_result;
    void *_a[] = {(void*)&qax_result, (void*)&sRQName, (void*)&sScreenNo, (void*)&sAccNo, (void*)&nOrderType, (void*)&sCode, (void*)&nQty, (void*)&nPrice, (void*)&sHogaGb, (void*)&sCreditGb, (void*)&sLoanDate, (void*)&sOrgOrderNo};
    qt_metacall(QMetaObject::InvokeMetaMethod, 95, _a);
    return qax_result;
}

inline int KiwoomApiWrapper::SendOrderFO(const QString& sRQName, const QString& sScreenNo, const QString& sAccNo, const QString& sCode, int lOrdKind, const QString& sSlbyTp, const QString& sOrdTp, int lQty, const QString& sPrice, const QString& sOrgOrdNo) {
    int qax_result;
    void *_a[] = {(void*)&qax_result, (void*)&sRQName, (void*)&sScreenNo, (void*)&sAccNo, (void*)&sCode, (void*)&lOrdKind, (void*)&sSlbyTp, (void*)&sOrdTp, (void*)&lQty, (void*)&sPrice, (void*)&sOrgOrdNo};
    qt_metacall(QMetaObject::InvokeMetaMethod, 96, _a);
    return qax_result;
}

inline int KiwoomApiWrapper::SetInfoData(const QString& sInfoData) {
    int qax_result;
    void *_a[] = {(void*)&qax_result, (void*)&sInfoData};
    qt_metacall(QMetaObject::InvokeMetaMethod, 97, _a);
    return qax_result;
}

inline void KiwoomApiWrapper::SetInputValue(const QString& sID, const QString& sValue) {
    void *_a[] = {0, (void*)&sID, (void*)&sValue};
    qt_metacall(QMetaObject::InvokeMetaMethod, 98, _a);
}

inline int KiwoomApiWrapper::SetOutputFID(const QString& sID) {
    int qax_result;
    void *_a[] = {(void*)&qax_result, (void*)&sID};
    qt_metacall(QMetaObject::InvokeMetaMethod, 99, _a);
    return qax_result;
}

inline int KiwoomApiWrapper::SetRealReg(const QString& strScreenNo, const QString& strCodeList, const QString& strFidList, const QString& strOptType) {
    int qax_result;
    void *_a[] = {(void*)&qax_result, (void*)&strScreenNo, (void*)&strCodeList, (void*)&strFidList, (void*)&strOptType};
    qt_metacall(QMetaObject::InvokeMetaMethod, 100, _a);
    return qax_result;
}

inline void KiwoomApiWrapper::SetRealRemove(const QString& strScrNo, const QString& strDelCode) {
    void *_a[] = {0, (void*)&strScrNo, (void*)&strDelCode};
    qt_metacall(QMetaObject::InvokeMetaMethod, 101, _a);
}

inline void KiwoomApiWrapper::OnEventConnect(int nErrCode) {
    if (nErrCode == 0) {
        emit sendMessage("Login Success.");
    } else if (nErrCode < 0) {
        emit sendMessage("Login Failed.");
    }
}

inline void KiwoomApiWrapper::OnReceiveChejanData(const QString &sGubun, int nItemCnt, const QString &sFIdList) {
    emit sendMessage("OnReceiveChejanData() : TODO");
}

inline void KiwoomApiWrapper::OnReceiveConditionVer(int lRet, const QString &sMsg) {
    emit sendMessage("OnReceiveConditionVer() : TODO");
}

inline void KiwoomApiWrapper::OnReceiveInvestRealData(const QString &sRealKey) {
    emit sendMessage("OnReceiveInvestRealData() : TODO");
}

inline void KiwoomApiWrapper::OnReceiveMsg(const QString &sScrNo, const QString &sRQName, const QString &sTrCode, const QString &sMsg) {
    emit sendMessage("OnReceiveMsg() : TODO");
}

inline void KiwoomApiWrapper::OnReceiveRealCondition(const QString &sTrCode, const QString &strType, const QString &strConditionName, const QString &strConditionIndex) {
    emit sendMessage("OnReceiveRealCondition() : TODO");
}

inline void KiwoomApiWrapper::OnReceiveRealData(const QString &sRealKey, const QString &sRealType, const QString &sRealData) {
    emit sendMessage("OnReceiveRealData() : TODO");
}

inline void KiwoomApiWrapper::OnReceiveTrCondition(const QString &sScrNo, const QString &strCodeList, const QString &strConditionName, int nIndex, int nNext) {
    emit sendMessage("OnReceiveTrCondition() : TODO");
}

inline void KiwoomApiWrapper::OnReceiveTrData(const QString &sScrNo, const QString &sRQName, const QString &sTrCode, const QString &sRecordName, const QString &sPrevNext, int nDataLength, const QString &sErrorCode, const QString &sMessage, const QString &sSplmMsg) {
    emit sendMessage("OnReceiveTrData() : TODO");
}
