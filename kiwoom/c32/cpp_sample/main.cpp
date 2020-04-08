#include "MainWindow.hpp"
#include "WinMsgHandler.hpp"
#include "KiwoomApiWrapper.hpp"

#include <QApplication>
#include <QAbstractEventDispatcher>

#include "PostConnectEvent.hpp"

// 1. Init Go DLL
// 2. Go DLL transfer msg by sendmessage in a separate transfer-only goroutine.
// 3. Install Event Filter to receive win32 msg from DLL. (DONE)

int main(int argc, char *argv[])
{
    QApplication a(argc, argv);

    MainWindow w;
    w.show();

    KiwoomApiWrapper *kiwoom = new KiwoomApiWrapper(&w, Qt::Widget);
    w.connect(kiwoom, SIGNAL(postDebugMessage(QString)), &w, SLOT(printDebugMessage(QString)));

    WinMsgHandler *filter = new WinMsgHandler();
    filter->setKiwoomApiWrapper(kiwoom);
    QAbstractEventDispatcher::instance()->installNativeEventFilter(filter);

    // PostMessage Task for Login test.
    // Task parented to the application so that it will be deleted by the application.
    PostConnectEvent *task = new PostConnectEvent(&a, w.winId());
    w.connect(task, SIGNAL(postDebugMessage(QString)), &w, SLOT(printDebugMessage(QString)));
    QTimer::singleShot(0, task, SLOT(run()));


    return a.exec();
}
