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
	by, err = conf.templateFile(string(by))
	conf.Lg.IfErrFatal(
		"can not expand config variables", logseal.F{"path": conf.FileName, "error": err},
	)
	err = yaml.Unmarshal(by, &content)
	conf.Lg.IfErrFatal(
		"can not unmarshal config", logseal.F{"path": conf.FileName, "error": err},
	)
	content.OpsList, err = conf.find(content.OpsDir, `\.yam?l$`)
	conf.Lg.IfErrFatal(
		"find operations failed", logseal.F{"path": conf.FileName, "opsdir": content.OpsDir, "error": err},
	)
	for _, el := range content.OpsList {
		content.Operations = append(content.Operations, conf.readOp(el))
	}
	return content
}

func (conf Conf) readOp(fn string) (op Operation) {
	by, err := os.ReadFile(fn)
	conf.Lg.IfErrFatal(
		"can not read file", logseal.F{"path": conf.FileName, "error": err},
	)
	by, err = conf.templateFile(string(by))
	conf.Lg.IfErrFatal(
		"can not expand config variables", logseal.F{"path": conf.FileName, "error": err},
	)
	err = yaml.Unmarshal(by, &op)
	conf.Lg.IfErrFatal(
		"can not unmarshal config", logseal.F{"path": conf.FileName, "error": err},
	)
	return
}
