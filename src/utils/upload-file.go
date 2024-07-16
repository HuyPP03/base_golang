package utils

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type UploadedFiles map[string][]string

func UploadFiles(c *gin.Context) (UploadedFiles, error) {
	form, err := c.MultipartForm()
	if err != nil {
		return nil, err
	}

	uploadedFiles := make(UploadedFiles)

	for fieldName, files := range form.File {
		for _, file := range files {
			fileExtension := filepath.Ext(file.Filename)
			filename := strings.TrimSuffix(file.Filename, fileExtension)
			newFilename := fmt.Sprintf("%s-%d%s", filename, time.Now().Unix(), fileExtension)
			dst := "./uploads/" + newFilename

			if err := c.SaveUploadedFile(file, dst); err != nil {
				return nil, err
			}

			if _, ok := uploadedFiles[fieldName]; !ok {
				uploadedFiles[fieldName] = []string{newFilename}
			} else {
				uploadedFiles[fieldName] = append(uploadedFiles[fieldName], newFilename)
			}
		}
	}

	return uploadedFiles, nil
}

func UploadFile(c *gin.Context, fieldName string) (string, error) {
	file, err := c.FormFile(fieldName)
	if err != nil {
		return "", err
	}

	fileExtension := filepath.Ext(file.Filename)
	filename := strings.TrimSuffix(file.Filename, fileExtension)
	newFilename := fmt.Sprintf("%s-%d%s", filename, time.Now().Unix(), fileExtension)

	dst := "./uploads/" + newFilename
	if err := c.SaveUploadedFile(file, dst); err != nil {
		return "", err
	}

	return newFilename, nil
}
