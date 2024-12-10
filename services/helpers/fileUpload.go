package helpers

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
)

// func FileUpload(file *multipart.FileHeader, c *gin.Context, destination string) (*string, error) {

// 	if err := c.SaveUploadedFile(file, destination); err != nil {
// 		return nil, err
// 	}

// 	return &destination, nil
// }

// func UploadByName(c *gin.Context, name string, file *multipart.FileHeader) (*string, error) {
// 	var image *multipart.FileHeader
// 	var err error

// 	if file == nil {
// 		image, err = c.FormFile(name)
// 		if err != nil {
// 			return nil, err
// 		}
// 	} else {
// 		image = file
// 	}

// 	imagePath := fmt.Sprintf("../uploads/%s_%s", uuid.New(), image.Filename)

// 	imgPath, uploadErr := FileUpload(image, c, imagePath)
// 	if uploadErr != nil {
// 		return nil, uploadErr
// 	}

// 	return imgPath, nil
// }

// func CheckFileExists(c *gin.Context, filename *string) bool {
// 	file, err := c.FormFile("featuredimage")
// 	if err != nil {
// 		return false
// 	}

// 	if file.Filename == *filename {
// 		return true
// 	}
// 	return false
// }

func UploadFile(data string) (*string, error) {
	decodedData, fileType, err := decodeBase64File(data)

	if err != nil {
		return nil, err
	}
	if fileType == nil {
		*fileType = "jpeg"
	}
	path := fmt.Sprintf("%s_%s.%s", uuid.New(), time.Now().Format("20060102_150405"), *fileType)
	err = saveFile(path, decodedData)

	if err != nil {
		return nil, err
	}

	return &path, nil
}

func decodeBase64File(data string) ([]byte, *string, error) {
	// Check if the base64 string has a data URI prefix
	var fileType string
	if strings.Contains(data, ",") {
		parts := strings.Split(data, ",")
		data = parts[1]
		fileType = GetFileExtension(parts[0])
	}

	decodedData, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, nil, err
	}
	return decodedData, &fileType, nil
}

func saveFile(filePath string, data []byte) error {
	// Ensure the directory exists
	err := os.MkdirAll("./uploads", os.ModePerm)
	if err != nil {
		return err
	}

	filePath = fmt.Sprintf("../uploads/%s", filePath)

	// Write the file
	err = ioutil.WriteFile(filePath, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func GetFileExtension(data string) string {
	parts := strings.Split(data, ";")
	if len(parts) > 0 {
		mimeType := parts[0] // "image/jpeg"
		mimeParts := strings.Split(mimeType, "/")
		if len(mimeParts) > 1 {
			return mimeParts[1] // "jpeg"
		}
	}
	return "" // Return an empty string if no extension is found
}
