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

Created on 2015. 3. 1.

@author: UnHa Kim <unha.kim@gh-system.com>
'''

from decimal import Decimal
from .func import *


class S종목:

    def __init__(self, 코드: str, 이름: str):
        self._코드 = 코드
        self._이름 = 이름

    def G코드(self):
        self._코드

    def G이름(self):
        self._이름


class S가격정보:

    def __init__(self, 종목: S종목, 가격: Decimal, UTC시점: datetime=now_utc()):
        self._종목 = 종목
        self.S가격(self, 가격)
        self.S_UTC시점(self, UTC시점)

    def G종목(self) -> S종목:
        self._종목

    def G가격(self) -> Decimal:
        self._가격

    def G_UTC시점(self) -> datetime:
        self._UTC시점

    def G_KO시점(self) -> datetime:
        ko시점(self._UTC시점)

    def S가격(self, 가격: Decimal) -> Decimal:
        self._가격 = 가격

    def S_UTC시점(self, UTC시점: datetime):
        if UTC시점.tzinfo != pytz.utc:
            format_string = "Wrong Timezone. Need UTC time." \
                            + "Current timezone is {0}"
            raise ValueError(format_string.format(UTC시점.tzinfo))

        self._UTC시점 = UTC시점


class S가격정보_모음:

    def __init__(self):
        self._가격정보_모음 = {}

    def S가격정보(self, 가격정보: S가격정보):
        self._가격정보_모음[가격정보.G종목().G코드()] = 가격정보

    def G가격정보(self, 종목: S종목) -> S가격정보:
        # get은 존재하지 않는 키에 대해서 None을 반환한다.
        self._가격정보_모음.get(종목.G코드())


class S자산내역_구성요소:

    def __init__(self, 종목: S종목, 수량: int, 최근_가격: Decimal):
        self._종목 = 종목
        self._수량 = 수량
        self._최근_가격 = 최근_가격

    def G종목(self) -> S종목:
        self._종목

    def G수량(self) -> int:
        self._수량

    def G최근_가격(self) -> Decimal:
        # 일단은 내부에 저장하는 형태로 구현.
        # 이후에는 'S가격정보'에서 최근 가격을 불러오는 형태가 될 수도 있을 듯.
        self._최근_가격

    def G평가액(self) -> Decimal:
        Decimal(self._수량) * self.G최근_가격()


# 자산내역 공유 데이터. (별도의 프로세스 필요?)
class S자산내역:

    def __init__(self):
        self._자산내역 = {}

    def S추가(self, 구성요소: S자산내역_구성요소):
        self._자산내역[구성요소.G종목().G코드()] = 구성요소

    def S삭제(self, 종목: S종목):
        del self._자산내역[종목.G코드()]

    def G보유종목(self) -> []:
        보유종목 = []

        for 구성요소 in self._자산내역.values():
            보유종목.append(구성요소.G종목())

        return 보유종목

    def G수량(self, 종목: S종목) -> int:
        if 종목.G코드() not in self._자산내역.keys():
            return Decimal(0)

        self._자산내역[종목.G코드()].G수량()

    def G종목별_평가액(self, 종목: S종목) -> Decimal:
        if 종목.G코드() not in self._자산내역.keys():
            return Decimal(0)

        self._자산내역[종목.G코드()].G평가액()

    def G전체_평가액(self) -> Decimal:
        전체_평가액 = Decimal("0.0")

        for 자산내역_구성요소 in self._자산내역.values():
            전체_평가액 = 전체_평가액 + 자산내역_구성요소.G평가액()
            전체_평가액


if __name__ == '__main__':
    # 테스트 코드를 여기에 넣을 것.
    pass
