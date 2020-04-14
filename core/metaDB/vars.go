package metaDB

import (
	"github.com/ZeroTechh/VelocityCore/logger"
	"github.com/ZeroTechh/hades"
)

var (
	// all the configs
	config = hades.GetConfig(
		"main.yaml",
		[]string{"config", "../config", "../../config"},
	)
	dbConfig = config.Map("database")

	log = logger.GetLogger(
		config.Map("service").Str("logFile"),
		config.Map("service").Bool("debug"),
	)
)
