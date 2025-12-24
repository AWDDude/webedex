package main

import (
	"flag"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/AWDDude/webedex/cmd/browsers/firefox"
)

func main() {
	debug := flag.Bool("debug", false, "sets log level to debug")
	ffProfPath := flag.String("firefox_prof_path", "", "base path to the target firefox profile")
	flag.Parse()

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	configPath, err := FindConfigPath()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to locate webedex config path")
	}

	var config WebedexConf
	config, err = GetConfig(configPath)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to load config")
	}
	
	

	if p, err := firefox.FindProfile(); ffProfPath == nil || *ffProfPath == "" {
		if err != nil {
			log.Fatal().Err(err).Msg("failed to locate firefox profile")
		}
		ffProfPath = &p
	}

	log.Info().Str("FirefoxProfilePath", *ffProfPath).Send()

}
