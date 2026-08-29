package cli

import (
	"fmt"
	"io"
	"os"
	"time"
)

type messageLevel int
type messageVisibility int

const (
	levelDebug messageLevel = iota
	levelInfo
	levelSuccess
	levelWarning
	levelError
)

const (
	visibilityNormal messageVisibility = iota
	visibilityVerbose
	visibilityDebug
)

const (
	colorReset  = "\033[0m"
	colorGray   = "\033[90m"
	colorBlue   = "\033[34m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorRed    = "\033[31m"
)

var outputOptions OutputOptions

func ConfigureOutput(options OutputOptions) {
	outputOptions = options
}

func Direct(format string, args ...any) {
	if outputOptions.Silent {
		return
	}
	fmt.Fprintf(os.Stdout, format, args...)
}

func Debug(format string, args ...any) {
	writeVisibleMessage(os.Stdout, visibilityDebug, levelDebug, format, args...)
}

func Verbose(format string, args ...any) {
	writeVisibleMessage(os.Stdout, visibilityVerbose, levelInfo, format, args...)
}

func Info(format string, args ...any) {
	writeVisibleMessage(os.Stdout, visibilityNormal, levelInfo, format, args...)
}

func Success(format string, args ...any) {
	writeVisibleMessage(os.Stdout, visibilityNormal, levelSuccess, format, args...)
}

func Warning(format string, args ...any) {
	writeVisibleMessage(os.Stderr, visibilityNormal, levelWarning, format, args...)
}

func Error(format string, args ...any) {
	writeVisibleMessage(os.Stderr, visibilityNormal, levelError, format, args...)
}

func writeVisibleMessage(writer io.Writer, visibility messageVisibility, level messageLevel, format string, args ...any) {
	if !messageVisible(visibility) {
		return
	}
	fmt.Fprint(writer, formatMessage(time.Now(), level, format, args...))
}

func messageVisible(visibility messageVisibility) bool {
	if outputOptions.Silent {
		return false
	}

	switch visibility {
	case visibilityNormal:
		return true
	case visibilityVerbose:
		return outputOptions.Verbose
	case visibilityDebug:
		return outputOptions.Debug
	default:
		return false
	}
}

func formatMessage(now time.Time, level messageLevel, format string, args ...any) string {
	message := fmt.Sprintf(format, args...)
	return fmt.Sprintf("%s - %s%s%s\n", now.Format("15:04:05"), messageColor(level), message, colorReset)
}

func messageColor(level messageLevel) string {
	switch level {
	case levelDebug:
		return colorGray
	case levelInfo:
		return colorBlue
	case levelSuccess:
		return colorGreen
	case levelWarning:
		return colorYellow
	case levelError:
		return colorRed
	default:
		return colorReset
	}
}
