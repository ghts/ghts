GHTS
====

- 알고리즘 트레이딩 시스템 구현에 유용한 라이브러리.  
- Go언어 기반

*********************************************************

디렉토리별 기능 설명  
- lib : 공용 기능.
- xing/base : Xing API c32/go 공용 자료형.
- xing/c32 : Xing API DLL을 호출하는 32비트 모듈.
- xing/go : Xing API 기능 호출 (32/64비트 모두 가능)

*********************************************************

사전준비물
- Go언어 : https://golang.org/dl/
- C/C++ 컴파일러 및 ZeroMQ (MSYS2) : https://www.msys2.org/

*********************************************************
MSYS2 설치 후 'MSYS2 MSYS' 터미널을 열고 아래 명령을 실행한다.

<pre><code>pacman -Syuu 
pacman -S base-devel
pacman -S mingw-w64-i686-toolchain
pacman -S mingw-w64-x86_64-toolchain
pacman -S mingw-w64-i686-zeromq
pacman -S mingw-w64-x86_64-zeromq
pacman -S mingw-w64-x86_64-{git,git-doc-html,git-doc-man,curl} git-extra</code></pre>

이후 가끔씩 모든 패키지를 업데이트 하려고 할 때는 다음 명령어를 실행한다.
<pre><code>pacman -Syuu</code></pre>

*********************************************************
GHTS 라이브러리 설치

<pre><code>go get github.com/ghts/ghts</code></pre>
 
*********************************************************

악성코드로 잘못 진단되는 경우.

금융기관이 설치하는 보안 방화벽 '안랩 세이프 트랜잭션'(Ahnlab Safe Transaction)에서 GHTS가 멀웨어로 잘못 진단됩니다.

'안랩 세이프 트랜잭션'(Ahnlab Safe Transaction)이 실행된 상태에서 GHTS를 '위협 행위 차단'을 체크 해제해야 합니다.

윈도우 기본 백신인 '윈도우 디펜더'에서는 아무런 문제가 없습니다. 

*********************************************************    
  
<주의>
------
저작권자, 개발자, 개발에 참여한 기여자들은 이 소프트웨어에 대한 어떠한 보증도 하지 않습니다.  
이 소프트웨어를 사용하면서 발생하는 그 어떠한 손실 및 손상에 대해서 책임지지 않습니다.
소스코드 파일에 별도의 언급이 없는 한, 모든 소스코드는 GNU LGPL V2.1 라이센스를 따릅니다.  
저작권에 대한 자세한 사항은 'LICENSE' 파일을 참고하십시오.
