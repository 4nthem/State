package main

import (
	rest "github.com/4nthem/State/rest"
	log "github.com/cihub/seelog"
)

func main() {
	defer log.Flush()

	logger, err := log.LoggerFromConfigAsFile("seelog.xml")

	if err != nil {
		log.Warn("Failed to load logger config", err)
	} else {
		log.ReplaceLogger(logger)
	}

	rest.StartServer()

}
