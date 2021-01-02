#ifndef KIWOOMAPIWRAPPER_HPP
#define KIWOOMAPIWRAPPER_HPP

#include <QObject>
#include <QWidget>

// dumpcpp로 자동 생성된 파일에는 이벤트 처리 기능이 빠져 있으며,
// 자동 생성된 파일을 수정하면 dumpcpp 재실행 후
// 수정 사항이 모두 사라지므로,
// 이벤트 처리를 하는 시그널/슬롯을 설정하는 기능을 독립적으로 분리시킨 클래스.
class KiwoomEvents : public QObject {
    Q_OBJECT

public:
    KiwoomEvents(QWidget *parent = 0) : QObject(parent) {}

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
};

#endif // KIWOOMAPIWRAPPER_HPP
