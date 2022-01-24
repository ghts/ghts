from PyQt5.QAxContainer import QAxWidget
from PyQt5.QtCore import QEventLoop


class StandardAPIWrapper:
    """ 한국투자증권 eFriend Expert 표준 API Reference Guide에 소개된 API 함수들 중 일부를 wrapping 한 객체
    """
    def __init__(self):
        self._instance = QAxWidget("ITGExpertCtl.ITGExpertCtlCtrl.1")
        self.EventLoop = None

    def _setEventHandler(self, event, handler):
        """이벤트 핸들러 등록 관련함수의 코드 중복을 제거하기 위한 함수"""
        # 데코레이터
        def decoratedHandler():
            handler()
            if self.EventLoop is not None:
                self.EventLoop.exit()     # handler 를 실행 후 eventLoop를 종료 하고 리셋
                self.EventLoop = None
        
        event.connect(decoratedHandler) # handler 연결


    def SetReceiveErrorDataEventHandler(self, handler):
        """ReceiveErrorData 이벤트에 대한 핸들러를 등록한다
        ReceiveErrorData: 사용자가 요청한 조회 데이터 오류시 발생되는 이벤트
        """
        self._setEventHandler(self._instance.ReceiveErrorData, handler)

    def SetReceiveDataEventHandler(self, handler):
        """ReceiveData 이벤트에 대한 핸들러를 등록한다
        ReceiveData 이벤트: 사용자가 요청한 조회 데이터를 받았을 때 발생되는 이벤트
        """
        self._setEventHandler(self._instance.ReceiveData, handler)

    # ========= 공식 API에 대한 wrapper 함수들 (* 모든 API를 wrapping한 것은 아님) =========

    def SetSingleData(self, fieldIndex, value):
        """사용자가 요청할 서비스의 Input이 단건(Single형) 데이터 값일 때 사용하는 공통함수
        
        [설명] 
        데이터를 요청하는 RequestData함수를 사용하기전에 단건 Input 데이터 값을 넣을 때 사용합니다. 
        전송할 서비스에 인덱스별 데이터 값을 설정하기 위해 사용합니다.
        """
        return self._instance.dynamicCall("SetSingleData(int,QString)", fieldIndex, value)

    def SetSingleDataEx(self, blockIndex, fieldIndex, value):
        """사용자가 요청할 서비스의 Input이 다건 데이터 값일 때 사용하는 공통함수
        
        [설명]
        데이터를 요청하는 RequestData함수를 사용하기전에 다건 Input 데이터 값을 넣을 때 사용합니다. 
        전송할 서비스에 인덱스별 데이터 값을 설정하기 위해 사용합니다.
        """
        return self._instance.dynamicCall("SetSingleDataEx(int, int, QString)", blockIndex, fieldIndex, value)

    def SetMultiData(self, recordIndex, fieldIndex, value):
        """사용자가 요청할 서비스의 Input이 다건(Multi형) 데이터 값일 때 사용하는 공통함수
        
        [설명] 
        데이터를 요청하는 RequestData함수를 사용하기전에 다건 Input 데이터 값을 넣을 때 사용합니다. 
        전송할 서비스에 인덱스별 데이터 값을 설정하기 위해 사용합니다.
        """
        return self._instance.dynamicCall("SetMultiData(int, int, QString)", recordIndex, fieldIndex, value)

    def GetSingleFieldCount(self):
        """요청한 데이터를 ReceiveData에서 받았을 때Output 블록이 단건 일 경우 필드 개수를 가져오는 공통함수(편의성을 위해 제공함)
        
        [설명] 
        ReceiveData로 수신된 단건 Output 필드 개수를 반환합니다.
        """
        return self._instance.dynamicCall("GetSingleFieldCount()")

    def GetMultiBlockCount(self):
        """요청한 데이터를 ReceiveData에서 받았을 때Output 블록에 Multi 블록이 하나 이상일 경우 블록 개수를 반환하는 공통함수
        
        [설명] 
        서비스의 Multi(다건) 블록이 여러 개 일 경우 사용함(보통은 하나이므로 예외적인 케이스에만 사용함)
        """
        return self._instance.dynamicCall("GetMultiBlockCount()")

    def GetMultiRecordCount(self, blockIndex):
        """요청한 데이터를 ReceiveData에서 받았을 때 멀티 블록 레코드(Row) 개수를 반환하는 공통함수
        
        [설명] 
        해당 서비스의 Output 블록이 Multi일 때 레코드 개수(Occurs Count)를 반환합니다.
        """
        return self._instance.dynamicCall("GetMultiRecordCount(int)", blockIndex)

    def GetMultiFieldCount(self, blockIndex, recordIndex):
        """요청한 데이터를 ReceiveData에서 받았을 때 멀티 블록의 Field 개수를 얻을 때 사용하는 공통함수(편의성을 위해 제공함)
        
        [설명] 
        선택한 멀티 블록 레코드의 필드 개수를 반환합니다.
        """
        return self._instance.dynamicCall("GetMultiFieldCount(int, int)", blockIndex, recordIndex)

    def GetSingleData(self, fieldIndex, attributeType):
        """요청한 데이터를 ReceiveData, ReceiveRealData에서 받았을 때 단건 데이터 필드 값을 얻을 때 사용하는 공통함수(단건 조회 및 실시간 데이터 값을 얻을 때 사용함)
        
        [매개변수]  
        fieldIndex: 수신된 필드의 Index 값.  attributeType: 얻고자 하는 데이터 타입을 입력(0: 데이터값, 1: 속성값)

        [반환값]
        수신된 서비스의 데이터 값과 속성 값을 반환합니다.

        [설명]
        ReceiveData, ReceiveRealData 이벤트 발생시 인덱스 값과 속성 구분을 설정하여 수신된 서비스의 데이터 값과 속성 값을 받습니다. 
        데이터의 형태는 모두 스트링 형태이므로 사용자가 데이터를 가공해서 사용할 때는 원하는 타입에 맞춰서 casting을 해서 사용해야 합니다.
        """
        return self._instance.dynamicCall("GetSingleData(int, int)", fieldIndex, attributeType)

    def GetSingleDataEx(self, blockIndex, fieldIndex, attributeType):
        """요청한 데이터를 ReceiveData, ReceiveRealData에서 받았을 때 다건 데이터 필드 값을 얻을 때 사용하는 공통함수
        
        (* 공식 메뉴얼에 자세한 설명이 없으나, GetSingleData과 blockIndex 이외에는 GetSingleData와 동일한 것으로 보임)
        """
        return self._instance.dynamicCall("GetSingleDataEx(int, int, int)", blockIndex, fieldIndex, attributeType)

    def GetMultiData(self, blockIndex, recordIndex, fieldIndex, attributeType):
        """요청한 데이터를 ReceiveData에서 받았을 때 멀티 데이터 필드 값을 얻는데 사용하는 공통함수
        
        [매개변수] 
        blockIndex: 멀티 블록의 인덱스.(보통 Multi Block이 하나이므로 0으로 셋팅함)
        recordIndex: 해당 멀티 블록의 레코드 인덱스
        fieldIndex: 수신된 필드의 Index 값.  
        attributeType: 해당 인덱스의 속성 값(0: 데이터값, 1: 속성값)

        [반환값]
        선택한 인덱스의 데이터 값을 반환합니다.

        [설명]
        해당 서비스가 멀티 블록일 경우 사용자가 원하는 데이터 필드 값을 설정하여 수신된 데이터 또는 속성값을 받을 때 사용합니다. 
        데이터의 형태는 모두 스트링 형태이므로 사용자가 데이터를 가공해서 사용할 때는 원하는 타입에 맞춰서 casting을 해서 사용해야 합니다. 
        다건 조회 데이터를 얻을 때만 사용되는 함수입니다.(실시간 데이터에서는 사용되지 않음)
        """
        return self._instance.dynamicCall("GetMultiData(int, int, int, int)", blockIndex, recordIndex, fieldIndex, attributeType)

    def GetReqMsgCode(self):
        """요청한 데이터를 ReceiveData, ReceiveErrorData에서 받았을 때 요청한 서비스의 메시지 코드 정보를 받는데 사용하는 공통함수
        
        [반환값]
        선택한 인덱스의 데이터 값을 반환합니다.

        [설명]
        요청한 서비스의 메시지 코드 값을 반환합니다. 요청 서비스가 정상적으로 처리됐는지 기타 에러 처리시에 사용합니다.
        """
        return self._instance.dynamicCall("GetReqMsgCode()")

    def GetRtCode(self):
        """요청한 데이터를 ReceiveErrorData에서 받았을 때 요청한 서비스의 메시지 코드 정보를 받는데 사용하는 공통함수
        
        [반환값]
        서버에서 수신된 서비스의 메시지 코드 값을 반환합니다.

        [설명]
        요청한 서비스의 메시지 코드 값을 반환합니다. 요청 서비스가 정상적으로 처리됐는지 기타 에러 처리시에 사용합니다.
        """
        return self._instance.dynamicCall("GetRtCode()")

    def GetReqMessage(self):
        """요청한 데이터를 ReceiveData, ReceiveErrorData에서 받았을 때 요청한 서비스의 메시지 정보를 받는 공통함수
        
        [반환값]
        서버에서 수신된 서비스의 통신 메시지를 받습니다.

        [설명]
        요청한 서비스의 응답 메시지를 반환합니다 응답 메시지를 통해 잘못된 정보를 확인할 수 있습니다.
        """
        return self._instance.dynamicCall("GetReqMessage()")

    def RequestData(self, quarry):
        """사용자가 SetSingleData , SetMultiData로 넣은 Input 정보를 가지고 해당 서비스의 조회를 요청하는 공통함수
        
        [매개변수]
        quarry: 요청할 서비스명

        [설명]
        사용자가 사용할 서비스 명을 설정하고 데이터를 요청한다.
        """
        self._instance.dynamicCall("RequestData(QString)", quarry)
        if self.EventLoop is not None:
            self.EventLoop.exit()
        self.EventLoop = QEventLoop()   # 이벤트루프 할당
        self.EventLoop.exec_()          # 이벤트루프 실행

    def SetMultiBlockData(self, blockIndex, recordIndex, fieldIndex, value):
        """멀티 블록이 여러 개 일 경우 Input 데이터를 설정하는 공통함수

        [반환값]
        0(FALSE) 이면 실패, 1(TRUE) 이면 성공

        [설명]
        멀티 블록 서비스인 경우 Input 항목에 데이터를 설정할 때 사용한다. 
        SetsingleData와 SetMultiData는 하나의 블록 기준이지만 여러 개의 멀티 블록이 사용되는 서비스가 추후에 만들어질 경우를 대비해서 만들어 놓은 함수이다.
        """
        return self._instance.dynamicCall("SetMultiBlockData(int, int, int, QString)", blockIndex, recordIndex, fieldIndex, value)

    def GetAccountCount(self):
        """로그인한 사용자 계좌 개수를 반환하는 공통함수
        
        [반환값]
        계좌 개수를 반환

        [설명]
        사용자 프로그램 시작시 계좌 리스트 정보를 가져올 때 사용하며 로그인한 사용자의 개설된 계좌 개수를 반환합니다.
        """
        return self._instance.dynamicCall("GetAccountCount()")

    def GetAccount(self, accountIndex):
        """인덱스 별 사용자 계좌를 가져오는 공통함수
        
        [매개변수]
        accountIndex: 계좌의 인덱스 값.

        [반환값]
        계좌 번호 반환 (단, 이때 총 10자리를 반환하며, 앞의 8자리는 종합계좌 번호 뒤의 2자리는 상품계좌 번호이다.)

        [설명]
        사용자 프로그램 시작시 계좌 리스트 정보를 가져올 때 사용하며 인덱스 별 사용자 계좌를 반환합니다.
        """
        return self._instance.dynamicCall("GetAccount(int)", accountIndex)

    def GetEncryptPassword(self, rawPassword):
        """암호화된 계좌 비밀번호를 반환하는 공통함수
        
        [매개변수]
        rawPassword: 암호화 안 된 비밀번호

        [반환값]
        암호화 처리된 비밀번호
        
        [설명]
        암호화 처리된 비밀번호를 반환합니다. 매수/매도/정정/취소 주문 서비스 요청시 필요한 암호화 된 비밀번호를 반환하는데 사용합니다. 
        이 함수를 사용하지 않고 암호화 되지 않은 비밀번호를 서버에 올릴 경우 정상적으로 주문이 처리되지 않습니다.
        """
        return self._instance.dynamicCall("GetEncryptPassword(QString)", rawPassword)

    def GetOverSeasStockSise(self):
        """해외주식 이용시 사용자 권한 정보를 가져오는 함수
        
        [반환값]
        해외주식 이용 사용자 권한 정보

        [설명]
        해외주식 이용 사용자 권한 정보를 반환합니다. 해외주식 이용신청시 신청자에 대한 암호값이 반환되며 미신청이용자의 경우 지연시세를 이용할 수 있는 값이 반환됩니다.
        """
        return self._instance.dynamicCall("GetOverSeasStockSise()")