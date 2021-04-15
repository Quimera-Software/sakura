// Copyright (c) 2020. Quimera Software S.p.A.

package sakura

import (
	"errors"
	"io"
	"log"
	"os"
)

const (
	LevelInfo = iota
	LevelWarn
	LevelError
	LevelFatal
)

func SetBroadcastLevel(lv int)  {
	cfg.BroadcastLevel = lv
}

func SetLoggingLevel(lv int)  {
	cfg.LogLevel = lv
}

func LogInfo(msg string) {
	if cfg.LogLevel >= LevelInfo {
		logToFile("[INFO] " + msg)
	}
}

func BroadcastInfo(msg string, details string) {
	if cfg.BroadcastLevel >= LevelInfo {
		LogInfo(msg)

		sErr := Broadcast(msg, "INFO", errors.New(details))
		if sErr != nil {
			LogWarn("(SAKURA) Unable broadcast info")
			LogWarn("(SAKURA) " + sErr.Error())
		}
	}
}

func LogWarn(msg string) {
	if cfg.LogLevel >= LevelWarn {
		logToFile("[WARN] " + msg)
	}
}

func BroadcastWarn(msg string, details string) {
	if cfg.BroadcastLevel >= LevelWarn {
		LogWarn(msg)

		sErr := Broadcast(msg, "WARN", errors.New(details))
		if sErr != nil {
			LogWarn("(SAKURA) Unable broadcast warning")
			LogWarn("(SAKURA) " + sErr.Error())
		}
	}
}

func LogError(msg string) {
	if cfg.LogLevel >= LevelError {
		logToFile("[ERROR] " + msg)
	}
}

func BroadcastError(msg string, err error) {
	if cfg.BroadcastLevel >= LevelError {
		LogError(msg)

		sErr := Broadcast(msg, "ERROR", err)
		if sErr != nil {
			LogError("(SAKURA) Unable to broadcast error")
			LogError("(SAKURA) " + sErr.Error())
		}
	}
}

func LogFatal(msg string) {
	logToFile("[FATAL] " + msg)
}

func BroadcastFatal(msg string, err error) {
	LogFatal(msg)

	sErr := Broadcast(msg, "FATAL", err)
	if sErr != nil {
		LogWarn("(SAKURA) Unable to broadcast fatal error")
		LogWarn("(SAKURA) " + sErr.Error())
	}
}

func logToFile(msg string) {
	log.Print(msg)
}

func newLogFile() error {
	file, err := os.Create(cfg.LogPath)
	if err != nil {
		return err
	}

	wrt := io.MultiWriter(os.Stdout, file)
	log.SetOutput(wrt)

	return nil
}
