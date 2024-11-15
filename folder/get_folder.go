package folder

import (
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

// func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) []Folder {
// 	// Your code here...
// 	// Slice to hold child folders
// 	var childFolders []Folder

// 	// Iterate through all folders
// 	for _, folder := range f.folders {
// 		// Check if the folder belongs to the same organization
// 		if folder.OrgId == orgID {
// 			// Check if the folder's path starts with the parent folder's path (name)
// 			if strings.HasPrefix(folder.Paths, name+".") {
// 				// Add to childFolders if it is a match
// 				childFolders = append(childFolders, folder)
// 			}
// 		}

// 		if (folder.OrgId != orgID || name != folder.Paths){

// 		}
// 	}
// 	//return []Folder{}
// 	return childFolders
// }

func (d *driver) GetAllChildFolders(orgID uuid.UUID, name string) []Folder {
	var childFolders []Folder
	var parentFolder *Folder

	// Find the parent folder based on orgID and name
	for _, folder := range d.folders {
		if folder.OrgId == orgID && folder.Paths == name {
			parentFolder = &folder
			break
		}
	}

	// If parent folder is not found, return an empty slice
	if parentFolder == nil {
		return []Folder{}
	}

	// Find all child folders where the path starts with the parent's path followed by "."
	for _, folder := range d.folders {
		if folder.OrgId == orgID && strings.HasPrefix(folder.Paths, parentFolder.Paths+".") {
			childFolders = append(childFolders, folder)
		}
	}

	//Return the list of child folders
	return childFolders
}
