package Utils

import (
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// SaveBase64ToFile extracts a base64 string and saves it to a file.
// It returns the relative path/URL to the saved file or an error.
func SaveBase64ToFile(base64String, uploadDir string) (string, error) {
	// Parse data URI if present (e.g. data:image/png;base64,iVBORw0KGgo...)
	var b64Data string
	var ext string = ".bin"

	if strings.Contains(base64String, ";base64,") {
		parts := strings.SplitN(base64String, ";base64,", 2)
		mimeTypePart := parts[0]
		b64Data = parts[1]

		// Extremely basic MIME-type to extension mapper
		if strings.Contains(mimeTypePart, "image/png") {
			ext = ".png"
		} else if strings.Contains(mimeTypePart, "image/jpeg") {
			ext = ".jpg"
		} else if strings.Contains(mimeTypePart, "application/pdf") {
			ext = ".pdf"
		} else if strings.Contains(mimeTypePart, "video/mp4") {
			ext = ".mp4"
		}
	} else {
		// Just a raw base64 string
		b64Data = base64String
		ext = ".png" // default to png if not specified for safety, though it could be anything
	}

	// Decode the base64 string
	decodedData, err := base64.StdEncoding.DecodeString(b64Data)
	if err != nil {
		return "", fmt.Errorf("failed to decode base64: %v", err)
	}

	// Ensure the directory exists
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		return "", fmt.Errorf("failed to create upload directory: %v", err)
	}

	// Generate a unique filename using timestamp
	fileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	filePath := filepath.Join(uploadDir, fileName)

	// Save the file
	if err := os.WriteFile(filePath, decodedData, 0644); err != nil {
		return "", fmt.Errorf("failed to write file: %v", err)
	}

	// Return a URL-friendly path, using forward slashes
	urlPath := strings.ReplaceAll(filePath, "\\", "/")
	if !strings.HasPrefix(urlPath, "/") && !strings.HasPrefix(urlPath, "./") {
		urlPath = "/" + urlPath
	}
	
	// Ensure standard absolute url representation based on root
	if strings.HasPrefix(urlPath, "uploads/") {
		urlPath = "/" + urlPath
	}

	return urlPath, nil
}
