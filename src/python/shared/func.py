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

import datetime
import pytz

# If want to check all of the timezones in pytz,
#    for tz in pytz.all_timezones:
#        print tz


_timezone_utc = pytz.utc
_timezone_ko = pytz.timezone('Asia/Seoul')


def ko시점(시점: datetime.datetime) -> datetime.datetime:
    if 시점.tzinfo() is None:
        raise ValueError("ko시점() : No timezone info.")

    return 시점.astimezone(_timezone_ko)


def utc시점(시점: datetime.datetime) -> datetime.datetime:
    if 시점.tzinfo() is None:
        raise ValueError("utc시점() : No timezone info.")

    return 시점.astimezone(_timezone_utc)


def now_utc() -> datetime.datetime:
    return datetime.datetime.utcnow().replace(tzinfo=pytz.utc)


def now_ko() -> datetime.datetime:
    return ko시점(now_utc())


if __name__ == '__main__':
    # 테스트 코드를 여기에 넣을 것.
    pass
