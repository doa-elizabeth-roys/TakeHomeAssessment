package folder

import (
	"errors"
	// "strings"
)

func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {
	var srcFolder *Folder
	var dstFolder *Folder

	// Find source folder
	for i := range f.folders {
		if f.folders[i].Name == name {
			srcFolder = &f.folders[i]
			break
		}
	}
	if srcFolder == nil {
		return nil, errors.New("source folder does not exist")
	}

	// Find destination folder
	for i := range f.folders {
		if f.folders[i].Name == dst {
			dstFolder = &f.folders[i]
			break
		}
	}
	if dstFolder == nil {
		return nil, errors.New("destination folder does not exist")
	}

	// Return the updated folder structure
	return f.folders, nil
}
