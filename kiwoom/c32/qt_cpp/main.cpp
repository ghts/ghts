#include "MainWindow.hpp"
#include "WinMsgHandler.hpp"
#include "KiwoomApiWrapper.hpp"

#include <QApplication>
#include <QAbstractEventDispatcher>
#include <QLibrary>
#include <QDebug>
#include <QDir>
#include <QFile>
#include <QDirIterator>

//#include "PostMessageTask.hpp"

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

    // PostMessageTask for Test.
    // Task parented to the application so that it will be deleted by the application.
    //PostMessageTask *task = new PostMessageTask(&a, w.winId());
    //w.connect(task, SIGNAL(postDebugMessage(QString)), &w, SLOT(printDebugMessage(QString)));
    //QTimer::singleShot(0, task, SLOT(run()));

    // 절대 경로 확인.
    //qDebug() << QFile::exists("C:\\Users\\kuh\\go\\src\\github.com\\ghts\\ghts\\kiwoom\\c32\\dll\\kiwoom_Go.dll");
    //QDir currentDirectory(QDir::currentPath());
    //qDebug() << currentDirectory.relativeFilePath("C:\\Users\\kuh\\go\\src\\github.com\\ghts\\ghts\\kiwoom\\c32\\dll\\kiwoom_Go.dll");

    QLibrary kiwoom_Go("../dll/kiwoom_Go.dll");

    bool loaded = kiwoom_Go.load();
    qDebug()<<(loaded ? "Loaded Successfully." : "Load Failure.")<<endl;

    if (!loaded) {
        return 0;
    }

    typedef void (*InitDLL)(void*);
    InitDLL initDLL = (InitDLL)kiwoom_Go.resolve("InitDLL");

    if (!initDLL) {
        qDebug("Resolve() Error.");
        return 0;
    } else {
        qDebug("Resolve() OK.");
    }

    initDLL((void *)w.winId());

    return a.exec();
}
