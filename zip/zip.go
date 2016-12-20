// Paquete zip encargado de comprimir un archivo zip en 32 bits
package zip

import (
	"archive/zip"
	"io"
	"os"
)

// Función encargada de elimina un directorio/archivo
func removeFileDirectory(path string) error {
	return os.Remove(path)
}

// Función encargada de crear un archivo
func createFile(filePath string) (*os.File, error) {
	// Eliminamos el archivo si existe
	removeFileDirectory(filePath)

	// Creamos archivo
	file, err := os.Create(filePath)
	if err != nil {
		return nil, err
	}
	// Asignamos permisos al archivo
	//err = file.Chmod(0777)
	return file, err
}

// Función encargada de crear un directorio
func createDirectory(path string, mode os.FileMode) error {
	if err := os.MkdirAll(path, mode); err != nil {
		return err
	}
	return nil
}

// Función encargada de leer el contenido de un archivo y escibirlo en el writer
func copyContent(file *os.File, writer io.Writer) (int64, error) {
	// Obtenemos la información del archivo
	fileInfo, err := file.Stat()
	if err != nil {
		return 0, err
	}

	// Leemos archivo
	bytes := make([]byte, fileInfo.Size())
	_, err = file.Read(bytes)
	if err != nil {
		return 0, err
	}

	// Guardamos contenido de archivo en writer (archivo contenido dentro del ZIP)
	size, err := writer.Write(bytes)
	return int64(size), err
}

// Función encargada de abrir una ruta (archivo/directorio)
func openPath(path string) (*os.File, error) {
	return os.Open(path)
}

// Función encargada de recorrer un directorio y agregar los archivos al ZIP
func walkDirectory(directory *os.File, zipWriter *zip.Writer) error {
	// Leemos contenido del directorio
	content, err := directory.Readdir(-1)
	if err != nil {
		return err
	}

	// Recorremos contenido del directorio
	for _, val := range content {
		if !val.IsDir() {
			// Es un archivo: anrimos el archivo y lo agregamos al Writer
			file, err := openPath(directory.Name() + "/" + val.Name())
			if err != nil {
				return err
			}

			// Agregamos un archivo al ZIP
			fileWriter, err := zipWriter.Create(file.Name())
			if err != nil {
				file.Close()
				return err
			}
			// Escrivimos sobre el nuevo archivo ZIP
			_, err = copyContent(file, fileWriter)
		} else {
			// Es un directorio: abrimos el directorio y recorremos su contenido
			dir, err := openPath(directory.Name() + "/" + val.Name())
			if err != nil {
				return err
			}
			err = walkDirectory(dir, zipWriter)
			if err != nil {
				return err
			}
		}
	}

	// Cerramos directorio
	directory.Close()
	return nil
}

// Función encargada de crear un archivo ZIP y escribir sobre él el contenido de un directorio/archivo
func Zip(startPath, finalFileName, finalFilePath string) error {
	// Abrimos la ruta inicial
	startDirectory, err := openPath(startPath)
	if err != nil {
		return err
	}
	// Cerramos la ruta inicial
	defer startDirectory.Close()

	// Creamos directorio destino
	if err := createDirectory(finalFilePath, 0777); err != nil {
		return err
	}

	// Creamos archivo ZIP
	zipFile, err := createFile(finalFilePath + finalFileName)
	if err != nil {
		return err
	}
	// Cerramos archivo ZIP
	defer zipFile.Close()

	// Creamos writer para el archivo ZIP
	zipWriter := zip.NewWriter(zipFile)
	// Cerramos Writer
	defer zipWriter.Close()

	// Iniciamos recorrido del directorio
	if err := walkDirectory(startDirectory, zipWriter); err != nil {
		return err
	}
	return nil
}
