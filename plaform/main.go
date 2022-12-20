package main

import (
	"platform/config"
	"platform/logging"
	"platform/services"
)

func writeMessage(logger logging.Logger, cfg config.Configuration) {
	section, ok := cfg.GetSection("main")
	if ok {
		message, ok := section.GetString("message")
		if ok {
			logger.Info(message)
		} else {
			logger.Panicf("Can't find configuration setting")
		}
	} else {
		logger.Panic("Config section not found")
	}
}

func main() {
	//var cfg config.Configuration
	//var err error
	//cfg, err = config.Load("config.json")
	//if err != nil {
	//	panic(err)
	//}
	//var logger logging.Logger = logging.NewDefaultLogger(cfg)

	services.RegisterDefaultServices()

	//var cfg config.Configuration
	//services.GetService(&cfg)
	//
	//var logger logging.Logger
	//services.GetService(&logger)
	//
	//writeMessage(logger, cfg)

	services.Call(writeMessage)
}
