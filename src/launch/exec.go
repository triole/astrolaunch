package launch

import (
	"bytes"
	"io"
	"os"
	"os/exec"
	"syscall"

	"github.com/triole/logseal"
)

func (la Launch) execute(cmds [][]string) (output []byte, exitcode int, err error) {
	for _, cmdArr := range cmds {
		if !la.Conf.DryRun {
			output, exitcode, err = la.runCmd(cmdArr)
		}
	}
	return output, exitcode, err
}

func (la Launch) runCmd(cmdArr []string) ([]byte, int, error) {
	var err error
	var exitcode int
	var stdBuffer bytes.Buffer

	cmd := exec.Command(cmdArr[0], cmdArr[1:]...)
	// mw := io.MultiWriter(&stdBuffer)
	mw := io.MultiWriter(os.Stdout, &stdBuffer)

	cmd.Stdout = mw
	cmd.Stderr = mw
	if err = cmd.Run(); err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {
			// the program has exited with an exit code != 0
			if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
				exitcode = status.ExitStatus()
			}
		}
	}
	by := stdBuffer.Bytes()
	if err != nil {
		la.Lg.IfErrError(
			"exec failed",
			logseal.F{"cmd": cmdArr, "error": err, "output": string(by)},
		)
	} else {
		la.Lg.Debug(
			"exec successful",
			logseal.F{"cmd": cmdArr, "error": err, "output": string(by)},
		)
	}
	return by, exitcode, err
}
