package main

import (
	"log"
	"os"

	"github.com/slack-go/slack"
)

func sendSlackNotification(message string) {
	slackWebhookURL := os.Getenv("SLACK_WEBHOOK_URL")
	slackChannel := os.Getenv("SLACK_CHANNEL")

	if slackWebhookURL == "" || slackChannel == "" {
		log.Println("Slack webhook URL or channel not configured in environment variables")
		return
	}

	msg := slack.WebhookMessage{
		Channel: slackChannel,
		Text:    message,
	}

	err := slack.PostWebhook(slackWebhookURL, &msg)
	if err != nil {
		log.Printf("Error sending notification to Slack: %v", err)
	}
}
