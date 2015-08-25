/************************************************************************************
	주의

	이 자료는 2013년 10월 15일 기준 자료이며 향후 변경될 가능성이 있습니다.
	자료 구조가 맞지 않을 경우 구조체가 변경되지 않았는지 확인하시기 바랍니다.

	최신 자료는 웹페이지를 통해 안내되며 자동 안내(OpenAPI Login시)를 하고 있으니
	게시를 꼭 확인하시기 바랍니다.

************************************************************************************/

////////////////////////////////////////
//	주식 ELW
////////////////////////////////////////

typedef struct tagc8101InBlock    //기본입력
{
	char pswd_noz8                        [ 44];	char _pswd_noz8;                          //비밀번호
	char issue_codez6                     [  6];	char _issue_codez6;                       //종목번호
	char order_qtyz12                     [ 12];	char _order_qtyz12;                       //주문수량
	char order_unit_pricez10              [ 10];	char _order_unit_pricez10;                //주문단가
	char trade_typez2                     [  2];	char _trade_typez2;                       //매매유형
	char shsll_pos_flagz1                 [  1];	char _shsll_pos_flagz1;                   //공매도가능여부
	char trad_pswd_no_1z8                 [ 44];	char _trad_pswd_no_1z8;                   //거래비밀번호1
	char trad_pswd_no_2z8                 [ 44];	char _trad_pswd_no_2z8;                   //거래비밀번호2
} Tc8101InBlock;

typedef struct tagc8101OutBlock    //화면출력
{
	char order_noz10                      [ 10];	char _order_noz10;                        //주문번호
	char order_qtyz12                     [ 12];	char _order_qtyz12;                       //주문수량
	char order_unit_pricez10              [ 10];	char _order_unit_pricez10;                //주문단가
} Tc8101OutBlock;

typedef struct tagc8101
{
	Tc8101InBlock                     c8101inblock                          ;  //기본입력
	Tc8101OutBlock                    c8101outblock                         ;  //화면출력
} Tc8101;

typedef struct tagc8102InBlock    //기본입력
{
	char pswd_noz8                        [ 44];	char _pswd_noz8;                          //비밀번호
	char issue_codez6                     [  6];	char _issue_codez6;                       //종목번호
	char order_qtyz12                     [ 12];	char _order_qtyz12;                       //주문수량
	char order_unit_pricez10              [ 10];	char _order_unit_pricez10;                //주문단가
	char trade_typez2                     [  2];	char _trade_typez2;                       //매매유형
	char trad_pswd_no_1z8                 [ 44];	char _trad_pswd_no_1z8;                   //거래비밀번호1
	char trad_pswd_no_2z8                 [ 44];	char _trad_pswd_no_2z8;                   //거래비밀번호2
} Tc8102InBlock;

typedef struct tagc8102OutBlock    //화면출력
{
	char order_noz10                      [ 10];	char _order_noz10;                        //주문번호
	char order_qtyz12                     [ 12];	char _order_qtyz12;                       //주문수량
	char order_unit_pricez10              [ 10];	char _order_unit_pricez10;                //주문단가
} Tc8102OutBlock;

typedef struct tagc8102
{
	Tc8102InBlock                     c8102inblock                          ;  //기본입력
	Tc8102OutBlock                    c8102outblock                         ;  //화면출력
} Tc8102;

typedef struct tagc8103InBlock    //기본입력
{
	char pswd_noz8                        [ 44];	char _pswd_noz8;                          //비밀번호
	char issue_codez6                     [  6];	char _issue_codez6;                       //종목번호
	char crctn_qtyz12                     [ 12];	char _crctn_qtyz12;                       //정정수량
	char crctn_pricez10                   [ 10];	char _crctn_pricez10;                     //정정단가
	char orgnl_order_noz10                [ 10];	char _orgnl_order_noz10;                  //원주문번호
	char all_part_typez1                  [  1];	char _all_part_typez1;                    //정정구분
	char trad_pswd_no_1z8                 [ 44];	char _trad_pswd_no_1z8;                   //거래비밀1
	char trad_pswd_no_2z8                 [ 44];	char _trad_pswd_no_2z8;                   //거래비밀2
} Tc8103InBlock;

typedef struct tagc8103OutBlock    //화면출력
{
	char orgnl_order_noz10                [ 10];	char _orgnl_order_noz10;                  //원주문번호
	char order_noz10                      [ 10];	char _order_noz10;                        //주문번호
	char mom_order_noz10                  [ 10];	char _mom_order_noz10;                    //모주문번호
	char issue_codez6                     [  6];	char _issue_codez6;                       //후종목번호
	char crctn_qtyz12                     [ 12];	char _crctn_qtyz12;                       //정정수량
	char crctn_pricez10                   [ 10];	char _crctn_pricez10;                     //정정단가
} Tc8103OutBlock;

typedef struct tagc8103
{
	Tc8103InBlock                     c8103inblock                          ;  //기본입력
	Tc8103OutBlock                    c8103outblock                         ;  //화면출력
} Tc8103;

typedef struct tagc8104InBlock    //기본입력
{
	char pswd_noz8                        [ 44];	char _pswd_noz8;                          //비밀번호
	char issue_codez6                     [  6];	char _issue_codez6;                       //종목번호
	char canc_qtyz12                      [ 12];	char _canc_qtyz12;                        //취소수량
	char orgnl_order_noz10                [ 10];	char _orgnl_order_noz10;                  //원주문번호
	char all_part_typez1                  [  1];	char _all_part_typez1;                    //취소구분
	char trad_pswd_no_1z8                 [ 44];	char _trad_pswd_no_1z8;                   //거래비밀번호1
	char trad_pswd_no_2z8                 [ 44];	char _trad_pswd_no_2z8;                   //거래비밀번호2
} Tc8104InBlock;

typedef struct tagc8104OutBlock    //화면출력
{
	char orgnl_order_noz10                [ 10];	char _orgnl_order_noz10;                  //원주문번호
	char order_noz10                      [ 10];	char _order_noz10;                        //주문번호
	char mom_order_noz10                  [ 10];	char _mom_order_noz10;                    //모주문번호
	char issue_codez6                     [  6];	char _issue_codez6;                       //후종목번호
	char canc_qtyz12                      [ 12];	char _canc_qtyz12;                        //취소수량
} Tc8104OutBlock;

typedef struct tagc8104
{
	Tc8104InBlock                     c8104inblock                          ;  //기본입력
	Tc8104OutBlock                    c8104outblock                         ;  //화면출력
} Tc8104;

typedef struct tagc8141InBlock    //기본입력
{
	char pswd_noz8                        [ 44];	char _pswd_noz8;                          //비밀번호
	char issue_codez12                    [ 12];	char _issue_codez12;                      //종목번호
	char buy_datez8                       [  8];	char _buy_datez8;                         //매수일자
	char order_qtyz12                     [ 12];	char _order_qtyz12;                       //주문수량
	char order_unit_pricez10              [ 10];	char _order_unit_pricez10;                //주문단가
	char trade_typez2                     [  2];	char _trade_typez2;                       //매매유형
	char order_condz1                     [  1];	char _order_condz1;                       //주문조건
	char shsll_pos_flagz1                 [  1];	char _shsll_pos_flagz1;                   //공매도가능여부
	char trad_pswd_no_1z8                 [ 44];	char _trad_pswd_no_1z8;                   //거래비밀번호1
	char trad_pswd_no_2z8                 [ 44];	char _trad_pswd_no_2z8;                   //거래비밀번호2
} Tc8141InBlock;

typedef struct tagc8141OutBlock    //화면출력
{
	char order_noz10                      [ 10];	char _order_noz10;                        //주문번호
	char order_qtyz12                     [ 12];	char _order_qtyz12;                       //주문수량
	char order_unit_pricez10              [ 10];	char _order_unit_pricez10;                //주문단가
} Tc8141OutBlock;

typedef struct tagc8141
{
	Tc8141InBlock                     c8141inblock                          ;  //기본입력
	Tc8141OutBlock                    c8141outblock                         ;  //화면출력
} Tc8141;

typedef struct tagc8142InBlock    //기본입력
{
	char password_noz8                    [ 44];	char _password_noz8;                      //비밀번호
	char issue_codez12                    [ 12];	char _issue_codez12;                      //종목번호
	char order_qtyz12                     [ 12];	char _order_qtyz12;                       //주문수량
	char order_unit_pricez10              [ 10];	char _order_unit_pricez10;                //주문단가
	char trade_typez2                     [  2];	char _trade_typez2;                       //매매유형
	char order_condz1                     [  1];	char _order_condz1;                       //주문조건
	char trad_pswd_no_1z8                 [ 44];	char _trad_pswd_no_1z8;                   //거래비밀번호1
	char trad_pswd_no_2z8                 [ 44];	char _trad_pswd_no_2z8;                   //거래비밀번호2
} Tc8142InBlock;

typedef struct tagc8142OutBlock    //화면출력
{
	char order_noz10                      [ 10];	char _order_noz10;                        //주문번호
	char order_qtyz12                     [ 12];	char _order_qtyz12;                       //주문수량
	char order_unit_pricez10              [ 10];	char _order_unit_pricez10;                //주문단가
} Tc8142OutBlock;

typedef struct tagc8142
{
	Tc8142InBlock                     c8142inblock                          ;  //기본입력
	Tc8142OutBlock                    c8142outblock                         ;  //화면출력
} Tc8142;

typedef struct tagc8143InBlock    //기본입력
{
	char pswd_noz8                        [ 44];	char _pswd_noz8;                          //비밀번호
	char issue_codez12                    [ 12];	char _issue_codez12;                      //종목번호
	char crctn_qtyz12                     [ 12];	char _crctn_qtyz12;                       //정정수량
	char crctn_pricez10                   [ 10];	char _crctn_pricez10;                     //정정단가
	char orgnl_order_noz10                [ 10];	char _orgnl_order_noz10;                  //원주문번호
	char all_part_typez1                  [  1];	char _all_part_typez1;                    //정정구분
	char order_condz1                     [  1];	char _order_condz1;                       //주문조건
	char trad_pswd_no_1z8                 [ 44];	char _trad_pswd_no_1z8;                   //거래비밀번호1
	char trad_pswd_no_2z8                 [ 44];	char _trad_pswd_no_2z8;                   //거래비밀번호2
} Tc8143InBlock;

typedef struct tagc8143OutBlock    //화면출력
{
	char orgnl_order_noz10                [ 10];	char _orgnl_order_noz10;                  //원주문번호
	char order_noz10                      [ 10];	char _order_noz10;                        //주문번호
	char mom_order_noz10                  [ 10];	char _mom_order_noz10;                    //모주문번호
	char af_issue_codez12                 [ 12];	char _af_issue_codez12;                   //후종목번호
	char crctn_qtyz12                     [ 12];	char _crctn_qtyz12;                       //정정수량
	char crctn_pricez10                   [ 10];	char _crctn_pricez10;                     //정정단가
} Tc8143OutBlock;

typedef struct tagc8143
{
	Tc8143InBlock                     c8143inblock                          ;  //기본입력
	Tc8143OutBlock                    c8143outblock                         ;  //화면출력
} Tc8143;

typedef struct tagc8201InBlock    //기본입력
{
	char pswd_noz44                       [ 44];	char _pswd_noz44;                         //비밀번호
	char bnc_bse_cdz1                     [  1];	char _bnc_bse_cdz1;                       //잔고구분
} Tc8201InBlock;

typedef struct tagc8201OutBlock    //화면출력
{
	char dpsit_amtz16                     [ 16];	char _dpsit_amtz16;                       //예수금
	char mrgn_amtz16                      [ 16];	char _mrgn_amtz16;                        //신용융자금
	char mgint_npaid_amtz16               [ 16];	char _mgint_npaid_amtz16;                 //이자미납금
	char chgm_pos_amtz16                  [ 16];	char _chgm_pos_amtz16;                    //출금가능금액
	char cash_mrgn_amtz16                 [ 16];	char _cash_mrgn_amtz16;                   //현금증거금
	char subst_mgamt_amtz16               [ 16];	char _subst_mgamt_amtz16;                 //대용증거금
	char coltr_ratez6                     [  6];	char _coltr_ratez6;                       //담보비율
	char rcble_amtz16                     [ 16];	char _rcble_amtz16;                       //현금미수금
	char order_pos_csamtz16               [ 16];	char _order_pos_csamtz16;                 //주문가능액
	char ecn_pos_csamtz16                 [ 16];	char _ecn_pos_csamtz16;                   //ECN주문가능액
	char nordm_loan_amtz16                [ 16];	char _nordm_loan_amtz16;                  //미상환금
	char etc_lend_amtz16                  [ 16];	char _etc_lend_amtz16;                    //기타대여금
	char subst_amtz16                     [ 16];	char _subst_amtz16;                       //대용금액
	char sln_sale_amtz16                  [ 16];	char _sln_sale_amtz16;                    //대주담보금
	char bal_buy_ttamtz16                 [ 16];	char _bal_buy_ttamtz16;                   //매입원가(계좌합산)
	char bal_ass_ttamtz16                 [ 16];	char _bal_ass_ttamtz16;                   //평가금액(계좌합산)
	char asset_tot_amtz16                 [ 16];	char _asset_tot_amtz16;                   //순자산액(계좌합산)
	char actvt_type10                     [ 10];	char _actvt_type10;                       //활동유형
	char lend_amtz16                      [ 16];	char _lend_amtz16;                        //대출금
	char accnt_mgamt_ratez6               [  6];	char _accnt_mgamt_ratez6;                 //계좌증거금율
	char sl_mrgn_amtz16                   [ 16];	char _sl_mrgn_amtz16;                     //매도증거금
	char pos_csamt1z16                    [ 16];	char _pos_csamt1z16;                      //20%주문가능금액
	char pos_csamt2z16                    [ 16];	char _pos_csamt2z16;                      //30%주문가능금액
	char pos_csamt3z16                    [ 16];	char _pos_csamt3z16;                      //40%주문가능금액
	char pos_csamt4z16                    [ 16];	char _pos_csamt4z16;                      //100%주문가능금액
	char dpsit_amtz_d1_16                 [ 16];	char _dpsit_amtz_d1_16;                   //D1예수금
	char dpsit_amtz_d2_16                 [ 16];	char _dpsit_amtz_d2_16;                   //D2예수금
	char noticez30                        [ 30];	char _noticez30;                          //공지사항             /*To-be에없음*/
	char tot_eal_plsz18                   [ 18];	char _tot_eal_plsz18;                     //총평가손익
	char pft_rtz15                        [ 15];	char _pft_rtz15;                          //수익율
} Tc8201OutBlock;

typedef struct tagc8201OutBlock1    //화면출력, [반복]
{
	char issue_codez6                     [  6];	char _issue_codez6;                       //종목번호
	char issue_namez40                    [ 40];	char _issue_namez40;                      //종목명
	char bal_typez6                       [  6];	char _bal_typez6;                         //잔고유형
	char loan_datez10                     [ 10];	char _loan_datez10;                       //대출일
	char bal_qtyz16                       [ 16];	char _bal_qtyz16;                         //잔고수량
	char unstl_qtyz16                     [ 16];	char _unstl_qtyz16;                       //미결제량
	char slby_amtz16                      [ 16];	char _slby_amtz16;                        //평균매입가
	char prsnt_pricez16                   [ 16];	char _prsnt_pricez16;                     //현재가
	char lsnpf_amtz16                     [ 16];	char _lsnpf_amtz16;                       //손익(천원)
	char earn_ratez9                      [  9];	char _earn_ratez9;                        //손익율
	char mrgn_codez4                      [  4];	char _mrgn_codez4;                        //신용유형
	char jan_qtyz16                       [ 16];	char _jan_qtyz16;                         //잔량
	char expr_datez10                     [ 10];	char _expr_datez10;                       //만기일
	char ass_amtz16                       [ 16];	char _ass_amtz16;                         //평가금액
	char issue_mgamt_ratez6               [  6];	char _issue_mgamt_ratez6;                 //종목증거금율         /*float->char*/
	char medo_slby_amtz16                 [ 16];	char _medo_slby_amtz16;                   //평균매도가
	char post_lsnpf_amtz16                [ 16];	char _post_lsnpf_amtz16;                  //매도손익
} Tc8201OutBlock1;

typedef struct tagc8201
{
	Tc8201InBlock                     c8201inblock                          ;  //기본입력
	Tc8201OutBlock                    c8201outblock                         ;  //화면출력
	Tc8201OutBlock1                   c8201outblock1                   [ 20];  //화면출력 , [반복]
} Tc8201;

typedef struct tags8120InBlock    //기본입력
{
	char inq_gubunz1                      [  1];	char _inq_gubunz1;                        //조회주체구분
	char pswd_noz8                        [ 44];	char _pswd_noz8;                          //비밀번호
	char group_noz4                       [  4];	char _group_noz4;                         //그룹번호
	char mkt_slctz1                       [  1];	char _mkt_slctz1;                         //시장구분
	char order_datez8                     [  8];	char _order_datez8;                       //주문일자
	char issue_codez12                    [ 12];	char _issue_codez12;                      //종목번호
	char comm_order_typez2                [  2];	char _comm_order_typez2;                  //매체구분
	char conc_gubunz1                     [  1];	char _conc_gubunz1;                       //체결구분
	char inq_seq_gubunz1                  [  1];	char _inq_seq_gubunz1;                    //조회순서
	char sort_gubunz1                     [  1];	char _sort_gubunz1;                       //정렬구분
	char sell_buy_typez1                  [  1];	char _sell_buy_typez1;                    //매수도구분
	char mrgn_typez1                      [  1];	char _mrgn_typez1;                        //신용구분
	char accnt_admin_typez1               [  1];	char _accnt_admin_typez1;                 //계좌구분
	char order_noz10                      [ 10];	char _order_noz10;                        //주문번호
	char ctsz56                           [ 56];	char _ctsz56;                             //CTS                  
	char trad_pswd1z8                     [ 44];	char _trad_pswd1z8;                       //거래비밀번호1
	char trad_pswd2z8                     [ 44];	char _trad_pswd2z8;                       //거래비밀번호2
	char IsPageUp                         [  1];	char _IsPageUp;                           //ISPAGEUP             
} Ts8120InBlock;

typedef struct tags8120OutBlock    //화면출력
{
	char emp_kor_namez20                  [ 20];	char _emp_kor_namez20;                    //한글사원성명
	char brch_namez30                     [ 30];	char _brch_namez30;                       //한글지점명
	char buy_conc_qtyz14                  [ 14];	char _buy_conc_qtyz14;                    //매수체결수량
	char buy_conc_amtz19                  [ 19];	char _buy_conc_amtz19;                    //매수체결금액
	char sell_conc_qtyz14                 [ 14];	char _sell_conc_qtyz14;                   //매도체결수량
	char sell_conc_amtz19                 [ 19];	char _sell_conc_amtz19;                   //매도체결금액
} Ts8120OutBlock;

typedef struct tags8120OutBlock1    //화면출력1, [반복]
{
	char order_datez8                     [  8];	char _order_datez8;                       //주문일자
	char order_noz10                      [ 10];	char _order_noz10;                        //주문번호
	char orgnl_order_noz10                [ 10];	char _orgnl_order_noz10;                  //원주문번호
	char accnt_noz11                      [ 11];	char _accnt_noz11;                        //계좌번호
	char accnt_namez20                    [ 20];	char _accnt_namez20;                      //계좌명
	char order_kindz20                    [ 20];	char _order_kindz20;                      //주문구분
	char trd_gubun_noz1                   [  1];	char _trd_gubun_noz1;                     //매매구분번호
	char trd_gubunz20                     [ 20];	char _trd_gubunz20;                       //매매구분
	char trade_type_noz1                  [  1];	char _trade_type_noz1;                    //거래구분번호
	char trade_type1z20                   [ 20];	char _trade_type1z20;                     //거래구분
	char issue_codez12                    [ 12];	char _issue_codez12;                      //종목번호
	char issue_namez40                    [ 40];	char _issue_namez40;                      //종목명
	char order_qtyz10                     [ 10];	char _order_qtyz10;                       //주문수량
	char conc_qtyz10                      [ 10];	char _conc_qtyz10;                        //체결수량
	char order_unit_pricez12              [ 12];	char _order_unit_pricez12;                //주문단가
	char conc_unit_pricez12               [ 12];	char _conc_unit_pricez12;                 //체결평균단가
	char crctn_canc_qtyz10                [ 10];	char _crctn_canc_qtyz10;                  //정정취소수량
	char cfirm_qtyz10                     [ 10];	char _cfirm_qtyz10;                       //확인수량
	char media_namez12                    [ 12];	char _media_namez12;                      //매체구분
	char proc_emp_noz5                    [  5];	char _proc_emp_noz5;                      //처리사번
	char proc_timez8                      [  8];	char _proc_timez8;                        //처리시간
	char proc_termz8                      [  8];	char _proc_termz8;                        //처리단말
	char proc_typez12                     [ 12];	char _proc_typez12;                       //처리구분
	char rejec_codez5                     [  5];	char _rejec_codez5;                       //거부코드
	char avail_qtyz10                     [ 10];	char _avail_qtyz10;                       //정취가능수량
	char mkt_typez1                       [  1];	char _mkt_typez1;                         //시장구분
	char shsll_typez20                    [ 20];	char _shsll_typez20;                      //공매도구분
	char passwd_noz8                      [  8];	char _passwd_noz8;                        //비밀번호
} Ts8120OutBlock1;

typedef struct tags8120OutBlock_IN    //Button정보
{
	char ctsz56                           [ 56];	char _ctsz56;                             //CTS                  
	char nextbutton                       [  1];	char _nextbutton;                         //NEXTBUTTON           
} Ts8120OutBlock_IN;

typedef struct tags8120
{
	Ts8120InBlock                     s8120inblock                          ;  //기본입력
	Ts8120OutBlock                    s8120outblock                         ;  //화면출력
	Ts8120OutBlock1                   s8120outblock1                   [ 20];  //화면출력1 , [반복]
	Ts8120OutBlock_IN                 s8120outblock_in                      ;  //Button정보
} Ts8120;

typedef struct tagp8101InBlock    //입력
{
	char pswd_noz8                        [ 44];	char _pswd_noz8;                          //비밀번호
	char gubunz1                          [  1];	char _gubunz1;                            //구분
} Tp8101InBlock;

typedef struct tagp8101OutBlock    //화면출력
{
	char accnt_namez30                    [ 30];	char _accnt_namez30;                      //계좌명               /*신OBM에존재하지않는항목*/
} Tp8101OutBlock;

typedef struct tagp8101OutBlock1    //GRID, [반복]
{
	char gubunz1                          [  1];	char _gubunz1;                            //구분                 1:현금2:융자3:채권4:대주5:ECN현금
	char gubun_namez6                     [  6];	char _gubun_namez6;                       //구분명
	char issue_codez12                    [ 12];	char _issue_codez12;                      //종목코드
	char issue_namez30                    [ 30];	char _issue_namez30;                      //종목명
	char mrgn_typez10                     [ 10];	char _mrgn_typez10;                       //신용구분
	char lend_datez10                     [ 10];	char _lend_datez10;                       //대출일자
	char taxtn_typez10                    [ 10];	char _taxtn_typez10;                      //과세유형
	char bal_qtyz12                       [ 12];	char _bal_qtyz12;                         //잔고수량
	char sell_rcble_qtyz12                [ 12];	char _sell_rcble_qtyz12;                  //매도미결제
	char buy_rcble_qtyz12                 [ 12];	char _buy_rcble_qtyz12;                   //매수미결제
	char sell_psqtyz12                    [ 12];	char _sell_psqtyz12;                      //매도가능수량
	char today_sell_rcble_qz12            [ 12];	char _today_sell_rcble_qz12;              //당일매도미체결수량
	char avrg_purch_uprc                  [ 10];	char _avrg_purch_uprc;                    //매입단가
} Tp8101OutBlock1;

typedef struct tagp8101
{
	Tp8101InBlock                     p8101inblock                          ;  //입력
	Tp8101OutBlock                    p8101outblock                         ;  //화면출력
	Tp8101OutBlock1                   p8101outblock1                   [ 17];  //GRID , [반복]
} Tp8101;

typedef struct tagp8105InBlock    //입력
{
	char pwdz8                            [ 44];	char _pwdz8;                              //비밀번호
	char ost_dit_cdz1                     [  1];	char _ost_dit_cdz1;                       //구분코드             /*1현금2:신용3:매입자금대출*/
	char sby_dit_cdz1                     [  1];	char _sby_dit_cdz1;                       //매매구분코드         /*1:매도상환2:매수신규*/
	char iem_gbz1                         [  1];	char _iem_gbz1;                           //종목구분             /*1:주식2:ELW3:신주인수4:기타*/
	char iem_cdz12                        [ 12];	char _iem_cdz12;                          //종목코드
	char nmn_pr_tp_gbz1                   [  1];	char _nmn_pr_tp_gbz1;                     //호가유형구분         /*1:구호가구분-1자리2:신시스템호가구분-2자리*/
	char nmn_pr_tp_cdz2                   [  2];	char _nmn_pr_tp_cdz2;                     //호가유형코드         /*01:보통05:시장가06:조건부10:S-OPTION자기11:금전신탁12:최유리13:최우선61:장전시간71:장후시간81:신간외단일*/
	char orr_prz18                        [ 18];	char _orr_prz18;                          //주문가격
	char mdi_tp_cdz1                      [  1];	char _mdi_tp_cdz1;                        //매체유형코드         /*1:지점2:HTS3:모바일ARS4:고객지원센터5:TXflat6:TXLever7:TXLever대출신용8:TXWin9:TXWinSMIT0:TX바로*/
	char cfd_lon_cdz2                     [  2];	char _cfd_lon_cdz2;                       //신용대출코드         /*01:유융02:자융03:유대04:자대*/
	char lon_dtz8                         [  8];	char _lon_dtz8;                           //대출일자
} Tp8105InBlock;

typedef struct tagp8105OutBlock    //화면출력
{
	char dcaz18                           [ 18];	char _dcaz18;                             //예수금               /*금일예수금*/
	char nxt_dd_dcaz18                    [ 18];	char _nxt_dd_dcaz18;                      //익일예수금           /*D+1예수금*/
	char nxt2_dd_dcaz18                   [ 18];	char _nxt2_dd_dcaz18;                     //익익일예수금         /*D+2예수금*/
	char max_pbl_amtz18                   [ 18];	char _max_pbl_amtz18;                     //최대가능금액         /*미수가능금액*/
	char max_pbl_qtyz18                   [ 18];	char _max_pbl_qtyz18;                     //최대가능수량         /*미수가능수량*/
	char rvb_orn_max_pbl_feez18           [ 18];	char _rvb_orn_max_pbl_feez18;             //미수발생최대가능수수료 /*미수수수료*/
	char csh_orr_pbl_amtz18               [ 18];	char _csh_orr_pbl_amtz18;                 //현금주문가능금액     /*현금가능금액*/
	char csh_orr_pbl_qtyz18               [ 18];	char _csh_orr_pbl_qtyz18;                 //현금주문가능수량     /*현금가능수량*/
	char ost_fee1z18                      [ 18];	char _ost_fee1z18;                        //현금수수료           /*현금수수료*/
	char cfd_rvb_orr_pbl_amtz18           [ 18];	char _cfd_rvb_orr_pbl_amtz18;             //신용미수주문가능금액 /*신용미수가능금액*/
	char cfd_rvb_orr_pbl_qtyz18           [ 18];	char _cfd_rvb_orr_pbl_qtyz18;             //신용미수주문가능수량 /*신용미수가능수량*/
	char cfd_max_pbl_feez18               [ 18];	char _cfd_max_pbl_feez18;                 //신용최대가능수수료   /*신용미수수수료*/
	char cfd_orr_pbl_amtz18               [ 18];	char _cfd_orr_pbl_amtz18;                 //신용주문가능금액     /*신용미발생가능금액*/
	char cfd_orr_pbl_qtyz18               [ 18];	char _cfd_orr_pbl_qtyz18;                 //신용주문가능수량     /*신용미발생가능수량*/
	char ost_fee2z18                      [ 18];	char _ost_fee2z18;                        //수수료2              /*신용미발생수수료*/
	char sdr_xps1z18                      [ 18];	char _sdr_xps1z18;                        //제비용1
	char sdr_xpsz18                       [ 18];	char _sdr_xpsz18;                         //제비용
} Tp8105OutBlock;

typedef struct tagp8105
{
	Tp8105InBlock                     p8105inblock                          ;  //입력
	Tp8105OutBlock                    p8105outblock                         ;  //화면출력
} Tp8105;

typedef struct tagp8104InBlock    //기본입력
{
	char pswd_noz8                        [ 44];	char _pswd_noz8;                          //비밀번호
	char issue_codez6                     [  6];	char _issue_codez6;                       //종목코드
	char gubunz1                          [  1];	char _gubunz1;                            //구분
	char the_datez8                       [  8];	char _the_datez8;                         //대출일               /*신OBM에존재하지않는항목*/
} Tp8104InBlock;

typedef struct tagp8104OutBlock    //화면출력
{
	char issue_codez6                     [  6];	char _issue_codez6;                       //종목코드
	char order_qtyz12                     [ 12];	char _order_qtyz12;                       //매도가능수량
} Tp8104OutBlock;

typedef struct tagp8104
{
	Tp8104InBlock                     p8104inblock                          ;  //기본입력
	Tp8104OutBlock                    p8104outblock                         ;  //화면출력
} Tp8104;








////////////////////////////////////////
//	선물옵션
////////////////////////////////////////

typedef struct tags8301InBlock    //기본입력
{
	char slbuy_typez1                     [  1];	char _slbuy_typez1;                       //매수매도유형
	char passwd_noz8                      [ 44];	char _passwd_noz8;                        //비밀번호
	char issue_codez9                     [  9];	char _issue_codez9;                       //종목코드
	char ord_typez1                       [  1];	char _ord_typez1;                         //주문유형
	char trade_typez1                     [  1];	char _trade_typez1;                       //거래유형
	char order_qtyz8                      [  8];	char _order_qtyz8;                        //주문수량
	char order_pricez12                   [ 12];	char _order_pricez12;                     //주문가격
	char trad_pswd_no_1z8                 [ 44];	char _trad_pswd_no_1z8;                   //거래비밀번호1
	char trad_pswd_no_2z8                 [ 44];	char _trad_pswd_no_2z8;                   //거래비밀번호2
} Ts8301InBlock;

typedef struct tags8301OutBlock    //화면출력
{
	char order_qtyz8                      [  8];	char _order_qtyz8;                        //주문수량
	char order_pricez12                   [ 12];	char _order_pricez12;                     //주문가격
	char order_noz10                      [ 10];	char _order_noz10;                        //주문번호
	char ord_pos_csamtz16                 [ 16];	char _ord_pos_csamtz16;                   //주문가능현금
	char ord_pos_ttamtz16                 [ 16];	char _ord_pos_ttamtz16;                   //주문가능총액
	char ord_insuf_csamtz16               [ 16];	char _ord_insuf_csamtz16;                 //주문부족현금
	char ord_insuf_ttamtz16               [ 16];	char _ord_insuf_ttamtz16;                 //주문부족총액
	char ord_pos_qtyz16                   [ 16];	char _ord_pos_qtyz16;                     //주문가능수량
} Ts8301OutBlock;

typedef struct tags8301
{
	Ts8301InBlock                     s8301inblock                          ;  //기본입력
	Ts8301OutBlock                    s8301outblock                         ;  //화면출력
} Ts8301;


typedef struct tags8302InBlock    //기본입력
{
	char gubunz1                          [  1];	char _gubunz1;                            //정정/취소구분
	char slbuy_typez1                     [  1];	char _slbuy_typez1;                       //매수매도유형
	char pswd_noz8                        [ 44];	char _pswd_noz8;                          //비밀번호
	char orgnl_order_noz10                [ 10];	char _orgnl_order_noz10;                  //원주문번호
	char issue_codez9                     [  9];	char _issue_codez9;                       //종목코드
	char orgnl_ord_typez1                 [  1];	char _orgnl_ord_typez1;                   //원주문유형
	char crctn_ord_typez1                 [  1];	char _crctn_ord_typez1;                   //정정주문유형
	char order_qtyz8                      [  8];	char _order_qtyz8;                        //주문수량
	char order_pricez12                   [ 12];	char _order_pricez12;                     //주문가격
	char trad_pswd_no_1z8                 [ 44];	char _trad_pswd_no_1z8;                   //거래비밀번호1
	char trad_pswd_no_2z8                 [ 44];	char _trad_pswd_no_2z8;                   //거래비밀번호2
} Ts8302InBlock;

typedef struct tags8302OutBlock    //화면출력
{
	char order_noz10                      [ 10];	char _order_noz10;                        //주문번호
	char ord_pos_csamtz16                 [ 16];	char _ord_pos_csamtz16;                   //주문가능현금
	char ord_pos_ttamtz16                 [ 16];	char _ord_pos_ttamtz16;                   //주문가능총액
	char ord_insuf_csamtz16               [ 16];	char _ord_insuf_csamtz16;                 //주문부족현금
	char ord_insuf_ttamtz16               [ 16];	char _ord_insuf_ttamtz16;                 //주문부족총액
	char ord_pos_qtyz16                   [ 16];	char _ord_pos_qtyz16;                     //주문가능수량
	char order_qtyz8                      [  8];	char _order_qtyz8;                        //주문수량
	char order_pricez12                   [ 12];	char _order_pricez12;                     //주문가격
} Ts8302OutBlock;

typedef struct tags8302
{
	Ts8302InBlock                     s8302inblock                          ;  //기본입력
	Ts8302OutBlock                    s8302outblock                         ;  //화면출력
} Ts8302;

typedef struct tagc8311InBlock    //입력
{
	char in_pswdz8                        [ 44];	char _in_pswdz8;                          //비밀번호
	char job_gubunz1                      [  1];	char _job_gubunz1;                        //구분
	char trobj_stock_codez12              [ 12];	char _trobj_stock_codez12;                //대상자산코드         /*사용안함*/
} Tc8311InBlock;

typedef struct tagc8311OutBlock    //화면출력
{
	char out_acnamez30                    [ 30];	char _out_acnamez30;                      //계좌명
	char out_ttamtz14                     [ 14];	char _out_ttamtz14;                       //예탁총액
	char out_cashz14                      [ 14];	char _out_cashz14;                        //예탁현금
	char out_mrgn_ttamtz14                [ 14];	char _out_mrgn_ttamtz14;                  //증거금총액
	char out_mrgn_cashz14                 [ 14];	char _out_mrgn_cashz14;                   //증거금현금
	char out_order_ttamtz14               [ 14];	char _out_order_ttamtz14;                 //주문가능총액
	char out_order_cashz14                [ 14];	char _out_order_cashz14;                  //주문가능현금
	char out_tot_ttamtz14                 [ 14];	char _out_tot_ttamtz14;                   //총평가총액
	char out_tot_cashz14                  [ 14];	char _out_tot_cashz14;                    //총평가현금
	char out_substz14                     [ 14];	char _out_substz14;                       //대용금액
	char out_renewz14                     [ 14];	char _out_renewz14;                       //정산차금
	char out_stlmtz14                     [ 14];	char _out_stlmtz14;                       //최종차금
	char out_opt_sellz14                  [ 14];	char _out_opt_sellz14;                    //옵션매도대금
	char out_opt_buyz14                   [ 14];	char _out_opt_buyz14;                     //옵션매수대금
	char out_opt_valuez14                 [ 14];	char _out_opt_valuez14;                   //옵션평가금액
	char out_pred_substz14                [ 14];	char _out_pred_substz14;                  //전일대용매도
	char out_thday_substz14               [ 14];	char _out_thday_substz14;                 //당일대용매도
	char out_pred_amtz14                  [ 14];	char _out_pred_amtz14;                    //전일가입금
	char out_opt_hangz14                  [ 14];	char _out_opt_hangz14;                    //옵션행사금액
	char out_opt_baiz14                   [ 14];	char _out_opt_baiz14;                     //옵션배정금액
	char out_thday_amtz14                 [ 14];	char _out_thday_amtz14;                   //당일가입금
	char out_rcblez14                     [ 14];	char _out_rcblez14;                       //미수금
	char out_ovamtz14                     [ 14];	char _out_ovamtz14;                       //연체료
	char out_ftr_cmsnz14                  [ 14];	char _out_ftr_cmsnz14;                    //선물수수료
	char out_opt_cmsnz14                  [ 14];	char _out_opt_cmsnz14;                    //옵션수수료
	char out_afterz14                     [ 14];	char _out_afterz14;                       //세후이용료
	char out_asset_ttamtz14               [ 14];	char _out_asset_ttamtz14;                 //순자산총액
	char out_asset_ttcashz14              [ 14];	char _out_asset_ttcashz14;                //순자산현금
} Tc8311OutBlock;

typedef struct tagc8311OutBlock1    //출력, [반복]
{
	char out_issuez9                      [  9];	char _out_issuez9;                        //종목코드
	char out_isnamez30                    [ 30];	char _out_isnamez30;                      //종목명
	char out_slbyz6                       [  6];	char _out_slbyz6;                         //매매구분
	char out_qtyz14                       [ 14];	char _out_qtyz14;                         //수량
	char out_averz12                      [ 12];	char _out_averz12;                        //평균가
	char out_pricez12                     [ 12];	char _out_pricez12;                       //현재가
	char out_lsnpfz14                     [ 14];	char _out_lsnpfz14;                       //평가손익
	char today_revs_odqtyz12              [ 12];	char _today_revs_odqtyz12;                //주문수량
	char sell_posbl_qtyz12                [ 12];	char _sell_posbl_qtyz12;                  //청산가능수량
} Tc8311OutBlock1;

typedef struct tagc8311
{
	Tc8311InBlock                     c8311inblock                          ;  //입력
	Tc8311OutBlock                    c8311outblock                         ;  //화면출력
	Tc8311OutBlock1                   c8311outblock1                   [ 10];  //출력 , [반복]
} Tc8311;

typedef struct tagc8322InBlock    //기본입력
{
	char pswd_noz8                        [ 44];	char _pswd_noz8;                          //비밀번호
	char order_datez8                     [  8];	char _order_datez8;                       //주문일자
	char issue_codez9                     [  9];	char _issue_codez9;                       //종목번호
	char conc_gubunz1                     [  1];	char _conc_gubunz1;                       //체결구분             0:전체1:미체결2:체결
	char sort_gubunz1                     [  1];	char _sort_gubunz1;                       //정렬구분             0:주문번호순1:주문번호역순
	char sl_buy_typez1                    [  1];	char _sl_buy_typez1;                      //매수도구분           0:전체1:매도2:매수
	char issue_gubunz1                    [  1];	char _issue_gubunz1;                      //종목구분             0:전체1:선물2:옵션
	char disp_gubunz1                     [  1];	char _disp_gubunz1;                       //조회구분             1:주문시간순2:체결시간순
	char cts_areaz67                      [ 67];	char _cts_areaz67;                        //CTS                  
	char trad_pswd1z8                     [ 44];	char _trad_pswd1z8;                       //거래비밀번호1
	char trad_pswd2z8                     [ 44];	char _trad_pswd2z8;                       //거래비밀번호2
	char IsPageUp                         [  1];	char _IsPageUp;                           //ISPAGEUP             
} Tc8322InBlock;

typedef struct tagc8322OutBlock    //화면출력
{
	char order_datez8                     [  8];	char _order_datez8;                       //주문일자
	char accnt_namez40                    [ 40];	char _accnt_namez40;                      //계좌명
} Tc8322OutBlock;

typedef struct tagc8322OutBlock1    //화면출력1, [반복]
{
	char order_noz10                      [ 10];	char _order_noz10;                        //주문번호
	char orgnl_order_noz10                [ 10];	char _orgnl_order_noz10;                  //원주문번호
	char conc_noz6                        [  6];	char _conc_noz6;                          //체결번호
	char order_datez8                     [  8];	char _order_datez8;                       //주문일자
	char slbuy_typez10                    [ 10];	char _slbuy_typez10;                      //주문구분
	char slby_typez15                     [ 15];	char _slby_typez15;                       //매매구분
	char issue_codez9                     [  9];	char _issue_codez9;                       //종목번호
	char issue_namez40                    [ 40];	char _issue_namez40;                      //종목명
	char order_qtyz12                     [ 12];	char _order_qtyz12;                       //주문수량
	char order_pricez12                   [ 12];	char _order_pricez12;                     //주문단가
	char conc_qtyz12                      [ 12];	char _conc_qtyz12;                        //체결수량
	char index_conc_pricez12              [ 12];	char _index_conc_pricez12;                //체결평균단가
	char unconc_qtyz12                    [ 12];	char _unconc_qtyz12;                      //미체결수량
	char rqst_typez6                      [  6];	char _rqst_typez6;                        //처리구분
	char rqst_timez8                      [  8];	char _rqst_timez8;                        //처리시간
} Tc8322OutBlock1;

typedef struct tagc8322OutBlock_IN    //Button정보
{
	char cts_areaz67                      [ 67];	char _cts_areaz67;                        //CTS                  
	char nextbutton                       [  1];	char _nextbutton;                         //NEXTBUTTON           
} Tc8322OutBlock_IN;

typedef struct tagc8322
{
	Tc8322InBlock                     c8322inblock                          ;  //기본입력
	Tc8322OutBlock                    c8322outblock                         ;  //화면출력
	Tc8322OutBlock1                   c8322outblock1                   [ 20];  //화면출력1 , [반복]
	Tc8322OutBlock_IN                 c8322outblock_in                      ;  //Button정보
} Tc8322;

typedef struct tagp8301InBlock    //입력
{
	char issue_codez9                     [  9];	char _issue_codez9;                       //종목코드
	char slbuy_typez1                     [  1];	char _slbuy_typez1;                       //매매구분
	char ord_typez1                       [  1];	char _ord_typez1;                         //주문유형
	char order_pricez12                   [ 12];	char _order_pricez12;                     //주문단가
	char passwd_noz8                      [ 44];	char _passwd_noz8;                        //비밀번호
} Tp8301InBlock;

typedef struct tagp8301OutBlock    //화면출력
{
	char accnt_namez20                    [ 20];	char _accnt_namez20;                      //계좌명
	char kor_issue_namez40                [ 40];	char _kor_issue_namez40;                  //종목명
	char be_dpsit_ttamtz14                [ 14];	char _be_dpsit_ttamtz14;                  //주문전
	char be_dpsit_cash_amtz14             [ 14];	char _be_dpsit_cash_amtz14;               //주문전
	char be_brkrg_mrgn_ttamz14            [ 14];	char _be_brkrg_mrgn_ttamz14;              //주문전
	char be_brkrg_cash_mgamz14            [ 14];	char _be_brkrg_cash_mgamz14;              //주문전
	char be_order_pos_ttamtz14            [ 14];	char _be_order_pos_ttamtz14;              //주문전
	char be_order_pos_csamtz14            [ 14];	char _be_order_pos_csamtz14;              //주문전
	char af_dpsit_ttamtz14                [ 14];	char _af_dpsit_ttamtz14;                  //주문후
	char af_dpsit_cash_amtz14             [ 14];	char _af_dpsit_cash_amtz14;               //주문후
	char af_brkrg_mrgn_ttamz14            [ 14];	char _af_brkrg_mrgn_ttamz14;              //주문후
	char af_brkrg_cash_mgamz14            [ 14];	char _af_brkrg_cash_mgamz14;              //주문후
	char af_order_pos_ttamtz14            [ 14];	char _af_order_pos_ttamtz14;              //주문후
	char af_order_pos_csamtz14            [ 14];	char _af_order_pos_csamtz14;              //주문후
	char new_pos_qtyz14                   [ 14];	char _new_pos_qtyz14;                     //신규주문
	char bal_pos_qtyz14                   [ 14];	char _bal_pos_qtyz14;                     //잔고주문
	char tot_pos_qtyz14                   [ 14];	char _tot_pos_qtyz14;                     //TOTAL                
} Tp8301OutBlock;

typedef struct tagp8301
{
	Tp8301InBlock                     p8301inblock                          ;  //입력
	Tp8301OutBlock                    p8301outblock                         ;  //화면출력
} Tp8301;

typedef struct tagp8302InBlock    //입력
{
	char proc_gubunz1                     [  1];	char _proc_gubunz1;                       //처리구분
	char trobj_stock_codez12              [ 12];	char _trobj_stock_codez12;                //대상자산코드
} Tp8302InBlock;

typedef struct tagp8302OutBlock    //화면출력
{
	char o_accnt_namez40                  [ 40];	char _o_accnt_namez40;                    //계좌명
} Tp8302OutBlock;

typedef struct tagp8302OutBlock1    //계좌명, [반복]
{
	char index_issue_codez9               [  9];	char _index_issue_codez9;                 //종목
	char index_issue_namez30              [ 30];	char _index_issue_namez30;                //종목명
	char index_slbuy_typez2               [  2];	char _index_slbuy_typez2;                 //매매구분
	char index_slbuy_namez6               [  6];	char _index_slbuy_namez6;                 //매매구분명
	char bal_qtyz12                       [ 12];	char _bal_qtyz12;                         //잔고수량
	char today_revs_odqtyz12              [ 12];	char _today_revs_odqtyz12;                //주문수량
	char sell_posbl_qtyz12                [ 12];	char _sell_posbl_qtyz12;                  //청산가능수량
	char avrgez14                         [ 14];	char _avrgez14;                           //평균가
} Tp8302OutBlock1;

typedef struct tagp8302
{
	Tp8302InBlock                     p8302inblock                          ;  //입력
	Tp8302OutBlock                    p8302outblock                         ;  //화면출력
	Tp8302OutBlock1                   p8302outblock1                   [ 20];  //계좌명 , [반복]
} Tp8302;








////////////////////////////////////////
//	실시간 패킷
////////////////////////////////////////

typedef struct tagd2OutBlock    //출력
{
	char userid                           [  8];   //사용자ID
	char itemgb                           [  1];   //ITEM구분
	char accountno                        [ 11];   //계좌번호
	char orderno                          [ 10];   //주문번호
	char issuecd                          [ 15];   //종목코드
	char slbygb                           [  1];   //매도수구분
	char concgty                          [ 10];   //체결수량
	char concprc                          [ 11];   //체결가격
	char conctime                         [  6];   //체결시간
	char ucgb                             [  1];   //정정취소구분
	char rejgb                            [  1];   //거부구분
	char fundcode                         [  3];   //펀드코드
	char sin_gb                           [  2];   //신용구분
	char loan_date                        [  8];   //대출일자
	char ato_ord_tpe_chg                  [  1];   //선물옵션주문유형변경여부
	char filler                           [ 34];   //filler           
} Td2OutBlock;

typedef struct tagd3OutBlock    //출력
{
	char userid                           [  8];   //USERID               
	char itemgb                           [  1];   //ITEM구분
	char accountno                        [ 11];   //계좌번호
	char orderno                          [ 10];   //주문번호
	char orgordno                         [ 10];   //원주문번호
	char ordercd                          [  2];   //주문구분
	char issuecd                          [ 15];   //종목코드
	char issuename                        [ 20];   //종목명
	char slbygb                           [  1];   //매매구분
	char order_type                       [  2];   //주문유형
	char ordergty                         [ 10];   //주문수량
	char orderprc                         [ 11];   //주문단가
	char procnm                           [  2];   //처리구분
	char commcd                           [  2];   //매체구분
	char order_cond                       [  1];   //주문조건1
	char fundcode                         [  3];   //펀드코드
	char sin_gb                           [  2];   //신용구분
	char order_time                       [  6];   //주문시간
	char loan_date                        [  8];   //대출일자
} Td3OutBlock;
