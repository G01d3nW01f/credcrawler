package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"github.com/fatih/color"
)

func findFilesWithKeywords(directory string, filenames []string, keywords []string) {
	_ = filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if info != nil && !info.IsDir() {
			for _, filename := range filenames {
				if strings.EqualFold(info.Name(), filename) {
					if fileContainsKeywords(path, keywords) {
						
						message := "[+]SUSPICIOUS: "+path
						coloredMessage := color.RedString(message)
						fmt.Println(coloredMessage)
						color.Unset()
						break
					}	
				}
			}
		}
		return nil
	})
}

func fileContainsKeywords(filepath string, keywords []string) bool {
	file, err := os.Open(filepath)
	if err != nil {
		return false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		for _, keyword := range keywords {
			if strings.Contains(strings.ToLower(line), strings.ToLower(keyword)) {
				return true
			}
		}
	}
	return false
}

func main() {
	directory := "/" // 検索を開始するディレクトリ
	filenames := []string{
		".env",
		"config.json",
		"database.yml",
		"settings.py",
		"config.php",
		"appsettings.json",
		"config.xml",
		".htpasswd",
	}
	keywords := []string{"pass", "user", "password"}
	findFilesWithKeywords(directory, filenames, keywords)
}

