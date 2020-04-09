#include "func.hpp"
#include "var.hpp"
#include <QDebug>

bool LoadDLL() {
    // 경로 확인.
    //qDebug() << QFile::exists("C:\\Users\\kuh\\go\\src\\github.com\\ghts\\ghts\\kiwoom\\c32\\dll\\kiwoom_Go.dll");
    //QDir currentDirectory(QDir::currentPath());
    //qDebug() << currentDirectory.relativeFilePath("C:\\Users\\kuh\\go\\src\\github.com\\ghts\\ghts\\kiwoom\\c32\\dll\\kiwoom_Go.dll");

    if (!kiwoom_Go->isLoaded()) {
        kiwoom_Go->load();
    }

    bool loaded = kiwoom_Go->isLoaded();
    qDebug()<<"kiwoom_Go.dll Load"<<OK_ERR(loaded);

    return loaded;
}

bool InitGoDLL(QWidget *w) {
    if (!LoadDLL()) { return false; }

    typedef bool (*InitGoDLL)(void*);
    InitGoDLL initGoDLL = (InitGoDLL)kiwoom_Go->resolve("InitGoDLL");
    qDebug()<<"Resolve('InitGoDLL')" << OK_ERR(initGoDLL);

    if (!initGoDLL) { return false; }

    bool result = initGoDLL((void *)w->winId());
    qDebug()<<"InitGoDLL()"<<OK_ERR(result);

    return result;
}

void NotifyLoginResult(bool loginResult) {
    typedef void (*Notifiy)(bool);
    Notifiy notifiy = (Notifiy)kiwoom_Go->resolve("NotifyLoginResult");
    qDebug()<<"Resolve('NotifiyLoginResult') " << OK_ERR(notifiy);

    if (!notifiy) { return; }

    notifiy(loginResult);
    qDebug("NotifiyLoginResult() OK.");
}
