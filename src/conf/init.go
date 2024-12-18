package conf

import (
	"path/filepath"
	"time"

	"github.com/triole/logseal"
)

func Init(now time.Time, confFile, opsFilter string, dryrun bool, lg logseal.Logseal) (conf Conf) {
	conf.SetNow(now)
	confFile, err := filepath.Abs(confFile)
	lg.IfErrFatal(
		"unable to determine absolute path", logseal.F{"path": confFile, "error": err},
	)
	conf.FileName = confFile
	conf.Lg = lg
	conf.DryRun = dryrun
	return
}
