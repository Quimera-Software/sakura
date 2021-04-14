// Copyright (c) 2020. Quimera Software S.p.A.

package sakura

import (
	"errors"
	"io"
	"log"
	"os"
)

func LogInfo(msg string) {
	logToFile("[INFO] " + msg)
}

func BroadcastInfo(msg string, details string) {
	LogInfo(msg)

	sErr := Broadcast(msg, "INFO", errors.New(details))
	if sErr != nil {
		LogWarn("(SAKURA) Unable broadcast info")
		LogWarn("(SAKURA) " + sErr.Error())
	}
}

func LogWarn(msg string) {
	logToFile("[WARN] " + msg)
}

func BroadcastWarn(msg string, details string) {
	LogWarn(msg)

	sErr := Broadcast(msg, "WARN", errors.New(details))
	if sErr != nil {
		LogWarn("(SAKURA) Unable broadcast warning")
		LogWarn("(SAKURA) " + sErr.Error())
	}
}

func LogError(msg string) {
	logToFile("[ERROR] " + msg)
}

func BroadcastError(msg string, err error) {
	LogError(msg)

	sErr := Broadcast(msg, "ERROR", err)
	if sErr != nil {
		LogError("(SAKURA) Unable to broadcast error")
		LogError("(SAKURA) " + sErr.Error())
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
