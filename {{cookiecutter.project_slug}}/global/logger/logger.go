// package logger
package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var ins = logrus.New()

func init() {
	ins.Out = os.Stdout
}

func GetInstance() *logrus.Logger {
	return ins
}
