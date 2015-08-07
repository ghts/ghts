''' This file is part of GHTS.

GHTS is free software: you can redistribute it and/or modify
it under the terms of the GNU Lesser General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

GHTS is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Lesser General Public License for more details.

You should have received a copy of the GNU Lesser General Public License
along with GHTS.  If not, see <http://www.gnu.org/licenses/>.

@author: UnHa Kim <unha.kim.ghts@gmail.com> '''

import random
import sys
import time
import zmq

P메시지_GET = "G"
P메시지_종료 = "Q"
P회신_OK = "O"
UTF8 = 'utf-8'

def 테스트용_주소정보_요청_모듈(P주소_주소정보, P주소_테스트_결과, 테스트_반복횟수):
    #print("주소정보 요청 : 시작")
    context = zmq.Context()
    
    주소정보_REQ = context.socket(zmq.REQ)
    주소정보_REQ.connect(P주소_주소정보)
    
    테스트_결과_REQ = context.socket(zmq.REQ)
    테스트_결과_REQ.connect(P주소_테스트_결과)
    
    #print("주소정보 요청 : 초기화 완료")
    
    질의값_모음 = []
    질의값_모음.append("주소정보")
    질의값_모음.append("종목정보")
    질의값_모음.append("가격정보_입수")
    질의값_모음.append("가격정보_배포")
    질의값_모음.append("가격정보")
    질의값_모음.append("테스트_결과")
    
    테스트_결과 = True
    
    for 반복횟수 in range(테스트_반복횟수):
        질의값 = 질의값_모음[random.randint(0, len(질의값_모음) - 1)]
        메시지 = [P메시지_GET.encode(UTF8), 질의값.encode(UTF8)]
        
        #print("제공 : send_multipart() 시작")
        주소정보_REQ.send_multipart(메시지)
        #print("제공 : send_multipart() 완료")
        
        #print("제공 : recv_multipart() 시작")
        메시지 = 주소정보_REQ.recv_multipart()
        #print("제공 : recv_multipart() 완료")
        
        구분 = 메시지[0].decode(UTF8)
        데이터 = 메시지[1].decode(UTF8)
        
        if 구분 == P회신_OK and 데이터.startswith("tcp://127.0.0.1:"):
            continue
        
        if 구분 != P회신_OK:
            print("제공 : 에러 메시지 회신.", 데이터)
            테스트_결과 = False
            break
        else:
            print("제공 : 회신값이 예상과 다름.", 데이터)
            테스트_결과 = False
            break    
    
    #print("주소정보 요청 : 테스트 결과 전송 시작")
    메시지 = [P메시지_GET.encode(UTF8), str(테스트_결과).encode(UTF8)]
    테스트_결과_REQ.send_multipart(메시지)
    테스트_결과_REQ.recv_multipart()
    #print("주소정보 요청 : 테스트 결과 전송 완료")
    
    # 리소스 정리 후 종료
    테스트_결과_REQ.close()
    context.destroy()
    
    #print("주소정보 요청 : 종료")

if __name__ == "__main__":
    P주소_주소정보 = sys.argv[1]
    P주소_테스트_결과 = sys.argv[2]
    테스트_반복횟수 = int(sys.argv[3])
    
    테스트용_주소정보_요청_모듈(P주소_주소정보, P주소_테스트_결과, 테스트_반복횟수)