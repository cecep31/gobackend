package utils

import (
	"crypto/rand"
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/oklog/ulid/v2"
)

func GenerateRandomFilename(originalFilename string) string {
	// Extract extension from the original filename
	extension := filepath.Ext(originalFilename)

	// Remove extension from the original filename
	filenameWithoutExtension := strings.TrimSuffix(originalFilename, extension)

	// Generate timestamp
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)

	// Combine the original filename (without extension), timestamp, and extension
	return fmt.Sprintf("%s%d%s", filenameWithoutExtension, timestamp, extension)
}

func GenerateRandomString(length int, filname string) (string, error) {
	generatedulit := ulid.Make()
	extention := filepath.Ext(filname)
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	// timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	return fmt.Sprintf("%s%s", generatedulit, extention), nil
}
