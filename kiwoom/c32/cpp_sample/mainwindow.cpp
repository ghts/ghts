#include "mainwindow.hpp"
#include "ui_mainwindow.h"
#include "KiwoomApiWrapper.hpp"
#include <QString>

KiwoomApiWrapper *kiwoomApiWrapper;

MainWindow::MainWindow(QWidget *parent) : QMainWindow(parent), ui(new Ui::MainWindow) {
    ui->setupUi(this);

    kiwoomApiWrapper = new KiwoomApiWrapper(this, ui->kiwoom);

    connect(kiwoomApiWrapper, SIGNAL(sendMessage(QString)), this, SLOT(printMessage(QString)));

    int connectCallResult = kiwoomApiWrapper->CommConnect();

    if (connectCallResult == 0) {
        printMessage("CommConnecet() OK");
    } else if (connectCallResult < 0) {
        printMessage("CommConnect() Error");
    } else {
        printMessage("CommConnect() Unexpected value.");
    }
}

MainWindow::~MainWindow() {
    delete ui;
}

void MainWindow::on_kiwoom_OnEventConnect(int nErrCode) {
    kiwoomApiWrapper->OnEventConnect(nErrCode);
}

void MainWindow::on_kiwoom_OnReceiveChejanData(const QString &sGubun, int nItemCnt, const QString &sFIdList) {
    kiwoomApiWrapper->OnReceiveChejanData(sGubun, nItemCnt, sFIdList);
}

void MainWindow::on_kiwoom_OnReceiveConditionVer(int lRet, const QString &sMsg) {
    kiwoomApiWrapper->OnReceiveConditionVer(lRet, sMsg);
}

void MainWindow::on_kiwoom_OnReceiveInvestRealData(const QString &sRealKey) {
    kiwoomApiWrapper->OnReceiveInvestRealData(sRealKey);
}

void MainWindow::on_kiwoom_OnReceiveMsg(const QString &sScrNo, const QString &sRQName, const QString &sTrCode, const QString &sMsg) {
    kiwoomApiWrapper->OnReceiveMsg(sScrNo, sRQName, sTrCode, sMsg);
}

void MainWindow::on_kiwoom_OnReceiveRealCondition(const QString &sTrCode, const QString &strType, const QString &strConditionName, const QString &strConditionIndex) {
    kiwoomApiWrapper->OnReceiveRealCondition(sTrCode, strType, strConditionName, strConditionIndex);
}

void MainWindow::on_kiwoom_OnReceiveRealData(const QString &sRealKey, const QString &sRealType, const QString &sRealData) {
    kiwoomApiWrapper->OnReceiveRealData(sRealKey, sRealType, sRealData);
}

void MainWindow::on_kiwoom_OnReceiveTrCondition(const QString &sScrNo, const QString &strCodeList, const QString &strConditionName, int nIndex, int nNext) {
    kiwoomApiWrapper->OnReceiveTrCondition(sScrNo, strCodeList, strConditionName, nIndex, nNext);
}

void MainWindow::on_kiwoom_OnReceiveTrData(const QString &sScrNo, const QString &sRQName, const QString &sTrCode, const QString &sRecordName, const QString &sPrevNext, int nDataLength, const QString &sErrorCode, const QString &sMessage, const QString &sSplmMsg) {
    kiwoomApiWrapper->OnReceiveTrData(sScrNo, sRQName, sTrCode, sRecordName, sPrevNext, nDataLength, sErrorCode, sMessage, sSplmMsg);
}

// 헬퍼 슬롯
void MainWindow::printMessage(QString string) {
    QString msg = ui->lblMessage->text();
    ui->lblMessage->setText(msg.append(string) + "\n");
}
