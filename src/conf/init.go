package conf

import (
	"time"

	"github.com/triole/logseal"
)

func Init(now time.Time, confFile string, lg logseal.Logseal) (conf Conf) {
	conf.Now = now
	conf.FileName = confFile
	conf.Lg = lg
	conf.Content = conf.readConf()
	return
}
