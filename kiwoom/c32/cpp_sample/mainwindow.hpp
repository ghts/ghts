#ifndef MAINWINDOW_HPP
#define MAINWINDOW_HPP

#include <QMainWindow>

QT_BEGIN_NAMESPACE
namespace Ui { class MainWindow; }
QT_END_NAMESPACE

class MainWindow : public QMainWindow {
    Q_OBJECT

public:
    MainWindow(QWidget *parent = nullptr);
    ~MainWindow();

public slots:
    void printMessage(QString message);

private slots:
    void on_kiwoom_OnEventConnect(int nErrCode);
    void on_kiwoom_OnReceiveChejanData(const QString &sGubun, int nItemCnt, const QString &sFIdList);
    void on_kiwoom_OnReceiveConditionVer(int lRet, const QString &sMsg);
    void on_kiwoom_OnReceiveInvestRealData(const QString &sRealKey);
    void on_kiwoom_OnReceiveMsg(const QString &sScrNo, const QString &sRQName, const QString &sTrCode, const QString &sMsg);
    void on_kiwoom_OnReceiveRealCondition(const QString &sTrCode, const QString &strType, const QString &strConditionName, const QString &strConditionIndex);
    void on_kiwoom_OnReceiveRealData(const QString &sRealKey, const QString &sRealType, const QString &sRealData);
    void on_kiwoom_OnReceiveTrCondition(const QString &sScrNo, const QString &strCodeList, const QString &strConditionName, int nIndex, int nNext);
    void on_kiwoom_OnReceiveTrData(const QString &sScrNo, const QString &sRQName, const QString &sTrCode, const QString &sRecordName, const QString &sPrevNext, int nDataLength, const QString &sErrorCode, const QString &sMessage, const QString &sSplmMsg);

private:
    Ui::MainWindow *ui;
};
#endif // MAINWINDOW_HPP
