package trash

import "errors"

var (
	ErrMakeTrashFilesDir     = errors.New("failed to make $XDG_DATA_HOME/Trash/files directory")
	ErrMakeTrashInfoDir      = errors.New("failed to make $XDG_DATA_HOME/Trash/info directory")
	ErrTrashFile             = errors.New("failed to trash file")
	ErrOpenTrashInfoFile     = errors.New("failed to open trash information file under $XDG_DATA_HOME/Trash/info")
	ErrGetTrashTargetAbsPath = errors.New("failed to get trash target absolute path")
	ErrWriteTrashInfoFile    = errors.New("failed to write trash information file under $XDG_DATA_HOME/Trash/info")
	ErrReadTrashFileDir      = errors.New("failed to read $XDG_DATA_HOME/Trash/files directory")
)
