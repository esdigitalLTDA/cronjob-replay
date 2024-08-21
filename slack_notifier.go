package main

import (
	"log"

	"github.com/nlopes/slack"
)

func sendSlackNotification(webhookURL, message string) {
	msg := slack.WebhookMessage{
		Text: message,
	}

	err := slack.PostWebhook(webhookURL, &msg)
	if err != nil {
		log.Printf("Error sending notification to Slack: %v", err)
	}
}
