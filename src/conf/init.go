package conf

import (
	"path/filepath"
	"time"

	"github.com/triole/logseal"
)

func Init(now time.Time, confFile, opsFilter string, lg logseal.Logseal) (conf Conf) {
	conf.Now.Local = now
	conf.Now.UTC = now.UTC()
	confFile, err := filepath.Abs(confFile)
	lg.IfErrFatal(
		"unable to determine absolute path", logseal.F{"path": confFile, "error": err},
	)
	conf.FileName = confFile
	conf.Lg = lg
	conf.Content = conf.readConf(opsFilter)
	return
}
