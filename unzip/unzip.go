// Paquete unzip encargado de descomprimir un archivo zip comprimido en 32 bits
package unzip

import (
	"archive/zip"
	"io"
	"os"
	"time"
)

// Estructura enacargada de archivar la información del directorio/arhivo comprimido/descomprimio
type ZipContent struct {
	FileName    string    `json:"file_name"`
	IsDirectory bool      `json:"is_directory"`
	Size        int64     `json:"size"`
	Error       error     `json:"error"`
	CreatedAt   time.Time `json:"created_at"`
}


var (
	files map[int]ZipFiles
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
	return os.Create(path)
}

// Función encargada de copiar un directorio
func copyDirectory(path string, dir *zip.File) (string, error) {
	return createDirectory(path+dir.Name, dir.Mode())
}

// Función enacrgada de copiar un archivo
func copyFile(path string, file *zip.File) (int64, string, error) {
	// Creamos archivo destino
	newFile, err := createFile(path + file.Name)
	if err != nil {
		return 0, (path + file.Name), err
	}
	// Cerramos el archivo destino
	defer newFile.Close()

	// Abrimos archivo original
	originalFile, err := file.Open()
	if err != nil {
		return 0, (path + file.Name), err
	}

	// Copiamos archivo
	_, err = io.Copy(newFile, originalFile)

	fileInfo, err := newFile.Stat()
	if err != nil {
		return 0, (path + file.Name), err
	}

	return fileInfo.Size(), (path + file.Name), nil
}

// Función encargada de descomprimir un archivo zip (32 bits)
func Unzip(pathFile, pathDestination string) (map[int]ZipContent, error) {
	files = make(map[int]ZipContent)

	// Creamos la ruta de destino
	_, err := createDirectory(pathDestination, os.ModePerm)
	if err != nil {
		return files, err
	}

	// Abriendo archivo ZIP
	reader, err := zip.OpenReader(os.Args[1])
	if err != nil {
		return files, err
	}
	defer reader.Close()

	// Recorremos el contenido del archivo ZIP
	for i, f := range reader.File {

		if f.FileInfo().IsDir() {
			// Creamos directorio
			path, err := copyDirectory(pathDestination, f)

			files[i] = compress.ZipContent{
				FileName:    path,
				IsDirectory: true,
				Size:        0,
				Error:       err,
				CreatedAt:   time.Now(),
			}
		} else {
			// Copiamos archivo
			size, path, err := copyFile(pathDestination, f)

			files[i] = compress.ZipContent{
				FileName:    path,
				IsDirectory: false,
				Size:        size,
				Error:       err,
				CreatedAt:   time.Now(),
			}
		}
	}

	return files, nil
}
