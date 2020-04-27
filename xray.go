package xray_loglevel

import (
	"github.com/aws/aws-xray-sdk-go/xray"
	"os"
)

const (
	XRayLogLevelEnv = "XRAY_LOG_LEVEL"
	LogLevelTrace   = "trace"
)

func Configure() error {
	logLevel := os.Getenv(XRayLogLevelEnv)
	if len(logLevel) == 0 {
		logLevel = LogLevelTrace
	}
	return xray.Configure(xray.Config{LogLevel: logLevel})
}
