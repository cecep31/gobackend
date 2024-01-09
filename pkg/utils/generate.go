package utils

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"
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
