package internal

import (
	"os"
	"time"

	logger "github.com/charmbracelet/log"
)

var (
	log = logger.NewWithOptions(os.Stderr, logger.Options{
		ReportCaller:    false,
		ReportTimestamp: false,
		TimeFormat:      time.Kitchen,
	})
)
