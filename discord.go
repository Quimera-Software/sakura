// Copyright (c) 2020. Quimera Software S.p.A.

package sakura

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/fatih/structs"
	"github.com/json-iterator/go"
)

const discColor = 16482710

type discordMessage struct {
	Message    string `structs:"content"`
	mentions   []int64
	TTS        bool   `structs:"tts"`
	Username   string `structs:"username"`
	AvatarURL  string `structs:"avatar_url"`
	webhookURL string
}

type discordEmbeds struct {
	Color       int    `json:"color"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Timestamp   string `json:"timestamp"`
}


func (dMsg discordMessage) send() error {
	mentions := parseMentions(dMsg.mentions)
	dMsg.Message = mentions + dMsg.Message

	body, err := jsoniter.Marshal(dMsg)
	if err != nil{
		return err
	}

	bodyReader := bytes.NewBuffer(body)

	resp, err := http.Post(dMsg.webhookURL, "application/json", bodyReader)
	if err != nil{
		return err
	}

	if resp.StatusCode != http.StatusNoContent{
		return errors.New(fmt.Sprintf("remote returned error: %s", resp.Status))
	}

	return nil
}

func (dMsg discordMessage) sendWithEmbeds(embeds *discordEmbeds) error {
	mentions := parseMentions(dMsg.mentions)
	dMsg.Message = mentions + dMsg.Message

	body := structs.Map(dMsg)
	body["embeds"] = []*discordEmbeds{embeds}

	bodyJson, err := jsoniter.Marshal(body)
	if err != nil{
		return err
	}

	bodyReader := bytes.NewBuffer(bodyJson)

	resp, err := http.Post(dMsg.webhookURL, "application/json", bodyReader)
	if err != nil{
		return err
	}

	if resp.StatusCode != http.StatusNoContent{
		return errors.New(fmt.Sprintf("remote returned error: %s", resp.Status))
	}

	return nil
}

func (dMsg *discordMessage) fill() error {
	dMsg.Username = cfg.Discord.Username
	dMsg.AvatarURL = cfg.Discord.AvatarURL
	dMsg.TTS = cfg.Discord.UseTTS

	dMsg.webhookURL = cfg.Discord.Webhook
	if dMsg.webhookURL == "" {
		return errors.New("no discord webhook URL set")
	}

	return nil
}

func (dEmb *discordEmbeds) fill() {
	dEmb.Color = discColor
	dEmb.Timestamp = time.Now().Format(time.RFC3339)
}

func parseMentions(mentions []int64) string {
	var mentionsString string
	for _, val := range mentions{
		mentionsString +=  fmt.Sprintf("<@&%d> ", val)
	}

	return mentionsString
}