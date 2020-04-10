#include "Func.hpp"

#include <QDebug>

QLibrary *kiwoom_Go = new QLibrary("../dll/kiwoom_Go.dll");

QLibrary *GetKiwoom_Go() {
    static bool loaded = kiwoom_Go->isLoaded();

    if (!loaded) {
        loaded = kiwoom_Go->load();
        qDebug()<<"kiwoom_Go.dll Load"<<OK_ERR(loaded);

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
    qDebug()<<"Resolve('Init')" << OK_ERR(init);

    if (!init) { return false; }

    bool result = init((void *)w->winId());
    qDebug()<<"Init()"<<OK_ERR(result);

    return result;
}

void Confirm(uint serialNo, void *ptrData, int sizeOfData) {
    QLibrary *kiwoom_Go = GetKiwoom_Go();
    if (kiwoom_Go == NULL) {
        qDebug()<<"Abort Confirm() : "<<serialNo;
        return;
    }

    typedef void (*Confirm)(uint, void*, int);
    Confirm confirm = (Confirm)kiwoom_Go->resolve("Confirm");
    qDebug()<<"Resolve('Confirm')" << OK_ERR(confirm);

    if (!confirm) { return; }

    confirm(serialNo, ptrData, sizeOfData);
    qDebug()<<"Confirm() OK.";
}

void ConfirmString(uint serialNo, QString qString) {
    int bufferSize = qString.length()*4;

    qDebug()<<"Buffer size : "<<bufferSize;

    char *buffer = new char[bufferSize]();
    qsnprintf(buffer, sizeof(buffer), "%s", qString.toUtf8().constData());

    Confirm(serialNo, (void*)(&buffer[0]), bufferSize);

    delete[] buffer;
}

