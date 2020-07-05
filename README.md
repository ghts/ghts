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
- C언어 컴파일러 (Msys2) : https://www.msys2.org/

*********************************************************
Msys2 C언어 컴파일러 추가 설치

“MSYS2 MSYS” 터미널에서 아래 명령을 실행하여 업데이트한다.
<pre><code>pacman –Syu</code></pre>

“MSYS2 MSYS” 터미널을 종료했다가 다시 실행한 후, 아래 명령을 실행하여 업데이트를 계속한다.
<pre><code>pacman -Su</code></pre>

“MSYS2 MSYS” 터미널을 종료했다가 다시 실행한 후, 아래 명령을 실행하여 compiler와 그밖에 필요한 package를 설치한다. 
<pre><code>pacman -S --needed base-devel
pacman -S --needed mingw-w64-i686-toolchain
pacman -S --needed mingw-w64-x86_64-toolchain</code></pre>

*********************************************************
GHTS 라이브러리 설치

설치방법
<pre><code>go get github.com/ghts/ghts</code></pre>
  
*********************************************************  
  
<주의>
------
저작권자, 개발자, 개발에 참여한 기여자들은 이 소프트웨어에 대한 어떠한 보증도 하지 않습니다.  
이 소프트웨어를 사용하면서 발생하는 그 어떠한 손실 및 손상에 대해서 책임지지 않습니다.
소스코드 파일에 별도의 언급이 없는 한, 모든 소스코드는 GNU LGPL V2.1 라이센스를 따릅니다.  
저작권에 대한 자세한 사항은 'LICENSE' 파일을 참고하십시오.
