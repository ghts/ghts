from eFriendPy.Core import *
import pandas as pd 
from datetime import datetime
import yfinance as yf 
from pandas_datareader import data as pdr 
yf.pdr_override()

class HighLevelAPI():
    """고수준 API 구현 예시"""
    def __init__(self):
        self._core = StandardAPIWrapper()
        self.Password = "0000"                  # 사용자가 직접 입력해야하는 부분
        self.logger = lambda msg : print(msg)   # 사용자가 logger를 외부에서 수정할 수 있도록함(ex> slack 연동 등)

        def SendLogWhenError():
            if self._core.GetRtCode() != "0":
                self.logger("[ERROR] " + self._core.GetReqMessage())

        # 메시지출력 외에 아무일도 안하는 이벤트 핸들러 등록 
        # 이벤트 driven하게 처리하지 않고 그냥 데이터 요청->받기 를 순차적으로 처리하기 위함
        # 아예 핸들러를 등록을 안하면 데이터를 못 받음 (확실하진 않으나 이벤트를 "받았다"는 사실 자체만 중요한 듯함)
        self._core.SetReceiveDataEventHandler(SendLogWhenError) 
        self._core.SetReceiveErrorDataEventHandler(SendLogWhenError)


    def IsConnected(self):
        """eFriend Expert 모듈과 제대로 통신이 되고 있는지 여부를 반환한다
        Account 정보를 가져올 수 있는지 확인하여 통신 여부를 판단.
        """
        accounts = self.GetAllAccounts()
        return len(accounts) > 0


    def GetAllAccounts(self):
        """모든 계좌번호를 list로 반환한다"""
        accountCount = self._core.GetAccountCount()
        return [self._core.GetAccount(i) for i in range(accountCount)]


    def SetAccountInfo(self, account):
        """데이터를 요청하기 위해 계좌 정보를 세팅한다"""
        accountNum, productCode = self.ParseAccountCode(account)

        # 파라미터 set
        self._core.SetSingleData(0, accountNum)
        self._core.SetSingleData(1, productCode)
        self._core.SetSingleData(2, self._core.GetEncryptPassword(self.Password))  


    def ParseAccountCode(self, account):
        """입력받은 계좌번호를 종합계좌번호와 상품코드로 파싱해서 반환한다"""
        accountNum = account[:8]  # 종합계좌번호 (계좌번호 앞 8자리) 
        productCode = account[8:] # 계좌상품코드(종합계좌번호 뒷 부분에 붙는 번호)
        return (accountNum, productCode)


    def BuyKRStock(self, account, productCode, amount, price = 0):
        """국내주식 매수, 설정한 price값이 0 이하이면 시장가로 매수"""
        self.SetAccountInfo(account)
        self._core.SetSingleData(3, productCode)

        orderType = "01" if price <= 0 else "00"    # 주문 구분(메뉴얼 참조), 00: 지정가, 01: 시장가

        self._core.SetSingleData(4, orderType)      
        self._core.SetSingleData(5, str(amount))    # 주문 수량
        self._core.SetSingleData(6, str(price))     # 주문 단가
        self._core.RequestData("SCABO")             # 현금매수 주문

        orderNum = self._core.GetSingleData(1, 0)
        return orderNum


    def SellKRStock(self, account, productCode, amount, price = 0):
        """국내주식 매도, 설정한 price값이 0 이하이면 시장가로 매도"""
        self.SetAccountInfo(account)
        self._core.SetSingleData(3, productCode)

        orderType = "01" if price <= 0 else "00"    # 주문 구분(메뉴얼 참조), 00: 지정가, 01: 시장가

        self._core.SetSingleData(4, "01")           # 매도유형, 01 (고정값으로 추정됨) 
        self._core.SetSingleData(5, orderType)      # 주문 구분(메뉴얼 참조), 00: 지정가, 01: 시장가
        self._core.SetSingleData(6, str(amount))    # 주문 수량
        self._core.SetSingleData(7, str(price))     # 주문 단가
        self._core.RequestData("SCAAO")             # 현금매도 주문

        orderNum = self._core.GetSingleData(1, 0)
        return orderNum


    def BuyUSStock(self, account, marketCode, productCode, amount, price):
        """미국주식 매수, marketCode: 해외거래소코드(NASD / NYSE / AMEX 등 4글자 문자열)"""
        self.SetAccountInfo(account)
        self._core.SetSingleData(3, marketCode)
        self._core.SetSingleData(4, productCode)
        self._core.SetSingleData(5, str(amount))
        self._core.SetSingleData(6, "{0:.2f}".format(price))    # 소숫점 2자리까지로 설정해야 오류가 안남
        self._core.SetSingleData(9, "0")                        # 주문서버구분코드, 0으로 입력
        self._core.SetSingleData(10, "00")                      # 주문구분, 00: 지정가
        self._core.RequestData("OS_US_BUY")                     # 미국매수 주문

        orderNum = self._core.GetSingleData(1, 0)
        return orderNum

    def SellUSStock(self, account, marketCode, productCode, amount, price):
        """미국주식 매도, marketCode: 해외거래소코드(NASD / NYSE / AMEX 등 4글자 문자열)"""
        self.SetAccountInfo(account)
        self._core.SetSingleData(3, marketCode)
        self._core.SetSingleData(4, productCode)
        self._core.SetSingleData(5, str(amount))
        self._core.SetSingleData(6, str(price))
        self._core.SetSingleData(9, "0")           # 주문서버구분코드, 0으로 입력
        self._core.SetSingleData(10, "00")         # 주문구분, 00: 지정가
        self._core.RequestData("OS_US_SEL")        # 미국매도 주문

        orderNum = self._core.GetSingleData(1, 0)
        return orderNum

    def GetProcessedUSOrders(self, account, marketCode, startDate = None):
        """미국주식 체결 내역 조회, marketCode: 해외거래소코드(NASD / NYSE / AMEX 등 4글자 문자열)"""
        today = datetime.today().strftime("%Y%m%d")

        if startDate is None:
            startDate = today   # 오늘
        
        endDate = today         # 오늘
        self.SetAccountInfo(account)

        self._core.SetSingleData(4, startDate)
        self._core.SetSingleData(5, endDate)
        self._core.SetSingleData(6, "00")          # 00: 전체, 01: 매도, 02: 매수
        self._core.SetSingleData(7, "01")          # 00: 전체, 01: 체결, 02: 미체결
        self._core.SetSingleData(8, marketCode) 

        self._core.RequestData("OS_US_CCLD")       # 미국 체결 내역 조회

        # 데이터 받아오기
        recordCount = self._core.GetMultiRecordCount(0)
        res = pd.DataFrame(index=range(recordCount), columns=["주문일자", "주문번호", "원주문번호", "상품번호", "매수매도구분코드명", "주문수량", "체결단가"])
        for recordIdx in range(recordCount):
            date = self._core.GetMultiData(0, recordIdx, 0, 0)
            orderNum = self._core.GetMultiData(0, recordIdx, 2, 0)
            orgOrderNum = self._core.GetMultiData(0, recordIdx, 3, 0)
            productNum = self._core.GetMultiData(0, recordIdx, 12, 0)
            buyOrSell = self._core.GetMultiData(0, recordIdx, 5, 0)
            orderAmount = self._core.GetMultiData(0, recordIdx, 10, 0)
            price =  self._core.GetMultiData(0, recordIdx, 13, 0)

            res.loc[recordIdx] = [date, orderNum, orgOrderNum, productNum, buyOrSell, orderAmount, price]

        return res

    def GetProcessedKROrders(self, account, startDate = None):
        """국내 주식 체결 내역 조회"""
        today = datetime.today().strftime("%Y%m%d")
        if startDate is None:
            startDate = today    # 오늘
        
        endDate = today          # 오늘
        self.SetAccountInfo(account)

        self._core.SetSingleData(3, startDate)
        self._core.SetSingleData(4, endDate)

        self._core.SetSingleData(5, "00")  # 매도매수구분코드. 전체: 00, 매도: 01, 매수: 02
        self._core.SetSingleData(6, "00")  # 조회구분.        역순: 00, 정순: 01
        self._core.SetSingleData(8, "01")  # 체결구분.        전체: 00, 체결: 01, 미체결: 02
        self._core.RequestData("TC8001R")  # 주식 일별 주문 체결 조회 (3개월 이내)

        # 데이터 받아오기
        recordCount = self._core.GetMultiRecordCount(0)
        res = pd.DataFrame(index=range(recordCount), columns=["주문일자", "주문번호", "원주문번호", "상품번호", "매수매도구분코드명", "주문수량"])
        for recordIdx in range(recordCount):
            date = self._core.GetMultiData(0, recordIdx, 0, 0)
            orderNum = self._core.GetMultiData(0, recordIdx, 2, 0)
            orgOrderNum = self._core.GetMultiData(0, recordIdx, 3, 0)
            productNum = self._core.GetMultiData(0, recordIdx, 7, 0)
            buyOrSell = self._core.GetMultiData(0, recordIdx, 6, 0)
            orderAmount = self._core.GetMultiData(0, recordIdx, 9, 0)

            res.loc[recordIdx] = [date, orderNum, orgOrderNum, productNum, buyOrSell, orderAmount]

        return res


    def GetUnprocessedUSOrders(self, account, marketCode):
        """미국 주식 미체결 내역 조회, marketCode: 해외거래소코드(NASD / NYSE / AMEX 등 4글자 문자열)"""
        self.SetAccountInfo(account)
        self._core.RequestData("OS_US_NCCS") # 미국 미체결 조회

        # 데이터 받아오기
        recordCount = self._core.GetMultiRecordCount(0)
        res = pd.DataFrame(index=range(recordCount), columns=["주문일자", "주문번호", "원주문번호", "상품번호", "매수매도구분코드명", "주문수량"])
        for recordIdx in range(recordCount):
            date = self._core.GetMultiData(0, recordIdx, 0, 0)
            orderNum = self._core.GetMultiData(0, recordIdx, 2, 0)
            orgOrderNum = self._core.GetMultiData(0, recordIdx, 3, 0)
            productNum = self._core.GetMultiData(0, recordIdx, 5, 0)
            buyOrSell = self._core.GetMultiData(0, recordIdx, 7, 0)
            orderAmount = self._core.GetMultiData(0, recordIdx, 17, 0)

            res.loc[recordIdx] = [date, orderNum, orgOrderNum, productNum, buyOrSell, orderAmount]

        return res

    def GetUnprocessedKROrders(self, account):
        """국내 주식 미체결 내역 조회
        (정정 취소 가능 주문 조회)
        """
        self.SetAccountInfo(account)
        self._core.SetSingleData(5, "0")  # 조회구분. 주문순: 0, 종목순: 1
        self._core.RequestData("SMCP") # 국내주식 정정 취소 가능 주문 조회

        # 데이터 받아오기
        recordCount = self._core.GetMultiRecordCount(0)
        res = pd.DataFrame(index=range(recordCount), columns=["주문번호", "원주문번호", "상품번호", "매수매도구분코드명", "주문수량"])
        for recordIdx in range(recordCount):
            orderNum = self._core.GetMultiData(0, recordIdx, 1, 0)
            orgOrderNum = self._core.GetMultiData(0, recordIdx, 2, 0)
            productNum = self._core.GetMultiData(0, recordIdx, 4, 0)
            buyOrSell = self._core.GetMultiData(0, recordIdx, 13, 0)
            orderAmount = self._core.GetMultiData(0, recordIdx, 7, 0)

            if orderNum == "":
                res = res.iloc[:recordIdx]
                break

            res.loc[recordIdx] = [orderNum, orgOrderNum, productNum, buyOrSell, orderAmount]

        return res



    def CancelUSOrder(self, account, marketCode, productCode, orgOrderNum, amount):
        """미국 주식 주문을 취소한다, marketCode: 해외거래소코드(NASD / NYSE / AMEX 등 4글자 문자열)"""
        self.SetAccountInfo(account)
        self._core.SetSingleData(3, marketCode)
        self._core.SetSingleData(4, productCode)
        self._core.SetSingleData(5, orgOrderNum)
        self._core.SetSingleData(6, "02")       # 02 : 취소, 01 : 정정 
        self._core.SetSingleData(7, str(amount))

        self._core.RequestData("OS_US_CNC")
        # orderNum = self._core.GetSingleData(1, 0)
        # return orderNum


    def CancelKROrder(self, account, orgOrderNum, amount):
        """국내 주식 주문을 취소한다."""
        self.SetAccountInfo(account)
        self._core.SetSingleData(4, orgOrderNum)
        self._core.SetSingleData(5, "00")       # 주문 구분, 취소인 경우는 00
        self._core.SetSingleData(5, "02")       # 정정취소구분코드. 02: 취소, 01: 정정
        self._core.SetSingleData(7, amount)     # 주문수량

        self._core.RequestData("SMCO")          # 국내 주식 주문 정정 취소
        

    def CancelAllUnprocessedUSOrders(self, account, marketCode):
        """미체결 미국 주식 주문을 모두 취소한다, marketCode: 해외거래소코드(NASD / NYSE / AMEX 등 4글자 문자열)"""
        unprocessed = self.GetUSUnprocessedOrders(account, marketCode)

        for i in unprocessed.index:
            orgOrderNum = unprocessed.loc[i, "원주문번호"]
            if orgOrderNum == "":
                orgOrderNum = unprocessed.loc[i, "주문번호"]

            productCode = unprocessed.loc[i, "상품번호"]
            amount = unprocessed.loc[i, "주문수량"]

            self.CancelKROrder(account, marketCode, productCode, orgOrderNum, amount)

    def CancelAllUnprocessedKROrders(self, account):
        """미체결 국내 주식 주문을 모두 취소한다."""
        unprocessed = self.GetUnprocessedKROrders(account)

        for i in unprocessed.index:
            orgOrderNum = unprocessed.loc[i, "원주문번호"]
            if orgOrderNum == "":
                orgOrderNum = unprocessed.loc[i, "주문번호"]

            amount = unprocessed.loc[i, "주문수량"]

            self.CancelKROrder(account, orgOrderNum, amount)
        

    def GetKRStockPrice(self, stockCode):
        """종목코드(문자열)에 해당하는 국내주식 현재가 시세 반환"""
        self._core.SetSingleData(0, "J")               # 시장분류코드, J: 주식, ETF, ETN
        self._core.SetSingleData(1, stockCode)
        self._core.RequestData("SCP")                  # 현재가 시세 요청

        return int(self._core.GetSingleData(11, 0))    # 현재가 데이터 반환


    def GetUSStockPrice(self, stockCode):
        """종목코드(문자열)에 해당하는 미국주식 현재가 반환 (yahoo finance 사용)"""
        today = datetime.today().strftime("%Y-%m-%d")
        df = pdr.get_data_yahoo(stockCode, start=today)
        if len(df) == 0:
            return Exception("미국 주식 {0} 의 현재가를 가져오는데 실패했습니다.".format(stockCode))

        return df["Close"].iloc[0]


    def GetUSDtoKRWRate(self, account):
        """1 달러 -> 원으로 환전할때의 현재 기준 예상환율을 반환
        예상환율은 최초고시 환율로 매일 08:15시경에 당일 환율이 제공됨
        """
        self.SetAccountInfo(account)
        self._core.SetSingleData(3, "512")          # 미국: 512, 홍콩: 501, 중국: 551, 일본: 515
        self._core.RequestData("OS_OS3004R")        # 해외증거금조회

        rate = self._core.GetMultiData(3, 0, 4, 0)  # 예상환율
        # CurrencyCode = self._core.GetMultiData(3, 0, 1, 0)
        # print("[DEBUG] CurrencyCode: {0}, rate: {1}".format(CurrencyCode, rate))

        return float(rate)

    def GetKRTotalEvaluatedPrice(self, account):
        """해당 계좌의 국내 주식(+원화예수금) 총평가금액을 반환"""
        deposit = self.GetCashKR(account)
        stocks = self.GetKRStocks(account)

        return deposit + stocks["평가금액"].sum()


    def GetUSTotalEvaluatedPrice(self, account):
        """해당 계좌의 미국 주식(+주문가능 달러예수금) 총평가금액을 반환"""
        deposit = self.GetCashUS(account)
        stocks = self.GetUSStocks(account)

        return deposit + stocks["평가금액"].sum()

    def GetCashKR(self, account):
        """주문가능 현금 반환(원화)"""
        self.SetAccountInfo(account)
        self._core.SetSingleData(5, "01")
        self._core.RequestData("SCAP")

        return int(self._core.GetSingleData(0, 0))  # 주문가능현금

    def GetCashUS(self, account):
        """주문 가능한 현금 반환(USD)"""  
        # 데이터 요청
        self.SetAccountInfo(account)   
        self._core.RequestData("OS_US_DNCL")

        # 데이터 받아오기
        recordCount = self._core.GetMultiRecordCount(0)

        for recordIdx in range(recordCount):
            CurrencyCode = self._core.GetMultiData(0, recordIdx, 0, 0)
            if CurrencyCode != "USD":
                continue
            
            cash = self._core.GetMultiData(0, recordIdx, 4, 0)  # 외화주문가능금액
            return float(cash)

        return 0.0


    def GetKRStocks(self, account):
        """국내주식 잔고 정보를 반환하는 함수
        각 보유 종목들에 대한 내용을 DataFrame 형태로 반환한다
        """
        # 데이터 요청
        self.SetAccountInfo(account)   
        self._core.RequestData("SATPS")           # 국내주식 잔고 현황

        # 데이터 받아오기
        # 보유 주식 정보
        recordCount = self._core.GetMultiRecordCount(0)
        stocks = pd.DataFrame(index=range(recordCount), columns=["종목코드", "종목명", "현재가", "보유수량", "평가금액"])

        for recordIdx in range(recordCount):
            ProductNum = self._core.GetMultiData(0, recordIdx, 0, 0)
            ProductName = self._core.GetMultiData(0, recordIdx, 1, 0)
            CurrentPrice = self._core.GetMultiData(0, recordIdx, 11, 0)
            Hold = self._core.GetMultiData(0, recordIdx, 7, 0)
            EvalPrice = self._core.GetMultiData(0, recordIdx, 12, 0)

            # 실패했을때에 대한 예외처리
            if ProductNum == "":        
                stocks = stocks.iloc[:recordIdx]
                break

            stocks.loc[recordIdx] = [ProductNum, ProductName, int(CurrentPrice), int(Hold), int(EvalPrice)]

        return stocks


    def GetUSStocks(self, account):
        """미국주식 잔고 정보를 반환하는 함수
        각 보유 종목들에 대한 내용을 DataFrame 형태로 반환한다
        """
        # 데이터 요청
        self.SetAccountInfo(account)  
        self._core.RequestData("OS_US_CBLC")           # 미국주식 잔고 현황

        # 데이터 받아오기 
        # 보유 주식 정보
        recordCount = self._core.GetMultiRecordCount(0)
        stocks = pd.DataFrame(index=range(recordCount), columns=["해외거래소코드","종목코드", "종목명", "현재가", "보유수량", "평가금액"])

        for recordIdx in range(recordCount):
            MarketCode = self._core.GetMultiData(0, recordIdx, 14, 0)
            ProductNum = self._core.GetMultiData(0, recordIdx, 3, 0)
            ProductName = self._core.GetMultiData(0, recordIdx, 4, 0)
            CurrentPrice = self._core.GetMultiData(0, recordIdx, 12, 0)
            Hold = self._core.GetMultiData(0, recordIdx, 8, 0)
            EvalPrice = self._core.GetMultiData(0, recordIdx, 11, 0)

            # recordCount가 실제 보유 종목수와 상관없이 100으로 들어가는듯함. 이에 대한 예외처리
            if ProductNum == "":        
                stocks = stocks.iloc[:recordIdx]
                break

            stocks.loc[recordIdx] = [MarketCode, ProductNum, ProductName, float(CurrentPrice), int(Hold), float(EvalPrice)]

        return stocks


   
        