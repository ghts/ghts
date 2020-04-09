#include "WinMsgHandler.hpp"
#include "KiwoomApiWrapper.hpp"
#include "func.hpp"

#include <QApplication>
#include <QAbstractEventDispatcher>
#include <QDebug>

int main(int argc, char *argv[])
{
    QApplication a(argc, argv);
    QWidget w;  // HWND를 제공하는 Win32 MSG 통로.
    KiwoomApiWrapper *kiwoom = new KiwoomApiWrapper(&w, Qt::Widget);
    WinMsgHandler *msgHandler = new WinMsgHandler();
    msgHandler->setKiwoomApiWrapper(kiwoom);
    QAbstractEventDispatcher::instance()->installNativeEventFilter(msgHandler);
    InitGoDLL(&w);

    return a.exec();
}
