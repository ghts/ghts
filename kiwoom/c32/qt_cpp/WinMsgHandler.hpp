#ifndef EVENTFILTER_HPP
#define EVENTFILTER_HPP

#include <QAbstractNativeEventFilter>
#include "Kiwoom.hpp"

//#include "WinMsg.hpp"
//#include "Func.hpp"
//#include <qt_windows.h>
//#include <QDebug>


class WinMsgHandler : public QAbstractNativeEventFilter {
public:
    virtual bool nativeEventFilter(const QByteArray &eventType, void *message, long *result) Q_DECL_OVERRIDE;
    void setKiwoom(KHOpenAPILib::KHOpenAPI *kiwoom) { this->kiwoom = kiwoom; }

private:
    KHOpenAPILib::KHOpenAPI *kiwoom;

};

#endif // EVENTFILTER_HPP
