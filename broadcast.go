package sakura

import "fmt"

func Broadcast(msg string, reportType string, reportError error) error {
	var errMsg string

	if reportError == nil {
		switch cfg.Language {
		case LanguageEsp:
			errMsg = "La aplicación no ha proporcionado más información"
		default:
			errMsg = "The application provided no further information"
		}
	} else {
		errMsg = reportError.Error()
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
