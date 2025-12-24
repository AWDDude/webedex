package firefox

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"

	"gopkg.in/ini.v1"
)

// FindProfile finds the base path to the main Firefox profile.
// Returns a string path to the base of the main Firefox profile and an error if it was unable to find a suitable profile.
func FindProfile() (profilePath string, err error) {
	profilesDir, err := getProfilesDir()
	if err != nil {
		return
	}

	profilesIni := filepath.Join(profilesDir, "profiles.ini")
	cfg, err := ini.Load(profilesIni)
	if err != nil {
		err = errors.New("could not open profiles.ini: " + err.Error())
		return
	}

	for _, section := range cfg.Sections() {
		if section.Key("Default").String() != "1" {
			continue
		}

		path := section.Key("Path").String()
		if path == "" {
			continue
		}

		isRelative := section.Key("IsRelative").MustInt(1) == 1
		profilePath = resolvePath(profilesDir, path, isRelative)

		if _, statErr := os.Stat(profilePath); os.IsNotExist(statErr) {
			err = errors.New("profile directory does not exist: " + profilePath)
			return
		}

		return
	}

	err = errors.New("no default profile found in profiles.ini")
	return
}

func getProfilesDir() (string, error) {
	switch runtime.GOOS {
	case "darwin":
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		return filepath.Join(home, "Library", "Application Support", "Firefox"), nil
	case "linux":
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		return filepath.Join(home, ".mozilla", "firefox"), nil
	case "windows":
		appData := os.Getenv("APPDATA")
		if appData == "" {
			return "", errors.New("APPDATA environment variable not set")
		}
		return filepath.Join(appData, "Mozilla", "Firefox"), nil
	default:
		return "", errors.New("unsupported operating system: " + runtime.GOOS)
	}
}

func resolvePath(profilesDir, path string, isRelative bool) string {
	if isRelative {
		return filepath.Join(profilesDir, path)
	}
	return path
}
