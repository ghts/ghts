#ifndef MAINWINDOW_HPP
#define MAINWINDOW_HPP

#include <QMainWindow>
#include "kiwoom.h"

QT_BEGIN_NAMESPACE
namespace Ui { class MainWindow; }
QT_END_NAMESPACE

class MainWindow : public QMainWindow
{
    Q_OBJECT

public:
    MainWindow(QWidget *parent = nullptr);
    ~MainWindow();

public slots:
    bool commConnect();
    void printMessage(QString message);

    // 이벤트 처리 슬롯
    void OnEventConnectHandler(int nErrCode);
    void OnReceiveChejanDataHandler(QString sGubun, int nItemCnt, QString sFIdList);
    void OnReceiveConditionVerHandler(int lRet, QString sMsg);
    void OnReceiveInvestRealDataHandler(QString sRealKey);
    void OnReceiveMsgHandler(QString sScrNo, QString sRQName, QString sTrCode, QString sMsg);
    void OnReceiveRealConditionHandler(QString sTrCode, QString strType, QString strConditionName, QString strConditionIndex);
    void OnReceiveRealDataHandler(QString sRealKey, QString sRealType, QString sRealData);
    void OnReceiveTrConditionHandler(QString sScrNo, QString strCodeList, QString strConditionName, int nIndex, int nNext);
    void OnReceiveTrDataHandler(QString sScrNo, QString sRQName, QString sTrCode, QString sRecordName, QString sPrevNext, int nDataLength, QString sErrorCode, QString sMessage, QString sSplmMsg);

private:
    Ui::MainWindow *ui;
    KHOpenAPILib::KHOpenAPI *kiwoom;

    void connectSignalSlot();
};
#endif // MAINWINDOW_HPP
