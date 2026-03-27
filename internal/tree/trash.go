package tree

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func trashFile(path string) error {
	trashDir := xdgTrashDir()
	filesDir := filepath.Join(trashDir, "files")
	infoDir := filepath.Join(trashDir, "info")

	if err := os.MkdirAll(filesDir, 0700); err != nil {
		return err
	}
	if err := os.MkdirAll(infoDir, 0700); err != nil {
		return err
	}

	absPath, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	base := filepath.Base(absPath)
	trashName := base
	for i := 2; ; i++ {
		if _, err := os.Lstat(filepath.Join(filesDir, trashName)); os.IsNotExist(err) {
			break
		}
		trashName = fmt.Sprintf("%s.%d", base, i)
	}

	infoPath := filepath.Join(infoDir, trashName+".trashinfo")
	trashInfo := fmt.Sprintf(
		"[Trash Info]\nPath=%s\nDeletionDate=%s\n",
		absPath,
		time.Now().Format("2006-01-02T15:04:05"),
	)
	if err := os.WriteFile(infoPath, []byte(trashInfo), 0600); err != nil {
		return err
	}

	destPath := filepath.Join(filesDir, trashName)
	if err := os.Rename(absPath, destPath); err != nil {
		os.Remove(infoPath)
		return err
	}

	return nil
}

func xdgTrashDir() string {
	if xdg := os.Getenv("XDG_DATA_HOME"); xdg != "" {
		return filepath.Join(xdg, "Trash")
	}
	return filepath.Join(os.Getenv("HOME"), ".local", "share", "Trash")
}
