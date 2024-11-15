package folder

import (
	"errors"
	"strings"
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

	// Move the folder: update the source folder's path
	newPath := dstFolder.Paths + "." + srcFolder.Name
	srcFolder.Paths = newPath

	// Update the paths of child folders
	for i := range f.folders {
		// Check if the folder's path starts with the source folder's path
		if strings.HasPrefix(f.folders[i].Paths, srcFolder.Paths+".") {
			// Update child folder's path to reflect new structure
			f.folders[i].Paths = dstFolder.Paths + "." + strings.TrimPrefix(f.folders[i].Paths, srcFolder.Paths)
		}
	}

	// Return the updated folder structure
	return f.folders, nil
}
