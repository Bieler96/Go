package Logger

import (
	"GO/File"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"text/tabwriter"
	"time"
)

var (
	colorReset  = "\033[0m"
	red         = "\033[31m"
	green       = "\033[32m"
	yellow      = "\033[33m"
	blue        = "\033[34m"
	purple      = "\033[35m"
	cyan        = "\033[36m"
	white       = "\033[37m"
	redBg       = string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
	greenBg     = string([]byte{27, 91, 57, 55, 59, 52, 50, 109})
	yellowBg    = string([]byte{27, 91, 57, 48, 59, 52, 51, 109})
	blueBg      = string([]byte{27, 91, 57, 55, 59, 52, 52, 109})
	purpleBg    = string([]byte{27, 91, 57, 55, 59, 52, 53, 109})
	cyanBg      = string([]byte{27, 91, 57, 55, 59, 52, 54, 109})
	whiteBg     = string([]byte{27, 91, 57, 48, 59, 52, 55, 109})
	debugMode   = true
	showTime    = false
	writeFile   = false
	timeFormat  = "01-02-2006 15:04:05"
	logFileName = "log.log"
)

func Debug(msg string, prefix ...string) {
	if debugMode {
		createMessage(cyan, cyanBg, "DEBUG", msg, prefix...)
	}
}

func Info(msg string, prefix ...string) {
	createMessage(white, whiteBg, "INFO", msg, prefix...)
}

func Notice(msg string, prefix ...string) {
	createMessage(green, greenBg, "NOTICE", msg, prefix...)
}

func Warning(msg string, prefix ...string) {
	createMessage(yellow, yellowBg, "WARNING", msg, prefix...)
}

func Error(msg string, prefix ...string) {
	createMessage(red, redBg, "ERROR", msg, prefix...)
}

func Critical(msg string, prefix ...string) {
	createMessage(purple, purpleBg, "CRITICAL", msg, prefix...)
}

func SetDebugMode(status bool) {
	debugMode = status
}

func fileAndLine() (string, int) {
	_, file, line, _ := runtime.Caller(3)
	return filepath.Base(file), line
}

func SetTimeFormat(format string) {
	timeFormat = format
}

func SetShowTime(status bool) {
	showTime = status
}

func SetWriteFile(status bool) {
	writeFile = status
}

func SetLogFileName(name string) {
	logFileName = name
}

func ClearLogFile() {
	File.ClearFile(logFileName)
}

func createMessage(color string, bgColor string, mode string, msg string, prefix ...string) {
	dt := time.Now().Format(timeFormat)
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 8, 8, 1, '\t', 0)

	defer w.Flush()

	state := mode
	if prefix != nil {
		state = prefix[0]
	}

	time := ""
	if showTime {
		time = dt
		//fmt.Fprintf(w, "%s%s\t%s\t%s\t%s\t%s\t \n", color, time, state, "▶", msg, colorReset)
		fmt.Fprintf(w, "%s%s\t%s", color, time, colorReset)
		fmt.Fprintf(w, "%s%s\t%s", bgColor, state, colorReset)
		fmt.Fprintf(w, "%s%s\t%s%s\n", color, "▶", msg, colorReset)
	} else {
		//fmt.Fprintf(w, "%s%s\t%s\t%s\t%s\t \n", color, state, "▶", msg, colorReset)
		fmt.Fprintf(w, "%s%s", color, colorReset)
		fmt.Fprintf(w, "%s%s\t%s", bgColor, state, colorReset)
		fmt.Fprintf(w, "%s%s\t%s%s\n", color, "▶", msg, colorReset)
	}

	if writeFile {
		File.WriteToFile(logFileName, "["+state+"] ▶ "+msg)
	}
}
