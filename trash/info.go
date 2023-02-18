package trash

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

const (
	trashInfoFileHeader string = "[Trash Info]"
	trashInfoTimeFormat string = "2006-01-02T15:04:05"
	trashInfoFileExt    string = ".trashinfo"
)

type Info struct {
	Path string
}

func NewInfo(trashPath string) *Info {
	return &Info{
		Path: filepath.Join(trashPath, "info"),
	}
}

func (i *Info) MakeDir() error {
	if err := os.MkdirAll(i.Path, 0700); err != nil {
		return fmt.Errorf("%w: %w", ErrMakeTrashInfoDir, err)
	}
	return nil
}

func (i *Info) MakeInfoFile(trashSrcPath, trashDestPath string) error {
	infoFilePath := filepath.Join(i.Path, filepath.Base(trashDestPath))

	f, err := os.OpenFile(infoFilePath, os.O_WRONLY, 0600)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrOpenTrashInfoFile, err)
	}
	defer f.Close()

	absOrgPath, err := filepath.Abs(trashSrcPath)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrGetTrashTargetAbsPath, err)
	}
	now := time.Now().Format(trashInfoTimeFormat)

	content := fmt.Sprintf("%s\n", trashInfoFileHeader)
	content += fmt.Sprintf("Path=%s\n", absOrgPath)
	content += fmt.Sprintf("DeletionDate=%s", now)

	_, err = fmt.Fprintf(f, content)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}
