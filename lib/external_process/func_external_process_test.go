package external_process

import (
	"encoding/gob"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"

	lb "github.com/ghts/ghts/lib"
	ps "github.com/mitchellh/go-ps"
)

// t격리 : 테스트마다 임시 디렉토리로 작업 디렉토리를 옮기고, 패키지 전역 상태(저장소, 초기화 플래그)를 초기화한다.
func t격리(t *testing.T) {
	t.Helper()

	종전_디렉토리, 에러 := os.Getwd()
	if 에러 != nil {
		t.Fatal(에러)
	}

	if 에러 = os.Chdir(t.TempDir()); 에러 != nil {
		t.Fatal(에러)
	}
	t.Cleanup(func() { _ = os.Chdir(종전_디렉토리) })

	// 저장소는 패키지 변수(고정 주소)에 걸린 포인터 수신자 메서드로 호출되므로,
	// 값 자체를 새 구조체로 대체하면 다른 루틴의 잠금 상태를 파괴한다. 내용만 초기화한다.
	프로세스ID_저장소.Lock()
	프로세스ID_저장소.저장소 = make(map[int]string)
	프로세스ID_저장소.Unlock()

	// atomic.Bool은 복사할 수 없으므로(값 재대체 금지), 내용만 초기화한다.
	프로세스ID_목록_파일_초기화_완료.Store(false)
}

// f테스트_핑_경로 : 테스트용 장수 프로세스(ping -n 60 localhost)의 절대 경로를 구한다. 없으면 생략.
func f테스트_핑_경로(t *testing.T) string {
	t.Helper()

	경로, 에러 := exec.LookPath("ping")
	if 에러 != nil {
		t.Skipf("ping.exe을 찾을 수 없어 이 테스트는 생략합니다. (%v)", 에러)
	}

	return 경로
}

// f테스트_정규화_경로 : 저장부가 기록하는 경로(Abs+Clean)와 동일한 정규화로 기대 값을 만든다.
func f테스트_정규화_경로(t *testing.T, 경로 string) string {
	t.Helper()

	절대_경로, 에러 := filepath.Abs(경로)
	if 에러 != nil {
		t.Fatal(에러)
	}

	return filepath.Clean(절대_경로)
}

// f테스트_종료_대기 : 프로세스가 소멸할 때까지 최대 3초 대기.
func f테스트_종료_대기(t *testing.T, 프로세스ID int) {
	t.Helper()

	for range 30 {
		if 프로세스, _ := ps.FindProcess(프로세스ID); 프로세스 == nil {
			return
		}

		time.Sleep(100 * time.Millisecond)
	}

	t.Fatalf("PID %v가 아직 종료되지 않았습니다.", 프로세스ID)
}

// f테스트_기다리기 : 조건이 참이 될 때까지 최대 2.5초 대기.
func f테스트_기다리기(t *testing.T, 설명 string, 조건 func() bool) {
	t.Helper()

	for range 50 {
		if 조건() {
			return
		}

		time.Sleep(50 * time.Millisecond)
	}

	t.Fatalf("시간 초과: %s", 설명)
}

// f테스트_파일_직접읽기 : pid_list.dat을 저장소 함수를 거치지 않고 gob로 직접 디코딩해 파일 포맷 자체를 검증한다.
func f테스트_파일_직접읽기(t *testing.T) (map[int]string, error) {
	t.Helper()

	파일, 에러 := os.Open(p프로세스ID_목록_파일명)
	if 에러 != nil {
		return nil, 에러
	}
	defer 파일.Close()

	매프 := make(map[int]string)
	return 매프, gob.NewDecoder(파일).Decode(&매프)
}

// TestI파일_초기읽기 : 파일이 없으면 자동 생성되고 빈 맵으로 읽혀야 한다.
func TestI파일_초기읽기(t *testing.T) {
	t격리(t)

	매프, 에러 := f프로세스ID_파일_읽기()
	if 에러 != nil {
		t.Fatalf("초기 파일 읽기 실패: %v", 에러)
	}

	if len(매프) != 0 {
		t.Fatalf("빈 상태에서는 빈 맵을 기대했으나 %d개 항목이 읽혔습니다.", len(매프))
	}

	if 존재함, _ := f프로세스ID_파일_존재함(); !존재함 {
		t.Fatal("읽기 시점에 파일이 자동 생성되어야 합니다.")
	}
}

// TestI추가_저장_형식 : S추가가 실행 화일 경로(절대경로 정규화)를 map[int]string으로 저장하고,
// 파일 포맷도 같은 모양이어야 한다. S제거는 항목을 삭제한다.
func TestI추가_저장_형식(t *testing.T) {
	t격리(t)

	테스트_경로 := filepath.Join(t.TempDir(), "테스트.exe")
	기대_값 := f테스트_정규화_경로(t, 테스트_경로)

	프로세스ID_저장소.S추가(1234, 테스트_경로)

	if 프로세스ID_저장소.저장소[1234] != 기대_값 {
		t.Fatalf("인메모리 저장값이 다릅니다. (기대=%s, 실제=%s)", 기대_값, 프로세스ID_저장소.저장소[1234])
	}

	if 매프, 에러 := f테스트_파일_직접읽기(t); 에러 != nil {
		t.Fatalf("pid_list.dat을 새 포맷(map[int]string)으로 직접 읽지 못했습니다: %v", 에러)
	} else if 매프[1234] != 기대_값 {
		t.Fatalf("파일 저장값이 다릅니다. (기대=%s, 실제=%s)", 기대_값, 매프[1234])
	}

	if 읽은_매프, 에러 := f프로세스ID_파일_읽기(); 에러 != nil {
		t.Fatalf("판독 함수로 파일 읽기 실패: %v", 에러)
	} else if 읽은_매프[1234] != 기대_값 {
		t.Fatalf("판독 함수 저장값이 다릅니다. (기대=%s, 실제=%s)", 기대_값, 읽은_매프[1234])
	}

	프로세스ID_저장소.S제거(1234)

	if 매프, 에러 := f테스트_파일_직접읽기(t); 에러 != nil {
		t.Fatalf("S제거 후 파일 읽기 실패: %v", 에러)
	} else if len(매프) != 0 {
		t.Fatalf("S제거 후에도 파일에 %d개 항목이 남았습니다.", len(매프))
	}
}

// TestI실행화일_동일함 : 소유권 검증 판정 로직을 결정적(프로세스 기동 없이)으로 확인한다.
// 핵심 불변식: 실제 값에 디렉터리가 있으면 전체 경로로만 비교 —
// 파일명이 같아도 다른 디렉터리의 실행 화일은 절대로 일치로 판정하지 않는다.
func TestI실행화일_동일함(t *testing.T) {
	경우_모음 := []struct {
		설명 string
		기대 string
		실제 string
		결과 bool
	}{
		{"전체 경로 완전 일치", `C:\Windows\System32\ping.exe`, `C:\Windows\System32\ping.exe`, true},
		{"전체 경로의 대소문자만 다름", `C:\WINDOWS\system32\ping.exe`, `c:\windows\SYSTEM32\PING.EXE`, true},
		{"같은 파일명이 다른 디렉터리 — 일치 불가", `C:\Windows\System32\ping.exe`, `D:\tools\ping.exe`, false},
		{"폴백: 기본 파일명 대문자(경로 없음)", `C:\WINDOWS\system32\ping.exe`, `PING.EXE`, true},
		{"폴백: 기본 파일명의 대소문자만 다름", `C:\Windows\System32\Ping.Exe`, `ping.exe`, true},
		{"기본 파일명도 다름 — 일치 불가", `C:\Windows\System32\python.exe`, `PING.EXE`, false},
		{"실제 값 빈 문자열(조회 실패) — 일치 불가", `C:\Windows\System32\ping.exe`, "", false},
	}

	for _, 경우 := range 경우_모음 {
		if 실제_결과 := f실행화일_동일함(경우.기대, 경우.실제); 실제_결과 != 경우.결과 {
			t.Errorf("[%s] 판정 불일치: 기대=%v 실제=%v (기대 경로=%q, 실제 값=%q)", 경우.설명, 경우.결과, 실제_결과, 경우.기대, 경우.실제)
		}
	}
}

// TestI정리_일치시_종료 : 실제 기동된 외부 프로세스(PID와 실행 화일이 일치)는 정상적으로 종료되어야 한다.
func TestI정리_일치시_종료(t *testing.T) {
	t격리(t)
	핑 := f테스트_핑_경로(t)

	pid, 에러 := F외부_프로세스_실행(핑, "-n", "60", "127.0.0.1")
	if 에러 != nil || pid <= 0 {
		t.Fatalf("외부 프로세스 기동 실패: %v", 에러)
	}
	t.Cleanup(func() { _ = F프로세스_종료by프로세스ID(pid) }) // 안전망: 테스트 실패 시에도 유령 프로세스 남기지 않기.

	// 기동 직후 S추가(파일 기록)는 별도 루틴에서 비동기로 수행되므로, 항목이 기록될 때까지 대기한다.
	기대_경로 := f테스트_정규화_경로(t, 핑)
	f테스트_기다리기(t, "기동된 프로세스의 경로가 pid_list.dat에 기록될 것", func() bool {
		매프, 에러 := f테스트_파일_직접읽기(t)
		if 에러 != nil {
			return false
		}

		return 매프[pid] == 기대_경로
	})

	수량, 에러 := f잔류_프로세스_정리_및_초기화()
	if 에러 != nil {
		t.Fatalf("정리 과정에서 에러: %v", 에러)
	}

	if 수량 != 1 {
		t.Fatalf("실행 화일이 일치하는 잔류 프로세스 1개를 종료할 것을 기대했으나 %d개입니다.", 수량)
	}

	f테스트_종료_대기(t, pid)

	// f외부_프로세스_생성 루틴의 Wait 완료 → S제거 처리가 끝날 때까지 대기 (다음 테스트로 유출 금지).
	f테스트_기다리기(t, "S제거(종료 처리)가 인메모리 저장소에서 완료될 것", func() bool {
		프로세스ID_저장소.Lock()
		개수 := len(프로세스ID_저장소.저장소)
		프로세스ID_저장소.Unlock()

		return 개수 == 0
	})
}

// TestI정리_불일치시_스킵 : 핵심 회귀 테스트. 무관 프로세스가 PID를 재사용한 경우(실행 화일 불일치)와
// 이미 종료된 PID는 어떤 경우에도 Kill하지 않고 스킵해야 한다.
func TestI정리_불일치시_스킵(t *testing.T) {
	t격리(t)

	자기_PID := os.Getpid()
	if 프로세스, _ := ps.FindProcess(자기_PID); 프로세스 == nil {
		t.Fatal("테스트 프로세스 자체의 핸들을 조회하지 못했습니다.")
	}

	// 시나리오 1: 무관 프로세스(= 이 테스트 프로세스)가 PID를 재사용한 상황 — 기대 경로와 실제 경로의 불일치.
	프로세스ID_저장소.S추가(자기_PID, filepath.Join(os.TempDir(), "존재하지", "않는.exe"))

	// 시나리오 2: 이미 종료된(존재하지 않는) PID.
	프로세스ID_저장소.S추가(99999999, filepath.Join(os.TempDir(), "아무", "exe.exe"))

	수량, 에러 := f잔류_프로세스_정리_및_초기화()

	if 수량 != 0 {
		t.Fatalf("검증에 실패한 프로세스가 종료되었습니다! (수량=%d) — PID 재사용 가드가 작동하지 않았습니다.", 수량)
	}

	if 에러 != nil {
		t.Fatalf("스킵은 정상 경로이므로 에러가 없어야 합니다: %v", 에러)
	}

	if 프로세스, _ := ps.FindProcess(자기_PID); 프로세스 == nil {
		t.Fatal("테스트 프로세스가 종료되었습니다 — 실행 화일이 불일치한 PID가 Kill된 것입니다.")
	}

	if 매프, 에러 := f테스트_파일_직접읽기(t); 에러 != nil {
		t.Fatalf("정리 후 파일 읽기 실패: %v", 에러)
	} else if len(매프) != 0 {
		t.Fatalf("정리 후 파일이 재초기화되어야 합니다. (남은 항목=%d)", len(매프))
	}
}

// TestI정리_대소문자_무시 : Windows 경로의 대소문자가 달라도 동일 실행 화일로 판정해 종료해야 한다.
func TestI정리_대소문자_무시(t *testing.T) {
	t격리(t)
	핑 := f테스트_핑_경로(t)

	pid, 에러 := F외부_프로세스_실행(핑, "-n", "60", "127.0.0.1")
	if 에러 != nil || pid <= 0 {
		t.Fatalf("외부 프로세스 기동 실패: %v", 에러)
	}
	t.Cleanup(func() { _ = F프로세스_종료by프로세스ID(pid) })

	var 기록_경로 string
	f테스트_기다리기(t, "기동된 프로세스의 경로가 pid_list.dat에 기록될 것", func() bool {
		매프, 에러 := f테스트_파일_직접읽기(t)
		if 에러 != nil {
			return false
		}

		기록_경로 = 매프[pid]
		return 기록_경로 != ""
	})

	// 저장 경로의 대소문자를 뒤집어 다시 기록 — EqualFold로 동일 경로가 되어야 한다.
	프로세스ID_저장소.S추가(pid, strings.ToLower(기록_경로))

	수량, _ := f잔류_프로세스_정리_및_초기화()
	if 수량 != 1 {
		t.Fatalf("대소문자가 다른 동일 경로도 종료 대상이 되어야 합니다. (수량=%d)", 수량)
	}

	f테스트_종료_대기(t, pid)

	// f외부_프로세스_생성 루틴의 Wait 완료 → S제거 처리가 끝날 때까지 대기 (다음 테스트로 유출 금지).
	f테스트_기다리기(t, "S제거(종료 처리)가 인메모리 저장소에서 완료될 것", func() bool {
		프로세스ID_저장소.Lock()
		개수 := len(프로세스ID_저장소.저장소)
		프로세스ID_저장소.Unlock()

		return 개수 == 0
	})
}

// TestI종료by프로세스ID_사망PID : 존재하지 않는 PID에는 에러를 반환하고 패닉해서는 안 된다(nil 가드 검증).
func TestI종료by프로세스ID_사망PID(t *testing.T) {
	t격리(t)

	명령어 := exec.Command("cmd.exe", "/c", "exit")
	if 에러 := 명령어.Start(); 에러 != nil {
		t.Fatal(에러)
	}

	사망_PID := 명령어.Process.Pid
	_ = 명령어.Wait() // 즉시 종료.

	if 에러 := F프로세스_종료by프로세스ID(사망_PID); 에러 == nil {
		t.Fatal("존재하지 않는 PID에 대한 종료 호출은 에러를 반환해야 합니다.")
	}
}

// TestI종료by프로세스_이름 : 이름 기반 종료의 판정이 대소문자를 무시하며,
// 파일명과 경로(경로는 파일명 부분만) 모두 입력으로 허용되어야 한다. 자기 자신은 절대 종료되지 않는다.
func TestI종료by프로세스_이름(t *testing.T) {
	t격리(t)
	핑 := f테스트_핑_경로(t)

	핑_기동 := func() int {
		명령어 := exec.Command(핑, "-n", "60", "127.0.0.1")
		if 에러 := 명령어.Start(); 에러 != nil {
			t.Fatalf("ping 기동 실패: %v", 에러)
		}

		pid := 명령어.Process.Pid
		t.Cleanup(func() { _ = F프로세스_종료by프로세스ID(pid) }) // 안전망.

		f테스트_기다리기(t, "기동된 ping이 프로세스 목록에 나타날 것", func() bool {
			프로세스, _ := ps.FindProcess(pid)
			return 프로세스 != nil
		})

		return pid
	}

	// 0) 빈 이름: 에러를 반환하고 아무 프로세스도 종료하지 않아야 한다.
	if 에러 := F프로세스_종료by프로세스_이름("   "); 에러 == nil {
		t.Fatal("빈(공백) 프로세스 이름은 에러를 반환해야 합니다.")
	}

	// 1) 아무 것과도 매칭되지 않는 이름: 자기 자신(테스트 프로세스)이 종료되어서는 안 된다.
	if 에러 := F프로세스_종료by프로세스_이름("존재하지_않는_파일.exe", false); 에러 != nil {
		t.Fatalf("매칭 없는 이름에는 에러가 없어야 합니다: %v", 에러)
	}

	// 2) 기동 중 프로세스의 파일명(대소문자 변형 입력): 종료되어야 한다.
	pid1 := 핑_기동()
	if 에러 := F프로세스_종료by프로세스_이름(strings.ToLower(filepath.Base(핑)), false); 에러 != nil {
		t.Fatalf("종료 호출 실패: %v", 에러)
	}
	f테스트_종료_대기(t, pid1)

	// 3) 전체 경로 입력(파일명 부분만 비교 대상): 종료되어야 한다.
	pid2 := 핑_기동()
	if 에러 := F프로세스_종료by프로세스_이름(핑, false); 에러 != nil {
		t.Fatalf("종료 호출 실패: %v", 에러)
	}
	f테스트_종료_대기(t, pid2)

	if 프로세스, _ := ps.FindProcess(os.Getpid()); 프로세스 == nil {
		t.Fatal("테스트 프로세스 자체가 종료되었습니다 — 이름 매칭 조건이 지나치게 넓습니다.")
	}
}

// TestI구형식_파일_자동초기화 : 구 형식 파일(값 = S비어있음 구조체)은 디코딩 실패로 표면화되고(침묵 금지),
// 이후 파일이 새 형식으로 자동 재초기화되어야 한다. (매일 재부팅 전제로 마이그레이션 코드는 두지 않는다.)
func TestI구형식_파일_자동초기화(t *testing.T) {
	t격리(t)

	구형식_매프 := map[int]lb.S비어있음{4287: {}}
	파일, 에러 := os.Create(p프로세스ID_목록_파일명)
	if 에러 != nil {
		t.Fatal(에러)
	}

	if 에러 = gob.NewEncoder(파일).Encode(구형식_매프); 에러 != nil {
		t.Fatalf("구 형식 파일 작성 실패: %v", 에러)
	} else if 에러 = 파일.Close(); 에러 != nil {
		t.Fatal(에러)
	}

	if _, 읽기_에러 := f잔류_프로세스_정리_및_초기화(); 읽기_에러 == nil {
		t.Fatal("구 형식 파일은 디코딩 실패가 표면화되어야 합니다 (침묵 처리 금지).")
	}

	if 매프, 에러 := f테스트_파일_직접읽기(t); 에러 != nil {
		t.Fatalf("재초기화된 파일을 새 형식으로 읽지 못했습니다: %v", 에러)
	} else if len(매프) != 0 {
		t.Fatalf("재초기화 후 파일이 비어 있어야 합니다. (남은 항목=%d)", len(매프))
	}
}
