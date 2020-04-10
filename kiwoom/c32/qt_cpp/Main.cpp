#include "Kiwoom.hpp"
#include "KiwoomEvents.hpp"
#include "WinMsgHandler.hpp"
#include "Func.hpp"

#include <QApplication>
#include <QAbstractEventDispatcher>
#include <QDebug>

void connectKiwoomEventSignalSlot(KHOpenAPILib::KHOpenAPI*, KiwoomEvents*);

int main(int argc, char *argv[])
{
    QApplication a(argc, argv);
    QWidget w;  // HWND를 제공하는 Win32 MSG 통로.

    KHOpenAPILib::KHOpenAPI *kiwoom = new KHOpenAPILib::KHOpenAPI(&w, Qt::Widget);
    KiwoomEvents *kiwoomEvents = new KiwoomEvents(&w);
    connectKiwoomEventSignalSlot(kiwoom, kiwoomEvents);

    WinMsgHandler *msgHandler = new WinMsgHandler();
    msgHandler->setKiwoom(kiwoom);
    QAbstractEventDispatcher::instance()->installNativeEventFilter(msgHandler);

    Init(&w);

    return a.exec();
}

void connectKiwoomEventSignalSlot(KHOpenAPILib::KHOpenAPI *kiwoom, KiwoomEvents *kiwoomEvents) {
    kiwoomEvents->connect(kiwoom, SIGNAL(OnEventConnect(int)), kiwoomEvents,  SLOT(OnEventConnectHandler(int)));
    kiwoomEvents->connect(kiwoom, SIGNAL(OnReceiveChejanData(QString, int, QString)), kiwoomEvents,  SLOT(OnReceiveChejanDataHandler(QString, int, QString)));
    kiwoomEvents->connect(kiwoom, SIGNAL(OnReceiveConditionVer(int, QString)), kiwoomEvents,  SLOT(OnReceiveConditionVerHandler(int, QString)));
    kiwoomEvents->connect(kiwoom, SIGNAL(OnReceiveInvestRealData(QString)), kiwoomEvents,  SLOT(OnReceiveInvestRealDataHandler(QString)));
    kiwoomEvents->connect(kiwoom, SIGNAL(OnReceiveMsg(QString, QString, QString, QString)), kiwoomEvents,  SLOT(OnReceiveMsgHandler(QString, QString, QString, QString)));
    kiwoomEvents->connect(kiwoom, SIGNAL(OnReceiveRealCondition(QString, QString, QString, QString)), kiwoomEvents,  SLOT(OnReceiveRealConditionHandler(QString, QString, QString, QString)));
    kiwoomEvents->connect(kiwoom, SIGNAL(OnReceiveRealData(QString, QString, QString)), kiwoomEvents,  SLOT(OnReceiveRealDataHandler(QString, QString, QString)));
    kiwoomEvents->connect(kiwoom, SIGNAL(OnReceiveTrCondition(QString, QString, QString, int, int)), kiwoomEvents,  SLOT(OnReceiveTrConditionHandler(QString, QString, QString, int, int)));
    kiwoomEvents->connect(kiwoom, SIGNAL(OnReceiveTrData(QString, QString, QString, QString, QString, int, QString, QString, QString)), kiwoomEvents,  SLOT(OnReceiveTrDataHandler(QString, QString, QString, QString, QString, int, QString, QString, QString)));
    kiwoomEvents->connect(kiwoom, SIGNAL(exception(int, QString, QString, QString)), kiwoomEvents,  SLOT(exceptionHandler(int, QString, QString, QString)));
    kiwoomEvents->connect(kiwoom, SIGNAL(propertyChanged(QString)), kiwoomEvents,  SLOT(propertyChangedHandler(QString)));
    kiwoomEvents->connect(kiwoom, SIGNAL(signal(QString, int, void*)), kiwoomEvents,  SLOT(signalHandler(QString, int, void*)));
}
