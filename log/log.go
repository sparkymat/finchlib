package log

import (
	"fmt"
	"os"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

func Init() {
	logger.Formatter = &logrus.TextFormatter{FullTimestamp: true}
	logger.Out = os.Stdout
	logger.Level = logrus.DebugLevel
}

func prefixRequestId(c echo.Context, inString string) string {
	if c == nil {
		return inString
	} else {
		return fmt.Sprintf("[%v] %v", c.Request().Header.Get(echo.HeaderXRequestID), inString)
	}
}

func Debug(c echo.Context, message string) {
	logger.Debugln(prefixRequestId(c, message))
}

func Info(c echo.Context, message string) {
	logger.Infoln(prefixRequestId(c, message))
}

func Warn(c echo.Context, message string) {
	logger.Warnln(prefixRequestId(c, message))
}

func Error(c echo.Context, message string) {
	logger.Errorln(prefixRequestId(c, message))
}

func Fatal(c echo.Context, message string) {
	logger.Fatalln(prefixRequestId(c, message))
}
