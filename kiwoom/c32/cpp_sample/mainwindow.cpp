#include "MainWindow.hpp"
#include "ui_MainWindow.h"

MainWindow::MainWindow(QWidget *parent)
    : QMainWindow(parent), ui(new Ui::MainWindow) {
    ui->setupUi(this);
    ui->lblMessage->setText("");
}

MainWindow::~MainWindow() {
    delete ui;
}


void MainWindow::printDebugMessage(QString string) {
    QString msg = ui->lblMessage->text();
    ui->lblMessage->setText(msg.append(string) + "\n");
}

