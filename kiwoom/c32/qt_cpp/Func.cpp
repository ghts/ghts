#include "Func.hpp"

#include <QDebug>

QLibrary *kiwoom_Go = new QLibrary("../dll/kiwoom_Go.dll");

QLibrary *GetKiwoom_Go() {
    static bool loaded = kiwoom_Go->isLoaded();

    if (!loaded) {
        loaded = kiwoom_Go->load();

        if (!loaded) {
            return NULL;
        }
    }

    return kiwoom_Go;
}

bool Init(QWidget *w) {
    QLibrary *kiwoom_Go = GetKiwoom_Go();
    if (kiwoom_Go == NULL) {
        qDebug()<<"Abort Init().";
        return false;
    }

    typedef bool (*Init)(void*);
    Init init = (Init)kiwoom_Go->resolve("Init");


    if (!init) {
         qDebug()<<"Resolve('Init') Fail.";
         return false;
    }

    bool result = init((void *)w->winId());

    return result;
}

void Confirm(uint serialNo, QString qString) {
    QLibrary *kiwoom_Go = GetKiwoom_Go();
    if (kiwoom_Go == NULL) {
        qDebug()<<"Abort Confirm() : "<<serialNo;
        return;
    }

    typedef void (*Confirm)(uint, char*, int);
    Confirm confirm = (Confirm)kiwoom_Go->resolve("Confirm");

    if (!confirm) {
        qDebug()<<"Abort Confirm() : "<<serialNo;
        return;
    }

    int bufferSize = qString.length()*4;
    char *buffer = new char[bufferSize]();
    qsnprintf(buffer, bufferSize, "%s", qString.toUtf8().constData());

    confirm(serialNo, buffer, bufferSize);

    delete[] buffer;
}

