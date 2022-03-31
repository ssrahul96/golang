package utils

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func GetCurrentPath() string {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return pwd
}

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
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
