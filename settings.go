package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

const xdgconfig = "XDG_CONFIG_HOME"
const settingsFileName string = "tiptap-settings.json"
const settingsSubfolder string = "tiptap/"
const defaultLocation = "~/.config/" + settingsSubfolder
const usersettingsPayload string = `
{
    "wordcount":50,
    "style":{
        "background":"#558B6A",
        "errorbackground":"#ED6B86",
        "cursor":"#D58936",
        "text":"#F7B2B7",
        "correct":"#DE639A",
        "error":"#7F2982"
    }
}
`

type SettingsFile struct {
	WordCount int       `json:"wordcount"`
	Style     StyleJson `json:"style"`
}

type StyleJson struct {
	Background      string `json:"background"`
	Errorbackground string `json:"errorbackground"`
	Cursor          string `json:"cursor"`
	Text            string `json:"text"`
	Correct         string `json:"correct"`
	Err             string `json:"error"`
}

func SaveUserSettings(settingsFile SettingsFile) {
	location := determineFilepath()
	ensureSettingsFileExists(location)

	settingBytes, err := json.MarshalIndent(settingsFile, "", "    ")
	if err != nil {
		panic(fmt.Sprintf("Unable to save settings to file %v : %v", location, err))
	}

	err = os.WriteFile(location, settingBytes, 0644)
	if err != nil {
		panic(fmt.Sprintf("Unable to save settings to file %v : %v", location, err))
	}
}

func ensureSettingsFileExists(location string) {
	//if file/folder does not exist, attempt to create it.
	if _, err := os.Stat(location); err != nil {
		err := os.MkdirAll(strings.TrimSuffix(location, settingsFileName), 0644)
		if err != nil {
			panic(fmt.Sprintf("Unable to write user settings file: %v :%v", location, err))
		}

		err = os.WriteFile(location, []byte(usersettingsPayload), 0644)
		if err != nil {
			panic(fmt.Sprintf("Unable to write user settings file: %v :%v", location, err))
		}
	}
}

func LoadUserSettings() SettingsFile {
	location := determineFilepath()
	ensureSettingsFileExists(location)

	// read information from json file
	locationFile, err := os.Open(location)
	if err != nil {
		panic(fmt.Sprintf("Could not read settings file at location %v : %v", location, err))
	}
	settingsBytes, err := io.ReadAll(locationFile)
	if err != nil {
		panic(fmt.Sprintf("Could not read settings file at location %v : %v", location, err))
	}

	//unmarshall
	var settingsFile SettingsFile
	err = json.Unmarshal(settingsBytes, &settingsFile)
	if err != nil {
		panic(fmt.Sprintf("Could not read settings file at location %v : %v", location, err))
	}

	return settingsFile
}

func determineFilepath() string {
	loc, present := os.LookupEnv(xdgconfig)
	if present {
		return filepath.Clean(loc + string(os.PathSeparator) + settingsSubfolder + string(os.PathSeparator) + settingsFileName)
	} else {
		return filepath.Clean(fixPath(defaultLocation) + string(os.PathSeparator) + settingsFileName)
	}
}

func fixPath(path string) string {
	if path == "~" {
		usr, _ := user.Current()
		dir := usr.HomeDir
		path = dir
	} else if strings.HasPrefix(path, "~/") {
		usr, _ := user.Current()
		dir := usr.HomeDir
		path = filepath.Join(dir, path[2:])
	}
	return path
}
