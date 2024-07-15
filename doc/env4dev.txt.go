package doc

// 개발용 환경변수.
//
// 테스트 실행을 위해서는 로그인 정보를 환경변수에 저장해야 한다.
// 모의서버에서 테스트할 때 공인인증암호는 없어도 되며, 계좌비밀번호도 실제가 아닌 가상의 번호를 넣어도 된다.
//
// XING_LOG_IN_ID=<ID>;XING_LOG_IN_PWD=<암호>;XING_CERT_PWD='';XING_ACCOUNT_PWD=<가상번호>;XING_TEST_PWD=<모의투자암호>
// DLL 테스트할 때는 32비트 컴파일이 필요하므로 'GOARCH=386'을 추가할 것.
