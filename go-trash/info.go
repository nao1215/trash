package trash

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"
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
	infoFilePath := filepath.Join(i.Path, fmt.Sprintf("%s%s", filepath.Base(trashDestPath), trashInfoFileExt))

	f, err := os.OpenFile(infoFilePath, os.O_WRONLY|os.O_CREATE, 0600)
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
	content += fmt.Sprintf("DeletionDate=%s\n", now)

	_, err = fmt.Fprint(f, content)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func (i *Info) originalFileLocation(fileNameInTrashCan string) (string, error) {
	targetFilePath := filepath.Join(i.Path, fileNameInTrashCan)
	infoFilePath := fmt.Sprintf("%s%s", targetFilePath, trashInfoFileExt)

	keyValue, err := i.getKeyValue(infoFilePath)
	if err != nil {
		return "", err
	}

	value, ok := keyValue["Path"]
	if !ok {
		return "", ErrOriginalLocationNotRecord
	}
	return value, nil
}

// spec.: The implementation MUST ignore any other lines in this file, except the first
// line (must be [Trash Info]) and these two key/value pairs. If a string that starts
// with “Path=” or “DeletionDate=” occurs several times, the first occurrence is to be used
func (i *Info) getKeyValue(path string) (map[string]string, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrReadTrashInfoFile, err)
	}

	list := strings.Split(string(bytes), "\n")

	// header check
	if !strings.HasPrefix(list[0], trashInfoFileHeader) {
		return nil, fmt.Errorf("%w: %s", ErrInvalidInfoFile, list[0])
	}

	keyValueMap := make(map[string]string, 0)
	for _, v := range list[1:] {
		v = strings.ReplaceAll(v, " ", "")

		if _, ok := keyValueMap["Path"]; ok {
			continue
		}
		if _, ok := keyValueMap["DeletionDate"]; ok {
			continue
		}

		keyValue := strings.Split(v, "=")
		if keyValue[0] != "Path" && keyValue[0] != "DeletionDate" {
			continue
		}

		if keyValue[0] == "Path" {
			path, err := url.QueryUnescape(strings.Join(keyValue[1:], ""))
			if err != nil {
				return nil, fmt.Errorf("%w: %w", ErrDecodePercentEncoding, err)
			}
			keyValueMap["Path"] = path
			continue
		}
		keyValueMap["DeletionDate"] = strings.Join(keyValue[1:], "")
	}
	return keyValueMap, nil
}
