#include <QString>
#include <QDebug>

void PrintDebugMsgUtf8(void *utf8Msg) {
    qDebug() << QString::fromUtf8((const char*)utf8Msg);
}
