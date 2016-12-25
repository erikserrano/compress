// Paquete unzip encargado de descomprimir un archivo zip comprimido en 32 bits
package unzip

import (
	"archive/zip"
	"io"
	"os"
	"strings"
)

// Función encargada de crear un directorio
func createDirectory(path string, mode os.FileMode) (string, error) {
	if err := os.MkdirAll(path, mode); err != nil {
		return path, err
	}
	return path, nil
}

// Función encargada de crear un archivo
func createFile(path string) (*os.File, error) {
	copyPath := path
	if index := strings.LastIndex(copyPath, "/"); index >= 0 {
		copyPath = copyPath[0:index]
	}
	// Creamos directorio destino
	_, err := createDirectory(copyPath, os.ModePerm)
	if err != nil {
		return nil, err
	}
	return os.Create(path)
}

// Función encargada de copiar un directorio
func copyDirectory(path string, dir *zip.File) (string, error) {
	return createDirectory(path+dir.Name, dir.Mode())
}

// Función enacrgada de copiar un archivo
func copyFile(path string, file *zip.File) (os.FileInfo, error) {
	newFilePath := path + file.Name
	if index := strings.Index(file.Name, "/"); index >= 0 {
		newFilePath = path + string(file.Name[index+1:])
	}

	// Creamos archivo destino
	newFile, err := createFile(newFilePath)
	if err != nil {
		return nil, err
	}
	// Cerramos el archivo destino
	defer newFile.Close()

	// Abrimos archivo original
	originalFile, err := file.Open()
	if err != nil {
		return nil, err
	}

	// Copiamos archivo
	_, err = io.Copy(newFile, originalFile)

	return newFile.Stat()
}

// Función encargada de descomprimir un archivo zip (32 bits)
func Unzip(pathFile, pathDestination string) ([]os.FileInfo, error) {
	files := make([]os.FileInfo, 0)

	// Creamos la ruta de destino
	_, err := createDirectory(pathDestination, os.ModePerm)
	if err != nil {
		return files, err
	}

	// Abriendo archivo ZIP
	reader, err := zip.OpenReader(pathFile)
	if err != nil {
		return files, err
	}
	defer reader.Close()

	// Recorremos el contenido del archivo ZIP
	for _, f := range reader.File {
		if f.FileInfo().IsDir() {
			// Creamos directorio
			if _, err := copyDirectory(pathDestination, f); err != nil {
				return nil, err
			}

			files = append(files, f.FileInfo())
		} else {
			// Copiamos archivo
			fileInfo, err := copyFile(pathDestination, f)
			if err != nil {
				return nil, err
			}
			files = append(files, fileInfo)
		}
	}

	return files, nil
}
