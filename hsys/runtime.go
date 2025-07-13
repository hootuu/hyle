package hsys

import (
	"errors"
	"github.com/hootuu/hyle/hcfg"
	"github.com/rs/xid"
	"os"
	"strings"
)

var gServerID string
var gRunMode Mode
var gWorkingDirectory string

func RunMode() Mode {
	return gRunMode
}

func ServerID() string {
	return gServerID
}

func WorkingDirectory() string {
	return gWorkingDirectory
}

func Exit(err error) {
	if err != nil {
		Error("Crash error: ", err.Error())
	}
	os.Exit(-119)
}

func init() {
	gServerID = strings.ToUpper(xid.New().String())
	gRunMode = ModeOf(hcfg.GetString(SysRunModeCfg, localName))
	wd, nErr := os.Getwd()
	if nErr != nil {
		Error("Get Current Working Directory Failed: ", nErr.Error())
		Exit(errors.New("get current working directory failed"))
		return
	}
	gWorkingDirectory = wd
	Success("# Server ID: ", ServerID)
	Success("# Run Mode: ", strings.ToUpper(gRunMode.String()))
	Success("# Working Directory: ", WorkingDirectory)
}
