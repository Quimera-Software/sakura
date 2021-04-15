package sakura

import (
	"fmt"
)

func Broadcast(msg string, reportType string, args ...string) error {
	var errMsg string

	if len(args) == 0 {
		switch cfg.Language {
		case LanguageEsp:
			errMsg = "La aplicación no ha proporcionado más información"
		default:
			errMsg = "The application provided no further information"
		}
	} else {
		for _, arg := range args {
			errMsg += arg + ". "
		}
	}

	if cfg.Discord.Enabled {
		dMsg := &discordMessage{}

		err := dMsg.fill()
		if err != nil {
			return err
		}

		dMsg.mentions = cfg.Discord.Mentions

		switch cfg.Language {
		case LanguageEsp:
			dMsg.Message = fmt.Sprintf("La aplicación '%s' ha reportado un error", cfg.AppName)
		default:
			dMsg.Message = fmt.Sprintf("The application '%s' has reported an error", cfg.AppName)
		}

		embeds := &discordEmbeds{}

		embeds.fill()
		embeds.Title = "[" + reportType + "] " + cfg.AppName + ": " + msg
		embeds.Description = errMsg

		err = dMsg.sendWithEmbeds(embeds)
		if err != nil {
			return err
		}
	}

	if cfg.Telegram.Enabled {
		err := telegramBroadcast("[" + reportType + "] " + cfg.AppName + ": " + msg + "\n" + errMsg)
		if err != nil {
			return err
		}
	}

	return nil
}

func BroadcastInfo(msg string, args ...string) {
	err := Broadcast(msg, "INFO", args...)
	if err != nil {
		fmt.Println("[ERROR] (SAKURA) Unable broadcast info:", err.Error())
		fmt.Println("[ERROR] (SAKURA) Unable broadcast info")
	}
}

func BroadcastWarn(msg string, args ...string) {
	err := Broadcast(msg, "WARN", args...)
	if err != nil {
		fmt.Println("[ERROR] (SAKURA) Unable broadcast warning:", err.Error())
		fmt.Println("[ERROR] (SAKURA) Unable broadcast warning")
	}
}

func BroadcastError(msg string, args ...string) {
	err := Broadcast(msg, "ERROR", args...)
	if err != nil {
		fmt.Println("[ERROR] (SAKURA) Unable broadcast error:", err.Error())
		fmt.Println("[ERROR] (SAKURA) Unable broadcast error")
	}
}

func BroadcastFatal(msg string, args ...string) {
	err := Broadcast(msg, "FATAL", args...)
	if err != nil {
		fmt.Println("[ERROR] (SAKURA) Unable broadcast fatal error:", err.Error())
		fmt.Println("[ERROR] (SAKURA) Unable broadcast fatal error")
	}
}
