package main

import (
	"GO/Logger"
)

func main() {
	Logger.SetShowTime(true)
	Logger.SetTimeFormat("01-02-2006")

	Logger.Debug("test")
	Logger.Info("test")
	Logger.Notice("test")
	Logger.Warning("test")
	Logger.Error("test")
	Logger.Critical("test")
}
