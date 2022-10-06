/* This file is auto-generated, manual edits in this file will be overwritten! */
package utils

import (
	"encoding/base64"
	"fmt"
	"mime"
	"net/http"

	"github.com/samber/lo"
)

// All the informations about a file
type File struct {
	Content     []byte
	Extension   *string
	ContentType string
}

func ToFile(data interface{}) (*File, error) {
	// Get bytes from data
	var bytes []byte
	switch t := data.(type) {
	case []byte:
		// Bytes given
		bytes = data.([]byte)
	case string:
		// Base64 given
		var err error
		bytes, err = base64.StdEncoding.DecodeString(data.(string))
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("Unsupported data type: %v", t)
	}

	// Get content type of the file
	contentType := http.DetectContentType(bytes)

	// Get the extensions of the file
	extension, err := GetExtensionFromContentType(contentType)
	if err != nil {
		return nil, err
	}

	return &File{
		Content:     bytes,
		Extension:   extension,
		ContentType: contentType,
	}, nil
}

func DecodeBase64(s string) (*File, error) {
	// Decode base64
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return nil, err
	}

	return ToFile(data)
}

func GetExtensionFromFile(data string) (*string, error) {
	contentType := http.DetectContentType([]byte(data))
	return GetExtensionFromContentType(contentType)
}

func GetExtensionFromContentType(contentType string) (*string, error) {
	mimeType, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		return nil, err
	}
	return GetExtensionFromMimeType(mimeType)
}

func GetExtensionFromMimeType(mimeType string) (*string, error) {
	switch mimeType {
	case "text/plain":
		return lo.ToPtr(".txt"), nil
	case "video/mp4":
		return lo.ToPtr(".mp4"), nil
	default:
		extensions, err := mime.ExtensionsByType(mimeType)
		if err != nil {
			return nil, err
		}
		if len(extensions) > 0 {
			return &extensions[0], nil
		}
	}
	return nil, nil
}
