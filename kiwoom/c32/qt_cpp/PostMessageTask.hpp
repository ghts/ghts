#ifndef POSTCONNECDEVENT_HPP
#define POSTCONNECDEVENT_HPP

#include <QtCore>
#include <QWidget>

#include <qt_windows.h>
#include <WinMsg.hpp>

class PostMessageTask : public QObject {
    Q_OBJECT
public:
    PostMessageTask(QObject *parent = 0, WId wId = 0) : QObject(parent) {
        this->wId = wId;
    }
    virtual ~PostMessageTask() {};

public slots:
    void run() {
        // Do processing here
        BOOL result = PostMessageA(HWND(wId), KWM_CONNECT,0, 0);

        if (result == FALSE) {
            qDebug("PostMessage() Error.");
            emit postDebugMessage("PostMessage() Error.");
        } else {
            qDebug("PostMessage() OK.");
            emit postDebugMessage("PostMessage() OK.");
        }
    }

signals:
    void postDebugMessage(QString debugMessage);

private:
    WId wId;

};

#endif // POSTCONNECDEVENT_HPP
