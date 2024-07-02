package cmd

import (
	"testing"

	ipfslog "github.com/ipfs/go-log/v2"
)

func TestLogger(t *testing.T) {
	loggeeeer := ipfslog.Logger("rollappd")
	logger:=ipfslog.WithSkip(loggeeeer,1)
	ipfslog.SetAllLoggers(ipfslog.LevelInfo)
	myLogger := MyLogger{Logger: logger}
	myLogger.Info("test")
}
