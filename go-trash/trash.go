// Package trash is a library following the FreeDesktop.org Trash specification.
// ref. https://specifications.freedesktop.org/trash-spec/trashspec-latest.html
package trash

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
	"github.com/nao1215/gorky/file"
)

type Trash struct {
	// HomeTrashCanPath is "user home trash" directory - the storage of files that were trashed (“deleted”) by the user.
	// These files can be listed, undeleted, or cleaned from the trash can.
	// The trash can directory is located at $XDG_DATA_HOME/Trash.
	HomeTrashCanPath string
	Info             *Info
}

func NewTrash() (*Trash, error) {
	trash := &Trash{
		HomeTrashCanPath: filepath.Join(xdg.DataHome, "Trash"),
	}
	trash.Info = NewInfo(trash.HomeTrashCanPath)

	if !file.IsDir(trash.HomeTrashCanPath) {
		if err := trash.makeTrashDirs(); err != nil {
			return nil, err
		}
	}
	return trash, nil
}

func (t *Trash) makeTrashDirs() error {
	filesDir := filepath.Join(t.HomeTrashCanPath, "files")

	if err := os.MkdirAll(filesDir, 0700); err != nil {
		return fmt.Errorf("%w: %w", ErrMakeTrashFilesDir, err)
	}

	if err := t.Info.MakeDir(); err != nil {
		return err
	}
	return nil
}

func (t *Trash) Trash(srcPath string) error {
	destPath := t.decideTrashDestPath(srcPath)

	if err := t.Info.MakeInfoFile(srcPath, destPath); err != nil {
		return err
	}

	if err := os.Rename(srcPath, destPath); err != nil {
		return fmt.Errorf("%w: %w", ErrTrashFile, err)
	}
	return nil
}

func (t *Trash) List() ([]string, error) {
	files, err := os.ReadDir(filepath.Join(t.HomeTrashCanPath, "files"))
	if err != nil {
		return nil, ErrReadTrashFileDir
	}

	fileNameList := make([]string, 0)
	for _, v := range files {
		fileNameList = append(fileNameList, filepath.Join(t.HomeTrashCanPath, "files", v.Name()))
	}
	return fileNameList, nil
}

func (t *Trash) decideTrashDestPath(targetFile string) string {
	basename := filepath.Base(targetFile)
	dest := filepath.Join(t.HomeTrashCanPath, basename)
	if !file.Exists(dest) {
		return dest
	}

	destFileName, destFileExt := splitFileName(basename)
	for i := 2; ; i++ {
		destPath := filepath.Join(t.HomeTrashCanPath, "files", fmt.Sprintf("%s.%d%s", destFileName, i, destFileExt))
		if !file.Exists(destPath) {
			return destPath
		}
	}
}

func splitFileName(file string) (name string, ext string) {
	ext = filepath.Ext(file)
	name = file[:len(file)-len(ext)]
	return name, ext
}
