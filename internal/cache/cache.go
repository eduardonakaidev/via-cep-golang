package cache

import (
	"log"
	"os"
	"strings"
	"time"
)

const cacheTime = 500 // Cache duration in seconds

// GetFromCache retrieves data from cache if available and valid.
func GetFromCache(id string) string {
	cacheFile := getCacheFilename(id)
	info, err := os.Stat(cacheFile)
	if err != nil || time.Since(info.ModTime()) > cacheTime*time.Second {
		return ""
	}

	content, err := os.ReadFile(cacheFile)
	if err != nil {
		return ""
	}

	return string(content)
}

// SaveOnCache saves data to a cache file.
func SaveOnCache(id, content string) string {
	cacheFile := getCacheFilename(id)
	err := os.WriteFile(cacheFile, []byte(content), 0644)
	if err != nil {
		log.Printf("Error saving cache: %v", err)
		return ""
	}

	return content
}

func getCacheFilename(id string) string {
	return os.TempDir() + "/cep_" + strings.ReplaceAll(id, "-", "")
}
