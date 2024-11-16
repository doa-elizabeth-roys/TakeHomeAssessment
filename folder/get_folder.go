package folder

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gofrs/uuid"
)

func GetAllFolders() []Folder {
	return GetSampleData()
}

func (f *driver) GetFoldersByOrgID(orgID uuid.UUID) []Folder {
	folders := f.folders

	res := []Folder{}
	for _, f := range folders {
		if f.OrgId == orgID {
			res = append(res, f)
		}
	}

	return res

}

func (d *driver) GetAllChildFolders(orgID uuid.UUID, name string) ([]Folder, error) {
	// Input validation
	if orgID == uuid.Nil {
		return nil, errors.New("invalid orgID")
	}
	if name == "" {
		return nil, errors.New("folder name cannot be empty")
	}

	// Check if folder exists
	var folderExists bool
	for _, folder := range d.folders {
		if folder.OrgId == orgID && folder.Paths == name {
			folderExists = true
			break
		}
	}

	if !folderExists {
		return nil, fmt.Errorf("folder '%s' does not exist in the specified organisation '%s'", name, orgID)
	}

	// Find all child folders
	var childFolders []Folder
	for _, folder := range d.folders {
		if folder.OrgId == orgID && strings.HasPrefix(folder.Paths, name+".") {
			childFolders = append(childFolders, folder)
		}
	}

	return childFolders, nil
}
