package main

import (
	log "github.com/sirupsen/logrus"
)

// init initializes the package and sets up the Logrus formatter for Windows logging.
func init() {
	// fixes Logrus windows logging
	//log.SetFormatter(&log.TextFormatter{ForceColors: true})
	log.SetFormatter(&log.TextFormatter{
		DisableColors:    true,
		DisableTimestamp: true,
		PadLevelText:     true,
	})
	//	log.SetOutput(ansicolor.NewAnsiColorWriter(os.Stdout))
}
