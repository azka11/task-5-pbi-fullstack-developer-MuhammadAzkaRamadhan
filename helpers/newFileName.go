package helpers

import (
	"path/filepath"
	// "time"

	uuid "github.com/satori/go.uuid"
)

func NewFileName(oldName string) string {
	ext := filepath.Ext(oldName)
	uuid := uuid.NewV4()
	newName := uuid.String() + ext
	return newName
}
