package api

import (
	"fmt"
	"mime/multipart"
	"path/filepath"
	"strings"
)

func ImageStorage(path, name string, image *multipart.FileHeader) string {
	splitName := strings.Split(name, " ")
	joinName := strings.Join(splitName, "_")
	splitImage := strings.Split(image.Filename, " ")
	joinImage := strings.Join(splitImage, "_")
	fileName := filepath.Base(joinName + "_" + joinImage)
	fileDir := fmt.Sprintf("storage/%s/%s", path, fileName)

	return fileDir
}
