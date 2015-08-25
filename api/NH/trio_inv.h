/************************************************************************************
	주의

	이 자료는 2013년 10월 15일 기준 자료이며 향후 변경될 가능성이 있습니다.
	자료 구조가 맞지 않을 경우 구조체가 변경되지 않았는지 확인하시기 바랍니다.

	최신 자료는 웹페이지를 통해 안내되며 자동 안내(OpenAPI Login시)를 하고 있으니
	게시를 꼭 확인하시기 바랍니다.

************************************************************************************/

//----------------------------------------------------------------------//
// 주식 현재가 조회 (c1101)
//----------------------------------------------------------------------//

typedef struct tagc1101InBlock							// 기본입력
{
	char Lang[1];	char _Lang;							// 한영구분
	char Code[6];	char _Code;							// 종목코드
} Tc1101InBlock;

typedef struct tagc1101OutBlock	// 종목마스타기본자료
{
	char Code[6];	char _Code;							// 종목코드
	char Title[13];	char _Title;						// 종목명. 첫자리는 kospi200은 ‘*’, 스타지수종목은 ‘#’. 실제 종목명은 12 byte임
	char MarketPrice[7];	char _MarketPrice;			// 현재가
	char DiffSign[1];	char _DiffSign;					// 등락부호. 0x18 :상한, 0x1E :상승, 0x20 :보함, 0x19 :하한, 0x1F :하락. 등락부호는 시장과 관계없이 동일한 코드체계 사용
	char Diff[6];	char _Diff;							// 등락폭
	char DiffRate[5];	char _DiffRate;					// 등락률
	char OfferPrice[7];	char _OfferPrice;				// 매도 호가
	char BidPrice[7];	char _BidPrice;					// 매수 호가
	char Volume[9];	char _Volume;						// 거래량
	char TrVolRate[6];	char _TrVolRate;				// 거래비율
	char FloatRate[5];	char _FloatRate;				// 유동주회전율
	char TrAmount[9];	char _TrAmount;					// 거래대금
	char UpLmtPrice[7];	char _UpLmtPrice;				// 상한가
	char High[7];	char _High;							// 장중고가
	char Open[7];	char _Open;							// 시가
	char VsOpenSign[1];	char _VsOpenSign;				// 시가대비부호
	char VsOpenDiff[6];	char _VsOpenDiff;				// 시가대비등락폭
	char Low[7];	char _Low;							// 장중저가
	char LowLmtPrice[7];	char _LowLmtPrice;			// 하한가
	char QuoteTime[8];	char _QuoteTime;				// 호가시간
	char OfferPrice1[7];	char _OfferPrice1;			// 매도 최우선호가
	char OfferPrice2[7];	char _OfferPrice2;			// 매도 차선 호가
	char OfferPrice3[7];	char _OfferPrice3;			// 매도 차차선 호가
	char OfferPrice4[7];	char _OfferPrice4;			// 매도 4차선 호가
	char OfferPrice5[7];	char _OfferPrice5;			// 매도 5차선 호가
	char OfferPrice6[7];	char _OfferPrice6;			// 매도 6차선 호가
	char OfferPrice7[7];	char _OfferPrice7;			// 매도 7차선 호가
	char OfferPrice8[7];	char _OfferPrice8;			// 매도 8차선 호가
	char OfferPrice9[7];	char _OfferPrice9;			// 매도 9차선 호가
	char OfferPrice10[7];	char _OfferPrice10;			// 매도 10차선 호가
	char BidPrice1[7];	char _BidPrice1;				// 매수 최우선 호가
	char BidPrice2[7];	char _BidPrice2;				// 매수 차선 호가
	char BidPrice3[7];	char _BidPrice3;				// 매수 차차선 호가
	char BidPrice4[7];	char _BidPrice4;				// 매수 4차선 호가
	char BidPrice5[7];	char _BidPrice5;				// 매수 5차선 호가
	char BidPrice6[7];	char _BidPrice6;				// 매수 6차선 호가
	char BidPrice7[7];	char _BidPrice7;				// 매수 7차선 호가
	char BidPrice8[7];	char _BidPrice8;				// 매수 8차선 호가
	char BidPrice9[7];	char _BidPrice9;				// 매수 9차선 호가
	char BidPrice10[7];	char _BidPrice10;				// 매수 10차선 호가
	char OfferVolume1[9];	char _OfferVolume1;			// 매도 최우선 잔량
	char OfferVolume2[9];	char _OfferVolume2;			// 매도 차선 잔량
	char OfferVolume3[9];	char _OfferVolume3;			// 매도 차차선 잔량
	char OfferVolume4[9];	char _OfferVolume4;			// 매도 4차선 잔량
	char OfferVolume5[9];	char _OfferVolume5;			// 매도 5차선 잔량
	char OfferVolume6[9];	char _OfferVolume6;			// 매도 6차선 잔량
	char OfferVolume7[9];	char _OfferVolume7;			// 매도 7차선 잔량
	char OfferVolume8[9];	char _OfferVolume8;			// 매도 8차선 잔량
	char OfferVolume9[9];	char _OfferVolume9;			// 매도 9차선 잔량
	char OfferVolume10[9];	char _OfferVolume10;		// 매도 10차선 잔량
	char BidVolume1[9];	char _BidVolume1;				// 매도 최우선 잔량
	char BidVolume2[9];	char _BidVolume2;				// 매수 차선 잔량
	char BidVolume3[9];	char _BidVolume3;				// 매수 차차선 잔량
	char BidVolume4[9];	char _BidVolume4;				// 매수 4차선 잔량
	char BidVolume5[9];	char _BidVolume5;				// 매수 5차선 잔량
	char BidVolume6[9];	char _BidVolume6;				// 매수 6차선 잔량
	char BidVolume7[9];	char _BidVolume7;				// 매수 7차선 잔량
	char BidVolume8[9];	char _BidVolume8;				// 매수 8차선 잔량
	char BidVolume9[9];	char _BidVolume9;				// 매수 9차선 잔량
	char BidVolume10[9];	char _BidVolume10;			// 매수 10차선 잔량
	char OfferVolTot[9];	char _OfferVolTot;			// 총 매도 잔량
	char BidVolTot[9];	char _BidVolTot;				// 총 매수 잔량
	char OfferVolAfterHour[9];	char _OfferVolAfterHour; // 시간외 매도 잔량
	char BidVolAfterHour[9];	char _BidVolAfterHour;	// 시간외 매수 잔량
	char PivotUp2[7];	char _PivotUp2;					// 피봇 2차 저항 : 피봇가 + 전일 고가 – 전일 저가
	char PivotUp1[7];	char _PivotUp1;					// 피봇 1차 저항 : (피봇가 * 2) – 전일 저가
	char PivotPrice[7];	char _PivotPrice;				// 피봇가 : (전일 고가 + 전일 저가 + 전일 종가) / 3
	char PivotDown1[7];	char _PivotDown1;				// 피봇 1차 지지 : (피봇가 * 2) – 전일 고가
	char PivotDown2[7];	char _PivotDown2;				// 피봇 2차 지지 : 피봇가 – 전일고가 + 전일 저가
	char Market[6];	char _Market;						// 코스피/코스닥 구분 : ‘코스피’ , ‘코스닥’
	char Sector[18];	char _Sector;					// 업종명
	char CapSize[6];	char _CapSize;					// 자본금규모
	char SettleMonth[16];	char _SettleMonth;			// 결산월
	char MarketAction1[16];	char _MarketAction1;		// 시장조치1
	char MarketAction2[16];	char _MarketAction2;		// 시장조치2
	char MarketAction3[16];	char _MarketAction3;		// 시장조치3
	char MarketAction4[16];	char _MarketAction4;		// 시장조치4
	char MarketAction5[16];	char _MarketAction5;		// 시장조치5
	char MarketAction6[16];	char _MarketAction6;		// 시장조치6
	char ConvertBond[6];	char _ConvertBond;			// CB구분
	char NominalPrice[7];	char _NominalPrice;			// 액면가
	char PrevPriceTitle[12];	char _PrevPriceTitle;	// 전일 종가 타이틀 (평가가격, 기준가, 전일종가)
	char PrevPrice[7];	char _PrevPrice;				// 전일종가
	char MortgageValue[7];	char _MortgageValue;		// 대용가
	char PublicOfferPrice[7];	char _PublicOfferPrice;	// 공모가
	char High5Day[7];	char _High5Day;					// 5일고가
	char Low5Day[7];	char _Low5Day;					// 5일저가
	char High20Day[7];	char _High20Day;				// 20일고가
	char Low20Day[7];	char _Low20Day;					// 20일저가
	char High1Year[7];	char _High1Year;				// 52주최고가
	char High1YearDate[4];	char _High1YearDate;		// 52주최고가일
	char Low1Year[7];	char _Low1Year;					// 52주최저가
	char Low1YearDate[4];	char _Low1YearDate;			// 52주최저가일
	char FloatVolume[8];	char _FloatVolume;			// 유동주식수
	char ListVolBy1000[12];	char _ListVolBy1000;		// 상장주식수. 1000주 단위?
	char MarketCapital[9];	char _MarketCapital;		// 시가총액
	char TraderInfoTime[5];	char _TraderInfoTime;		// 거래원 정보 최종 수신 시간
	char Seller1[6];	char _Seller1;					// 매도 거래원1
	char Buyer1[6];	char _Buyer1;						// 매수 거래원1
	char Seller1Volumn[9];	char _Seller1Volume;		// 매도 거래량1
	char Buyer1Volume[9];	char _Buyer1Volume;			// 매수 거래량1
	char Seller2[6];	char _Seller2;					// 매도 거래원2
	char Buyer2[6];	char _Buyer2;						// 매수 거래원2
	char Seller2Volumn[9];	char _Seller2Volume;		// 매도 거래량2
	char Buyer2Volume[9];	char _Buyer2Volume;			// 매수 거래량2
	char Seller3[6];	char _Seller3;					// 매도 거래원3
	char Buyer3[6];	char _Buyer3;						// 매수 거래원3
	char Seller3Volumn[9];	char _Seller3Volume;		// 매도 거래량3
	char Buyer3Volume[9];	char _Buyer3Volume;			// 매수 거래량3
	char Seller4[6];	char _Seller4;					// 매도 거래원4
	char Buyer4[6];	char _Buyer4;						// 매수 거래원4
	char Seller4Volumn[9];	char _Seller4Volume;		// 매도 거래량4
	char Buyer4Volume[9];	char _Buyer4Volume;			// 매수 거래량4
	char Seller5[6];	char _Seller5;					// 매도 거래원5
	char Buyer5[6];	char _Buyer5;						// 매수 거래원5
	char Seller5Volumn[9];	char _Seller5Volume;		// 매도 거래량5
	char Buyer5Volume[9];	char _Buyer5Volume;			// 매수 거래량5
	char ForeignSellVolume[9];	char _ForeignSellVolume; // 외국인 매도 거래량
	char ForeignBuyVolume[9];	char _ForeignBuyVolume;	// 외국인 매수 거래량
	char ForeignTime[6];	char _ForeignTime;			// 외국인 시간 ???
	char ForeignHoldingRate[5];	char _ForeignHoldingRate; // 외국인 지분율
	char SettleDate[4];	char _SettleDate;				// 결제일
	char DebtPercent[5];	char _DebtPercent;			// 잔고 비율(%)
	char RightsIssueDate[4];	char _RightsIssueDate;	// 유상 기준일
	char BonusIssueDate[4];	char _BonusIssueDate;		// 무상 기준일
	char RightsIssueRate[5];	char _RightsIssueRate;	// 유상 배정비율
	char ForeignFloatVol[10];	char _ForeignFloatVol;	// 외국인 변동주 수
	char BonusIssueRate[5];	char _BonusIssueRate;		// 무상 배정비율
	char TreasuryStock[1];	char _TreasuryStock;		// 당일 자사주 신청 여부  1: 자사주 신청
	char IPO_Date[8];	char _IPO_Date;					// 상장일
	char MajorHoldRate[5];	char _MajorHoldRate;		// 대주주지분율
	char MajorHoldInfoDate[6];	char _MajorHoldInfoDate; // 대주주지분일자
	char FourLeafClover[1];	char _FourLeafClover;		// 네잎클로버 종목 여부 1: 네잎클로버 종목
	char MarginRate[1];	char _MarginRate;				// 증거금율
	char Capital[9];	char _Capital;					// 자본금
	char SellTotalSum[9];	char _SellTotalSum;			// 전체 거래원 매도 합계
	char BuyTotalSum[9];	char _BuyTotalSum;			// 전체 거래원 매수 합계
	char Title2[21];	char _Title2;					// 종목명2. 앞에 한자리를 제외하고 18byte가 종목명
	char BackdoorListing[1];	char _BackdoorListing;	// 우회상장여부
	char FloatRate2[6];	char _FloatRate2;				// 유동주회전율2
	char Market2[6];	char _Market2;					// 코스피 구분 ?? 앞에 나왔는 데...
	char DebtTrDate[4];	char _DebtTrDate;				// 공여율기준일
	char DebtTrPercent[5];	char _DebtTrPercent;		// 공여율(%)
	char PER[5];	char _PER;							// PER
	char DebtLimit[1];	char _DebtLimit;				// 종목별신용한도
	char WeightAvgPrice[7];	char _WeightAvgPrice;		// 가중가
	char ListedVolume[12];	char _ListedVolume;			// 상장주식 수  _주
	char AddListing[12];	char _AddListing;			// 추가상장 주식 수
	char Comment[100];	char _Comment;					// 종목 comment
	char PrevVolume[9];	char _PrevVolume;				// 전일 거래량
	char VsPrevSign[1];	char _VsPrevSign;				// 전일대비 등락부호
	char VsPrevDiff[6];	char _VsPrevDiff;				// 전일대비 등락폭
	char High1Year2[7];	char _High1Year2;				// 연중 최고가 (52주 최고가와 중복 아닌가?
	char High1YearDate2[4];	char _High1YearDate2;		// 연중 최고가일
	char Low1Year2[7];	char _Low1Year2;				// 연중 최저가
	char Low1YearDate2[4];	char _Low1YearDate2;		// 연중 최저가일
	char ForeignHoldNo[15];	char _ForeignHoldNo;		// 외국인 보유 주식수
	char ForeignLmtPercent[5];	char _ForeignLmtPercent; // 외국인 한도율(%)
	char TrUnitVolume[5];	char _TrUnitVolume;			// 매매 수량 단위
	char BlackPoolOfferBid[1];	char _BlackPoolOfferBid; // 경쟁대량방향구분. 0: 해당없음, 1: 매도, 2: 매수
	char BlackPoolExist[1];	char _BlackPoolExist;		// 대량매매구분. 1: 대량매매有, 0:대량매매無
} Tc1101OutBlock;

typedef struct tagc1101OutBlock2// 변동거래량자료,[반복]
{
	char Time[8];	char _Time;							// 시간
	char MarketPrice[7];	char _MarketPrice;			// 현재가
	char DiffSign[1];	char _DiffSign;					// 등락부호
	char Diff[6];	char _Diff;							// 등락폭
	char OfferPrice[7];	char _OfferPrice;				// 매도 호가
	char BidPrice[7];	char _BidPrice;					// 매수 호가
	char DiffVolume[8];	char _DiffVolume;				// 변동거래량
	char Volume[9];	char _Volume;						// 거래량
} Tc1101OutBlock2;

typedef struct tagc1101OutBlock3// 종목지표
{
	char SyncOfferBid[1];	char _SyncOfferBid;			// 동시호가 구분.  0:동시호가 아님, 1:동시호가, 2:동시호가연장, 3:시가범위연장, 4:종가범위연장, 5:배분개시, 6:변동성 완화장치 발동
	char EstmPrice[7];	char _EstmPrice;		// 예상체결가
	char EstmSign[1];	char _EstmSign;			// 예상체결 부호
	char EstmDiff[6];	char _EstmDiff;			// 예상체결 등락폭
	char EstmDiffRate[5];	char _EstmDiffRate;	// 예상체결 등락률
	char EstmVol[9];	char _EstmVol;			// 예상체결수량
	char ECN_InfoExist[1];	char _ECN_InfoExist;		// ECN정보 유무 구분 (우리나라에는 ECN이 아직 없을텐데...)
	char ECN_PrevPrice[9];	char _ECN_PrevPrice;		// ECN 전일종가
	char ECN_DiffSign[1];	char _ECN_DiffSign;			// ECN 부호
	char ECN_Diff[9];	char _ECN_Diff;					// ECN 등락폭
	char ECN_DiffRate[5];	char _ECN_DiffRate;			// ECN 등락률
	char ECN_Volume[10];	char _ECN_Volume;			// ECN 체결수량
	char VsECN_EstmSign[1];	char _VsECN_EstmSign; 		// ECN대비 예상 체결 부호
	char VsECN_EstmDiff[6];	char _VsECN_EstmDiff;		// ECN대비 예상 체결 등락폭
	char VsECN_EstmDiffRate[5];	char _ECN_EstmDiffRate;	// ECN대비 예상 체결 등락률
} Tc1101OutBlock3;

typedef struct tagc1101
{
	Tc1101InBlock c1101inblock;							// 기본입력
	Tc1101OutBlock c1101outblock;						// 종목마스타기본자료
	Tc1101OutBlock2 c1101outblock2[20];					// 변동거래량자료 ,[반복]
	Tc1101OutBlock3 c1101outblock3;						// 종목지표
} Tc1101;

typedef struct tags4101InBlock// 기본입력
{
	char Lang[1];	char _Lang;		// 한영구분
	char fuitemz9[9];	char _fuitemz9;		// 종목코드
} Ts4101InBlock;

typedef struct tags4101OutBlock// 종목마스타기본자료
{
	char fuitem[4];	char _fuitem;		// 종목코드
	char fuhname[12];	char _fuhname;		// 종목명
	char fucurr[5];	char _fucurr;		// 현재가
	char fusign[1];	char _fusign;		// 등락부호
	char fuchange[5];	char _fuchange;		// 등락폭
	char fuchrate[5];	char _fuchrate;		// 등락률
	char fubasis[5];	char _fubasis;		// 베이시스
	char futheoryprice[5];	char _futheoryprice;		// 이론가
	char fugrate[5];	char _fugrate;		// 괴리도
	char fugratio[5];	char _fugratio;		// 괴리율
	char fuvolall[7];	char _fuvolall;		// 거래량
	char fuvalall[12];	char _fuvalall;		// 누적거래대금(백만)
	char fuopenyak[7];	char _fuopenyak;		// 미결제약정수량
	char fupreopenyak[7];	char _fupreopenyak;		// 미결제약정전일
	char fuhprice[5];	char _fuhprice;		// 상한가
	char fuhigh[5];	char _fuhigh;		// 고가
	char fuopen[5];	char _fuopen;		// 시가
	char fuopensign[1];	char _fuopensign;		// 시가대비부호
	char fuopenchange[5];	char _fuopenchange;		// 시가대비등락
	char fulow[5];	char _fulow;		// 저가
	char fulprice[5];	char _fulprice;		// 하한가
	char fucbhprice[5];	char _fucbhprice;		// CB발동상한
	char fucblprice[5];	char _fucblprice;		// CB발동하한
	char fudehprice[5];	char _fudehprice;		// DEMARK저항
	char fudelprice[5];	char _fudelprice;		// DEMARK지지
	char fulisthprice[5];	char _fulisthprice;		// 상장후최고가
	char fulisthdate[8];	char _fulisthdate;		// 상장후최고일
	char fulistlprice[5];	char _fulistlprice;		// 상장후최저가
	char fulistldate[8];	char _fulistldate;		// 상장후최저일
	char fulastdate[8];	char _fulastdate;		// 최종거래일
	char fujandatecnt[3];	char _fujandatecnt;		// 잔존일
	char fucdratio[6];	char _fucdratio;		// 무위험이자율
	char fuchetime[8];	char _fuchetime;		// 호가시간
	char fuoffer[5];	char _fuoffer;		// 매도 최우선호가
	char fujoffer[5];	char _fujoffer;		// 매도 차선 호가
	char fujjoffer[5];	char _fujjoffer;		// 매도 차차선 호가
	char fuj4offer[5];	char _fuj4offer;		// 매도 4차선 호가
	char fuj5offer[5];	char _fuj5offer;		// 매도 5차선 호가
	char fubid[5];	char _fubid;		// 매수 최우선호가
	char fujbid[5];	char _fujbid;		// 매수 차선 호가
	char fujjbid[5];	char _fujjbid;		// 매수 차차선 호가
	char fuj4bid[5];	char _fuj4bid;		// 매수 4차선 호가
	char fuj5bid[5];	char _fuj5bid;		// 매수 5차선 호가
	char fuofferjan[6];	char _fuofferjan;		// 매도 최우선 잔량
	char fujofferjan[6];	char _fujofferjan;		// 매도 차선 잔량
	char fujjofferjan[6];	char _fujjofferjan;		// 매도 차차선 잔량
	char fuj4offerjan[6];	char _fuj4offerjan;		// 매도 4차선 잔량
	char fuj5offerjan[6];	char _fuj5offerjan;		// 매도 5차선 잔량
	char fubidjan[6];	char _fubidjan;		// 매수 최우선 잔량
	char fujbidjan[6];	char _fujbidjan;		// 매수 차선 잔량
	char fujjbidjan[6];	char _fujjbidjan;		// 매수 차차선 잔량
	char fuj4bidjan[6];	char _fuj4bidjan;		// 매수 4차선 잔량
	char fuj5bidjan[6];	char _fuj5bidjan;		// 매수 5차선 잔량
	char futofferjan[6];	char _futofferjan;		// 총 매도 잔량
	char futbidjan[6];	char _futbidjan;		// 총 매수 잔량
	char fuoffersu[4];	char _fuoffersu;		// 매도 최우선건수
	char fujoffersu[4];	char _fujoffersu;		// 매도 차선건수
	char fujjoffersu[4];	char _fujjoffersu;		// 매도 차차선건수
	char fuj4offersu[4];	char _fuj4offersu;		// 매도 4차선건수
	char fuj5offersu[4];	char _fuj5offersu;		// 매도 5차선건수
	char fubidsu[4];	char _fubidsu;		// 매수 최우선건수
	char fujbidsu[4];	char _fujbidsu;		// 매수 차선건수
	char fujjbidsu[4];	char _fujjbidsu;		// 매수 차차선건수
	char fuj4bidsu[4];	char _fuj4bidsu;		// 매수 4차선건수
	char fuj5bidsu[4];	char _fuj5bidsu;		// 매수 5차선건수
	char futoffersu[5];	char _futoffersu;		// 총 매도 건수
	char futbidsu[5];	char _futbidsu;		// 총 매수 건수
	char fupivot2upz5[5];	char _fupivot2upz5;		// 피봇2차저항
	char fupivot1upz5[5];	char _fupivot1upz5;		// 피봇1차저항
	char fupivotz5[5];	char _fupivotz5;		// 피봇가
	char fupivot1dnz5[5];	char _fupivot1dnz5;		// 피봇1차지지
	char fupivot2dnz5[5];	char _fupivot2dnz5;		// 피봇2차지지
	char fujgubun[8];	char _fujgubun;		// CB발동여부
	char fuspvolall[7];	char _fuspvolall;		// 스프레드거래량
	char fudivideratio[9];	char _fudivideratio;		// 배당액지수
	char preclose[5];	char _preclose;		// 전일종가
	char fudynhprice[5];	char _fudynhprice;		// 실시간상한가
	char fudynlprice[5];	char _fudynlprice;		// 실시간하한가
	char fudynpriceflag[1];	char _fudynpriceflag;		// 동적가격제한여부
	char fubulkvol[7];	char _fubulkvol;		// 협의거래량
	char exlmtgb[1];	char _exlmtgb;		// 가격확대예정구분
	char uplmtgb[1];	char _uplmtgb;		// 가격제한확대상한단계
	char dnlmtgb[1];	char _dnlmtgb;		// 가격제한확대하한단계
} Ts4101OutBlock;

typedef struct tags4101OutBlock1// 코스피200지수
{
	char fuitem[4];	char _fuitem;		// 코스피200코드
	char fucurr[5];	char _fucurr;		// 코스피200지수
	char fusign[1];	char _fusign;		// 코스피200등락부호
	char fuchange[5];	char _fuchange;		// 코스피200등락폭
	char fuchrate[5];	char _fuchrate;		// 코스피200등락률
} Ts4101OutBlock1;

typedef struct tags4101OutBlock2// 변동거래량자료,[반복]
{
	char fuchetime[8];	char _fuchetime;		// 시간
	char fucurr[5];	char _fucurr;		// 현재가
	char fusign[1];	char _fusign;		// 등락부호
	char fuchange[5];	char _fuchange;		// 등락폭
	char fuoffer[5];	char _fuoffer;		// 매도 호가
	char fubid[5];	char _fubid;		// 매수 호가
	char fuvol[6];	char _fuvol;		// 거래량
	char fuvolall[7];	char _fuvolall;		// 누적거래량
	char fuopenyak[7];	char _fuopenyak;		// 미결제약정
} Ts4101OutBlock2;

typedef struct tags4101OutBlock3// 시간대별투자자현황최근메모리,[반복]
{
	char titlez6[6];	char _titlez6;		// TITLE
	char amesuvalpure[9];	char _amesuvalpure;		// 순 매수
	char cmesuvalpure[9];	char _cmesuvalpure;		// 매도
	char imesuvalpure[9];	char _imesuvalpure;		// 매수
} Ts4101OutBlock3;

typedef struct tags4101OutBlock4// 시간대별투자자현황시간별,[반복]
{
	char timez8[8];	char _timez8;		// 시간별
	char amesuvalpure[9];	char _amesuvalpure;		// 외국인순 매수
	char cmesuvalpure[9];	char _cmesuvalpure;		// 증권순 매수
	char imesuvalpure[9];	char _imesuvalpure;		// 개인순 매수
} Ts4101OutBlock4;

typedef struct tags4101OutBlock5// KOSPI200시가총액상위10종목,[반복]
{
	char Code[6];	char _Code;		// 종목코드
	char Title[13];	char _Title;		// 종목명
	char parvalue[7];	char _parvalue;		// 액면가
	char price[7];	char _price;		// 현재가
	char sign[1];	char _sign;		// 등락부호
	char change[6];	char _change;		// 등락폭
	char chrate[5];	char _chrate;		// 등락률
} Ts4101OutBlock5;

typedef struct tags4101OutBlock6// 예상체결
{
	char dongsi[1];	char _dongsi;		// 동시호가구분
	char jeqprice[5];	char _jeqprice;		// 예상체결가
	char jeqsign[1];	char _jeqsign;		// 예상체결부호
	char jeqchange[5];	char _jeqchange;		// 예상체결등락폭
	char jeqchrate[5];	char _jeqchrate;		// 예상체결등락률
} Ts4101OutBlock6;

typedef struct tags4101
{
	Ts4101InBlock                     s4101inblock                          ;		// 기본입력
	Ts4101OutBlock                    s4101outblock                         ;		// 종목마스타기본자료
	Ts4101OutBlock1                   s4101outblock1                        ;		// 코스피200지수
	Ts4101OutBlock2                   s4101outblock2[30];		// 변동거래량자료 ,[반복]
	Ts4101OutBlock3                   s4101outblock3[3];		// 시간대별투자자현황최근메모리 ,[반복]
	Ts4101OutBlock4                   s4101outblock4[20];		// 시간대별투자자현황시간별 ,[반복]
	Ts4101OutBlock5                   s4101outblock5[10];		// KOSPI200시가총액상위10종목 ,[반복]
	Ts4101OutBlock6                   s4101outblock6                        ;		// 예상체결
} Ts4101;


typedef struct tagc4113InBlock// 입력데이타
{
	char fuitemz9[9];	char _fuitemz9;		// 입력코드
} Tc4113InBlock;

typedef struct tagc4113OutKospi200// c4113OutKospi200
{
	char fuitem[4];	char _fuitem;		// 종목코드
	char fucurr[5];	char _fucurr;		// 현물지수
	char fusign[1];	char _fusign;		// 전일비부호
	char fuchange[5];	char _fuchange;		// 전일비
	char fuopen[5];	char _fuopen;		// 시가
	char fuhigh[5];	char _fuhigh;		// 고가
	char fulow[5];	char _fulow;		// 저가
	char fuvolall[7];	char _fuvolall;		// 거래량
} Tc4113OutKospi200;

typedef struct tagc4113OutSMaster// c4113OutSMaster
{
	char fuitemz8[8];	char _fuitemz8;		// 종목코드
	char fuspcurr[6];	char _fuspcurr;		// 지수
	char fuspsign[1];	char _fuspsign;		// 전일비부호
	char fuspchange[5];	char _fuspchange;		// 전일비
	char fuspchrate[5];	char _fuspchrate;		// 등락률
	char fuspopen[6];	char _fuspopen;		// 시가
	char fusphigh[6];	char _fusphigh;		// 고가
	char fusplow[6];	char _fusplow;		// 저가
	char fuspvolall[7];	char _fuspvolall;		// 거래량
	char fuspvalall[12];	char _fuspvalall;		// 누적거래대금(백만원)
	char fuspcurr1[5];	char _fuspcurr1;		// 의제약정가(근월물)
	char fuspcurr2[5];	char _fuspcurr2;		// 의제약정가(원월물)
	char fudynhprice[6];	char _fudynhprice;		// 실시간상한가
	char fudynlprice[6];	char _fudynlprice;		// 실시간하한가
	char fudynpriceflag[1];	char _fudynpriceflag;		// 동적가격제한여부
} Tc4113OutSMaster;

typedef struct tagc4113OutBlock1// 코스피선물Master1
{
	char fuitem[4];	char _fuitem;		// 종목코드
	char fuchetime[8];	char _fuchetime;		// 체결시간
	char fuhname[12];	char _fuhname;		// 한글명
	char fucurr[5];	char _fucurr;		// 현재가
	char fusign[1];	char _fusign;		// 전일대비부호
	char fuchange[5];	char _fuchange;		// 전일대비
	char fuchrate[5];	char _fuchrate;		// 등락률
	char fubasis[5];	char _fubasis;		// 베이시스
	char futheoryprice[5];	char _futheoryprice;		// 이론가
	char fugrate[5];	char _fugrate;		// 괴리도
	char fugratio[5];	char _fugratio;		// 괴리율
	char fuvolall[7];	char _fuvolall;		// 누적체결수량
	char fuvalall[12];	char _fuvalall;		// 누적거래대금(백만원)
	char fuopenyak[7];	char _fuopenyak;		// 미결제약정수량
	char fupreopenyak[7];	char _fupreopenyak;		// 미결제약정전일
	char fujgubun[8];	char _fujgubun;		// 장운용
	char fuopen[5];	char _fuopen;		// 시가
	char fuhigh[5];	char _fuhigh;		// 고가
	char fulow[5];	char _fulow;		// 저가
	char fudynhprice[5];	char _fudynhprice;		// 실시간상한가
	char fudynlprice[5];	char _fudynlprice;		// 실시간하한가
	char fudynpriceflag[1];	char _fudynpriceflag;		// 동적가격제한여부
} Tc4113OutBlock1;

typedef struct tagc4113OutBlock2// 코스피선물Master2
{
	char fuitem[4];	char _fuitem;		// 종목코드
	char fuchetime[8];	char _fuchetime;		// 체결시간
	char fuhname[12];	char _fuhname;		// 한글명
	char fucurr[5];	char _fucurr;		// 현재가
	char fusign[1];	char _fusign;		// 전일대비부호
	char fuchange[5];	char _fuchange;		// 전일대비
	char fuchrate[5];	char _fuchrate;		// 등락률
	char fubasis[5];	char _fubasis;		// 베이시스
	char futheoryprice[5];	char _futheoryprice;		// 이론가
	char fugrate[5];	char _fugrate;		// 괴리도
	char fugratio[5];	char _fugratio;		// 괴리율
	char fuvolall[7];	char _fuvolall;		// 누적체결수량
	char fuvalall[12];	char _fuvalall;		// 누적거래대금(백만원)
	char fuopenyak[7];	char _fuopenyak;		// 미결제약정수량
	char fupreopenyak[7];	char _fupreopenyak;		// 미결제약정전일
	char fujgubun[8];	char _fujgubun;		// 장운용
	char fuopen[5];	char _fuopen;		// 시가
	char fuhigh[5];	char _fuhigh;		// 고가
	char fulow[5];	char _fulow;		// 저가
	char fudynhprice[5];	char _fudynhprice;		// 실시간상한가
	char fudynlprice[5];	char _fudynlprice;		// 실시간하한가
	char fudynpriceflag[1];	char _fudynpriceflag;		// 동적가격제한여부
} Tc4113OutBlock2;

typedef struct tagc4113OutHoga1// 코스피선물호가1
{
	char fuitem[4];	char _fuitem;		// 종목코드
	char fuhotime[8];	char _fuhotime;		// 호가시간
	char fuoffer[5];	char _fuoffer;		// 매도 우선호가
	char fujoffer[5];	char _fujoffer;		// 차선 매도 호가
	char fujjoffer[5];	char _fujjoffer;		// 차차선 매도 호가
	char fuj4offer[5];	char _fuj4offer;		// 4차선 매도 호가
	char fuj5offer[5];	char _fuj5offer;		// 5차선 매도 호가
	char fuofferjan[6];	char _fuofferjan;		// 매도 잔량
	char fujofferjan[6];	char _fujofferjan;		// 차선 매도 호가잔량
	char fujjofferjan[6];	char _fujjofferjan;		// 차차선 매도 호가잔량
	char fuj4offerjan[6];	char _fuj4offerjan;		// 4차선 매도 호가잔량
	char fuj5offerjan[6];	char _fuj5offerjan;		// 5차선 매도 호가잔량
	char fubid[5];	char _fubid;		// 매수 우선호가
	char fujbid[5];	char _fujbid;		// 차선 매수 호가
	char fujjbid[5];	char _fujjbid;		// 차차선 매수 호가
	char fuj4bid[5];	char _fuj4bid;		// 4차선 매수 호가
	char fuj5bid[5];	char _fuj5bid;		// 5차선 매수 호가
	char fubidjan[6];	char _fubidjan;		// 매수 잔량
	char fujbidjan[6];	char _fujbidjan;		// 차선 매수 호가잔량
	char fujjbidjan[6];	char _fujjbidjan;		// 차차선 매수 호가잔량
	char fuj4bidjan[6];	char _fuj4bidjan;		// 4차선 매수 호가잔량
	char fuj5bidjan[6];	char _fuj5bidjan;		// 5차선 매수 호가잔량
	char futofferjan[6];	char _futofferjan;		// 총 매도 잔량
	char futbidjan[6];	char _futbidjan;		// 총 매수 잔량
	char fuoffersu[4];	char _fuoffersu;		// 매도 최우선건수
	char fujoffersu[4];	char _fujoffersu;		// 매도 차선건수
	char fujjoffersu[4];	char _fujjoffersu;		// 매도 차차선건수
	char fuj4offersu[4];	char _fuj4offersu;		// 매도 4차선건수
	char fuj5offersu[4];	char _fuj5offersu;		// 매도 5차선건수
	char fubidsu[4];	char _fubidsu;		// 매수 최우선건수
	char fujbidsu[4];	char _fujbidsu;		// 매수 차선건수
	char fujjbidsu[4];	char _fujjbidsu;		// 매수 차차선건수
	char fuj4bidsu[4];	char _fuj4bidsu;		// 매수 4차선건수
	char fuj5bidsu[4];	char _fuj5bidsu;		// 매수 5차선건수
	char futoffersu[5];	char _futoffersu;		// 총 매도 건수
	char futbidsu[5];	char _futbidsu;		// 총 매수 건수
	char fuhname[12];	char _fuhname;		// 한글명
} Tc4113OutHoga1;

typedef struct tagc4113OutHoga2// 코스피선물호가2
{
	char fuitem[4];	char _fuitem;		// 종목코드
	char fuhotime[8];	char _fuhotime;		// 호가시간
	char fuoffer[5];	char _fuoffer;		// 매도 우선호가
	char fujoffer[5];	char _fujoffer;		// 차선 매도 호가
	char fujjoffer[5];	char _fujjoffer;		// 차차선 매도 호가
	char fuj4offer[5];	char _fuj4offer;		// 4차선 매도 호가
	char fuj5offer[5];	char _fuj5offer;		// 5차선 매도 호가
	char fuofferjan[6];	char _fuofferjan;		// 매도 잔량
	char fujofferjan[6];	char _fujofferjan;		// 차선 매도 호가잔량
	char fujjofferjan[6];	char _fujjofferjan;		// 차차선 매도 호가잔량
	char fuj4offerjan[6];	char _fuj4offerjan;		// 4차선 매도 호가잔량
	char fuj5offerjan[6];	char _fuj5offerjan;		// 5차선 매도 호가잔량
	char fubid[5];	char _fubid;		// 매수 우선호가
	char fujbid[5];	char _fujbid;		// 차선 매수 호가
	char fujjbid[5];	char _fujjbid;		// 차차선 매수 호가
	char fuj4bid[5];	char _fuj4bid;		// 4차선 매수 호가
	char fuj5bid[5];	char _fuj5bid;		// 5차선 매수 호가
	char fubidjan[6];	char _fubidjan;		// 매수 잔량
	char fujbidjan[6];	char _fujbidjan;		// 차선 매수 호가잔량
	char fujjbidjan[6];	char _fujjbidjan;		// 차차선 매수 호가잔량
	char fuj4bidjan[6];	char _fuj4bidjan;		// 4차선 매수 호가잔량
	char fuj5bidjan[6];	char _fuj5bidjan;		// 5차선 매수 호가잔량
	char futofferjan[6];	char _futofferjan;		// 총 매도 잔량
	char futbidjan[6];	char _futbidjan;		// 총 매수 잔량
	char fuoffersu[4];	char _fuoffersu;		// 매도 최우선건수
	char fujoffersu[4];	char _fujoffersu;		// 매도 차선건수
	char fujjoffersu[4];	char _fujjoffersu;		// 매도 차차선건수
	char fuj4offersu[4];	char _fuj4offersu;		// 매도 4차선건수
	char fuj5offersu[4];	char _fuj5offersu;		// 매도 5차선건수
	char fubidsu[4];	char _fubidsu;		// 매수 최우선건수
	char fujbidsu[4];	char _fujbidsu;		// 매수 차선건수
	char fujjbidsu[4];	char _fujjbidsu;		// 매수 차차선건수
	char fuj4bidsu[4];	char _fuj4bidsu;		// 매수 4차선건수
	char fuj5bidsu[4];	char _fuj5bidsu;		// 매수 5차선건수
	char futoffersu[5];	char _futoffersu;		// 총 매도 건수
	char futbidsu[5];	char _futbidsu;		// 총 매수 건수
	char fuhname[12];	char _fuhname;		// 한글명
} Tc4113OutHoga2;

typedef struct tagc4113OutHoga3// 코스피선물스프레드호가3
{
	char fuspfuitem[8];	char _fuspfuitem;		// 종목코드
	char fusphname[14];	char _fusphname;		// 한글명
	char fusphotime[8];	char _fusphotime;		// 호가시간
	char fuspoffer[6];	char _fuspoffer;		// 매도 우선호가
	char fuspjoffer[6];	char _fuspjoffer;		// 차선 매도 호가
	char fuspjjoffer[6];	char _fuspjjoffer;		// 차차선 매도 호가
	char fuspj4offer[6];	char _fuspj4offer;		// 4차선 매도 호가
	char fuspj5offer[6];	char _fuspj5offer;		// 5차선 매도 호가
	char fuspofferjan[6];	char _fuspofferjan;		// 매도 잔량
	char fuspjofferjan[6];	char _fuspjofferjan;		// 차선 매도 호가잔량
	char fuspjjofferjan[6];	char _fuspjjofferjan;		// 차차선 매도 호가잔량
	char fuspj4offerjan[6];	char _fuspj4offerjan;		// 4차선 매도 호가잔량
	char fuspj5offerjan[6];	char _fuspj5offerjan;		// 5차선 매도 호가잔량
	char fuspbid[6];	char _fuspbid;		// 매수 우선호가
	char fuspjbid[6];	char _fuspjbid;		// 차선 매수 호가
	char fuspjjbid[6];	char _fuspjjbid;		// 차차선 매수 호가
	char fuspj4bid[6];	char _fuspj4bid;		// 4차선 매수 호가
	char fuspj5bid[6];	char _fuspj5bid;		// 5차선 매수 호가
	char fuspbidjan[6];	char _fuspbidjan;		// 매수 잔량
	char fuspjbidjan[6];	char _fuspjbidjan;		// 차선 매수 호가잔량
	char fuspjjbidjan[6];	char _fuspjjbidjan;		// 차차선 매수 호가잔량
	char fuspj4bidjan[6];	char _fuspj4bidjan;		// 4차선 매수 호가잔량
	char fuspj5bidjan[6];	char _fuspj5bidjan;		// 5차선 매수 호가잔량
	char fusptofferjan[6];	char _fusptofferjan;		// 총 매도 잔량
	char fusptbidjan[6];	char _fusptbidjan;		// 총 매수 잔량
	char fuspoffersu[4];	char _fuspoffersu;		// 매도 최우선건수
	char fuspjoffersu[4];	char _fuspjoffersu;		// 매도 차선건수
	char fuspjjoffersu[4];	char _fuspjjoffersu;		// 매도 차차선건수
	char fuspj4offersu[4];	char _fuspj4offersu;		// 매도 4차선건수
	char fuspj5offersu[4];	char _fuspj5offersu;		// 매도 5차선건수
	char fuspbidsu[4];	char _fuspbidsu;		// 매수 최우선건수
	char fuspjbidsu[4];	char _fuspjbidsu;		// 매수 차선건수
	char fuspjjbidsu[4];	char _fuspjjbidsu;		// 매수 차차선건수
	char fuspj4bidsu[4];	char _fuspj4bidsu;		// 매수 4차선건수
	char fuspj5bidsu[4];	char _fuspj5bidsu;		// 매수 5차선건수
	char fusptoffersu[5];	char _fusptoffersu;		// 총 매도 건수
	char fusptbidsu[5];	char _fusptbidsu;		// 총 매수 건수
} Tc4113OutHoga3;

typedef struct tagc4113OutFuteq1// 선물예상체결1
{
	char dongsi[1];	char _dongsi;		// 동시호가구분
	char jeqprice[5];	char _jeqprice;		// 예상체결가
	char jeqsign[1];	char _jeqsign;		// 예상체결부호
	char jeqchange[5];	char _jeqchange;		// 예상체결등락폭
	char jeqchrate[5];	char _jeqchrate;		// 예상체결등락률
} Tc4113OutFuteq1;

typedef struct tagc4113OutFuteq2// 선물예상체결2
{
	char dongsi[1];	char _dongsi;		// 동시호가구분
	char jeqprice[5];	char _jeqprice;		// 예상체결가
	char jeqsign[1];	char _jeqsign;		// 예상체결부호
	char jeqchange[5];	char _jeqchange;		// 예상체결등락폭
	char jeqchrate[5];	char _jeqchrate;		// 예상체결등락률
} Tc4113OutFuteq2;

typedef struct tagc4113
{
	Tc4113InBlock                     c4113inblock                          ;		// 입력데이타
	Tc4113OutKospi200                 c4113outkospi200                      ;		// c4113OutKospi200
	Tc4113OutSMaster                  c4113outsmaster                       ;		// c4113OutSMaster
	Tc4113OutBlock1                   c4113outblock1                        ;		// 코스피선물Master1
	Tc4113OutBlock2                   c4113outblock2                        ;		// 코스피선물Master2
	Tc4113OutHoga1                    c4113outhoga1                         ;		// 코스피선물호가1
	Tc4113OutHoga2                    c4113outhoga2                         ;		// 코스피선물호가2
	Tc4113OutHoga3                    c4113outhoga3                         ;		// 코스피선물스프레드호가3
	Tc4113OutFuteq1                   c4113outfuteq1                        ;		// 선물예상체결1
	Tc4113OutFuteq2                   c4113outfuteq2                        ;		// 선물예상체결2
} Tc4113;

typedef struct tags4201InBlock// 기본입력
{
	char Lang[1];	char _Lang;		// 한영구분
	char opitemz9[9];	char _opitemz9;		// 종목코드
} Ts4201InBlock;

typedef struct tags4201OutBlock// 종목마스타기본자료
{
	char opitem[8];	char _opitem;		// 종목코드
	char ophname[14];	char _ophname;		// 종목명
	char opcurr[5];	char _opcurr;		// 현재가
	char opsign[1];	char _opsign;		// 등락부호
	char opchange[5];	char _opchange;		// 등락폭
	char opchrate[5];	char _opchrate;		// 등락률
	char opopen[5];	char _opopen;		// 시가
	char ophigh[5];	char _ophigh;		// 고가
	char oplow[5];	char _oplow;		// 저가
	char optheoryprice[5];	char _optheoryprice;		// 이론가
	char opvolallz8[8];	char _opvolallz8;		// 거래량
	char opvalall[12];	char _opvalall;		// 누적거래대금(백만)
	char opopenyak[7];	char _opopenyak;		// 미결제약정수량
	char oppreopenyak[7];	char _oppreopenyak;		// 미결제약정전일
	char oplisthdatez11[11];	char _oplisthdatez11;		// 상장후최고일
	char oplistldatez11[11];	char _oplistldatez11;		// 상장후최저일
	char oplistdate[8];	char _oplistdate;		// 거래개시일
	char oplastdate[8];	char _oplastdate;		// 최종거래일
	char opjandatecnt[4];	char _opjandatecnt;		// 잔존일
	char ophprice[5];	char _ophprice;		// 상한가
	char oplprice[5];	char _oplprice;		// 하한가
	char opgrate[5];	char _opgrate;		// 괴리도
	char opimpv[5];	char _opimpv;		// 내재변동성
	char oppastv90[5];	char _oppastv90;		// 과거변동성90
	char opdelta[8];	char _opdelta;		// 델타지수
	char opgmma[8];	char _opgmma;		// 감마지수
	char opvega[8];	char _opvega;		// 베가변동성
	char optheta[8];	char _optheta;		// 쎄타시간
	char oprho[8];	char _oprho;		// 로이자율
	char opcdratio[6];	char _opcdratio;		// 이자율
	char opdivideratio[9];	char _opdivideratio;		// 배당액지수
	char opchetime[8];	char _opchetime;		// 호가시간
	char opoffer[5];	char _opoffer;		// 매도 최우선호가
	char opjoffer[5];	char _opjoffer;		// 매도 차선 호가
	char opjjoffer[5];	char _opjjoffer;		// 매도 차차선 호가
	char opj4offer[5];	char _opj4offer;		// 매도 4차선 호가
	char opj5offer[5];	char _opj5offer;		// 매도 5차선 호가
	char opbid[5];	char _opbid;		// 매수 최우선호가
	char opjbid[5];	char _opjbid;		// 매수 차선 호가
	char opjjbid[5];	char _opjjbid;		// 매수 차차선 호가
	char opj4bid[5];	char _opj4bid;		// 매수 4차선 호가
	char opj5bid[5];	char _opj5bid;		// 매수 5차선 호가
	char opofferjan[7];	char _opofferjan;		// 매도 최우선 잔량
	char opjofferjan[7];	char _opjofferjan;		// 매도 차선 잔량
	char opjjofferjan[7];	char _opjjofferjan;		// 매도 차차선 잔량
	char opj4offerjan[7];	char _opj4offerjan;		// 매도 4차선 잔량
	char opj5offerjan[7];	char _opj5offerjan;		// 매도 5차선 잔량
	char opbidjan[7];	char _opbidjan;		// 매수 최우선 잔량
	char opjbidjan[7];	char _opjbidjan;		// 매수 차선 잔량
	char opjjbidjan[7];	char _opjjbidjan;		// 매수 차차선 잔량
	char opj4bidjan[7];	char _opj4bidjan;		// 매수 4차선 잔량
	char opj5bidjan[7];	char _opj5bidjan;		// 매수 5차선 잔량
	char optofferjan[7];	char _optofferjan;		// 총 매도 잔량
	char optbidjan[7];	char _optbidjan;		// 총 매수 잔량
	char opoffersu[4];	char _opoffersu;		// 매도 최우선건수
	char opjoffersu[4];	char _opjoffersu;		// 매도 차선건수
	char opjjoffersu[4];	char _opjjoffersu;		// 매도 차차선건수
	char opj4offersu[4];	char _opj4offersu;		// 매도 4차선건수
	char opj5offersu[4];	char _opj5offersu;		// 매도 5차선건수
	char opbidsu[4];	char _opbidsu;		// 매수 최우선건수
	char opjbidsu[4];	char _opjbidsu;		// 매수 차선건수
	char opjjbidsu[4];	char _opjjbidsu;		// 매수 차차선건수
	char opj4bidsu[4];	char _opj4bidsu;		// 매수 4차선건수
	char opj5bidsu[4];	char _opj5bidsu;		// 매수 5차선건수
	char optoffersu[5];	char _optoffersu;		// 총 매도 건수
	char optbidsu[5];	char _optbidsu;		// 총 매수 건수
	char opjgubun[8];	char _opjgubun;		// CB발동여부
	char opopensign[1];	char _opopensign;		// 시가대비부호
	char opopenchange[5];	char _opopenchange;		// 시가대비등락
	char opgratio[5];	char _opgratio;		// 괴리율
	char preclose[5];	char _preclose;		// 전일종가
	char fudynhprice[5];	char _fudynhprice;		// 실시간상한가
	char fudynlprice[5];	char _fudynlprice;		// 실시간하한가
	char fudynpriceflag[1];	char _fudynpriceflag;		// 동적가격제한여부
	char opbulkvol[8];	char _opbulkvol;		// 협의거래량
	char uplmtgb[1];	char _uplmtgb;		// 가격제한확대상한단계
	char dnlmtgb[1];	char _dnlmtgb;		// 가격제한확대하한단계
} Ts4201OutBlock;

typedef struct tags4201OutBlock1// 코스피200지수
{
	char fuitem[4];	char _fuitem;		// 코스피200코드
	char fucurr[5];	char _fucurr;		// 코스피200지수
	char fusign[1];	char _fusign;		// 코스피200등락부호
	char fuchange[5];	char _fuchange;		// 코스피200등락폭
	char fuchrate[5];	char _fuchrate;		// 코스피200등락률
} Ts4201OutBlock1;

typedef struct tags4201OutBlock2// 옵션변동거래량자료,[반복]
{
	char opchetime[8];	char _opchetime;		// 시간
	char opcurr[5];	char _opcurr;		// 현재가
	char opsign[1];	char _opsign;		// 등락부호
	char opchange[5];	char _opchange;		// 등락폭
	char opoffer[5];	char _opoffer;		// 매도 호가
	char opbid[5];	char _opbid;		// 매수 호가
	char opvol[6];	char _opvol;		// 거래량
	char opvolallz8[8];	char _opvolallz8;		// 누적거래량
	char opopenyak[7];	char _opopenyak;		// 미결제약정
} Ts4201OutBlock2;

typedef struct tags4201OutBlock3// 선물최근월물
{
	char fuitem[4];	char _fuitem;		// 선물최근월물코드
	char fuitemz9[9];	char _fuitemz9;		// 선물최근월물확장코드
	char fuhname[12];	char _fuhname;		// 선물최근월물명
	char fucurr[5];	char _fucurr;		// 선물최근월물지수
	char fusign[1];	char _fusign;		// 선물최근월물등락부호
	char fuchange[5];	char _fuchange;		// 선물최근월물등락폭
	char fuchrate[5];	char _fuchrate;		// 선물최근월물등락률
	char fuvolall[7];	char _fuvolall;		// 선물최근월물거래량
	char fuvalall[12];	char _fuvalall;		// 선물최근월물누적거래대금(백만)
	char fuchetime[8];	char _fuchetime;		// 선물최근월물호가시간
	char fuoffer[5];	char _fuoffer;		// 선물최근월물 매도 최우선호가
	char fujoffer[5];	char _fujoffer;		// 선물최근월물 매도 차선 호가
	char fujjoffer[5];	char _fujjoffer;		// 선물최근월물 매도 차차선 호가
	char fuj4offer[5];	char _fuj4offer;		// 선물최근월물 매도 4차선 호가
	char fuj5offer[5];	char _fuj5offer;		// 선물최근월물 매도 5차선 호가
	char fubid[5];	char _fubid;		// 선물최근월물 매수 최우선호가
	char fujbid[5];	char _fujbid;		// 선물최근월물 매수 차선 호가
	char fujjbid[5];	char _fujjbid;		// 선물최근월물 매수 차차선 호가
	char fuj4bid[5];	char _fuj4bid;		// 선물최근월물 매수 4차선 호가
	char fuj5bid[5];	char _fuj5bid;		// 선물최근월물 매수 5차선 호가
	char fuofferjan[6];	char _fuofferjan;		// 선물최근월물 매도 최우선 잔량
	char fujofferjan[6];	char _fujofferjan;		// 선물최근월물 매도 차선 잔량
	char fujjofferjan[6];	char _fujjofferjan;		// 선물최근월물 매도 차차선 잔량
	char fuj4offerjan[6];	char _fuj4offerjan;		// 선물최근월물 매도 4차선 잔량
	char fuj5offerjan[6];	char _fuj5offerjan;		// 선물최근월물 매도 5차선 잔량
	char fubidjan[6];	char _fubidjan;		// 선물최근월물 매수 최우선 잔량
	char fujbidjan[6];	char _fujbidjan;		// 선물최근월물 매수 차선 잔량
	char fujjbidjan[6];	char _fujjbidjan;		// 선물최근월물 매수 차차선 잔량
	char fuj4bidjan[6];	char _fuj4bidjan;		// 선물최근월물 매수 4차선 잔량
	char fuj5bidjan[6];	char _fuj5bidjan;		// 선물최근월물 매수 5차선 잔량
	char futofferjan[6];	char _futofferjan;		// 선물최근월물총 매도 잔량
	char futbidjan[6];	char _futbidjan;		// 선물최근월물총 매수 잔량
	char fuoffersu[4];	char _fuoffersu;		// 선물최근월물 매도 최우선건수
	char fujoffersu[4];	char _fujoffersu;		// 선물최근월물 매도 차선건수
	char fujjoffersu[4];	char _fujjoffersu;		// 선물최근월물 매도 차차선건수
	char fuj4offersu[4];	char _fuj4offersu;		// 선물최근월물 매도 4차선건수
	char fuj5offersu[4];	char _fuj5offersu;		// 선물최근월물 매도 5차선건수
	char fubidsu[4];	char _fubidsu;		// 선물최근월물 매수 최우선건수
	char fujbidsu[4];	char _fujbidsu;		// 선물최근월물 매수 차선건수
	char fujjbidsu[4];	char _fujjbidsu;		// 선물최근월물 매수 차차선건수
	char fuj4bidsu[4];	char _fuj4bidsu;		// 선물최근월물 매수 4차선건수
	char fuj5bidsu[4];	char _fuj5bidsu;		// 선물최근월물 매수 5차선건수
	char futoffersu[5];	char _futoffersu;		// 선물최근월물총 매도 건수
	char futbidsu[5];	char _futbidsu;		// 선물최근월물총 매수 건수
	char fuhprice[5];	char _fuhprice;		// 상한가
	char fuhigh[5];	char _fuhigh;		// 고가
	char fuopen[5];	char _fuopen;		// 시가
	char fuopensign[1];	char _fuopensign;		// 시가대비부호
	char fuopenchange[5];	char _fuopenchange;		// 시가대비등락
	char fulow[5];	char _fulow;		// 저가
	char fulprice[5];	char _fulprice;		// 하한가
	char fupivot2upz5[5];	char _fupivot2upz5;		// 피봇2차저항
	char fupivot1upz5[5];	char _fupivot1upz5;		// 피봇1차저항
	char fupivotz5[5];	char _fupivotz5;		// 피봇가
	char fupivot1dnz5[5];	char _fupivot1dnz5;		// 피봇1차지지
	char fupivot2dnz5[5];	char _fupivot2dnz5;		// 피봇2차지지
	char fudynhprice[5];	char _fudynhprice;		// 실시간상한가
	char fudynlprice[5];	char _fudynlprice;		// 실시간하한가
	char fudynpriceflag[1];	char _fudynpriceflag;		// 동적가격제한여부
} Ts4201OutBlock3;

typedef struct tags4201OutBlock4// 선물변동거래량자료,[반복]
{
	char fuchetime[8];	char _fuchetime;		// 시간
	char fucurr[5];	char _fucurr;		// 현재가
	char fusign[1];	char _fusign;		// 등락부호
	char fuchange[5];	char _fuchange;		// 등락폭
	char fuoffer[5];	char _fuoffer;		// 매도 호가
	char fubid[5];	char _fubid;		// 매수 호가
	char fuvol[6];	char _fuvol;		// 거래량
	char fuvolall[7];	char _fuvolall;		// 누적거래량
	char fuopenyak[7];	char _fuopenyak;		// 미결제약정
} Ts4201OutBlock4;

typedef struct tags4201OutBlock5// 옵션예상체결
{
	char dongsi[1];	char _dongsi;		// 동시호가구분
	char jeqprice[5];	char _jeqprice;		// 예상체결가
	char jeqsign[1];	char _jeqsign;		// 예상체결부호
	char jeqchange[5];	char _jeqchange;		// 예상체결등락폭
	char jeqchrate[5];	char _jeqchrate;		// 예상체결등락률
} Ts4201OutBlock5;

typedef struct tags4201OutBlock6// 선물예상체결
{
	char dongsi[1];	char _dongsi;		// 동시호가구분
	char jeqprice[5];	char _jeqprice;		// 예상체결가
	char jeqsign[1];	char _jeqsign;		// 예상체결부호
	char jeqchange[5];	char _jeqchange;		// 예상체결등락폭
	char jeqchrate[5];	char _jeqchrate;		// 예상체결등락률
} Ts4201OutBlock6;

typedef struct tags4201
{
	Ts4201InBlock                     s4201inblock                          ;		// 기본입력
	Ts4201OutBlock                    s4201outblock                         ;		// 종목마스타기본자료
	Ts4201OutBlock1                   s4201outblock1                        ;		// 코스피200지수
	Ts4201OutBlock2                   s4201outblock2[20];		// 옵션변동거래량자료 ,[반복]
	Ts4201OutBlock3                   s4201outblock3                        ;		// 선물최근월물
	Ts4201OutBlock4                   s4201outblock4[20];		// 선물변동거래량자료 ,[반복]
	Ts4201OutBlock5                   s4201outblock5                        ;		// 옵션예상체결
	Ts4201OutBlock6                   s4201outblock6                        ;		// 선물예상체결
} Ts4201;


typedef struct tagc4801InBlock// 기본입력
{
	char Lang[1];	char _Lang;		// 한영구분
	char fuitemz9[9];	char _fuitemz9;		// 종목코드
} Tc4801InBlock;

typedef struct tagc4801OutBlock// 주식선물MASTER기본자료
{
	char expcode[8];	char _expcode;		// 종목코드
	char Title[50];	char _Title;		// 한글명
	char ename[50];	char _ename;		// 영문명
	char sname[25];	char _sname;		// 단축명
	char baseprice[7];	char _baseprice;		// 기준가격
	char hprice[7];	char _hprice;		// 상한가
	char lprice[7];	char _lprice;		// 하한가
	char preclose[7];	char _preclose;		// 전일종가
	char unit[16];	char _unit;		// 거래단위
	char openyak[7];	char _openyak;		// 미결제약정수량
	char fusign[1];	char _fusign;		// 전일대비부호
	char fuchange[7];	char _fuchange;		// 전일대비
	char fucurr[7];	char _fucurr;		// 현재가
	char fuopen[7];	char _fuopen;		// 시가
	char fuhigh[7];	char _fuhigh;		// 고가
	char fulow[7];	char _fulow;		// 저가
	char fuvolall[7];	char _fuvolall;		// 누적체결수량(계약)
	char fuspvolall[7];	char _fuspvolall;		// 스프레드체결수량
	char fuvalall[12];	char _fuvalall;		// 누적거래대금(천원)
	char hotime[8];	char _hotime;		// 호가시간
	char offer[7];	char _offer;		// 매도 우선호가
	char bid[7];	char _bid;		// 매수 우선호가
	char offerjan[6];	char _offerjan;		// 매도 잔량
	char bidjan[6];	char _bidjan;		// 매수 잔량
	char S2offer[7];	char _S2offer;		// 매도 2차호가
	char S2bid[7];	char _S2bid;		// 매수 2차호가
	char S2offerjan[6];	char _S2offerjan;		// 매도 2차잔량
	char S2bidjan[6];	char _S2bidjan;		// 매수 2차잔량
	char S3offer[7];	char _S3offer;		// 매도 3차호가
	char S3bid[7];	char _S3bid;		// 매수 3차호가
	char S3offerjan[6];	char _S3offerjan;		// 매도 3차잔량
	char S3bidjan[6];	char _S3bidjan;		// 매수 3차잔량
	char S4offer[7];	char _S4offer;		// 매도 4차호가
	char S4bid[7];	char _S4bid;		// 매수 4차호가
	char S4offerjan[6];	char _S4offerjan;		// 매도 4차잔량
	char S4bidjan[6];	char _S4bidjan;		// 매수 4차잔량
	char S5offer[7];	char _S5offer;		// 매도 5차호가
	char S5bid[7];	char _S5bid;		// 매수 5차호가
	char S5offerjan[6];	char _S5offerjan;		// 매도 5차잔량
	char S5bidjan[6];	char _S5bidjan;		// 매수 5차잔량
	char S6offer[7];	char _S6offer;		// 매도 6차호가
	char S6bid[7];	char _S6bid;		// 매수 6차호가
	char S6offerjan[6];	char _S6offerjan;		// 매도 6차잔량
	char S6bidjan[6];	char _S6bidjan;		// 매수 6차잔량
	char S7offer[7];	char _S7offer;		// 매도 7차호가
	char S7bid[7];	char _S7bid;		// 매수 7차호가
	char S7offerjan[6];	char _S7offerjan;		// 매도 7차잔량
	char S7bidjan[6];	char _S7bidjan;		// 매수 7차잔량
	char S8offer[7];	char _S8offer;		// 매도 8차호가
	char S8bid[7];	char _S8bid;		// 매수 8차호가
	char S8offerjan[6];	char _S8offerjan;		// 매도 8차잔량
	char S8bidjan[6];	char _S8bidjan;		// 매수 8차잔량
	char S9offer[7];	char _S9offer;		// 매도 9차호가
	char S9bid[7];	char _S9bid;		// 매수 9차호가
	char S9offerjan[6];	char _S9offerjan;		// 매도 9차잔량
	char S9bidjan[6];	char _S9bidjan;		// 매수 9차잔량
	char S0offer[7];	char _S0offer;		// 매도 10차호가
	char S0bid[7];	char _S0bid;		// 매수 10차호가
	char S0offerjan[6];	char _S0offerjan;		// 매도 10차잔량
	char S0bidjan[6];	char _S0bidjan;		// 매수 10차잔량
	char offersu[4];	char _offersu;		// 매도 건수
	char bidsu[4];	char _bidsu;		// 매수 건수
	char S2offersu[4];	char _S2offersu;		// 매도 2차건수
	char S2bidsu[4];	char _S2bidsu;		// 매수 2차건수
	char S3offersu[4];	char _S3offersu;		// 매도 3차건수
	char S3bidsu[4];	char _S3bidsu;		// 매수 3차건수
	char S4offersu[4];	char _S4offersu;		// 매도 4차건수
	char S4bidsu[4];	char _S4bidsu;		// 매수 4차건수
	char S5offersu[4];	char _S5offersu;		// 매도 5차건수
	char S5bidsu[4];	char _S5bidsu;		// 매수 5차건수
	char S6offersu[4];	char _S6offersu;		// 매도 6차건수
	char S6bidsu[4];	char _S6bidsu;		// 매수 6차건수
	char S7offersu[4];	char _S7offersu;		// 매도 7차건수
	char S7bidsu[4];	char _S7bidsu;		// 매수 7차건수
	char S8offersu[4];	char _S8offersu;		// 매도 8차건수
	char S8bidsu[4];	char _S8bidsu;		// 매수 8차건수
	char S9offersu[4];	char _S9offersu;		// 매도 9차건수
	char S9bidsu[4];	char _S9bidsu;		// 매수 9차건수
	char S0offersu[4];	char _S0offersu;		// 매도 10차건수
	char S0bidsu[4];	char _S0bidsu;		// 매수 10차건수
	char tofferjan[6];	char _tofferjan;		// 총 매도 잔량
	char tobidjan[6];	char _tobidjan;		// 총 매수 잔량
	char toffersu[5];	char _toffersu;		// 총 매도 건수
	char tbidsu[5];	char _tbidsu;		// 총 매수 건수
	char theorytime[6];	char _theorytime;		// 이론가시간
	char theoryprice[7];	char _theoryprice;		// 이론가
	char fuchrate[5];	char _fuchrate;		// 등락률
	char fupivot2upz7[7];	char _fupivot2upz7;		// 피봇2차저항
	char fupivot1upz7[7];	char _fupivot1upz7;		// 피봇1차저항
	char fupivotz7[7];	char _fupivotz7;		// 피봇가
	char fupivot1dnz7[7];	char _fupivot1dnz7;		// 피봇1차지지
	char fupivot2dnz7[7];	char _fupivot2dnz7;		// 피봇2차지지
	char fubasis[7];	char _fubasis;		// 베이시스
	char fugrate[7];	char _fugrate;		// 괴리도
	char fugratio[6];	char _fugratio;		// 괴리율
	char fupreopenyak[7];	char _fupreopenyak;		// 미결제약정전일
	char fulisthprice[7];	char _fulisthprice;		// 상장후최고가
	char fulisthdate[8];	char _fulisthdate;		// 상장후최고일
	char fulistlprice[7];	char _fulistlprice;		// 상장후최저가
	char fulistldate[8];	char _fulistldate;		// 상장후최저일
	char fulastdate[8];	char _fulastdate;		// 최종거래일
	char fujandatecnt[3];	char _fujandatecnt;		// 잔존일
	char fucdratio[6];	char _fucdratio;		// 무위험이자율
	char fuopenchange[7];	char _fuopenchange;		// 시가대비등락
	char fudynhprice[7];	char _fudynhprice;		// 실시간상한가
	char fudynlprice[7];	char _fudynlprice;		// 실시간하한가
	char fudynpriceflag[1];	char _fudynpriceflag;		// 동적가격제한여부
	char exlmtgb[1];	char _exlmtgb;		// 가격확대예정구분
	char uplmtgb[1];	char _uplmtgb;		// 가격제한확대상한단계
	char dnlmtgb[1];	char _dnlmtgb;		// 가격제한확대하한단계
} Tc4801OutBlock;

typedef struct tagc4801OutBlock1// 기초자산
{
	char shcode[6];	char _shcode;		// 종목코드
	char Title[20];	char _Title;		// 종목명
	char price[7];	char _price;		// 현재가
	char sign[1];	char _sign;		// 등락부호
	char change[6];	char _change;		// 등락폭
	char market[16];	char _market;		// 락구분
	char chrate[5];	char _chrate;		// 등락률
	char offer[7];	char _offer;		// 매도 호가
	char bid[7];	char _bid;		// 매수 호가
	char volume[9];	char _volume;		// 거래량
	char volrate[6];	char _volrate;		// 거래비율
	char uplmtprice[7];	char _uplmtprice;		// 상한가
	char high[7];	char _high;		// 고가
	char open[7];	char _open;		// 시가
	char low[7];	char _low;		// 저가
	char dnlmtprice[7];	char _dnlmtprice;		// 하한가
} Tc4801OutBlock1;

typedef struct tagc4801OutBlock2// 주식선물예상체결
{
	char dongsi[1];	char _dongsi;		// 동시호가구분
	char jeqprice[7];	char _jeqprice;		// 예상체결가
	char jeqsign[1];	char _jeqsign;		// 예상체결부호
	char jeqchange[7];	char _jeqchange;		// 예상체결등락폭
	char jeqchrate[5];	char _jeqchrate;		// 예상체결등락률
} Tc4801OutBlock2;

typedef struct tagc4801
{
	Tc4801InBlock                     c4801inblock                          ;		// 기본입력
	Tc4801OutBlock                    c4801outblock                         ;		// 주식선물MASTER기본자료
	Tc4801OutBlock1                   c4801outblock1                        ;		// 기초자산
	Tc4801OutBlock2                   c4801outblock2                        ;		// 주식선물예상체결
} Tc4801;


typedef struct tagc4805InBlock// 입력데이타
{
	char fuitemz9[9];	char _fuitemz9;		// 입력코드
} Tc4805InBlock;

typedef struct tagc4805OutUnder// c4805OutUnder
{
	char shcode[6];	char _shcode;		// 종목코드
	char Title[20];	char _Title;		// 종목명
	char price[7];	char _price;		// 현재가
	char sign[1];	char _sign;		// 전일비부호
	char change[6];	char _change;		// 전일비
} Tc4805OutUnder;

typedef struct tagc4805OutSMaster// c4805OutSMaster
{
	char fuitemz8[8];	char _fuitemz8;		// 종목코드
	char fuspcurr[8];	char _fuspcurr;		// 지수
	char fuspsign[1];	char _fuspsign;		// 전일비부호
	char fuspchange[7];	char _fuspchange;		// 전일비
	char fuspchrate[5];	char _fuspchrate;		// 등락률
	char fuspopen[7];	char _fuspopen;		// 시가
	char fusphigh[7];	char _fusphigh;		// 고가
	char fusplow[7];	char _fusplow;		// 저가
	char fuspvolall[7];	char _fuspvolall;		// 거래량
	char fuspvalall[12];	char _fuspvalall;		// 누적거래대금(천원)
	char fuspcurr1[7];	char _fuspcurr1;		// 의제약정가(근월물)
	char fuspcurr2[7];	char _fuspcurr2;		// 의제약정가(원월물)
	char fuitem1[8];	char _fuitem1;		// 종목코드(근월물)
	char fuitem2[8];	char _fuitem2;		// 종목코드(원월물)
	char fudynhprice[7];	char _fudynhprice;		// 실시간상한가
	char fudynlprice[7];	char _fudynlprice;		// 실시간하한가
	char fudynpriceflag[1];	char _fudynpriceflag;		// 동적가격제한여부
} Tc4805OutSMaster;

typedef struct tagc4805OutHoga3// 주식선물호가3
{
	char fuspfuitem[8];	char _fuspfuitem;		// 종목코드
	char fusphname[50];	char _fusphname;		// 한글명
	char fusphotime[8];	char _fusphotime;		// 호가시간
	char offer[8];	char _offer;		// 매도 우선호가
	char bid[8];	char _bid;		// 매수 우선호가
	char offerjan[6];	char _offerjan;		// 매도 잔량
	char bidjan[6];	char _bidjan;		// 매수 잔량
	char S2offer[8];	char _S2offer;		// 2차 매도 호가
	char S2bid[8];	char _S2bid;		// 2차 매수 호가
	char S2offerjan[6];	char _S2offerjan;		// 2차 매도 잔량
	char S2bidjan[6];	char _S2bidjan;		// 2차 매수 잔량
	char S3offer[8];	char _S3offer;		// 3차 매도 호가
	char S3bid[8];	char _S3bid;		// 3차 매수 호가
	char S3offerjan[6];	char _S3offerjan;		// 3차 매도 잔량
	char S3bidjan[6];	char _S3bidjan;		// 3차 매수 잔량
	char S4offer[8];	char _S4offer;		// 4차 매도 호가
	char S4bid[8];	char _S4bid;		// 4차 매수 호가
	char S4offerjan[6];	char _S4offerjan;		// 4차 매도 잔량
	char S4bidjan[6];	char _S4bidjan;		// 4차 매수 잔량
	char S5offer[8];	char _S5offer;		// 5차 매도 호가
	char S5bid[8];	char _S5bid;		// 5차 매수 호가
	char S5offerjan[6];	char _S5offerjan;		// 5차 매도 잔량
	char S5bidjan[6];	char _S5bidjan;		// 5차 매수 잔량
	char S6offer[8];	char _S6offer;		// 6차 매도 호가
	char S6bid[8];	char _S6bid;		// 6차 매수 호가
	char S6offerjan[6];	char _S6offerjan;		// 6차 매도 잔량
	char S6bidjan[6];	char _S6bidjan;		// 6차 매수 잔량
	char S7offer[8];	char _S7offer;		// 7차 매도 호가
	char S7bid[8];	char _S7bid;		// 7차 매수 호가
	char S7offerjan[6];	char _S7offerjan;		// 7차 매도 잔량
	char S7bidjan[6];	char _S7bidjan;		// 7차 매수 잔량
	char S8offer[8];	char _S8offer;		// 8차 매도 호가
	char S8bid[8];	char _S8bid;		// 8차 매수 호가
	char S8offerjan[6];	char _S8offerjan;		// 8차 매도 잔량
	char S8bidjan[6];	char _S8bidjan;		// 8차 매수 잔량
	char S9offer[8];	char _S9offer;		// 9차 매도 호가
	char S9bid[8];	char _S9bid;		// 9차 매수 호가
	char S9offerjan[6];	char _S9offerjan;		// 9차 매도 잔량
	char S9bidjan[6];	char _S9bidjan;		// 9차 매수 잔량
	char S0offer[8];	char _S0offer;		// 10차 매도 호가
	char S0bid[8];	char _S0bid;		// 10차 매수 호가
	char S0offerjan[6];	char _S0offerjan;		// 10차 매도 잔량
	char S0bidjan[6];	char _S0bidjan;		// 10차 매수 잔량
	char offersu[4];	char _offersu;		// 매도 건수
	char bidsu[4];	char _bidsu;		// 매수 건수
	char S2offersu[4];	char _S2offersu;		// 2차 매도 건수
	char S2bidsu[4];	char _S2bidsu;		// 2차 매수 건수
	char S3offersu[4];	char _S3offersu;		// 3차 매도 건수
	char S3bidsu[4];	char _S3bidsu;		// 3차 매수 건수
	char S4offersu[4];	char _S4offersu;		// 4차 매도 건수
	char S4bidsu[4];	char _S4bidsu;		// 4차 매수 건수
	char S5offersu[4];	char _S5offersu;		// 5차 매도 건수
	char S5bidsu[4];	char _S5bidsu;		// 5차 매수 건수
	char S6offersu[4];	char _S6offersu;		// 6차 매도 건수
	char S6bidsu[4];	char _S6bidsu;		// 6차 매수 건수
	char S7offersu[4];	char _S7offersu;		// 7차 매도 건수
	char S7bidsu[4];	char _S7bidsu;		// 7차 매수 건수
	char S8offersu[4];	char _S8offersu;		// 8차 매도 건수
	char S8bidsu[4];	char _S8bidsu;		// 8차 매수 건수
	char S9offersu[4];	char _S9offersu;		// 9차 매도 건수
	char S9bidsu[4];	char _S9bidsu;		// 9차 매수 건수
	char S0offersu[4];	char _S0offersu;		// 10차 매도 건수
	char S0bidsu[4];	char _S0bidsu;		// 10차 매수 건수
	char tofferjan[6];	char _tofferjan;		// 총 매도 잔량
	char tobidjan[6];	char _tobidjan;		// 총 매수 잔량
	char toffersu[5];	char _toffersu;		// 총 매도 건수
	char tbidsu[5];	char _tbidsu;		// 총 매수 건수
	char undershcode[6];	char _undershcode;		// 기초자산종목코드
	char underhname[20];	char _underhname;		// 기초자산종목명
	char eitem[2];	char _eitem;		// 기초대상주식
	char lgcode[9];	char _lgcode;		// lgcode
	char bp_jgubun[1];	char _bp_jgubun;		// BP용장구분
} Tc4805OutHoga3;

typedef struct tagc4805OutSpread// 선물SPREAD
{
	char thspread[7];	char _thspread;		// 이론스프레드
	char respread[7];	char _respread;		// 실제스프레드
	char fugrate1[7];	char _fugrate1;		// 괴리
} Tc4805OutSpread;

typedef struct tagc4805
{
	Tc4805InBlock                     c4805inblock                          ;		// 입력데이타
	Tc4805OutUnder                    c4805outunder                         ;		// c4805OutUnder
	Tc4805OutSMaster                  c4805outsmaster                       ;		// c4805OutSMaster
	Tc4805OutHoga3                    c4805outhoga3                         ;		// 주식선물호가3
	Tc4805OutSpread                   c4805outspread                        ;		// 선물SPREAD
} Tc4805;

typedef struct tags1701InBlock// 기본입력
{
	char Code[6];	char _Code;		// 종목코드
} Ts1701InBlock;

typedef struct tags1701OutBlock// 종목마스타기본자료
{
	char Code[6];	char _Code;		// 종목코드
	char Title[40];	char _Title;		// 종목명
	char price[7];	char _price;		// 현재가
	char sign[1];	char _sign;		// 등락부호
	char change[6];	char _change;		// 등락폭
	char chrate[5];	char _chrate;		// 등락률
	char offer[7];	char _offer;		// 매도 호가
	char bid[7];	char _bid;		// 매수 호가
	char volume[9];	char _volume;		// 거래량
	char volrate[6];	char _volrate;		// 거래비율
	char value[9];	char _value;		// 거래대금
	char open[7];	char _open;		// 시가
	char high[7];	char _high;		// 고가
	char low[7];	char _low;		// 저가
	char sale[7];	char _sale;		// 발행가
	char dnlmtprice[7];	char _dnlmtprice;		// 하한가
	char theoryprice[7];	char _theoryprice;		// 이론가
	char grate[7];	char _grate;		// 괴리도
	char actprice[10];	char _actprice;		// 행사가
	char listhprice[7];	char _listhprice;		// 상장후최고가
	char listhdate[4];	char _listhdate;		// 상장후최고가일
	char listlprice[7];	char _listlprice;		// 상장후최저가
	char listldate[4];	char _listldate;		// 상장후최저가일
	char preprice[7];	char _preprice;		// 전일종가
	char hotime[8];	char _hotime;		// 호가시간
	char offho1[7];	char _offho1;		// 매도 최우선호가
	char offho2[7];	char _offho2;		// 매도 차선 호가
	char offho3[7];	char _offho3;		// 매도 차차선 호가
	char offho4[7];	char _offho4;		// 매도 4차선 호가
	char offho5[7];	char _offho5;		// 매도 5차선 호가
	char offho6[7];	char _offho6;		// 매도 6차선 호가
	char offho7[7];	char _offho7;		// 매도 7차선 호가
	char offho8[7];	char _offho8;		// 매도 8차선 호가
	char offho9[7];	char _offho9;		// 매도 9차선 호가
	char offho10[7];	char _offho10;		// 매도 10차선 호가
	char bidho1[7];	char _bidho1;		// 매수 최우선호가
	char bidho2[7];	char _bidho2;		// 매수 차선 호가
	char bidho3[7];	char _bidho3;		// 매수 차차선 호가
	char bidho4[7];	char _bidho4;		// 매수 4차선 호가
	char bidho5[7];	char _bidho5;		// 매수 5차선 호가
	char bidho6[7];	char _bidho6;		// 매수 6차선 호가
	char bidho7[7];	char _bidho7;		// 매수 7차선 호가
	char bidho8[7];	char _bidho8;		// 매수 8차선 호가
	char bidho9[7];	char _bidho9;		// 매수 9차선 호가
	char bidho10[7];	char _bidho10;		// 매수 10차선 호가
	char offremain1[9];	char _offremain1;		// 매도 최우선 잔량
	char offremain2[9];	char _offremain2;		// 매도 차선 잔량
	char offremain3[9];	char _offremain3;		// 매도 차차선 잔량
	char offremain4[9];	char _offremain4;		// 매도 4차선 잔량
	char offremain5[9];	char _offremain5;		// 매도 5차선 잔량
	char offremain6[9];	char _offremain6;		// 매도 6차선 잔량
	char offremain7[9];	char _offremain7;		// 매도 7차선 잔량
	char offremain8[9];	char _offremain8;		// 매도 8차선 잔량
	char offremain9[9];	char _offremain9;		// 매도 9차선 잔량
	char offremain10[9];	char _offremain10;		// 매도 10차선 잔량
	char bidremain1[9];	char _bidremain1;		// 매수 최우선 잔량
	char bidremain2[9];	char _bidremain2;		// 매수 차선 잔량
	char bidremain3[9];	char _bidremain3;		// 매수 차차선 잔량
	char bidremain4[9];	char _bidremain4;		// 매수 4차선 잔량
	char bidremain5[9];	char _bidremain5;		// 매수 5차선 잔량
	char bidremain6[9];	char _bidremain6;		// 매수 6차선 잔량
	char bidremain7[9];	char _bidremain7;		// 매수 7차선 잔량
	char bidremain8[9];	char _bidremain8;		// 매수 8차선 잔량
	char bidremain9[9];	char _bidremain9;		// 매수 9차선 잔량
	char bidremain10[9];	char _bidremain10;		// 매수 10차선 잔량
	char LP_OfferVolume1[9];	char _LP_OfferVolume1;		// LP 매도 최우선 잔량
	char LP_OfferVolume2[9];	char _LP_OfferVolume2;		// LP 매도 차선 잔량
	char LP_OfferVolume3[9];	char _LP_OfferVolume3;		// LP 매도 차차선 잔량
	char LP_OfferVolume4[9];	char _LP_OfferVolume4;		// LP 매도 4차선 잔량
	char LP_OfferVolume5[9];	char _LP_OfferVolume5;		// LP 매도 5차선 잔량
	char LP_OfferVolume6[9];	char _LP_OfferVolume6;		// LP 매도 6차선 잔량
	char LP_OfferVolume7[9];	char _LP_OfferVolume7;		// LP 매도 7차선 잔량
	char LP_OfferVolume8[9];	char _LP_OfferVolume8;		// LP 매도 8차선 잔량
	char LP_OfferVolume9[9];	char _LP_OfferVolume9;		// LP 매도 9차선 잔량
	char LP_OfferVolume10[9];	char _LP_OfferVolume10;		// LP 매도 10차선 잔량
	char LP_BidVolume1[9];	char _LP_BidVolume1;			// LP 매수 최우선 잔량
	char LP_BidVolume2[9];	char _LP_BidVolume2;			// LP 매수 차선 잔량
	char LP_BidVolume3[9];	char _LP_BidVolume3;			// LP 매수 차차선 잔량
	char LP_BidVolume4[9];	char _LP_BidVolume4;			// LP 매수 4차선 잔량
	char LP_BidVolume5[9];	char _LP_BidVolume5;			// LP 매수 5차선 잔량
	char LP_BidVolume6[9];	char _LP_BidVolume6;			// LP 매수 6차선 잔량
	char LP_BidVolume7[9];	char _LP_BidVolume7;			// LP 매수 7차선 잔량
	char LP_BidVolume8[9];	char _LP_BidVolume8;			// LP 매수 8차선 잔량
	char LP_BidVolume9[9];	char _LP_BidVolume9;			// LP 매수 9차선 잔량
	char LP_BidVolume10[9];	char _LP_BidVolume10;			// LP 매수 10차선 잔량
	char offtot[9];	char _offtot;		// 총 매도 잔량
	char bidtot[9];	char _bidtot;		// 총 매수 잔량
	char impv[10];	char _impv;		// 내재변동성
	char delta[9];	char _delta;		// 델타지수
	char gamma[9];	char _gamma;		// 감마지수
	char vega[9];	char _vega;		// 베가변동성
	char theta[9];	char _theta;		// 쎄타시간
	char rho[9];	char _rho;		// 로이자율
	char cdratio[6];	char _cdratio;		// 이자율
	char divideratio[9];	char _divideratio;		// 배당액지수
	char jandatecnt[4];	char _jandatecnt;		// 잔존일
	char elwsdate[8];	char _elwsdate;		// 행사기간개시일
	char elwedate[8];	char _elwedate;		// 행사기간종료일
	char lastdate[8];	char _lastdate;		// 최종거래일
	char balname[18];	char _balname;		// 발행기관
	char listing[9];	char _listing;		// 발행수량
	char rightgb[4];	char _rightgb;		// 권리유형             /*콜,풋,기타*/
	char righttype[6];	char _righttype;		// 권리행사방식         /*유럽형,미국형,기타*/
	char settletype[9];	char _settletype;		// 결제방법             /*현금,실물,현금+실물*/
	char changerate[8];	char _changerate;		// 전환비율
	char rewardrate[5];	char _rewardrate;		// 최소지급율
	char uppartrate[5];	char _uppartrate;		// 가격상승참여율
	char paydate[8];	char _paydate;		// 최종지급일
	char lpjumun[4];	char _lpjumun;		// LP주문가능여부// 불가,가능
	char parity[8];	char _parity;		// 패리티
	char gearingrate[8];	char _gearingrate;		// 기어링비율
	char profitrate[8];	char _profitrate;		// 손익분기율
	char basepoint[8];	char _basepoint;		// 자본지지점
	char lp_name1[6];	char _lp_name1;		// LP회원사1
	char lp_name2[6];	char _lp_name2;		// LP회원사2
	char lp_name3[6];	char _lp_name3;		// LP회원사3
	char lp_name4[6];	char _lp_name4;		// LP회원사4
	char lp_name5[6];	char _lp_name5;		// LP회원사5
	char dongsi[1];	char _dongsi;		// 동시호가구분
	char eqprice[7];	char _eqprice;		// 예상체결가
	char eqsign[1];	char _eqsign;		// 예상체결부호
	char eqchange[6];	char _eqchange;		// 예상체결등락폭
	char eqchrate[5];	char _eqchrate;		// 예상체결등락률
	char eqvol[9];	char _eqvol;		// 예상체결수량
	char lphold[10];	char _lphold;		// LP보유수량
	char lprate[5];	char _lprate;		// LP보유율
	char egearing[8];	char _egearing;		// E기어링
	char fixpay[8];	char _fixpay;		// 확정지급액
	char listdate[8];	char _listdate;		// 상장일
	char listhdatez8[8];	char _listhdatez8;		// 상장후최고가일
	char listldatez8[8];	char _listldatez8;		// 상장후최저가일
	char intval[10];	char _intval;		// 내재가치
	char leverage[8];	char _leverage;		// 레버리지
	char timeval[10];	char _timeval;		// 시간가치
	char gratio[6];	char _gratio;		// 괴리율
	char profitpt[8];	char _profitpt;		// 손익분기점(정수)
	char payproxy[20];	char _payproxy;		// 지급대리인
	char standardopt[2];	char _standardopt;		// 종목구분             /**01:표준,03:조기종료**/
	char koprice[6];	char _koprice;		// 조기종료가
	char koappr[5];	char _koappr;		// KO접근도
	char expcode[12];	char _expcode;		// 확장코드
	char minpayment[8];	char _minpayment;		// 최소지급액
	char stop[1];	char _stop;		// 거래정지구분
	char gratio2[8];	char _gratio2;		// 괴리율2
	char lpstop[8];	char _lpstop;		// LP종료일
	char gonggb[1];	char _gonggb;		// 추가상장공시
	char lp_impv[5];	char _lp_impv;		// LP내재변동성
	char r_intval[8];	char _r_intval;		// 실시간용내재가치
	char jandatecnt2[4];	char _jandatecnt2;		// 잔존일(영업일)
	char profitpt2[10];	char _profitpt2;		// 손익분기점(소수점)
	char alertgb[1];	char _alertgb;		// 투자주의구분
} Ts1701OutBlock;

typedef struct tags1701OutBlock1// 기초자산정보,[반복]
{
	char code1[6];	char _code1;		// 기초자산코드1
	char hname1[20];	char _hname1;		// 기초자산명1
	char price1[7];	char _price1;		// 현재가1
	char sign1[1];	char _sign1;		// 등락부호1
	char change1[6];	char _change1;		// 등락폭1
	char chrate1[5];	char _chrate1;		// 등락률1
	char comrate1[5];	char _comrate1;		// 구성비1
	char pastv1[5];	char _pastv1;		// 과거변동성1
	char basegubun[1];	char _basegubun;		// 기초자산시장구분     /*1:코스피,2:코스닥*/
} Ts1701OutBlock1;

typedef struct tags1701OutBlock2// 거래원정보
{
	char tratimez5[5];	char _tratimez5;		// 시간
	char off_tra1[6];	char _off_tra1;		// 매도 거래원1
	char bid_tra1[6];	char _bid_tra1;		// 매수 거래원1
	char offvolume1[9];	char _offvolume1;		// 매도 거래량1
	char bidvolume1[9];	char _bidvolume1;		// 매수 거래량1
	char off_tra2[6];	char _off_tra2;		// 매도 거래원2
	char bid_tra2[6];	char _bid_tra2;		// 매수 거래원2
	char offvolume2[9];	char _offvolume2;		// 매도 거래량2
	char bidvolume2[9];	char _bidvolume2;		// 매수 거래량2
	char off_tra3[6];	char _off_tra3;		// 매도 거래원3
	char bid_tra3[6];	char _bid_tra3;		// 매수 거래원3
	char offvolume3[9];	char _offvolume3;		// 매도 거래량3
	char bidvolume3[9];	char _bidvolume3;		// 매수 거래량3
	char off_tra4[6];	char _off_tra4;		// 매도 거래원4
	char bid_tra4[6];	char _bid_tra4;		// 매수 거래원4
	char offvolume4[9];	char _offvolume4;		// 매도 거래량4
	char bidvolume4[9];	char _bidvolume4;		// 매수 거래량4
	char off_tra5[6];	char _off_tra5;		// 매도 거래원5
	char bid_tra5[6];	char _bid_tra5;		// 매수 거래원5
	char offvolume5[9];	char _offvolume5;		// 매도 거래량5
	char bidvolume5[9];	char _bidvolume5;		// 매수 거래량5
	char offvolall[9];	char _offvolall;		// 매도 외국인거래량
	char bidvolall[9];	char _bidvolall;		// 매수 외국인거래량
	char alloffvol[9];	char _alloffvol;		// 전체거래원 매도 합
	char allbidvol[9];	char _allbidvol;		// 전체거래원 매수 합
} Ts1701OutBlock2;

typedef struct tags1701OutBlock3// ELW변동거래량자료,[반복]
{
	char chetime[8];	char _chetime;		// 시간
	char price[7];	char _price;		// 현재가
	char sign[1];	char _sign;		// 등락부호
	char change[6];	char _change;		// 등락폭
	char offer[7];	char _offer;		// 매도 호가
	char bid[7];	char _bid;		// 매수 호가
	char volume[9];	char _volume;		// 거래량
	char movvol[9];	char _movvol;		// 변동량
} Ts1701OutBlock3;

typedef struct tags1701OutBlock4// K200기초자산정보
{
	char code6[2];	char _code6;		// 기초자산코드6
	char hname6[20];	char _hname6;		// 기초자산명6
	char price6[7];	char _price6;		// 현재가6
	char sign6[1];	char _sign6;		// 등락부호6
	char change6[6];	char _change6;		// 등락폭6
	char chrate6[5];	char _chrate6;		// 등락률6
	char comrate6[5];	char _comrate6;		// 구성비6
	char pastv6[5];	char _pastv6;		// 과거변동성6
} Ts1701OutBlock4;

typedef struct tags1701OutBlock5// 해외지수기초자산정보
{
	char code7[6];	char _code7;		// 기초자산코드7
	char hname7[16];	char _hname7;		// 기초자산명7
	char price7[9];	char _price7;		// 현재가7
	char sign7[1];	char _sign7;		// 등락부호7
	char change7[9];	char _change7;		// 등락폭7
	char chrate7[5];	char _chrate7;		// 등락률7
	char time7[4];	char _time7;		// 데이타시간
} Ts1701OutBlock5;

typedef struct tags1701
{
	Ts1701InBlock                     s1701inblock                          ;		// 기본입력
	Ts1701OutBlock                    s1701outblock                         ;		// 종목마스타기본자료
	Ts1701OutBlock1                   s1701outblock1[20];		// 기초자산정보 ,[반복]
	Ts1701OutBlock2                   s1701outblock2                        ;		// 거래원정보
	Ts1701OutBlock3                   s1701outblock3[20];		// ELW변동거래량자료 ,[반복]
	Ts1701OutBlock4                   s1701outblock4                        ;		// K200기초자산정보
	Ts1701OutBlock5                   s1701outblock5                        ;		// 해외지수기초자산정보
} Ts1701;

typedef struct tagp1003InBlock// 입력Data
{
	char Lang[1];	char _Lang;		// 한영구분
	char gubun[1];	char _gubun;		// 선옵구분             /*f:KRX선물,o:KRX옵션,u:내부선물,p:내부옵션*/
} Tp1003InBlock;

typedef struct tagp1003OutBlock// 코드출력Data,[반복]
{
	char codez8[8];	char _codez8;		// code
	char namez30[30];	char _namez30;		// name
} Tp1003OutBlock;

typedef struct tagp1003
{
	Tp1003InBlock                     p1003inblock                          ;		// 입력Data
	Tp1003OutBlock                    p1003outblock[20];		// 코드출력Data ,[반복]
} Tp1003;

//----------------------------------------------------------------------//
// ETF 현재가 조회 (c1151)
//----------------------------------------------------------------------//

typedef struct tagc1151InBlock// 기본입력
{
	char Lang[1];	char _Lang;							// 한영구분. 기본값 'K'
	char Code[6];	char _Code;							// 종목코드
} Tc1151InBlock;

typedef struct tagc1151OutBlock// 종목마스타기본자료
{
	char Code[6];	char _Code;							// 종목코드
	char Title[13];	char _Title;						// 종목명
	char MarketPrice[7];	char _MarketPrice;			// 현재가
	char DiffSign[1];	char _DiffSign;					// 등락부호
	char Diff[6];	char _Diff;							// 등락폭
	char DiffRate[5];	char _DiffRate;					// 등락률
	char OfferPrice[7];	char _OfferPrice;				// 매도 호가
	char BidPrice[7];	char _BidPrice;					// 매수 호가
	char Volume[9];	char _Volume;						// 거래량
	char TrVolRate[6];	char _TrVolRate;				// 거래비율
	char FloatVolRate[5];	char _FloatVolRate;			// 유동주회전율
	char TrAmount[9];	char _TrAmount;					// 거래대금
	char UpLmtPrice[7];	char _UpLmtPrice;				// 상한가
	char High[7];	char _High;							// 장중고가
	char Open[7];	char _Open;							// 시가
	char VsOpenSign[1];	char _VsOpenSign;				// 시가대비부호
	char VsOpenDiff[6];	char _VsOpenDiff;				// 시가대비등락폭
	char Low[7];	char _Low;							// 장중저가
	char LowLmtPrice[7];	char _LowLmtPrice;			// 하한가
	char QuoteTime[8];	char _QuoteTime;				// 호가시간
	char OfferPrice1[7];	char _OfferPrice1;			// 매도 최우선 호가
	char OfferPrice2[7];	char _OfferPrice2;			// 매도 차선 호가
	char OfferPrice3[7];	char _OfferPrice3;			// 매도 차차선 호가
	char OfferPrice4[7];	char _OfferPrice4;			// 매도 4차선 호가
	char OfferPrice5[7];	char _OfferPrice5;			// 매도 5차선 호가
	char OfferPrice6[7];	char _OfferPrice6;			// 매도 6차선 호가
	char OfferPrice7[7];	char _OfferPrice7;			// 매도 7차선 호가
	char OfferPrice8[7];	char _OfferPrice8;			// 매도 8차선 호가
	char OfferPrice9[7];	char _OfferPrice9;			// 매도 9차선 호가
	char OfferPrice10[7];	char _OfferPrice10;			// 매도 10차선 호가
	char BidPrice1[7];	char _BidPrice1;				// 매수 최우선 호가
	char BidPrice2[7];	char _BidPrice2;				// 매수 차선 호가
	char BidPrice3[7];	char _BidPrice3;				// 매수 차차선 호가
	char BidPrice4[7];	char _BidPrice4;				// 매수 4차선 호가
	char BidPrice5[7];	char _BidPrice5;				// 매수 5차선 호가
	char BidPrice6[7];	char _BidPrice6;				// 매수 6차선 호가
	char BidPrice7[7];	char _BidPrice7;				// 매수 7차선 호가
	char BidPrice8[7];	char _BidPrice8;				// 매수 8차선 호가
	char BidPrice9[7];	char _BidPrice9;				// 매수 9차선 호가
	char BidPrice10[7];	char _BidPrice10;				// 매수 10차선 호가
	char OfferVolume1[9];	char _OfferVolume1;			// 매도 최우선 잔량
	char OfferVolume2[9];	char _OfferVolume2;			// 매도 차선 잔량
	char OfferVolume3[9];	char _OfferVolume3;			// 매도 차차선 잔량
	char OfferVolume4[9];	char _OfferVolume4;			// 매도 4차선 잔량
	char OfferVolume5[9];	char _OfferVolume5;			// 매도 5차선 잔량
	char OfferVolume6[9];	char _OfferVolume6;			// 매도 6차선 잔량
	char OfferVolume7[9];	char _OfferVolume7;			// 매도 7차선 잔량
	char OfferVolume8[9];	char _OfferVolume8;			// 매도 8차선 잔량
	char OfferVolume9[9];	char _OfferVolume9;			// 매도 9차선 잔량
	char OfferVolume10[9];	char _OfferVolume10;		// 매도 10차선 잔량
	char BidVolume1[9];	char _BidVolume1;				// 매수 최우선 잔량
	char BidVolume2[9];	char _BidVolume2;				// 매수 차선 잔량
	char BidVolume3[9];	char _BidVolume3;				// 매수 차차선 잔량
	char BidVolume4[9];	char _BidVolume4;				// 매수 4차선 잔량
	char BidVolume5[9];	char _BidVolume5;				// 매수 5차선 잔량
	char BidVolume6[9];	char _BidVolume6;				// 매수 6차선 잔량
	char BidVolume7[9];	char _BidVolume7;				// 매수 7차선 잔량
	char BidVolume8[9];	char _BidVolume8;				// 매수 8차선 잔량
	char BidVolume9[9];	char _BidVolume9;				// 매수 9차선 잔량
	char BidVolume10[9];	char _BidVolume10;			// 매수 10차선 잔량
	char OfferVolTot[9];	char _OfferVolTot;			// 총 매도 잔량
	char BidVolTot[9];	char _BidVolTot;				// 총 매수 잔량
	char OfferVolAfterHour[9];	char _OfferVolAfterHour; // 시간외 매도 잔량
	char BidVolAfterHour[9];	char _BidVolAfterHour;	// 시간외 매수 잔량
	char PivotUp2[7];	char _PivotUp2;					// 피봇 2차 저항
	char PivotUp1[7];	char _PivotUp1;					// 피봇 1차 저항
	char PivotPrice[7];	char _PivotPrice;				// 피봇가
	char PivotDown1[7];	char _PivotDown1;				// 피봇 1차 지지
	char PivotDown2[7];	char _PivotDown2;				// 피봇 2차 지지
	char Market[6];	char _Market;						// 코스피/코스닥 구분
	char Sector[18];	char _Sector;					// 업종명
	char CapSize[6];	char _CapSize;					// 자본금규모
	char SettleMonth[16];	char _SettleMonth;			// 결산월
	char MarketAction1[16];	char _MarketAction1;		// 시장조치1
	char MarketAction2[16];	char _MarketAction2;		// 시장조치2
	char MarketAction3[16];	char _MarketAction3;		// 시장조치3
	char MarketAction4[16];	char _MarketAction4;		// 시장조치4
	char MarketAction5[16];	char _MarketAction5;		// 시장조치5
	char MarketAction6[16];	char _MarketAction6;		// 시장조치6
	char ConvertBond[6];	char _ConvertBond;			// CB구분
	char NominalPrice[7];	char _NominalPrice;			// 액면가
	char PrevPriceTitle[12];	char _PrevPriceTitle;	// 전일 종가 타이틀
	char PrevPrice[7];	char _PrevPrice;				// 전일종가
	char MortgageValue[7];	char _MortgageValue;		// 대용가
	char PublicOfferPrice[7];	char _PublicOfferPrice;	// 공모가
	char High5Day[7];	char _High5Day;					// 5일고가
	char Low5Day[7];	char _Low5Day;					// 5일저가
	char High20Day[7];	char _High20Day;				// 20일고가
	char Low20Day[7];	char _Low20Day;					// 20일저가
	char High1Year[7];	char _High1Year;				// 52주최고가
	char High1YearDate[4];	char _High1YearDate;		// 52주최고가일
	char Low1Year[7];	char _Low1Year;					// 52주최저가
	char Low1YearDate[4];	char _Low1YearDate;			// 52주최저가일
	char FloatVolume[8];	char _FloatVolume;			// 유동주식수
	char ListVolBy1000[12];	char _ListVolBy1000;		// 상장주식수_천주
	char MarketCapital[9];	char _MarketCapital;		// 시가총액
	char TraderInfoTime[5];	char _TraderInfoTime;		// 거래원 정보 최종 수신 시간
	char Seller1[6];	char _Seller1;					// 매도 거래원1
	char Buyer1[6];	char _Buyer1;						// 매수 거래원1
	char Seller1Volumn[9];	char _Seller1Volume;		// 매도 거래량1
	char Buyer1Volume[9];	char _Buyer1Volume;			// 매수 거래량1
	char Seller2[6];	char _Seller2;					// 매도 거래원2
	char Buyer2[6];	char _Buyer2;						// 매수 거래원2
	char Seller2Volumn[9];	char _Seller2Volume;		// 매도 거래량2
	char Buyer2Volume[9];	char _Buyer2Volume;			// 매수 거래량2
	char Seller3[6];	char _Seller3;					// 매도 거래원3
	char Buyer3[6];	char _Buyer3;						// 매수 거래원3
	char Seller3Volumn[9];	char _Seller3Volume;		// 매도 거래량3
	char Buyer3Volume[9];	char _Buyer3Volume;			// 매수 거래량3
	char Seller4[6];	char _Seller4;					// 매도 거래원4
	char Buyer4[6];	char _Buyer4;						// 매수 거래원4
	char Seller4Volumn[9];	char _Seller4Volume;		// 매도 거래량4
	char Buyer4Volume[9];	char _Buyer4Volume;			// 매수 거래량4
	char Seller5[6];	char _Seller5;					// 매도 거래원5
	char Buyer5[6];	char _Buyer5;						// 매수 거래원5
	char Seller5Volumn[9];	char _Seller5Volume;		// 매도 거래량5
	char Buyer5Volume[9];	char _Buyer5Volume;			// 매수 거래량5
	char ForeignSellVolume[9];	char _ForeignSellVolume; // 외국인 매도 거래량
	char ForeignBuyVolume[9];	char _ForeignBuyVolume;	// 외국인 매수 거래량
	char ForeignTime[6];	char _ForeignTime;			// 외국인 시간 ???
	char ForeignHoldingRate[5];	char _ForeignHoldingRate; // 외국인 지분율
	char SettleDate[4];	char _SettleDate;				// 결제일
	char DebtPercent[5];	char _DebtPercent;			// 잔고비율(%)
	char RightsIssueDate[4];	char _RightsIssueDate;	// 유상기준일
	char BonusIssueDate[4];	char _BonusIssueDate;		// 무상기준일
	char RightsIssueRate[5];	char _RightsIssueRate;	// 유상배정비율
	char BonusIssueRate[5];	char _BonusIssueRate;		// 무상배정비율
	char IPO_Date[8];	char _IPO_Date;					// 상장일
	char ListedVolume[12];	char _ListedVolume;			// 상장주식수_주
	char SellTotalSum[9];	char _SellTotalSum;			// 전체 거래원 매도 합계
	char BuyTotalSum[9];	char _BuyTotalSum;			// 전체 거래원 매수 합계
} Tc1151OutBlock;

typedef struct tagc1151OutBlock2// 변동거래량자료
{
	char Time[8];	char _Time;							// 시간
	char MarketPrice[7];	char _MarketPrice;			// 현재가
	char DiffSign[1];	char _DiffSign;					// 등락부호
	char Diff[6];	char _Diff;							// 등락폭
	char OfferPrice[7];	char _OfferPrice;				// 매도 호가
	char BidPrice[7];	char _BidPrice;					// 매수 호가
	char DiffVolume[8];	char _DiffVolumn;				// 변동거래량
	char Volume[9];	char _Volume;						// 거래량
} Tc1151OutBlock2;

typedef struct tagc1151OutBlock3// 예상체결
{
	char SyncOfferBid[1];	char _SyncOfferBid;			// 동시 호가 구분
	char EstmPrice[7];	char _EstmPrice;		// 예상 체결가
	char EstmSign[1];	char _EstmSign;			// 예상 체결 부호
	char EstmDiff[6];	char _EstmDiff;			// 예상 체결 등락폭
	char EstmDiffRate[5];	char _EstmDiffRate;	// 예상 체결 등락률
	char EstmVolume[9];	char _EstmVolume;		// 예상체결 수량
} Tc1151OutBlock3;

typedef struct tagc1151OutBlock4// ETF자료
{
	char ETF[1];	char _ETF;							// ETF 구분
	char NAV[9];	char _NAV;							// 장중/최종 NAV
	char DiffSign[1];	char _DiffSign;					// NAV 등락 부호
	char Diff[9];	char _Diff;							// NAV 등락폭
	char PrevNAV[9];	char _PrevNAV;					// 전일 NAV
	char DivergeRate[9];	char _DivergeRate;			// 괴리율
	char DivergeSign[1];	char _DivergeSign;			// 괴리율 부호
	char DividendPerCU[18];	char _DividendPerCU;		// CU(Creation Unit : 설정단위)당 현금 배당액(원)
	char ConstituentNo[4];	char _ConstituentNo;		// 구성 종목수
	char NAVBy100Million[7];	char _NAVBy100Million;	// 순자산총액(억원)
	char TrackingErrRate[9];	char _TrackingErrRate;	// 추적오차율
	char LP_OfferVolume1[9];	char _LP_OfferVolume1;	// LP 매도 최우선 잔량
	char LP_OfferVolume2[9];	char _LP_OfferVolume2;	// LP 매도 차선 잔량
	char LP_OfferVolume3[9];	char _LP_OfferVolume3;	// LP 매도 차차선 잔량
	char LP_OfferVolume4[9];	char _LP_OfferVolume4;	// LP 매도 4차선 잔량
	char LP_OfferVolume5[9];	char _LP_OfferVolume5;	// LP 매도 5차선 잔량
	char LP_OfferVolume6[9];	char _LP_OfferVolume6;	// LP 매도 6차선 잔량
	char LP_OfferVolume7[9];	char _LP_OfferVolume7;	// LP 매도 7차선 잔량
	char LP_OfferVolume8[9];	char _LP_OfferVolume8;	// LP 매도 8차선 잔량
	char LP_OfferVolume9[9];	char _LP_OfferVolume9;	// LP 매도 9차선 잔량
	char LP_OfferVolume10[9];	char _LP_OfferVolume10;	// LP 매도 10차선 잔량
	char LP_BidVolume1[9];	char _LP_BidVolume1;		// LP 매수 최우선 잔량
	char LP_BidVolume2[9];	char _LP_BidVolume2;		// LP 매수 차선 잔량
	char LP_BidVolume3[9];	char _LP_BidVolume3;		// LP 매수 차차선 잔량
	char LP_BidVolume4[9];	char _LP_BidVolume4;		// LP 매수 4차선 잔량
	char LP_BidVolume5[9];	char _LP_BidVolume5;		// LP 매수 5차선 잔량
	char LP_BidVolume6[9];	char _LP_BidVolume6;		// LP 매수 6차선 잔량
	char LP_BidVolume7[9];	char _LP_BidVolume7;		// LP 매수 7차선 잔량
	char LP_BidVolume8[9];	char _LP_BidVolume8;		// LP 매수 8차선 잔량
	char LP_BidVolume9[9];	char _LP_BidVolume9;		// LP 매수 9차선 잔량
	char LP_BidVolume10[9];	char _LP_BidVolume10;		// LP 매수 10차선 잔량
	char TrackingMethod[8];	char _TrackingMethod;		// ETF 복제 방법 구분 코드
	char ETF_Type[6];	char _ETF_Type;					// ETF 상품 유형 코드
} Tc1151OutBlock4;

typedef struct tagc1151OutBlock5// 베이스지수자료
{
	char IndexCode[2];	char _IndexCode;				// 지수코드
	char SectorCode[4];	char _SectorCode;				// 섹터코드
	char IndexName[20];	char _IndexName;				// 지수명
	char KP200Index[8];	char _KP200Index;				// 지수
	char KP200Sign[1];	char _KP200Sign;				// 등락부호
	char KP200Diff[8];	char _KP200Diff;				// 등락폭
	char BondIndex[10];	char _BondIndex;				// 채권지수
	char BondSign[1];	char _BondSign;					// 채권등락부호
	char BondDiff[10];	char _BondDiff;					// 채권등락폭
	char ForeignIndexSymbol[12];	char _ForeignIndexSymbol; // 해외지수심볼
	char EtcSectorCode[3];	char _EtcSectorCode;		// 기타업종코드
	char BondIndexCode[6];	char _BondIndexCode;		// 채권지수코드
	char BondDetailCode[1];	char _BondDetailCode;		// 채권지수세부코드
} Tc1151OutBlock5;

typedef struct tagc1151
{
	Tc1151InBlock c1151inblock;							// 기본입력
	Tc1151OutBlock c1151outblock;						// 종목마스타기본자료
	Tc1151OutBlock2 c1151outblock2;						// 변동거래량자료
	Tc1151OutBlock3 c1151outblock3;						// 예상체결
	Tc1151OutBlock4 c1151outblock4;						// ETF자료
	Tc1151OutBlock5 c1151outblock5;						// 베이스지수자료
} Tc1151;


typedef struct tagh1InBlock// 입력
{
	char Code[6];		// 종목코드
} Th1InBlock;

typedef struct tagh1OutBlock// 출력
{
	char Code[6];		// 종목코드
	char hotime[8];		// 시간
	char offer[7];		// 매도 호가
	char bid[7];		// 매수 호가
	char offerrem[9];		// 매도 호가잔량
	char bidrem[9];		// 매수 호가잔량
	char P_offer[7];		// 차 매도 호가
	char P_bid[7];		// 차 매수 호가
	char P_offerrem[9];		// 차 매도 호가잔량
	char P_bidrem[9];		// 차 매수 호가잔량
	char S_offer[7];		// 차차 매도 호가
	char S_bid[7];		// 차차 매수 호가
	char S_offerrem[9];		// 차차 매도 호가잔량
	char S_bidrem[9];		// 차차 매수 호가잔량
	char S4_offer[7];		// 4차 매도 호가
	char S4_bid[7];		// 4차 매수 호가
	char S4_offerrem[9];		// 4차 매도 호가잔량
	char S4_bidrem[9];		// 4차 매수 호가잔량
	char S5_offer[7];		// 5차 매도 호가
	char S5_bid[7];		// 5차 매수 호가
	char S5_offerrem[9];		// 5차 매도 호가잔량
	char S5_bidrem[9];		// 5차 매수 호가잔량
	char T_offerrem[9];		// 총 매도 호가잔량
	char T_bidrem[9];		// 총 매수 호가잔량
	char S6_offer[7];		// 6차 매도 호가
	char S6_bid[7];		// 6차 매수 호가
	char S6_offerrem[9];		// 6차 매도 호가잔량
	char S6_bidrem[9];		// 6차 매수 호가잔량
	char S7_offer[7];		// 7차 매도 호가
	char S7_bid[7];		// 7차 매수 호가
	char S7_offerrem[9];		// 7차 매도 호가잔량
	char S7_bidrem[9];		// 7차 매수 호가잔량
	char S8_offer[7];		// 8차 매도 호가
	char S8_bid[7];		// 8차 매수 호가
	char S8_offerrem[9];		// 8차 매도 호가잔량
	char S8_bidrem[9];		// 8차 매수 호가잔량
	char S9_offer[7];		// 9차 매도 호가
	char S9_bid[7];		// 9차 매수 호가
	char S9_offerrem[9];		// 9차 매도 호가잔량
	char S9_bidrem[9];		// 9차 매수 호가잔량
	char S10_offer[7];		// 10차 매도 호가
	char S10_bid[7];		// 10차 매수 호가
	char S10_offerrem[9];		// 10차 매도 호가잔량
	char S10_bidrem[9];		// 10차 매수 호가잔량
	char volume[9];		// 누적거래량
} Th1OutBlock;

typedef struct tagh1
{
	Th1InBlock                        h1inblock                             ;		// 입력
	Th1OutBlock                       h1outblock                            ;		// 출력
} Th1;

typedef struct tagk3InBlock// 입력
{
	char Code[6];		// 종목코드
} Tk3InBlock;

typedef struct tagk3OutBlock// 출력
{
	char Code[6];		// 종목코드
	char time[8];		// 시간
	char offer[7];		// 매도 호가
	char bid[7];		// 매수 호가
	char offerrem[9];		// 매도 호가잔량
	char bidrem[9];		// 매수 호가잔량
	char P_offer[7];		// 차 매도 호가
	char P_bid[7];		// 차 매수 호가
	char P_offerrem[9];		// 차 매도 호가잔량
	char P_bidrem[9];		// 차 매수 호가잔량
	char S_offer[7];		// 차차 매도 호가
	char S_bid[7];		// 차차 매수 호가
	char S_offerrem[9];		// 차차 매도 호가잔량
	char S_bidrem[9];		// 차차 매수 호가잔량
	char S4_offer[7];		// 4차 매도 호가
	char S4_bid[7];		// 4차 매수 호가
	char S4_offerrem[9];		// 4차 매도 호가잔량
	char S4_bidrem[9];		// 4차 매수 호가잔량
	char S5_offer[7];		// 5차 매도 호가
	char S5_bid[7];		// 5차 매수 호가
	char S5_offerrem[9];		// 5차 매도 호가잔량
	char S5_bidrem[9];		// 5차 매수 호가잔량
	char T_offerrem[9];		// 총 매도 호가잔량
	char T_bidrem[9];		// 총 매수 호가잔량
	char S6_offer[7];		// 6차 매도 호가
	char S6_bid[7];		// 6차 매수 호가
	char S6_offerrem[9];		// 6차 매도 호가잔량
	char S6_bidrem[9];		// 6차 매수 호가잔량
	char S7_offer[7];		// 7차 매도 호가
	char S7_bid[7];		// 7차 매수 호가
	char S7_offerrem[9];		// 7차 매도 호가잔량
	char S7_bidrem[9];		// 7차 매수 호가잔량
	char S8_offer[7];		// 8차 매도 호가
	char S8_bid[7];		// 8차 매수 호가
	char S8_offerrem[9];		// 8차 매도 호가잔량
	char S8_bidrem[9];		// 8차 매수 호가잔량
	char S9_offer[7];		// 9차 매도 호가
	char S9_bid[7];		// 9차 매수 호가
	char S9_offerrem[9];		// 9차 매도 호가잔량
	char S9_bidrem[9];		// 9차 매수 호가잔량
	char S10_offer[7];		// 10차 매도 호가
	char S10_bid[7];		// 10차 매수 호가
	char S10_offerrem[9];		// 10차 매도 호가잔량
	char S10_bidrem[9];		// 10차 매수 호가잔량
	char volume[9];		// 거래량
} Tk3OutBlock;

typedef struct tagk3
{
	Tk3InBlock                        k3inblock                             ;		// 입력
	Tk3OutBlock                       k3outblock                            ;		// 출력
} Tk3;

typedef struct tagh2InBlock// 입력
{
	char Code[6];		// 종목코드
} Th2InBlock;

typedef struct tagh2OutBlock// 출력
{
	char Code[6];		// 종목코드
	char hotime[8];		// 시간
	char O_offerrem[9];		// 총 매도 호가잔량
	char O_bidrem[9];		// 총 매수 호가잔량
} Th2OutBlock;

typedef struct tagh2
{
	Th2InBlock                        h2inblock                             ;		// 입력
	Th2OutBlock                       h2outblock                            ;		// 출력
} Th2;

typedef struct tagk4InBlock// 입력
{
	char Code[6];		// 종목코드
} Tk4InBlock;

typedef struct tagk4OutBlock// 출력
{
	char Code[6];		// 종목코드
	char hotime[8];		// 시간
	char O_offerrem[9];		// 총 매도 호가잔량
	char O_bidrem[9];		// 총 매수 호가잔량
} Tk4OutBlock;

typedef struct tagk4
{
	Tk4InBlock                        k4inblock                             ;		// 입력
	Tk4OutBlock                       k4outblock                            ;		// 출력
} Tk4;

typedef struct tagh3InBlock// 입력
{
	char Code[6];		// 종목코드
} Th3InBlock;

typedef struct tagh3OutBlock// 출력
{
	char Code[6];		// 종목코드
	char hotime[8];		// 시간
	char dongsi[1];		// 동시구분
	char jeqprice[7];		// 예상체결가
	char jeqsign[1];		// 예상등락부호
	char jeqchange[6];		// 예상등락폭
	char jeqchrate[5];		// 예상등락률
	char jeqvol[9];		// 예상체결수량
	char offer[7];		// 매도 호가
	char bid[7];		// 매수 호가
	char offerrem[9];		// 매도 호가잔량
	char bidrem[9];		// 매수 호가잔량
} Th3OutBlock;

typedef struct tagh3
{
	Th3InBlock                        h3inblock                             ;		// 입력
	Th3OutBlock                       h3outblock                            ;		// 출력
} Th3;

typedef struct tagk5InBlock// 입력
{
	char Code[6];		// 종목코드
} Tk5InBlock;

typedef struct tagk5OutBlock// 출력
{
	char Code[6];		// 종목코드
	char hotime[8];		// 시간
	char dongsi[1];		// 동시구분
	char jeqprice[7];		// 예상체결가
	char jeqsign[1];		// 예상등락부호
	char jeqchange[6];		// 예상등락폭
	char jeqchrate[5];		// 예상등락률
	char jeqvol[9];		// 예상체결수량
	char offer[7];		// 매도 호가
	char bid[7];		// 매수 호가
	char offerrem[9];		// 매도 호가잔량
	char bidrem[9];		// 매수 호가잔량
} Tk5OutBlock;

typedef struct tagk5
{
	Tk5InBlock                        k5inblock                             ;		// 입력
	Tk5OutBlock                       k5outblock                            ;		// 출력
} Tk5;

typedef struct tagj8InBlock// 입력
{
	char Code[6];	char _Code;		// 종목코드
} Tj8InBlock;

typedef struct tagj8OutBlock// 출력
{
	char Code[6];	char _Code;		// 종목코드
	char time[8];	char _time;		// 시간
	char sign[1];	char _sign;		// 등락부호
	char change[6];	char _change;		// 등락폭
	char price[7];	char _price;		// 현재가
	char chrate[5];	char _chrate;		// 등락률
	char high[7];	char _high;		// 고가
	char low[7];	char _low;		// 저가
	char offer[7];	char _offer;		// 매도 호가
	char bid[7];	char _bid;		// 매수 호가
	char volume[9];	char _volume;		// 거래량
	char volrate[6];	char _volrate;		// 거래량전일비
	char movolume[8];	char _movolume;		// 변동거래량
	char value[9];	char _value;		// 거래대금
	char open[7];	char _open;		// 시가
	char avgprice[7];	char _avgprice;		// 가중평균가
	char janggubun[1];	char _janggubun;		// 장구분
} Tj8OutBlock;

typedef struct tagj8
{
	Tj8InBlock                        j8inblock                             ;		// 입력
	Tj8OutBlock                       j8outblock                            ;		// 출력
} Tj8;

typedef struct tagk8InBlock// 입력
{
	char Code[6];	char _Code;		// 종목코드
} Tk8InBlock;

typedef struct tagk8OutBlock// 출력
{
	char Code[6];	char _Code;		// 종목코드
	char time[8];	char _time;		// 시간
	char price[7];	char _price;		// 현재가
	char sign[1];	char _sign;		// 등락부호
	char change[6];	char _change;		// 등락폭
	char chrate[5];	char _chrate;		// 등락률
	char high[7];	char _high;		// 고가
	char low[7];	char _low;		// 저가
	char offer[7];	char _offer;		// 매도 호가
	char bid[7];	char _bid;		// 매수 호가
	char volume[9];	char _volume;		// 거래량
	char volrate[6];	char _volrate;		// 거래량전일비
	char movolume[8];	char _movolume;		// 변동거래량
	char value[9];	char _value;		// 거래대금
	char open[7];	char _open;		// 시가
	char avgprice[7];	char _avgprice;		// 가중평균가
	char janggubun[1];	char _janggubun;		// 장구분
} Tk8OutBlock;

typedef struct tagk8
{
	Tk8InBlock                        k8inblock                             ;		// 입력
	Tk8OutBlock                       k8outblock                            ;		// 출력
} Tk8;

typedef struct tagf1InBlock// 입력
{
	char fuitem[4];		// 종목코드
} Tf1InBlock;

typedef struct tagf1OutBlock// 출력
{
	char fuitem[4];		// 종목코드
	char fuhotime[8];		// 시간
	char fuoffer[5];		// 매도 우선호가
	char fubid[5];		// 매수 우선호가
	char fuofferjan[6];		// 매도 우선잔량
	char fubidjan[6];		// 매수 우선잔량
	char fujoffer[5];		// 차선 매도 호가
	char fujbid[5];		// 차선 매수 호가
	char fujofferjan[6];		// 차선 매도 잔량
	char fujbidjan[6];		// 차선 매수 잔량
	char fujjoffer[5];		// 차차선 매도 호가
	char fujjbid[5];		// 차차선 매수 호가
	char fujjofferjan[6];		// 차차선 매도 잔량
	char fujjbidjan[6];		// 차차선 매수 잔량
	char futofferjan[6];		// 총 매도 잔량
	char futbidjan[6];		// 총 매수 잔량
	char fuj4offer[5];		// 4차선 매도 호가
	char fuj4bid[5];		// 4차선 매수 호가
	char fuj4offerjan[6];		// 4차선 매도 잔량
	char fuj4bidjan[6];		// 4차선 매수 잔량
	char fuj5offer[5];		// 5차선 매도 호가
	char fuj5bid[5];		// 5차선 매수 호가
	char fuj5offerjan[6];		// 5차선 매도 잔량
	char fuj5bidjan[6];		// 5차선 매수 잔량
	char fuoffersu[4];		// 우선 매도 건수
	char fujoffersu[4];		// 차선 매도 건수
	char fujjoffersu[4];		// 차차선 매도 건수
	char fuj4offersu[4];		// 4차선 매도 건수
	char fuj5offersu[4];		// 5차선 매도 건수
	char futoffersu[5];		// 총 매도 건수
	char fubidsu[4];		// 우선 매수 건수
	char fujbidsu[4];		// 차선 매수 건수
	char fujjbidsu[4];		// 차차선 매수 건수
	char fuj4bidsu[4];		// 4차선 매수 건수
	char fuj5bidsu[4];		// 5차선 매수 건수
	char futbidsu[5];		// 총 매수 건수
} Tf1OutBlock;

typedef struct tagf1
{
	Tf1InBlock                        f1inblock                             ;		// 입력
	Tf1OutBlock                       f1outblock                            ;		// 출력
} Tf1;

typedef struct tagf3InBlock// 입력
{
	char fuitem[4];		// 종목코드
} Tf3InBlock;

typedef struct tagf3OutBlock// 출력
{
	char fuitem[4];		// 종목코드
	char futheoryprice[5];		// 선물이론가
	char futheorytime[8];		// 이론가시간
	char fugrate[5];		// 괴리도
	char fugratio[5];		// 괴리율
} Tf3OutBlock;

typedef struct tagf3
{
	Tf3InBlock                        f3inblock                             ;		// 입력
	Tf3OutBlock                       f3outblock                            ;		// 출력
} Tf3;

typedef struct tagf4InBlock// 입력
{
	char fuitem[4];		// 종목코드
} Tf4InBlock;

typedef struct tagf4OutBlock// 출력
{
	char fuitem[4];		// 종목코드
	char fuchetime[8];		// 체결시간
	char fuopenyak[7];		// 미결제약정수량
	char fupreopenyak[7];		// 전일미결제약정수량
} Tf4OutBlock;

typedef struct tagf4
{
	Tf4InBlock                        f4inblock                             ;		// 입력
	Tf4OutBlock                       f4outblock                            ;		// 출력
} Tf4;

typedef struct tagf8InBlock// 입력
{
	char fuitem[4];	char _fuitem;		// 종목코드
} Tf8InBlock;

typedef struct tagf8OutBlock// 출력
{
	char fuitem[4];	char _fuitem;		// 종목코드
	char fuchetime[8];	char _fuchetime;		// 시간
	char fusign[1];	char _fusign;		// 등락부호
	char fuchange[5];	char _fuchange;		// 등락폭
	char fucurr[5];	char _fucurr;		// 현재가
	char fuhigh[5];	char _fuhigh;		// 고가
	char fulow[5];	char _fulow;		// 저가
	char fuvol[6];	char _fuvol;		// 체결수량
	char fuvolall[7];	char _fuvolall;		// 누적체결수량
	char fuvalall[12];	char _fuvalall;		// 누적거래대금
	char fuopenyak[7];	char _fuopenyak;		// 미결제약정수량
	char fuoffer[5];	char _fuoffer;		// 우선 매도 호가
	char fubid[5];	char _fubid;		// 우선 매수 호가
	char fuofferjan[6];	char _fuofferjan;		// 우선 매도 잔량
	char fubidjan[6];	char _fubidjan;		// 우선 매수 잔량
	char futofferjan[6];	char _futofferjan;		// 총 매도 잔량
	char futbidjan[6];	char _futbidjan;		// 총 매수 잔량
	char fuoffersu[4];	char _fuoffersu;		// 우선 매도 건수
	char fubidsu[4];	char _fubidsu;		// 우선 매수 건수
	char futoffersu[5];	char _futoffersu;		// 총 매도 건수
	char futbidsu[5];	char _futbidsu;		// 총 매수 건수
	char fuchrate[5];	char _fuchrate;		// 등락률
	char fubasis[5];	char _fubasis;		// 베이시스
	char fugrate[5];	char _fugrate;		// 괴리도
	char fugratio[5];	char _fugratio;		// 괴리율
	char fupreopenyak[7];	char _fupreopenyak;		// 미결제약정전일
	char fuspvolall[7];	char _fuspvolall;		// 스프레드수량
	char fuopen[5];	char _fuopen;		// 시가
	char bulkvol[7];	char _bulkvol;		// 협의대량누적체결수량
} Tf8OutBlock;

typedef struct tagf8
{
	Tf8InBlock                        f8inblock                             ;		// 입력
	Tf8OutBlock                       f8outblock                            ;		// 출력
} Tf8;

typedef struct tagq1InBlock// 입력
{
	char fuspcode[8];		// 종목코드
} Tq1InBlock;

typedef struct tagq1OutBlock// 출력
{
	char fuspcode[8];		// 종목코드
	char fusphotime[8];		// 시간
	char fuspoffer[6];		// 우선 매도 호가
	char fuspbid[6];		// 우선 매수 호가
	char fuspofferjan[6];		// 우선 매도 잔량
	char fuspbidjan[6];		// 우선 매수 잔량
	char fuspjoffer[6];		// 차선 매도 호가
	char fuspjbid[6];		// 차선 매수 호가
	char fuspjofferjan[6];		// 차선 매도 잔량
	char fuspjbidjan[6];		// 차선 매수 잔량
	char fuspjjoffer[6];		// 차차선 매도 호가
	char fuspjjbid[6];		// 차차선 매수 호가
	char fuspjjofferjan[6];		// 차차선 매도 잔량
	char fuspjjbidjan[6];		// 차차선 매수 잔량
	char fuspj4offer[6];		// 4차선 매도 호가
	char fuspj4bid[6];		// 4차선 매수 호가
	char fuspj4offerjan[6];		// 4차선 매도 잔량
	char fuspj4bidjan[6];		// 4차선 매수 잔량
	char fuspj5offer[6];		// 5차선 매도 호가
	char fuspj5bid[6];		// 5차선 매수 호가
	char fuspj5offerjan[6];		// 5차선 매도 잔량
	char fuspj5bidjan[6];		// 5차선 매수 잔량
	char fusptofferjan[6];		// 총 매도 잔량
	char fusptbidjan[6];		// 총 매수 잔량
	char fuspoffersu[4];		// 우선 매도 건수
	char fuspjoffersu[4];		// 차선 매도 건수
	char fuspjjoffersu[4];		// 차차선 매도 건수
	char fuspj4offersu[4];		// 4차선 매도 건수
	char fuspj5offersu[4];		// 5차선 매도 건수
	char fusptoffersu[5];		// 총 매도 건수
	char fuspbidsu[4];		// 우선 매수 건수
	char fuspjbidsu[4];		// 차선 매수 건수
	char fuspjjbidsu[4];		// 차차선 매수 건수
	char fuspj4bidsu[4];		// 4차선 매수 건수
	char fuspj5bidsu[4];		// 5차선 매수 건수
	char fusptbidsu[5];		// 총 매수 건수
} Tq1OutBlock;

typedef struct tagq1
{
	Tq1InBlock                        q1inblock                             ;		// 입력
	Tq1OutBlock                       q1outblock                            ;		// 출력
} Tq1;

typedef struct tagq2InBlock// 입력
{
	char fuspcode[8];	char _fuspcode;		// 종목코드
} Tq2InBlock;

typedef struct tagq2OutBlock// 출력
{
	char fuspcode[8];	char _fuspcode;		// 종목코드
	char fusphotime[8];	char _fusphotime;		// 시간
	char fuspjgubun[8];	char _fuspjgubun;		// 장운용
	char fuspsign[1];	char _fuspsign;		// 전일부호
	char fuspchange[5];	char _fuspchange;		// 전일대비
	char fuspcurr[6];	char _fuspcurr;		// 현재가
	char fuspcurr1[5];	char _fuspcurr1;		// 의제약정가(근월)
	char fuspcurr2[5];	char _fuspcurr2;		// 의제약정가(원월)
	char fuspopen[6];	char _fuspopen;		// 시가
	char fusphigh[6];	char _fusphigh;		// 고가
	char fusplow[6];	char _fusplow;		// 저가
	char fuspvol[6];	char _fuspvol;		// 체결수량
	char fuspvolall[7];	char _fuspvolall;		// 누적체결수량
	char fuspvalall[12];	char _fuspvalall;		// 누적거래대금
	char fuspchrate[5];	char _fuspchrate;		// 등락율
	char fuspbp_jgubun[1];	char _fuspbp_jgubun;		// BP용장구분
	char fuspoffer[6];	char _fuspoffer;		// 우선 매도 호가
	char fuspbid[6];	char _fuspbid;		// 우선 매수 호가
} Tq2OutBlock;

typedef struct tagq2
{
	Tq2InBlock                        q2inblock                             ;		// 입력
	Tq2OutBlock                       q2outblock                            ;		// 출력
} Tq2;

typedef struct tago1InBlock// 입력
{
	char opitem[8];		// 코드
} To1InBlock;

typedef struct tago1OutBlock// 출력
{
	char opitem[8];		// 코드
	char ophotime[8];		// 시간
	char opoffer[5];		// 매도 우선호가
	char opbid[5];		// 매수 우선호가
	char opofferjan[7];		// 매도 우선잔량
	char opbidjan[7];		// 매수 우선잔량
	char opjoffer[5];		// 차선 매도 호가
	char opjbid[5];		// 차선 매수 호가
	char opjofferjan[7];		// 차선 매도 잔량
	char opjbidjan[7];		// 차선 매수 잔량
	char opjjoffer[5];		// 차차선 매도 호가
	char opjjbid[5];		// 차차선 매수 호가
	char opjjofferjan[7];		// 차차선 매도 잔량
	char opjjbidjan[7];		// 차차선 매수 잔량
	char optofferjan[7];		// 총 매도 잔량
	char optbidjan[7];		// 총 매수 잔량
	char opj4offer[5];		// 4차선 매도 호가
	char opj4bid[5];		// 4차선 매수 호가
	char opj4offerjan[7];		// 4차선 매도 잔량
	char opj4bidjan[7];		// 4차선 매수 잔량
	char opj5offer[5];		// 5차선 매도 호가
	char opj5bid[5];		// 5차선 매수 호가
	char opj5offerjan[7];		// 5차선 매도 잔량
	char opj5bidjan[7];		// 5차선 매수 잔량
	char opoffersu[4];		// 우선 매도 건수
	char opjoffersu[4];		// 차선 매도 건수
	char opjjoffersu[4];		// 차차선 매도 건수
	char opj4offersu[4];		// 4차선 매도 건수
	char opj5offersu[4];		// 5차선 매도 건수
	char optoffersu[5];		// 총 매도 건수
	char opbidsu[4];		// 우선 매수 건수
	char opjbidsu[4];		// 차선 매수 건수
	char opjjbidsu[4];		// 차차선 매수 건수
	char opj4bidsu[4];		// 4차선 매수 건수
	char opj5bidsu[4];		// 5차선 매수 건수
	char optbidsu[5];		// 총 매수 건수
} To1OutBlock;

typedef struct tago1
{
	To1InBlock                        o1inblock                             ;		// 입력
	To1OutBlock                       o1outblock                            ;		// 출력
} To1;

typedef struct tago2InBlock// 입력
{
	char opitem[8];	char _opitem;		// 종목코드
} To2InBlock;

typedef struct tago2OutBlock// 출력
{
	char opitem[8];	char _opitem;		// 종목코드
	char opchetime[8];	char _opchetime;		// 시간
	char opjgubun[8];	char _opjgubun;		// 장운용
	char opsign[1];	char _opsign;		// 등락부호
	char opchange[5];	char _opchange;		// 등락폭
	char opcurr[5];	char _opcurr;		// 현재가
	char opopen[5];	char _opopen;		// 시가
	char ophigh[5];	char _ophigh;		// 고가
	char oplow[5];	char _oplow;		// 저가
	char opvol[6];	char _opvol;		// 체결수량
	char opvolallz8[8];	char _opvolallz8;		// 누적체결수량
	char opvalall[12];	char _opvalall;		// 누적거래대금
	char opopenyak[7];	char _opopenyak;		// 미결제약정수량
	char opoffer[5];	char _opoffer;		// 우선 매도 호가
	char opbid[5];	char _opbid;		// 우선 매수 호가
	char opofferjan[7];	char _opofferjan;		// 우선 매도 잔량
	char opbidjan[7];	char _opbidjan;		// 우선 매수 잔량
	char opjoffer[5];	char _opjoffer;		// 차선 매도 호가
	char opjbid[5];	char _opjbid;		// 차선 매수 호가
	char opjofferjan[7];	char _opjofferjan;		// 차선 매도 잔량
	char opjbidjan[7];	char _opjbidjan;		// 차선 매수 잔량
	char opjjoffer[5];	char _opjjoffer;		// 차차선 매도 호가
	char opjjbid[5];	char _opjjbid;		// 차차선 매수 호가
	char opjjofferjan[7];	char _opjjofferjan;		// 차차선 매도 잔량
	char opjjbidjan[7];	char _opjjbidjan;		// 차차선 매수 잔량
	char optofferjan[7];	char _optofferjan;		// 총 매도 잔량
	char optbidjan[7];	char _optbidjan;		// 총 매수 잔량
	char opj4offer[5];	char _opj4offer;		// 4차선 매도 호가
	char opj4bid[5];	char _opj4bid;		// 4차선 매수 호가
	char opj4offerjan[7];	char _opj4offerjan;		// 4차선 매도 잔량
	char opj4bidjan[7];	char _opj4bidjan;		// 4차선 매수 잔량
	char opj5offer[5];	char _opj5offer;		// 5차선 매도 호가
	char opj5bid[5];	char _opj5bid;		// 5차선 매수 호가
	char opj5offerjan[7];	char _opj5offerjan;		// 5차선 매도 잔량
	char opj5bidjan[7];	char _opj5bidjan;		// 5차선 매수 잔량
	char opoffersu[4];	char _opoffersu;		// 우선 매도 건수
	char opjoffersu[4];	char _opjoffersu;		// 차선 매도 건수
	char opjjoffersu[4];	char _opjjoffersu;		// 차차선 매도 건수
	char opj4offersu[4];	char _opj4offersu;		// 4차선 매도 건수
	char opj5offersu[4];	char _opj5offersu;		// 5차선 매도 건수
	char optoffersu[5];	char _optoffersu;		// 총 매도 건수
	char opbidsu[4];	char _opbidsu;		// 우선 매수 건수
	char opjbidsu[4];	char _opjbidsu;		// 차선 매수 건수
	char opjjbidsu[4];	char _opjjbidsu;		// 차차선 매수 건수
	char opj4bidsu[4];	char _opj4bidsu;		// 4차선 매수 건수
	char opj5bidsu[4];	char _opj5bidsu;		// 5차선 매수 건수
	char optbidsu[5];	char _optbidsu;		// 총 매수 건수
	char opchrate[5];	char _opchrate;		// 등락률
	char opgrate[5];	char _opgrate;		// 괴리도
	char opgratio[5];	char _opgratio;		// 괴리율
	char oppreopenyak[7];	char _oppreopenyak;		// 미결제약정전일
	char opbp_jgubun[1];	char _opbp_jgubun;		// BP용장구분
	char bulkvolz8[8];	char _bulkvolz8;		// 혐의대량누적체결수량
} To2OutBlock;

typedef struct tago2
{
	To2InBlock                        o2inblock                             ;		// 입력
	To2OutBlock                       o2outblock                            ;		// 출력
} To2;

typedef struct tago3InBlock// 입력
{
	char opitem[8];		// 종목코드
} To3InBlock;

typedef struct tago3OutBlock// 출력
{
	char opitem[8];		// 종목코드
	char optheorytime[8];		// 이론가시간
	char optheoryprice[5];		// 옵션이론가
	char opimpv[5];		// 내재변동성
	char opdelta[8];		// 부호+델타
	char opgmma[8];		// 부호+감마
	char opvega[8];		// 부호+베가
	char optheta[8];		// 부호+세타
	char oprho[8];		// 부호+로
	char opgrate[5];		// 괴리도
	char opgratio[5];		// 괴리율
} To3OutBlock;

typedef struct tago3
{
	To3InBlock                        o3inblock                             ;		// 입력
	To3OutBlock                       o3outblock                            ;		// 출력
} To3;

typedef struct tago4InBlock// 입력
{
	char opitem[8];		// 종목코드
} To4InBlock;

typedef struct tago4OutBlock// 출력
{
	char opitem[8];		// 종목코드
	char opchetime[8];		// 체결시간
	char opopenyak[7];		// 미결제약정수량
	char oppreopenyak[7];		// 전일미결제약정수량
} To4OutBlock;

typedef struct tago4
{
	To4InBlock                        o4inblock                             ;		// 입력
	To4OutBlock                       o4outblock                            ;		// 출력
} To4;

typedef struct tagvHInBlock// 입력
{
	char fuitem[8];		// 종목코드
} TvHInBlock;

typedef struct tagvHOutBlock// 출력
{
	char fuitem[8];		// 종목코드
	char futime[8];		// 시간 HH:MM:SS
	char offer[7];		// 매도 호가
	char bid[7];		// 매수 호가
	char offerjan[6];		// 매도 잔량
	char bidjan[6];		// 매수 잔량
	char S2offer[7];		// 차 매도 호가
	char S2bid[7];		// 차 매수 호가
	char S2offerjan[6];		// 차 매도 잔량
	char S2bidjan[6];		// 차 매수 잔량
	char S3offer[7];		// 차차 매도 호가
	char S3bid[7];		// 차차 매수 호가
	char S3offerjan[6];		// 차차 매도 잔량
	char S3bidjan[6];		// 차차 매수 잔량
	char S4offer[7];		// 4차 매도 호가
	char S4bid[7];		// 4차 매수 호가
	char S4offerjan[6];		// 4차 매도 잔량
	char S4bidjan[6];		// 4차 매수 잔량
	char S5offer[7];		// 5차 매도 호가
	char S5bid[7];		// 5차 매수 호가
	char S5offerjan[6];		// 5차 매도 잔량
	char S5bidjan[6];		// 5차 매수 잔량
	char S6offer[7];		// 6차 매도 호가
	char S6bid[7];		// 6차 매수 호가
	char S6offerjan[6];		// 6차 매도 잔량
	char S6bidjan[6];		// 6차 매수 잔량
	char S7offer[7];		// 7차 매도 호가
	char S7bid[7];		// 7차 매수 호가
	char S7offerjan[6];		// 7차 매도 잔량
	char S7bidjan[6];		// 7차 매수 잔량
	char S8offer[7];		// 8차 매도 호가
	char S8bid[7];		// 8차 매수 호가
	char S8offerjan[6];		// 8차 매도 잔량
	char S8bidjan[6];		// 8차 매수 잔량
	char S9offer[7];		// 9차 매도 호가
	char S9bid[7];		// 9차 매수 호가
	char S9offerjan[6];		// 9차 매도 잔량
	char S9bidjan[6];		// 9차 매수 잔량
	char S0offer[7];		// 10차 매도 호가
	char S0bid[7];		// 10차 매수 호가
	char S0offerjan[6];		// 10차 매도 잔량
	char S0bidjan[6];		// 10차 매수 잔량
	char offersu[4];		// 우 선 매도 건수
	char bidsu[4];		// 우 선 매수 건수
	char S2offersu[4];		// 차 선 매도 건수
	char S2bidsu[4];		// 차 선 매수 건수
	char S3offersu[4];		// 3차선 매도 건수
	char S3bidsu[4];		// 3차선 매수 건수
	char S4offersu[4];		// 4차선 매도 건수
	char S4bidsu[4];		// 4차선 매수 건수
	char S5offersu[4];		// 5차선 매도 건수
	char S5bidsu[4];		// 5차선 매수 건수
	char S6offersu[4];		// 6차선 매도 건수
	char S6bidsu[4];		// 6차선 매수 건수
	char S7offersu[4];		// 7차선 매도 건수
	char S7bidsu[4];		// 7차선 매수 건수
	char S8offersu[4];		// 8차선 매도 건수
	char S8bidsu[4];		// 8차선 매수 건수
	char S9offersu[4];		// 9차선 매도 건수
	char S9bidsu[4];		// 9차선 매수 건수
	char S0offersu[4];		// 10차선 매도 건수
	char S0bidsu[4];		// 10차선 매수 건수
	char tofferjan[6];		// 총 매도 호가 잔량
	char tobidjan[6];		// 총 매수 호가 잔량
	char toffersu[5];		// 총 매도 건수
	char tbidsu[5];		// 총 매수 건수
} TvHOutBlock;

typedef struct tagvH
{
	TvHInBlock                        vhinblock                             ;		// 입력
	TvHOutBlock                       vhoutblock                            ;		// 출력
} TvH;

typedef struct tagvCInBlock// 입력
{
	char fuitem[8];	char _fuitem;		// 종목코드
} TvCInBlock;

typedef struct tagvCOutBlock// 출력
{
	char fuitem[8];	char _fuitem;		// 종목코드
	char futime[8];	char _futime;		// 시간 HH:MM:SS
	char jgubun[8];	char _jgubun;		// 장운용
	char fusign[1];	char _fusign;		// 전일대비 부호
	char fuchange[7];	char _fuchange;		// 전일대비
	char fucurr[7];	char _fucurr;		// 현재가
	char fuopen[7];	char _fuopen;		// 시가
	char fuhigh[7];	char _fuhigh;		// 고가
	char fulow[7];	char _fulow;		// 저가
	char fuvol[6];	char _fuvol;		// 체결수량
	char fuvolall[7];	char _fuvolall;		// 누적 체결수량
	char fuvalall[12];	char _fuvalall;		// 누적거래대금
	char openyak[7];	char _openyak;		// 미결제약정수량
	char jandatecnt[3];	char _jandatecnt;		// 잔존일수
	char offer[7];	char _offer;		// 매도 호가
	char bid[7];	char _bid;		// 매수 호가
	char offerjan[6];	char _offerjan;		// 매도 잔량
	char bidjan[6];	char _bidjan;		// 매수 잔량
	char S2offer[7];	char _S2offer;		// 차 매도 호가
	char S2bid[7];	char _S2bid;		// 차 매수 호가
	char S2offerjan[6];	char _S2offerjan;		// 차 매도 잔량
	char S2bidjan[6];	char _S2bidjan;		// 차 매수 잔량
	char S3offer[7];	char _S3offer;		// 차차 매도 호가
	char S3bid[7];	char _S3bid;		// 차차 매수 호가
	char S3offerjan[6];	char _S3offerjan;		// 차차 매도 잔량
	char S3bidjan[6];	char _S3bidjan;		// 차차 매수 잔량
	char S4offer[7];	char _S4offer;		// 4차 매도 호가
	char S4bid[7];	char _S4bid;		// 4차 매수 호가
	char S4offerjan[6];	char _S4offerjan;		// 4차 매도 잔량
	char S4bidjan[6];	char _S4bidjan;		// 4차 매수 잔량
	char S5offer[7];	char _S5offer;		// 5차 매도 호가
	char S5bid[7];	char _S5bid;		// 5차 매수 호가
	char S5offerjan[6];	char _S5offerjan;		// 5차 매도 잔량
	char S5bidjan[6];	char _S5bidjan;		// 5차 매수 잔량
	char S6offer[7];	char _S6offer;		// 6차 매도 호가
	char S6bid[7];	char _S6bid;		// 6차 매수 호가
	char S6offerjan[6];	char _S6offerjan;		// 6차 매도 잔량
	char S6bidjan[6];	char _S6bidjan;		// 6차 매수 잔량
	char S7offer[7];	char _S7offer;		// 7차 매도 호가
	char S7bid[7];	char _S7bid;		// 7차 매수 호가
	char S7offerjan[6];	char _S7offerjan;		// 7차 매도 잔량
	char S7bidjan[6];	char _S7bidjan;		// 7차 매수 잔량
	char S8offer[7];	char _S8offer;		// 8차 매도 호가
	char S8bid[7];	char _S8bid;		// 8차 매수 호가
	char S8offerjan[6];	char _S8offerjan;		// 8차 매도 잔량
	char S8bidjan[6];	char _S8bidjan;		// 8차 매수 잔량
	char S9offer[7];	char _S9offer;		// 9차 매도 호가
	char S9bid[7];	char _S9bid;		// 9차 매수 호가
	char S9offerjan[6];	char _S9offerjan;		// 9차 매도 잔량
	char S9bidjan[6];	char _S9bidjan;		// 9차 매수 잔량
	char S0offer[7];	char _S0offer;		// 10차 매도 호가
	char S0bid[7];	char _S0bid;		// 10차 매수 호가
	char S0offerjan[6];	char _S0offerjan;		// 10차 매도 잔량
	char S0bidjan[6];	char _S0bidjan;		// 10차 매수 잔량
	char offersu[4];	char _offersu;		// 우 선 매도 건수
	char bidsu[4];	char _bidsu;		// 우 선 매수 건수
	char S2offersu[4];	char _S2offersu;		// 차 선 매도 건수
	char S2bidsu[4];	char _S2bidsu;		// 차 선 매수 건수
	char S3offersu[4];	char _S3offersu;		// 3차선 매도 건수
	char S3bidsu[4];	char _S3bidsu;		// 3차선 매수 건수
	char S4offersu[4];	char _S4offersu;		// 4차선 매도 건수
	char S4bidsu[4];	char _S4bidsu;		// 4차선 매수 건수
	char S5offersu[4];	char _S5offersu;		// 5차선 매도 건수
	char S5bidsu[4];	char _S5bidsu;		// 5차선 매수 건수
	char S6offersu[4];	char _S6offersu;		// 6차선 매도 건수
	char S6bidsu[4];	char _S6bidsu;		// 6차선 매수 건수
	char S7offersu[4];	char _S7offersu;		// 7차선 매도 건수
	char S7bidsu[4];	char _S7bidsu;		// 7차선 매수 건수
	char S8offersu[4];	char _S8offersu;		// 8차선 매도 건수
	char S8bidsu[4];	char _S8bidsu;		// 8차선 매수 건수
	char S9offersu[4];	char _S9offersu;		// 9차선 매도 건수
	char S9bidsu[4];	char _S9bidsu;		// 9차선 매수 건수
	char S0offersu[4];	char _S0offersu;		// 10차선 매도 건수
	char S0bidsu[4];	char _S0bidsu;		// 10차선 매수 건수
	char tofferjan[6];	char _tofferjan;		// 총 매도 호가 잔량
	char tobidjan[6];	char _tobidjan;		// 총 매수 호가 잔량
	char toffersu[5];	char _toffersu;		// 총 매도 건수
	char tbidsu[5];	char _tbidsu;		// 총 매수 건수
	char chrate[5];	char _chrate;		// 등락률
	char basis[7];	char _basis;		// 베이시스
	char grate[7];	char _grate;		// 괴리도
	char gratio[6];	char _gratio;		// 괴리율
	char preopenyak[7];	char _preopenyak;		// 미결제약정전일비
	char bp_jgubun[1];	char _bp_jgubun;		// BP용 장구분
	char fspvolall[7];	char _fspvolall;		// 스프레드 체결수량
} TvCOutBlock;

typedef struct tagvC
{
	TvCInBlock                        vcinblock                             ;		// 입력
	TvCOutBlock                       vcoutblock                            ;		// 출력
} TvC;

typedef struct tagvVInBlock// 입력
{
	char fuitem[8];		// 종목코드
} TvVInBlock;

typedef struct tagvVOutBlock// 출력
{
	char fuitem[8];		// 종목코드
	char theoryprice[7];		// 선물 이론가
	char theorytime[8];		// 이론가 시간
	char grate[7];		// 괴리도
	char gratio[6];		// 괴리율
} TvVOutBlock;

typedef struct tagvV
{
	TvVInBlock                        vvinblock                             ;		// 입력
	TvVOutBlock                       vvoutblock                            ;		// 출력
} TvV;

typedef struct tagvMInBlock// 입력
{
	char fuitem[8];		// 종목코드
} TvMInBlock;

typedef struct tagvMOutBlock// 출력
{
	char fuitem[8];		// 종목코드
	char chetime[8];		// 체결시간
	char openyak[7];		// 미결제약정수량
	char preopenyak[7];		// 전일미결제약정수량
} TvMOutBlock;

typedef struct tagvM
{
	TvMInBlock                        vminblock                             ;		// 입력
	TvMOutBlock                       vmoutblock                            ;		// 출력
} TvM;

typedef struct tagv7InBlock// 입력
{
	char fspitem[8];		// 스프레드종목코드
} Tv7InBlock;

typedef struct tagv7OutBlock// 출력
{
	char fspitem[8];		// 스프레드종목코드
	char fsptime[8];		// 시간 HH:MM:SS
	char offer[8];		// 매도 호가
	char bid[8];		// 매수 호가
	char offerjan[6];		// 매도 잔량
	char bidjan[6];		// 매수 잔량
	char S2offer[8];		// 차 매도 호가
	char S2bid[8];		// 차 매수 호가
	char S2offerjan[6];		// 차 매도 잔량
	char S2bidjan[6];		// 차 매수 잔량
	char S3offer[8];		// 차차 매도 호가
	char S3bid[8];		// 차차 매수 호가
	char S3offerjan[6];		// 차차 매도 잔량
	char S3bidjan[6];		// 차차 매수 잔량
	char S4offer[8];		// 4차 매도 호가
	char S4bid[8];		// 4차 매수 호가
	char S4offerjan[6];		// 4차 매도 잔량
	char S4bidjan[6];		// 4차 매수 잔량
	char S5offer[8];		// 5차 매도 호가
	char S5bid[8];		// 5차 매수 호가
	char S5offerjan[6];		// 5차 매도 잔량
	char S5bidjan[6];		// 5차 매수 잔량
	char S6offer[8];		// 6차 매도 호가
	char S6bid[8];		// 6차 매수 호가
	char S6offerjan[6];		// 6차 매도 잔량
	char S6bidjan[6];		// 6차 매수 잔량
	char S7offer[8];		// 7차 매도 호가
	char S7bid[8];		// 7차 매수 호가
	char S7offerjan[6];		// 7차 매도 잔량
	char S7bidjan[6];		// 7차 매수 잔량
	char S8offer[8];		// 8차 매도 호가
	char S8bid[8];		// 8차 매수 호가
	char S8offerjan[6];		// 8차 매도 잔량
	char S8bidjan[6];		// 8차 매수 잔량
	char S9offer[8];		// 9차 매도 호가
	char S9bid[8];		// 9차 매수 호가
	char S9offerjan[6];		// 9차 매도 잔량
	char S9bidjan[6];		// 9차 매수 잔량
	char S0offer[8];		// 10차 매도 호가
	char S0bid[8];		// 10차 매수 호가
	char S0offerjan[6];		// 10차 매도 잔량
	char S0bidjan[6];		// 10차 매수 잔량
	char offersu[4];		// 우 선 매도 건수
	char bidsu[4];		// 우 선 매수 건수
	char S2offersu[4];		// 차 선 매도 건수
	char S2bidsu[4];		// 차 선 매수 건수
	char S3offersu[4];		// 3차선 매도 건수
	char S3bidsu[4];		// 3차선 매수 건수
	char S4offersu[4];		// 4차선 매도 건수
	char S4bidsu[4];		// 4차선 매수 건수
	char S5offersu[4];		// 5차선 매도 건수
	char S5bidsu[4];		// 5차선 매수 건수
	char S6offersu[4];		// 6차선 매도 건수
	char S6bidsu[4];		// 6차선 매수 건수
	char S7offersu[4];		// 7차선 매도 건수
	char S7bidsu[4];		// 7차선 매수 건수
	char S8offersu[4];		// 8차선 매도 건수
	char S8bidsu[4];		// 8차선 매수 건수
	char S9offersu[4];		// 9차선 매도 건수
	char S9bidsu[4];		// 9차선 매수 건수
	char S0offersu[4];		// 10차선 매도 건수
	char S0bidsu[4];		// 10차선 매수 건수
	char tofferjan[6];		// 총 매도 호가 잔량
	char tobidjan[6];		// 총 매수 호가 잔량
	char toffersu[5];		// 총 매도 건수
	char tbidsu[5];		// 총 매수 건수
} Tv7OutBlock;

typedef struct tagv7
{
	Tv7InBlock                        v7inblock                             ;		// 입력
	Tv7OutBlock                       v7outblock                            ;		// 출력
} Tv7;

typedef struct tagv8InBlock// 입력
{
	char fspitem[8];	char _fspitem;		// 스프레드종목코드
} Tv8InBlock;

typedef struct tagv8OutBlock// 출력
{
	char fspitem[8];	char _fspitem;		// 스프레드종목코드
	char fsptime[8];	char _fsptime;		// 시간 HH:MM:SS
	char jgubun[8];	char _jgubun;		// 장운용
	char fspsign[1];	char _fspsign;		// 전일대비 부호
	char fspchange[7];	char _fspchange;		// 전일대비
	char fspcurr[8];	char _fspcurr;		// 현재가
	char fspcurr1[7];	char _fspcurr1;		// 의제약정가-근월물
	char fspcurr2[7];	char _fspcurr2;		// 의제약정가-원월물
	char fspopen[8];	char _fspopen;		// 시가
	char fsphigh[8];	char _fsphigh;		// 고가
	char fsplow[8];	char _fsplow;		// 저가
	char fspvol[6];	char _fspvol;		// 체결수량
	char fspvolall[7];	char _fspvolall;		// 누적 체결수량
	char fspvalall[12];	char _fspvalall;		// 누적거래대금
	char offer[8];	char _offer;		// 매도 호가
	char bid[8];	char _bid;		// 매수 호가
	char offerjan[6];	char _offerjan;		// 매도 잔량
	char bidjan[6];	char _bidjan;		// 매수 잔량
	char S2offer[8];	char _S2offer;		// 차 매도 호가
	char S2bid[8];	char _S2bid;		// 차 매수 호가
	char S2offerjan[6];	char _S2offerjan;		// 차 매도 잔량
	char S2bidjan[6];	char _S2bidjan;		// 차 매수 잔량
	char S3offer[8];	char _S3offer;		// 차차 매도 호가
	char S3bid[8];	char _S3bid;		// 차차 매수 호가
	char S3offerjan[6];	char _S3offerjan;		// 차차 매도 잔량
	char S3bidjan[6];	char _S3bidjan;		// 차차 매수 잔량
	char S4offer[8];	char _S4offer;		// 4차 매도 호가
	char S4bid[8];	char _S4bid;		// 4차 매수 호가
	char S4offerjan[6];	char _S4offerjan;		// 4차 매도 잔량
	char S4bidjan[6];	char _S4bidjan;		// 4차 매수 잔량
	char S5offer[8];	char _S5offer;		// 5차 매도 호가
	char S5bid[8];	char _S5bid;		// 5차 매수 호가
	char S5offerjan[6];	char _S5offerjan;		// 5차 매도 잔량
	char S5bidjan[6];	char _S5bidjan;		// 5차 매수 잔량
	char S6offer[8];	char _S6offer;		// 6차 매도 호가
	char S6bid[8];	char _S6bid;		// 6차 매수 호가
	char S6offerjan[6];	char _S6offerjan;		// 6차 매도 잔량
	char S6bidjan[6];	char _S6bidjan;		// 6차 매수 잔량
	char S7offer[8];	char _S7offer;		// 7차 매도 호가
	char S7bid[8];	char _S7bid;		// 7차 매수 호가
	char S7offerjan[6];	char _S7offerjan;		// 7차 매도 잔량
	char S7bidjan[6];	char _S7bidjan;		// 7차 매수 잔량
	char S8offer[8];	char _S8offer;		// 8차 매도 호가
	char S8bid[8];	char _S8bid;		// 8차 매수 호가
	char S8offerjan[6];	char _S8offerjan;		// 8차 매도 잔량
	char S8bidjan[6];	char _S8bidjan;		// 8차 매수 잔량
	char S9offer[8];	char _S9offer;		// 9차 매도 호가
	char S9bid[8];	char _S9bid;		// 9차 매수 호가
	char S9offerjan[6];	char _S9offerjan;		// 9차 매도 잔량
	char S9bidjan[6];	char _S9bidjan;		// 9차 매수 잔량
	char S0offer[8];	char _S0offer;		// 10차 매도 호가
	char S0bid[8];	char _S0bid;		// 10차 매수 호가
	char S0offerjan[6];	char _S0offerjan;		// 10차 매도 잔량
	char S0bidjan[6];	char _S0bidjan;		// 10차 매수 잔량
	char offersu[4];	char _offersu;		// 우 선 매도 건수
	char bidsu[4];	char _bidsu;		// 우 선 매수 건수
	char S2offersu[4];	char _S2offersu;		// 차 선 매도 건수
	char S2bidsu[4];	char _S2bidsu;		// 차 선 매수 건수
	char S3offersu[4];	char _S3offersu;		// 3차선 매도 건수
	char S3bidsu[4];	char _S3bidsu;		// 3차선 매수 건수
	char S4offersu[4];	char _S4offersu;		// 4차선 매도 건수
	char S4bidsu[4];	char _S4bidsu;		// 4차선 매수 건수
	char S5offersu[4];	char _S5offersu;		// 5차선 매도 건수
	char S5bidsu[4];	char _S5bidsu;		// 5차선 매수 건수
	char S6offersu[4];	char _S6offersu;		// 6차선 매도 건수
	char S6bidsu[4];	char _S6bidsu;		// 6차선 매수 건수
	char S7offersu[4];	char _S7offersu;		// 7차선 매도 건수
	char S7bidsu[4];	char _S7bidsu;		// 7차선 매수 건수
	char S8offersu[4];	char _S8offersu;		// 8차선 매도 건수
	char S8bidsu[4];	char _S8bidsu;		// 8차선 매수 건수
	char S9offersu[4];	char _S9offersu;		// 9차선 매도 건수
	char S9bidsu[4];	char _S9bidsu;		// 9차선 매수 건수
	char S0offersu[4];	char _S0offersu;		// 10차선 매도 건수
	char S0bidsu[4];	char _S0bidsu;		// 10차선 매수 건수
	char tofferjan[6];	char _tofferjan;		// 총 매도 호가 잔량
	char tobidjan[6];	char _tobidjan;		// 총 매수 호가 잔량
	char toffersu[5];	char _toffersu;		// 총 매도 건수
	char tbidsu[5];	char _tbidsu;		// 총 매수 건수
	char chrate[5];	char _chrate;		// 등락률
	char bp_jgubun[1];	char _bp_jgubun;		// BP용 장구분
} Tv8OutBlock;

typedef struct tagv8
{
	Tv8InBlock                        v8inblock                             ;		// 입력
	Tv8OutBlock                       v8outblock                            ;		// 출력
} Tv8;

typedef struct tageCInBlock// 입력
{
	char Code[6];	char _Code;		// 종목코드
} TeCInBlock;

typedef struct tageCOutBlock// 출력
{
	char Code[6];	char _Code;		// 단축종목코드
	char time[8];	char _time;		// 시간(HH:MM:SS)
	char price[7];	char _price;		// 현재가
	char sign[1];	char _sign;		// 등락부호
	char change[6];	char _change;		// 등락폭
	char chrate[5];	char _chrate;		// 등락률
	char open[7];	char _open;		// 시가
	char high[7];	char _high;		// 고가
	char low[7];	char _low;		// 저가
	char offer[7];	char _offer;		// 매도 호가
	char bid[7];	char _bid;		// 매수 호가
	char volume[9];	char _volume;		// 누적거래량
	char volrate[6];	char _volrate;		// 거래량 전일비
	char movolume[8];	char _movolume;		// 변동거래량
	char value[9];	char _value;		// 거래대금 백만원
	char janggubun[1];	char _janggubun;		// 장구분
	char cbgubun[1];	char _cbgubun;		// CB구분
	char stop[1];	char _stop;		// STOP
	char grate[6];	char _grate;		// 괴리도 9(6)
	char gratio[8];	char _gratio;		// 괴리율S9(5)V9(2)
	char lphold[9];	char _lphold;		// LP보유수량
	char lprate[5];	char _lprate;		// LP보유률
} TeCOutBlock;

typedef struct tageC
{
	TeCInBlock                        ecinblock                             ;		// 입력
	TeCOutBlock                       ecoutblock                            ;		// 출력
} TeC;

typedef struct tageHInBlock// 입력
{
	char Code[6];		// 종목코드
} TeHInBlock;

typedef struct tageHOutBlock// 출력
{
	char Code[6];		// 단축코드
	char time[8];		// 시간
	char S1_off[7];		// 매도 호가
	char S1_bid[7];		// 매수 호가
	char S1_offrem[9];		// 매도 호가 잔량
	char S1_bidrem[9];		// 매수 호가 잔량
	char S2_off[7];		// 2차 매도 호가
	char S2_bid[7];		// 2차 매수 호가
	char S2_offrem[9];		// 2차 매도 호가 잔량
	char S2_bidrem[9];		// 2차 매수 호가 잔량
	char S3_off[7];		// 3차 매도 호가
	char S3_bid[7];		// 3차 매수 호가
	char S3_offrem[9];		// 3차 매도 호가 잔량
	char S3_bidrem[9];		// 3차 매수 호가 잔량
	char S4_off[7];		// 4차 매도 호가
	char S4_bid[7];		// 4차 매수 호가
	char S4_offrem[9];		// 4차 매도 호가 잔량
	char S4_bidrem[9];		// 4차 매수 호가 잔량
	char S5_off[7];		// 5차 매도 호가
	char S5_bid[7];		// 5차 매수 호가
	char S5_offrem[9];		// 5차 매도 호가 잔량
	char S5_bidrem[9];		// 5차 매수 호가 잔량
	char S6_off[7];		// 6차 매도 호가
	char S6_bid[7];		// 6차 매수 호가
	char S6_offrem[9];		// 6차 매도 호가 잔량
	char S6_bidrem[9];		// 6차 매수 호가 잔량
	char S7_off[7];		// 7차 매도 호가
	char S7_bid[7];		// 7차 매수 호가
	char S7_offrem[9];		// 7차 매도 호가 잔량
	char S7_bidrem[9];		// 7차 매수 호가 잔량
	char S8_off[7];		// 8차 매도 호가
	char S8_bid[7];		// 8차 매수 호가
	char S8_offrem[9];		// 8차 매도 호가 잔량
	char S8_bidrem[9];		// 8차 매수 호가 잔량
	char S9_off[7];		// 9차 매도 호가
	char S9_bid[7];		// 9차 매수 호가
	char S9_offrem[9];		// 9차 매도 호가 잔량
	char S9_bidrem[9];		// 9차 매수 호가 잔량
	char S10_off[7];		// 10차 매도 호가
	char S10_bid[7];		// 10차 매수 호가
	char S10_offrem[9];		// 10차 매도 호가 잔량
	char S10_bidrem[9];		// 10차 매수 호가 잔량
	char T_offrem[9];		// 총 매도 호가 잔량
	char T_bidrem[9];		// 총 매수 호가 잔량
	char dongsi[1];		// 동시구분
	char eqprice[7];		// 동시호가시예상체결가
	char sign[1];		// 등락부호
	char change[6];		// 등락폭
	char chrate[5];		// 등락률
	char eqvol[9];		// 동시호가시예상체결수량
	char S1_lpoffrem[9];		// 1차LP 매도 호가 잔량
	char S1_lpbidrem[9];		// 1차LP 매수 호가 잔량
	char S2_lpoffrem[9];		// 2차LP 매도 호가 잔량
	char S2_lpbidrem[9];		// 2차LP 매수 호가 잔량
	char S3_lpoffrem[9];		// 3차LP 매도 호가 잔량
	char S3_lpbidrem[9];		// 3차LP 매수 호가 잔량
	char S4_lpoffrem[9];		// 4차LP 매도 호가 잔량
	char S4_lpbidrem[9];		// 4차LP 매수 호가 잔량
	char S5_lpoffrem[9];		// 5차LP 매도 호가 잔량
	char S5_lpbidrem[9];		// 5차LP 매수 호가 잔량
	char S6_lpoffrem[9];		// 6차LP 매도 호가 잔량
	char S6_lpbidrem[9];		// 6차LP 매수 호가 잔량
	char S7_lpoffrem[9];		// 7차LP 매도 호가 잔량
	char S7_lpbidrem[9];		// 7차LP 매수 호가 잔량
	char S8_lpoffrem[9];		// 8차LP 매도 호가 잔량
	char S8_lpbidrem[9];		// 8차LP 매수 호가 잔량
	char S9_lpoffrem[9];		// 9차LP 매도 호가 잔량
	char S9_lpbidrem[9];		// 9차LP 매수 호가 잔량
	char S10_lpoffrem[9];		// 10차LP 매도 호가 잔량
	char S10_lpbidrem[9];		// 10차LP 매수 호가 잔량
} TeHOutBlock;

typedef struct tageH
{
	TeHInBlock                        ehinblock                             ;		// 입력
	TeHOutBlock                       ehoutblock                            ;		// 출력
} TeH;

typedef struct tageVInBlock// 입력
{
	char Code[6];		// 종목코드
} TeVInBlock;

typedef struct tageVOutBlock// 출력
{
	char Code[6];		// 종목코드
	char theorytime[8];		// 이론가시간
	char theoryprice[7];		// 이론가
	char impv[10];		// 내재변동성
	char delta[9];		// 부호+델타
	char gmma[9];		// 부호+감마
	char vega[9];		// 부호+베가
	char theta[9];		// 부호+세타
	char rho[9];		// 부호+로
	char grate[6];		// 괴리도
	char gratio[8];		// 괴리율
} TeVOutBlock;

typedef struct tageV
{
	TeVInBlock                        evinblock                             ;		// 입력
	TeVOutBlock                       evoutblock                            ;		// 출력
} TeV;

typedef struct tageLInBlock// 입력
{
	char Code[6];		// 종목코드
} TeLInBlock;

typedef struct tageLOutBlock// 출력
{
	char Code[6];		// 종목코드
	char jipyotime[8];		// 투자지표시간
	char parity[8];		// 패리티
	char egearing[8];		// E기어링
	char gearingrate[8];		// 기어링비율
	char profitrate[8];		// 손익분기율
	char basepoint[8];		// 자본지지점
	char filler[6];		// FILLER
} TeLOutBlock;

typedef struct tageL
{
	TeLInBlock                        elinblock                             ;		// 입력
	TeLOutBlock                       eloutblock                            ;		// 출력
} TeL;

typedef struct tageTInBlock// 입력
{
	char Code[6];	char _Code;		// 종목코드
} TeTInBlock;

typedef struct tageTOutBlock// 출력
{
	char Code[6];	char _Code;		// 종목코드
	char time[8];	char _time;		// 시간
	char off_trano1[4];	char _off_trano1;		// 매도 회원사코드1
	char off_tra1[6];	char _off_tra1;		// 매도 회원사단이름1
	char N_off_tra1[20];	char _N_off_tra1;		// 매도 회원사장이름1
	char N_otraflag1[1];	char _N_otraflag1;		// 매도 회원사외구분1
	char N_offvolume1[9];	char _N_offvolume1;		// 매도 거래량1
	char N_offvolcha1[9];	char _N_offvolcha1;		// 직전 매도 차1
	char bid_trano1[4];	char _bid_trano1;		// 매수 회원사코드1
	char bid_tra1[6];	char _bid_tra1;		// 매수 회원사단이름1
	char N_bid_tra1[20];	char _N_bid_tra1;		// 매수 회원사장이름1
	char N_btraflag1[1];	char _N_btraflag1;		// 매수 회원사외구분1
	char N_bidvolume1[9];	char _N_bidvolume1;		// 매수 거래량1
	char N_bidvolcha1[9];	char _N_bidvolcha1;		// 직전 매수 차1
	char off_trano2[4];	char _off_trano2;		// 매도 회원사코드2
	char off_tra2[6];	char _off_tra2;		// 매도 회원사단이름2
	char N_off_tra2[20];	char _N_off_tra2;		// 매도 회원사장이름2
	char N_otraflag2[1];	char _N_otraflag2;		// 매도 회원사외구분2
	char N_offvolume2[9];	char _N_offvolume2;		// 매도 거래량2
	char N_offvolcha2[9];	char _N_offvolcha2;		// 직전 매도 차2
	char bid_trano2[4];	char _bid_trano2;		// 매수 회원사코드2
	char bid_tra2[6];	char _bid_tra2;		// 매수 회원사단이름2
	char N_bid_tra2[20];	char _N_bid_tra2;		// 매수 회원사장이름2
	char N_btraflag2[1];	char _N_btraflag2;		// 매수 회원사외구분2
	char N_bidvolume2[9];	char _N_bidvolume2;		// 매수 거래량2
	char N_bidvolcha2[9];	char _N_bidvolcha2;		// 직전 매수 차2
	char off_trano3[4];	char _off_trano3;		// 매도 회원사코드3
	char off_tra3[6];	char _off_tra3;		// 매도 회원사단이름3
	char N_off_tra3[20];	char _N_off_tra3;		// 매도 회원사장이름3
	char N_otraflag3[1];	char _N_otraflag3;		// 매도 회원사외구분3
	char N_offvolume3[9];	char _N_offvolume3;		// 매도 거래량3
	char N_offvolcha3[9];	char _N_offvolcha3;		// 직전 매도 차3
	char bid_trano3[4];	char _bid_trano3;		// 매수 회원사코드3
	char bid_tra3[6];	char _bid_tra3;		// 매수 회원사단이름3
	char N_bid_tra3[20];	char _N_bid_tra3;		// 매수 회원사장이름3
	char N_btraflag3[1];	char _N_btraflag3;		// 매수 회원사외구분3
	char N_bidvolume3[9];	char _N_bidvolume3;		// 매수 거래량3
	char N_bidvolcha3[9];	char _N_bidvolcha3;		// 직전 매수 차3
	char off_trano4[4];	char _off_trano4;		// 매도 회원사코드4
	char off_tra4[6];	char _off_tra4;		// 매도 회원사단이름4
	char N_off_tra4[20];	char _N_off_tra4;		// 매도 회원사장이름4
	char N_otraflag4[1];	char _N_otraflag4;		// 매도 회원사외구분4
	char N_offvolume4[9];	char _N_offvolume4;		// 매도 거래량4
	char N_offvolcha4[9];	char _N_offvolcha4;		// 직전 매도 차4
	char bid_trano4[4];	char _bid_trano4;		// 매수 회원사코드4
	char bid_tra4[6];	char _bid_tra4;		// 매수 회원사단이름4
	char N_bid_tra4[20];	char _N_bid_tra4;		// 매수 회원사장이름4
	char N_btraflag4[1];	char _N_btraflag4;		// 매수 회원사외구분4
	char N_bidvolume4[9];	char _N_bidvolume4;		// 매수 거래량4
	char N_bidvolcha4[9];	char _N_bidvolcha4;		// 직전 매수 차4
	char off_trano5[4];	char _off_trano5;		// 매도 회원사코드5
	char off_tra5[6];	char _off_tra5;		// 매도 회원사단이름5
	char N_off_tra5[20];	char _N_off_tra5;		// 매도 회원사장이름5
	char N_otraflag5[1];	char _N_otraflag5;		// 매도 회원사외구분5
	char N_offvolume5[9];	char _N_offvolume5;		// 매도 거래량5
	char N_offvolcha5[9];	char _N_offvolcha5;		// 직전 매도 차5
	char bid_trano5[4];	char _bid_trano5;		// 매수 회원사코드5
	char bid_tra5[6];	char _bid_tra5;		// 매수 회원사단이름5
	char N_bid_tra5[20];	char _N_bid_tra5;		// 매수 회원사장이름5
	char N_btraflag5[1];	char _N_btraflag5;		// 매수 회원사외구분5
	char N_bidvolume5[9];	char _N_bidvolume5;		// 매수 거래량5
	char N_bidvolcha5[9];	char _N_bidvolcha5;		// 직전 매수 차5
	char N_offvolall[9];	char _N_offvolall;		// 외국계회원사 매도 합
	char N_offvolcha[9];	char _N_offvolcha;		// 외국계직전 매도 차
	char N_bidvolall[9];	char _N_bidvolall;		// 외국계회원사 매수 합
	char N_bidvolcha[9];	char _N_bidvolcha;		// 외국계직전 매수 차
	char N_soonmaesu[9];	char _N_soonmaesu;		// 외국계회원순 매수
	char N_soonmaecha[9];	char _N_soonmaecha;		// 외국계직전순 매수 차
	char N_alloffvol[9];	char _N_alloffvol;		// 매도 전체합
	char N_allbidvol[9];	char _N_allbidvol;		// 매수 전체합
	char Title[13];	char _Title;		// 종목명
	char kpgubun[1];	char _kpgubun;		// 시장구분
} TeTOutBlock;

typedef struct tageT
{
	TeTInBlock                        etinblock                             ;		// 입력
	TeTOutBlock                       etoutblock                            ;		// 출력
} TeT;

typedef struct tagfEInBlock// 입력
{
	char fuitem[4];		// 종목코드
} TfEInBlock;

typedef struct tagfEOutBlock// 출력
{
	char fuitem[4];		// 종목코드
	char time[8];		// 시간
	char dongsi[1];		// 동시호가구분
	char eqsign[1];		// 예상등락부호
	char eqprice[5];		// 예상체결가
	char eqchange[5];		// 예상등락폭
	char eqchrate[5];		// 예상등락률
} TfEOutBlock;

typedef struct tagfE
{
	TfEInBlock                        feinblock                             ;		// 입력
	TfEOutBlock                       feoutblock                            ;		// 출력
} TfE;

typedef struct tagoEInBlock// 입력
{
	char opitem[8];		// 종목코드
} ToEInBlock;

typedef struct tagoEOutBlock// 출력
{
	char opitem[8];		// 종목코드
	char time[8];		// 시간
	char dongsi[1];		// 동시호가구분
	char eqsign[1];		// 예상등락부호
	char eqprice[5];		// 예상체결가
	char eqchange[5];		// 예상등락폭
	char eqchrate[5];		// 예상등락률
} ToEOutBlock;

typedef struct tagoE
{
	ToEInBlock                        oeinblock                             ;		// 입력
	ToEOutBlock                       oeoutblock                            ;		// 출력
} ToE;

typedef struct tagvEInBlock// 입력
{
	char expcode[8];		// 종목코드
} TvEInBlock;

typedef struct tagvEOutBlock// 출력
{
	char expcode[8];		// 종목코드
	char time[8];		// 시간
	char dongsi[1];		// 동시호가구분
	char eqsign[1];		// 예상등락부호
	char eqprice[7];		// 예상체결가
	char eqchange[7];		// 예상등락폭
	char eqchrate[5];		// 예상등락률
} TvEOutBlock;

typedef struct tagvE
{
	TvEInBlock                        veinblock                             ;		// 입력
	TvEOutBlock                       veoutblock                            ;		// 출력
} TvE;

typedef struct tagf7InBlock// 입력
{
	char fuitem[4];	char _fuitem;		// 종목코드
} Tf7InBlock;

typedef struct tagf7OutBlock// 출력
{
	char fuitem[4];	char _fuitem;		// 종목코드
	char futime[8];	char _futime;		// 시간
	char exlmtstep[1];	char _exlmtstep;		// 가격확대예정단계
	char exlmtgb[1];	char _exlmtgb;		// 가격확대예정 구분
	char uplmtgb[1];	char _uplmtgb;		// 가격제한확대상한단계
	char dnlmtgb[1];	char _dnlmtgb;		// 가격제한확대하한단계
	char uplmtprice[5];	char _uplmtprice;		// 적용된 단계 상한가
	char dnlmtprice[5];	char _dnlmtprice;		// 적용된 단계 하한가
} Tf7OutBlock;

typedef struct tagf7
{
	Tf7InBlock                        f7inblock                             ;		// 입력
	Tf7OutBlock                       f7outblock                            ;		// 출력
} Tf7;

typedef struct tago7InBlock// 입력
{
	char opitem[8];	char _opitem;		// 종목코드
} To7InBlock;

typedef struct tago7OutBlock// 출력
{
	char opitem[8];	char _opitem;		// 종목코드
	char optime[8];	char _optime;		// 시간
	char uplmtgb[1];	char _uplmtgb;		// 가격제한확대상한단계
	char dnlmtgb[1];	char _dnlmtgb;		// 가격제한확대하한단계
	char uplmtprice[5];	char _uplmtprice;		// 적용된 단계 상한가
	char dnlmtprice[5];	char _dnlmtprice;		// 적용된 단계 하한가
} To7OutBlock;

typedef struct tago7
{
	To7InBlock                        o7inblock                             ;		// 입력
	To7OutBlock                       o7outblock                            ;		// 출력
} To7;

typedef struct tagvIInBlock// 입력
{
	char fuitem[8];	char _fuitem;		// 종목코드
} TvIInBlock;

typedef struct tagvIOutBlock// 출력
{
	char fuitem[8];	char _fuitem;		// 종목코드
	char futime[8];	char _futime;		// 시간
	char exlmtstep[1];	char _exlmtstep;		// 가격확대예정단계
	char exlmtgb[1];	char _exlmtgb;		// 가격확대예정 구분
	char uplmtgb[1];	char _uplmtgb;		// 가격제한확대상한단계
	char dnlmtgb[1];	char _dnlmtgb;		// 가격제한확대하한단계
	char uplmtprice[7];	char _uplmtprice;		// 적용된 단계 상한가
	char dnlmtprice[7];	char _dnlmtprice;		// 적용된 단계 하한가
} TvIOutBlock;

typedef struct tagvI
{
	TvIInBlock                        viinblock                             ;		// 입력
	TvIOutBlock                       vioutblock                            ;		// 출력
} TvI;

typedef struct tagu1InBlock// 입력
{
	char jisucode[2];	char _jisucode;		// 업종코드
} Tu1InBlock;

typedef struct tagu1OutBlock// 출력
{
	char jisucode[2];	char _jisucode;		// 업종코드
	char jisutime[8];	char _jisutime;		// 시간
	char jisu[8];	char _jisu;		// 지수
	char jisusign[1];	char _jisusign;		// 등락부호
	char jisuchange[8];	char _jisuchange;		// 등락폭
	char jisuvolume[8];	char _jisuvolume;		// 거래량
	char jisuvalue[8];	char _jisuvalue;		// 거래대금
	char jisuopen[8];	char _jisuopen;		// 시가지수
	char jisuhigh[8];	char _jisuhigh;		// 고가지수
	char jisuhightime[8];	char _jisuhightime;		// 고가시간
	char jisulow[8];	char _jisulow;		// 저가지수
	char jisulowtime[8];	char _jisulowtime;		// 저가시간
	char jisuchrate[5];	char _jisuchrate;		// 지수등락률
	char jisubrkvol[5];	char _jisubrkvol;		// 거래비중
} Tu1OutBlock;

typedef struct tagu1
{
	Tu1InBlock                        u1inblock                             ;		// 입력
	Tu1OutBlock                       u1outblock                            ;		// 출력
} Tu1;

typedef struct tagk1InBlock// 입력
{
	char jisukcode[2];	char _jisukcode;		// 업종코드
} Tk1InBlock;

typedef struct tagk1OutBlock// 출력
{
	char jisukcode[2];	char _jisukcode;		// 업종코드
	char jisuktime[8];	char _jisuktime;		// 시간
	char jisuk[8];	char _jisuk;		// 지수
	char jisuksign[1];	char _jisuksign;		// 등락부호
	char jisukchange[8];	char _jisukchange;		// 등락폭
	char jisukvolume[8];	char _jisukvolume;		// 거래량
	char jisukvalue[8];	char _jisukvalue;		// 거래대금
	char jisukopen[8];	char _jisukopen;		// 시가지수
	char jisukhigh[8];	char _jisukhigh;		// 고가지수
	char jisukhightime[8];	char _jisukhightime;		// 고가시간
	char jisuklow[8];	char _jisuklow;		// 저가지수
	char jisuklowtime[8];	char _jisuklowtime;		// 저가시간
	char jisukchrate[5];	char _jisukchrate;		// 지수등락률
	char jisukbrkvol[5];	char _jisukbrkvol;		// 거래비중
} Tk1OutBlock;

typedef struct tagk1
{
	Tk1InBlock k1inblock;		// 입력
	Tk1OutBlock k1outblock;		// 출력
} Tk1;
