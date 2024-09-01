package conf

import (
	"astrolaunch/src/astro"
	"time"

	"github.com/triole/logseal"
)

func Init(confFile string, lg logseal.Logseal) (conf Conf) {
	conf.Now = time.Now()
	conf.FileName = confFile
	conf.Lg = lg
	conf.Content = conf.readConf()
	conf.Astro = astro.Init(
		conf.Now, conf.Content.Location.Lat, conf.Content.Location.Lon,
	)
	return
}
