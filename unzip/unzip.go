// Package unzip reads and writes standard 32 bit zip file
package unzip

import (
	"archive/zip"
	"io"
	"os"
	"strings"
)

// createFile create new empty file
func createFile(path string) (*os.File, error) {
	copyPath := path
	if index := strings.LastIndex(copyPath, "/"); index >= 0 {
		copyPath = copyPath[0:index]
	} else if index = strings.LastIndex(copyPath, "\\"); index >= 0 {
		copyPath = copyPath[0:index]
	}

	// Creamos directorio destino
	if err := os.MkdirAll(copyPath, os.ModePerm); err != nil {
		return nil, err
	}

	return os.Create(path)
}

// copyFile create a file with data from zip´s item
func copyFile(path string, file *zip.File) error {
	// create empty file
	newFile, err := createFile(path)
	if err != nil {
		return err
	}
	defer newFile.Close()

	originalFile, err := file.Open()
	if err != nil {
		return err
	}
	defer originalFile.Close()

	_, err = io.Copy(newFile, originalFile)

	return err
}

// Unzip reads standard zip 32 bit file and extract its content
func Unzip(pathFile, pathDestination string) ([]os.FileInfo, error) {
	var files []os.FileInfo

	// create destination directory
	if err := os.MkdirAll(pathDestination, os.ModePerm); err != nil {
		return files, err
	}

	// open zip file
	reader, err := zip.OpenReader(pathFile)
	if err != nil {
		return files, err
	}
	defer reader.Close()

	// read zip file items
	for _, f := range reader.File {
		files = append(files, f.FileInfo())
		if f.FileInfo().IsDir() {
			// create directory
			if err := os.MkdirAll(pathDestination+f.Name, f.Mode()); err != nil {
				return files, err
			}
		} else {
			// create file
			if err := copyFile(pathDestination+f.Name, f); err != nil {
				return files, err
			}
		}
	}

	return files, nil
}
