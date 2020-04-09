#ifndef KIWOOMAPIWRAPPER_HPP
#define KIWOOMAPIWRAPPER_HPP

#include "KiwoomApi.hpp"

// dumpcpp로 자동 생성된 KiwoomApi.[c/h] 파일을 건드리지 않고,
// 이벤트 처리를 위해서는 시그널/슬롯을 설정하는 래퍼.
class KiwoomApiWrapper : public QObject {
    Q_OBJECT

public:
    KiwoomApiWrapper(QWidget *parent = 0, Qt::WindowFlags f = 0) : QObject(parent) {
        kiwoom = new KHOpenAPILib::KHOpenAPI(parent, f);
        connectEventSignalSlot();
    }

    int CommConnect();
    QString CommGetData(const QString& sJongmokCode, const QString& sRealType, const QString& sFieldName, int nIndex, const QString& sInnerFieldName);
    int CommInvestRqData(const QString& sMarketGb, const QString& sRQName, const QString& sScreenNo);
    int CommKwRqData(const QString& sArrCode, int bNext, int nCodeCount, int nTypeFlag, const QString& sRQName, const QString& sScreenNo);
    int CommRqData(const QString& sRQName, const QString& sTrCode, int nPrevNext, const QString& sScreenNo);
    void CommTerminate();
    void DisconnectRealData(const QString& sScnNo);
    QString GetAPIModulePath();
    QString GetActPriceList();
    QString GetBranchCodeName();
    QString GetChejanData(int nFid);
    QString GetCodeListByMarket(const QString& sMarket);
    QString GetCommData(const QString& strTrCode, const QString& strRecordName, int nIndex, const QString& strItemName);
    QVariant GetCommDataEx(const QString& strTrCode, const QString& strRecordName);
    QString GetCommRealData(const QString& sTrCode, int nFid);
    int GetConditionLoad();
    QString GetConditionNameList();
    int GetConnectState();
    int GetDataCount(const QString& strRecordName);
    QString GetFutureCodeByIndex(int nIndex);
    QString GetFutureList();
    QString GetLoginInfo(const QString& sTag);
    int GetMarketType(const QString& sTrCode);
    QString GetMasterCodeName(const QString& sTrCode);
    QString GetMasterConstruction(const QString& sTrCode);
    QString GetMasterLastPrice(const QString& sTrCode);
    int GetMasterListedStockCnt(const QString& sTrCode);
    QString GetMasterListedStockDate(const QString& sTrCode);
    QString GetMasterStockState(const QString& sTrCode);
    QString GetMonthList();
    QString GetOptionATM();
    QString GetOptionCode(const QString& strActPrice, int nCp, const QString& strMonth);
    QString GetOptionCodeByActPrice(const QString& sTrCode, int nCp, int nTick);
    QString GetOptionCodeByMonth(const QString& sTrCode, int nCp, const QString& strMonth);
    QString GetOutputValue(const QString& strRecordName, int nRepeatIdx, int nItemIdx);
    int GetRepeatCnt(const QString& sTrCode, const QString& sRecordName);
    QString GetSActPriceList(const QString& strBaseAssetGb);
    QString GetSFOBasisAssetList();
    QString GetSFutureCodeByIndex(const QString& strBaseAssetCode, int nIndex);
    QString GetSFutureList(const QString& strBaseAssetCode);
    QString GetSMonthList(const QString& strBaseAssetGb);
    QString GetSOptionATM(const QString& strBaseAssetGb);
    QString GetSOptionCode(const QString& strBaseAssetGb, const QString& strActPrice, int nCp, const QString& strMonth);
    QString GetSOptionCodeByActPrice(const QString& strBaseAssetGb, const QString& sTrCode, int nCp, int nTick);
    QString GetSOptionCodeByMonth(const QString& strBaseAssetGb, const QString& sTrCode, int nCp, const QString& strMonth);
    QString GetThemeGroupCode(const QString& strThemeCode);
    QString GetThemeGroupList(int nType);
    QString KOA_Functions(const QString& sFunctionName, const QString& sParam);
    int SendCondition(const QString& strScrNo, const QString& strConditionName, int nIndex, int nSearch);
    void SendConditionStop(const QString& strScrNo, const QString& strConditionName, int nIndex);
    int SendOrder(const QString& sRQName, const QString& sScreenNo, const QString& sAccNo, int nOrderType, const QString& sCode, int nQty, int nPrice, const QString& sHogaGb, const QString& sOrgOrderNo);
    int SendOrderCredit(const QString& sRQName, const QString& sScreenNo, const QString& sAccNo, int nOrderType, const QString& sCode, int nQty, int nPrice, const QString& sHogaGb, const QString& sCreditGb, const QString& sLoanDate, const QString& sOrgOrderNo);
    int SendOrderFO(const QString& sRQName, const QString& sScreenNo, const QString& sAccNo, const QString& sCode, int lOrdKind, const QString& sSlbyTp, const QString& sOrdTp, int lQty, const QString& sPrice, const QString& sOrgOrdNo);
    int SetInfoData(const QString& sInfoData);
    void SetInputValue(const QString& sID, const QString& sValue);
    int SetOutputFID(const QString& sID);
    int SetRealReg(const QString& (strScreenNo), const QString& strCodeList, const QString& strFidList, const QString& strOptType);
    void SetRealRemove(const QString& strScrNo, const QString& strDelCode);

public slots:
    // 이벤트 핸들러 슬롯
    void OnEventConnectHandler(int nErrCode);
    void OnReceiveChejanDataHandler(QString sGubun, int nItemCnt, QString sFIdList);
    void OnReceiveConditionVerHandler(int lRet, QString sMsg);
    void OnReceiveInvestRealDataHandler(QString sRealKey);
    void OnReceiveMsgHandler(QString sScrNo, QString sRQName, QString sTrCode, QString sMsg);
    void OnReceiveRealConditionHandler(QString sTrCode, QString strType, QString strConditionName, QString strConditionIndex);
    void OnReceiveRealDataHandler(QString sRealKey, QString sRealType, QString sRealData);
    void OnReceiveTrConditionHandler(QString sScrNo, QString strCodeList, QString strConditionName, int nIndex, int nNext);
    void OnReceiveTrDataHandler(QString sScrNo, QString sRQName, QString sTrCode, QString sRecordName, QString sPrevNext, int nDataLength, QString sErrorCode, QString sMessage, QString sSplmMsg);
    void exceptionHandler(int code, QString source, QString disc, QString help);
    void propertyChangedHandler(QString name);
    void signalHandler(QString name, int argc, void* argv);

private:
    KHOpenAPILib::KHOpenAPI *kiwoom;
    void connectEventSignalSlot();

};

#endif // KIWOOMAPIWRAPPER_HPP
