package log

import (
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func SetLogLevel(cmd *cobra.Command, args []string) {
	logLevels := map[string]log.Level{"info": log.InfoLevel, "debug": log.DebugLevel}
	log.SetLevel(logLevels[strings.ToLower(cmd.Flag("log-level").Value.String())])
}

var logs = log.WithFields(log.Fields{
	"pivot": "0.0.1",
})

func Info(msg ...interface{}) {
	logs.Info(msg)

}

func Debug(msg ...interface{}) {
	logs.Debug(msg)
}

func Warn(msg ...interface{}) {
	logs.Warn(msg)
}

func Fatal(msg ...interface{}) {
	logs.Fatal(msg)
}

func Panic(msg ...interface{}) {
	logs.Panic(msg)
}

func GetLogLevel() log.Level {
	return log.GetLevel()
}

func Print(msg string, exit bool) {
	fmt.Println(msg)
	os.Exit(2)
}
