''' This file is part of GHTS.

GHTS is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

GHTS is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with GHTS.  If not, see <http://www.gnu.org/licenses/>.

@author: UnHa Kim <unha.kim.ghts@gmail.com> '''

import random
import sys
import time
import zmq

P메시지_구분_일반 = "N"
P메시지_구분_종료 = "Q"
P회신_OK = "O"
UTF8 = 'utf-8'

def 테스트용_종목정보_요청_모듈(P주소_종목정보, P주소_테스트_결과_회신, 테스트_반복횟수):
    #print("종목정보 요청 : 시작")
    
    context = zmq.Context()
    
    종목정보_REQ = context.socket(zmq.REQ)
    종목정보_REQ.connect(P주소_종목정보)
    
    테스트_결과_REQ = context.socket(zmq.REQ)
    테스트_결과_REQ.connect(P주소_테스트_결과_회신)
    
    #print("종목정보 요청 : 초기화 완료")
    
    질의예상값_모음 = []
    질의예상값_모음.append(["000020", "동화약품"])
    질의예상값_모음.append(["000030", "우리은행"])
    질의예상값_모음.append(["000040", "KR모터스"])
    질의예상값_모음.append(["000050", "경방"])
    질의예상값_모음.append(["000060", "메리츠화재"])    
    
    테스트_결과 = True
    
    for 반복횟수 in range(테스트_반복횟수):
        질의예상값 = 질의예상값_모음[random.randint(0, len(질의예상값_모음) - 1)]
        질의값 = 질의예상값[0]
        
        송신_메시지 = [P메시지_구분_일반.encode(UTF8), 질의값.encode(UTF8)]
        
        #print("제공 : send_multipart() 시작")
        종목정보_REQ.send_multipart(송신_메시지)
        #print("제공 : send_multipart() 완료")
        
        #print("제공 : recv_multipart() 시작")
        수신_메시지 = 종목정보_REQ.recv_multipart()
        #print("제공 : recv_multipart() 완료")
        
        구분 = 수신_메시지[0].decode(UTF8)
        
        if 구분 != P회신_OK:
            에러_메시지 = 수신_메시지[1].decode(UTF8)
            print("제공 : 에러 메시지 회신.", 구분, 에러_메시지)
            테스트_결과 = False
            break
        
        종목코드 = 수신_메시지[1].decode(UTF8)
        종목이름 = 수신_메시지[2].decode(UTF8)
        
        if 종목코드 == 질의예상값[0] and 종목이름 == 질의예상값[1]:
            continue       
        
        print("제공 : 회신값이 예상과 다름.", 질의값, 종목코드, 종목이름)
        테스트_결과 = False
        break
    
    #print("종목정보 요청 : 테스트 결과 전송 시작")
    메시지 = [P메시지_구분_일반.encode(UTF8), str(    테스트_결과).encode(UTF8)]
    테스트_결과_REQ.send_multipart(메시지)
    테스트_결과_REQ.recv_multipart()
    #print("종목정보 요청 : 테스트 결과 전송 완료")
    
    # 리소스 정리 후 종료
    테스트_결과_REQ.close()
    context.destroy()
    
    #print("종목정보 요청 : 종료")

if __name__ == "__main__":
    P주소_종목정보 = sys.argv[1]
    P주소_테스트_결과_회신 = sys.argv[2]
    테스트_반복횟수 = int(sys.argv[3])
    테스트용_종목정보_요청_모듈(P주소_종목정보, P주소_테스트_결과_회신, 테스트_반복횟수)