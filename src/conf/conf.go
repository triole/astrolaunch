package conf

import (
	"os"

	"github.com/triole/logseal"
	yaml "gopkg.in/yaml.v3"
)

func (conf Conf) readConf() (content ConfContent) {
	by, err := os.ReadFile(conf.FileName)
	conf.Lg.IfErrFatal(
		"can not read file", logseal.F{"path": conf.FileName, "error": err},
	)
	err = yaml.Unmarshal(by, &content)
	conf.Lg.IfErrFatal(
		"can not unmarshal config", logseal.F{"path": conf.FileName, "error": err},
	)
	return content
}
