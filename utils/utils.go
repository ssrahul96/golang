package utils

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

func GetCurrentPath() string {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return pwd
}

func FormatJson(class interface{}) string {
	s, _ := json.MarshalIndent(class, "", "\t")
	return string(s)
}

func FormatYaml(class interface{}) string {
	yamlData, _ := yaml.Marshal(&class)
	return string(yamlData)
}

func FileExists(configPath string) bool {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return false
	}

	return true
}

func GetFiles(root string, ext string) []string {
	var a []string
	filepath.WalkDir(root, func(s string, d fs.DirEntry, e error) error {
		if e != nil {
			return e
		}
		if filepath.Ext(d.Name()) == ext {
			a = append(a, s)
		}
		return nil
	})
	return a
}

func DeleteFile(path string) {
	if FileExists(path) {
		os.Remove(path)
	}
}
