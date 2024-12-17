package conf

import (
	"bytes"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"text/template"
)

func (conf Conf) templateFile(str string) (by []byte, err error) {
	ud := conf.getUserdataMap()
	buf := &bytes.Buffer{}
	templ, err := template.New("conf").Parse(str)
	if err == nil {
		templ.Execute(buf, map[string]interface{}{
			"confdir": filepath.Dir(conf.pabs(conf.FileName)),
			"CONFDIR": filepath.Dir(conf.pabs(conf.FileName)),
			"SELFDIR": filepath.Dir(conf.pabs(conf.FileName)),
			"selfdir": filepath.Dir(conf.pabs(conf.FileName)),
			"workdir": conf.pwd(),
			"WORKDIR": conf.pwd(),
			"home":    ud["home"],
			"HOME":    ud["home"],
			"uid":     ud["uid"],
			"UID":     ud["uid"],
			"gid":     ud["gid"],
			"GID":     ud["gid"],
			"user":    ud["username"],
			"USER":    ud["username"],
		})
		by = buf.Bytes()
	}
	return
}

func (conf Conf) getUserdataMap() map[string]string {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	m := make(map[string]string)
	m["home"] = user.HomeDir + "/"
	m["uid"] = user.Uid
	m["gid"] = user.Gid
	m["username"] = user.Username
	m["name"] = user.Name
	return m
}

func (conf Conf) pwd() string {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return pwd
}

func (conf Conf) pabs(pathstring string) string {
	r, err := filepath.Abs(pathstring)
	if err != nil {
		fmt.Printf("Unable to make absolute path. %s\n", err)
		os.Exit(1)
	}
	return r
}
