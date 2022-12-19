package bt_factor_common

import (
	"github.com/ghts/ghts/lib"
	"path"
)

func DB파일명() string {
	return path.Join(lib.F홈_디렉토리(), "backtest_factor.dat")
}
