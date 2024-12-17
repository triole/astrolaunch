package conf

import (
	"os"
	"path/filepath"
	"regexp"

	"github.com/triole/logseal"
)

func (conf Conf) find(basedir string, rxFilter string) (filelist []string, err error) {
	filelist = []string{}
	inf, err := os.Stat(basedir)
	if err != nil {
		conf.Lg.Error(
			"unable to access folder", logseal.F{
				"path": basedir, "error": err,
			},
		)
		return
	}
	if !inf.IsDir() {
		conf.Lg.Error(
			"not a folder, please provide a directory to look for md files.",
			logseal.F{"path": basedir},
		)
		return
	}

	rxf, _ := regexp.Compile(rxFilter)
	err = filepath.Walk(basedir, func(path string, f os.FileInfo, err error) error {
		if rxf.MatchString(path) {
			inf, err := os.Stat(path)
			if err == nil {
				if !inf.IsDir() {
					filelist = append(filelist, path)
				}
			} else {
				conf.Lg.IfErrInfo("stat file failed", logseal.F{"path": path})
			}
		}
		return nil
	})
	conf.Lg.IfErrError("find files failed", logseal.F{"path": basedir, "error": err})
	return
}
