package main

import "os"

type WebedexConf struct {
	BookMarks []Bookmark
}

type Bookmark struct {
	Kind     string
	Name     string
	Url      string
	Children []Bookmark
}

// Returns the configuration 
func GetConfig(configPath string) (config WebedexConf, err error) {

	return
}

// Returns the path to the config file as a string, it will respect $XDG_CONFIG_HOME
// if availible otherwise it will default to $HOME/.config/webedex/config.yaml
func FindConfigPath() (path string, err error) {
	path, err = os.UserConfigDir()

	path += "/webedex/config.yaml"
	return
}
