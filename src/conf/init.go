package conf

import (
	"time"

	"github.com/triole/logseal"
)

func Init(confFile string, lg logseal.Logseal) (conf Conf) {
	conf.Now = time.Now()
	conf.FileName = confFile
	conf.Lg = lg
	conf.Content = conf.readConf()
	return
}
