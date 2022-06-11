/* 증권사 제공 DLL함수 레퍼런스와 예제소스 코드의 'IXingAPI.h'파일,
* DevCenter의 C++ 헤더를 참조해서 약간 수정.
* COPIED FROM API provider's reference and sample source code.
* MODIFIED by GHTS Authors.
* LICENSING TERM follows that of original code.
*
* 저작권 관련 규정은 레퍼런스 및 헤더 파일, 샘플 소스코드의 원래 저작권 규정을 따름.
*
* 변수명명규칙
* Go언어와 데이터를 주고 받는 구조체의 멤버 변수는 Go언어와의 호환성을 위해서,
* Go언어에서 public형은 대문자로 시작해야 함.
* C 헤더 파일은 'go tool cgo -godefs'로 바로 Go자료형으로 변환해서 사용하므로,
* 구조체 멤버 필드의 경우 Go언어의 변수명명 규칙을 여기에서도 적용해서 첫 글자를 대문자로 함.
* 그 외 최근 주류 언어인 Java, C#의 관례에 따라 CamelCase를 적용함. */

// '#pragma pack()'으로 지정된 PACKED C구조체는 Go언어에서 직접 사용할 수 없음.
// binary.Read()를 이용해서 Go 구조체로 읽어들인 후 사용할 수 있음.
#pragma pack(push, 1)

//------------------------------------------------------------------------------
// 기본 구조체
//------------------------------------------------------------------------------
// '#pragma pack(push, 1)'이 적용된 기본 구조체들은 Go언어에서 읽을 수 없으며, C언어에서만 사용함..
//  Go언어에서 읽을 때는 '#pragma pack(push, 1)'이 적용되지 않은 UNPACKED 자료형으로 변환해서 사용함.
//  메모리 저장 방식으로 변환은 dll32 패키지 내 콜백 Go함수에서 binary.encoding 로 수행.
typedef struct {    // 조회TR 수신 패킷
	int					RequestID;					// Request ID
	int					DataLength;				    // 받은 데이터 크기
	int					TotalDataBufferSize;		// lpData에 할당된 크기
	int					ElapsedTime;				// 전송에서 수신까지 걸린시간(1/1000초)
	int					DataMode;					// 현재 의미 없음. (1:BLOCK MODE, 2:NON-BLOCK MODE) 20
	char TrCode[10]; char _TrCode[1];			    // TR Code
	char				Cont[1];       			    // 다음조회 없음 : '0', 'N', 다음조회 있음 : '1', '2'
	char ContKey[18]; char _ContKey[1];		        // 연속키, Data Header가 B 인 경우에만 사용
	char				None[31];                   //szUserData[31]; // 사용자 데이터 (사용 안 함) 62
	char BlockName[16]; char _BlockName[1];		    // Block 명, Block Mode 일때만 사용  17
	unsigned char*		Data;                        // 수신된 TR 데이터
} TR_DATA;

typedef struct {    // 실시간TR 수신 패킷
	char TrCode[3]; char _TrCode[1];		    // TR Code
	int					KeyLength;                 // 뭐지??
	char KeyData[32]; char _KeyData[1];         // 뭐지??
	char RegKey[32]; char _RegKey[1];          // 뭐지??
	int					DataLength;                // 받은 데이터 크기
	char*				Data;                    // 실시간 데이터
} REALTIME_DATA;

typedef struct {    // 메시지 수신 패킷
	int					RequestID;			        // Request ID
	int					SystemError;				// 0:일반메시지, 1:시스템 에러 메시지
	char MsgCode[5]; char _MsgCode[1];              // 메시지 코드
	int					MsgLength;					// 메시지 데이터 길이
	char*		        MsgData;                // 메시지 데이터
} MSG_DATA;

typedef	struct {    // HTS-> API로 연동데이터 수신 패킷
    char                LinkName[32];             // 연동 키 ex) 주식 종목코드 연동 시, &STOCK_CODE
    char                LinkData[32];             // 연동 값 ex) 주식 종목코드 연동 시, 종목코드
    char                None[64];                   // 사용 안 함
} LINK_DATA;

//------------------------------------------------------------------------------
// 현물 정상주문 (CSPAT00600,ENCRYPT,SIGNATURE,HEADTYPE=B)
//------------------------------------------------------------------------------
typedef struct {
    char    acntNo[20];    //[string,   20] 계좌번호   StartPos 0, Length 20
    char    inptPwd[8];    //[string,    8] 입력비밀번호   StartPos 20, Length 8
    char    isuNo[12];    //[string,   12] 종목번호   StartPos 28, Length 12
    char    ordQty[16];    //[long  ,   16] 주문수량   StartPos 40, Length 16
    char    ordPrc[13];    //[double, 13.2] 주문가   StartPos 56, Length 13
    char    bnsTpCode[1];    //[string,    1] 매매구분   StartPos 69, Length 1
    char    ordprcPtnCode[2];    //[string,    2] 호가유형코드   StartPos 70, Length 2
    char    mgntrnCode[3];    //[string,    3] 신용거래코드   StartPos 72, Length 3
    char    loanDt[8];    //[string,    8] 대출일   StartPos 75, Length 8
    char    ordCndiTpCode[1];    //[string,    1] 주문조건구분   StartPos 83, Length 1
} CSPAT00600InBlock1;

typedef struct {
    char    recCnt[5];    //[long  ,    5] 레코드갯수   StartPos 0, Length 5
    char    acntNo[20];    //[string,   20] 계좌번호   StartPos 5, Length 20
    char    inptPwd[8];    //[string,    8] 입력비밀번호   StartPos 25, Length 8
    char    isuNo[12];    //[string,   12] 종목번호   StartPos 33, Length 12
    char    ordQty[16];    //[long  ,   16] 주문수량   StartPos 45, Length 16
    char    ordPrc[13];    //[double, 13.2] 주문가   StartPos 61, Length 13
    char    bnsTpCode[1];    //[string,    1] 매매구분   StartPos 74, Length 1
    char    ordprcPtnCode[2];    //[string,    2] 호가유형코드   StartPos 75, Length 2
    char    prgmOrdprcPtnCode[2];    //[string,    2] 프로그램호가유형코드   StartPos 77, Length 2
    char    stslAbleYn[1];    //[string,    1] 공매도가능여부   StartPos 79, Length 1
    char    stslOrdprcTpCode[1];    //[string,    1] 공매도호가구분   StartPos 80, Length 1
    char    commdaCode[2];    //[string,    2] 통신매체코드   StartPos 81, Length 2
    char    mgntrnCode[3];    //[string,    3] 신용거래코드   StartPos 83, Length 3
    char    loanDt[8];    //[string,    8] 대출일   StartPos 86, Length 8
    char    mbrNo[3];    //[string,    3] 회원번호   StartPos 94, Length 3
    char    ordCndiTpCode[1];    //[string,    1] 주문조건구분   StartPos 97, Length 1
    char    strtgCode[6];    //[string,    6] 전략코드   StartPos 98, Length 6
    char    grpId[20];    //[string,   20] 그룹ID   StartPos 104, Length 20
    char    ordSeqNo[10];    //[long  ,   10] 주문회차   StartPos 124, Length 10
    char    ptflNo[10];    //[long  ,   10] 포트폴리오번호   StartPos 134, Length 10
    char    bskNo[10];    //[long  ,   10] 바스켓번호   StartPos 144, Length 10
    char    trchNo[10];    //[long  ,   10] 트렌치번호   StartPos 154, Length 10
    char    itemNo[10];    //[long  ,   10] 아이템번호   StartPos 164, Length 10
    char    opDrtnNo[12];    //[string,   12] 운용지시번호   StartPos 174, Length 12
    char    lpYn[1];    //[string,    1] 유동성공급자여부   StartPos 186, Length 1
    char    cvrgTpCode[1];    //[string,    1] 반대매매구분   StartPos 187, Length 1
} CSPAT00600OutBlock1;

typedef struct {
    char    recCnt[5];    //[long  ,    5] 레코드갯수   StartPos 0, Length 5
    char    ordNo[10];    //[long  ,   10] 주문번호   StartPos 5, Length 10
    char    ordTime[9];    //[string,    9] 주문시각   StartPos 15, Length 9
    char    ordMktCode[2];    //[string,    2] 주문시장코드   StartPos 24, Length 2
    char    ordPtnCode[2];    //[string,    2] 주문유형코드   StartPos 26, Length 2
    char    shtnIsuNo[9];    //[string,    9] 단축종목번호   StartPos 28, Length 9
    char    mgempNo[9];    //[string,    9] 관리사원번호   StartPos 37, Length 9
    char    ordAmt[16];    //[long  ,   16] 주문금액   StartPos 46, Length 16
    char    spareOrdNo[10];    //[long  ,   10] 예비주문번호   StartPos 62, Length 10
    char    cvrgSeqno[10];    //[long  ,   10] 반대매매일련번호   StartPos 72, Length 10
    char    rsvOrdNo[10];    //[long  ,   10] 예약주문번호   StartPos 82, Length 10
    char    spotOrdQty[16];    //[long  ,   16] 실물주문수량   StartPos 92, Length 16
    char    ruseOrdQty[16];    //[long  ,   16] 재사용주문수량   StartPos 108, Length 16
    char    mnyOrdAmt[16];    //[long  ,   16] 현금주문금액   StartPos 124, Length 16
    char    substOrdAmt[16];    //[long  ,   16] 대용주문금액   StartPos 140, Length 16
    char    ruseOrdAmt[16];    //[long  ,   16] 재사용주문금액   StartPos 156, Length 16
    char    acntNm[40];    //[string,   40] 계좌명   StartPos 172, Length 40
    char    isuNm[40];    //[string,   40] 종목명   StartPos 212, Length 40
} CSPAT00600OutBlock2;

typedef struct {
    CSPAT00600OutBlock1	outBlock1;
    CSPAT00600OutBlock2	outBlock2;
} CSPAT00600OutBlock;

//------------------------------------------------------------------------------
// 현물 정정주문 (CSPAT00700,ENCRYPT,SIGNATURE,HEADTYPE=B)
//------------------------------------------------------------------------------
typedef struct {
    char    orgOrdNo[10];    //[long  ,   10] 원주문번호   StartPos 0, Length 10
    char    acntNo[20];    //[string,   20] 계좌번호   StartPos 10, Length 20
    char    inptPwd[8];    //[string,    8] 입력비밀번호   StartPos 30, Length 8
    char    isuNo[12];    //[string,   12] 종목번호   StartPos 38, Length 12
    char    ordQty[16];    //[long  ,   16] 주문수량   StartPos 50, Length 16
    char    ordprcPtnCode[2];    //[string,    2] 호가유형코드   StartPos 66, Length 2
    char    ordCndiTpCode[1];    //[string,    1] 주문조건구분   StartPos 68, Length 1
    char    ordPrc[13];    //[double, 13.2] 주문가   StartPos 69, Length 13
} CSPAT00700InBlock1;

typedef struct {
    char    recCnt[5];    //[long  ,    5] 레코드갯수   StartPos 0, Length 5
    char    orgOrdNo[10];    //[long  ,   10] 원주문번호   StartPos 5, Length 10
    char    acntNo[20];    //[string,   20] 계좌번호   StartPos 15, Length 20
    char    inptPwd[8];    //[string,    8] 입력비밀번호   StartPos 35, Length 8
    char    isuNo[12];    //[string,   12] 종목번호   StartPos 43, Length 12
    char    ordQty[16];    //[long  ,   16] 주문수량   StartPos 55, Length 16
    char    ordprcPtnCode[2];    //[string,    2] 호가유형코드   StartPos 71, Length 2
    char    ordCndiTpCode[1];    //[string,    1] 주문조건구분   StartPos 73, Length 1
    char    ordPrc[13];    //[double, 13.2] 주문가   StartPos 74, Length 13
    char    commdaCode[2];    //[string,    2] 통신매체코드   StartPos 87, Length 2
    char    strtgCode[6];    //[string,    6] 전략코드   StartPos 89, Length 6
    char    grpId[20];    //[string,   20] 그룹ID   StartPos 95, Length 20
    char    ordSeqNo[10];    //[long  ,   10] 주문회차   StartPos 115, Length 10
    char    ptflNo[10];    //[long  ,   10] 포트폴리오번호   StartPos 125, Length 10
    char    bskNo[10];    //[long  ,   10] 바스켓번호   StartPos 135, Length 10
    char    trchNo[10];    //[long  ,   10] 트렌치번호   StartPos 145, Length 10
    char    itemNo[10];    //[long  ,   10] 아이템번호   StartPos 155, Length 10
} CSPAT00700OutBlock1;

typedef struct {
    char    recCnt[5];    //[long  ,    5] 레코드갯수   StartPos 0, Length 5
    char    ordNo[10];    //[long  ,   10] 주문번호   StartPos 5, Length 10
    char    prntOrdNo[10];    //[long  ,   10] 모주문번호   StartPos 15, Length 10
    char    ordTime[9];    //[string,    9] 주문시각   StartPos 25, Length 9
    char    ordMktCode[2];    //[string,    2] 주문시장코드   StartPos 34, Length 2
    char    ordPtnCode[2];    //[string,    2] 주문유형코드   StartPos 36, Length 2
    char    shtnIsuNo[9];    //[string,    9] 단축종목번호   StartPos 38, Length 9
    char    prgmOrdprcPtnCode[2];    //[string,    2] 프로그램호가유형코드   StartPos 47, Length 2
    char    stslOrdprcTpCode[1];    //[string,    1] 공매도호가구분   StartPos 49, Length 1
    char    stslAbleYn[1];    //[string,    1] 공매도가능여부   StartPos 50, Length 1
    char    mgntrnCode[3];    //[string,    3] 신용거래코드   StartPos 51, Length 3
    char    loanDt[8];    //[string,    8] 대출일   StartPos 54, Length 8
    char    cvrgOrdTp[1];    //[string,    1] 반대매매주문구분   StartPos 62, Length 1
    char    lpYn[1];    //[string,    1] 유동성공급자여부   StartPos 63, Length 1
    char    mgempNo[9];    //[string,    9] 관리사원번호   StartPos 64, Length 9
    char    ordAmt[16];    //[long  ,   16] 주문금액   StartPos 73, Length 16
    char    bnsTpCode[1];    //[string,    1] 매매구분   StartPos 89, Length 1
    char    spareOrdNo[10];    //[long  ,   10] 예비주문번호   StartPos 90, Length 10
    char    cvrgSeqno[10];    //[long  ,   10] 반대매매일련번호   StartPos 100, Length 10
    char    rsvOrdNo[10];    //[long  ,   10] 예약주문번호   StartPos 110, Length 10
    char    mnyOrdAmt[16];    //[long  ,   16] 현금주문금액   StartPos 120, Length 16
    char    substOrdAmt[16];    //[long  ,   16] 대용주문금액   StartPos 136, Length 16
    char    ruseOrdAmt[16];    //[long  ,   16] 재사용주문금액   StartPos 152, Length 16
    char    acntNm[40];    //[string,   40] 계좌명   StartPos 168, Length 40
    char    isuNm[40];    //[string,   40] 종목명   StartPos 208, Length 40
} CSPAT00700OutBlock2;

typedef struct {
    CSPAT00700OutBlock1	outBlock1;
    CSPAT00700OutBlock2	outBlock2;
} CSPAT00700OutBlock;

//------------------------------------------------------------------------------
// 현물 취소주문 (CSPAT00800,ENCRYPT,SIGNATURE,HEADTYPE=B)
//------------------------------------------------------------------------------
typedef struct {
    char    orgOrdNo[10];    //[long  ,   10] 원주문번호   StartPos 0, Length 10
    char    acntNo[20];    //[string,   20] 계좌번호   StartPos 10, Length 20
    char    inptPwd[8];    //[string,    8] 입력비밀번호   StartPos 30, Length 8
    char    isuNo[12];    //[string,   12] 종목번호   StartPos 38, Length 12
    char    ordQty[16];    //[long  ,   16] 주문수량   StartPos 50, Length 16
} CSPAT00800InBlock1;

typedef struct {
    char    recCnt[5];    //[long  ,    5] 레코드갯수   StartPos 0, Length 5
    char    orgOrdNo[10];    //[long  ,   10] 원주문번호   StartPos 5, Length 10
    char    acntNo[20];    //[string,   20] 계좌번호   StartPos 15, Length 20
    char    inptPwd[8];    //[string,    8] 입력비밀번호   StartPos 35, Length 8
    char    isuNo[12];    //[string,   12] 종목번호   StartPos 43, Length 12
    char    ordQty[16];    //[long  ,   16] 주문수량   StartPos 55, Length 16
    char    commdaCode[2];    //[string,    2] 통신매체코드   StartPos 71, Length 2
    char    grpId[20];    //[string,   20] 그룹ID   StartPos 73, Length 20
    char    strtgCode[6];    //[string,    6] 전략코드   StartPos 93, Length 6
    char    ordSeqNo[10];    //[long  ,   10] 주문회차   StartPos 99, Length 10
    char    ptflNo[10];    //[long  ,   10] 포트폴리오번호   StartPos 109, Length 10
    char    bskNo[10];    //[long  ,   10] 바스켓번호   StartPos 119, Length 10
    char    trchNo[10];    //[long  ,   10] 트렌치번호   StartPos 129, Length 10
    char    itemNo[10];    //[long  ,   10] 아이템번호   StartPos 139, Length 10
} CSPAT00800OutBlock1;

typedef struct {
    char    recCnt[5];    //[long  ,    5] 레코드갯수   StartPos 0, Length 5
    char    ordNo[10];    //[long  ,   10] 주문번호   StartPos 5, Length 10
    char    prntOrdNo[10];    //[long  ,   10] 모주문번호   StartPos 15, Length 10
    char    ordTime[9];    //[string,    9] 주문시각   StartPos 25, Length 9
    char    ordMktCode[2];    //[string,    2] 주문시장코드   StartPos 34, Length 2
    char    ordPtnCode[2];    //[string,    2] 주문유형코드   StartPos 36, Length 2
    char    shtnIsuNo[9];    //[string,    9] 단축종목번호   StartPos 38, Length 9
    char    prgmOrdprcPtnCode[2];    //[string,    2] 프로그램호가유형코드   StartPos 47, Length 2
    char    stslOrdprcTpCode[1];    //[string,    1] 공매도호가구분   StartPos 49, Length 1
    char    stslAbleYn[1];    //[string,    1] 공매도가능여부   StartPos 50, Length 1
    char    mgntrnCode[3];    //[string,    3] 신용거래코드   StartPos 51, Length 3
    char    loanDt[8];    //[string,    8] 대출일   StartPos 54, Length 8
    char    cvrgOrdTp[1];    //[string,    1] 반대매매주문구분   StartPos 62, Length 1
    char    lpYn[1];    //[string,    1] 유동성공급자여부   StartPos 63, Length 1
    char    mgempNo[9];    //[string,    9] 관리사원번호   StartPos 64, Length 9
    char    bnsTpCode[1];    //[string,    1] 매매구분   StartPos 73, Length 1
    char    spareOrdNo[10];    //[long  ,   10] 예비주문번호   StartPos 74, Length 10
    char    cvrgSeqno[10];    //[long  ,   10] 반대매매일련번호   StartPos 84, Length 10
    char    rsvOrdNo[10];    //[long  ,   10] 예약주문번호   StartPos 94, Length 10
    char    acntNm[40];    //[string,   40] 계좌명   StartPos 104, Length 40
    char    isuNm[40];    //[string,   40] 종목명   StartPos 144, Length 40
} CSPAT00800OutBlock2;

typedef struct {
    CSPAT00800OutBlock1	outBlock1;
    CSPAT00800OutBlock2	outBlock2;
} CSPAT00800OutBlock;

//------------------------------------------------------------------------------
// 현물 주문 접수 실시간 정보 (SC0)
//------------------------------------------------------------------------------
typedef struct {
    char    lineseq             [  10];    // [long  ,   10] 라인일련번호                   StartPos 0, Length 10
    char    accno               [  11];    // [string,   11] 계좌번호                       StartPos 10, Length 11
    char    user                [   8];    // [string,    8] 조작자ID                       StartPos 21, Length 8
    char    len                 [   6];    // [long  ,    6] 헤더길이                       StartPos 29, Length 6
    char    gubun               [   1];    // [string,    1] 헤더구분                       StartPos 35, Length 1
    char    compress            [   1];    // [string,    1] 압축구분                       StartPos 36, Length 1
    char    encrypt             [   1];    // [string,    1] 암호구분                       StartPos 37, Length 1
    char    offset              [   3];    // [long  ,    3] 공통시작지점                   StartPos 38, Length 3
    char    trcode              [   8];    // [string,    8] TRCODE                         StartPos 41, Length 8
    char    compid              [   3];    // [string,    3] 이용사번호                     StartPos 49, Length 3
    char    userid              [  16];    // [string,   16] 사용자ID                       StartPos 52, Length 16
    char    media               [   2];    // [string,    2] 접속매체                       StartPos 68, Length 2
    char    ifid                [   3];    // [string,    3] I/F일련번호                    StartPos 70, Length 3
    char    seq                 [   9];    // [string,    9] 전문일련번호                   StartPos 73, Length 9
    char    trid                [  16];    // [string,   16] TR추적ID                       StartPos 82, Length 16
    char    pubip               [  12];    // [string,   12] 공인IP                         StartPos 98, Length 12
    char    prvip               [  12];    // [string,   12] 사설IP                         StartPos 110, Length 12
    char    pcbpno              [   3];    // [string,    3] 처리지점번호                   StartPos 122, Length 3
    char    bpno                [   3];    // [string,    3] 지점번호                       StartPos 125, Length 3
    char    termno              [   8];    // [string,    8] 단말번호                       StartPos 128, Length 8
    char    lang                [   1];    // [string,    1] 언어구분                       StartPos 136, Length 1
    char    proctm              [   9];    // [long  ,    9] AP처리시간                     StartPos 137, Length 9
    char    msgcode             [   4];    // [string,    4] 메세지코드                     StartPos 146, Length 4
    char    outgu               [   1];    // [string,    1] 메세지출력구분                 StartPos 150, Length 1
    char    compreq             [   1];    // [string,    1] 압축요청구분                   StartPos 151, Length 1
    char    funckey             [   4];    // [string,    4] 기능키                         StartPos 152, Length 4
    char    reqcnt              [   4];    // [long  ,    4] 요청레코드개수                 StartPos 156, Length 4
    char    filler              [   6];    // [string,    6] 예비영역                       StartPos 160, Length 6
    char    cont                [   1];    // [string,    1] 연속구분                       StartPos 166, Length 1
    char    contkey             [  18];    // [string,   18] 연속키값                       StartPos 167, Length 18
    char    varlen              [   2];    // [long  ,    2] 가변시스템길이                 StartPos 185, Length 2
    char    varhdlen            [   2];    // [long  ,    2] 가변해더길이                   StartPos 187, Length 2
    char    varmsglen           [   2];    // [long  ,    2] 가변메시지길이                 StartPos 189, Length 2
    char    trsrc               [   1];    // [string,    1] 조회발원지                     StartPos 191, Length 1
    char    eventid             [   4];    // [string,    4] I/F이벤트ID                    StartPos 192, Length 4
    char    ifinfo              [   4];    // [string,    4] I/F정보                        StartPos 196, Length 4
    char    filler1             [  41];    // [string,   41] 예비영역                       StartPos 200, Length 41
    char    ordchegb            [   2];    // [string,    2] 주문체결구분                   StartPos 241, Length 2
    char    marketgb            [   2];    // [string,    2] 시장구분                       StartPos 243, Length 2
    char    ordgb               [   2];    // [string,    2] 주문구분                       StartPos 245, Length 2
    char    orgordno            [  10];    // [long  ,   10] 원주문번호                     StartPos 247, Length 10
    char    accno1              [  11];    // [string,   11] 계좌번호                       StartPos 257, Length 11
    char    accno2              [   9];    // [string,    9] 계좌번호                       StartPos 268, Length 9
    char    passwd              [   8];    // [string,    8] 비밀번호                       StartPos 277, Length 8
    char    expcode             [  12];    // [string,   12] 종목번호                       StartPos 285, Length 12
    char    shtcode             [   9];    // [string,    9] 단축종목번호                   StartPos 297, Length 9
    char    hname               [  40];    // [string,   40] 종목명                         StartPos 306, Length 40
    char    ordqty              [  16];    // [long  ,   16] 주문수량                       StartPos 346, Length 16
    char    ordprice            [  13];    // [long  ,   13] 주문가격                       StartPos 362, Length 13
    char    hogagb              [   1];    // [string,    1] 주문조건                       StartPos 375, Length 1
    char    etfhogagb           [   2];    // [string,    2] 호가유형코드                   StartPos 376, Length 2
    char    pgmtype             [   2];    // [long  ,    2] 프로그램호가구분               StartPos 378, Length 2
    char    gmhogagb            [   1];    // [long  ,    1] 공매도호가구분                 StartPos 380, Length 1
    char    gmhogayn            [   1];    // [long  ,    1] 공매도가능여부                 StartPos 381, Length 1
    char    singb               [   3];    // [string,    3] 신용구분                       StartPos 382, Length 3
    char    loandt              [   8];    // [string,    8] 대출일                         StartPos 385, Length 8
    char    cvrgordtp           [   1];    // [string,    1] 반대매매주문구분               StartPos 393, Length 1
    char    strtgcode           [   6];    // [string,    6] 전략코드                       StartPos 394, Length 6
    char    groupid             [  20];    // [string,   20] 그룹ID                         StartPos 400, Length 20
    char    ordseqno            [  10];    // [long  ,   10] 주문회차                       StartPos 420, Length 10
    char    prtno               [  10];    // [long  ,   10] 포트폴리오번호                 StartPos 430, Length 10
    char    basketno            [  10];    // [long  ,   10] 바스켓번호                     StartPos 440, Length 10
    char    trchno              [  10];    // [long  ,   10] 트렌치번호                     StartPos 450, Length 10
    char    itemno              [  10];    // [long  ,   10] 아아템번호                     StartPos 460, Length 10
    char    brwmgmyn            [   1];    // [long  ,    1] 차입구분                       StartPos 470, Length 1
    char    mbrno               [   3];    // [long  ,    3] 회원사번호                     StartPos 471, Length 3
    char    procgb              [   1];    // [string,    1] 처리구분                       StartPos 474, Length 1
    char    admbrchno           [   3];    // [string,    3] 관리지점번호                   StartPos 475, Length 3
    char    futaccno            [  20];    // [string,   20] 선물계좌번호                   StartPos 478, Length 20
    char    futmarketgb         [   1];    // [string,    1] 선물상품구분                   StartPos 498, Length 1
    char    tongsingb           [   2];    // [string,    2] 통신매체구분                   StartPos 499, Length 2
    char    lpgb                [   1];    // [string,    1] 유동성공급자구분               StartPos 501, Length 1
    char    dummy               [  20];    // [string,   20] DUMMY                          StartPos 502, Length 20
    char    ordno               [  10];    // [long  ,   10] 주문번호                       StartPos 522, Length 10
    char    ordtm               [   9];    // [string,    9] 주문시각                       StartPos 532, Length 9
    char    prntordno           [  10];    // [long  ,   10] 모주문번호                     StartPos 541, Length 10
    char    mgempno             [   9];    // [string,    9] 관리사원번호                   StartPos 551, Length 9
    char    orgordundrqty       [  16];    // [long  ,   16] 원주문미체결수량               StartPos 560, Length 16
    char    orgordmdfyqty       [  16];    // [long  ,   16] 원주문정정수량                 StartPos 576, Length 16
    char    ordordcancelqty     [  16];    // [long  ,   16] 원주문취소수량                 StartPos 592, Length 16
    char    nmcpysndno          [  10];    // [long  ,   10] 비회원사송신번호               StartPos 608, Length 10
    char    ordamt              [  16];    // [long  ,   16] 주문금액                       StartPos 618, Length 16
    char    bnstp               [   1];    // [string,    1] 매매구분                       StartPos 634, Length 1
    char    spareordno          [  10];    // [long  ,   10] 예비주문번호                   StartPos 635, Length 10
    char    cvrgseqno           [  10];    // [long  ,   10] 반대매매일련번호               StartPos 645, Length 10
    char    rsvordno            [  10];    // [long  ,   10] 예약주문번호                   StartPos 655, Length 10
    char    mtordseqno          [  10];    // [long  ,   10] 복수주문일련번호               StartPos 665, Length 10
    char    spareordqty         [  16];    // [long  ,   16] 예비주문수량                   StartPos 675, Length 16
    char    orduserid           [  16];    // [string,   16] 주문사원번호                   StartPos 691, Length 16
    char    spotordqty          [  16];    // [long  ,   16] 실물주문수량                   StartPos 707, Length 16
    char    ordruseqty          [  16];    // [long  ,   16] 재사용주문수량                 StartPos 723, Length 16
    char    mnyordamt           [  16];    // [long  ,   16] 현금주문금액                   StartPos 739, Length 16
    char    ordsubstamt         [  16];    // [long  ,   16] 주문대용금액                   StartPos 755, Length 16
    char    ruseordamt          [  16];    // [long  ,   16] 재사용주문금액                 StartPos 771, Length 16
    char    ordcmsnamt          [  16];    // [long  ,   16] 수수료주문금액                 StartPos 787, Length 16
    char    crdtuseamt          [  16];    // [long  ,   16] 사용신용담보재사용금           StartPos 803, Length 16
    char    secbalqty           [  16];    // [long  ,   16] 잔고수량                       StartPos 819, Length 16
    char    spotordableqty      [  16];    // [long  ,   16] 실물가능수량                   StartPos 835, Length 16
    char    ordableruseqty      [  16];    // [long  ,   16] 재사용가능수량(매도)           StartPos 851, Length 16
    char    flctqty             [  16];    // [long  ,   16] 변동수량                       StartPos 867, Length 16
    char    secbalqtyd2         [  16];    // [long  ,   16] 잔고수량(D2)                   StartPos 883, Length 16
    char    sellableqty         [  16];    // [long  ,   16] 매도주문가능수량               StartPos 899, Length 16
    char    unercsellordqty     [  16];    // [long  ,   16] 미체결매도주문수량             StartPos 915, Length 16
    char    avrpchsprc          [  13];    // [long  ,   13] 평균매입가                     StartPos 931, Length 13
    char    pchsamt             [  16];    // [long  ,   16] 매입금액                       StartPos 944, Length 16
    char    deposit             [  16];    // [long  ,   16] 예수금                         StartPos 960, Length 16
    char    substamt            [  16];    // [long  ,   16] 대용금                         StartPos 976, Length 16
    char    csgnmnymgn          [  16];    // [long  ,   16] 위탁증거금현금                 StartPos 992, Length 16
    char    csgnsubstmgn        [  16];    // [long  ,   16] 위탁증거금대용                 StartPos 1008, Length 16
    char    crdtpldgruseamt     [  16];    // [long  ,   16] 신용담보재사용금               StartPos 1024, Length 16
    char    ordablemny          [  16];    // [long  ,   16] 주문가능현금                   StartPos 1040, Length 16
    char    ordablesubstamt     [  16];    // [long  ,   16] 주문가능대용                   StartPos 1056, Length 16
    char    ruseableamt         [  16];    // [long  ,   16] 재사용가능금액                 StartPos 1072, Length 16
} SC0_OutBlock;

//------------------------------------------------------------------------------
// 현물 주문 체결 실시간 정보 (SC1)
//------------------------------------------------------------------------------
typedef struct {
    char    lineseq             [  10];    // [long  ,   10] 라인일련번호                   StartPos 0, Length 10
    char    accno               [  11];    // [string,   11] 계좌번호                       StartPos 10, Length 11
    char    user                [   8];    // [string,    8] 조작자ID                       StartPos 21, Length 8
    char    len                 [   6];    // [long  ,    6] 헤더길이                       StartPos 29, Length 6
    char    gubun               [   1];    // [string,    1] 헤더구분                       StartPos 35, Length 1
    char    compress            [   1];    // [string,    1] 압축구분                       StartPos 36, Length 1
    char    encrypt             [   1];    // [string,    1] 암호구분                       StartPos 37, Length 1
    char    offset              [   3];    // [long  ,    3] 공통시작지점                   StartPos 38, Length 3
    char    trcode              [   8];    // [string,    8] TRCODE                         StartPos 41, Length 8
    char    compid              [   3];    // [string,    3] 이용사번호                     StartPos 49, Length 3
    char    userid              [  16];    // [string,   16] 사용자ID                       StartPos 52, Length 16
    char    media               [   2];    // [string,    2] 접속매체                       StartPos 68, Length 2
    char    ifid                [   3];    // [string,    3] I/F일련번호                    StartPos 70, Length 3
    char    seq                 [   9];    // [string,    9] 전문일련번호                   StartPos 73, Length 9
    char    trid                [  16];    // [string,   16] TR추적ID                       StartPos 82, Length 16
    char    pubip               [  12];    // [string,   12] 공인IP                         StartPos 98, Length 12
    char    prvip               [  12];    // [string,   12] 사설IP                         StartPos 110, Length 12
    char    pcbpno              [   3];    // [string,    3] 처리지점번호                   StartPos 122, Length 3
    char    bpno                [   3];    // [string,    3] 지점번호                       StartPos 125, Length 3
    char    termno              [   8];    // [string,    8] 단말번호                       StartPos 128, Length 8
    char    lang                [   1];    // [string,    1] 언어구분                       StartPos 136, Length 1
    char    proctm              [   9];    // [long  ,    9] AP처리시간                     StartPos 137, Length 9
    char    msgcode             [   4];    // [string,    4] 메세지코드                     StartPos 146, Length 4
    char    outgu               [   1];    // [string,    1] 메세지출력구분                 StartPos 150, Length 1
    char    compreq             [   1];    // [string,    1] 압축요청구분                   StartPos 151, Length 1
    char    funckey             [   4];    // [string,    4] 기능키                         StartPos 152, Length 4
    char    reqcnt              [   4];    // [long  ,    4] 요청레코드개수                 StartPos 156, Length 4
    char    filler              [   6];    // [string,    6] 예비영역                       StartPos 160, Length 6
    char    cont                [   1];    // [string,    1] 연속구분                       StartPos 166, Length 1
    char    contkey             [  18];    // [string,   18] 연속키값                       StartPos 167, Length 18
    char    varlen              [   2];    // [long  ,    2] 가변시스템길이                 StartPos 185, Length 2
    char    varhdlen            [   2];    // [long  ,    2] 가변해더길이                   StartPos 187, Length 2
    char    varmsglen           [   2];    // [long  ,    2] 가변메시지길이                 StartPos 189, Length 2
    char    trsrc               [   1];    // [string,    1] 조회발원지                     StartPos 191, Length 1
    char    eventid             [   4];    // [string,    4] I/F이벤트ID                    StartPos 192, Length 4
    char    ifinfo              [   4];    // [string,    4] I/F정보                        StartPos 196, Length 4
    char    filler1             [  41];    // [string,   41] 예비영역                       StartPos 200, Length 41
    char    ordxctptncode       [   2];    // [string,    2] 주문체결유형코드               StartPos 241, Length 2
    char    ordmktcode          [   2];    // [string,    2] 주문시장코드                   StartPos 243, Length 2
    char    ordptncode          [   2];    // [string,    2] 주문유형코드                   StartPos 245, Length 2
    char    mgmtbrnno           [   3];    // [string,    3] 관리지점번호                   StartPos 247, Length 3
    char    accno1              [  11];    // [string,   11] 계좌번호                       StartPos 250, Length 11
    char    accno2              [   9];    // [string,    9] 계좌번호                       StartPos 261, Length 9
    char    acntnm              [  40];    // [string,   40] 계좌명                         StartPos 270, Length 40
    char    Isuno               [  12];    // [string,   12] 종목번호                       StartPos 310, Length 12
    char    Isunm               [  40];    // [string,   40] 종목명                         StartPos 322, Length 40
    char    ordno               [  10];    // [long  ,   10] 주문번호                       StartPos 362, Length 10
    char    orgordno            [  10];    // [long  ,   10] 원주문번호                     StartPos 372, Length 10
    char    execno              [  10];    // [long  ,   10] 체결번호                       StartPos 382, Length 10
    char    ordqty              [  16];    // [long  ,   16] 주문수량                       StartPos 392, Length 16
    char    ordprc              [  13];    // [long  ,   13] 주문가격                       StartPos 408, Length 13
    char    execqty             [  16];    // [long  ,   16] 체결수량                       StartPos 421, Length 16
    char    execprc             [  13];    // [long  ,   13] 체결가격                       StartPos 437, Length 13
    char    mdfycnfqty          [  16];    // [long  ,   16] 정정확인수량                   StartPos 450, Length 16
    char    mdfycnfprc          [  16];    // [long  ,   16] 정정확인가격                   StartPos 466, Length 16
    char    canccnfqty          [  16];    // [long  ,   16] 취소확인수량                   StartPos 482, Length 16
    char    rjtqty              [  16];    // [long  ,   16] 거부수량                       StartPos 498, Length 16
    char    ordtrxptncode       [   4];    // [long  ,    4] 주문처리유형코드               StartPos 514, Length 4
    char    mtiordseqno         [  10];    // [long  ,   10] 복수주문일련번호               StartPos 518, Length 10
    char    ordcndi             [   1];    // [string,    1] 주문조건                       StartPos 528, Length 1
    char    ordprcptncode       [   2];    // [string,    2] 호가유형코드                   StartPos 529, Length 2
    char    nsavtrdqty          [  16];    // [long  ,   16] 비저축체결수량                 StartPos 531, Length 16
    char    shtnIsuno           [   9];    // [string,    9] 단축종목번호                   StartPos 547, Length 9
    char    opdrtnno            [  12];    // [string,   12] 운용지시번호                   StartPos 556, Length 12
    char    cvrgordtp           [   1];    // [string,    1] 반대매매주문구분               StartPos 568, Length 1
    char    unercqty            [  16];    // [long  ,   16] 미체결수량(주문)               StartPos 569, Length 16
    char    orgordunercqty      [  16];    // [long  ,   16] 원주문미체결수량               StartPos 585, Length 16
    char    orgordmdfyqty       [  16];    // [long  ,   16] 원주문정정수량                 StartPos 601, Length 16
    char    orgordcancqty       [  16];    // [long  ,   16] 원주문취소수량                 StartPos 617, Length 16
    char    ordavrexecprc       [  13];    // [long  ,   13] 주문평균체결가격               StartPos 633, Length 13
    char    ordamt              [  16];    // [long  ,   16] 주문금액                       StartPos 646, Length 16
    char    stdIsuno            [  12];    // [string,   12] 표준종목번호                   StartPos 662, Length 12
    char    bfstdIsuno          [  12];    // [string,   12] 전표준종목번호                 StartPos 674, Length 12
    char    bnstp               [   1];    // [string,    1] 매매구분                       StartPos 686, Length 1
    char    ordtrdptncode       [   2];    // [string,    2] 주문거래유형코드               StartPos 687, Length 2
    char    mgntrncode          [   3];    // [string,    3] 신용거래코드                   StartPos 689, Length 3
    char    adduptp             [   2];    // [string,    2] 수수료합산코드                 StartPos 692, Length 2
    char    commdacode          [   2];    // [string,    2] 통신매체코드                   StartPos 694, Length 2
    char    Loandt              [   8];    // [string,    8] 대출일                         StartPos 696, Length 8
    char    mbrnmbrno           [   3];    // [long  ,    3] 회원/비회원사번호              StartPos 704, Length 3
    char    ordacntno           [  20];    // [string,   20] 주문계좌번호                   StartPos 707, Length 20
    char    agrgbrnno           [   3];    // [string,    3] 집계지점번호                   StartPos 727, Length 3
    char    mgempno             [   9];    // [string,    9] 관리사원번호                   StartPos 730, Length 9
    char    futsLnkbrnno        [   3];    // [string,    3] 선물연계지점번호               StartPos 739, Length 3
    char    futsLnkacntno       [  20];    // [string,   20] 선물연계계좌번호               StartPos 742, Length 20
    char    futsmkttp           [   1];    // [string,    1] 선물시장구분                   StartPos 762, Length 1
    char    regmktcode          [   2];    // [string,    2] 등록시장코드                   StartPos 763, Length 2
    char    mnymgnrat           [   7];    // [long  ,    7] 현금증거금률                   StartPos 765, Length 7
    char    substmgnrat         [   9];    // [long  ,    9] 대용증거금률                   StartPos 772, Length 9
    char    mnyexecamt          [  16];    // [long  ,   16] 현금체결금액                   StartPos 781, Length 16
    char    ubstexecamt         [  16];    // [long  ,   16] 대용체결금액                   StartPos 797, Length 16
    char    cmsnamtexecamt      [  16];    // [long  ,   16] 수수료체결금액                 StartPos 813, Length 16
    char    crdtpldgexecamt     [  16];    // [long  ,   16] 신용담보체결금액               StartPos 829, Length 16
    char    crdtexecamt         [  16];    // [long  ,   16] 신용체결금액                   StartPos 845, Length 16
    char    prdayruseexecval    [  16];    // [long  ,   16] 전일재사용체결금액             StartPos 861, Length 16
    char    crdayruseexecval    [  16];    // [long  ,   16] 금일재사용체결금액             StartPos 877, Length 16
    char    spotexecqty         [  16];    // [long  ,   16] 실물체결수량                   StartPos 893, Length 16
    char    stslexecqty         [  16];    // [long  ,   16] 공매도체결수량                 StartPos 909, Length 16
    char    strtgcode           [   6];    // [string,    6] 전략코드                       StartPos 925, Length 6
    char    grpId               [  20];    // [string,   20] 그룹Id                         StartPos 931, Length 20
    char    ordseqno            [  10];    // [long  ,   10] 주문회차                       StartPos 951, Length 10
    char    ptflno              [  10];    // [long  ,   10] 포트폴리오번호                 StartPos 961, Length 10
    char    bskno               [  10];    // [long  ,   10] 바스켓번호                     StartPos 971, Length 10
    char    trchno              [  10];    // [long  ,   10] 트렌치번호                     StartPos 981, Length 10
    char    itemno              [  10];    // [long  ,   10] 아이템번호                     StartPos 991, Length 10
    char    orduserId           [  16];    // [string,   16] 주문자Id                       StartPos 1001, Length 16
    char    brwmgmtYn           [   1];    // [long  ,    1] 차입관리여부                   StartPos 1017, Length 1
    char    frgrunqno           [   6];    // [string,    6] 외국인고유번호                 StartPos 1018, Length 6
    char    trtzxLevytp         [   1];    // [string,    1] 거래세징수구분                 StartPos 1024, Length 1
    char    lptp                [   1];    // [string,    1] 유동성공급자구분               StartPos 1025, Length 1
    char    exectime            [   9];    // [string,    9] 체결시각                       StartPos 1026, Length 9
    char    rcptexectime        [   9];    // [string,    9] 거래소수신체결시각             StartPos 1035, Length 9
    char    rmndLoanamt         [  16];    // [long  ,   16] 잔여대출금액                   StartPos 1044, Length 16
    char    secbalqty           [  16];    // [long  ,   16] 잔고수량                       StartPos 1060, Length 16
    char    spotordableqty      [  16];    // [long  ,   16] 실물가능수량                   StartPos 1076, Length 16
    char    ordableruseqty      [  16];    // [long  ,   16] 재사용가능수량(매도)           StartPos 1092, Length 16
    char    flctqty             [  16];    // [long  ,   16] 변동수량                       StartPos 1108, Length 16
    char    secbalqtyd2         [  16];    // [long  ,   16] 잔고수량(d2)                   StartPos 1124, Length 16
    char    sellableqty         [  16];    // [long  ,   16] 매도주문가능수량               StartPos 1140, Length 16
    char    unercsellordqty     [  16];    // [long  ,   16] 미체결매도주문수량             StartPos 1156, Length 16
    char    avrpchsprc          [  13];    // [long  ,   13] 평균매입가                     StartPos 1172, Length 13
    char    pchsant             [  16];    // [long  ,   16] 매입금액                       StartPos 1185, Length 16
    char    deposit             [  16];    // [long  ,   16] 예수금                         StartPos 1201, Length 16
    char    substamt            [  16];    // [long  ,   16] 대용금                         StartPos 1217, Length 16
    char    csgnmnymgn          [  16];    // [long  ,   16] 위탁증거금현금                 StartPos 1233, Length 16
    char    csgnsubstmgn        [  16];    // [long  ,   16] 위탁증거금대용                 StartPos 1249, Length 16
    char    crdtpldgruseamt     [  16];    // [long  ,   16] 신용담보재사용금               StartPos 1265, Length 16
    char    ordablemny          [  16];    // [long  ,   16] 주문가능현금                   StartPos 1281, Length 16
    char    ordablesubstamt     [  16];    // [long  ,   16] 주문가능대용                   StartPos 1297, Length 16
    char    ruseableamt         [  16];    // [long  ,   16] 재사용가능금액                 StartPos 1313, Length 16
} SC1_OutBlock;

//------------------------------------------------------------------------------
// 현물 주문 정정 실시간 정보 (SC2)
//------------------------------------------------------------------------------
typedef struct _SC2_OutBlock
{
    char    lineseq             [  10];    // [long  ,   10] 라인일련번호                   StartPos 0, Length 10
    char    accno               [  11];    // [string,   11] 계좌번호                       StartPos 10, Length 11
    char    user                [   8];    // [string,    8] 조작자ID                       StartPos 21, Length 8
    char    len                 [   6];    // [long  ,    6] 헤더길이                       StartPos 29, Length 6
    char    gubun               [   1];    // [string,    1] 헤더구분                       StartPos 35, Length 1
    char    compress            [   1];    // [string,    1] 압축구분                       StartPos 36, Length 1
    char    encrypt             [   1];    // [string,    1] 암호구분                       StartPos 37, Length 1
    char    offset              [   3];    // [long  ,    3] 공통시작지점                   StartPos 38, Length 3
    char    trcode              [   8];    // [string,    8] TRCODE                         StartPos 41, Length 8
    char    compid              [   3];    // [string,    3] 이용사번호                     StartPos 49, Length 3
    char    userid              [  16];    // [string,   16] 사용자ID                       StartPos 52, Length 16
    char    media               [   2];    // [string,    2] 접속매체                       StartPos 68, Length 2
    char    ifid                [   3];    // [string,    3] I/F일련번호                    StartPos 70, Length 3
    char    seq                 [   9];    // [string,    9] 전문일련번호                   StartPos 73, Length 9
    char    trid                [  16];    // [string,   16] TR추적ID                       StartPos 82, Length 16
    char    pubip               [  12];    // [string,   12] 공인IP                         StartPos 98, Length 12
    char    prvip               [  12];    // [string,   12] 사설IP                         StartPos 110, Length 12
    char    pcbpno              [   3];    // [string,    3] 처리지점번호                   StartPos 122, Length 3
    char    bpno                [   3];    // [string,    3] 지점번호                       StartPos 125, Length 3
    char    termno              [   8];    // [string,    8] 단말번호                       StartPos 128, Length 8
    char    lang                [   1];    // [string,    1] 언어구분                       StartPos 136, Length 1
    char    proctm              [   9];    // [long  ,    9] AP처리시간                     StartPos 137, Length 9
    char    msgcode             [   4];    // [string,    4] 메세지코드                     StartPos 146, Length 4
    char    outgu               [   1];    // [string,    1] 메세지출력구분                 StartPos 150, Length 1
    char    compreq             [   1];    // [string,    1] 압축요청구분                   StartPos 151, Length 1
    char    funckey             [   4];    // [string,    4] 기능키                         StartPos 152, Length 4
    char    reqcnt              [   4];    // [long  ,    4] 요청레코드개수                 StartPos 156, Length 4
    char    filler              [   6];    // [string,    6] 예비영역                       StartPos 160, Length 6
    char    cont                [   1];    // [string,    1] 연속구분                       StartPos 166, Length 1
    char    contkey             [  18];    // [string,   18] 연속키값                       StartPos 167, Length 18
    char    varlen              [   2];    // [long  ,    2] 가변시스템길이                 StartPos 185, Length 2
    char    varhdlen            [   2];    // [long  ,    2] 가변해더길이                   StartPos 187, Length 2
    char    varmsglen           [   2];    // [long  ,    2] 가변메시지길이                 StartPos 189, Length 2
    char    trsrc               [   1];    // [string,    1] 조회발원지                     StartPos 191, Length 1
    char    eventid             [   4];    // [string,    4] I/F이벤트ID                    StartPos 192, Length 4
    char    ifinfo              [   4];    // [string,    4] I/F정보                        StartPos 196, Length 4
    char    filler1             [  41];    // [string,   41] 예비영역                       StartPos 200, Length 41
    char    ordxctptncode       [   2];    // [string,    2] 주문체결유형코드               StartPos 241, Length 2
    char    ordmktcode          [   2];    // [string,    2] 주문시장코드                   StartPos 243, Length 2
    char    ordptncode          [   2];    // [string,    2] 주문유형코드                   StartPos 245, Length 2
    char    mgmtbrnno           [   3];    // [string,    3] 관리지점번호                   StartPos 247, Length 3
    char    accno1              [  11];    // [string,   11] 계좌번호                       StartPos 250, Length 11
    char    accno2              [   9];    // [string,    9] 계좌번호                       StartPos 261, Length 9
    char    acntnm              [  40];    // [string,   40] 계좌명                         StartPos 270, Length 40
    char    Isuno               [  12];    // [string,   12] 종목번호                       StartPos 310, Length 12
    char    Isunm               [  40];    // [string,   40] 종목명                         StartPos 322, Length 40
    char    ordno               [  10];    // [long  ,   10] 주문번호                       StartPos 362, Length 10
    char    orgordno            [  10];    // [long  ,   10] 원주문번호                     StartPos 372, Length 10
    char    execno              [  10];    // [long  ,   10] 체결번호                       StartPos 382, Length 10
    char    ordqty              [  16];    // [long  ,   16] 주문수량                       StartPos 392, Length 16
    char    ordprc              [  13];    // [long  ,   13] 주문가격                       StartPos 408, Length 13
    char    execqty             [  16];    // [long  ,   16] 체결수량                       StartPos 421, Length 16
    char    execprc             [  13];    // [long  ,   13] 체결가격                       StartPos 437, Length 13
    char    mdfycnfqty          [  16];    // [long  ,   16] 정정확인수량                   StartPos 450, Length 16
    char    mdfycnfprc          [  16];    // [long  ,   16] 정정확인가격                   StartPos 466, Length 16
    char    canccnfqty          [  16];    // [long  ,   16] 취소확인수량                   StartPos 482, Length 16
    char    rjtqty              [  16];    // [long  ,   16] 거부수량                       StartPos 498, Length 16
    char    ordtrxptncode       [   4];    // [long  ,    4] 주문처리유형코드               StartPos 514, Length 4
    char    mtiordseqno         [  10];    // [long  ,   10] 복수주문일련번호               StartPos 518, Length 10
    char    ordcndi             [   1];    // [string,    1] 주문조건                       StartPos 528, Length 1
    char    ordprcptncode       [   2];    // [string,    2] 호가유형코드                   StartPos 529, Length 2
    char    nsavtrdqty          [  16];    // [long  ,   16] 비저축체결수량                 StartPos 531, Length 16
    char    shtnIsuno           [   9];    // [string,    9] 단축종목번호                   StartPos 547, Length 9
    char    opdrtnno            [  12];    // [string,   12] 운용지시번호                   StartPos 556, Length 12
    char    cvrgordtp           [   1];    // [string,    1] 반대매매주문구분               StartPos 568, Length 1
    char    unercqty            [  16];    // [long  ,   16] 미체결수량(주문)               StartPos 569, Length 16
    char    orgordunercqty      [  16];    // [long  ,   16] 원주문미체결수량               StartPos 585, Length 16
    char    orgordmdfyqty       [  16];    // [long  ,   16] 원주문정정수량                 StartPos 601, Length 16
    char    orgordcancqty       [  16];    // [long  ,   16] 원주문취소수량                 StartPos 617, Length 16
    char    ordavrexecprc       [  13];    // [long  ,   13] 주문평균체결가격               StartPos 633, Length 13
    char    ordamt              [  16];    // [long  ,   16] 주문금액                       StartPos 646, Length 16
    char    stdIsuno            [  12];    // [string,   12] 표준종목번호                   StartPos 662, Length 12
    char    bfstdIsuno          [  12];    // [string,   12] 전표준종목번호                 StartPos 674, Length 12
    char    bnstp               [   1];    // [string,    1] 매매구분                       StartPos 686, Length 1
    char    ordtrdptncode       [   2];    // [string,    2] 주문거래유형코드               StartPos 687, Length 2
    char    mgntrncode          [   3];    // [string,    3] 신용거래코드                   StartPos 689, Length 3
    char    adduptp             [   2];    // [string,    2] 수수료합산코드                 StartPos 692, Length 2
    char    commdacode          [   2];    // [string,    2] 통신매체코드                   StartPos 694, Length 2
    char    Loandt              [   8];    // [string,    8] 대출일                         StartPos 696, Length 8
    char    mbrnmbrno           [   3];    // [long  ,    3] 회원/비회원사번호              StartPos 704, Length 3
    char    ordacntno           [  20];    // [string,   20] 주문계좌번호                   StartPos 707, Length 20
    char    agrgbrnno           [   3];    // [string,    3] 집계지점번호                   StartPos 727, Length 3
    char    mgempno             [   9];    // [string,    9] 관리사원번호                   StartPos 730, Length 9
    char    futsLnkbrnno        [   3];    // [string,    3] 선물연계지점번호               StartPos 739, Length 3
    char    futsLnkacntno       [  20];    // [string,   20] 선물연계계좌번호               StartPos 742, Length 20
    char    futsmkttp           [   1];    // [string,    1] 선물시장구분                   StartPos 762, Length 1
    char    regmktcode          [   2];    // [string,    2] 등록시장코드                   StartPos 763, Length 2
    char    mnymgnrat           [   7];    // [long  ,    7] 현금증거금률                   StartPos 765, Length 7
    char    substmgnrat         [   9];    // [long  ,    9] 대용증거금률                   StartPos 772, Length 9
    char    mnyexecamt          [  16];    // [long  ,   16] 현금체결금액                   StartPos 781, Length 16
    char    ubstexecamt         [  16];    // [long  ,   16] 대용체결금액                   StartPos 797, Length 16
    char    cmsnamtexecamt      [  16];    // [long  ,   16] 수수료체결금액                 StartPos 813, Length 16
    char    crdtpldgexecamt     [  16];    // [long  ,   16] 신용담보체결금액               StartPos 829, Length 16
    char    crdtexecamt         [  16];    // [long  ,   16] 신용체결금액                   StartPos 845, Length 16
    char    prdayruseexecval    [  16];    // [long  ,   16] 전일재사용체결금액             StartPos 861, Length 16
    char    crdayruseexecval    [  16];    // [long  ,   16] 금일재사용체결금액             StartPos 877, Length 16
    char    spotexecqty         [  16];    // [long  ,   16] 실물체결수량                   StartPos 893, Length 16
    char    stslexecqty         [  16];    // [long  ,   16] 공매도체결수량                 StartPos 909, Length 16
    char    strtgcode           [   6];    // [string,    6] 전략코드                       StartPos 925, Length 6
    char    grpId               [  20];    // [string,   20] 그룹Id                         StartPos 931, Length 20
    char    ordseqno            [  10];    // [long  ,   10] 주문회차                       StartPos 951, Length 10
    char    ptflno              [  10];    // [long  ,   10] 포트폴리오번호                 StartPos 961, Length 10
    char    bskno               [  10];    // [long  ,   10] 바스켓번호                     StartPos 971, Length 10
    char    trchno              [  10];    // [long  ,   10] 트렌치번호                     StartPos 981, Length 10
    char    itemno              [  10];    // [long  ,   10] 아이템번호                     StartPos 991, Length 10
    char    orduserId           [  16];    // [string,   16] 주문자Id                       StartPos 1001, Length 16
    char    brwmgmtYn           [   1];    // [long  ,    1] 차입관리여부                   StartPos 1017, Length 1
    char    frgrunqno           [   6];    // [string,    6] 외국인고유번호                 StartPos 1018, Length 6
    char    trtzxLevytp         [   1];    // [string,    1] 거래세징수구분                 StartPos 1024, Length 1
    char    lptp                [   1];    // [string,    1] 유동성공급자구분               StartPos 1025, Length 1
    char    exectime            [   9];    // [string,    9] 체결시각                       StartPos 1026, Length 9
    char    rcptexectime        [   9];    // [string,    9] 거래소수신체결시각             StartPos 1035, Length 9
    char    rmndLoanamt         [  16];    // [long  ,   16] 잔여대출금액                   StartPos 1044, Length 16
    char    secbalqty           [  16];    // [long  ,   16] 잔고수량                       StartPos 1060, Length 16
    char    spotordableqty      [  16];    // [long  ,   16] 실물가능수량                   StartPos 1076, Length 16
    char    ordableruseqty      [  16];    // [long  ,   16] 재사용가능수량(매도)           StartPos 1092, Length 16
    char    flctqty             [  16];    // [long  ,   16] 변동수량                       StartPos 1108, Length 16
    char    secbalqtyd2         [  16];    // [long  ,   16] 잔고수량(d2)                   StartPos 1124, Length 16
    char    sellableqty         [  16];    // [long  ,   16] 매도주문가능수량               StartPos 1140, Length 16
    char    unercsellordqty     [  16];    // [long  ,   16] 미체결매도주문수량             StartPos 1156, Length 16
    char    avrpchsprc          [  13];    // [long  ,   13] 평균매입가                     StartPos 1172, Length 13
    char    pchsant             [  16];    // [long  ,   16] 매입금액                       StartPos 1185, Length 16
    char    deposit             [  16];    // [long  ,   16] 예수금                         StartPos 1201, Length 16
    char    substamt            [  16];    // [long  ,   16] 대용금                         StartPos 1217, Length 16
    char    csgnmnymgn          [  16];    // [long  ,   16] 위탁증거금현금                 StartPos 1233, Length 16
    char    csgnsubstmgn        [  16];    // [long  ,   16] 위탁증거금대용                 StartPos 1249, Length 16
    char    crdtpldgruseamt     [  16];    // [long  ,   16] 신용담보재사용금               StartPos 1265, Length 16
    char    ordablemny          [  16];    // [long  ,   16] 주문가능현금                   StartPos 1281, Length 16
    char    ordablesubstamt     [  16];    // [long  ,   16] 주문가능대용                   StartPos 1297, Length 16
    char    ruseableamt         [  16];    // [long  ,   16] 재사용가능금액                 StartPos 1313, Length 16
} SC2_OutBlock;

//------------------------------------------------------------------------------
// 현물 주문 취소 실시간 정보 (SC3)
//------------------------------------------------------------------------------
typedef struct {
    char    lineseq[10];    //[long  ,   10] 라인일련번호   StartPos 0, Length 10
    char    accno[11];    //[string,   11] 계좌번호   StartPos 10, Length 11
    char    user[8];    //[string,    8] 조작자ID   StartPos 21, Length 8
    char    len[6];    //[long  ,    6] 헤더길이   StartPos 29, Length 6
    char    gubun[1];    //[string,    1] 헤더구분   StartPos 35, Length 1
    char    compress[1];    //[string,    1] 압축구분   StartPos 36, Length 1
    char    encrypt[1];    //[string,    1] 암호구분   StartPos 37, Length 1
    char    offset[3];    //[long  ,    3] 공통시작지점   StartPos 38, Length 3
    char    trcode[8];    //[string,    8] TRCODE   StartPos 41, Length 8
    char    compid[3];    //[string,    3] 이용사번호   StartPos 49, Length 3
    char    userid[16];    //[string,   16] 사용자ID   StartPos 52, Length 16
    char    media[2];    //[string,    2] 접속매체   StartPos 68, Length 2
    char    ifid[3];    //[string,    3] I/F일련번호   StartPos 70, Length 3
    char    seq[9];    //[string,    9] 전문일련번호   StartPos 73, Length 9
    char    trid[16];    //[string,   16] TR추적ID   StartPos 82, Length 16
    char    pubip[12];    //[string,   12] 공인IP   StartPos 98, Length 12
    char    prvip[12];    //[string,   12] 사설IP   StartPos 110, Length 12
    char    pcbpno[3];    //[string,    3] 처리지점번호   StartPos 122, Length 3
    char    bpno[3];    //[string,    3] 지점번호   StartPos 125, Length 3
    char    termno[8];    //[string,    8] 단말번호   StartPos 128, Length 8
    char    lang[1];    //[string,    1] 언어구분   StartPos 136, Length 1
    char    proctm[9];    //[long  ,    9] AP처리시간   StartPos 137, Length 9
    char    msgcode[4];    //[string,    4] 메세지코드   StartPos 146, Length 4
    char    outgu[1];    //[string,    1] 메세지출력구분   StartPos 150, Length 1
    char    compreq[1];    //[string,    1] 압축요청구분   StartPos 151, Length 1
    char    funckey[4];    //[string,    4] 기능키   StartPos 152, Length 4
    char    reqcnt[4];    //[long  ,    4] 요청레코드개수   StartPos 156, Length 4
    char    filler[6];    //[string,    6] 예비영역   StartPos 160, Length 6
    char    cont[1];    //[string,    1] 연속구분   StartPos 166, Length 1
    char    contkey[18];    //[string,   18] 연속키값   StartPos 167, Length 18
    char    varlen[2];    //[long  ,    2] 가변시스템길이   StartPos 185, Length 2
    char    varhdlen[2];    //[long  ,    2] 가변해더길이   StartPos 187, Length 2
    char    varmsglen[2];    //[long  ,    2] 가변메시지길이   StartPos 189, Length 2
    char    trsrc[1];    //[string,    1] 조회발원지   StartPos 191, Length 1
    char    eventid[4];    //[string,    4] I/F이벤트ID   StartPos 192, Length 4
    char    ifinfo[4];    //[string,    4] I/F정보   StartPos 196, Length 4
    char    filler1[41];    //[string,   41] 예비영역   StartPos 200, Length 41
    char    ordxctptncode[2];    //[string,    2] 주문체결유형코드   StartPos 241, Length 2
    char    ordmktcode[2];    //[string,    2] 주문시장코드   StartPos 243, Length 2
    char    ordptncode[2];    //[string,    2] 주문유형코드   StartPos 245, Length 2
    char    mgmtbrnno[3];    //[string,    3] 관리지점번호   StartPos 247, Length 3
    char    accno1[11];    //[string,   11] 계좌번호   StartPos 250, Length 11
    char    accno2[9];    //[string,    9] 계좌번호   StartPos 261, Length 9
    char    acntnm[40];    //[string,   40] 계좌명   StartPos 270, Length 40
    char    Isuno[12];    //[string,   12] 종목번호   StartPos 310, Length 12
    char    Isunm[40];    //[string,   40] 종목명   StartPos 322, Length 40
    char    ordno[10];    //[long  ,   10] 주문번호   StartPos 362, Length 10
    char    orgordno[10];    //[long  ,   10] 원주문번호   StartPos 372, Length 10
    char    execno[10];    //[long  ,   10] 체결번호   StartPos 382, Length 10
    char    ordqty[16];    //[long  ,   16] 주문수량   StartPos 392, Length 16
    char    ordprc[13];    //[long  ,   13] 주문가격   StartPos 408, Length 13
    char    execqty[16];    //[long  ,   16] 체결수량   StartPos 421, Length 16
    char    execprc[13];    //[long  ,   13] 체결가격   StartPos 437, Length 13
    char    mdfycnfqty[16];    //[long  ,   16] 정정확인수량   StartPos 450, Length 16
    char    mdfycnfprc[16];    //[long  ,   16] 정정확인가격   StartPos 466, Length 16
    char    canccnfqty[16];    //[long  ,   16] 취소확인수량   StartPos 482, Length 16
    char    rjtqty[16];    //[long  ,   16] 거부수량   StartPos 498, Length 16
    char    ordtrxptncode[4];    //[long  ,    4] 주문처리유형코드   StartPos 514, Length 4
    char    mtiordseqno[10];    //[long  ,   10] 복수주문일련번호   StartPos 518, Length 10
    char    ordcndi[1];    //[string,    1] 주문조건   StartPos 528, Length 1
    char    ordprcptncode[2];    //[string,    2] 호가유형코드   StartPos 529, Length 2
    char    nsavtrdqty[16];    //[long  ,   16] 비저축체결수량   StartPos 531, Length 16
    char    shtnIsuno[9];    //[string,    9] 단축종목번호   StartPos 547, Length 9
    char    opdrtnno[12];    //[string,   12] 운용지시번호   StartPos 556, Length 12
    char    cvrgordtp[1];    //[string,    1] 반대매매주문구분   StartPos 568, Length 1
    char    unercqty[16];    //[long  ,   16] 미체결수량(주문)   StartPos 569, Length 16
    char    orgordunercqty[16];    //[long  ,   16] 원주문미체결수량   StartPos 585, Length 16
    char    orgordmdfyqty[16];    //[long  ,   16] 원주문정정수량   StartPos 601, Length 16
    char    orgordcancqty[16];    //[long  ,   16] 원주문취소수량   StartPos 617, Length 16
    char    ordavrexecprc[13];    //[long  ,   13] 주문평균체결가격   StartPos 633, Length 13
    char    ordamt[16];    //[long  ,   16] 주문금액   StartPos 646, Length 16
    char    stdIsuno[12];    //[string,   12] 표준종목번호   StartPos 662, Length 12
    char    bfstdIsuno[12];    //[string,   12] 전표준종목번호   StartPos 674, Length 12
    char    bnstp[1];    //[string,    1] 매매구분   StartPos 686, Length 1
    char    ordtrdptncode[2];    //[string,    2] 주문거래유형코드   StartPos 687, Length 2
    char    mgntrncode[3];    //[string,    3] 신용거래코드   StartPos 689, Length 3
    char    adduptp[2];    //[string,    2] 수수료합산코드   StartPos 692, Length 2
    char    commdacode[2];    //[string,    2] 통신매체코드   StartPos 694, Length 2
    char    Loandt[8];    //[string,    8] 대출일   StartPos 696, Length 8
    char    mbrnmbrno[3];    //[long  ,    3] 회원/비회원사번호   StartPos 704, Length 3
    char    ordacntno[20];    //[string,   20] 주문계좌번호   StartPos 707, Length 20
    char    agrgbrnno[3];    //[string,    3] 집계지점번호   StartPos 727, Length 3
    char    mgempno[9];    //[string,    9] 관리사원번호   StartPos 730, Length 9
    char    futsLnkbrnno[3];    //[string,    3] 선물연계지점번호   StartPos 739, Length 3
    char    futsLnkacntno[20];    //[string,   20] 선물연계계좌번호   StartPos 742, Length 20
    char    futsmkttp[1];    //[string,    1] 선물시장구분   StartPos 762, Length 1
    char    regmktcode[2];    //[string,    2] 등록시장코드   StartPos 763, Length 2
    char    mnymgnrat[7];    //[long  ,    7] 현금증거금률   StartPos 765, Length 7
    char    substmgnrat[9];    //[long  ,    9] 대용증거금률   StartPos 772, Length 9
    char    mnyexecamt[16];    //[long  ,   16] 현금체결금액   StartPos 781, Length 16
    char    ubstexecamt[16];    //[long  ,   16] 대용체결금액   StartPos 797, Length 16
    char    cmsnamtexecamt[16];    //[long  ,   16] 수수료체결금액   StartPos 813, Length 16
    char    crdtpldgexecamt[16];    //[long  ,   16] 신용담보체결금액   StartPos 829, Length 16
    char    crdtexecamt[16];    //[long  ,   16] 신용체결금액   StartPos 845, Length 16
    char    prdayruseexecval[16];    //[long  ,   16] 전일재사용체결금액   StartPos 861, Length 16
    char    crdayruseexecval[16];    //[long  ,   16] 금일재사용체결금액   StartPos 877, Length 16
    char    spotexecqty[16];    //[long  ,   16] 실물체결수량   StartPos 893, Length 16
    char    stslexecqty[16];    //[long  ,   16] 공매도체결수량   StartPos 909, Length 16
    char    strtgcode[6];    //[string,    6] 전략코드   StartPos 925, Length 6
    char    grpId[20];    //[string,   20] 그룹Id   StartPos 931, Length 20
    char    ordseqno[10];    //[long  ,   10] 주문회차   StartPos 951, Length 10
    char    ptflno[10];    //[long  ,   10] 포트폴리오번호   StartPos 961, Length 10
    char    bskno[10];    //[long  ,   10] 바스켓번호   StartPos 971, Length 10
    char    trchno[10];    //[long  ,   10] 트렌치번호   StartPos 981, Length 10
    char    itemno[10];    //[long  ,   10] 아이템번호   StartPos 991, Length 10
    char    orduserId[16];    //[string,   16] 주문자Id   StartPos 1001, Length 16
    char    brwmgmtYn[1];    //[long  ,    1] 차입관리여부   StartPos 1017, Length 1
    char    frgrunqno[6];    //[string,    6] 외국인고유번호   StartPos 1018, Length 6
    char    trtzxLevytp[1];    //[string,    1] 거래세징수구분   StartPos 1024, Length 1
    char    lptp[1];    //[string,    1] 유동성공급자구분   StartPos 1025, Length 1
    char    exectime[9];    //[string,    9] 체결시각   StartPos 1026, Length 9
    char    rcptexectime[9];    //[string,    9] 거래소수신체결시각   StartPos 1035, Length 9
    char    dummy_rmndLoanamt[16];    //[long  ,   16] 잔여대출금액   StartPos 1044, Length 16
    char    dummy_secbalqty[16];    //[long  ,   16] 잔고수량   StartPos 1060, Length 16
    char    dummy_spotordableqty[16];    //[long  ,   16] 실물가능수량   StartPos 1076, Length 16
    char    dummy_ordableruseqty[16];    //[long  ,   16] 재사용가능수량(매도)   StartPos 1092, Length 16
    char    flctqty[16];    //[long  ,   16] 변동수량   StartPos 1108, Length 16
    char    dummy_secbalqtyd2[16];    //[long  ,   16] 잔고수량(d2)   StartPos 1124, Length 16
    char    dummy_sellableqty[16];    //[long  ,   16] 매도주문가능수량   StartPos 1140, Length 16
    char    dummy_unercsellordqty[16];    //[long  ,   16] 미체결매도주문수량   StartPos 1156, Length 16
    char    dummy_avrpchsprc[13];    //[long  ,   13] 평균매입가   StartPos 1172, Length 13
    char    dummy_pchsant[16];    //[long  ,   16] 매입금액   StartPos 1185, Length 16
    char    deposit[16];    //[long  ,   16] 예수금   StartPos 1201, Length 16
    char    substamt[16];    //[long  ,   16] 대용금   StartPos 1217, Length 16
    char    csgnmnymgn[16];    //[long  ,   16] 위탁증거금현금   StartPos 1233, Length 16
    char    csgnsubstmgn[16];    //[long  ,   16] 위탁증거금대용   StartPos 1249, Length 16
    char    crdtpldgruseamt[16];    //[long  ,   16] 신용담보재사용금   StartPos 1265, Length 16
    char    ordablemny[16];    //[long  ,   16] 주문가능현금   StartPos 1281, Length 16
    char    ordablesubstamt[16];    //[long  ,   16] 주문가능대용   StartPos 1297, Length 16
    char    ruseableamt[16];    //[long  ,   16] 재사용가능금액   StartPos 1313, Length 16
} SC3_OutBlock;

//------------------------------------------------------------------------------
// 현물 주문 거부 실시간 정보 (SC4)
//------------------------------------------------------------------------------
typedef struct {
    char    lineseq[10];    //[long  ,   10] 라인일련번호   StartPos 0, Length 10
    char    accno[11];    //[string,   11] 계좌번호   StartPos 10, Length 11
    char    user[8];    //[string,    8] 조작자ID   StartPos 21, Length 8
    char    len[6];    //[long  ,    6] 헤더길이   StartPos 29, Length 6
    char    gubun[1];    //[string,    1] 헤더구분   StartPos 35, Length 1
    char    compress[1];    //[string,    1] 압축구분   StartPos 36, Length 1
    char    encrypt[1];    //[string,    1] 암호구분   StartPos 37, Length 1
    char    offset[3];    //[long  ,    3] 공통시작지점   StartPos 38, Length 3
    char    trcode[8];    //[string,    8] TRCODE   StartPos 41, Length 8
    char    compid[3];    //[string,    3] 이용사번호   StartPos 49, Length 3
    char    userid[16];    //[string,   16] 사용자ID   StartPos 52, Length 16
    char    media[2];    //[string,    2] 접속매체   StartPos 68, Length 2
    char    ifid[3];    //[string,    3] I/F일련번호   StartPos 70, Length 3
    char    seq[9];    //[string,    9] 전문일련번호   StartPos 73, Length 9
    char    trid[16];    //[string,   16] TR추적ID   StartPos 82, Length 16
    char    pubip[12];    //[string,   12] 공인IP   StartPos 98, Length 12
    char    prvip[12];    //[string,   12] 사설IP   StartPos 110, Length 12
    char    pcbpno[3];    //[string,    3] 처리지점번호   StartPos 122, Length 3
    char    bpno[3];    //[string,    3] 지점번호   StartPos 125, Length 3
    char    termno[8];    //[string,    8] 단말번호   StartPos 128, Length 8
    char    lang[1];    //[string,    1] 언어구분   StartPos 136, Length 1
    char    proctm[9];    //[long  ,    9] AP처리시간   StartPos 137, Length 9
    char    msgcode[4];    //[string,    4] 메세지코드   StartPos 146, Length 4
    char    outgu[1];    //[string,    1] 메세지출력구분   StartPos 150, Length 1
    char    compreq[1];    //[string,    1] 압축요청구분   StartPos 151, Length 1
    char    funckey[4];    //[string,    4] 기능키   StartPos 152, Length 4
    char    reqcnt[4];    //[long  ,    4] 요청레코드개수   StartPos 156, Length 4
    char    filler[6];    //[string,    6] 예비영역   StartPos 160, Length 6
    char    cont[1];    //[string,    1] 연속구분   StartPos 166, Length 1
    char    contkey[18];    //[string,   18] 연속키값   StartPos 167, Length 18
    char    varlen[2];    //[long  ,    2] 가변시스템길이   StartPos 185, Length 2
    char    varhdlen[2];    //[long  ,    2] 가변해더길이   StartPos 187, Length 2
    char    varmsglen[2];    //[long  ,    2] 가변메시지길이   StartPos 189, Length 2
    char    trsrc[1];    //[string,    1] 조회발원지   StartPos 191, Length 1
    char    eventid[4];    //[string,    4] I/F이벤트ID   StartPos 192, Length 4
    char    ifinfo[4];    //[string,    4] I/F정보   StartPos 196, Length 4
    char    filler1[41];    //[string,   41] 예비영역   StartPos 200, Length 41
    char    ordxctptncode[2];    //[string,    2] 주문체결유형코드   StartPos 241, Length 2
    char    ordmktcode[2];    //[string,    2] 주문시장코드   StartPos 243, Length 2
    char    ordptncode[2];    //[string,    2] 주문유형코드   StartPos 245, Length 2
    char    mgmtbrnno[3];    //[string,    3] 관리지점번호   StartPos 247, Length 3
    char    accno1[11];    //[string,   11] 계좌번호   StartPos 250, Length 11
    char    accno2[9];    //[string,    9] 계좌번호   StartPos 261, Length 9
    char    acntnm[40];    //[string,   40] 계좌명   StartPos 270, Length 40
    char    Isuno[12];    //[string,   12] 종목번호   StartPos 310, Length 12
    char    Isunm[40];    //[string,   40] 종목명   StartPos 322, Length 40
    char    ordno[10];    //[long  ,   10] 주문번호   StartPos 362, Length 10
    char    orgordno[10];    //[long  ,   10] 원주문번호   StartPos 372, Length 10
    char    execno[10];    //[long  ,   10] 체결번호   StartPos 382, Length 10
    char    ordqty[16];    //[long  ,   16] 주문수량   StartPos 392, Length 16
    char    ordprc[13];    //[long  ,   13] 주문가격   StartPos 408, Length 13
    char    execqty[16];    //[long  ,   16] 체결수량   StartPos 421, Length 16
    char    execprc[13];    //[long  ,   13] 체결가격   StartPos 437, Length 13
    char    mdfycnfqty[16];    //[long  ,   16] 정정확인수량   StartPos 450, Length 16
    char    mdfycnfprc[16];    //[long  ,   16] 정정확인가격   StartPos 466, Length 16
    char    canccnfqty[16];    //[long  ,   16] 취소확인수량   StartPos 482, Length 16
    char    rjtqty[16];    //[long  ,   16] 거부수량   StartPos 498, Length 16
    char    ordtrxptncode[4];    //[long  ,    4] 주문처리유형코드   StartPos 514, Length 4
    char    mtiordseqno[10];    //[long  ,   10] 복수주문일련번호   StartPos 518, Length 10
    char    ordcndi[1];    //[string,    1] 주문조건   StartPos 528, Length 1
    char    ordprcptncode[2];    //[string,    2] 호가유형코드   StartPos 529, Length 2
    char    nsavtrdqty[16];    //[long  ,   16] 비저축체결수량   StartPos 531, Length 16
    char    shtnIsuno[9];    //[string,    9] 단축종목번호   StartPos 547, Length 9
    char    opdrtnno[12];    //[string,   12] 운용지시번호   StartPos 556, Length 12
    char    cvrgordtp[1];    //[string,    1] 반대매매주문구분   StartPos 568, Length 1
    char    unercqty[16];    //[long  ,   16] 미체결수량(주문)   StartPos 569, Length 16
    char    orgordunercqty[16];    //[long  ,   16] 원주문미체결수량   StartPos 585, Length 16
    char    orgordmdfyqty[16];    //[long  ,   16] 원주문정정수량   StartPos 601, Length 16
    char    orgordcancqty[16];    //[long  ,   16] 원주문취소수량   StartPos 617, Length 16
    char    ordavrexecprc[13];    //[long  ,   13] 주문평균체결가격   StartPos 633, Length 13
    char    ordamt[16];    //[long  ,   16] 주문금액   StartPos 646, Length 16
    char    stdIsuno[12];    //[string,   12] 표준종목번호   StartPos 662, Length 12
    char    bfstdIsuno[12];    //[string,   12] 전표준종목번호   StartPos 674, Length 12
    char    bnstp[1];    //[string,    1] 매매구분   StartPos 686, Length 1
    char    ordtrdptncode[2];    //[string,    2] 주문거래유형코드   StartPos 687, Length 2
    char    mgntrncode[3];    //[string,    3] 신용거래코드   StartPos 689, Length 3
    char    adduptp[2];    //[string,    2] 수수료합산코드   StartPos 692, Length 2
    char    commdacode[2];    //[string,    2] 통신매체코드   StartPos 694, Length 2
    char    Loandt[8];    //[string,    8] 대출일   StartPos 696, Length 8
    char    mbrnmbrno[3];    //[long  ,    3] 회원/비회원사번호   StartPos 704, Length 3
    char    ordacntno[20];    //[string,   20] 주문계좌번호   StartPos 707, Length 20
    char    agrgbrnno[3];    //[string,    3] 집계지점번호   StartPos 727, Length 3
    char    mgempno[9];    //[string,    9] 관리사원번호   StartPos 730, Length 9
    char    futsLnkbrnno[3];    //[string,    3] 선물연계지점번호   StartPos 739, Length 3
    char    futsLnkacntno[20];    //[string,   20] 선물연계계좌번호   StartPos 742, Length 20
    char    futsmkttp[1];    //[string,    1] 선물시장구분   StartPos 762, Length 1
    char    regmktcode[2];    //[string,    2] 등록시장코드   StartPos 763, Length 2
    char    mnymgnrat[7];    //[long  ,    7] 현금증거금률   StartPos 765, Length 7
    char    substmgnrat[9];    //[long  ,    9] 대용증거금률   StartPos 772, Length 9
    char    mnyexecamt[16];    //[long  ,   16] 현금체결금액   StartPos 781, Length 16
    char    ubstexecamt[16];    //[long  ,   16] 대용체결금액   StartPos 797, Length 16
    char    cmsnamtexecamt[16];    //[long  ,   16] 수수료체결금액   StartPos 813, Length 16
    char    crdtpldgexecamt[16];    //[long  ,   16] 신용담보체결금액   StartPos 829, Length 16
    char    crdtexecamt[16];    //[long  ,   16] 신용체결금액   StartPos 845, Length 16
    char    prdayruseexecval[16];    //[long  ,   16] 전일재사용체결금액   StartPos 861, Length 16
    char    crdayruseexecval[16];    //[long  ,   16] 금일재사용체결금액   StartPos 877, Length 16
    char    spotexecqty[16];    //[long  ,   16] 실물체결수량   StartPos 893, Length 16
    char    stslexecqty[16];    //[long  ,   16] 공매도체결수량   StartPos 909, Length 16
    char    strtgcode[6];    //[string,    6] 전략코드   StartPos 925, Length 6
    char    grpId[20];    //[string,   20] 그룹Id   StartPos 931, Length 20
    char    ordseqno[10];    //[long  ,   10] 주문회차   StartPos 951, Length 10
    char    ptflno[10];    //[long  ,   10] 포트폴리오번호   StartPos 961, Length 10
    char    bskno[10];    //[long  ,   10] 바스켓번호   StartPos 971, Length 10
    char    trchno[10];    //[long  ,   10] 트렌치번호   StartPos 981, Length 10
    char    itemno[10];    //[long  ,   10] 아이템번호   StartPos 991, Length 10
    char    orduserId[16];    //[string,   16] 주문자Id   StartPos 1001, Length 16
    char    brwmgmtYn[1];    //[long  ,    1] 차입관리여부   StartPos 1017, Length 1
    char    frgrunqno[6];    //[string,    6] 외국인고유번호   StartPos 1018, Length 6
    char    trtzxLevytp[1];    //[string,    1] 거래세징수구분   StartPos 1024, Length 1
    char    lptp[1];    //[string,    1] 유동성공급자구분   StartPos 1025, Length 1
    char    exectime[9];    //[string,    9] 체결시각   StartPos 1026, Length 9
    char    rcptexectime[9];    //[string,    9] 거래소수신체결시각   StartPos 1035, Length 9
    char    dummy_rmndLoanamt[16];    //[long  ,   16] 잔여대출금액   StartPos 1044, Length 16
    char    dummy_secbalqty[16];    //[long  ,   16] 잔고수량   StartPos 1060, Length 16
    char    dummy_spotordableqty[16];    //[long  ,   16] 실물가능수량   StartPos 1076, Length 16
    char    dummy_ordableruseqty[16];    //[long  ,   16] 재사용가능수량(매도)   StartPos 1092, Length 16
    char    flctqty[16];    //[long  ,   16] 변동수량   StartPos 1108, Length 16
    char    dummy_secbalqtyd2[16];    //[long  ,   16] 잔고수량(d2)   StartPos 1124, Length 16
    char    dummy_sellableqty[16];    //[long  ,   16] 매도주문가능수량   StartPos 1140, Length 16
    char    dummy_unercsellordqty[16];    //[long  ,   16] 미체결매도주문수량   StartPos 1156, Length 16
    char    dummy_avrpchsprc[13];    //[long  ,   13] 평균매입가   StartPos 1172, Length 13
    char    dummy_pchsant[16];    //[long  ,   16] 매입금액   StartPos 1185, Length 16
    char    deposit[16];    //[long  ,   16] 예수금   StartPos 1201, Length 16
    char    substamt[16];    //[long  ,   16] 대용금   StartPos 1217, Length 16
    char    csgnmnymgn[16];    //[long  ,   16] 위탁증거금현금   StartPos 1233, Length 16
    char    csgnsubstmgn[16];    //[long  ,   16] 위탁증거금대용   StartPos 1249, Length 16
    char    crdtpldgruseamt[16];    //[long  ,   16] 신용담보재사용금   StartPos 1265, Length 16
    char    ordablemny[16];    //[long  ,   16] 주문가능현금   StartPos 1281, Length 16
    char    ordablesubstamt[16];    //[long  ,   16] 주문가능대용   StartPos 1297, Length 16
    char    ruseableamt[16];    //[long  ,   16] 재사용가능금액   StartPos 1313, Length 16
} SC4_OutBlock;

//------------------------------------------------------------------------------
// 선물 옵션 계좌 주문 체결 내역 조회 (CFOAQ00600)
//------------------------------------------------------------------------------

typedef struct {
    char    RecCnt              [   5];    // [long  ,    5] 레코드갯수                      StartPos 0, Length 5
    char    AcntNo              [  20];    // [string,   20] 계좌번호                        StartPos 5, Length 20
    char    InptPwd             [   8];    // [string,    8] 입력비밀번호                    StartPos 25, Length 8
    char    QrySrtDt            [   8];    // [string,    8] 조회시작일                      StartPos 33, Length 8
    char    QryEndDt            [   8];    // [string,    8] 조회종료일                      StartPos 41, Length 8
    char    FnoClssCode         [   2];    // [string,    2] 선물옵션분류코드                StartPos 49, Length 2
    char    PrdgrpCode          [   2];    // [string,    2] 상품군코드                      StartPos 51, Length 2
    char    PrdtExecTpCode      [   1];    // [string,    1] 체결구분                        StartPos 53, Length 1
    char    StnlnSeqTp          [   1];    // [string,    1] 정렬순서구분                    StartPos 54, Length 1
    char    CommdaCode          [   2];    // [string,    2] 통신매체코드                    StartPos 55, Length 2
} CFOAQ00600InBlock1;

typedef struct {
    char    RecCnt              [   5];    // [long  ,    5] 레코드갯수                      StartPos 0, Length 5
    char    AcntNo              [  20];    // [string,   20] 계좌번호                        StartPos 5, Length 20
    char    InptPwd             [   8];    // [string,    8] 입력비밀번호                    StartPos 25, Length 8
    char    QrySrtDt            [   8];    // [string,    8] 조회시작일                      StartPos 33, Length 8
    char    QryEndDt            [   8];    // [string,    8] 조회종료일                      StartPos 41, Length 8
    char    FnoClssCode         [   2];    // [string,    2] 선물옵션분류코드                StartPos 49, Length 2
    char    PrdgrpCode          [   2];    // [string,    2] 상품군코드                      StartPos 51, Length 2
    char    PrdtExecTpCode      [   1];    // [string,    1] 체결구분                        StartPos 53, Length 1
    char    StnlnSeqTp          [   1];    // [string,    1] 정렬순서구분                    StartPos 54, Length 1
    char    CommdaCode          [   2];    // [string,    2] 통신매체코드                    StartPos 55, Length 2
} CFOAQ00600OutBlock1;

typedef struct {
    char    RecCnt              [   5];    // [long  ,    5] 레코드갯수                      StartPos 0, Length 5
    char    AcntNm              [  40];    // [string,   40] 계좌명                          StartPos 5, Length 40
    char    FutsOrdQty          [  16];    // [long  ,   16] 선물주문수량                    StartPos 45, Length 16
    char    FutsExecQty         [  16];    // [long  ,   16] 선물체결수량                    StartPos 61, Length 16
    char    OptOrdQty           [  16];    // [long  ,   16] 옵션주문수량                    StartPos 77, Length 16
    char    OptExecQty          [  16];    // [long  ,   16] 옵션체결수량                    StartPos 93, Length 16
} CFOAQ00600OutBlock2;

typedef struct {
    char    OrdDt               [   8];    // [string,    8] 주문일                          StartPos 0, Length 8
    char    OrdNo               [  10];    // [long  ,   10] 주문번호                        StartPos 8, Length 10
    char    OrgOrdNo            [  10];    // [long  ,   10] 원주문번호                      StartPos 18, Length 10
    char    OrdTime             [   9];    // [string,    9] 주문시각                        StartPos 28, Length 9
    char    FnoIsuNo            [  12];    // [string,   12] 선물옵션종목번호                StartPos 37, Length 12
    char    IsuNm               [  40];    // [string,   40] 종목명                          StartPos 49, Length 40
    char    BnsTpNm             [  10];    // [string,   10] 매매구분                        StartPos 89, Length 10
    char    MrcTpNm             [  10];    // [string,   10] 정정취소구분명                  StartPos 99, Length 10
    char    FnoOrdprcPtnCode    [   2];    // [string,    2] 선물옵션호가유형코드            StartPos 109, Length 2
    char    FnoOrdprcPtnNm      [  40];    // [string,   40] 선물옵션호가유형명              StartPos 111, Length 40
    char    OrdPrc              [  13];    // [double, 13.2] 주문가                          StartPos 151, Length 13
    char    OrdQty              [  16];    // [long  ,   16] 주문수량                        StartPos 164, Length 16
    char    OrdTpNm             [  10];    // [string,   10] 주문구분명                      StartPos 180, Length 10
    char    ExecTpNm            [  10];    // [string,   10] 체결구분명                      StartPos 190, Length 10
    char    ExecPrc             [  13];    // [double, 13.2] 체결가                          StartPos 200, Length 13
    char    ExecQty             [  16];    // [long  ,   16] 체결수량                        StartPos 213, Length 16
    char    CtrctTime           [   9];    // [string,    9] 약정시각                        StartPos 229, Length 9
    char    CtrctNo             [  10];    // [long  ,   10] 약정번호                        StartPos 238, Length 10
    char    ExecNo              [  10];    // [long  ,   10] 체결번호                        StartPos 248, Length 10
    char    BnsplAmt            [  16];    // [long  ,   16] 매매손익금액                    StartPos 258, Length 16
    char    UnercQty            [  16];    // [long  ,   16] 미체결수량                      StartPos 274, Length 16
    char    UserId              [  16];    // [string,   16] 사용자ID                        StartPos 290, Length 16
    char    CommdaCode          [   2];    // [string,    2] 통신매체코드                    StartPos 306, Length 2
    char    CommdaCodeNm        [  40];    // [string,   40] 통신매체코드명                  StartPos 308, Length 40
} CFOAQ00600OutBlock3;

//------------------------------------------------------------------------------
// 선물 옵션 정상 주문 (CFOAT00100)
//------------------------------------------------------------------------------

typedef struct {
    char    AcntNo              [  20];    // [string,   20] 계좌번호                        StartPos 0, Length 20
    char    Pwd                 [   8];    // [string,    8] 비밀번호                        StartPos 20, Length 8
    char    FnoIsuNo            [  12];    // [string,   12] 선물옵션종목번호                StartPos 28, Length 12
    char    BnsTpCode           [   1];    // [string,    1] 매매구분                        StartPos 40, Length 1
    char    FnoOrdprcPtnCode    [   2];    // [string,    2] 선물옵션호가유형코드            StartPos 41, Length 2
    char    OrdPrc              [  15];    // [double, 15.2] 주문가격                        StartPos 43, Length 15
    char    OrdQty              [  16];    // [long  ,   16] 주문수량                        StartPos 58, Length 16
} CFOAT00100InBlock1;

typedef struct {
    char    RecCnt              [   5];    // [long  ,    5] 레코드갯수                      StartPos 0, Length 5
    char    OrdMktCode          [   2];    // [string,    2] 주문시장코드                    StartPos 5, Length 2
    char    AcntNo              [  20];    // [string,   20] 계좌번호                        StartPos 7, Length 20
    char    Pwd                 [   8];    // [string,    8] 비밀번호                        StartPos 27, Length 8
    char    FnoIsuNo            [  12];    // [string,   12] 선물옵션종목번호                StartPos 35, Length 12
    char    BnsTpCode           [   1];    // [string,    1] 매매구분                        StartPos 47, Length 1
    char    FnoOrdPtnCode       [   2];    // [string,    2] 선물옵션주문유형코드            StartPos 48, Length 2
    char    FnoOrdprcPtnCode    [   2];    // [string,    2] 선물옵션호가유형코드            StartPos 50, Length 2
    char    FnoTrdPtnCode       [   2];    // [string,    2] 선물옵션거래유형코드            StartPos 52, Length 2
    char    OrdPrc              [  15];    // [double, 15.2] 주문가격                        StartPos 54, Length 15
    char    OrdQty              [  16];    // [long  ,   16] 주문수량                        StartPos 69, Length 16
    char    CommdaCode          [   2];    // [string,    2] 통신매체코드                    StartPos 85, Length 2
    char    DscusBnsCmpltTime   [   9];    // [string,    9] 협의매매완료시각                StartPos 87, Length 9
    char    GrpId               [  20];    // [string,   20] 그룹ID                          StartPos 96, Length 20
    char    OrdSeqno            [  10];    // [long  ,   10] 주문일련번호                    StartPos 116, Length 10
    char    PtflNo              [  10];    // [long  ,   10] 포트폴리오번호                  StartPos 126, Length 10
    char    BskNo               [  10];    // [long  ,   10] 바스켓번호                      StartPos 136, Length 10
    char    TrchNo              [  10];    // [long  ,   10] 트렌치번호                      StartPos 146, Length 10
    char    ItemNo              [  16];    // [long  ,   16] 항목번호                        StartPos 156, Length 16
    char    OpDrtnNo            [  12];    // [string,   12] 운용지시번호                    StartPos 172, Length 12
    char    MgempNo             [   9];    // [string,    9] 관리사원번호                    StartPos 184, Length 9
    char    FundId              [  12];    // [string,   12] 펀드ID                          StartPos 193, Length 12
    char    FundOrdNo           [  10];    // [long  ,   10] 펀드주문번호                    StartPos 205, Length 10
} CFOAT00100OutBlock1;

typedef struct {
    char    RecCnt              [   5];    // [long  ,    5] 레코드갯수                      StartPos 0, Length 5
    char    OrdNo               [  10];    // [long  ,   10] 주문번호                        StartPos 5, Length 10
    char    BrnNm               [  40];    // [string,   40] 지점명                          StartPos 15, Length 40
    char    AcntNm              [  40];    // [string,   40] 계좌명                          StartPos 55, Length 40
    char    IsuNm               [  50];    // [string,   50] 종목명                          StartPos 95, Length 50
    char    OrdAbleAmt          [  16];    // [long  ,   16] 주문가능금액                    StartPos 145, Length 16
    char    MnyOrdAbleAmt       [  16];    // [long  ,   16] 현금주문가능금액                StartPos 161, Length 16
    char    OrdMgn              [  16];    // [long  ,   16] 주문증거금                      StartPos 177, Length 16
    char    MnyOrdMgn           [  16];    // [long  ,   16] 현금주문증거금                  StartPos 193, Length 16
    char    OrdAbleQty          [  16];    // [long  ,   16] 주문가능수량                    StartPos 209, Length 16
} CFOAT00100OutBlock2;

typedef struct {
    CFOAT00100OutBlock1	outBlock1;
    CFOAT00100OutBlock2	outBlock2;
} CFOAT00100OutBlock;

//------------------------------------------------------------------------------
// 선물 옵션 정정 주문 (CFOAT00200)
//------------------------------------------------------------------------------

typedef struct {
    char    AcntNo              [  20];    // [string,   20] 계좌번호                        StartPos 0, Length 20
    char    Pwd                 [   8];    // [string,    8] 비밀번호                        StartPos 20, Length 8
    char    FnoIsuNo            [  12];    // [string,   12] 선물옵션종목번호                StartPos 28, Length 12
    char    OrgOrdNo            [  10];    // [long  ,   10] 원주문번호                      StartPos 40, Length 10
    char    FnoOrdprcPtnCode    [   2];    // [string,    2] 선물옵션호가유형코드            StartPos 50, Length 2
    char    OrdPrc              [  15];    // [double, 15.2] 주문가격                        StartPos 52, Length 15
    char    MdfyQty             [  16];    // [long  ,   16] 정정수량                        StartPos 67, Length 16
} CFOAT00200InBlock1;

typedef struct {
    char    RecCnt              [   5];    // [long  ,    5] 레코드갯수                      StartPos 0, Length 5
    char    OrdMktCode          [   2];    // [string,    2] 주문시장코드                    StartPos 5, Length 2
    char    AcntNo              [  20];    // [string,   20] 계좌번호                        StartPos 7, Length 20
    char    Pwd                 [   8];    // [string,    8] 비밀번호                        StartPos 27, Length 8
    char    FnoIsuNo            [  12];    // [string,   12] 선물옵션종목번호                StartPos 35, Length 12
    char    FnoOrdPtnCode       [   2];    // [string,    2] 선물옵션주문유형코드            StartPos 47, Length 2
    char    OrgOrdNo            [  10];    // [long  ,   10] 원주문번호                      StartPos 49, Length 10
    char    FnoOrdprcPtnCode    [   2];    // [string,    2] 선물옵션호가유형코드            StartPos 59, Length 2
    char    OrdPrc              [  15];    // [double, 15.2] 주문가격                        StartPos 61, Length 15
    char    MdfyQty             [  16];    // [long  ,   16] 정정수량                        StartPos 76, Length 16
    char    CommdaCode          [   2];    // [string,    2] 통신매체코드                    StartPos 92, Length 2
    char    DscusBnsCmpltTime   [   9];    // [string,    9] 협의매매완료시각                StartPos 94, Length 9
    char    GrpId               [  20];    // [string,   20] 그룹ID                          StartPos 103, Length 20
    char    OrdSeqno            [  10];    // [long  ,   10] 주문일련번호                    StartPos 123, Length 10
    char    PtflNo              [  10];    // [long  ,   10] 포트폴리오번호                  StartPos 133, Length 10
    char    BskNo               [  10];    // [long  ,   10] 바스켓번호                      StartPos 143, Length 10
    char    TrchNo              [  10];    // [long  ,   10] 트렌치번호                      StartPos 153, Length 10
    char    ItemNo              [  10];    // [long  ,   10] 아이템번호                      StartPos 163, Length 10
    char    MgempNo             [   9];    // [string,    9] 관리사원번호                    StartPos 173, Length 9
    char    FundId              [  12];    // [string,   12] 펀드ID                          StartPos 182, Length 12
    char    FundOrgOrdNo        [  10];    // [long  ,   10] 펀드원주문번호                  StartPos 194, Length 10
    char    FundOrdNo           [  10];    // [long  ,   10] 펀드주문번호                    StartPos 204, Length 10
} CFOAT00200OutBlock1;

typedef struct {
    char    RecCnt              [   5];    // [long  ,    5] 레코드갯수                      StartPos 0, Length 5
    char    OrdNo               [  10];    // [long  ,   10] 주문번호                        StartPos 5, Length 10
    char    BrnNm               [  40];    // [string,   40] 지점명                          StartPos 15, Length 40
    char    AcntNm              [  40];    // [string,   40] 계좌명                          StartPos 55, Length 40
    char    IsuNm               [  50];    // [string,   50] 종목명                          StartPos 95, Length 50
    char    OrdAbleAmt          [  16];    // [long  ,   16] 주문가능금액                    StartPos 145, Length 16
    char    MnyOrdAbleAmt       [  16];    // [long  ,   16] 현금주문가능금액                StartPos 161, Length 16
    char    OrdMgn              [  16];    // [long  ,   16] 주문증거금액                    StartPos 177, Length 16
    char    MnyOrdMgn           [  16];    // [long  ,   16] 현금주문증거금액                StartPos 193, Length 16
    char    OrdAbleQty          [  16];    // [long  ,   16] 주문가능수량                    StartPos 209, Length 16
} CFOAT00200OutBlock2;

typedef struct {
    CFOAT00200OutBlock1	outBlock1;
    CFOAT00200OutBlock2	outBlock2;
} CFOAT00200OutBlock;

//------------------------------------------------------------------------------
// 선물 옵션 취소 주문 (CFOAT00300)
//------------------------------------------------------------------------------

typedef struct {
    char    AcntNo              [  20];    // [string,   20] 계좌번호                        StartPos 0, Length 20
    char    Pwd                 [   8];    // [string,    8] 비밀번호                        StartPos 20, Length 8
    char    FnoIsuNo            [  12];    // [string,   12] 선물옵션종목번호                StartPos 28, Length 12
    char    OrgOrdNo            [  10];    // [long  ,   10] 원주문번호                      StartPos 40, Length 10
    char    CancQty             [  16];    // [long  ,   16] 취소수량                        StartPos 50, Length 16
} CFOAT00300InBlock1;

typedef struct {
    char    RecCnt              [   5];    // [long  ,    5] 레코드갯수                      StartPos 0, Length 5
    char    OrdMktCode          [   2];    // [string,    2] 주문시장코드                    StartPos 5, Length 2
    char    AcntNo              [  20];    // [string,   20] 계좌번호                        StartPos 7, Length 20
    char    Pwd                 [   8];    // [string,    8] 비밀번호                        StartPos 27, Length 8
    char    FnoIsuNo            [  12];    // [string,   12] 선물옵션종목번호                StartPos 35, Length 12
    char    FnoOrdPtnCode       [   2];    // [string,    2] 선물옵션주문유형코드            StartPos 47, Length 2
    char    OrgOrdNo            [  10];    // [long  ,   10] 원주문번호                      StartPos 49, Length 10
    char    CancQty             [  16];    // [long  ,   16] 취소수량                        StartPos 59, Length 16
    char    CommdaCode          [   2];    // [string,    2] 통신매체코드                    StartPos 75, Length 2
    char    DscusBnsCmpltTime   [   9];    // [string,    9] 협의매매완료시각                StartPos 77, Length 9
    char    GrpId               [  20];    // [string,   20] 그룹ID                          StartPos 86, Length 20
    char    OrdSeqno            [  10];    // [long  ,   10] 주문일련번호                    StartPos 106, Length 10
    char    PtflNo              [  10];    // [long  ,   10] 포트폴리오번호                  StartPos 116, Length 10
    char    BskNo               [  10];    // [long  ,   10] 바스켓번호                      StartPos 126, Length 10
    char    TrchNo              [  10];    // [long  ,   10] 트렌치번호                      StartPos 136, Length 10
    char    ItemNo              [  10];    // [long  ,   10] 아이템번호                      StartPos 146, Length 10
    char    MgempNo             [   9];    // [string,    9] 관리사원번호                    StartPos 156, Length 9
    char    FundId              [  12];    // [string,   12] 펀드ID                          StartPos 165, Length 12
    char    FundOrgOrdNo        [  10];    // [long  ,   10] 펀드원주문번호                  StartPos 177, Length 10
    char    FundOrdNo           [  10];    // [long  ,   10] 펀드주문번호                    StartPos 187, Length 10
} CFOAT00300OutBlock1;

typedef struct {
    char    RecCnt              [   5];    // [long  ,    5] 레코드갯수                      StartPos 0, Length 5
    char    OrdNo               [  10];    // [long  ,   10] 주문번호                        StartPos 5, Length 10
    char    BrnNm               [  40];    // [string,   40] 지점명                          StartPos 15, Length 40
    char    AcntNm              [  40];    // [string,   40] 계좌명                          StartPos 55, Length 40
    char    IsuNm               [  50];    // [string,   50] 종목명                          StartPos 95, Length 50
    char    OrdAbleAmt          [  16];    // [long  ,   16] 주문가능금액                    StartPos 145, Length 16
    char    MnyOrdAbleAmt       [  16];    // [long  ,   16] 현금주문가능금액                StartPos 161, Length 16
    char    OrdMgn              [  16];    // [long  ,   16] 주문증거금액                    StartPos 177, Length 16
    char    MnyOrdMgn           [  16];    // [long  ,   16] 현금주문증거금액                StartPos 193, Length 16
    char    OrdAbleQty          [  16];    // [long  ,   16] 주문가능수량                    StartPos 209, Length 16
} CFOAT00300OutBlock2;

typedef struct {
    CFOAT00300OutBlock1	outBlock1;
    CFOAT00300OutBlock2	outBlock2;
} CFOAT00300OutBlock;

//------------------------------------------------------------------------------
// 선물 옵션 계좌 예탁금 증거금 조회 (CFOBQ10500)
//------------------------------------------------------------------------------

typedef struct {
    char    RecCnt              [   5];    // [long  ,    5] 레코드갯수                      StartPos 0, Length 5
    char    AcntNo              [  20];    // [string,   20] 계좌번호                        StartPos 5, Length 20
    char    Pwd                 [   8];    // [string,    8] 비밀번호                        StartPos 25, Length 8
} CFOBQ10500InBlock1;

typedef struct {
    char    RecCnt              [   5];    // [long  ,    5] 레코드갯수                      StartPos 0, Length 5
    char    AcntNo              [  20];    // [string,   20] 계좌번호                        StartPos 5, Length 20
    char    Pwd                 [   8];    // [string,    8] 비밀번호                        StartPos 25, Length 8
} CFOBQ10500OutBlock1;

typedef struct {
    char    RecCnt              [   5];    // [long  ,    5] 레코드갯수                      StartPos 0, Length 5
    char    AcntNm              [  40];    // [string,   40] 계좌명                          StartPos 5, Length 40
    char    DpsamtTotamt        [  16];    // [long  ,   16] 예탁금총액                      StartPos 45, Length 16
    char    Dps                 [  16];    // [long  ,   16] 예수금                          StartPos 61, Length 16
    char    SubstAmt            [  16];    // [long  ,   16] 대용금액                        StartPos 77, Length 16
    char    FilupDpsamtTotamt   [  16];    // [long  ,   16] 충당예탁금총액                  StartPos 93, Length 16
    char    FilupDps            [  16];    // [long  ,   16] 충당예수금                      StartPos 109, Length 16
    char    FutsPnlAmt          [  16];    // [long  ,   16] 선물손익금액                    StartPos 125, Length 16
    char    WthdwAbleAmt        [  16];    // [long  ,   16] 인출가능금액                    StartPos 141, Length 16
    char    PsnOutAbleCurAmt    [  16];    // [long  ,   16] 인출가능현금액                  StartPos 157, Length 16
    char    PsnOutAbleSubstAmt  [  16];    // [long  ,   16] 인출가능대용금액                StartPos 173, Length 16
    char    Mgn                 [  16];    // [long  ,   16] 증거금액                        StartPos 189, Length 16
    char    MnyMgn              [  16];    // [long  ,   16] 현금증거금액                    StartPos 205, Length 16
    char    OrdAbleAmt          [  16];    // [long  ,   16] 주문가능금액                    StartPos 221, Length 16
    char    MnyOrdAbleAmt       [  16];    // [long  ,   16] 현금주문가능금액                StartPos 237, Length 16
    char    AddMgn              [  16];    // [long  ,   16] 추가증거금액                    StartPos 253, Length 16
    char    MnyAddMgn           [  16];    // [long  ,   16] 현금추가증거금액                StartPos 269, Length 16
    char    AmtPrdayChckInAmt   [  16];    // [long  ,   16] 금전일수표입금액                StartPos 285, Length 16
    char    FnoPrdaySubstSellAmt[  16];    // [long  ,   16] 선물옵션전일대용매도금액        StartPos 301, Length 16
    char    FnoCrdaySubstSellAmt[  16];    // [long  ,   16] 선물옵션금일대용매도금액        StartPos 317, Length 16
    char    FnoPrdayFdamt       [  16];    // [long  ,   16] 선물옵션전일가입금액            StartPos 333, Length 16
    char    FnoCrdayFdamt       [  16];    // [long  ,   16] 선물옵션금일가입금액            StartPos 349, Length 16
    char    FcurrSubstAmt       [  16];    // [long  ,   16] 외화대용금액                    StartPos 365, Length 16
    char    FnoAcntAfmgnNm      [  20];    // [string,   20] 선물옵션계좌사후증거금명        StartPos 381, Length 20
} CFOBQ10500OutBlock2;

typedef struct {
    char    PdGrpCodeNm         [  20];    // [string,   20] 상품군코드명                    StartPos 0, Length 20
    char    NetRiskMgn          [  16];    // [long  ,   16] 순위험증거금액                  StartPos 20, Length 16
    char    PrcMgn              [  16];    // [long  ,   16] 가격증거금액                    StartPos 36, Length 16
    char    SprdMgn             [  16];    // [long  ,   16] 스프레드증거금액                StartPos 52, Length 16
    char    PrcFlctMgn          [  16];    // [long  ,   16] 가격변동증거금액                StartPos 68, Length 16
    char    MinMgn              [  16];    // [long  ,   16] 최소증거금액                    StartPos 84, Length 16
    char    OrdMgn              [  16];    // [long  ,   16] 주문증거금액                    StartPos 100, Length 16
    char    OptNetBuyAmt        [  16];    // [long  ,   16] 옵션순매수금액                  StartPos 116, Length 16
    char    CsgnMgn             [  16];    // [long  ,   16] 위탁증거금액                    StartPos 132, Length 16
    char    MaintMgn            [  16];    // [long  ,   16] 유지증거금액                    StartPos 148, Length 16
    char    FutsBuyExecAmt      [  16];    // [long  ,   16] 선물매수체결금액                StartPos 164, Length 16
    char    FutsSellExecAmt     [  16];    // [long  ,   16] 선물매도체결금액                StartPos 180, Length 16
    char    OptBuyExecAmt       [  16];    // [long  ,   16] 옵션매수체결금액                StartPos 196, Length 16
    char    OptSellExecAmt      [  16];    // [long  ,   16] 옵션매도체결금액                StartPos 212, Length 16
    char    FutsPnlAmt          [  16];    // [long  ,   16] 선물손익금액                    StartPos 228, Length 16
    char    TotRiskCsgnMgn      [  16];    // [long  ,   16] 총위험위탁증거금                StartPos 244, Length 16
    char    UndCsgnMgn          [  16];    // [long  ,   16] 인수도위탁증거금                StartPos 260, Length 16
    char    MgnRdctAmt          [  16];    // [long  ,   16] 증거금감면금액                  StartPos 276, Length 16
} CFOBQ10500OutBlock3;

//------------------------------------------------------------------------------
// 선물 옵션 계좌 미결제 약정 현황 (평균가)  (CFOFQ02400)
//------------------------------------------------------------------------------

typedef struct {
    char    RecCnt              [   5];    // [long  ,    5] 레코드갯수                      StartPos 0, Length 5
    char    AcntNo              [  20];    // [string,   20] 계좌번호                        StartPos 5, Length 20
    char    Pwd                 [   8];    // [string,    8] 비밀번호                        StartPos 25, Length 8
    char    RegMktCode          [   2];    // [string,    2] 등록시장코드                    StartPos 33, Length 2
    char    BuyDt               [   8];    // [string,    8] 매수일자                        StartPos 35, Length 8
} CFOFQ02400InBlock1;

typedef struct {
    char    RecCnt              [   5];    // [long  ,    5] 레코드갯수                      StartPos 0, Length 5
    char    AcntNo              [  20];    // [string,   20] 계좌번호                        StartPos 5, Length 20
    char    Pwd                 [   8];    // [string,    8] 비밀번호                        StartPos 25, Length 8
    char    RegMktCode          [   2];    // [string,    2] 등록시장코드                    StartPos 33, Length 2
    char    BuyDt               [   8];    // [string,    8] 매수일자                        StartPos 35, Length 8
} CFOFQ02400OutBlock1;

typedef struct {
    char    RecCnt              [   5];    // [long  ,    5] 레코드갯수                      StartPos 0, Length 5
    char    AcntNm              [  40];    // [string,   40] 계좌명                          StartPos 5, Length 40
    char    FutsCtrctQty        [  16];    // [long  ,   16] 선물약정수량                    StartPos 45, Length 16
    char    OptCtrctQty         [  16];    // [long  ,   16] 옵션약정수량                    StartPos 61, Length 16
    char    CtrctQty            [  16];    // [long  ,   16] 약정수량                        StartPos 77, Length 16
    char    FutsCtrctAmt        [  16];    // [long  ,   16] 선물약정금액                    StartPos 93, Length 16
    char    FutsBuyctrAmt       [  16];    // [long  ,   16] 선물매수약정금액                StartPos 109, Length 16
    char    FutsSlctrAmt        [  16];    // [long  ,   16] 선물매도약정금액                StartPos 125, Length 16
    char    CalloptCtrctAmt     [  16];    // [long  ,   16] 콜옵션약정금액                  StartPos 141, Length 16
    char    CallBuyAmt          [  16];    // [long  ,   16] 콜매수금액                      StartPos 157, Length 16
    char    CallSellAmt         [  16];    // [long  ,   16] 콜매도금액                      StartPos 173, Length 16
    char    PutoptCtrctAmt      [  16];    // [long  ,   16] 풋옵션약정금액                  StartPos 189, Length 16
    char    PutBuyAmt           [  16];    // [long  ,   16] 풋매수금액                      StartPos 205, Length 16
    char    PutSellAmt          [  16];    // [long  ,   16] 풋매도금액                      StartPos 221, Length 16
    char    AllCtrctAmt         [  16];    // [long  ,   16] 전체약정금액                    StartPos 237, Length 16
    char    BuyctrAsmAmt        [  16];    // [long  ,   16] 매수약정누계금액                StartPos 253, Length 16
    char    SlctrAsmAmt         [  16];    // [long  ,   16] 매도약정누계금액                StartPos 269, Length 16
    char    FutsPnlSum          [  16];    // [long  ,   16] 선물손익합계                    StartPos 285, Length 16
    char    OptPnlSum           [  16];    // [long  ,   16] 옵션손익합계                    StartPos 301, Length 16
    char    AllPnlSum           [  16];    // [long  ,   16] 전체손익합계                    StartPos 317, Length 16
} CFOFQ02400OutBlock2;

typedef struct {
    char    FnoClssCode         [   1];    // [string,    1] 선물옵션품목구분                StartPos 0, Length 1
    char    FutsSellQty         [  16];    // [long  ,   16] 선물매도수량                    StartPos 1, Length 16
    char    FutsSellPnl         [  16];    // [long  ,   16] 선물매도손익                    StartPos 17, Length 16
    char    FutsBuyQty          [  16];    // [long  ,   16] 선물매수수량                    StartPos 33, Length 16
    char    FutsBuyPnl          [  16];    // [long  ,   16] 선물매수손익                    StartPos 49, Length 16
    char    CallSellQty         [  16];    // [long  ,   16] 콜매도수량                      StartPos 65, Length 16
    char    CallSellPnl         [  16];    // [long  ,   16] 콜매도손익                      StartPos 81, Length 16
    char    CallBuyQty          [  16];    // [long  ,   16] 콜매수수량                      StartPos 97, Length 16
    char    CallBuyPnl          [  16];    // [long  ,   16] 콜매수손익                      StartPos 113, Length 16
    char    PutSellQty          [  16];    // [long  ,   16] 풋매도수량                      StartPos 129, Length 16
    char    PutSellPnl          [  16];    // [long  ,   16] 풋매도손익                      StartPos 145, Length 16
    char    PutBuyQty           [  16];    // [long  ,   16] 풋매수수량                      StartPos 161, Length 16
    char    PutBuyPnl           [  16];    // [long  ,   16] 풋매수손익                      StartPos 177, Length 16
} CFOFQ02400OutBlock3;

typedef struct {
    char    IsuNo               [  12];    // [string,   12] 종목번호                        StartPos 0, Length 12
    char    IsuNm               [  40];    // [string,   40] 종목명                          StartPos 12, Length 40
    char    BnsTpCode           [   1];    // [string,    1] 매매구분                        StartPos 52, Length 1
    char    BnsTpNm             [  10];    // [string,   10] 매매구분                        StartPos 53, Length 10
    char    BalQty              [  16];    // [long  ,   16] 잔고수량                        StartPos 63, Length 16
    char    FnoAvrPrc           [  19];    // [double, 19.8] 평균가                          StartPos 79, Length 19
    char    BgnAmt              [  16];    // [long  ,   16] 당초금액                        StartPos 98, Length 16
    char    ThdayLqdtQty        [  16];    // [long  ,   16] 당일청산수량                    StartPos 114, Length 16
    char    Curprc              [  13];    // [double, 13.2] 현재가                          StartPos 130, Length 13
    char    EvalAmt             [  16];    // [long  ,   16] 평가금액                        StartPos 143, Length 16
    char    EvalPnlAmt          [  16];    // [long  ,   16] 평가손익금액                    StartPos 159, Length 16
    char    EvalErnrat          [  12];    // [double, 12.6] 평가수익률                      StartPos 175, Length 12
} CFOFQ02400OutBlock4;

//------------------------------------------------------------------------------
// 현물 계좌 총평가 (CSPAQ12200)
//------------------------------------------------------------------------------

typedef struct {
    char    RecCnt              [   5];    // [long  ,    5] 레코드갯수                      StartPos 0, Length 5
    char    MgmtBrnNo           [   3];    // [string,    3] 관리지점번호                    StartPos 5, Length 3
    char    AcntNo              [  20];    // [string,   20] 계좌번호                        StartPos 8, Length 20
    char    Pwd                 [   8];    // [string,    8] 비밀번호                        StartPos 28, Length 8
    char    BalCreTp            [   1];    // [string,    1] 잔고생성구분                    StartPos 36, Length 1
} CSPAQ12200InBlock1;

typedef struct {
    char    RecCnt              [   5];    // [long  ,    5] 레코드갯수                      StartPos 0, Length 5
    char    MgmtBrnNo           [   3];    // [string,    3] 관리지점번호                    StartPos 5, Length 3
    char    AcntNo              [  20];    // [string,   20] 계좌번호                        StartPos 8, Length 20
    char    Pwd                 [   8];    // [string,    8] 비밀번호                        StartPos 28, Length 8
    char    BalCreTp            [   1];    // [string,    1] 잔고생성구분                    StartPos 36, Length 1
} CSPAQ12200OutBlock1;

typedef struct {
    char    RecCnt              [   5];    // [long  ,    5] 레코드갯수                      StartPos 0, Length 5
    char    BrnNm               [  40];    // [string,   40] 지점명                          StartPos 5, Length 40
    char    AcntNm              [  40];    // [string,   40] 계좌명                          StartPos 45, Length 40
    char    MnyOrdAbleAmt       [  16];    // [long  ,   16] 현금주문가능금액                StartPos 85, Length 16
    char    MnyoutAbleAmt       [  16];    // [long  ,   16] 출금가능금액                    StartPos 101, Length 16
    char    SeOrdAbleAmt        [  16];    // [long  ,   16] 거래소금액                      StartPos 117, Length 16
    char    KdqOrdAbleAmt       [  16];    // [long  ,   16] 코스닥금액                      StartPos 133, Length 16
    char    BalEvalAmt          [  16];    // [long  ,   16] 잔고평가금액                    StartPos 149, Length 16
    char    RcvblAmt            [  16];    // [long  ,   16] 미수금액                        StartPos 165, Length 16
    char    DpsastTotamt        [  16];    // [long  ,   16] 예탁자산총액                    StartPos 181, Length 16
    char    PnlRat              [  18];    // [double, 18.6] 손익율                          StartPos 197, Length 18
    char    InvstOrgAmt         [  20];    // [long  ,   20] 투자원금                        StartPos 215, Length 20
    char    InvstPlAmt          [  16];    // [long  ,   16] 투자손익금액                    StartPos 235, Length 16
    char    CrdtPldgOrdAmt      [  16];    // [long  ,   16] 신용담보주문금액                StartPos 251, Length 16
    char    Dps                 [  16];    // [long  ,   16] 예수금                          StartPos 267, Length 16
    char    SubstAmt            [  16];    // [long  ,   16] 대용금액                        StartPos 283, Length 16
    char    D1Dps               [  16];    // [long  ,   16] D1예수금                        StartPos 299, Length 16
    char    D2Dps               [  16];    // [long  ,   16] D2예수금                        StartPos 315, Length 16
    char    MnyrclAmt           [  16];    // [long  ,   16] 현금미수금액                    StartPos 331, Length 16
    char    MgnMny              [  16];    // [long  ,   16] 증거금현금                      StartPos 347, Length 16
    char    MgnSubst            [  16];    // [long  ,   16] 증거금대용                      StartPos 363, Length 16
    char    ChckAmt             [  16];    // [long  ,   16] 수표금액                        StartPos 379, Length 16
    char    SubstOrdAbleAmt     [  16];    // [long  ,   16] 대용주문가능금액                StartPos 395, Length 16
    char    MgnRat100pctOrdAbleAmt[  16];    // [long  ,   16] 증거금률100퍼센트주문가능금액   StartPos 411, Length 16
    char    MgnRat35ordAbleAmt  [  16];    // [long  ,   16] 증거금률35%주문가능금액         StartPos 427, Length 16
    char    MgnRat50ordAbleAmt  [  16];    // [long  ,   16] 증거금률50%주문가능금액         StartPos 443, Length 16
    char    PrdaySellAdjstAmt   [  16];    // [long  ,   16] 전일매도정산금액                StartPos 459, Length 16
    char    PrdayBuyAdjstAmt    [  16];    // [long  ,   16] 전일매수정산금액                StartPos 475, Length 16
    char    CrdaySellAdjstAmt   [  16];    // [long  ,   16] 금일매도정산금액                StartPos 491, Length 16
    char    CrdayBuyAdjstAmt    [  16];    // [long  ,   16] 금일매수정산금액                StartPos 507, Length 16
    char    D1ovdRepayRqrdAmt   [  16];    // [long  ,   16] D1연체변제소요금액              StartPos 523, Length 16
    char    D2ovdRepayRqrdAmt   [  16];    // [long  ,   16] D2연체변제소요금액              StartPos 539, Length 16
    char    D1PrsmptWthdwAbleAmt[  16];    // [long  ,   16] D1추정인출가능금액              StartPos 555, Length 16
    char    D2PrsmptWthdwAbleAmt[  16];    // [long  ,   16] D2추정인출가능금액              StartPos 571, Length 16
    char    DpspdgLoanAmt       [  16];    // [long  ,   16] 예탁담보대출금액                StartPos 587, Length 16
    char    Imreq               [  16];    // [long  ,   16] 신용설정보증금                  StartPos 603, Length 16
    char    MloanAmt            [  16];    // [long  ,   16] 융자금액                        StartPos 619, Length 16
    char    ChgAfPldgRat        [   9];    // [double,  9.3] 변경후담보비율                  StartPos 635, Length 9
    char    OrgPldgAmt          [  16];    // [long  ,   16] 원담보금액                      StartPos 644, Length 16
    char    SubPldgAmt          [  16];    // [long  ,   16] 부담보금액                      StartPos 660, Length 16
    char    RqrdPldgAmt         [  16];    // [long  ,   16] 소요담보금액                    StartPos 676, Length 16
    char    OrgPdlckAmt         [  16];    // [long  ,   16] 원담보부족금액                  StartPos 692, Length 16
    char    PdlckAmt            [  16];    // [long  ,   16] 담보부족금액                    StartPos 708, Length 16
    char    AddPldgMny          [  16];    // [long  ,   16] 추가담보현금                    StartPos 724, Length 16
    char    D1OrdAbleAmt        [  16];    // [long  ,   16] D1주문가능금액                  StartPos 740, Length 16
    char    CrdtIntdltAmt       [  16];    // [long  ,   16] 신용이자미납금액                StartPos 756, Length 16
    char    EtclndAmt           [  16];    // [long  ,   16] 기타대여금액                    StartPos 772, Length 16
    char    NtdayPrsmptCvrgAmt  [  16];    // [long  ,   16] 익일추정반대매매금액            StartPos 788, Length 16
    char    OrgPldgSumAmt       [  16];    // [long  ,   16] 원담보합계금액                  StartPos 804, Length 16
    char    CrdtOrdAbleAmt      [  16];    // [long  ,   16] 신용주문가능금액                StartPos 820, Length 16
    char    SubPldgSumAmt       [  16];    // [long  ,   16] 부담보합계금액                  StartPos 836, Length 16
    char    CrdtPldgAmtMny      [  16];    // [long  ,   16] 신용담보금현금                  StartPos 852, Length 16
    char    CrdtPldgSubstAmt    [  16];    // [long  ,   16] 신용담보대용금액                StartPos 868, Length 16
    char    AddCrdtPldgMny      [  16];    // [long  ,   16] 추가신용담보현금                StartPos 884, Length 16
    char    CrdtPldgRuseAmt     [  16];    // [long  ,   16] 신용담보재사용금액              StartPos 900, Length 16
    char    AddCrdtPldgSubst    [  16];    // [long  ,   16] 추가신용담보대용                StartPos 916, Length 16
    char    CslLoanAmtdt1       [  16];    // [long  ,   16] 매도대금담보대출금액            StartPos 932, Length 16
    char    DpslRestrcAmt       [  16];    // [long  ,   16] 처분제한금액                    StartPos 948, Length 16
} CSPAQ12200OutBlock2;

typedef struct {
    CSPAQ12200OutBlock1	outBlock1;
    CSPAQ12200OutBlock2	outBlock2;
} CSPAQ12200OutBlock;

//------------------------------------------------------------------------------
// 현물 계좌 잔고 내역 조회 (CSPAQ12300)
//------------------------------------------------------------------------------

typedef struct {
    char    RecCnt              [   5];    // [long  ,    5] 레코드갯수                      StartPos 0, Length 5
    char    AcntNo              [  20];    // [string,   20] 계좌번호                        StartPos 5, Length 20
    char    Pwd                 [   8];    // [string,    8] 비밀번호                        StartPos 25, Length 8
    char    BalCreTp            [   1];    // [string,    1] 잔고생성구분                    StartPos 33, Length 1
    char    CmsnAppTpCode       [   1];    // [string,    1] 수수료적용구분                  StartPos 34, Length 1
    char    D2balBaseQryTp      [   1];    // [string,    1] D2잔고기준조회구분              StartPos 35, Length 1
    char    UprcTpCode          [   1];    // [string,    1] 단가구분                        StartPos 36, Length 1
} CSPAQ12300InBlock1;

typedef struct {
    char    RecCnt              [   5];    // [long  ,    5] 레코드갯수                      StartPos 0, Length 5
    char    AcntNo              [  20];    // [string,   20] 계좌번호                        StartPos 5, Length 20
    char    Pwd                 [   8];    // [string,    8] 비밀번호                        StartPos 25, Length 8
    char    BalCreTp            [   1];    // [string,    1] 잔고생성구분                    StartPos 33, Length 1
    char    CmsnAppTpCode       [   1];    // [string,    1] 수수료적용구분                  StartPos 34, Length 1
    char    D2balBaseQryTp      [   1];    // [string,    1] D2잔고기준조회구분              StartPos 35, Length 1
    char    UprcTpCode          [   1];    // [string,    1] 단가구분                        StartPos 36, Length 1
} CSPAQ12300OutBlock1;

typedef struct {
    char    RecCnt              [   5];    // [long  ,    5] 레코드갯수                      StartPos 0, Length 5
    char    BrnNm               [  40];    // [string,   40] 지점명                          StartPos 5, Length 40
    char    AcntNm              [  40];    // [string,   40] 계좌명                          StartPos 45, Length 40
    char    MnyOrdAbleAmt       [  16];    // [long  ,   16] 현금주문가능금액                StartPos 85, Length 16
    char    MnyoutAbleAmt       [  16];    // [long  ,   16] 출금가능금액                    StartPos 101, Length 16
    char    SeOrdAbleAmt        [  16];    // [long  ,   16] 거래소금액                      StartPos 117, Length 16
    char    KdqOrdAbleAmt       [  16];    // [long  ,   16] 코스닥금액                      StartPos 133, Length 16
    char    HtsOrdAbleAmt       [  16];    // [long  ,   16] HTS주문가능금액                 StartPos 149, Length 16
    char    MgnRat100pctOrdAbleAmt[  16];    // [long  ,   16] 증거금률100퍼센트주문가능금액   StartPos 165, Length 16
    char    BalEvalAmt          [  16];    // [long  ,   16] 잔고평가금액                    StartPos 181, Length 16
    char    PchsAmt             [  16];    // [long  ,   16] 매입금액                        StartPos 197, Length 16
    char    RcvblAmt            [  16];    // [long  ,   16] 미수금액                        StartPos 213, Length 16
    char    PnlRat              [  18];    // [double, 18.6] 손익율                          StartPos 229, Length 18
    char    InvstOrgAmt         [  20];    // [long  ,   20] 투자원금                        StartPos 247, Length 20
    char    InvstPlAmt          [  16];    // [long  ,   16] 투자손익금액                    StartPos 267, Length 16
    char    CrdtPldgOrdAmt      [  16];    // [long  ,   16] 신용담보주문금액                StartPos 283, Length 16
    char    Dps                 [  16];    // [long  ,   16] 예수금                          StartPos 299, Length 16
    char    D1Dps               [  16];    // [long  ,   16] D1예수금                        StartPos 315, Length 16
    char    D2Dps               [  16];    // [long  ,   16] D2예수금                        StartPos 331, Length 16
    char    OrdDt               [   8];    // [string,    8] 주문일                          StartPos 347, Length 8
    char    MnyMgn              [  16];    // [long  ,   16] 현금증거금액                    StartPos 355, Length 16
    char    SubstMgn            [  16];    // [long  ,   16] 대용증거금액                    StartPos 371, Length 16
    char    SubstAmt            [  16];    // [long  ,   16] 대용금액                        StartPos 387, Length 16
    char    PrdayBuyExecAmt     [  16];    // [long  ,   16] 전일매수체결금액                StartPos 403, Length 16
    char    PrdaySellExecAmt    [  16];    // [long  ,   16] 전일매도체결금액                StartPos 419, Length 16
    char    CrdayBuyExecAmt     [  16];    // [long  ,   16] 금일매수체결금액                StartPos 435, Length 16
    char    CrdaySellExecAmt    [  16];    // [long  ,   16] 금일매도체결금액                StartPos 451, Length 16
    char    EvalPnlSum          [  15];    // [long  ,   15] 평가손익합계                    StartPos 467, Length 15
    char    DpsastTotamt        [  16];    // [long  ,   16] 예탁자산총액                    StartPos 482, Length 16
    char    Evrprc              [  19];    // [long  ,   19] 제비용                          StartPos 498, Length 19
    char    RuseAmt             [  16];    // [long  ,   16] 재사용금액                      StartPos 517, Length 16
    char    EtclndAmt           [  16];    // [long  ,   16] 기타대여금액                    StartPos 533, Length 16
    char    PrcAdjstAmt         [  16];    // [long  ,   16] 가정산금액                      StartPos 549, Length 16
    char    D1CmsnAmt           [  16];    // [long  ,   16] D1수수료                        StartPos 565, Length 16
    char    D2CmsnAmt           [  16];    // [long  ,   16] D2수수료                        StartPos 581, Length 16
    char    D1EvrTax            [  16];    // [long  ,   16] D1제세금                        StartPos 597, Length 16
    char    D2EvrTax            [  16];    // [long  ,   16] D2제세금                        StartPos 613, Length 16
    char    D1SettPrergAmt      [  16];    // [long  ,   16] D1결제예정금액                  StartPos 629, Length 16
    char    D2SettPrergAmt      [  16];    // [long  ,   16] D2결제예정금액                  StartPos 645, Length 16
    char    PrdayKseMnyMgn      [  16];    // [long  ,   16] 전일KSE현금증거금               StartPos 661, Length 16
    char    PrdayKseSubstMgn    [  16];    // [long  ,   16] 전일KSE대용증거금               StartPos 677, Length 16
    char    PrdayKseCrdtMnyMgn  [  16];    // [long  ,   16] 전일KSE신용현금증거금           StartPos 693, Length 16
    char    PrdayKseCrdtSubstMgn[  16];    // [long  ,   16] 전일KSE신용대용증거금           StartPos 709, Length 16
    char    CrdayKseMnyMgn      [  16];    // [long  ,   16] 금일KSE현금증거금               StartPos 725, Length 16
    char    CrdayKseSubstMgn    [  16];    // [long  ,   16] 금일KSE대용증거금               StartPos 741, Length 16
    char    CrdayKseCrdtMnyMgn  [  16];    // [long  ,   16] 금일KSE신용현금증거금           StartPos 757, Length 16
    char    CrdayKseCrdtSubstMgn[  16];    // [long  ,   16] 금일KSE신용대용증거금           StartPos 773, Length 16
    char    PrdayKdqMnyMgn      [  16];    // [long  ,   16] 전일코스닥현금증거금            StartPos 789, Length 16
    char    PrdayKdqSubstMgn    [  16];    // [long  ,   16] 전일코스닥대용증거금            StartPos 805, Length 16
    char    PrdayKdqCrdtMnyMgn  [  16];    // [long  ,   16] 전일코스닥신용현금증거금        StartPos 821, Length 16
    char    PrdayKdqCrdtSubstMgn[  16];    // [long  ,   16] 전일코스닥신용대용증거금        StartPos 837, Length 16
    char    CrdayKdqMnyMgn      [  16];    // [long  ,   16] 금일코스닥현금증거금            StartPos 853, Length 16
    char    CrdayKdqSubstMgn    [  16];    // [long  ,   16] 금일코스닥대용증거금            StartPos 869, Length 16
    char    CrdayKdqCrdtMnyMgn  [  16];    // [long  ,   16] 금일코스닥신용현금증거금        StartPos 885, Length 16
    char    CrdayKdqCrdtSubstMgn[  16];    // [long  ,   16] 금일코스닥신용대용증거금        StartPos 901, Length 16
    char    PrdayFrbrdMnyMgn    [  16];    // [long  ,   16] 전일프리보드현금증거금          StartPos 917, Length 16
    char    PrdayFrbrdSubstMgn  [  16];    // [long  ,   16] 전일프리보드대용증거금          StartPos 933, Length 16
    char    CrdayFrbrdMnyMgn    [  16];    // [long  ,   16] 금일프리보드현금증거금          StartPos 949, Length 16
    char    CrdayFrbrdSubstMgn  [  16];    // [long  ,   16] 금일프리보드대용증거금          StartPos 965, Length 16
    char    PrdayCrbmkMnyMgn    [  16];    // [long  ,   16] 전일장외현금증거금              StartPos 981, Length 16
    char    PrdayCrbmkSubstMgn  [  16];    // [long  ,   16] 전일장외대용증거금              StartPos 997, Length 16
    char    CrdayCrbmkMnyMgn    [  16];    // [long  ,   16] 금일장외현금증거금              StartPos 1013, Length 16
    char    CrdayCrbmkSubstMgn  [  16];    // [long  ,   16] 금일장외대용증거금              StartPos 1029, Length 16
    char    DpspdgQty           [  16];    // [long  ,   16] 예탁담보수량                    StartPos 1045, Length 16
    char    BuyAdjstAmtD2       [  16];    // [long  ,   16] 매수정산금(D+2)                 StartPos 1061, Length 16
    char    SellAdjstAmtD2      [  16];    // [long  ,   16] 매도정산금(D+2)                 StartPos 1077, Length 16
    char    RepayRqrdAmtD1      [  16];    // [long  ,   16] 변제소요금(D+1)                 StartPos 1093, Length 16
    char    RepayRqrdAmtD2      [  16];    // [long  ,   16] 변제소요금(D+2)                 StartPos 1109, Length 16
    char    LoanAmt             [  16];    // [long  ,   16] 대출금액                        StartPos 1125, Length 16
} CSPAQ12300OutBlock2;

typedef struct {
    char    IsuNo               [  12];    // [string,   12] 종목번호                        StartPos 0, Length 12
    char    IsuNm               [  40];    // [string,   40] 종목명                          StartPos 12, Length 40
    char    SecBalPtnCode       [   2];    // [string,    2] 유가증권잔고유형코드            StartPos 52, Length 2
    char    SecBalPtnNm         [  40];    // [string,   40] 유가증권잔고유형명              StartPos 54, Length 40
    char    BalQty              [  16];    // [long  ,   16] 잔고수량                        StartPos 94, Length 16
    char    BnsBaseBalQty       [  16];    // [long  ,   16] 매매기준잔고수량                StartPos 110, Length 16
    char    CrdayBuyExecQty     [  16];    // [long  ,   16] 금일매수체결수량                StartPos 126, Length 16
    char    CrdaySellExecQty    [  16];    // [long  ,   16] 금일매도체결수량                StartPos 142, Length 16
    char    SellPrc             [  21];    // [double, 21.4] 매도가                          StartPos 158, Length 21
    char    BuyPrc              [  21];    // [double, 21.4] 매수가                          StartPos 179, Length 21
    char    SellPnlAmt          [  16];    // [long  ,   16] 매도손익금액                    StartPos 200, Length 16
    char    PnlRat              [  18];    // [double, 18.6] 손익율                          StartPos 216, Length 18
    char    NowPrc              [  15];    // [double, 15.2] 현재가                          StartPos 234, Length 15
    char    CrdtAmt             [  16];    // [long  ,   16] 신용금액                        StartPos 249, Length 16
    char    DueDt               [   8];    // [string,    8] 만기일                          StartPos 265, Length 8
    char    PrdaySellExecPrc    [  13];    // [double, 13.2] 전일매도체결가                  StartPos 273, Length 13
    char    PrdaySellQty        [  16];    // [long  ,   16] 전일매도수량                    StartPos 286, Length 16
    char    PrdayBuyExecPrc     [  13];    // [double, 13.2] 전일매수체결가                  StartPos 302, Length 13
    char    PrdayBuyQty         [  16];    // [long  ,   16] 전일매수수량                    StartPos 315, Length 16
    char    LoanDt              [   8];    // [string,    8] 대출일                          StartPos 331, Length 8
    char    AvrUprc             [  13];    // [double, 13.2] 평균단가                        StartPos 339, Length 13
    char    SellAbleQty         [  16];    // [long  ,   16] 매도가능수량                    StartPos 352, Length 16
    char    SellOrdQty          [  16];    // [long  ,   16] 매도주문수량                    StartPos 368, Length 16
    char    CrdayBuyExecAmt     [  16];    // [long  ,   16] 금일매수체결금액                StartPos 384, Length 16
    char    CrdaySellExecAmt    [  16];    // [long  ,   16] 금일매도체결금액                StartPos 400, Length 16
    char    PrdayBuyExecAmt     [  16];    // [long  ,   16] 전일매수체결금액                StartPos 416, Length 16
    char    PrdaySellExecAmt    [  16];    // [long  ,   16] 전일매도체결금액                StartPos 432, Length 16
    char    BalEvalAmt          [  16];    // [long  ,   16] 잔고평가금액                    StartPos 448, Length 16
    char    EvalPnl             [  16];    // [long  ,   16] 평가손익                        StartPos 464, Length 16
    char    MnyOrdAbleAmt       [  16];    // [long  ,   16] 현금주문가능금액                StartPos 480, Length 16
    char    OrdAbleAmt          [  16];    // [long  ,   16] 주문가능금액                    StartPos 496, Length 16
    char    SellUnercQty        [  16];    // [long  ,   16] 매도미체결수량                  StartPos 512, Length 16
    char    SellUnsttQty        [  16];    // [long  ,   16] 매도미결제수량                  StartPos 528, Length 16
    char    BuyUnercQty         [  16];    // [long  ,   16] 매수미체결수량                  StartPos 544, Length 16
    char    BuyUnsttQty         [  16];    // [long  ,   16] 매수미결제수량                  StartPos 560, Length 16
    char    UnsttQty            [  16];    // [long  ,   16] 미결제수량                      StartPos 576, Length 16
    char    UnercQty            [  16];    // [long  ,   16] 미체결수량                      StartPos 592, Length 16
    char    PrdayCprc           [  15];    // [double, 15.2] 전일종가                        StartPos 608, Length 15
    char    PchsAmt             [  16];    // [long  ,   16] 매입금액                        StartPos 623, Length 16
    char    RegMktCode          [   2];    // [string,    2] 등록시장코드                    StartPos 639, Length 2
    char    LoanDtlClssCode     [   2];    // [string,    2] 대출상세분류코드                StartPos 641, Length 2
    char    DpspdgLoanQty       [  16];    // [long  ,   16] 예탁담보대출수량                StartPos 643, Length 16
} CSPAQ12300OutBlock3;

//------------------------------------------------------------------------------
// 현물 계좌 주문 체결 내역 조회 (CSPAQ13700)
//------------------------------------------------------------------------------

typedef struct {
    char    RecCnt              [   5];    // [long  ,    5] 레코드갯수                      StartPos 0, Length 5
    char    AcntNo              [  20];    // [string,   20] 계좌번호                        StartPos 5, Length 20
    char    InptPwd             [   8];    // [string,    8] 입력비밀번호                    StartPos 25, Length 8
    char    OrdMktCode          [   2];    // [string,    2] 주문시장코드                    StartPos 33, Length 2
    char    BnsTpCode           [   1];    // [string,    1] 매매구분                        StartPos 35, Length 1
    char    IsuNo               [  12];    // [string,   12] 종목번호                        StartPos 36, Length 12
    char    ExecYn              [   1];    // [string,    1] 체결여부                        StartPos 48, Length 1
    char    OrdDt               [   8];    // [string,    8] 주문일                          StartPos 49, Length 8
    char    SrtOrdNo2           [  10];    // [long  ,   10] 시작주문번호2                   StartPos 57, Length 10
    char    BkseqTpCode         [   1];    // [string,    1] 역순구분                        StartPos 67, Length 1
    char    OrdPtnCode          [   2];    // [string,    2] 주문유형코드                    StartPos 68, Length 2
} CSPAQ13700InBlock1;

typedef struct {
    char    RecCnt              [   5];    // [long  ,    5] 레코드갯수                      StartPos 0, Length 5
    char    AcntNo              [  20];    // [string,   20] 계좌번호                        StartPos 5, Length 20
    char    InptPwd             [   8];    // [string,    8] 입력비밀번호                    StartPos 25, Length 8
    char    OrdMktCode          [   2];    // [string,    2] 주문시장코드                    StartPos 33, Length 2
    char    BnsTpCode           [   1];    // [string,    1] 매매구분                        StartPos 35, Length 1
    char    IsuNo               [  12];    // [string,   12] 종목번호                        StartPos 36, Length 12
    char    ExecYn              [   1];    // [string,    1] 체결여부                        StartPos 48, Length 1
    char    OrdDt               [   8];    // [string,    8] 주문일                          StartPos 49, Length 8
    char    SrtOrdNo2           [  10];    // [long  ,   10] 시작주문번호2                   StartPos 57, Length 10
    char    BkseqTpCode         [   1];    // [string,    1] 역순구분                        StartPos 67, Length 1
    char    OrdPtnCode          [   2];    // [string,    2] 주문유형코드                    StartPos 68, Length 2
} CSPAQ13700OutBlock1;

typedef struct {
    char    RecCnt              [   5];    // [long  ,    5] 레코드갯수                      StartPos 0, Length 5
    char    SellExecAmt         [  16];    // [long  ,   16] 매도체결금액                    StartPos 5, Length 16
    char    BuyExecAmt          [  16];    // [long  ,   16] 매수체결금액                    StartPos 21, Length 16
    char    SellExecQty         [  16];    // [long  ,   16] 매도체결수량                    StartPos 37, Length 16
    char    BuyExecQty          [  16];    // [long  ,   16] 매수체결수량                    StartPos 53, Length 16
    char    SellOrdQty          [  16];    // [long  ,   16] 매도주문수량                    StartPos 69, Length 16
    char    BuyOrdQty           [  16];    // [long  ,   16] 매수주문수량                    StartPos 85, Length 16
} CSPAQ13700OutBlock2;

typedef struct {
    char    OrdDt               [   8];    // [string,    8] 주문일                          StartPos 0, Length 8
    char    MgmtBrnNo           [   3];    // [string,    3] 관리지점번호                    StartPos 8, Length 3
    char    OrdMktCode          [   2];    // [string,    2] 주문시장코드                    StartPos 11, Length 2
    char    OrdNo               [  10];    // [long  ,   10] 주문번호                        StartPos 13, Length 10
    char    OrgOrdNo            [  10];    // [long  ,   10] 원주문번호                      StartPos 23, Length 10
    char    IsuNo               [  12];    // [string,   12] 종목번호                        StartPos 33, Length 12
    char    IsuNm               [  40];    // [string,   40] 종목명                          StartPos 45, Length 40
    char    BnsTpCode           [   1];    // [string,    1] 매매구분                        StartPos 85, Length 1
    char    BnsTpNm             [  10];    // [string,   10] 매매구분                        StartPos 86, Length 10
    char    OrdPtnCode          [   2];    // [string,    2] 주문유형코드                    StartPos 96, Length 2
    char    OrdPtnNm            [  40];    // [string,   40] 주문유형명                      StartPos 98, Length 40
    char    OrdTrxPtnCode       [   9];    // [long  ,    9] 주문처리유형코드                StartPos 138, Length 9
    char    OrdTrxPtnNm         [  50];    // [string,   50] 주문처리유형명                  StartPos 147, Length 50
    char    MrcTpCode           [   1];    // [string,    1] 정정취소구분                    StartPos 197, Length 1
    char    MrcTpNm             [  10];    // [string,   10] 정정취소구분명                  StartPos 198, Length 10
    char    MrcQty              [  16];    // [long  ,   16] 정정취소수량                    StartPos 208, Length 16
    char    MrcAbleQty          [  16];    // [long  ,   16] 정정취소가능수량                StartPos 224, Length 16
    char    OrdQty              [  16];    // [long  ,   16] 주문수량                        StartPos 240, Length 16
    char    OrdPrc              [  15];    // [double, 15.2] 주문가격                        StartPos 256, Length 15
    char    ExecQty             [  16];    // [long  ,   16] 체결수량                        StartPos 271, Length 16
    char    ExecPrc             [  15];    // [double, 15.2] 체결가                          StartPos 287, Length 15
    char    ExecTrxTime         [   9];    // [string,    9] 체결처리시각                    StartPos 302, Length 9
    char    LastExecTime        [   9];    // [string,    9] 최종체결시각                    StartPos 311, Length 9
    char    OrdprcPtnCode       [   2];    // [string,    2] 호가유형코드                    StartPos 320, Length 2
    char    OrdprcPtnNm         [  40];    // [string,   40] 호가유형명                      StartPos 322, Length 40
    char    OrdCndiTpCode       [   1];    // [string,    1] 주문조건구분                    StartPos 362, Length 1
    char    AllExecQty          [  16];    // [long  ,   16] 전체체결수량                    StartPos 363, Length 16
    char    RegCommdaCode       [   2];    // [string,    2] 통신매체코드                    StartPos 379, Length 2
    char    CommdaNm            [  40];    // [string,   40] 통신매체명                      StartPos 381, Length 40
    char    MbrNo               [   3];    // [string,    3] 회원번호                        StartPos 421, Length 3
    char    RsvOrdYn            [   1];    // [string,    1] 예약주문여부                    StartPos 424, Length 1
    char    LoanDt              [   8];    // [string,    8] 대출일                          StartPos 425, Length 8
    char    OrdTime             [   9];    // [string,    9] 주문시각                        StartPos 433, Length 9
    char    OpDrtnNo            [  12];    // [string,   12] 운용지시번호                    StartPos 442, Length 12
    char    OdrrId              [  16];    // [string,   16] 주문자ID                        StartPos 454, Length 16
} CSPAQ13700OutBlock3;

//------------------------------------------------------------------------------
//  현물 계좌 예수금/주문가능금 총평가 (CSPAQ22200)
//------------------------------------------------------------------------------

typedef struct {
    char    RecCnt              [   5];    // [long  ,    5] 레코드갯수                      StartPos 0, Length 5
    char    MgmtBrnNo           [   3];    // [string,    3] 관리지점번호                    StartPos 5, Length 3
    char    AcntNo              [  20];    // [string,   20] 계좌번호                        StartPos 8, Length 20
    char    Pwd                 [   8];    // [string,    8] 비밀번호                        StartPos 28, Length 8
    char    BalCreTp            [   1];    // [string,    1] 잔고생성구분                    StartPos 36, Length 1
} CSPAQ22200InBlock1;

typedef struct {
    char    RecCnt              [   5];    // [long  ,    5] 레코드갯수                      StartPos 0, Length 5
    char    MgmtBrnNo           [   3];    // [string,    3] 관리지점번호                    StartPos 5, Length 3
    char    AcntNo              [  20];    // [string,   20] 계좌번호                        StartPos 8, Length 20
    char    Pwd                 [   8];    // [string,    8] 비밀번호                        StartPos 28, Length 8
    char    BalCreTp            [   1];    // [string,    1] 잔고생성구분                    StartPos 36, Length 1
} CSPAQ22200OutBlock1;

typedef struct {
    char    RecCnt              [   5];    // [long  ,    5] 레코드갯수                      StartPos 0, Length 5
    char    BrnNm               [  40];    // [string,   40] 지점명                          StartPos 5, Length 40
    char    AcntNm              [  40];    // [string,   40] 계좌명                          StartPos 45, Length 40
    char    MnyOrdAbleAmt       [  16];    // [long  ,   16] 현금주문가능금액                StartPos 85, Length 16
    char    SubstOrdAbleAmt     [  16];    // [long  ,   16] 대용주문가능금액                StartPos 101, Length 16
    char    SeOrdAbleAmt        [  16];    // [long  ,   16] 거래소금액                      StartPos 117, Length 16
    char    KdqOrdAbleAmt       [  16];    // [long  ,   16] 코스닥금액                      StartPos 133, Length 16
    char    CrdtPldgOrdAmt      [  16];    // [long  ,   16] 신용담보주문금액                StartPos 149, Length 16
    char    MgnRat100pctOrdAbleAmt[  16];    // [long  ,   16] 증거금률100퍼센트주문가능금액   StartPos 165, Length 16
    char    MgnRat35ordAbleAmt  [  16];    // [long  ,   16] 증거금률35%주문가능금액         StartPos 181, Length 16
    char    MgnRat50ordAbleAmt  [  16];    // [long  ,   16] 증거금률50%주문가능금액         StartPos 197, Length 16
    char    CrdtOrdAbleAmt      [  16];    // [long  ,   16] 신용주문가능금액                StartPos 213, Length 16
    char    Dps                 [  16];    // [long  ,   16] 예수금                          StartPos 229, Length 16
    char    SubstAmt            [  16];    // [long  ,   16] 대용금액                        StartPos 245, Length 16
    char    MgnMny              [  16];    // [long  ,   16] 증거금현금                      StartPos 261, Length 16
    char    MgnSubst            [  16];    // [long  ,   16] 증거금대용                      StartPos 277, Length 16
    char    D1Dps               [  16];    // [long  ,   16] D1예수금                        StartPos 293, Length 16
    char    D2Dps               [  16];    // [long  ,   16] D2예수금                        StartPos 309, Length 16
    char    RcvblAmt            [  16];    // [long  ,   16] 미수금액                        StartPos 325, Length 16
    char    D1ovdRepayRqrdAmt   [  16];    // [long  ,   16] D1연체변제소요금액              StartPos 341, Length 16
    char    D2ovdRepayRqrdAmt   [  16];    // [long  ,   16] D2연체변제소요금액              StartPos 357, Length 16
    char    MloanAmt            [  16];    // [long  ,   16] 융자금액                        StartPos 373, Length 16
    char    ChgAfPldgRat        [   9];    // [double,  9.3] 변경후담보비율                  StartPos 389, Length 9
    char    RqrdPldgAmt         [  16];    // [long  ,   16] 소요담보금액                    StartPos 398, Length 16
    char    PdlckAmt            [  16];    // [long  ,   16] 담보부족금액                    StartPos 414, Length 16
    char    OrgPldgSumAmt       [  16];    // [long  ,   16] 원담보합계금액                  StartPos 430, Length 16
    char    SubPldgSumAmt       [  16];    // [long  ,   16] 부담보합계금액                  StartPos 446, Length 16
    char    CrdtPldgAmtMny      [  16];    // [long  ,   16] 신용담보금현금                  StartPos 462, Length 16
    char    CrdtPldgSubstAmt    [  16];    // [long  ,   16] 신용담보대용금액                StartPos 478, Length 16
    char    Imreq               [  16];    // [long  ,   16] 신용설정보증금                  StartPos 494, Length 16
    char    CrdtPldgRuseAmt     [  16];    // [long  ,   16] 신용담보재사용금액              StartPos 510, Length 16
    char    DpslRestrcAmt       [  16];    // [long  ,   16] 처분제한금액                    StartPos 526, Length 16
    char    PrdaySellAdjstAmt   [  16];    // [long  ,   16] 전일매도정산금액                StartPos 542, Length 16
    char    PrdayBuyAdjstAmt    [  16];    // [long  ,   16] 전일매수정산금액                StartPos 558, Length 16
    char    CrdaySellAdjstAmt   [  16];    // [long  ,   16] 금일매도정산금액                StartPos 574, Length 16
    char    CrdayBuyAdjstAmt    [  16];    // [long  ,   16] 금일매수정산금액                StartPos 590, Length 16
    char    CslLoanAmtdt1       [  16];    // [long  ,   16] 매도대금담보대출금액            StartPos 606, Length 16
} CSPAQ22200OutBlock2;

typedef struct {
    CSPAQ22200OutBlock1	outBlock1;
    CSPAQ22200OutBlock2	outBlock2;
} CSPAQ22200OutBlock;

//------------------------------------------------------------------------------
//  현물 계좌 기간별 수익률 상세 (FOCCQ33600)
//------------------------------------------------------------------------------

typedef struct {
    char    RecCnt              [   5];    // [long  ,    5] 레코드갯수                      StartPos 0, Length 5
    char    AcntNo              [  20];    // [string,   20] 계좌번호                        StartPos 5, Length 20
    char    Pwd                 [   8];    // [string,    8] 비밀번호                        StartPos 25, Length 8
    char    QrySrtDt            [   8];    // [string,    8] 조회시작일                      StartPos 33, Length 8
    char    QryEndDt            [   8];    // [string,    8] 조회종료일                      StartPos 41, Length 8
    char    TermTp              [   1];    // [string,    1] 기간구분                        StartPos 49, Length 1
} FOCCQ33600InBlock1;

typedef struct {
    char    RecCnt              [   5];    // [long  ,    5] 레코드갯수                      StartPos 0, Length 5
    char    AcntNo              [  20];    // [string,   20] 계좌번호                        StartPos 5, Length 20
    char    Pwd                 [   8];    // [string,    8] 비밀번호                        StartPos 25, Length 8
    char    QrySrtDt            [   8];    // [string,    8] 조회시작일                      StartPos 33, Length 8
    char    QryEndDt            [   8];    // [string,    8] 조회종료일                      StartPos 41, Length 8
    char    TermTp              [   1];    // [string,    1] 기간구분                        StartPos 49, Length 1
} FOCCQ33600OutBlock1;

typedef struct {
    char    RecCnt              [   5];    // [long  ,    5] 레코드갯수                      StartPos 0, Length 5
    char    AcntNm              [  40];    // [string,   40] 계좌명                          StartPos 5, Length 40
    char    BnsctrAmt           [  16];    // [long  ,   16] 매매약정금액                    StartPos 45, Length 16
    char    MnyinAmt            [  16];    // [long  ,   16] 입금금액                        StartPos 61, Length 16
    char    MnyoutAmt           [  16];    // [long  ,   16] 출금금액                        StartPos 77, Length 16
    char    InvstAvrbalPramt    [  16];    // [long  ,   16] 투자원금평잔금액                StartPos 93, Length 16
    char    InvstPlAmt          [  16];    // [long  ,   16] 투자손익금액                    StartPos 109, Length 16
    char    InvstErnrat         [   9];    // [double,  9.2] 투자수익률                      StartPos 125, Length 9
} FOCCQ33600OutBlock2;

typedef struct {
    char    BaseDt              [   8];    // [string,    8] 기준일                          StartPos 0, Length 8
    char    FdEvalAmt           [  19];    // [long  ,   19] 기초평가금액                    StartPos 8, Length 19
    char    EotEvalAmt          [  19];    // [long  ,   19] 기말평가금액                    StartPos 27, Length 19
    char    InvstAvrbalPramt    [  16];    // [long  ,   16] 투자원금평잔금액                StartPos 46, Length 16
    char    BnsctrAmt           [  16];    // [long  ,   16] 매매약정금액                    StartPos 62, Length 16
    char    MnyinSecinAmt       [  16];    // [long  ,   16] 입금고액                        StartPos 78, Length 16
    char    MnyoutSecoutAmt     [  16];    // [long  ,   16] 출금고액                        StartPos 94, Length 16
    char    EvalPnlAmt          [  16];    // [long  ,   16] 평가손익금액                    StartPos 110, Length 16
    char    TermErnrat          [  11];    // [double, 11.3] 기간수익률                      StartPos 126, Length 11
    char    Idx                 [  13];    // [double, 13.2] 지수                            StartPos 137, Length 13
} FOCCQ33600OutBlock3;

//------------------------------------------------------------------------------
// 현물 당일 매매일지 수수료 (t0150)
//------------------------------------------------------------------------------

typedef struct {
    char    accno               [  11];    // [string,   11] 계좌번호                        StartPos 0, Length 11
    char    cts_medosu          [   1];    // [string,    1] CTS_매매구분                    StartPos 11, Length 1
    char    cts_expcode         [  12];    // [string,   12] CTS_종목번호                    StartPos 12, Length 12
    char    cts_price           [   9];    // [string,    9] CTS_단가                        StartPos 24, Length 9
    char    cts_middiv          [   2];    // [string,    2] CTS_매체                        StartPos 33, Length 2
} T0150InBlock;

typedef struct {
    char    mdqty               [   9];    // [long  ,    9] 매도수량                        StartPos 0, Length 9
    char    mdamt               [  18];    // [long  ,   18] 매도약정금액                    StartPos 9, Length 18
    char    mdfee               [  18];    // [long  ,   18] 매도수수료                      StartPos 27, Length 18
    char    mdtax               [  18];    // [long  ,   18] 매도거래세                      StartPos 45, Length 18
    char    mdargtax            [  18];    // [long  ,   18] 매도농특세                      StartPos 63, Length 18
    char    tmdtax              [  18];    // [long  ,   18] 매도제비용합                    StartPos 81, Length 18
    char    mdadjamt            [  18];    // [long  ,   18] 매도정산금액                    StartPos 99, Length 18
    char    msqty               [   9];    // [long  ,    9] 매수수량                        StartPos 117, Length 9
    char    msamt               [  18];    // [long  ,   18] 매수약정금액                    StartPos 126, Length 18
    char    msfee               [  18];    // [long  ,   18] 매수수수료                      StartPos 144, Length 18
    char    tmstax              [  18];    // [long  ,   18] 매수제비용합                    StartPos 162, Length 18
    char    msadjamt            [  18];    // [long  ,   18] 매수정산금액                    StartPos 180, Length 18
    char    tqty                [   9];    // [long  ,    9] 합계수량                        StartPos 198, Length 9
    char    tamt                [  18];    // [long  ,   18] 합계약정금액                    StartPos 207, Length 18
    char    tfee                [  18];    // [long  ,   18] 합계수수료                      StartPos 225, Length 18
    char    tottax              [  18];    // [long  ,   18] 합계거래세                      StartPos 243, Length 18
    char    targtax             [  18];    // [long  ,   18] 합계농특세                      StartPos 261, Length 18
    char    ttax                [  18];    // [long  ,   18] 합계제비용합                    StartPos 279, Length 18
    char    tadjamt             [  18];    // [long  ,   18] 합계정산금액                    StartPos 297, Length 18
    char    cts_medosu          [   1];    // [string,    1] CTS_매매구분                    StartPos 315, Length 1
    char    cts_expcode         [  12];    // [string,   12] CTS_종목번호                    StartPos 316, Length 12
    char    cts_price           [   9];    // [string,    9] CTS_단가                        StartPos 328, Length 9
    char    cts_middiv          [   2];    // [string,    2] CTS_매체                        StartPos 337, Length 2
} T0150OutBlock;

typedef struct {
    char    medosu              [  10];    // [string,   10] 매매구분                        StartPos 0, Length 10
    char    expcode             [  12];    // [string,   12] 종목번호                        StartPos 10, Length 12
    char    qty                 [   9];    // [long  ,    9] 수량                            StartPos 22, Length 9
    char    price               [   9];    // [long  ,    9] 단가                            StartPos 31, Length 9
    char    amt                 [  18];    // [long  ,   18] 약정금액                        StartPos 40, Length 18
    char    fee                 [  18];    // [long  ,   18] 수수료                          StartPos 58, Length 18
    char    tax                 [  18];    // [long  ,   18] 거래세                          StartPos 76, Length 18
    char    argtax              [  18];    // [long  ,   18] 농특세                          StartPos 94, Length 18
    char    adjamt              [  18];    // [long  ,   18] 정산금액                        StartPos 112, Length 18
    char    middiv              [  20];    // [string,   20] 매체                            StartPos 130, Length 20
} T0150OutBlock1;

//------------------------------------------------------------------------------
// 주식 당일/전일 매매일지 수수료 (t0151)
//------------------------------------------------------------------------------

typedef struct {
    char    date                [   8];    // [string,    8] 일자                            StartPos 0, Length 8
    char    accno               [  11];    // [string,   11] 계좌번호                        StartPos 8, Length 11
    char    cts_medosu          [   1];    // [string,    1] CTS_매매구분                    StartPos 19, Length 1
    char    cts_expcode         [  12];    // [string,   12] CTS_종목번호                    StartPos 20, Length 12
    char    cts_price           [   9];    // [string,    9] CTS_단가                        StartPos 32, Length 9
    char    cts_middiv          [   2];    // [string,    2] CTS_매체                        StartPos 41, Length 2
} T0151InBlock;

typedef struct {
    char    mdqty               [   9];    // [long  ,    9] 매도수량                        StartPos 0, Length 9
    char    mdamt               [  18];    // [long  ,   18] 매도약정금액                    StartPos 9, Length 18
    char    mdfee               [  18];    // [long  ,   18] 매도수수료                      StartPos 27, Length 18
    char    mdtax               [  18];    // [long  ,   18] 매도거래세                      StartPos 45, Length 18
    char    mdargtax            [  18];    // [long  ,   18] 매도농특세                      StartPos 63, Length 18
    char    tmdtax              [  18];    // [long  ,   18] 매도제비용합                    StartPos 81, Length 18
    char    mdadjamt            [  18];    // [long  ,   18] 매도정산금액                    StartPos 99, Length 18
    char    msqty               [   9];    // [long  ,    9] 매수수량                        StartPos 117, Length 9
    char    msamt               [  18];    // [long  ,   18] 매수약정금액                    StartPos 126, Length 18
    char    msfee               [  18];    // [long  ,   18] 매수수수료                      StartPos 144, Length 18
    char    tmstax              [  18];    // [long  ,   18] 매수제비용합                    StartPos 162, Length 18
    char    msadjamt            [  18];    // [long  ,   18] 매수정산금액                    StartPos 180, Length 18
    char    tqty                [   9];    // [long  ,    9] 합계수량                        StartPos 198, Length 9
    char    tamt                [  18];    // [long  ,   18] 합계약정금액                    StartPos 207, Length 18
    char    tfee                [  18];    // [long  ,   18] 합계수수료                      StartPos 225, Length 18
    char    tottax              [  18];    // [long  ,   18] 합계거래세                      StartPos 243, Length 18
    char    targtax             [  18];    // [long  ,   18] 합계농특세                      StartPos 261, Length 18
    char    ttax                [  18];    // [long  ,   18] 합계제비용합                    StartPos 279, Length 18
    char    tadjamt             [  18];    // [long  ,   18] 합계정산금액                    StartPos 297, Length 18
    char    cts_medosu          [   1];    // [string,    1] CTS_매매구분                    StartPos 315, Length 1
    char    cts_expcode         [  12];    // [string,   12] CTS_종목번호                    StartPos 316, Length 12
    char    cts_price           [   9];    // [string,    9] CTS_단가                        StartPos 328, Length 9
    char    cts_middiv          [   2];    // [string,    2] CTS_매체                        StartPos 337, Length 2
} T0151OutBlock;

typedef struct {
    char    medosu              [  10];    // [string,   10] 매매구분                        StartPos 0, Length 10
    char    expcode             [  12];    // [string,   12] 종목번호                        StartPos 10, Length 12
    char    qty                 [   9];    // [long  ,    9] 수량                            StartPos 22, Length 9
    char    price               [   9];    // [long  ,    9] 단가                            StartPos 31, Length 9
    char    amt                 [  18];    // [long  ,   18] 약정금액                        StartPos 40, Length 18
    char    fee                 [  18];    // [long  ,   18] 수수료                          StartPos 58, Length 18
    char    tax                 [  18];    // [long  ,   18] 거래세                          StartPos 76, Length 18
    char    argtax              [  18];    // [long  ,   18] 농특세                          StartPos 94, Length 18
    char    adjamt              [  18];    // [long  ,   18] 정산금액                        StartPos 112, Length 18
    char    middiv              [  20];    // [string,   20] 매체                            StartPos 130, Length 20
} T0151OutBlock1;

//------------------------------------------------------------------------------
// 시간 조회 (t0167)
//------------------------------------------------------------------------------
typedef struct {
    char    date[8];                            // 일자(YYYYMMDD)
    char    time[12];                           // 시간(HHMMSSssssss)
} T0167OutBlock;

//------------------------------------------------------------------------------
// 현물 체결 / 미체결 (t0425)
//------------------------------------------------------------------------------
typedef struct {
    char    accno               [  11];    char    _accno               ;    // [string,   11] 계좌번호                        StartPos 0, Length 11
    char    passwd              [   8];    char    _passwd              ;    // [string,    8] 비밀번호                        StartPos 12, Length 8
    char    expcode             [  12];    char    _expcode             ;    // [string,   12] 종목번호                        StartPos 21, Length 12
    char    chegb               [   1];    char    _chegb               ;    // [string,    1] 체결구분                        StartPos 34, Length 1
    char    medosu              [   1];    char    _medosu              ;    // [string,    1] 매매구분                        StartPos 36, Length 1
    char    sortgb              [   1];    char    _sortgb              ;    // [string,    1] 정렬순서                        StartPos 38, Length 1
    char    cts_ordno           [  10];    char    _cts_ordno           ;    // [string,   10] 주문번호                        StartPos 40, Length 10
} T0425InBlock;

typedef struct {
    char    tqty                [  18];    char    _tqty                ;    // [long  ,   18] 총주문수량                      StartPos 0, Length 18
    char    tcheqty             [  18];    char    _tcheqty             ;    // [long  ,   18] 총체결수량                      StartPos 19, Length 18
    char    tordrem             [  18];    char    _tordrem             ;    // [long  ,   18] 총미체결수량                    StartPos 38, Length 18
    char    cmss                [  18];    char    _cmss                ;    // [long  ,   18] 추정수수료                      StartPos 57, Length 18
    char    tamt                [  18];    char    _tamt                ;    // [long  ,   18] 총주문금액                      StartPos 76, Length 18
    char    tmdamt              [  18];    char    _tmdamt              ;    // [long  ,   18] 총매도체결금액                  StartPos 95, Length 18
    char    tmsamt              [  18];    char    _tmsamt              ;    // [long  ,   18] 총매수체결금액                  StartPos 114, Length 18
    char    tax                 [  18];    char    _tax                 ;    // [long  ,   18] 추정제세금                      StartPos 133, Length 18
    char    cts_ordno           [  10];    char    _cts_ordno           ;    // [string,   10] 주문번호                        StartPos 152, Length 10
} T0425OutBlock;

typedef struct {
    char    ordno               [  10];    char    _ordno               ;    // [long  ,   10] 주문번호                        StartPos 0, Length 10
    char    expcode             [  12];    char    _expcode             ;    // [string,   12] 종목번호                        StartPos 11, Length 12
    char    medosu              [  10];    char    _medosu              ;    // [string,   10] 구분                            StartPos 24, Length 10
    char    qty                 [   9];    char    _qty                 ;    // [long  ,    9] 주문수량                        StartPos 35, Length 9
    char    price               [   9];    char    _price               ;    // [long  ,    9] 주문가격                        StartPos 45, Length 9
    char    cheqty              [   9];    char    _cheqty              ;    // [long  ,    9] 체결수량                        StartPos 55, Length 9
    char    cheprice            [   9];    char    _cheprice            ;    // [long  ,    9] 체결가격                        StartPos 65, Length 9
    char    ordrem              [   9];    char    _ordrem              ;    // [long  ,    9] 미체결잔량                      StartPos 75, Length 9
    char    cfmqty              [   9];    char    _cfmqty              ;    // [long  ,    9] 확인수량                        StartPos 85, Length 9
    char    status              [  10];    char    _status              ;    // [string,   10] 상태                            StartPos 95, Length 10
    char    orgordno            [  10];    char    _orgordno            ;    // [long  ,   10] 원주문번호                      StartPos 106, Length 10
    char    ordgb               [  20];    char    _ordgb               ;    // [string,   20] 유형                            StartPos 117, Length 20
    char    ordtime             [   8];    char    _ordtime             ;    // [string,    8] 주문시간                        StartPos 138, Length 8
    char    ordermtd            [  10];    char    _ordermtd            ;    // [string,   10] 주문매체                        StartPos 147, Length 10
    char    sysprocseq          [  10];    char    _sysprocseq          ;    // [long  ,   10] 처리순번                        StartPos 158, Length 10
    char    hogagb              [   2];    char    _hogagb              ;    // [string,    2] 호가유형                        StartPos 169, Length 2
    char    price1              [   8];    char    _price1              ;    // [long  ,    8] 현재가                          StartPos 172, Length 8
    char    orggb               [   2];    char    _orggb               ;    // [string,    2] 주문구분                        StartPos 181, Length 2
    char    singb               [   2];    char    _singb               ;    // [string,    2] 신용구분                        StartPos 184, Length 2
    char    loandt              [   8];    char    _loandt              ;    // [string,    8] 대출일자                        StartPos 187, Length 8
} T0425OutBlock1;

//------------------------------------------------------------------------------
// 선물옵션 체결 / 미체결 (t0434)
//------------------------------------------------------------------------------
typedef struct {
    char    accno               [  11];    // [string,   11] 계좌번호                        StartPos 0, Length 11
    char    passwd              [   8];    // [string,    8] 비밀번호                        StartPos 11, Length 8
    char    expcode             [   8];    // [string,    8] 종목번호                        StartPos 19, Length 8
    char    chegb               [   1];    // [string,    1] 체결구분                        StartPos 27, Length 1
    char    sortgb              [   1];    // [string,    1] 정렬순서                        StartPos 28, Length 1
    char    cts_ordno           [   7];    // [string,    7] CTS_주문번호                    StartPos 29, Length 7
} T0434InBlock;

typedef struct {
    char    cts_ordno           [   7];    // [string,    7] CTS_주문번호                    StartPos 0, Length 7
} T0434OutBlock;

typedef struct {
    char    ordno               [   7];    // [long  ,    7] 주문번호                        StartPos 0, Length 7
    char    orgordno            [   7];    // [long  ,    7] 원주문번호                      StartPos 7, Length 7
    char    medosu              [  10];    // [string,   10] 구분                            StartPos 14, Length 10
    char    ordgb               [  20];    // [string,   20] 유형                            StartPos 24, Length 20
    char    qty                 [   9];    // [long  ,    9] 주문수량                        StartPos 44, Length 9
    char    price               [   9];    // [float ,  9.2] 주문가격                        StartPos 53, Length 9
    char    cheqty              [   9];    // [long  ,    9] 체결수량                        StartPos 62, Length 9
    char    cheprice            [   9];    // [float ,  9.2] 체결가격                        StartPos 71, Length 9
    char    ordrem              [   9];    // [long  ,    9] 미체결잔량                      StartPos 80, Length 9
    char    status              [  10];    // [string,   10] 상태                            StartPos 89, Length 10
    char    ordtime             [   8];    // [string,    8] 주문시간                        StartPos 99, Length 8
    char    ordermtd            [  10];    // [string,   10] 주문매체                        StartPos 107, Length 10
    char    expcode             [   8];    // [string,    8] 종목번호                        StartPos 117, Length 8
    char    rtcode              [   3];    // [string,    3] 사유코드                        StartPos 125, Length 3
    char    sysprocseq          [  10];    // [long  ,   10] 처리순번                        StartPos 128, Length 10
    char    hogatype            [   1];    // [string,    1] 호가타입                        StartPos 138, Length 1
} T0434OutBlock1;

//------------------------------------------------------------------------------
// 현물 현재가 호가 조회 (t1101)
//------------------------------------------------------------------------------
typedef struct {
    char    shcode[6];     char    _shcode;    //[string,    6] 단축코드   StartPos 0, Length 6
} T1101InBlock;

typedef struct {
    char hname[20];     char _hname;        //[string,   20] 한글명   StartPos 0, Length 20
    char price[8];      char _price;        //[long  ,    8] 현재가   StartPos 21, Length 8
    char sign[1];       char _sign;         //[string,    1] 전일대비구분   StartPos 30, Length 1
    char change[8];     char _change;       //[long  ,    8] 전일대비   StartPos 32, Length 8
    char diff[6];       char _diff;         //[float ,  6.2] 등락율   StartPos 41, Length 6
    char volume[12];    char _volume;       //[long  ,   12] 누적거래량   StartPos 48, Length 12
    char jnilclose[8];     char _jnilclose;    //[long  ,    8] 전일종가   StartPos 61, Length 8
    char offerho1[8];    char _offerho1;     //[long  ,    8] 매도호가1   StartPos 70, Length 8
    char bidho1[8];     char _bidho1;       //[long  ,    8] 매수호가1   StartPos 79, Length 8
    char offerrem1[12];    char _offerrem1;    //[long  ,   12] 매도호가수량1   StartPos 88, Length 12
    char bidrem1[12];    char _bidrem1;      //[long  ,   12] 매수호가수량1   StartPos 101, Length 12
    char preoffercha1[12];  char _preoffercha1; //[long  ,   12] 직전매도대비수량1   StartPos 114, Length 12
    char prebidcha1[12];    char _prebidcha1;   //[long  ,   12] 직전매수대비수량1   StartPos 127, Length 12
    char offerho2[8];    char _offerho2;     //[long  ,    8] 매도호가2   StartPos 140, Length 8
    char bidho2[8];     char _bidho2;       //[long  ,    8] 매수호가2   StartPos 149, Length 8
    char offerrem2[12];    char _offerrem2;    //[long  ,   12] 매도호가수량2   StartPos 158, Length 12
    char bidrem2[12];    char _bidrem2;      //[long  ,   12] 매수호가수량2   StartPos 171, Length 12
    char preoffercha2[12];  char _preoffercha2; //[long  ,   12] 직전매도대비수량2   StartPos 184, Length 12
    char prebidcha2[12];    char _prebidcha2;   //[long  ,   12] 직전매수대비수량2   StartPos 197, Length 12
    char offerho3[8];    char _offerho3;     //[long  ,    8] 매도호가3   StartPos 210, Length 8
    char bidho3[8];     char _bidho3;       //[long  ,    8] 매수호가3   StartPos 219, Length 8
    char offerrem3[12];    char _offerrem3;    //[long  ,   12] 매도호가수량3   StartPos 228, Length 12
    char bidrem3[12];    char _bidrem3;      //[long  ,   12] 매수호가수량3   StartPos 241, Length 12
    char preoffercha3[12];  char _preoffercha3; //[long  ,   12] 직전매도대비수량3   StartPos 254, Length 12
    char prebidcha3[12];    char _prebidcha3;   //[long  ,   12] 직전매수대비수량3   StartPos 267, Length 12
    char offerho4[8];    char _offerho4;     //[long  ,    8] 매도호가4   StartPos 280, Length 8
    char bidho4[8];     char _bidho4;       //[long  ,    8] 매수호가4   StartPos 289, Length 8
    char offerrem4[12];    char _offerrem4;    //[long  ,   12] 매도호가수량4   StartPos 298, Length 12
    char bidrem4[12];    char _bidrem4;      //[long  ,   12] 매수호가수량4   StartPos 311, Length 12
    char preoffercha4[12];  char _preoffercha4; //[long  ,   12] 직전매도대비수량4   StartPos 324, Length 12
    char prebidcha4[12];    char _prebidcha4;   //[long  ,   12] 직전매수대비수량4   StartPos 337, Length 12
    char offerho5[8];    char _offerho5;     //[long  ,    8] 매도호가5   StartPos 350, Length 8
    char bidho5[8];     char _bidho5;       //[long  ,    8] 매수호가5   StartPos 359, Length 8
    char offerrem5[12];    char _offerrem5;    //[long  ,   12] 매도호가수량5   StartPos 368, Length 12
    char bidrem5[12];    char _bidrem5;      //[long  ,   12] 매수호가수량5   StartPos 381, Length 12
    char preoffercha5[12];  char _preoffercha5; //[long  ,   12] 직전매도대비수량5   StartPos 394, Length 12
    char prebidcha5[12];    char _prebidcha5;   //[long  ,   12] 직전매수대비수량5   StartPos 407, Length 12
    char offerho6[8];    char _offerho6;     //[long  ,    8] 매도호가6   StartPos 420, Length 8
    char bidho6[8];     char _bidho6;       //[long  ,    8] 매수호가6   StartPos 429, Length 8
    char offerrem6[12];    char _offerrem6;    //[long  ,   12] 매도호가수량6   StartPos 438, Length 12
    char bidrem6[12];    char _bidrem6;      //[long  ,   12] 매수호가수량6   StartPos 451, Length 12
    char preoffercha6[12];  char _preoffercha6; //[long  ,   12] 직전매도대비수량6   StartPos 464, Length 12
    char prebidcha6[12];    char _prebidcha6;   //[long  ,   12] 직전매수대비수량6   StartPos 477, Length 12
    char offerho7[8];    char _offerho7;     //[long  ,    8] 매도호가7   StartPos 490, Length 8
    char bidho7[8];     char _bidho7;       //[long  ,    8] 매수호가7   StartPos 499, Length 8
    char offerrem7[12];    char _offerrem7;    //[long  ,   12] 매도호가수량7   StartPos 508, Length 12
    char bidrem7[12];    char _bidrem7;      //[long  ,   12] 매수호가수량7   StartPos 521, Length 12
    char preoffercha7[12];  char _preoffercha7; //[long  ,   12] 직전매도대비수량7   StartPos 534, Length 12
    char prebidcha7[12];    char _prebidcha7;   //[long  ,   12] 직전매수대비수량7   StartPos 547, Length 12
    char offerho8[8];    char _offerho8;     //[long  ,    8] 매도호가8   StartPos 560, Length 8
    char bidho8[8];     char _bidho8;       //[long  ,    8] 매수호가8   StartPos 569, Length 8
    char offerrem8[12];    char _offerrem8;    //[long  ,   12] 매도호가수량8   StartPos 578, Length 12
    char bidrem8[12];    char _bidrem8;      //[long  ,   12] 매수호가수량8   StartPos 591, Length 12
    char preoffercha8[12];  char _preoffercha8; //[long  ,   12] 직전매도대비수량8   StartPos 604, Length 12
    char prebidcha8[12];    char _prebidcha8;   //[long  ,   12] 직전매수대비수량8   StartPos 617, Length 12
    char offerho9[8];    char _offerho9;     //[long  ,    8] 매도호가9   StartPos 630, Length 8
    char bidho9[8];     char _bidho9;       //[long  ,    8] 매수호가9   StartPos 639, Length 8
    char offerrem9[12];    char _offerrem9;    //[long  ,   12] 매도호가수량9   StartPos 648, Length 12
    char bidrem9[12];    char _bidrem9;      //[long  ,   12] 매수호가수량9   StartPos 661, Length 12
    char preoffercha9[12];  char _preoffercha9; //[long  ,   12] 직전매도대비수량9   StartPos 674, Length 12
    char prebidcha9[12];    char _prebidcha9;   //[long  ,   12] 직전매수대비수량9   StartPos 687, Length 12
    char offerho10[8];     char _offerho10;    //[long  ,    8] 매도호가10   StartPos 700, Length 8
    char bidho10[8];    char _bidho10;      //[long  ,    8] 매수호가10   StartPos 709, Length 8
    char offerrem10[12];    char _offerrem10;   //[long  ,   12] 매도호가수량10   StartPos 718, Length 12
    char bidrem10[12];     char _bidrem10;     //[long  ,   12] 매수호가수량10   StartPos 731, Length 12
    char preoffercha10[12]; char _preoffercha10; //[long  ,   12] 직전매도대비수량10   StartPos 744, Length 12
    char prebidcha10[12];   char _prebidcha10;  //[long  ,   12] 직전매수대비수량10   StartPos 757, Length 12
    char offer[12];     char _offer;        //[long  ,   12] 매도호가수량합   StartPos 770, Length 12
    char bid[12];       char _bid;          //[long  ,   12] 매수호가수량합   StartPos 783, Length 12
    char preoffercha[12];   char _preoffercha;  //[long  ,   12] 직전매도대비수량합   StartPos 796, Length 12
    char prebidcha[12];    char _prebidcha;    //[long  ,   12] 직전매수대비수량합   StartPos 809, Length 12
    char hotime[8];     char _hotime;       //[string,    8] 수신시간   StartPos 822, Length 8
    char yeprice[8];    char _yeprice;      //[long  ,    8] 예상체결가격   StartPos 831, Length 8
    char yevolume[12];     char _yevolume;     //[long  ,   12] 예상체결수량   StartPos 840, Length 12
    char yesign[1];     char _yesign;       //[string,    1] 예상체결전일구분   StartPos 853, Length 1
    char yechange[8];    char _yechange;     //[long  ,    8] 예상체결전일대비   StartPos 855, Length 8
    char yediff[6];     char _yediff;       //[float ,  6.2] 예상체결등락율   StartPos 864, Length 6
    char tmoffer[12];    char _tmoffer;      //[long  ,   12] 시간외매도잔량   StartPos 871, Length 12
    char tmbid[12];     char _tmbid;        //[long  ,   12] 시간외매수잔량   StartPos 884, Length 12
    char ho_status[1];     char _ho_status;    //[string,    1] 동시구분   StartPos 897, Length 1
    char shcode[6];     char _shcode;       //[string,    6] 단축코드   StartPos 899, Length 6
    char uplmtprice[8];    char _uplmtprice;   //[long  ,    8] 상한가   StartPos 906, Length 8
    char dnlmtprice[8];    char _dnlmtprice;   //[long  ,    8] 하한가   StartPos 915, Length 8
    char open[8];       char _open;         //[long  ,    8] 시가   StartPos 924, Length 8
    char high[8];       char _high;         //[long  ,    8] 고가   StartPos 933, Length 8
    char low[8];        char _low;          //[long  ,    8] 저가   StartPos 942, Length 8
} T1101OutBlock;

//------------------------------------------------------------------------------
// 현물 현재가 시세 조회 (t1102)
//------------------------------------------------------------------------------
typedef struct {
    char    shcode              [   6];    char    _shcode              ;    // [string,    6] 단축코드                        StartPos 0, Length 6
} T1102InBlock;

typedef struct {
   char    hname               [  20];    char    _hname               ;    // [string,   20] 한글명                          StartPos 0, Length 20
   char    price               [   8];    char    _price               ;    // [long  ,    8] 현재가                          StartPos 21, Length 8
   char    sign                [   1];    char    _sign                ;    // [string,    1] 전일대비구분                    StartPos 30, Length 1
   char    change              [   8];    char    _change              ;    // [long  ,    8] 전일대비                        StartPos 32, Length 8
   char    diff                [   6];    char    _diff                ;    // [float ,  6.2] 등락율                          StartPos 41, Length 6
   char    volume              [  12];    char    _volume              ;    // [long  ,   12] 누적거래량                      StartPos 48, Length 12
   char    recprice            [   8];    char    _recprice            ;    // [long  ,    8] 기준가(평가가격)                StartPos 61, Length 8
   char    avg                 [   8];    char    _avg                 ;    // [long  ,    8] 가중평균                        StartPos 70, Length 8
   char    uplmtprice          [   8];    char    _uplmtprice          ;    // [long  ,    8] 상한가(최고호가가격)            StartPos 79, Length 8
   char    dnlmtprice          [   8];    char    _dnlmtprice          ;    // [long  ,    8] 하한가(최저호가가격)            StartPos 88, Length 8
   char    jnilvolume          [  12];    char    _jnilvolume          ;    // [long  ,   12] 전일거래량                      StartPos 97, Length 12
   char    volumediff          [  12];    char    _volumediff          ;    // [long  ,   12] 거래량차                        StartPos 110, Length 12
   char    open                [   8];    char    _open                ;    // [long  ,    8] 시가                            StartPos 123, Length 8
   char    opentime            [   6];    char    _opentime            ;    // [string,    6] 시가시간                        StartPos 132, Length 6
   char    high                [   8];    char    _high                ;    // [long  ,    8] 고가                            StartPos 139, Length 8
   char    hightime            [   6];    char    _hightime            ;    // [string,    6] 고가시간                        StartPos 148, Length 6
   char    low                 [   8];    char    _low                 ;    // [long  ,    8] 저가                            StartPos 155, Length 8
   char    lowtime             [   6];    char    _lowtime             ;    // [string,    6] 저가시간                        StartPos 164, Length 6
   char    high52w             [   8];    char    _high52w             ;    // [long  ,    8] 52최고가                        StartPos 171, Length 8
   char    high52wdate         [   8];    char    _high52wdate         ;    // [string,    8] 52최고가일                      StartPos 180, Length 8
   char    low52w              [   8];    char    _low52w              ;    // [long  ,    8] 52최저가                        StartPos 189, Length 8
   char    low52wdate          [   8];    char    _low52wdate          ;    // [string,    8] 52최저가일                      StartPos 198, Length 8
   char    exhratio            [   6];    char    _exhratio            ;    // [float ,  6.2] 소진율                          StartPos 207, Length 6
   char    per                 [   6];    char    _per                 ;    // [float ,  6.2] PER                             StartPos 214, Length 6
   char    pbrx                [   6];    char    _pbrx                ;    // [float ,  6.2] PBRX                            StartPos 221, Length 6
   char    listing             [  12];    char    _listing             ;    // [long  ,   12] 상장주식수(천)                  StartPos 228, Length 12
   char    jkrate              [   8];    char    _jkrate              ;    // [long  ,    8] 증거금율                        StartPos 241, Length 8
   char    memedan             [   5];    char    _memedan             ;    // [string,    5] 수량단위                        StartPos 250, Length 5
   char    offernocd1          [   3];    char    _offernocd1          ;    // [string,    3] 매도증권사코드1                 StartPos 256, Length 3
   char    bidnocd1            [   3];    char    _bidnocd1            ;    // [string,    3] 매수증권사코드1                 StartPos 260, Length 3
   char    offerno1            [   6];    char    _offerno1            ;    // [string,    6] 매도증권사명1                   StartPos 264, Length 6
   char    bidno1              [   6];    char    _bidno1              ;    // [string,    6] 매수증권사명1                   StartPos 271, Length 6
   char    dvol1               [   8];    char    _dvol1               ;    // [long  ,    8] 총매도수량1                     StartPos 278, Length 8
   char    svol1               [   8];    char    _svol1               ;    // [long  ,    8] 총매수수량1                     StartPos 287, Length 8
   char    dcha1               [   8];    char    _dcha1               ;    // [long  ,    8] 매도증감1                       StartPos 296, Length 8
   char    scha1               [   8];    char    _scha1               ;    // [long  ,    8] 매수증감1                       StartPos 305, Length 8
   char    ddiff1              [   6];    char    _ddiff1              ;    // [float ,  6.2] 매도비율1                       StartPos 314, Length 6
   char    sdiff1              [   6];    char    _sdiff1              ;    // [float ,  6.2] 매수비율1                       StartPos 321, Length 6
   char    offernocd2          [   3];    char    _offernocd2          ;    // [string,    3] 매도증권사코드2                 StartPos 328, Length 3
   char    bidnocd2            [   3];    char    _bidnocd2            ;    // [string,    3] 매수증권사코드2                 StartPos 332, Length 3
   char    offerno2            [   6];    char    _offerno2            ;    // [string,    6] 매도증권사명2                   StartPos 336, Length 6
   char    bidno2              [   6];    char    _bidno2              ;    // [string,    6] 매수증권사명2                   StartPos 343, Length 6
   char    dvol2               [   8];    char    _dvol2               ;    // [long  ,    8] 총매도수량2                     StartPos 350, Length 8
   char    svol2               [   8];    char    _svol2               ;    // [long  ,    8] 총매수수량2                     StartPos 359, Length 8
   char    dcha2               [   8];    char    _dcha2               ;    // [long  ,    8] 매도증감2                       StartPos 368, Length 8
   char    scha2               [   8];    char    _scha2               ;    // [long  ,    8] 매수증감2                       StartPos 377, Length 8
   char    ddiff2              [   6];    char    _ddiff2              ;    // [float ,  6.2] 매도비율2                       StartPos 386, Length 6
   char    sdiff2              [   6];    char    _sdiff2              ;    // [float ,  6.2] 매수비율2                       StartPos 393, Length 6
   char    offernocd3          [   3];    char    _offernocd3          ;    // [string,    3] 매도증권사코드3                 StartPos 400, Length 3
   char    bidnocd3            [   3];    char    _bidnocd3            ;    // [string,    3] 매수증권사코드3                 StartPos 404, Length 3
   char    offerno3            [   6];    char    _offerno3            ;    // [string,    6] 매도증권사명3                   StartPos 408, Length 6
   char    bidno3              [   6];    char    _bidno3              ;    // [string,    6] 매수증권사명3                   StartPos 415, Length 6
   char    dvol3               [   8];    char    _dvol3               ;    // [long  ,    8] 총매도수량3                     StartPos 422, Length 8
   char    svol3               [   8];    char    _svol3               ;    // [long  ,    8] 총매수수량3                     StartPos 431, Length 8
   char    dcha3               [   8];    char    _dcha3               ;    // [long  ,    8] 매도증감3                       StartPos 440, Length 8
   char    scha3               [   8];    char    _scha3               ;    // [long  ,    8] 매수증감3                       StartPos 449, Length 8
   char    ddiff3              [   6];    char    _ddiff3              ;    // [float ,  6.2] 매도비율3                       StartPos 458, Length 6
   char    sdiff3              [   6];    char    _sdiff3              ;    // [float ,  6.2] 매수비율3                       StartPos 465, Length 6
   char    offernocd4          [   3];    char    _offernocd4          ;    // [string,    3] 매도증권사코드4                 StartPos 472, Length 3
   char    bidnocd4            [   3];    char    _bidnocd4            ;    // [string,    3] 매수증권사코드4                 StartPos 476, Length 3
   char    offerno4            [   6];    char    _offerno4            ;    // [string,    6] 매도증권사명4                   StartPos 480, Length 6
   char    bidno4              [   6];    char    _bidno4              ;    // [string,    6] 매수증권사명4                   StartPos 487, Length 6
   char    dvol4               [   8];    char    _dvol4               ;    // [long  ,    8] 총매도수량4                     StartPos 494, Length 8
   char    svol4               [   8];    char    _svol4               ;    // [long  ,    8] 총매수수량4                     StartPos 503, Length 8
   char    dcha4               [   8];    char    _dcha4               ;    // [long  ,    8] 매도증감4                       StartPos 512, Length 8
   char    scha4               [   8];    char    _scha4               ;    // [long  ,    8] 매수증감4                       StartPos 521, Length 8
   char    ddiff4              [   6];    char    _ddiff4              ;    // [float ,  6.2] 매도비율4                       StartPos 530, Length 6
   char    sdiff4              [   6];    char    _sdiff4              ;    // [float ,  6.2] 매수비율4                       StartPos 537, Length 6
   char    offernocd5          [   3];    char    _offernocd5          ;    // [string,    3] 매도증권사코드5                 StartPos 544, Length 3
   char    bidnocd5            [   3];    char    _bidnocd5            ;    // [string,    3] 매수증권사코드5                 StartPos 548, Length 3
   char    offerno5            [   6];    char    _offerno5            ;    // [string,    6] 매도증권사명5                   StartPos 552, Length 6
   char    bidno5              [   6];    char    _bidno5              ;    // [string,    6] 매수증권사명5                   StartPos 559, Length 6
   char    dvol5               [   8];    char    _dvol5               ;    // [long  ,    8] 총매도수량5                     StartPos 566, Length 8
   char    svol5               [   8];    char    _svol5               ;    // [long  ,    8] 총매수수량5                     StartPos 575, Length 8
   char    dcha5               [   8];    char    _dcha5               ;    // [long  ,    8] 매도증감5                       StartPos 584, Length 8
   char    scha5               [   8];    char    _scha5               ;    // [long  ,    8] 매수증감5                       StartPos 593, Length 8
   char    ddiff5              [   6];    char    _ddiff5              ;    // [float ,  6.2] 매도비율5                       StartPos 602, Length 6
   char    sdiff5              [   6];    char    _sdiff5              ;    // [float ,  6.2] 매수비율5                       StartPos 609, Length 6
   char    fwdvl               [  12];    char    _fwdvl               ;    // [long  ,   12] 외국계매도합계수량              StartPos 616, Length 12
   char    ftradmdcha          [  12];    char    _ftradmdcha          ;    // [long  ,   12] 외국계매도직전대비              StartPos 629, Length 12
   char    ftradmddiff         [   6];    char    _ftradmddiff         ;    // [float ,  6.2] 외국계매도비율                  StartPos 642, Length 6
   char    fwsvl               [  12];    char    _fwsvl               ;    // [long  ,   12] 외국계매수합계수량              StartPos 649, Length 12
   char    ftradmscha          [  12];    char    _ftradmscha          ;    // [long  ,   12] 외국계매수직전대비              StartPos 662, Length 12
   char    ftradmsdiff         [   6];    char    _ftradmsdiff         ;    // [float ,  6.2] 외국계매수비율                  StartPos 675, Length 6
   char    vol                 [   6];    char    _vol                 ;    // [float ,  6.2] 회전율                          StartPos 682, Length 6
   char    shcode              [   6];    char    _shcode              ;    // [string,    6] 단축코드                        StartPos 689, Length 6
   char    value               [  12];    char    _value               ;    // [long  ,   12] 누적거래대금                    StartPos 696, Length 12
   char    jvolume             [  12];    char    _jvolume             ;    // [long  ,   12] 전일동시간거래량                StartPos 709, Length 12
   char    highyear            [   8];    char    _highyear            ;    // [long  ,    8] 연중최고가                      StartPos 722, Length 8
   char    highyeardate        [   8];    char    _highyeardate        ;    // [string,    8] 연중최고일자                    StartPos 731, Length 8
   char    lowyear             [   8];    char    _lowyear             ;    // [long  ,    8] 연중최저가                      StartPos 740, Length 8
   char    lowyeardate         [   8];    char    _lowyeardate         ;    // [string,    8] 연중최저일자                    StartPos 749, Length 8
   char    target              [   8];    char    _target              ;    // [long  ,    8] 목표가                          StartPos 758, Length 8
   char    capital             [  12];    char    _capital             ;    // [long  ,   12] 자본금                          StartPos 767, Length 12
   char    abscnt              [  12];    char    _abscnt              ;    // [long  ,   12] 유동주식수                      StartPos 780, Length 12
   char    parprice            [   8];    char    _parprice            ;    // [long  ,    8] 액면가                          StartPos 793, Length 8
   char    gsmm                [   2];    char    _gsmm                ;    // [string,    2] 결산월                          StartPos 802, Length 2
   char    subprice            [   8];    char    _subprice            ;    // [long  ,    8] 대용가                          StartPos 805, Length 8
   char    total               [  12];    char    _total               ;    // [long  ,   12] 시가총액                        StartPos 814, Length 12
   char    listdate            [   8];    char    _listdate            ;    // [string,    8] 상장일                          StartPos 827, Length 8
   char    name                [  10];    char    _name                ;    // [string,   10] 전분기명                        StartPos 836, Length 10
   char    bfsales             [  12];    char    _bfsales             ;    // [long  ,   12] 전분기매출액                    StartPos 847, Length 12
   char    bfoperatingincome   [  12];    char    _bfoperatingincome   ;    // [long  ,   12] 전분기영업이익                  StartPos 860, Length 12
   char    bfordinaryincome    [  12];    char    _bfordinaryincome    ;    // [long  ,   12] 전분기경상이익                  StartPos 873, Length 12
   char    bfnetincome         [  12];    char    _bfnetincome         ;    // [long  ,   12] 전분기순이익                    StartPos 886, Length 12
   char    bfeps               [  13];    char    _bfeps               ;    // [float , 13.2] 전분기EPS                       StartPos 899, Length 13
   char    name2               [  10];    char    _name2               ;    // [string,   10] 전전분기명                      StartPos 913, Length 10
   char    bfsales2            [  12];    char    _bfsales2            ;    // [long  ,   12] 전전분기매출액                  StartPos 924, Length 12
   char    bfoperatingincome2  [  12];    char    _bfoperatingincome2  ;    // [long  ,   12] 전전분기영업이익                StartPos 937, Length 12
   char    bfordinaryincome2   [  12];    char    _bfordinaryincome2   ;    // [long  ,   12] 전전분기경상이익                StartPos 950, Length 12
   char    bfnetincome2        [  12];    char    _bfnetincome2        ;    // [long  ,   12] 전전분기순이익                  StartPos 963, Length 12
   char    bfeps2              [  13];    char    _bfeps2              ;    // [float , 13.2] 전전분기EPS                     StartPos 976, Length 13
   char    salert              [   7];    char    _salert              ;    // [float ,  7.2] 전년대비매출액                  StartPos 990, Length 7
   char    opert               [   7];    char    _opert               ;    // [float ,  7.2] 전년대비영업이익                StartPos 998, Length 7
   char    ordrt               [   7];    char    _ordrt               ;    // [float ,  7.2] 전년대비경상이익                StartPos 1006, Length 7
   char    netrt               [   7];    char    _netrt               ;    // [float ,  7.2] 전년대비순이익                  StartPos 1014, Length 7
   char    epsrt               [   7];    char    _epsrt               ;    // [float ,  7.2] 전년대비EPS                     StartPos 1022, Length 7
   char    info1               [  10];    char    _info1               ;    // [string,   10] 락구분                          StartPos 1030, Length 10
   char    info2               [  10];    char    _info2               ;    // [string,   10] 관리/급등구분                   StartPos 1041, Length 10
   char    info3               [  10];    char    _info3               ;    // [string,   10] 정지/연장구분                   StartPos 1052, Length 10
   char    info4               [  12];    char    _info4               ;    // [string,   12] 투자/불성실구분                 StartPos 1063, Length 12
   char    janginfo            [  10];    char    _janginfo            ;    // [string,   10] 장구분                          StartPos 1076, Length 10
   char    t_per               [   6];    char    _t_per               ;    // [float ,  6.2] T.PER                           StartPos 1087, Length 6
   char    tonghwa             [   3];    char    _tonghwa             ;    // [string,    3] 통화ISO코드                     StartPos 1094, Length 3
   char    dval1               [  18];    char    _dval1               ;    // [long  ,   18] 총매도대금1                     StartPos 1098, Length 18
   char    sval1               [  18];    char    _sval1               ;    // [long  ,   18] 총매수대금1                     StartPos 1117, Length 18
   char    dval2               [  18];    char    _dval2               ;    // [long  ,   18] 총매도대금2                     StartPos 1136, Length 18
   char    sval2               [  18];    char    _sval2               ;    // [long  ,   18] 총매수대금2                     StartPos 1155, Length 18
   char    dval3               [  18];    char    _dval3               ;    // [long  ,   18] 총매도대금3                     StartPos 1174, Length 18
   char    sval3               [  18];    char    _sval3               ;    // [long  ,   18] 총매수대금3                     StartPos 1193, Length 18
   char    dval4               [  18];    char    _dval4               ;    // [long  ,   18] 총매도대금4                     StartPos 1212, Length 18
   char    sval4               [  18];    char    _sval4               ;    // [long  ,   18] 총매수대금4                     StartPos 1231, Length 18
   char    dval5               [  18];    char    _dval5               ;    // [long  ,   18] 총매도대금5                     StartPos 1250, Length 18
   char    sval5               [  18];    char    _sval5               ;    // [long  ,   18] 총매수대금5                     StartPos 1269, Length 18
   char    davg1               [   8];    char    _davg1               ;    // [long  ,    8] 총매도평단가1                   StartPos 1288, Length 8
   char    savg1               [   8];    char    _savg1               ;    // [long  ,    8] 총매수평단가1                   StartPos 1297, Length 8
   char    davg2               [   8];    char    _davg2               ;    // [long  ,    8] 총매도평단가2                   StartPos 1306, Length 8
   char    savg2               [   8];    char    _savg2               ;    // [long  ,    8] 총매수평단가2                   StartPos 1315, Length 8
   char    davg3               [   8];    char    _davg3               ;    // [long  ,    8] 총매도평단가3                   StartPos 1324, Length 8
   char    savg3               [   8];    char    _savg3               ;    // [long  ,    8] 총매수평단가3                   StartPos 1333, Length 8
   char    davg4               [   8];    char    _davg4               ;    // [long  ,    8] 총매도평단가4                   StartPos 1342, Length 8
   char    savg4               [   8];    char    _savg4               ;    // [long  ,    8] 총매수평단가4                   StartPos 1351, Length 8
   char    davg5               [   8];    char    _davg5               ;    // [long  ,    8] 총매도평단가5                   StartPos 1360, Length 8
   char    savg5               [   8];    char    _savg5               ;    // [long  ,    8] 총매수평단가5                   StartPos 1369, Length 8
   char    ftradmdval          [  18];    char    _ftradmdval          ;    // [long  ,   18] 외국계매도대금                  StartPos 1378, Length 18
   char    ftradmsval          [  18];    char    _ftradmsval          ;    // [long  ,   18] 외국계매수대금                  StartPos 1397, Length 18
   char    ftradmdavg          [   8];    char    _ftradmdavg          ;    // [long  ,    8] 외국계매도평단가                StartPos 1416, Length 8
   char    ftradmsavg          [   8];    char    _ftradmsavg          ;    // [long  ,    8] 외국계매수평단가                StartPos 1425, Length 8
   char    info5               [   8];    char    _info5               ;    // [string,    8] 투자주의환기                    StartPos 1434, Length 8
   char    spac_gubun          [   1];    char    _spac_gubun          ;    // [string,    1] 기업인수목적회사여부            StartPos 1443, Length 1
   char    issueprice          [   8];    char    _issueprice          ;    // [long  ,    8] 발행가격                        StartPos 1445, Length 8
   char    alloc_gubun         [   1];    char    _alloc_gubun         ;    // [string,    1] 배분적용구분코드(1:배분발생2:배 StartPos 1454, Length 1
   char    alloc_text          [   8];    char    _alloc_text          ;    // [string,    8] 배분적용구분                    StartPos 1456, Length 8
   char    shterm_text         [  10];    char    _shterm_text         ;    // [string,   10] 단기과열/VI발동                 StartPos 1465, Length 10
   char    svi_uplmtprice      [   8];    char    _svi_uplmtprice      ;    // [long  ,    8] 정적VI상한가                    StartPos 1476, Length 8
   char    svi_dnlmtprice      [   8];    char    _svi_dnlmtprice      ;    // [long  ,    8] 정적VI하한가                    StartPos 1485, Length 8
   char    low_lqdt_gu         [   1];    char    _low_lqdt_gu         ;    // [string,    1] 저유동성종목여부                StartPos 1494, Length 1
   char    abnormal_rise_gu    [   1];    char    _abnormal_rise_gu    ;    // [string,    1] 이상급등종목여부                StartPos 1496, Length 1
   char    lend_text           [   8];    char    _lend_text           ;    // [string,    8] 대차불가표시                    StartPos 1498, Length 8
   char    ty_text             [   8];    char    _ty_text             ;    // [string,    8] ETF/ETN투자유의                 StartPos 1507, Length 8
} T1102OutBlock;

//------------------------------------------------------------------------------
// 현물 시간대별 체결 조회 (t1301)
//------------------------------------------------------------------------------
typedef struct {
    char    shcode              [   6];    char    _shcode              ;    // [string,    6] 단축코드                        StartPos 0, Length 6
    char    cvolume             [  12];    char    _cvolume             ;    // [long  ,   12] 특이거래량                      StartPos 7, Length 12
    char    starttime           [   4];    char    _starttime           ;    // [string,    4] 시작시간                        StartPos 20, Length 4
    char    endtime             [   4];    char    _endtime             ;    // [string,    4] 종료시간                        StartPos 25, Length 4
    char    cts_time            [  10];    char    _cts_time            ;    // [string,   10] 시간CTS                         StartPos 30, Length 10
} T1301InBlock;

typedef struct {
    char    cts_time            [  10];    char    _cts_time            ;    // [string,   10] 시간CTS                         StartPos 0, Length 10
} T1301OutBlock;

typedef struct {
    char    chetime             [  10];    char    _chetime             ;    // [string,   10] 시간                            StartPos 0, Length 10
    char    price               [   8];    char    _price               ;    // [long  ,    8] 현재가                          StartPos 11, Length 8
    char    sign                [   1];    char    _sign                ;    // [string,    1] 전일대비구분                    StartPos 20, Length 1
    char    change              [   8];    char    _change              ;    // [long  ,    8] 전일대비                        StartPos 22, Length 8
    char    diff                [   6];    char    _diff                ;    // [float ,  6.2] 등락율                          StartPos 31, Length 6
    char    cvolume             [  12];    char    _cvolume             ;    // [long  ,   12] 체결수량                        StartPos 38, Length 12
    char    chdegree            [   8];    char    _chdegree            ;    // [float ,  8.2] 체결강도                        StartPos 51, Length 8
    char    volume              [  12];    char    _volume              ;    // [long  ,   12] 거래량                          StartPos 60, Length 12
    char    mdvolume            [  12];    char    _mdvolume            ;    // [long  ,   12] 매도체결수량                    StartPos 73, Length 12
    char    mdchecnt            [   8];    char    _mdchecnt            ;    // [long  ,    8] 매도체결건수                    StartPos 86, Length 8
    char    msvolume            [  12];    char    _msvolume            ;    // [long  ,   12] 매수체결수량                    StartPos 95, Length 12
    char    mschecnt            [   8];    char    _mschecnt            ;    // [long  ,    8] 매수체결건수                    StartPos 108, Length 8
    char    revolume            [  12];    char    _revolume            ;    // [long  ,   12] 순체결량                        StartPos 117, Length 12
    char    rechecnt            [   8];    char    _rechecnt            ;    // [long  ,    8] 순체결건수                      StartPos 130, Length 8
} T1301OutBlock1;

//------------------------------------------------------------------------------
// 기간별 주가 (t1305)
//------------------------------------------------------------------------------
typedef struct {
    char    shcode              [   6];    char    _shcode              ;    // [string,    6] 단축코드                        StartPos 0, Length 6
    char    dwmcode             [   1];    char    _dwmcode             ;    // [long  ,    1] 일주월구분                      StartPos 7, Length 1
    char    date                [   8];    char    _date                ;    // [string,    8] 날짜                            StartPos 9, Length 8
    char    idx                 [   4];    char    _idx                 ;    // [long  ,    4] IDX                             StartPos 18, Length 4
    char    cnt                 [   4];    char    _cnt                 ;    // [long  ,    4] 건수                            StartPos 23, Length 4
} T1305InBlock;

typedef struct {
    char    cnt                 [   4];    char    _cnt                 ;    // [long  ,    4] CNT                             StartPos 0, Length 4
    char    date                [   8];    char    _date                ;    // [string,    8] 날짜                            StartPos 5, Length 8
    char    idx                 [   4];    char    _idx                 ;    // [long  ,    4] IDX                             StartPos 14, Length 4
} T1305OutBlock;

typedef struct {
    char    date                [   8];    char    _date                ;    // [string,    8] 날짜                            StartPos 0, Length 8
    char    open                [   8];    char    _open                ;    // [long  ,    8] 시가                            StartPos 9, Length 8
    char    high                [   8];    char    _high                ;    // [long  ,    8] 고가                            StartPos 18, Length 8
    char    low                 [   8];    char    _low                 ;    // [long  ,    8] 저가                            StartPos 27, Length 8
    char    close               [   8];    char    _close               ;    // [long  ,    8] 종가                            StartPos 36, Length 8
    char    sign                [   1];    char    _sign                ;    // [string,    1] 전일대비구분                    StartPos 45, Length 1
    char    change              [   8];    char    _change              ;    // [long  ,    8] 전일대비                        StartPos 47, Length 8
    char    diff                [   6];    char    _diff                ;    // [float ,  6.2] 등락율                          StartPos 56, Length 6
    char    volume              [  12];    char    _volume              ;    // [long  ,   12] 누적거래량                      StartPos 63, Length 12
    char    diff_vol            [  10];    char    _diff_vol            ;    // [float , 10.2] 거래증가율                      StartPos 76, Length 10
    char    chdegree            [   6];    char    _chdegree            ;    // [float ,  6.2] 체결강도                        StartPos 87, Length 6
    char    sojinrate           [   6];    char    _sojinrate           ;    // [float ,  6.2] 소진율                          StartPos 94, Length 6
    char    changerate          [   6];    char    _changerate          ;    // [float ,  6.2] 회전율                          StartPos 101, Length 6
    char    fpvolume            [  12];    char    _fpvolume            ;    // [long  ,   12] 외인순매수                      StartPos 108, Length 12
    char    covolume            [  12];    char    _covolume            ;    // [long  ,   12] 기관순매수                      StartPos 121, Length 12
    char    shcode              [   6];    char    _shcode              ;    // [string,    6] 종목코드                        StartPos 134, Length 6
    char    value               [  12];    char    _value               ;    // [long  ,   12] 누적거래대금(단위:백만)         StartPos 141, Length 12
    char    ppvolume            [  12];    char    _ppvolume            ;    // [long  ,   12] 개인순매수                      StartPos 154, Length 12
    char    o_sign              [   1];    char    _o_sign              ;    // [string,    1] 시가대비구분                    StartPos 167, Length 1
    char    o_change            [   8];    char    _o_change            ;    // [long  ,    8] 시가대비                        StartPos 169, Length 8
    char    o_diff              [   6];    char    _o_diff              ;    // [float ,  6.2] 시가기준등락율                  StartPos 178, Length 6
    char    h_sign              [   1];    char    _h_sign              ;    // [string,    1] 고가대비구분                    StartPos 185, Length 1
    char    h_change            [   8];    char    _h_change            ;    // [long  ,    8] 고가대비                        StartPos 187, Length 8
    char    h_diff              [   6];    char    _h_diff              ;    // [float ,  6.2] 고가기준등락율                  StartPos 196, Length 6
    char    l_sign              [   1];    char    _l_sign              ;    // [string,    1] 저가대비구분                    StartPos 203, Length 1
    char    l_change            [   8];    char    _l_change            ;    // [long  ,    8] 저가대비                        StartPos 205, Length 8
    char    l_diff              [   6];    char    _l_diff              ;    // [float ,  6.2] 저가기준등락율                  StartPos 214, Length 6
    char    marketcap           [  12];    char    _marketcap           ;    // [long  ,   12] 시가총액(단위:백만)             StartPos 221, Length 12
} T1305OutBlock1;

//------------------------------------------------------------------------------
// 현물 당일전일분틱조회 (t1310)
//------------------------------------------------------------------------------
typedef struct {
    char    daygb               [   1];    char    _daygb               ;    // [string,    1] 당일전일구분                    StartPos 0, Length 1
    char    timegb              [   1];    char    _timegb              ;    // [string,    1] 분틱구분                        StartPos 2, Length 1
    char    shcode              [   6];    char    _shcode              ;    // [string,    6] 단축코드                        StartPos 4, Length 6
    char    endtime             [   4];    char    _endtime             ;    // [string,    4] 종료시간                        StartPos 11, Length 4
    char    cts_time            [  10];    char    _cts_time            ;    // [string,   10] 시간CTS                         StartPos 16, Length 10
} T1310InBlock;

typedef struct {
    char    cts_time            [  10];    char    _cts_time            ;    // [string,   10] 시간CTS                         StartPos 0, Length 10
} T1310OutBlock;

typedef struct {
    char    chetime             [  10];    char    _chetime             ;    // [string,   10] 시간                            StartPos 0, Length 10
    char    price               [   8];    char    _price               ;    // [long  ,    8] 현재가                          StartPos 11, Length 8
    char    sign                [   1];    char    _sign                ;    // [string,    1] 전일대비구분                    StartPos 20, Length 1
    char    change              [   8];    char    _change              ;    // [long  ,    8] 전일대비                        StartPos 22, Length 8
    char    diff                [   6];    char    _diff                ;    // [float ,  6.2] 등락율                          StartPos 31, Length 6
    char    cvolume             [  12];    char    _cvolume             ;    // [long  ,   12] 체결수량                        StartPos 38, Length 12
    char    chdegree            [   8];    char    _chdegree            ;    // [float ,  8.2] 체결강도                        StartPos 51, Length 8
    char    volume              [  12];    char    _volume              ;    // [long  ,   12] 거래량                          StartPos 60, Length 12
    char    mdvolume            [  12];    char    _mdvolume            ;    // [long  ,   12] 매도체결수량                    StartPos 73, Length 12
    char    mdchecnt            [   8];    char    _mdchecnt            ;    // [long  ,    8] 매도체결건수                    StartPos 86, Length 8
    char    msvolume            [  12];    char    _msvolume            ;    // [long  ,   12] 매수체결수량                    StartPos 95, Length 12
    char    mschecnt            [   8];    char    _mschecnt            ;    // [long  ,    8] 매수체결건수                    StartPos 108, Length 8
    char    revolume            [  12];    char    _revolume            ;    // [long  ,   12] 순체결량                        StartPos 117, Length 12
    char    rechecnt            [   8];    char    _rechecnt            ;    // [long  ,    8] 순체결건수                      StartPos 130, Length 8
} T1310OutBlock1;

//------------------------------------------------------------------------------
// 관리/불성실/투자유의 조회 (t1404)
//------------------------------------------------------------------------------

// 기본입력
typedef struct {
    char    gubun               [   1];    char    _gubun               ;    // [string,    1] 구분                            StartPos 0, Length 1
    char    jongchk             [   1];    char    _jongchk             ;    // [string,    1] 종목체크                        StartPos 2, Length 1
    char    cts_shcode          [   6];    char    _cts_shcode          ;    // [string,    6] 종목코드_CTS                    StartPos 4, Length 6
} T1404InBlock;

// 출력
typedef struct {
    char    cts_shcode          [   6];    char    _cts_shcode          ;    // [string,    6] 종목코드_CTS                    StartPos 0, Length 6
} T1404OutBlock;

// 출력1                          , occurs
typedef struct {
    char    hname               [  20];    char    _hname               ;    // [string,   20] 한글명                          StartPos 0, Length 20
    char    price               [   8];    char    _price               ;    // [long  ,    8] 현재가                          StartPos 21, Length 8
    char    sign                [   1];    char    _sign                ;    // [string,    1] 전일대비구분                    StartPos 30, Length 1
    char    change              [   8];    char    _change              ;    // [long  ,    8] 전일대비                        StartPos 32, Length 8
    char    diff                [   6];    char    _diff                ;    // [float ,  6.2] 등락율                          StartPos 41, Length 6
    char    volume              [  12];    char    _volume              ;    // [long  ,   12] 누적거래량                      StartPos 48, Length 12
    char    date                [   8];    char    _date                ;    // [string,    8] 지정일                          StartPos 61, Length 8
    char    tprice              [   8];    char    _tprice              ;    // [long  ,    8] 지정일주가                      StartPos 70, Length 8
    char    tchange             [   8];    char    _tchange             ;    // [long  ,    8] 지정일대비                      StartPos 79, Length 8
    char    tdiff               [   6];    char    _tdiff               ;    // [float ,  6.2] 대비율                          StartPos 88, Length 6
    char    reason              [   4];    char    _reason              ;    // [string,    4] 사유                            StartPos 95, Length 4
    char    shcode              [   6];    char    _shcode              ;    // [string,    6] 종목코드                        StartPos 100, Length 6
    char    edate               [   8];    char    _edate               ;    // [string,    8] 해제일                          StartPos 107, Length 8
} T1404OutBlock1;

//------------------------------------------------------------------------------
// 투자경고/매매정지/정리매매조회 (t1405)
//------------------------------------------------------------------------------

// 기본입력
typedef struct {
    char    gubun               [   1];    char    _gubun               ;    // [string,    1] 구분                            StartPos 0, Length 1
    char    jongchk             [   1];    char    _jongchk             ;    // [string,    1] 종목체크                        StartPos 2, Length 1
    char    cts_shcode          [   6];    char    _cts_shcode          ;    // [string,    6] 종목코드_CTS                    StartPos 4, Length 6
} T1405InBlock;

// 출력
typedef struct {
    char    cts_shcode          [   6];    char    _cts_shcode          ;    // [string,    6] 종목코드_CTS                    StartPos 0, Length 6
} T1405OutBlock;

// 출력1                          , occurs
typedef struct {
    char    hname               [  20];    char    _hname               ;    // [string,   20] 한글명                          StartPos 0, Length 20
    char    price               [   8];    char    _price               ;    // [long  ,    8] 현재가                          StartPos 21, Length 8
    char    sign                [   1];    char    _sign                ;    // [string,    1] 전일대비구분                    StartPos 30, Length 1
    char    change              [   8];    char    _change              ;    // [long  ,    8] 전일대비                        StartPos 32, Length 8
    char    diff                [   6];    char    _diff                ;    // [float ,  6.2] 등락율                          StartPos 41, Length 6
    char    volume              [  12];    char    _volume              ;    // [long  ,   12] 누적거래량                      StartPos 48, Length 12
    char    date                [   8];    char    _date                ;    // [string,    8] 지정일                          StartPos 61, Length 8
    char    edate               [   8];    char    _edate               ;    // [string,    8] 해제일                          StartPos 70, Length 8
    char    shcode              [   6];    char    _shcode              ;    // [string,    6] 종목코드                        StartPos 79, Length 6
} T1405OutBlock1;

//------------------------------------------------------------------------------
// ETF 현재가(시세) 조회 (t1901)
//------------------------------------------------------------------------------
typedef struct {
    char    shcode              [   6];    char    _shcode              ;    // [string,    6] 단축코드                        StartPos 0, Length 6
} T1901InBlock;

typedef struct {
    char    hname               [  20];    char    _hname               ;    // [string,   20] 한글명                          StartPos 0, Length 20
    char    price               [   8];    char    _price               ;    // [long  ,    8] 현재가                          StartPos 21, Length 8
    char    sign                [   1];    char    _sign                ;    // [string,    1] 전일대비구분                    StartPos 30, Length 1
    char    change              [   8];    char    _change              ;    // [long  ,    8] 전일대비                        StartPos 32, Length 8
    char    diff                [   6];    char    _diff                ;    // [float ,  6.2] 등락율                          StartPos 41, Length 6
    char    volume              [  12];    char    _volume              ;    // [float ,   12] 누적거래량                      StartPos 48, Length 12
    char    recprice            [   8];    char    _recprice            ;    // [long  ,    8] 기준가                          StartPos 61, Length 8
    char    avg                 [   8];    char    _avg                 ;    // [long  ,    8] 가중평균                        StartPos 70, Length 8
    char    uplmtprice          [   8];    char    _uplmtprice          ;    // [long  ,    8] 상한가                          StartPos 79, Length 8
    char    dnlmtprice          [   8];    char    _dnlmtprice          ;    // [long  ,    8] 하한가                          StartPos 88, Length 8
    char    jnilvolume          [  12];    char    _jnilvolume          ;    // [float ,   12] 전일거래량                      StartPos 97, Length 12
    char    volumediff          [  12];    char    _volumediff          ;    // [long  ,   12] 거래량차                        StartPos 110, Length 12
    char    open                [   8];    char    _open                ;    // [long  ,    8] 시가                            StartPos 123, Length 8
    char    opentime            [   6];    char    _opentime            ;    // [string,    6] 시가시간                        StartPos 132, Length 6
    char    high                [   8];    char    _high                ;    // [long  ,    8] 고가                            StartPos 139, Length 8
    char    hightime            [   6];    char    _hightime            ;    // [string,    6] 고가시간                        StartPos 148, Length 6
    char    low                 [   8];    char    _low                 ;    // [long  ,    8] 저가                            StartPos 155, Length 8
    char    lowtime             [   6];    char    _lowtime             ;    // [string,    6] 저가시간                        StartPos 164, Length 6
    char    high52w             [   8];    char    _high52w             ;    // [long  ,    8] 52최고가                        StartPos 171, Length 8
    char    high52wdate         [   8];    char    _high52wdate         ;    // [string,    8] 52최고가일                      StartPos 180, Length 8
    char    low52w              [   8];    char    _low52w              ;    // [long  ,    8] 52최저가                        StartPos 189, Length 8
    char    low52wdate          [   8];    char    _low52wdate          ;    // [string,    8] 52최저가일                      StartPos 198, Length 8
    char    exhratio            [   6];    char    _exhratio            ;    // [float ,  6.2] 소진율                          StartPos 207, Length 6
    char    flmtvol             [  12];    char    _flmtvol             ;    // [float ,   12] 외국인보유수량                  StartPos 214, Length 12
    char    per                 [   6];    char    _per                 ;    // [float ,  6.2] PER                             StartPos 227, Length 6
    char    listing             [  12];    char    _listing             ;    // [long  ,   12] 상장주식수(천)                  StartPos 234, Length 12
    char    jkrate              [   8];    char    _jkrate              ;    // [long  ,    8] 증거금율                        StartPos 247, Length 8
    char    vol                 [   6];    char    _vol                 ;    // [float ,  6.2] 회전율                          StartPos 256, Length 6
    char    shcode              [   6];    char    _shcode              ;    // [string,    6] 단축코드                        StartPos 263, Length 6
    char    value               [  12];    char    _value               ;    // [long  ,   12] 누적거래대금                    StartPos 270, Length 12
    char    highyear            [   8];    char    _highyear            ;    // [long  ,    8] 연중최고가                      StartPos 283, Length 8
    char    highyeardate        [   8];    char    _highyeardate        ;    // [string,    8] 연중최고일자                    StartPos 292, Length 8
    char    lowyear             [   8];    char    _lowyear             ;    // [long  ,    8] 연중최저가                      StartPos 301, Length 8
    char    lowyeardate         [   8];    char    _lowyeardate         ;    // [string,    8] 연중최저일자                    StartPos 310, Length 8
    char    upname              [  20];    char    _upname              ;    // [string,   20] 업종명                          StartPos 319, Length 20
    char    upcode              [   3];    char    _upcode              ;    // [string,    3] 업종코드                        StartPos 340, Length 3
    char    upprice             [   7];    char    _upprice             ;    // [float ,  7.2] 업종현재가                      StartPos 344, Length 7
    char    upsign              [   1];    char    _upsign              ;    // [string,    1] 업종전일비구분                  StartPos 352, Length 1
    char    upchange            [   6];    char    _upchange            ;    // [float ,  6.2] 업종전일대비                    StartPos 354, Length 6
    char    updiff              [   6];    char    _updiff              ;    // [float ,  6.2] 업종등락율                      StartPos 361, Length 6
    char    futname             [  20];    char    _futname             ;    // [string,   20] 선물최근월물명                  StartPos 368, Length 20
    char    futcode             [   8];    char    _futcode             ;    // [string,    8] 선물최근월물코드                StartPos 389, Length 8
    char    futprice            [   6];    char    _futprice            ;    // [float ,  6.2] 선물현재가                      StartPos 398, Length 6
    char    futsign             [   1];    char    _futsign             ;    // [string,    1] 선물전일비구분                  StartPos 405, Length 1
    char    futchange           [   6];    char    _futchange           ;    // [float ,  6.2] 선물전일대비                    StartPos 407, Length 6
    char    futdiff             [   6];    char    _futdiff             ;    // [float ,  6.2] 선물등락율                      StartPos 414, Length 6
    char    nav                 [   8];    char    _nav                 ;    // [float ,  8.2] NAV                             StartPos 421, Length 8
    char    navsign             [   1];    char    _navsign             ;    // [string,    1] NAV전일대비구분                 StartPos 430, Length 1
    char    navchange           [   8];    char    _navchange           ;    // [float ,  8.2] NAV전일대비                     StartPos 432, Length 8
    char    navdiff             [   6];    char    _navdiff             ;    // [float ,  6.2] NAV등락율                       StartPos 441, Length 6
    char    cocrate             [   6];    char    _cocrate             ;    // [float ,  6.2] 추적오차율                      StartPos 448, Length 6
    char    kasis               [   6];    char    _kasis               ;    // [float ,  6.2] 괴리율                          StartPos 455, Length 6
    char    subprice            [  10];    char    _subprice            ;    // [long  ,   10] 대용가                          StartPos 462, Length 10
    char    offerno1            [   6];    char    _offerno1            ;    // [string,    6] 매도증권사코드1                 StartPos 473, Length 6
    char    bidno1              [   6];    char    _bidno1              ;    // [string,    6] 매수증권사코드1                 StartPos 480, Length 6
    char    dvol1               [   8];    char    _dvol1               ;    // [long  ,    8] 총매도수량1                     StartPos 487, Length 8
    char    svol1               [   8];    char    _svol1               ;    // [long  ,    8] 총매수수량1                     StartPos 496, Length 8
    char    dcha1               [   8];    char    _dcha1               ;    // [long  ,    8] 매도증감1                       StartPos 505, Length 8
    char    scha1               [   8];    char    _scha1               ;    // [long  ,    8] 매수증감1                       StartPos 514, Length 8
    char    ddiff1              [   6];    char    _ddiff1              ;    // [float ,  6.2] 매도비율1                       StartPos 523, Length 6
    char    sdiff1              [   6];    char    _sdiff1              ;    // [float ,  6.2] 매수비율1                       StartPos 530, Length 6
    char    offerno2            [   6];    char    _offerno2            ;    // [string,    6] 매도증권사코드2                 StartPos 537, Length 6
    char    bidno2              [   6];    char    _bidno2              ;    // [string,    6] 매수증권사코드2                 StartPos 544, Length 6
    char    dvol2               [   8];    char    _dvol2               ;    // [long  ,    8] 총매도수량2                     StartPos 551, Length 8
    char    svol2               [   8];    char    _svol2               ;    // [long  ,    8] 총매수수량2                     StartPos 560, Length 8
    char    dcha2               [   8];    char    _dcha2               ;    // [long  ,    8] 매도증감2                       StartPos 569, Length 8
    char    scha2               [   8];    char    _scha2               ;    // [long  ,    8] 매수증감2                       StartPos 578, Length 8
    char    ddiff2              [   6];    char    _ddiff2              ;    // [float ,  6.2] 매도비율2                       StartPos 587, Length 6
    char    sdiff2              [   6];    char    _sdiff2              ;    // [float ,  6.2] 매수비율2                       StartPos 594, Length 6
    char    offerno3            [   6];    char    _offerno3            ;    // [string,    6] 매도증권사코드3                 StartPos 601, Length 6
    char    bidno3              [   6];    char    _bidno3              ;    // [string,    6] 매수증권사코드3                 StartPos 608, Length 6
    char    dvol3               [   8];    char    _dvol3               ;    // [long  ,    8] 총매도수량3                     StartPos 615, Length 8
    char    svol3               [   8];    char    _svol3               ;    // [long  ,    8] 총매수수량3                     StartPos 624, Length 8
    char    dcha3               [   8];    char    _dcha3               ;    // [long  ,    8] 매도증감3                       StartPos 633, Length 8
    char    scha3               [   8];    char    _scha3               ;    // [long  ,    8] 매수증감3                       StartPos 642, Length 8
    char    ddiff3              [   6];    char    _ddiff3              ;    // [float ,  6.2] 매도비율3                       StartPos 651, Length 6
    char    sdiff3              [   6];    char    _sdiff3              ;    // [float ,  6.2] 매수비율3                       StartPos 658, Length 6
    char    offerno4            [   6];    char    _offerno4            ;    // [string,    6] 매도증권사코드4                 StartPos 665, Length 6
    char    bidno4              [   6];    char    _bidno4              ;    // [string,    6] 매수증권사코드4                 StartPos 672, Length 6
    char    dvol4               [   8];    char    _dvol4               ;    // [long  ,    8] 총매도수량4                     StartPos 679, Length 8
    char    svol4               [   8];    char    _svol4               ;    // [long  ,    8] 총매수수량4                     StartPos 688, Length 8
    char    dcha4               [   8];    char    _dcha4               ;    // [long  ,    8] 매도증감4                       StartPos 697, Length 8
    char    scha4               [   8];    char    _scha4               ;    // [long  ,    8] 매수증감4                       StartPos 706, Length 8
    char    ddiff4              [   6];    char    _ddiff4              ;    // [float ,  6.2] 매도비율4                       StartPos 715, Length 6
    char    sdiff4              [   6];    char    _sdiff4              ;    // [float ,  6.2] 매수비율4                       StartPos 722, Length 6
    char    offerno5            [   6];    char    _offerno5            ;    // [string,    6] 매도증권사코드5                 StartPos 729, Length 6
    char    bidno5              [   6];    char    _bidno5              ;    // [string,    6] 매수증권사코드5                 StartPos 736, Length 6
    char    dvol5               [   8];    char    _dvol5               ;    // [long  ,    8] 총매도수량5                     StartPos 743, Length 8
    char    svol5               [   8];    char    _svol5               ;    // [long  ,    8] 총매수수량5                     StartPos 752, Length 8
    char    dcha5               [   8];    char    _dcha5               ;    // [long  ,    8] 매도증감5                       StartPos 761, Length 8
    char    scha5               [   8];    char    _scha5               ;    // [long  ,    8] 매수증감5                       StartPos 770, Length 8
    char    ddiff5              [   6];    char    _ddiff5              ;    // [float ,  6.2] 매도비율5                       StartPos 779, Length 6
    char    sdiff5              [   6];    char    _sdiff5              ;    // [float ,  6.2] 매수비율5                       StartPos 786, Length 6
    char    fwdvl               [  12];    char    _fwdvl               ;    // [long  ,   12] 외국계매도합계수량              StartPos 793, Length 12
    char    ftradmdcha          [  12];    char    _ftradmdcha          ;    // [long  ,   12] 외국계매도직전대비              StartPos 806, Length 12
    char    ftradmddiff         [   6];    char    _ftradmddiff         ;    // [float ,  6.2] 외국계매도비율                  StartPos 819, Length 6
    char    fwsvl               [  12];    char    _fwsvl               ;    // [long  ,   12] 외국계매수합계수량              StartPos 826, Length 12
    char    ftradmscha          [  12];    char    _ftradmscha          ;    // [long  ,   12] 외국계매수직전대비              StartPos 839, Length 12
    char    ftradmsdiff         [   6];    char    _ftradmsdiff         ;    // [float ,  6.2] 외국계매수비율                  StartPos 852, Length 6
    char    upname2             [  20];    char    _upname2             ;    // [string,   20] 참고지수명                      StartPos 859, Length 20
    char    upcode2             [   3];    char    _upcode2             ;    // [string,    3] 참고지수코드                    StartPos 880, Length 3
    char    upprice2            [   7];    char    _upprice2            ;    // [float ,  7.2] 참고지수현재가                  StartPos 884, Length 7
    char    jnilnav             [   8];    char    _jnilnav             ;    // [float ,  8.2] 전일NAV                         StartPos 892, Length 8
    char    jnilnavsign         [   1];    char    _jnilnavsign         ;    // [string,    1] 전일NAV전일대비구분             StartPos 901, Length 1
    char    jnilnavchange       [   8];    char    _jnilnavchange       ;    // [float ,  8.2] 전일NAV전일대비                 StartPos 903, Length 8
    char    jnilnavdiff         [   6];    char    _jnilnavdiff         ;    // [float ,  6.2] 전일NAV등락율                   StartPos 912, Length 6
    char    etftotcap           [  12];    char    _etftotcap           ;    // [long  ,   12] 순자산총액(억원)                StartPos 919, Length 12
    char    spread              [   6];    char    _spread              ;    // [float ,  6.2] 스프레드                        StartPos 932, Length 6
    char    leverage            [   2];    char    _leverage            ;    // [long  ,    2] 레버리지                        StartPos 939, Length 2
    char    taxgubun            [   1];    char    _taxgubun            ;    // [string,    1] 과세구분                        StartPos 942, Length 1
    char    opcom_nmk           [  20];    char    _opcom_nmk           ;    // [string,   20] 운용사                          StartPos 944, Length 20
    char    lp_nm1              [  20];    char    _lp_nm1              ;    // [string,   20] LP1                             StartPos 965, Length 20
    char    lp_nm2              [  20];    char    _lp_nm2              ;    // [string,   20] LP2                             StartPos 986, Length 20
    char    lp_nm3              [  20];    char    _lp_nm3              ;    // [string,   20] LP3                             StartPos 1007, Length 20
    char    lp_nm4              [  20];    char    _lp_nm4              ;    // [string,   20] LP4                             StartPos 1028, Length 20
    char    lp_nm5              [  20];    char    _lp_nm5              ;    // [string,   20] LP5                             StartPos 1049, Length 20
    char    etf_cp              [  10];    char    _etf_cp              ;    // [string,   10] 복제방법                        StartPos 1070, Length 10
    char    etf_kind            [  10];    char    _etf_kind            ;    // [string,   10] 상품유형                        StartPos 1081, Length 10
    char    vi_gubun            [  10];    char    _vi_gubun            ;    // [string,   10] VI발동해제                      StartPos 1092, Length 10
    char    etn_kind_cd         [  20];    char    _etn_kind_cd         ;    // [string,   20] ETN상품분류                     StartPos 1103, Length 20
    char    lastymd             [   8];    char    _lastymd             ;    // [string,    8] ETN만기일                       StartPos 1124, Length 8
    char    payday              [   8];    char    _payday              ;    // [string,    8] ETN지급일                       StartPos 1133, Length 8
    char    lastdate            [   8];    char    _lastdate            ;    // [string,    8] ETN최종거래일                   StartPos 1142, Length 8
    char    issuernmk           [  20];    char    _issuernmk           ;    // [string,   20] ETN발행시장참가자               StartPos 1151, Length 20
    char    last_sdate          [   8];    char    _last_sdate          ;    // [string,    8] ETN만기상환가격결정시작일       StartPos 1172, Length 8
    char    last_edate          [   8];    char    _last_edate          ;    // [string,    8] ETN만기상환가격결정종료일       StartPos 1181, Length 8
    char    lp_holdvol          [  12];    char    _lp_holdvol          ;    // [string,   12] ETNLP보유수량                   StartPos 1190, Length 12
    char    listdate            [   8];    char    _listdate            ;    // [string,    8] 상장일                          StartPos 1203, Length 8
    char    etp_gb              [   1];    char    _etp_gb              ;    // [string,    1] ETP상품구분코드                 StartPos 1212, Length 1
    char    etn_elback_yn       [   1];    char    _etn_elback_yn       ;    // [string,    1] ETN조기상환가능여부             StartPos 1214, Length 1
    char    settletype          [   2];    char    _settletype          ;    // [string,    2] 최종결제                        StartPos 1216, Length 2
    char    idx_asset_class1    [   2];    char    _idx_asset_class1    ;    // [string,    2] 지수자산분류코드(대분류)        StartPos 1219, Length 2
    char    ty_text             [   8];    char    _ty_text             ;    // [string,    8] ETF/ETN투자유의                 StartPos 1222, Length 8
} T1901OutBlock;

//------------------------------------------------------------------------------
// ETF 시간별 추이 (t1902)
//------------------------------------------------------------------------------
typedef struct {
    char    shCode[6];  char _shcode;       //[string,    6] 단축코드   StartPos 0, Length 6
    char    time[6];    char _time;         //[string,    6] 시간   StartPos 7, Length 6
} T1902InBlock;

typedef struct {
    char    time[6];    char _time;         //[string,    6] 시간   StartPos 0, Length 6
    char    hName[20];  char _hname;        //[string,   20] 종목명   StartPos 7, Length 20
    char    upName[20]; char _upname;       //[string,   20] 업종지수명   StartPos 28, Length 20
} T1902OutBlock;

typedef struct {    // occurs
    char    time[8];    char _time;         //[string,    8] 시간   StartPos 0, Length 8
    char    price[8];   char _price;        //[long  ,    8] 현재가   StartPos 9, Length 8
    char    sign[1];    char _sign;         //[string,    1] 전일대비구분   StartPos 18, Length 1
    char    change[8];  char _change;       //[long  ,    8] 전일대비   StartPos 20, Length 8
    char    volume[12]; char _volume;       //[float ,   12] 누적거래량   StartPos 29, Length 12
    char    navDiff[9]; char _navdiff;      //[float ,  9.2] NAV대비   StartPos 42, Length 9
    char    nav[9];    char _nav;          //[float ,  9.2] NAV   StartPos 52, Length 9
    char    navChange[9];   char _navchange;    //[float ,  9.2] 전일대비   StartPos 62, Length 9
    char    crate[9];   char _crate;        //[float ,  9.2] 추적오차   StartPos 72, Length 9
    char    grate[9];   char _grate;        //[float ,  9.2] 괴리   StartPos 82, Length 9
    char    jisu[8];    char _jisu;         //[float ,  8.2] 지수   StartPos 92, Length 8
    char    jiChange[8];    char _jichange; //[float ,  8.2] 전일대비   StartPos 101, Length 8
    char    jiRate[8];  char _jirate;       //[float ,  8.2] 전일대비율   StartPos 110, Length 8
} T1902OutBlock1;

//------------------------------------------------------------------------------
// ETF LP 호가 (t1906)
//------------------------------------------------------------------------------
typedef struct {
    char    shcode              [   6];    char    _shcode              ;    // [string,    6] 단축코드                        StartPos 0, Length 6
} T1906InBlock;

typedef struct {
    char    hname               [  20];    char    _hname               ;    // [string,   20] 한글명                          StartPos 0, Length 20
    char    price               [   8];    char    _price               ;    // [long  ,    8] 현재가                          StartPos 21, Length 8
    char    sign                [   1];    char    _sign                ;    // [string,    1] 전일대비구분                    StartPos 30, Length 1
    char    change              [   8];    char    _change              ;    // [long  ,    8] 전일대비                        StartPos 32, Length 8
    char    diff                [   6];    char    _diff                ;    // [float ,  6.2] 등락율                          StartPos 41, Length 6
    char    volume              [  12];    char    _volume              ;    // [long  ,   12] 누적거래량                      StartPos 48, Length 12
    char    lp_offerrem1        [  12];    char    _lp_offerrem1        ;    // [long  ,   12] LP매도호가수량1                 StartPos 61, Length 12
    char    lp_bidrem1          [  12];    char    _lp_bidrem1          ;    // [long  ,   12] LP매수호가수량1                 StartPos 74, Length 12
    char    lp_offerrem2        [  12];    char    _lp_offerrem2        ;    // [long  ,   12] LP매도호가수량2                 StartPos 87, Length 12
    char    lp_bidrem2          [  12];    char    _lp_bidrem2          ;    // [long  ,   12] LP매수호가수량2                 StartPos 100, Length 12
    char    lp_offerrem3        [  12];    char    _lp_offerrem3        ;    // [long  ,   12] LP매도호가수량3                 StartPos 113, Length 12
    char    lp_bidrem3          [  12];    char    _lp_bidrem3          ;    // [long  ,   12] LP매수호가수량3                 StartPos 126, Length 12
    char    lp_offerrem4        [  12];    char    _lp_offerrem4        ;    // [long  ,   12] LP매도호가수량4                 StartPos 139, Length 12
    char    lp_bidrem4          [  12];    char    _lp_bidrem4          ;    // [long  ,   12] LP매수호가수량4                 StartPos 152, Length 12
    char    lp_offerrem5        [  12];    char    _lp_offerrem5        ;    // [long  ,   12] LP매도호가수량5                 StartPos 165, Length 12
    char    lp_bidrem5          [  12];    char    _lp_bidrem5          ;    // [long  ,   12] LP매수호가수량5                 StartPos 178, Length 12
    char    lp_offerrem6        [  12];    char    _lp_offerrem6        ;    // [long  ,   12] LP매도호가수량6                 StartPos 191, Length 12
    char    lp_bidrem6          [  12];    char    _lp_bidrem6          ;    // [long  ,   12] LP매수호가수량6                 StartPos 204, Length 12
    char    lp_offerrem7        [  12];    char    _lp_offerrem7        ;    // [long  ,   12] LP매도호가수량7                 StartPos 217, Length 12
    char    lp_bidrem7          [  12];    char    _lp_bidrem7          ;    // [long  ,   12] LP매수호가수량7                 StartPos 230, Length 12
    char    lp_offerrem8        [  12];    char    _lp_offerrem8        ;    // [long  ,   12] LP매도호가수량8                 StartPos 243, Length 12
    char    lp_bidrem8          [  12];    char    _lp_bidrem8          ;    // [long  ,   12] LP매수호가수량8                 StartPos 256, Length 12
    char    lp_offerrem9        [  12];    char    _lp_offerrem9        ;    // [long  ,   12] LP매도호가수량9                 StartPos 269, Length 12
    char    lp_bidrem9          [  12];    char    _lp_bidrem9          ;    // [long  ,   12] LP매수호가수량9                 StartPos 282, Length 12
    char    lp_offerrem10       [  12];    char    _lp_offerrem10       ;    // [long  ,   12] LP매도호가수량10                StartPos 295, Length 12
    char    lp_bidrem10         [  12];    char    _lp_bidrem10         ;    // [long  ,   12] LP매수호가수량10                StartPos 308, Length 12
    char    jnilclose           [   8];    char    _jnilclose           ;    // [long  ,    8] 전일종가                        StartPos 321, Length 8
    char    offerho1            [   8];    char    _offerho1            ;    // [long  ,    8] 매도호가1                       StartPos 330, Length 8
    char    bidho1              [   8];    char    _bidho1              ;    // [long  ,    8] 매수호가1                       StartPos 339, Length 8
    char    offerrem1           [  12];    char    _offerrem1           ;    // [long  ,   12] 매도호가수량1                   StartPos 348, Length 12
    char    bidrem1             [  12];    char    _bidrem1             ;    // [long  ,   12] 매수호가수량1                   StartPos 361, Length 12
    char    preoffercha1        [  12];    char    _preoffercha1        ;    // [long  ,   12] 직전매도대비수량1               StartPos 374, Length 12
    char    prebidcha1          [  12];    char    _prebidcha1          ;    // [long  ,   12] 직전매수대비수량1               StartPos 387, Length 12
    char    offerho2            [   8];    char    _offerho2            ;    // [long  ,    8] 매도호가2                       StartPos 400, Length 8
    char    bidho2              [   8];    char    _bidho2              ;    // [long  ,    8] 매수호가2                       StartPos 409, Length 8
    char    offerrem2           [  12];    char    _offerrem2           ;    // [long  ,   12] 매도호가수량2                   StartPos 418, Length 12
    char    bidrem2             [  12];    char    _bidrem2             ;    // [long  ,   12] 매수호가수량2                   StartPos 431, Length 12
    char    preoffercha2        [  12];    char    _preoffercha2        ;    // [long  ,   12] 직전매도대비수량2               StartPos 444, Length 12
    char    prebidcha2          [  12];    char    _prebidcha2          ;    // [long  ,   12] 직전매수대비수량2               StartPos 457, Length 12
    char    offerho3            [   8];    char    _offerho3            ;    // [long  ,    8] 매도호가3                       StartPos 470, Length 8
    char    bidho3              [   8];    char    _bidho3              ;    // [long  ,    8] 매수호가3                       StartPos 479, Length 8
    char    offerrem3           [  12];    char    _offerrem3           ;    // [long  ,   12] 매도호가수량3                   StartPos 488, Length 12
    char    bidrem3             [  12];    char    _bidrem3             ;    // [long  ,   12] 매수호가수량3                   StartPos 501, Length 12
    char    preoffercha3        [  12];    char    _preoffercha3        ;    // [long  ,   12] 직전매도대비수량3               StartPos 514, Length 12
    char    prebidcha3          [  12];    char    _prebidcha3          ;    // [long  ,   12] 직전매수대비수량3               StartPos 527, Length 12
    char    offerho4            [   8];    char    _offerho4            ;    // [long  ,    8] 매도호가4                       StartPos 540, Length 8
    char    bidho4              [   8];    char    _bidho4              ;    // [long  ,    8] 매수호가4                       StartPos 549, Length 8
    char    offerrem4           [  12];    char    _offerrem4           ;    // [long  ,   12] 매도호가수량4                   StartPos 558, Length 12
    char    bidrem4             [  12];    char    _bidrem4             ;    // [long  ,   12] 매수호가수량4                   StartPos 571, Length 12
    char    preoffercha4        [  12];    char    _preoffercha4        ;    // [long  ,   12] 직전매도대비수량4               StartPos 584, Length 12
    char    prebidcha4          [  12];    char    _prebidcha4          ;    // [long  ,   12] 직전매수대비수량4               StartPos 597, Length 12
    char    offerho5            [   8];    char    _offerho5            ;    // [long  ,    8] 매도호가5                       StartPos 610, Length 8
    char    bidho5              [   8];    char    _bidho5              ;    // [long  ,    8] 매수호가5                       StartPos 619, Length 8
    char    offerrem5           [  12];    char    _offerrem5           ;    // [long  ,   12] 매도호가수량5                   StartPos 628, Length 12
    char    bidrem5             [  12];    char    _bidrem5             ;    // [long  ,   12] 매수호가수량5                   StartPos 641, Length 12
    char    preoffercha5        [  12];    char    _preoffercha5        ;    // [long  ,   12] 직전매도대비수량5               StartPos 654, Length 12
    char    prebidcha5          [  12];    char    _prebidcha5          ;    // [long  ,   12] 직전매수대비수량5               StartPos 667, Length 12
    char    offerho6            [   8];    char    _offerho6            ;    // [long  ,    8] 매도호가6                       StartPos 680, Length 8
    char    bidho6              [   8];    char    _bidho6              ;    // [long  ,    8] 매수호가6                       StartPos 689, Length 8
    char    offerrem6           [  12];    char    _offerrem6           ;    // [long  ,   12] 매도호가수량6                   StartPos 698, Length 12
    char    bidrem6             [  12];    char    _bidrem6             ;    // [long  ,   12] 매수호가수량6                   StartPos 711, Length 12
    char    preoffercha6        [  12];    char    _preoffercha6        ;    // [long  ,   12] 직전매도대비수량6               StartPos 724, Length 12
    char    prebidcha6          [  12];    char    _prebidcha6          ;    // [long  ,   12] 직전매수대비수량6               StartPos 737, Length 12
    char    offerho7            [   8];    char    _offerho7            ;    // [long  ,    8] 매도호가7                       StartPos 750, Length 8
    char    bidho7              [   8];    char    _bidho7              ;    // [long  ,    8] 매수호가7                       StartPos 759, Length 8
    char    offerrem7           [  12];    char    _offerrem7           ;    // [long  ,   12] 매도호가수량7                   StartPos 768, Length 12
    char    bidrem7             [  12];    char    _bidrem7             ;    // [long  ,   12] 매수호가수량7                   StartPos 781, Length 12
    char    preoffercha7        [  12];    char    _preoffercha7        ;    // [long  ,   12] 직전매도대비수량7               StartPos 794, Length 12
    char    prebidcha7          [  12];    char    _prebidcha7          ;    // [long  ,   12] 직전매수대비수량7               StartPos 807, Length 12
    char    offerho8            [   8];    char    _offerho8            ;    // [long  ,    8] 매도호가8                       StartPos 820, Length 8
    char    bidho8              [   8];    char    _bidho8              ;    // [long  ,    8] 매수호가8                       StartPos 829, Length 8
    char    offerrem8           [  12];    char    _offerrem8           ;    // [long  ,   12] 매도호가수량8                   StartPos 838, Length 12
    char    bidrem8             [  12];    char    _bidrem8             ;    // [long  ,   12] 매수호가수량8                   StartPos 851, Length 12
    char    preoffercha8        [  12];    char    _preoffercha8        ;    // [long  ,   12] 직전매도대비수량8               StartPos 864, Length 12
    char    prebidcha8          [  12];    char    _prebidcha8          ;    // [long  ,   12] 직전매수대비수량8               StartPos 877, Length 12
    char    offerho9            [   8];    char    _offerho9            ;    // [long  ,    8] 매도호가9                       StartPos 890, Length 8
    char    bidho9              [   8];    char    _bidho9              ;    // [long  ,    8] 매수호가9                       StartPos 899, Length 8
    char    offerrem9           [  12];    char    _offerrem9           ;    // [long  ,   12] 매도호가수량9                   StartPos 908, Length 12
    char    bidrem9             [  12];    char    _bidrem9             ;    // [long  ,   12] 매수호가수량9                   StartPos 921, Length 12
    char    preoffercha9        [  12];    char    _preoffercha9        ;    // [long  ,   12] 직전매도대비수량9               StartPos 934, Length 12
    char    prebidcha9          [  12];    char    _prebidcha9          ;    // [long  ,   12] 직전매수대비수량9               StartPos 947, Length 12
    char    offerho10           [   8];    char    _offerho10           ;    // [long  ,    8] 매도호가10                      StartPos 960, Length 8
    char    bidho10             [   8];    char    _bidho10             ;    // [long  ,    8] 매수호가10                      StartPos 969, Length 8
    char    offerrem10          [  12];    char    _offerrem10          ;    // [long  ,   12] 매도호가수량10                  StartPos 978, Length 12
    char    bidrem10            [  12];    char    _bidrem10            ;    // [long  ,   12] 매수호가수량10                  StartPos 991, Length 12
    char    preoffercha10       [  12];    char    _preoffercha10       ;    // [long  ,   12] 직전매도대비수량10              StartPos 1004, Length 12
    char    prebidcha10         [  12];    char    _prebidcha10         ;    // [long  ,   12] 직전매수대비수량10              StartPos 1017, Length 12
    char    offer               [  12];    char    _offer               ;    // [long  ,   12] 매도호가수량합                  StartPos 1030, Length 12
    char    bid                 [  12];    char    _bid                 ;    // [long  ,   12] 매수호가수량합                  StartPos 1043, Length 12
    char    preoffercha         [  12];    char    _preoffercha         ;    // [long  ,   12] 직전매도대비수량합              StartPos 1056, Length 12
    char    prebidcha           [  12];    char    _prebidcha           ;    // [long  ,   12] 직전매수대비수량합              StartPos 1069, Length 12
    char    hotime              [   8];    char    _hotime              ;    // [string,    8] 수신시간                        StartPos 1082, Length 8
    char    yeprice             [   8];    char    _yeprice             ;    // [long  ,    8] 예상체결가격                    StartPos 1091, Length 8
    char    yevolume            [  12];    char    _yevolume            ;    // [long  ,   12] 예상체결수량                    StartPos 1100, Length 12
    char    yesign              [   1];    char    _yesign              ;    // [string,    1] 예상체결전일구분                StartPos 1113, Length 1
    char    yechange            [   8];    char    _yechange            ;    // [long  ,    8] 예상체결전일대비                StartPos 1115, Length 8
    char    yediff              [   6];    char    _yediff              ;    // [float ,  6.2] 예상체결등락율                  StartPos 1124, Length 6
    char    tmoffer             [  12];    char    _tmoffer             ;    // [long  ,   12] 시간외매도잔량                  StartPos 1131, Length 12
    char    tmbid               [  12];    char    _tmbid               ;    // [long  ,   12] 시간외매수잔량                  StartPos 1144, Length 12
    char    ho_status           [   1];    char    _ho_status           ;    // [string,    1] 동시구분                        StartPos 1157, Length 1
    char    shcode              [   6];    char    _shcode              ;    // [string,    6] 단축코드                        StartPos 1159, Length 6
    char    uplmtprice          [   8];    char    _uplmtprice          ;    // [long  ,    8] 상한가                          StartPos 1166, Length 8
    char    dnlmtprice          [   8];    char    _dnlmtprice          ;    // [long  ,    8] 하한가                          StartPos 1175, Length 8
    char    open                [   8];    char    _open                ;    // [long  ,    8] 시가                            StartPos 1184, Length 8
    char    high                [   8];    char    _high                ;    // [long  ,    8] 고가                            StartPos 1193, Length 8
    char    low                 [   8];    char    _low                 ;    // [long  ,    8] 저가                            StartPos 1202, Length 8
} T1906OutBlock;

//------------------------------------------------------------------------------
// 기업 정보 요약 (t3320)
//------------------------------------------------------------------------------
typedef struct {
    char    gicode              [   7];    char    _gicode              ;    // [string,    7] 종목코드                        StartPos 0, Length 7
} T3320InBlock;

typedef struct {
    char    upgubunnm           [  20];    char    _upgubunnm           ;    // [string,   20] 업종구분명                      StartPos 0, Length 20
    char    sijangcd            [   1];    char    _sijangcd            ;    // [string,    1] 시장구분                        StartPos 21, Length 1
    char    marketnm            [  10];    char    _marketnm            ;    // [string,   10] 시장구분명                      StartPos 23, Length 10
    char    company             [ 100];    char    _company             ;    // [string,  100] 한글기업명                      StartPos 34, Length 100
    char    baddress            [ 100];    char    _baddress            ;    // [string,  100] 본사주소                        StartPos 135, Length 100
    char    btelno              [  20];    char    _btelno              ;    // [string,   20] 본사전화번호                    StartPos 236, Length 20
    char    gsyyyy              [   4];    char    _gsyyyy              ;    // [string,    4] 최근결산년도                    StartPos 257, Length 4
    char    gsmm                [   2];    char    _gsmm                ;    // [string,    2] 결산월                          StartPos 262, Length 2
    char    gsym                [   6];    char    _gsym                ;    // [string,    6] 최근결산년월                    StartPos 265, Length 6
    char    lstprice            [  12];    char    _lstprice            ;    // [long  ,   12] 주당액면가                      StartPos 272, Length 12
    char    gstock              [  12];    char    _gstock              ;    // [long  ,   12] 주식수                          StartPos 285, Length 12
    char    homeurl             [  50];    char    _homeurl             ;    // [string,   50] Homepage                        StartPos 298, Length 50
    char    grdnm               [  30];    char    _grdnm               ;    // [string,   30] 그룹명                          StartPos 349, Length 30
    char    foreignratio        [   6];    char    _foreignratio        ;    // [float ,  6.2] 외국인                          StartPos 380, Length 6
    char    irtel               [  30];    char    _irtel               ;    // [string,   30] 주담전화                        StartPos 387, Length 30
    char    capital             [  12];    char    _capital             ;    // [float ,   12] 자본금                          StartPos 418, Length 12
    char    sigavalue           [  12];    char    _sigavalue           ;    // [float ,   12] 시가총액                        StartPos 431, Length 12
    char    cashsis             [  12];    char    _cashsis             ;    // [float ,   12] 배당금                          StartPos 444, Length 12
    char    cashrate            [  13];    char    _cashrate            ;    // [float , 13.2] 배당수익율                      StartPos 457, Length 13
    char    price               [   8];    char    _price               ;    // [long  ,    8] 현재가                          StartPos 471, Length 8
    char    jnilclose           [   8];    char    _jnilclose           ;    // [long  ,    8] 전일종가                        StartPos 480, Length 8
    char    padding_added       [   6]; // 수동 추가한 패딩
} T3320OutBlock;

typedef struct {
    char    gicode              [   7];    char    _gicode              ;    // [string,    7] 기업코드                        StartPos 0, Length 7
    char    gsym                [   6];    char    _gsym                ;    // [string,    6] 결산년월                        StartPos 8, Length 6
    char    gsgb                [   1];    char    _gsgb                ;    // [string,    1] 결산구분                        StartPos 15, Length 1
    char    per                 [  13];    char    _per                 ;    // [float , 13.2] PER                             StartPos 17, Length 13
    char    eps                 [  13];    char    _eps                 ;    // [float ,   13] EPS                             StartPos 31, Length 13
    char    pbr                 [  13];    char    _pbr                 ;    // [float , 13.2] PBR                             StartPos 45, Length 13
    char    roa                 [  13];    char    _roa                 ;    // [float , 13.2] ROA                             StartPos 59, Length 13
    char    roe                 [  13];    char    _roe                 ;    // [float , 13.2] ROE                             StartPos 73, Length 13
    char    ebitda              [  13];    char    _ebitda              ;    // [float , 13.2] EBITDA                          StartPos 87, Length 13
    char    evebitda            [  13];    char    _evebitda            ;    // [float , 13.2] EVEBITDA                        StartPos 101, Length 13
    char    par                 [  13];    char    _par                 ;    // [float , 13.2] 액면가                          StartPos 115, Length 13
    char    sps                 [  13];    char    _sps                 ;    // [float , 13.2] SPS                             StartPos 129, Length 13
    char    cps                 [  13];    char    _cps                 ;    // [float , 13.2] CPS                             StartPos 143, Length 13
    char    bps                 [  13];    char    _bps                 ;    // [float ,   13] BPS                             StartPos 157, Length 13
    char    tper               [  13];    char    _t_per               ;    // [float , 13.2] T.PER                           StartPos 171, Length 13
    char    teps               [  13];    char    _t_eps               ;    // [float ,   13] T.EPS                           StartPos 185, Length 13
    char    peg                 [  13];    char    _peg                 ;    // [float , 13.2] PEG                             StartPos 199, Length 13
    char    tpeg               [  13];    char    _t_peg               ;    // [float , 13.2] T.PEG                           StartPos 213, Length 13
    char    tgsym              [   6];    char    _t_gsym              ;    // [string,    6] 최근분기년도                    StartPos 227, Length 6
} T3320OutBlock1;

//------------------------------------------------------------------------------
// 재무 순위 종합 (t3341)
//------------------------------------------------------------------------------
typedef struct {
    char    gubun               [   1];    char    _gubun               ;    // [string,    1] 시장구분                        StartPos 0, Length 1
    char    gubun1              [   1];    char    _gubun1              ;    // [string,    1] 순위구분(1:매출액증가율2:영업이 StartPos 2, Length 1
    char    gubun2              [   1];    char    _gubun2              ;    // [string,    1] 대비구분                        StartPos 4, Length 1
    char    idx                 [   4];    char    _idx                 ;    // [long  ,    4] IDX                             StartPos 6, Length 4
} T3341InBlock;

typedef struct {
    char    cnt                 [   4];    char    _cnt                 ;    // [long  ,    4] CNT                             StartPos 0, Length 4
    char    idx                 [   4];    char    _idx                 ;    // [long  ,    4] IDX                             StartPos 5, Length 4
} T3341OutBlock;

typedef struct {
    char    rank                [   4];    char    _rank                ;    // [long  ,    4] 순위                            StartPos 0, Length 4
    char    hname               [  20];    char    _hname               ;    // [string,   20] 기업명                          StartPos 5, Length 20
    char    salesgrowth         [  12];    char    _salesgrowth         ;    // [long  ,   12] 매출액증가율                    StartPos 26, Length 12
    char    operatingincomegrowt[  12];    char    _operatingincomegrowt;    // [long  ,   12] 영업이익증가율                  StartPos 39, Length 12
    char    ordinaryincomegrowth[  12];    char    _ordinaryincomegrowth;    // [long  ,   12] 경상이익증가율                  StartPos 52, Length 12
    char    liabilitytoequity   [  12];    char    _liabilitytoequity   ;    // [long  ,   12] 부채비율                        StartPos 65, Length 12
    char    enterpriseratio     [  12];    char    _enterpriseratio     ;    // [long  ,   12] 유보율                          StartPos 78, Length 12
    char    eps                 [  12];    char    _eps                 ;    // [long  ,   12] EPS                             StartPos 91, Length 12
    char    bps                 [  12];    char    _bps                 ;    // [long  ,   12] BPS                             StartPos 104, Length 12
    char    roe                 [  12];    char    _roe                 ;    // [long  ,   12] ROE                             StartPos 117, Length 12
    char    shcode              [   6];    char    _shcode              ;    // [string,    6] 종목코드                        StartPos 130, Length 6
    char    per                 [  13];    char    _per                 ;    // [float , 13.2] PER                             StartPos 137, Length 13
    char    pbr                 [  13];    char    _pbr                 ;    // [float , 13.2] PBR                             StartPos 151, Length 13
    char    peg                 [  13];    char    _peg                 ;    // [float , 13.2] PEG                             StartPos 165, Length 13
} T3341OutBlock1;

//------------------------------------------------------------------------------
// 현물 멀티 현재가 조회 (t8407)
//------------------------------------------------------------------------------
typedef struct {
    char    nrec                [   3];    char    _nrec                ;    // [long  ,    3] 건수                            StartPos 0, Length 3
    char    shcode              [ 300];    char    _shcode              ;    // [string,  300] 종목코드                        StartPos 4, Length 300
} T8407InBlock;

typedef struct {
    char    shcode              [   6];    char    _shcode              ;    // [string,    6] 종목코드                        StartPos 0, Length 6
    char    hname               [  40];    char    _hname               ;    // [string,   40] 종목명                          StartPos 7, Length 40
    char    price               [   8];    char    _price               ;    // [long  ,    8] 현재가                          StartPos 48, Length 8
    char    sign                [   1];    char    _sign                ;    // [string,    1] 전일대비구분                    StartPos 57, Length 1
    char    change              [   8];    char    _change              ;    // [long  ,    8] 전일대비                        StartPos 59, Length 8
    char    diff                [   6];    char    _diff                ;    // [float ,  6.2] 등락율                          StartPos 68, Length 6
    char    volume              [  12];    char    _volume              ;    // [long  ,   12] 누적거래량                      StartPos 75, Length 12
    char    offerho             [   8];    char    _offerho             ;    // [long  ,    8] 매도호가                        StartPos 88, Length 8
    char    bidho               [   8];    char    _bidho               ;    // [long  ,    8] 매수호가                        StartPos 97, Length 8
    char    cvolume             [   8];    char    _cvolume             ;    // [long  ,    8] 체결수량                        StartPos 106, Length 8
    char    chdegree            [   9];    char    _chdegree            ;    // [float ,  9.2] 체결강도                        StartPos 115, Length 9
    char    open                [   8];    char    _open                ;    // [long  ,    8] 시가                            StartPos 125, Length 8
    char    high                [   8];    char    _high                ;    // [long  ,    8] 고가                            StartPos 134, Length 8
    char    low                 [   8];    char    _low                 ;    // [long  ,    8] 저가                            StartPos 143, Length 8
    char    value               [  12];    char    _value               ;    // [long  ,   12] 거래대금(백만)                  StartPos 152, Length 12
    char    offerrem            [  12];    char    _offerrem            ;    // [long  ,   12] 우선매도잔량                    StartPos 165, Length 12
    char    bidrem              [  12];    char    _bidrem              ;    // [long  ,   12] 우선매수잔량                    StartPos 178, Length 12
    char    totofferrem         [  12];    char    _totofferrem         ;    // [long  ,   12] 총매도잔량                      StartPos 191, Length 12
    char    totbidrem           [  12];    char    _totbidrem           ;    // [long  ,   12] 총매수잔량                      StartPos 204, Length 12
    char    jnilclose           [   8];    char    _jnilclose           ;    // [long  ,    8] 전일종가                        StartPos 217, Length 8
    char    uplmtprice          [   8];    char    _uplmtprice          ;    // [long  ,    8] 상한가                          StartPos 226, Length 8
    char    dnlmtprice          [   8];    char    _dnlmtprice          ;    // [long  ,    8] 하한가                          StartPos 235, Length 8
} T8407OutBlock1;

//------------------------------------------------------------------------------
// 현물 차트 틱 (t8411)
//------------------------------------------------------------------------------
typedef struct {
    char    shcode              [   6];    char    _shcode              ;    // [string,    6] 단축코드                        StartPos 0, Length 6
    char    ncnt                [   4];    char    _ncnt                ;    // [long  ,    4] 단위(n틱)                       StartPos 7, Length 4
    char    qrycnt              [   4];    char    _qrycnt              ;    // [long  ,    4] 요청건수(최대-압축:2000비압축:5 StartPos 12, Length 4
    char    nday                [   1];    char    _nday                ;    // [string,    1] 조회영업일수(0:미사용1>=사용)   StartPos 17, Length 1
    char    sdate               [   8];    char    _sdate               ;    // [string,    8] 시작일자                        StartPos 19, Length 8
    char    stime               [   6];    char    _stime               ;    // [string,    6] 시작시간(현재미사용)            StartPos 28, Length 6
    char    edate               [   8];    char    _edate               ;    // [string,    8] 종료일자                        StartPos 35, Length 8
    char    etime               [   6];    char    _etime               ;    // [string,    6] 종료시간(현재미사용)            StartPos 44, Length 6
    char    cts_date            [   8];    char    _cts_date            ;    // [string,    8] 연속일자                        StartPos 51, Length 8
    char    cts_time            [  10];    char    _cts_time            ;    // [string,   10] 연속시간                        StartPos 60, Length 10
    char    comp_yn             [   1];    char    _comp_yn             ;    // [string,    1] 압축여부(Y:압축N:비압축)        StartPos 71, Length 1
} T8411InBlock;

typedef struct {
    char    shcode              [   6];    char    _shcode              ;    // [string,    6] 단축코드                        StartPos 0, Length 6
    char    jisiga              [   8];    char    _jisiga              ;    // [long  ,    8] 전일시가                        StartPos 7, Length 8
    char    jihigh              [   8];    char    _jihigh              ;    // [long  ,    8] 전일고가                        StartPos 16, Length 8
    char    jilow               [   8];    char    _jilow               ;    // [long  ,    8] 전일저가                        StartPos 25, Length 8
    char    jiclose             [   8];    char    _jiclose             ;    // [long  ,    8] 전일종가                        StartPos 34, Length 8
    char    jivolume            [  12];    char    _jivolume            ;    // [long  ,   12] 전일거래량                      StartPos 43, Length 12
    char    disiga              [   8];    char    _disiga              ;    // [long  ,    8] 당일시가                        StartPos 56, Length 8
    char    dihigh              [   8];    char    _dihigh              ;    // [long  ,    8] 당일고가                        StartPos 65, Length 8
    char    dilow               [   8];    char    _dilow               ;    // [long  ,    8] 당일저가                        StartPos 74, Length 8
    char    diclose             [   8];    char    _diclose             ;    // [long  ,    8] 당일종가                        StartPos 83, Length 8
    char    highend             [   8];    char    _highend             ;    // [long  ,    8] 상한가                          StartPos 92, Length 8
    char    lowend              [   8];    char    _lowend              ;    // [long  ,    8] 하한가                          StartPos 101, Length 8
    char    cts_date            [   8];    char    _cts_date            ;    // [string,    8] 연속일자                        StartPos 110, Length 8
    char    cts_time            [  10];    char    _cts_time            ;    // [string,   10] 연속시간                        StartPos 119, Length 10
    char    s_time              [   6];    char    _s_time              ;    // [string,    6] 장시작시간(HHMMSS)              StartPos 130, Length 6
    char    e_time              [   6];    char    _e_time              ;    // [string,    6] 장종료시간(HHMMSS)              StartPos 137, Length 6
    char    dshmin              [   2];    char    _dshmin              ;    // [string,    2] 동시호가처리시간(MM:분)         StartPos 144, Length 2
    char    rec_count           [   7];    char    _rec_count           ;    // [long  ,    7] 레코드카운트                    StartPos 147, Length 7
} T8411OutBlock;

typedef struct {
    char    date                [   8];    char    _date                ;    // [string,    8] 날짜                            StartPos 0, Length 8
    char    time                [  10];    char    _time                ;    // [string,   10] 시간                            StartPos 9, Length 10
    char    open                [   8];    char    _open                ;    // [long  ,    8] 시가                            StartPos 20, Length 8
    char    high                [   8];    char    _high                ;    // [long  ,    8] 고가                            StartPos 29, Length 8
    char    low                 [   8];    char    _low                 ;    // [long  ,    8] 저가                            StartPos 38, Length 8
    char    close               [   8];    char    _close               ;    // [long  ,    8] 종가                            StartPos 47, Length 8
    char    jdiff_vol           [  12];    char    _jdiff_vol           ;    // [long  ,   12] 거래량                          StartPos 56, Length 12
    char    jongchk             [  13];    char    _jongchk             ;    // [long  ,   13] 수정구분                        StartPos 69, Length 13
    char    rate                [   6];    char    _rate                ;    // [double,  6.2] 수정비율                        StartPos 83, Length 6
    char    pricechk            [  13];    char    _pricechk            ;    // [long  ,   13] 수정주가반영항목                StartPos 90, Length 13
} T8411OutBlock1;

//------------------------------------------------------------------------------
// 현물 차트 분 (t8412)
//------------------------------------------------------------------------------
typedef struct {
    char    shcode              [   6];    char    _shcode              ;    // [string,    6] 단축코드                        StartPos 0, Length 6
    char    ncnt                [   4];    char    _ncnt                ;    // [long  ,    4] 단위(n분)                       StartPos 7, Length 4
    char    qrycnt              [   4];    char    _qrycnt              ;    // [long  ,    4] 요청건수(최대-압축:2000비압축:5 StartPos 12, Length 4
    char    nday                [   1];    char    _nday                ;    // [string,    1] 조회영업일수(0:미사용1>=사용)   StartPos 17, Length 1
    char    sdate               [   8];    char    _sdate               ;    // [string,    8] 시작일자                        StartPos 19, Length 8
    char    stime               [   6];    char    _stime               ;    // [string,    6] 시작시간(현재미사용)            StartPos 28, Length 6
    char    edate               [   8];    char    _edate               ;    // [string,    8] 종료일자                        StartPos 35, Length 8
    char    etime               [   6];    char    _etime               ;    // [string,    6] 종료시간(현재미사용)            StartPos 44, Length 6
    char    cts_date            [   8];    char    _cts_date            ;    // [string,    8] 연속일자                        StartPos 51, Length 8
    char    cts_time            [  10];    char    _cts_time            ;    // [string,   10] 연속시간                        StartPos 60, Length 10
    char    comp_yn             [   1];    char    _comp_yn             ;    // [string,    1] 압축여부(Y:압축N:비압축)        StartPos 71, Length 1
} T8412InBlock;

typedef struct {
    char    shcode              [   6];    char    _shcode              ;    // [string,    6] 단축코드                        StartPos 0, Length 6
    char    jisiga              [   8];    char    _jisiga              ;    // [long  ,    8] 전일시가                        StartPos 7, Length 8
    char    jihigh              [   8];    char    _jihigh              ;    // [long  ,    8] 전일고가                        StartPos 16, Length 8
    char    jilow               [   8];    char    _jilow               ;    // [long  ,    8] 전일저가                        StartPos 25, Length 8
    char    jiclose             [   8];    char    _jiclose             ;    // [long  ,    8] 전일종가                        StartPos 34, Length 8
    char    jivolume            [  12];    char    _jivolume            ;    // [long  ,   12] 전일거래량                      StartPos 43, Length 12
    char    disiga              [   8];    char    _disiga              ;    // [long  ,    8] 당일시가                        StartPos 56, Length 8
    char    dihigh              [   8];    char    _dihigh              ;    // [long  ,    8] 당일고가                        StartPos 65, Length 8
    char    dilow               [   8];    char    _dilow               ;    // [long  ,    8] 당일저가                        StartPos 74, Length 8
    char    diclose             [   8];    char    _diclose             ;    // [long  ,    8] 당일종가                        StartPos 83, Length 8
    char    highend             [   8];    char    _highend             ;    // [long  ,    8] 상한가                          StartPos 92, Length 8
    char    lowend              [   8];    char    _lowend              ;    // [long  ,    8] 하한가                          StartPos 101, Length 8
    char    cts_date            [   8];    char    _cts_date            ;    // [string,    8] 연속일자                        StartPos 110, Length 8
    char    cts_time            [  10];    char    _cts_time            ;    // [string,   10] 연속시간                        StartPos 119, Length 10
    char    s_time              [   6];    char    _s_time              ;    // [string,    6] 장시작시간(HHMMSS)              StartPos 130, Length 6
    char    e_time              [   6];    char    _e_time              ;    // [string,    6] 장종료시간(HHMMSS)              StartPos 137, Length 6
    char    dshmin              [   2];    char    _dshmin              ;    // [string,    2] 동시호가처리시간(MM:분)         StartPos 144, Length 2
    char    rec_count           [   7];    char    _rec_count           ;    // [long  ,    7] 레코드카운트                    StartPos 147, Length 7
} T8412OutBlock;

typedef struct {
    char    date                [   8];    char    _date                ;    // [string,    8] 날짜                            StartPos 0, Length 8
    char    time                [  10];    char    _time                ;    // [string,   10] 시간                            StartPos 9, Length 10
    char    open                [   8];    char    _open                ;    // [long  ,    8] 시가                            StartPos 20, Length 8
    char    high                [   8];    char    _high                ;    // [long  ,    8] 고가                            StartPos 29, Length 8
    char    low                 [   8];    char    _low                 ;    // [long  ,    8] 저가                            StartPos 38, Length 8
    char    close               [   8];    char    _close               ;    // [long  ,    8] 종가                            StartPos 47, Length 8
    char    jdiff_vol           [  12];    char    _jdiff_vol           ;    // [long  ,   12] 거래량                          StartPos 56, Length 12
    char    value               [  12];    char    _value               ;    // [long  ,   12] 거래대금                        StartPos 69, Length 12
    char    jongchk             [  13];    char    _jongchk             ;    // [long  ,   13] 수정구분                        StartPos 82, Length 13
    char    rate                [   6];    char    _rate                ;    // [double,  6.2] 수정비율                        StartPos 96, Length 6
    char    sign                [   1];    char    _sign                ;    // [string,    1] 종가등락구분(1:상한2:상승3:보합 StartPos 103, Length 1
} T8412OutBlock1;

//------------------------------------------------------------------------------
// 현물 차트 일주월 (t8413)
//------------------------------------------------------------------------------
typedef struct {
    char    shcode              [   6];    char    _shcode              ;    // [string,    6] 단축코드                        StartPos 0, Length 6
    char    gubun               [   1];    char    _gubun               ;    // [string,    1] 주기구분(2:일3:주4:월)          StartPos 7, Length 1
    char    qrycnt              [   4];    char    _qrycnt              ;    // [long  ,    4] 요청건수(최대-압축:2000비압축:5 StartPos 9, Length 4
    char    sdate               [   8];    char    _sdate               ;    // [string,    8] 시작일자                        StartPos 14, Length 8
    char    edate               [   8];    char    _edate               ;    // [string,    8] 종료일자                        StartPos 23, Length 8
    char    cts_date            [   8];    char    _cts_date            ;    // [string,    8] 연속일자                        StartPos 32, Length 8
    char    comp_yn             [   1];    char    _comp_yn             ;    // [string,    1] 압축여부(Y:압축N:비압축)        StartPos 41, Length 1
} T8413InBlock;

typedef struct {
    char    shcode              [   6];    char    _shcode              ;    // [string,    6] 단축코드                        StartPos 0, Length 6
    char    jisiga              [   8];    char    _jisiga              ;    // [long  ,    8] 전일시가                        StartPos 7, Length 8
    char    jihigh              [   8];    char    _jihigh              ;    // [long  ,    8] 전일고가                        StartPos 16, Length 8
    char    jilow               [   8];    char    _jilow               ;    // [long  ,    8] 전일저가                        StartPos 25, Length 8
    char    jiclose             [   8];    char    _jiclose             ;    // [long  ,    8] 전일종가                        StartPos 34, Length 8
    char    jivolume            [  12];    char    _jivolume            ;    // [long  ,   12] 전일거래량                      StartPos 43, Length 12
    char    disiga              [   8];    char    _disiga              ;    // [long  ,    8] 당일시가                        StartPos 56, Length 8
    char    dihigh              [   8];    char    _dihigh              ;    // [long  ,    8] 당일고가                        StartPos 65, Length 8
    char    dilow               [   8];    char    _dilow               ;    // [long  ,    8] 당일저가                        StartPos 74, Length 8
    char    diclose             [   8];    char    _diclose             ;    // [long  ,    8] 당일종가                        StartPos 83, Length 8
    char    highend             [   8];    char    _highend             ;    // [long  ,    8] 상한가                          StartPos 92, Length 8
    char    lowend              [   8];    char    _lowend              ;    // [long  ,    8] 하한가                          StartPos 101, Length 8
    char    cts_date            [   8];    char    _cts_date            ;    // [string,    8] 연속일자                        StartPos 110, Length 8
    char    s_time              [   6];    char    _s_time              ;    // [string,    6] 장시작시간(HHMMSS)              StartPos 119, Length 6
    char    e_time              [   6];    char    _e_time              ;    // [string,    6] 장종료시간(HHMMSS)              StartPos 126, Length 6
    char    dshmin              [   2];    char    _dshmin              ;    // [string,    2] 동시호가처리시간(MM:분)         StartPos 133, Length 2
    char    rec_count           [   7];    char    _rec_count           ;    // [long  ,    7] 레코드카운트                    StartPos 136, Length 7
} T8413OutBlock;

typedef struct {
    char    date                [   8];    char    _date                ;    // [string,    8] 날짜                            StartPos 0, Length 8
    char    open                [   8];    char    _open                ;    // [long  ,    8] 시가                            StartPos 9, Length 8
    char    high                [   8];    char    _high                ;    // [long  ,    8] 고가                            StartPos 18, Length 8
    char    low                 [   8];    char    _low                 ;    // [long  ,    8] 저가                            StartPos 27, Length 8
    char    close               [   8];    char    _close               ;    // [long  ,    8] 종가                            StartPos 36, Length 8
    char    jdiff_vol           [  12];    char    _jdiff_vol           ;    // [long  ,   12] 거래량                          StartPos 45, Length 12
    char    value               [  12];    char    _value               ;    // [long  ,   12] 거래대금                        StartPos 58, Length 12
    char    jongchk             [  13];    char    _jongchk             ;    // [long  ,   13] 수정구분                        StartPos 71, Length 13
    char    rate                [   6];    char    _rate                ;    // [double,  6.2] 수정비율                        StartPos 85, Length 6
    char    pricechk            [  13];    char    _pricechk            ;    // [long  ,   13] 수정주가반영항목                StartPos 92, Length 13
    char    ratevalue           [  12];    char    _ratevalue           ;    // [long  ,   12] 수정비율반영거래대금            StartPos 106, Length 12
    char    sign                [   1];    char    _sign                ;    // [string,    1] 종가등락구분(1:상한2:상승3:보합 StartPos 119, Length 1
} T8413OutBlock1;

//------------------------------------------------------------------------------
// 증시 주변 자금 추이 (t8428)
//------------------------------------------------------------------------------
typedef struct {
    char    fdate               [   8];    char    _fdate               ;    // [string,    8] from일자                        StartPos 0, Length 8
    char    tdate               [   8];    char    _tdate               ;    // [string,    8] to일자                          StartPos 9, Length 8
    char    gubun               [   1];    char    _gubun               ;    // [string,    1] 구분                            StartPos 18, Length 1
    char    keyDate            [   8];    char    _key_date            ;    // [string,    8] 날짜                            StartPos 20, Length 8
    char    upcode              [   3];    char    _upcode              ;    // [string,    3] 업종코드                        StartPos 29, Length 3
    char    cnt                 [   3];    char    _cnt                 ;    // [string,    3] 조회건수                        StartPos 33, Length 3
} T8428InBlock;

typedef struct {
    char    date                [   8];    char    _date                ;    // [string,    8] 날짜CTS                         StartPos 0, Length 8
    char    idx                 [   4];    char    _idx                 ;    // [long  ,    4] IDX                             StartPos 9, Length 4
} T8428OutBlock;

typedef struct {
    char    date                [   8];    char    _date                ;    // [string,    8] 일자                            StartPos 0, Length 8
    char    jisu                [   7];    char    _jisu                ;    // [float ,  7.2] 지수                            StartPos 9, Length 7
    char    sign                [   1];    char    _sign                ;    // [string,    1] 대비구분                        StartPos 17, Length 1
    char    change              [   6];    char    _change              ;    // [float ,  6.2] 대비                            StartPos 19, Length 6
    char    diff                [   6];    char    _diff                ;    // [float ,  6.2] 등락율                          StartPos 26, Length 6
    char    volume              [  12];    char    _volume              ;    // [long  ,   12] 거래량                          StartPos 33, Length 12
    char    custmoney           [  12];    char    _custmoney           ;    // [long  ,   12] 고객예탁금_억원                 StartPos 46, Length 12
    char    yecha               [  12];    char    _yecha               ;    // [long  ,   12] 예탁증감_억원                   StartPos 59, Length 12
    char    vol                 [   6];    char    _vol                 ;    // [float ,  6.2] 회전율                          StartPos 72, Length 6
    char    outmoney            [  12];    char    _outmoney            ;    // [long  ,   12] 미수금_억원                     StartPos 79, Length 12
    char    trjango             [  12];    char    _trjango             ;    // [long  ,   12] 신용잔고_억원                   StartPos 92, Length 12
    char    futymoney           [  12];    char    _futymoney           ;    // [long  ,   12] 선물예수금_억원                 StartPos 105, Length 12
    char    stkmoney            [   8];    char    _stkmoney            ;    // [long  ,    8] 주식형_억원                     StartPos 118, Length 8
    char    mstkmoney           [   8];    char    _mstkmoney           ;    // [long  ,    8] 혼합형_억원(주식)               StartPos 127, Length 8
    char    mbndmoney           [   8];    char    _mbndmoney           ;    // [long  ,    8] 혼합형_억원(채권)               StartPos 136, Length 8
    char    bndmoney            [   8];    char    _bndmoney            ;    // [long  ,    8] 채권형_억원                     StartPos 145, Length 8
    char    bndsmoney           [   8];    char    _bndsmoney           ;    // [long  ,    8] 필러(구.단기채권)               StartPos 154, Length 8
    char    mmfmsoney            [   8];    char    _mmfmoney            ;    // [long  ,    8] MMF_억원(주식)                  StartPos 163, Length 8
} T8428OutBlock1;

//------------------------------------------------------------------------------
// 지수선물 조회 API용 (t8432)
//------------------------------------------------------------------------------
typedef struct {
    char    gubun               [   1];    // [string,    1] 구분                            StartPos 0, Length 1
} T8432InBlock;

typedef struct {
    char    hname               [  20];    // [string,   20] 종목명                          StartPos 0, Length 20
    char    shcode              [   8];    // [string,    8] 단축코드                        StartPos 20, Length 8
    char    expcode             [  12];    // [string,   12] 확장코드                        StartPos 28, Length 12
    char    uplmtprice          [   6];    // [float ,  6.2] 상한가                          StartPos 40, Length 6
    char    dnlmtprice          [   6];    // [float ,  6.2] 하한가                          StartPos 46, Length 6
    char    jnilclose           [   6];    // [float ,  6.2] 전일종가                        StartPos 52, Length 6
    char    jnilhigh            [   6];    // [float ,  6.2] 전일고가                        StartPos 58, Length 6
    char    jnillow             [   6];    // [float ,  6.2] 전일저가                        StartPos 64, Length 6
    char    recprice            [   6];    // [float ,  6.2] 기준가                          StartPos 70, Length 6
} T8432OutBlock;

//------------------------------------------------------------------------------
// 현물 종목조회 API용 (t8436)
//------------------------------------------------------------------------------
typedef struct {
    char    gubun[1];    //[string,    1] 구분(0:전체1:코스피2:코스닥)
} T8436InBlock;

typedef struct {
    char    hName[20];    //[string,   20] 종목명
    char    shCode[6];    //[string,    6] 단축코드
    char    expCode[12];    //[string,   12] 확장코드
    char    etfGubun[1];    //[string,    1] ETF구분(1:ETF2:ETN)
    char    upLmtPrice[8];    //[long  ,    8] 상한가
    char    dnLmtPrice[8];    //[long  ,    8] 하한가
    char    jnilClose[8];    //[long  ,    8] 전일가
    char    meMeDan[5];    //[string,    5] 주문수량단위
    char    recPrice[8];    //[long  ,    8] 기준가
    char    gubun[1];    //[string,    1] 구분(1:코스피2:코스닥)
    char    bu12Gubun[2];    //[string,    2] 증권그룹
    char    spacGubun[1];    //[string,    1] 기업인수목적회사여부(Y/N)
    char    filler[32];    //[string,   32] filler(미사용)
} T8436OutBlock;

//------------------------------------------------------------------------------
// 코스피 호가 잔량 (H1_)
//------------------------------------------------------------------------------
typedef struct {
    char    hotime              [   6];    char    _hotime              ;    // [string,    6] 호가시간                        StartPos 0, Length 6
    char    offerho1            [   7];    char    _offerho1            ;    // [long  ,    7] 매도호가1                       StartPos 7, Length 7
    char    bidho1              [   7];    char    _bidho1              ;    // [long  ,    7] 매수호가1                       StartPos 15, Length 7
    char    offerrem1           [   9];    char    _offerrem1           ;    // [long  ,    9] 매도호가잔량1                   StartPos 23, Length 9
    char    bidrem1             [   9];    char    _bidrem1             ;    // [long  ,    9] 매수호가잔량1                   StartPos 33, Length 9
    char    offerho2            [   7];    char    _offerho2            ;    // [long  ,    7] 매도호가2                       StartPos 43, Length 7
    char    bidho2              [   7];    char    _bidho2              ;    // [long  ,    7] 매수호가2                       StartPos 51, Length 7
    char    offerrem2           [   9];    char    _offerrem2           ;    // [long  ,    9] 매도호가잔량2                   StartPos 59, Length 9
    char    bidrem2             [   9];    char    _bidrem2             ;    // [long  ,    9] 매수호가잔량2                   StartPos 69, Length 9
    char    offerho3            [   7];    char    _offerho3            ;    // [long  ,    7] 매도호가3                       StartPos 79, Length 7
    char    bidho3              [   7];    char    _bidho3              ;    // [long  ,    7] 매수호가3                       StartPos 87, Length 7
    char    offerrem3           [   9];    char    _offerrem3           ;    // [long  ,    9] 매도호가잔량3                   StartPos 95, Length 9
    char    bidrem3             [   9];    char    _bidrem3             ;    // [long  ,    9] 매수호가잔량3                   StartPos 105, Length 9
    char    offerho4            [   7];    char    _offerho4            ;    // [long  ,    7] 매도호가4                       StartPos 115, Length 7
    char    bidho4              [   7];    char    _bidho4              ;    // [long  ,    7] 매수호가4                       StartPos 123, Length 7
    char    offerrem4           [   9];    char    _offerrem4           ;    // [long  ,    9] 매도호가잔량4                   StartPos 131, Length 9
    char    bidrem4             [   9];    char    _bidrem4             ;    // [long  ,    9] 매수호가잔량4                   StartPos 141, Length 9
    char    offerho5            [   7];    char    _offerho5            ;    // [long  ,    7] 매도호가5                       StartPos 151, Length 7
    char    bidho5              [   7];    char    _bidho5              ;    // [long  ,    7] 매수호가5                       StartPos 159, Length 7
    char    offerrem5           [   9];    char    _offerrem5           ;    // [long  ,    9] 매도호가잔량5                   StartPos 167, Length 9
    char    bidrem5             [   9];    char    _bidrem5             ;    // [long  ,    9] 매수호가잔량5                   StartPos 177, Length 9
    char    offerho6            [   7];    char    _offerho6            ;    // [long  ,    7] 매도호가6                       StartPos 187, Length 7
    char    bidho6              [   7];    char    _bidho6              ;    // [long  ,    7] 매수호가6                       StartPos 195, Length 7
    char    offerrem6           [   9];    char    _offerrem6           ;    // [long  ,    9] 매도호가잔량6                   StartPos 203, Length 9
    char    bidrem6             [   9];    char    _bidrem6             ;    // [long  ,    9] 매수호가잔량6                   StartPos 213, Length 9
    char    offerho7            [   7];    char    _offerho7            ;    // [long  ,    7] 매도호가7                       StartPos 223, Length 7
    char    bidho7              [   7];    char    _bidho7              ;    // [long  ,    7] 매수호가7                       StartPos 231, Length 7
    char    offerrem7           [   9];    char    _offerrem7           ;    // [long  ,    9] 매도호가잔량7                   StartPos 239, Length 9
    char    bidrem7             [   9];    char    _bidrem7             ;    // [long  ,    9] 매수호가잔량7                   StartPos 249, Length 9
    char    offerho8            [   7];    char    _offerho8            ;    // [long  ,    7] 매도호가8                       StartPos 259, Length 7
    char    bidho8              [   7];    char    _bidho8              ;    // [long  ,    7] 매수호가8                       StartPos 267, Length 7
    char    offerrem8           [   9];    char    _offerrem8           ;    // [long  ,    9] 매도호가잔량8                   StartPos 275, Length 9
    char    bidrem8             [   9];    char    _bidrem8             ;    // [long  ,    9] 매수호가잔량8                   StartPos 285, Length 9
    char    offerho9            [   7];    char    _offerho9            ;    // [long  ,    7] 매도호가9                       StartPos 295, Length 7
    char    bidho9              [   7];    char    _bidho9              ;    // [long  ,    7] 매수호가9                       StartPos 303, Length 7
    char    offerrem9           [   9];    char    _offerrem9           ;    // [long  ,    9] 매도호가잔량9                   StartPos 311, Length 9
    char    bidrem9             [   9];    char    _bidrem9             ;    // [long  ,    9] 매수호가잔량9                   StartPos 321, Length 9
    char    offerho10           [   7];    char    _offerho10           ;    // [long  ,    7] 매도호가10                      StartPos 331, Length 7
    char    bidho10             [   7];    char    _bidho10             ;    // [long  ,    7] 매수호가10                      StartPos 339, Length 7
    char    offerrem10          [   9];    char    _offerrem10          ;    // [long  ,    9] 매도호가잔량10                  StartPos 347, Length 9
    char    bidrem10            [   9];    char    _bidrem10            ;    // [long  ,    9] 매수호가잔량10                  StartPos 357, Length 9
    char    totofferrem         [   9];    char    _totofferrem         ;    // [long  ,    9] 총매도호가잔량                  StartPos 367, Length 9
    char    totbidrem           [   9];    char    _totbidrem           ;    // [long  ,    9] 총매수호가잔량                  StartPos 377, Length 9
    char    donsigubun          [   1];    char    _donsigubun          ;    // [string,    1] 동시호가구분                    StartPos 387, Length 1
    char    shcode              [   6];    char    _shcode              ;    // [string,    6] 단축코드                        StartPos 389, Length 6
    char    alloc_gubun         [   1];    char    _alloc_gubun         ;    // [string,    1] 배분적용구분                    StartPos 396, Length 1
    char    volume              [  12];    char    _volume              ;    // [long  ,   12] 누적거래량                      StartPos 398, Length 12
} H1_OutBlock;

//------------------------------------------------------------------------------
// 코스피 시간외 호가 잔량 (H2_)
//------------------------------------------------------------------------------
typedef struct {
    char    hotime              [   6];    char    _hotime              ;    // [string,    6] 호가시간                        StartPos 0, Length 6
    char    tmofferrem          [  12];    char    _tmofferrem          ;    // [long  ,   12] 시간외매도잔량                  StartPos 7, Length 12
    char    tmbidrem            [  12];    char    _tmbidrem            ;    // [long  ,   12] 시간외매수잔량                  StartPos 20, Length 12
    char    pretmoffercha       [  12];    char    _pretmoffercha       ;    // [long  ,   12] 시간외매도수량직전대비          StartPos 33, Length 12
    char    pretmbidcha         [  12];    char    _pretmbidcha         ;    // [long  ,   12] 시간외매수수량직전대비          StartPos 46, Length 12
    char    shcode              [   6];    char    _shcode              ;    // [string,    6] 단축코드                        StartPos 59, Length 6
} H2_OutBlock;

//------------------------------------------------------------------------------
// 코스닥 호가 잔량 (HA_)
//------------------------------------------------------------------------------
typedef struct {
    char    hotime              [   6];    char    _hotime              ;    // [string,    6] 호가시간                        StartPos 0, Length 6
    char    offerho1            [   7];    char    _offerho1            ;    // [long  ,    7] 매도호가1                       StartPos 7, Length 7
    char    bidho1              [   7];    char    _bidho1              ;    // [long  ,    7] 매수호가1                       StartPos 15, Length 7
    char    offerrem1           [   9];    char    _offerrem1           ;    // [long  ,    9] 매도호가잔량1                   StartPos 23, Length 9
    char    bidrem1             [   9];    char    _bidrem1             ;    // [long  ,    9] 매수호가잔량1                   StartPos 33, Length 9
    char    offerho2            [   7];    char    _offerho2            ;    // [long  ,    7] 매도호가2                       StartPos 43, Length 7
    char    bidho2              [   7];    char    _bidho2              ;    // [long  ,    7] 매수호가2                       StartPos 51, Length 7
    char    offerrem2           [   9];    char    _offerrem2           ;    // [long  ,    9] 매도호가잔량2                   StartPos 59, Length 9
    char    bidrem2             [   9];    char    _bidrem2             ;    // [long  ,    9] 매수호가잔량2                   StartPos 69, Length 9
    char    offerho3            [   7];    char    _offerho3            ;    // [long  ,    7] 매도호가3                       StartPos 79, Length 7
    char    bidho3              [   7];    char    _bidho3              ;    // [long  ,    7] 매수호가3                       StartPos 87, Length 7
    char    offerrem3           [   9];    char    _offerrem3           ;    // [long  ,    9] 매도호가잔량3                   StartPos 95, Length 9
    char    bidrem3             [   9];    char    _bidrem3             ;    // [long  ,    9] 매수호가잔량3                   StartPos 105, Length 9
    char    offerho4            [   7];    char    _offerho4            ;    // [long  ,    7] 매도호가4                       StartPos 115, Length 7
    char    bidho4              [   7];    char    _bidho4              ;    // [long  ,    7] 매수호가4                       StartPos 123, Length 7
    char    offerrem4           [   9];    char    _offerrem4           ;    // [long  ,    9] 매도호가잔량4                   StartPos 131, Length 9
    char    bidrem4             [   9];    char    _bidrem4             ;    // [long  ,    9] 매수호가잔량4                   StartPos 141, Length 9
    char    offerho5            [   7];    char    _offerho5            ;    // [long  ,    7] 매도호가5                       StartPos 151, Length 7
    char    bidho5              [   7];    char    _bidho5              ;    // [long  ,    7] 매수호가5                       StartPos 159, Length 7
    char    offerrem5           [   9];    char    _offerrem5           ;    // [long  ,    9] 매도호가잔량5                   StartPos 167, Length 9
    char    bidrem5             [   9];    char    _bidrem5             ;    // [long  ,    9] 매수호가잔량5                   StartPos 177, Length 9
    char    offerho6            [   7];    char    _offerho6            ;    // [long  ,    7] 매도호가6                       StartPos 187, Length 7
    char    bidho6              [   7];    char    _bidho6              ;    // [long  ,    7] 매수호가6                       StartPos 195, Length 7
    char    offerrem6           [   9];    char    _offerrem6           ;    // [long  ,    9] 매도호가잔량6                   StartPos 203, Length 9
    char    bidrem6             [   9];    char    _bidrem6             ;    // [long  ,    9] 매수호가잔량6                   StartPos 213, Length 9
    char    offerho7            [   7];    char    _offerho7            ;    // [long  ,    7] 매도호가7                       StartPos 223, Length 7
    char    bidho7              [   7];    char    _bidho7              ;    // [long  ,    7] 매수호가7                       StartPos 231, Length 7
    char    offerrem7           [   9];    char    _offerrem7           ;    // [long  ,    9] 매도호가잔량7                   StartPos 239, Length 9
    char    bidrem7             [   9];    char    _bidrem7             ;    // [long  ,    9] 매수호가잔량7                   StartPos 249, Length 9
    char    offerho8            [   7];    char    _offerho8            ;    // [long  ,    7] 매도호가8                       StartPos 259, Length 7
    char    bidho8              [   7];    char    _bidho8              ;    // [long  ,    7] 매수호가8                       StartPos 267, Length 7
    char    offerrem8           [   9];    char    _offerrem8           ;    // [long  ,    9] 매도호가잔량8                   StartPos 275, Length 9
    char    bidrem8             [   9];    char    _bidrem8             ;    // [long  ,    9] 매수호가잔량8                   StartPos 285, Length 9
    char    offerho9            [   7];    char    _offerho9            ;    // [long  ,    7] 매도호가9                       StartPos 295, Length 7
    char    bidho9              [   7];    char    _bidho9              ;    // [long  ,    7] 매수호가9                       StartPos 303, Length 7
    char    offerrem9           [   9];    char    _offerrem9           ;    // [long  ,    9] 매도호가잔량9                   StartPos 311, Length 9
    char    bidrem9             [   9];    char    _bidrem9             ;    // [long  ,    9] 매수호가잔량9                   StartPos 321, Length 9
    char    offerho10           [   7];    char    _offerho10           ;    // [long  ,    7] 매도호가10                      StartPos 331, Length 7
    char    bidho10             [   7];    char    _bidho10             ;    // [long  ,    7] 매수호가10                      StartPos 339, Length 7
    char    offerrem10          [   9];    char    _offerrem10          ;    // [long  ,    9] 매도호가잔량10                  StartPos 347, Length 9
    char    bidrem10            [   9];    char    _bidrem10            ;    // [long  ,    9] 매수호가잔량10                  StartPos 357, Length 9
    char    totofferrem         [   9];    char    _totofferrem         ;    // [long  ,    9] 총매도호가잔량                  StartPos 367, Length 9
    char    totbidrem           [   9];    char    _totbidrem           ;    // [long  ,    9] 총매수호가잔량                  StartPos 377, Length 9
    char    donsigubun          [   1];    char    _donsigubun          ;    // [string,    1] 동시호가구분                    StartPos 387, Length 1
    char    shcode              [   6];    char    _shcode              ;    // [string,    6] 단축코드                        StartPos 389, Length 6
    char    alloc_gubun         [   1];    char    _alloc_gubun         ;    // [string,    1] 배분적용구분                    StartPos 396, Length 1
    char    volume              [  12];    char    _volume              ;    // [long  ,   12] 누적거래량                      StartPos 398, Length 12
} HA_OutBlock;

//------------------------------------------------------------------------------
// 코스닥 시간외 호가 잔량 (HB_)
//------------------------------------------------------------------------------
typedef struct {
    char    hotime              [   6];    char    _hotime              ;    // [string,    6] 호가시간                        StartPos 0, Length 6
    char    tmofferrem          [  12];    char    _tmofferrem          ;    // [long  ,   12] 시간외매도잔량                  StartPos 7, Length 12
    char    tmbidrem            [  12];    char    _tmbidrem            ;    // [long  ,   12] 시간외매수잔량                  StartPos 20, Length 12
    char    pretmoffercha       [  12];    char    _pretmoffercha       ;    // [long  ,   12] 시간외매도수량직전대비          StartPos 33, Length 12
    char    pretmbidcha         [  12];    char    _pretmbidcha         ;    // [long  ,   12] 시간외매수수량직전대비          StartPos 46, Length 12
    char    shcode              [   6];    char    _shcode              ;    // [string,    6] 단축코드                        StartPos 59, Length 6
} HB_OutBlock;

//------------------------------------------------------------------------------
// 코스피 체결 (S3_)
//------------------------------------------------------------------------------
typedef struct {
    char    chetime             [   6];    char    _chetime             ;    // [string,    6] 체결시간                        StartPos 0, Length 6
    char    sign                [   1];    char    _sign                ;    // [string,    1] 전일대비구분                    StartPos 7, Length 1
    char    change              [   8];    char    _change              ;    // [long  ,    8] 전일대비                        StartPos 9, Length 8
    char    drate               [   6];    char    _drate               ;    // [float ,  6.2] 등락율                          StartPos 18, Length 6
    char    price               [   8];    char    _price               ;    // [long  ,    8] 현재가                          StartPos 25, Length 8
    char    opentime            [   6];    char    _opentime            ;    // [string,    6] 시가시간                        StartPos 34, Length 6
    char    open                [   8];    char    _open                ;    // [long  ,    8] 시가                            StartPos 41, Length 8
    char    hightime            [   6];    char    _hightime            ;    // [string,    6] 고가시간                        StartPos 50, Length 6
    char    high                [   8];    char    _high                ;    // [long  ,    8] 고가                            StartPos 57, Length 8
    char    lowtime             [   6];    char    _lowtime             ;    // [string,    6] 저가시간                        StartPos 66, Length 6
    char    low                 [   8];    char    _low                 ;    // [long  ,    8] 저가                            StartPos 73, Length 8
    char    cgubun              [   1];    char    _cgubun              ;    // [string,    1] 체결구분                        StartPos 82, Length 1
    char    cvolume             [   8];    char    _cvolume             ;    // [long  ,    8] 체결량                          StartPos 84, Length 8
    char    volume              [  12];    char    _volume              ;    // [long  ,   12] 누적거래량                      StartPos 93, Length 12
    char    value               [  12];    char    _value               ;    // [long  ,   12] 누적거래대금                    StartPos 106, Length 12
    char    mdvolume            [  12];    char    _mdvolume            ;    // [long  ,   12] 매도누적체결량                  StartPos 119, Length 12
    char    mdchecnt            [   8];    char    _mdchecnt            ;    // [long  ,    8] 매도누적체결건수                StartPos 132, Length 8
    char    msvolume            [  12];    char    _msvolume            ;    // [long  ,   12] 매수누적체결량                  StartPos 141, Length 12
    char    mschecnt            [   8];    char    _mschecnt            ;    // [long  ,    8] 매수누적체결건수                StartPos 154, Length 8
    char    cpower              [   9];    char    _cpower              ;    // [float ,  9.2] 체결강도                        StartPos 163, Length 9
    char    wAvrg              [   8];    char    _w_avrg              ;    // [long  ,    8] 가중평균가                      StartPos 173, Length 8
    char    offerho             [   8];    char    _offerho             ;    // [long  ,    8] 매도호가                        StartPos 182, Length 8
    char    bidho               [   8];    char    _bidho               ;    // [long  ,    8] 매수호가                        StartPos 191, Length 8
    char    status              [   2];    char    _status              ;    // [string,    2] 장정보                          StartPos 200, Length 2
    char    jnilvolume          [  12];    char    _jnilvolume          ;    // [long  ,   12] 전일동시간대거래량              StartPos 203, Length 12
    char    shcode              [   6];    char    _shcode              ;    // [string,    6] 단축코드                        StartPos 216, Length 6
} S3_OutBlock;

//------------------------------------------------------------------------------
// 코스피 예상 체결 (YS3)
//------------------------------------------------------------------------------
typedef struct {
    char    hotime              [   6];    char    _hotime              ;    // [string,    6] 호가시간                        StartPos 0, Length 6
    char    yeprice             [   8];    char    _yeprice             ;    // [long  ,    8] 예상체결가격                    StartPos 7, Length 8
    char    yevolume            [  12];    char    _yevolume            ;    // [long  ,   12] 예상체결수량                    StartPos 16, Length 12
    char    jnilysign           [   1];    char    _jnilysign           ;    // [string,    1] 예상체결가전일종가대비구분      StartPos 29, Length 1
    char    preychange          [   8];    char    _preychange          ;    // [long  ,    8] 예상체결가전일종가대비          StartPos 31, Length 8
    char    jnilydrate          [   6];    char    _jnilydrate          ;    // [float ,  6.2] 예상체결가전일종가등락율        StartPos 40, Length 6
    char    yofferho0           [   8];    char    _yofferho0           ;    // [long  ,    8] 예상매도호가                    StartPos 47, Length 8
    char    ybidho0             [   8];    char    _ybidho0             ;    // [long  ,    8] 예상매수호가                    StartPos 56, Length 8
    char    yofferrem0          [  12];    char    _yofferrem0          ;    // [long  ,   12] 예상매도호가수량                StartPos 65, Length 12
    char    ybidrem0            [  12];    char    _ybidrem0            ;    // [long  ,   12] 예상매수호가수량                StartPos 78, Length 12
    char    shcode              [   6];    char    _shcode              ;    // [string,    6] 단축코드                        StartPos 91, Length 6
} YS3OutBlock;

//------------------------------------------------------------------------------
// 코스닥 체결 (K3_)
//------------------------------------------------------------------------------
typedef struct {
    char    chetime             [   6];    char    _chetime             ;    // [string,    6] 체결시간                        StartPos 0, Length 6
    char    sign                [   1];    char    _sign                ;    // [string,    1] 전일대비구분                    StartPos 7, Length 1
    char    change              [   8];    char    _change              ;    // [long  ,    8] 전일대비                        StartPos 9, Length 8
    char    drate               [   6];    char    _drate               ;    // [float ,  6.2] 등락율                          StartPos 18, Length 6
    char    price               [   8];    char    _price               ;    // [long  ,    8] 현재가                          StartPos 25, Length 8
    char    opentime            [   6];    char    _opentime            ;    // [string,    6] 시가시간                        StartPos 34, Length 6
    char    open                [   8];    char    _open                ;    // [long  ,    8] 시가                            StartPos 41, Length 8
    char    hightime            [   6];    char    _hightime            ;    // [string,    6] 고가시간                        StartPos 50, Length 6
    char    high                [   8];    char    _high                ;    // [long  ,    8] 고가                            StartPos 57, Length 8
    char    lowtime             [   6];    char    _lowtime             ;    // [string,    6] 저가시간                        StartPos 66, Length 6
    char    low                 [   8];    char    _low                 ;    // [long  ,    8] 저가                            StartPos 73, Length 8
    char    cgubun              [   1];    char    _cgubun              ;    // [string,    1] 체결구분                        StartPos 82, Length 1
    char    cvolume             [   8];    char    _cvolume             ;    // [long  ,    8] 체결량                          StartPos 84, Length 8
    char    volume              [  12];    char    _volume              ;    // [long  ,   12] 누적거래량                      StartPos 93, Length 12
    char    value               [  12];    char    _value               ;    // [long  ,   12] 누적거래대금                    StartPos 106, Length 12
    char    mdvolume            [  12];    char    _mdvolume            ;    // [long  ,   12] 매도누적체결량                  StartPos 119, Length 12
    char    mdchecnt            [   8];    char    _mdchecnt            ;    // [long  ,    8] 매도누적체결건수                StartPos 132, Length 8
    char    msvolume            [  12];    char    _msvolume            ;    // [long  ,   12] 매수누적체결량                  StartPos 141, Length 12
    char    mschecnt            [   8];    char    _mschecnt            ;    // [long  ,    8] 매수누적체결건수                StartPos 154, Length 8
    char    cpower              [   9];    char    _cpower              ;    // [float ,  9.2] 체결강도                        StartPos 163, Length 9
    char    wAvrg              [   8];    char    _w_avrg              ;    // [long  ,    8] 가중평균가                      StartPos 173, Length 8
    char    offerho             [   8];    char    _offerho             ;    // [long  ,    8] 매도호가                        StartPos 182, Length 8
    char    bidho               [   8];    char    _bidho               ;    // [long  ,    8] 매수호가                        StartPos 191, Length 8
    char    status              [   2];    char    _status              ;    // [string,    2] 장정보                          StartPos 200, Length 2
    char    jnilvolume          [  12];    char    _jnilvolume          ;    // [long  ,   12] 전일동시간대거래량              StartPos 203, Length 12
    char    shcode              [   6];    char    _shcode              ;    // [string,    6] 단축코드                        StartPos 216, Length 6
} K3_OutBlock;

//------------------------------------------------------------------------------
// 코스닥 예상 체결 (YK3)
//------------------------------------------------------------------------------
typedef struct {
    char    hotime              [   6];    char    _hotime              ;    // [string,    6] 호가시간                        StartPos 0, Length 6
    char    yeprice             [   8];    char    _yeprice             ;    // [long  ,    8] 예상체결가격                    StartPos 7, Length 8
    char    yevolume            [  12];    char    _yevolume            ;    // [long  ,   12] 예상체결수량                    StartPos 16, Length 12
    char    jnilysign           [   1];    char    _jnilysign           ;    // [string,    1] 예상체결가전일종가대비구분      StartPos 29, Length 1
    char    preychange          [   8];    char    _preychange          ;    // [long  ,    8] 예상체결가전일종가대비          StartPos 31, Length 8
    char    jnilydrate          [   6];    char    _jnilydrate          ;    // [float ,  6.2] 예상체결가전일종가등락율        StartPos 40, Length 6
    char    yofferho0           [   8];    char    _yofferho0           ;    // [long  ,    8] 예상매도호가                    StartPos 47, Length 8
    char    ybidho0             [   8];    char    _ybidho0             ;    // [long  ,    8] 예상매수호가                    StartPos 56, Length 8
    char    yofferrem0          [  12];    char    _yofferrem0          ;    // [long  ,   12] 예상매도호가수량                StartPos 65, Length 12
    char    ybidrem0            [  12];    char    _ybidrem0            ;    // [long  ,   12] 예상매수호가수량                StartPos 78, Length 12
    char    shcode              [   6];    char    _shcode              ;    // [string,    6] 단축코드                        StartPos 91, Length 6
} YK3OutBlock;

//------------------------------------------------------------------------------
// 코스피 ETF종목 실시간 NAV (I5_)
//------------------------------------------------------------------------------
typedef struct {
    char    time                [   8];    char    _time                ;    // [string,    8] 시간                            StartPos 0, Length 8
    char    price               [   8];    char    _price               ;    // [long  ,    8] 현재가                          StartPos 9, Length 8
    char    sign                [   1];    char    _sign                ;    // [string,    1] 전일대비구분                    StartPos 18, Length 1
    char    change              [   8];    char    _change              ;    // [long  ,    8] 전일대비                        StartPos 20, Length 8
    char    volume              [  12];    char    _volume              ;    // [float ,   12] 누적거래량                      StartPos 29, Length 12
    char    navdiff             [   9];    char    _navdiff             ;    // [float ,  9.2] NAV대비                         StartPos 42, Length 9
    char    nav                 [   9];    char    _nav                 ;    // [float ,  9.2] NAV                             StartPos 52, Length 9
    char    navchange           [   9];    char    _navchange           ;    // [float ,  9.2] 전일대비                        StartPos 62, Length 9
    char    crate               [   9];    char    _crate               ;    // [float ,  9.2] 추적오차                        StartPos 72, Length 9
    char    grate               [   9];    char    _grate               ;    // [float ,  9.2] 괴리                            StartPos 82, Length 9
    char    jisu                [   8];    char    _jisu                ;    // [float ,  8.2] 지수                            StartPos 92, Length 8
    char    jichange            [   8];    char    _jichange            ;    // [float ,  8.2] 전일대비                        StartPos 101, Length 8
    char    jirate              [   8];    char    _jirate              ;    // [float ,  8.2] 전일대비율                      StartPos 110, Length 8
    char    shcode              [   6];    char    _shcode              ;    // [string,    6] 단축코드                        StartPos 119, Length 6
} I5_OutBlock;

//------------------------------------------------------------------------------
// 현물 VI발동해제 (VI_)
//------------------------------------------------------------------------------
typedef struct {
    char    vi_gubun            [   1];    char    _vi_gubun            ;    // [string,    1] 구분(0:해제 1:정적발동 2:동적발 StartPos 0, Length 1
    char    svi_recprice        [   8];    char    _svi_recprice        ;    // [long  ,    8] 정적VI발동기준가격              StartPos 2, Length 8
    char    dvi_recprice        [   8];    char    _dvi_recprice        ;    // [long  ,    8] 동적VI발동기준가격              StartPos 11, Length 8
    char    vi_trgprice         [   8];    char    _vi_trgprice         ;    // [long  ,    8] VI발동가격                      StartPos 20, Length 8
    char    shcode              [   6];    char    _shcode              ;    // [string,    6] 단축코드(KEY)                   StartPos 29, Length 6
    char    ref_shcode          [   6];    char    _ref_shcode          ;    // [string,    6] 참조코드                        StartPos 36, Length 6
    char    time                [   6];    char    _time                ;    // [string,    6] 시간                            StartPos 43, Length 6
} VI_OutBlock;

//------------------------------------------------------------------------------
// 시간외 단일가 VI발동해제 (DVI)
//------------------------------------------------------------------------------
typedef struct {
    char    vi_gubun            [   1];    char    _vi_gubun            ;    // [string,    1] 구분(0:해제 1:정적발동 2:동적발 StartPos 0, Length 1
    char    svi_recprice        [   8];    char    _svi_recprice        ;    // [long  ,    8] 정적VI발동기준가격              StartPos 2, Length 8
    char    dvi_recprice        [   8];    char    _dvi_recprice        ;    // [long  ,    8] 동적VI발동기준가격              StartPos 11, Length 8
    char    vi_trgprice         [   8];    char    _vi_trgprice         ;    // [long  ,    8] VI발동가격                      StartPos 20, Length 8
    char    shcode              [   6];    char    _shcode              ;    // [string,    6] 단축코드(KEY)                   StartPos 29, Length 6
    char    ref_shcode          [   6];    char    _ref_shcode          ;    // [string,    6] 참조코드(미사용)                StartPos 36, Length 6
    char    time                [   6];    char    _time                ;    // [string,    6] 시간                            StartPos 43, Length 6
} DVIOutBlock;

//------------------------------------------------------------------------------
// 장 운영 정보 (JIF)
//------------------------------------------------------------------------------
typedef struct {
    char    jangubun[1];    //[string,    1] 장구분   StartPos 0, Length 1
} JIFInBlock;

typedef struct {
    char    jangubun[1];    //[string,    1] 장구분   StartPos 0, Length 1
    char    jstatus[2];    //[string,    2] 장상태   StartPos 1, Length 2
} JIFOutBlock;

//------------------------------------------------------------------------------
// 업종별 투자자별 매매현황 (BM_)
//------------------------------------------------------------------------------
typedef struct {
    char    upCode              [   3];    char    _upcode              ;    // [string,    3] 업종코드                        StartPos 0, Length 3
} BM_InBlock;

typedef struct {
    char    tjjCode             [   4];    char    _tjjcode             ;    // [string,    4] 투자자코드                      StartPos 0, Length 4
    char    tjjTime             [   8];    char    _tjjtime             ;    // [string,    8] 수신시간                        StartPos 5, Length 8
    char    msVolume            [   8];    char    _msvolume            ;    // [long  ,    8] 매수 거래량                     StartPos 14, Length 8
    char    mdVolume            [   8];    char    _mdvolume            ;    // [long  ,    8] 매도 거래량                     StartPos 23, Length 8
    char    msVol               [   8];    char    _msvol               ;    // [long  ,    8] 거래량 순매수                   StartPos 32, Length 8
    char    pMsVol             [   8];    char    _p_msvol             ;    // [long  ,    8] 거래량 순매수 직전대비          StartPos 41, Length 8
    char    msValue             [   6];    char    _msvalue             ;    // [long  ,    6] 매수 거래대금                   StartPos 50, Length 6
    char    mdValue             [   6];    char    _mdvalue             ;    // [long  ,    6] 매도 거래대금                   StartPos 57, Length 6
    char    msVal               [   6];    char    _msval               ;    // [long  ,    6] 거래대금 순매수                 StartPos 64, Length 6
    char    pMsVal             [   6];    char    _p_msval             ;    // [long  ,    6] 거래대금 순매수 직전대비        StartPos 71, Length 6
    char    upCode              [   3];    char    _upcode              ;    // [string,    3] 업종코드                        StartPos 78, Length 3
} BM_OutBlock;

// 앞서 '#pragma pack(push, 1)'로 1바이트 단위로 설정한
// 구조체 메모리 저장방식을 원래대로 되돌림.
#pragma pack(pop)

