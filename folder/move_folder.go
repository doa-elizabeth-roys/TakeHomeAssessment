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
		return nil, errors.New("Source folder does not exist")
	}

	// Find destination folder
	for i := range f.folders {
		if f.folders[i].Name == dst {
			dstFolder = &f.folders[i]
			break
		}
	}

	if dstFolder == nil {
		return nil, errors.New("Destination folder does not exist")
	}

	// Check if both folders are in the same org (optional)
	if srcFolder.OrgId != dstFolder.OrgId {
		return nil, errors.New("Cannot move a folder to a different organization")
	}

	// Check if destination folder is not the source folder (or a child of it)
	if srcFolder.Name == dstFolder.Name {
		return nil, errors.New("Cannot move a folder to itself")
	}

	if strings.HasPrefix(srcFolder.Paths, dstFolder.Paths) {
		return nil, errors.New("Cannot move a folder to a child of itself")
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
