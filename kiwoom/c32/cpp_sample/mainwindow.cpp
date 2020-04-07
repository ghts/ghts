#include "MainWindow.hpp"
#include "ui_MainWindow.h"


MainWindow::MainWindow(QWidget *parent)
    : QMainWindow(parent), ui(new Ui::MainWindow) {
    ui->setupUi(this);
    kiwoom = new KHOpenAPILib::KHOpenAPI(this, Qt::Widget);
    connectSignalSlot();
    commConnect();
}

MainWindow::~MainWindow() {
    delete ui;
}

void MainWindow::connectSignalSlot() {
    connect(kiwoom, SIGNAL(OnEventConnect(int)), this, SLOT(OnEventConnectHandler(int)));
    connect(kiwoom, SIGNAL(OnReceiveChejanData(QString, int, QString)), this, SLOT(OnReceiveChejanDataHandler(QString, int, QString)));
    connect(kiwoom, SIGNAL(OnReceiveConditionVer(int, QString)), this, SLOT(OnReceiveConditionVerHandler(int, QString)));
    connect(kiwoom, SIGNAL(OnReceiveInvestRealData(QString)), this, SLOT(OnReceiveInvestRealDataHandler(QString)));
    connect(kiwoom, SIGNAL(OnReceiveMsg(QString, QString, QString, QString)), this, SLOT(OnReceiveMsgHandler(QString, QString, QString, QString)));
    connect(kiwoom, SIGNAL(OnReceiveRealCondition(QString, QString, QString, QString)), this, SLOT(OnReceiveRealConditionHandler(QString, QString, QString, QString)));
    connect(kiwoom, SIGNAL(OnReceiveRealData(QString, QString, QString)), this, SLOT(OnReceiveRealDataHandler(QString, QString, QString)));
    connect(kiwoom, SIGNAL(OnReceiveTrCondition(QString, QString, QString, int, int)), this, SLOT(OnReceiveTrConditionHandler(QString, QString, QString, int, int)));
    connect(kiwoom, SIGNAL(OnReceiveTrData(QString, QString, QString, QString, QString, int, QString, QString, QString)), this, SLOT(OnReceiveTrDataHandler(QString, QString, QString, QString, QString, int, QString, QString, QString)));
}

bool MainWindow::commConnect() {
    int result = kiwoom->CommConnect();

    if (result == 0) {
        printMessage("CommConnecet() OK");
    } else if (result < 0) {
        printMessage("CommConnect() Error");
    } else {
        printMessage("CommConnect() Unexpected value.");
    }

    return (result == 0);
}

void MainWindow::printMessage(QString string) {
    QString msg = ui->lblMessage->text();
    ui->lblMessage->setText(msg.append(string) + "\n");
}

void MainWindow::OnEventConnectHandler(int nErrCode) {
    if (nErrCode == 0) {
        printMessage("Login Success.");
    } else if (nErrCode < 0) {
        printMessage("Login Failed.");
    }
}

void MainWindow::OnReceiveChejanDataHandler(QString sGubun, int nItemCnt, QString sFIdList) {
    printMessage("TODO : OnReceiveChejanDataHandler()");
}

void MainWindow::OnReceiveConditionVerHandler(int lRet, QString sMsg) {
    printMessage("TODO : OnReceiveConditionVerHandler()");
}

void MainWindow::OnReceiveInvestRealDataHandler(QString sRealKey) {
    printMessage("TODO : OnReceiveInvestRealDataHandler()");
}

void MainWindow::OnReceiveMsgHandler(QString sScrNo, QString sRQName, QString sTrCode, QString sMsg) {
    printMessage("TODO : OnReceiveMsgHandler()");
}

void MainWindow::OnReceiveRealConditionHandler(QString sTrCode, QString strType, QString strConditionName, QString strConditionIndex) {
    printMessage("TODO : OnReceiveRealConditionHandler()");
}

void MainWindow::OnReceiveRealDataHandler(QString sRealKey, QString sRealType, QString sRealData) {
    printMessage("TODO : OnReceiveRealDataHandler()");
}

void MainWindow::OnReceiveTrConditionHandler(QString sScrNo, QString strCodeList, QString strConditionName, int nIndex, int nNext) {
    printMessage("TODO : OnReceiveTrConditionHandler()");
}

void MainWindow::OnReceiveTrDataHandler(QString sScrNo, QString sRQName, QString sTrCode, QString sRecordName, QString sPrevNext, int nDataLength, QString sErrorCode, QString sMessage, QString sSplmMsg) {
    printMessage("TODO : OnReceiveTrDataHandler()");
}
