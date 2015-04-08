'''
This file is part of GHTS.

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

Created on 2015. 4. 5.

@author: UnHa Kim <unha.kim@gh-system.com>
'''

import sys
import time
import zmq

P메시지_구분_일반 = "N"
P메시지_구분_종료 = "Q"
P회신_OK = "O"
UTF8 = 'utf-8'

def 테스트용_가격정보_제공_모듈(P가격정보_입수_주소, P가격정보_배포횟수):
    #print("제공 : 시작")
    
    context = zmq.Context()
    가격정보_입수_소켓 = context.socket(zmq.REQ)
    가격정보_입수_소켓.connect(P가격정보_입수_주소)
    
    # 다른 모듈들이 안정화 될 때까지 잠시 대기.
    #print("제공 : 2초간 대기 시작")
    time.sleep(2)
    #print("제공 : 2초간 대기 완료")
    
    for 반복횟수 in range(P가격정보_배포횟수):
        모의_가격_데이터 = str(반복횟수 * 10)
        
        메시지 = [P메시지_구분_일반.encode(UTF8), 모의_가격_데이터.encode(UTF8)]
        
        #print("제공 : send_multipart() 시작")
        가격정보_입수_소켓.send_multipart(메시지)
        #print("제공 : send_multipart() 완료")
        
        #print("제공 : recv_multipart() 시작")
        메시지 = 가격정보_입수_소켓.recv_multipart()
        #print("제공 : recv_multipart() 완료")
        
        구분 = 메시지[0].decode(UTF8)
        데이터 = 메시지[1].decode(UTF8)
        
        if 구분 != P회신_OK:
            print("제공 : 회신 메시지 구분이 예상과 다름.", 구분)
        
        if 데이터 != "":
            print("제공 : 회신 메시지 구분이 예상과 다름.", 데이터)
     
    #print("제공 : 종료 메시지 send_multipart() 시작")
    메시지 = [P메시지_구분_종료.encode(UTF8), "".encode(UTF8)]
    가격정보_입수_소켓.send_multipart(메시지)
    #print("제공 : 종료 메시지 send_multipart() 완료")
        
    #print("제공 : 종료 메시지 회신 recv_multipart() 시작")
    메시지 = 가격정보_입수_소켓.recv_multipart()
    #print("제공 : 종료 메시지 회신 recv_multipart() 완료")
           
    가격정보_입수_소켓.close()
    context.destroy()
    #print("제공 : 종료")
    
def 테스트용_가격정보_구독_모듈(P가격정보_배포_주소, P테스트_결과_회신_주소):
    #print("구독 : 시작")
    
    context = zmq.Context()
    
    가격정보_구독_소켓 = context.socket(zmq.SUB)
    가격정보_구독_소켓.connect(P가격정보_배포_주소)
    가격정보_구독_소켓.setsockopt(zmq.SUBSCRIBE, b"")
    
    테스트_결과_회신_소켓 = context.socket(zmq.REQ)
    테스트_결과_회신_소켓.connect(P테스트_결과_회신_주소)
    
    #print("구독 : 초기화 완료")
    
    반복횟수 = 0
    
    while True:
        #print("구독 : recv_multipart() 시작")
        메시지 = 가격정보_구독_소켓.recv_multipart()
        #print("구독 : recv_multipart() 완료")
        
        구분 = 메시지[0].decode(UTF8)
        데이터 = 메시지[1].decode(UTF8)
        
        if 구분 == P메시지_구분_일반:
            반복횟수 += 1
            #print("구독 : continue")
            continue
        
        elif 구분 == P메시지_구분_종료:
            if 데이터 != "":
                print("테스트용_가격정보_구독_모듈() : 종료 메시지 데이터가 예상과 다름.", 데이터)
            
            #print("구독 : 종료신호 수신")
            break
    
    반복횟수 = str(반복횟수)
    메시지 = [P메시지_구분_일반.encode(UTF8), 반복횟수.encode(UTF8)]
    
    #print("구독 : 결과 send_multipart() 시작")
    테스트_결과_회신_소켓.send_multipart(메시지)
    #print("구독 : 결과 send_multipart() 종료")
    
    #print("구독 : 결과 recv_multipart() 시작")
    테스트_결과_회신_소켓.recv_multipart()
    #print("구독 : 결과 recv_multipart() 종료")
    
    가격정보_구독_소켓.close()
    테스트_결과_회신_소켓.close()
    context.destroy()
    
    #print("구독 : 종료")

if __name__ == "__main__":    
    구분 = sys.argv[1]
    
    if 구분 == "provider":
        P가격정보_입수_주소 = sys.argv[2]
        P가격정보_배포횟수 = int(sys.argv[3])
    
        테스트용_가격정보_제공_모듈(P가격정보_입수_주소, P가격정보_배포횟수)
    elif 구분 == "subscriber" :
        P가격정보_배포_주소 = sys.argv[2]
        P테스트_결과_회신_주소 = sys.argv[3]

        테스트용_가격정보_구독_모듈(P가격정보_배포_주소, P테스트_결과_회신_주소)
    else:
        print("예상치 못한 파라메터.", 구분)