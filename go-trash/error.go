package trash

import "errors"

var (
	ErrMakeTrashFilesDir             = errors.New("failed to make $XDG_DATA_HOME/Trash/files directory")
	ErrMakeTrashInfoDir              = errors.New("failed to make $XDG_DATA_HOME/Trash/info directory")
	ErrTrashFile                     = errors.New("failed to trash file")
	ErrOpenTrashInfoFile             = errors.New("failed to open trash information file under $XDG_DATA_HOME/Trash/info")
	ErrGetTrashTargetAbsPath         = errors.New("failed to get trash target absolute path")
	ErrWriteTrashInfoFile            = errors.New("failed to write trash information file under $XDG_DATA_HOME/Trash/info")
	ErrReadTrashFileDir              = errors.New("failed to read $XDG_DATA_HOME/Trash/files directory")
	ErrReadTrashInfoFile             = errors.New("failed to read trash info file")
	ErrRestoreTargetFileAlreadyExist = errors.New("attempted to restore, however same name file already exists")
	ErrOriginalLocationNotRecord     = errors.New("trashed file original location is not recorded")
	ErrRestoreFile                   = errors.New("failed to restore file at original location")
	ErrInvalidInfoFile               = errors.New("invalid trash info file format")
	ErrDecodePercentEncoding         = errors.New("decoding error for percent-encoding in path")
	ErrEraseFileInTrash              = errors.New("failed to erase file in trash")
)
