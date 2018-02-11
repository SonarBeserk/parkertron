package main

import (
	"flag"
	"strings"
)

var (
	//BotID is the Discord Bot ID
	BotID string
	//ShowConfig is part of the init process
	ShowConfig string
	//response is the bot response on the channel
	response string
)

func init() {
	verbose := flag.String("v", "info", "set the console verbosity of the bot")
	flag.Parse()
	if *verbose == "debug" {
		setLogLevel("debug")
	} else {
		setLogLevel("info")
	}

	setupLogger()

	setConfig()

	writeLog("debug", "bot prefix is "+getConfig("prefix"), nil)
	writeLog("debug", "services loaded are "+getConfig("services"), nil)
}

func main() {
	if strings.Contains(getConfig("services"), "discord") == true {
		writeLog("info", "Starting discord connector", nil)
		go startDiscordConnection()
	}

	writeLog("info", "Bot is now running.  Press CTRL-C to exit.", nil)
	// Simple way to keep program running until CTRL-C is pressed.
	<-make(chan struct{})
	return
}
