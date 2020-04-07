#include "mainwindow.hpp"
#include "ui_mainwindow.h"
#include <QAxWidget>
#include <QVariant>

void commConnect();
void appendMessage(QString);

QAxWidget *kiwoom;
QLabel *lblMessage;

MainWindow::MainWindow(QWidget *parent) : QMainWindow(parent),
    ui(new Ui::MainWindow) {
    ui->setupUi(this);

    kiwoom = ui->kiwoom;
    lblMessage = ui->lblMessage;

    commConnect();
}

MainWindow::~MainWindow() {
    delete ui;
}

void commConnect() {
    QVariant result = kiwoom->dynamicCall("CommConnect()");

    if (result.toInt() == 0) {
        appendMessage("CommConnecet() OK");
    } else if (result.toInt() < 0) {
        appendMessage("CommConnect() Error");
    } else {
        appendMessage("CommConnect() Unexpected value.");
    }
}

void appendMessage(QString string) {
    QString msg = lblMessage->text();
    lblMessage->setText(msg.append(string) + "\n");
}

void MainWindow::on_kiwoom_OnEventConnect(int nErrCode) {
    if (nErrCode == 0) {
        appendMessage("Login Success.");
    } else if (nErrCode < 0) {
        appendMessage("Login Failed.");
    }
}
