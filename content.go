// Paquet encargado de comprimir/descomprimir archivos ZIP de 32 bits
package compress

import (
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