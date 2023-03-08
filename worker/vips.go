package worker

import (
	"github.com/davidbyttow/govips/v2/vips"
	l "pc-ziegert.de/media_service/common/log"
)

func VipsLogger(messageDomain string, verbosity vips.LogLevel, message string) {
	switch verbosity {
	case vips.LogLevelError:
		l.Errorf("%v: %v", messageDomain, message)
	case vips.LogLevelCritical:
		l.Errorf("%v: %v", messageDomain, message)
	case vips.LogLevelWarning:
		l.Warnf("%v: %v", messageDomain, message)
	case vips.LogLevelMessage:
		l.Verbf("%v: %v", messageDomain, message)
	case vips.LogLevelInfo:
		l.Infof("%v: %v", messageDomain, message)
	case vips.LogLevelDebug:
		l.Debugf("%v: %v", messageDomain, message)
	}
}

func GetVipsLogLevel(level string) vips.LogLevel {
	switch level {
	case "verbose":
		return vips.LogLevelMessage
	case "debug":
		return vips.LogLevelDebug
	case "info":
		return vips.LogLevelInfo
	case "warning":
		return vips.LogLevelWarning
	case "error":
		return vips.LogLevelError
	case "fatal":
		return vips.LogLevelCritical
	case "off":
		return vips.LogLevelCritical
	default:
		l.Fatalf("Invalid logging level '%s'!", level)
		return vips.LogLevelCritical
	}
}
