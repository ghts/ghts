#ifndef EVENTFILTER_HPP
#define EVENTFILTER_HPP

#include <qt_windows.h>
#include <QByteArray>
#include <QAbstractNativeEventFilter>
#include <KiwoomApiWrapper.hpp>
#include <WinMsg.hpp>

class WinMsgHandler : public QAbstractNativeEventFilter {
public:
    virtual bool nativeEventFilter(const QByteArray &, void *message, long *) Q_DECL_OVERRIDE {
        MSG *msg = static_cast<MSG*>(message);
        UINT uMsg = msg->message;

        switch (uMsg) {
        case KWM_CONNECT:
            kiwoom->CommConnect();
            // 호출 스레드가 다를 경우(Go DLL) 결과값 회신할 것.
            return true;
        }

        return false;
    }

    void setKiwoomApiWrapper(KiwoomApiWrapper *kiwoom) { this->kiwoom = kiwoom; }

private:
    KiwoomApiWrapper *kiwoom;

};

#endif // EVENTFILTER_HPP
