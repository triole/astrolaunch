package conf

import (
	"os"
	"time"

	"github.com/triole/logseal"
	yaml "gopkg.in/yaml.v3"
)

func (conf *Conf) ReadConf() {
	by, err := os.ReadFile(conf.FileName)
	conf.Lg.IfErrFatal(
		"can not read file", logseal.F{"path": conf.FileName, "error": err},
	)
	by, err = conf.templateFile(string(by))
	conf.Lg.IfErrFatal(
		"can not expand config variables", logseal.F{"path": conf.FileName, "error": err},
	)
	err = yaml.Unmarshal(by, &conf.Content)
	conf.Lg.IfErrFatal(
		"can not unmarshal config", logseal.F{"path": conf.FileName, "error": err},
	)
	conf.Content.OpsList, err = conf.find(conf.Content.OpsDir, conf.OpsFilter+".*\\.yam?l$")
	conf.Lg.IfErrFatal(
		"find operations failed",
		logseal.F{
			"path": conf.FileName, "opsdir": conf.Content.OpsDir, "error": err,
		},
	)
}

func (conf *Conf) ReadOps() {
	for _, el := range conf.Content.OpsList {
		op := conf.readOp(el)
		if op.Range.Pre == "" {
			op.Range.Pre = conf.Content.DefaultRange.Pre
		}
		if op.Range.Post == "" {
			op.Range.Post = conf.Content.DefaultRange.Post
		}
		conf.Content.Operations = append(conf.Content.Operations, op)
	}
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

func (conf *Conf) SetNow(tim time.Time) {
	conf.Now.Local = tim
	conf.Now.UTC = tim.UTC()
}
