package folder_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_folder_GetFoldersByOrgID(t *testing.T) {
	t.Parallel()

	tests := [...]struct {
		name    string
		orgID   uuid.UUID
		folders []folder.Folder
		want    []folder.Folder
	}{
		{
			name:  "get folders by existing orgID",
			orgID: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
			folders: []folder.Folder{
				{
					OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
					Paths: "creative-scalphunter",
					Name:  "creative-scalphunter",
				},
				{
					OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
					Paths: "creative-scalphunter.subfolder1",
					Name:  "subfolder1",
				},
				{
					OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths: "noble-vixen",
					Name:  "noble-vixen",
				},
			},
			want: []folder.Folder{
				{
					OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
					Paths: "creative-scalphunter",
					Name:  "creative-scalphunter",
				},
				{
					OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
					Paths: "creative-scalphunter.subfolder1",
					Name:  "subfolder1",
				},
			},
		},
		{
			name:  "get folders by empty orgID",
			orgID: uuid.Nil,
			folders: []folder.Folder{
				{
					OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
					Paths: "creative-scalphunter",
					Name:  "creative-scalphunter",
				},
				{
					OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths: "noble-vixen",
					Name:  "noble-vixen",
				},
			},
			want: []folder.Folder{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			get := f.GetFoldersByOrgID(tt.orgID)
			assert.ElementsMatch(t, tt.want, get)
		})
	}
}

func Test_folder_GetAllChildFolders(t *testing.T) {
	t.Parallel()

	tests := [...]struct {
		name       string
		orgID      uuid.UUID
		nameToFind string
		folders    []folder.Folder
		want       []folder.Folder
		wantErr    bool
	}{
		{
			name:       "get child folders by existing orgID and folder name",
			orgID:      uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
			nameToFind: "creative-scalphunter",
			folders: []folder.Folder{
				{
					OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
					Paths: "creative-scalphunter",
					Name:  "creative-scalphunter",
				},
				{
					OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
					Paths: "creative-scalphunter.subfolder1",
					Name:  "subfolder1",
				},
				{
					OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths: "noble-vixen",
					Name:  "noble-vixen",
				},
			},
			want: []folder.Folder{
				{
					OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
					Paths: "creative-scalphunter.subfolder1",
					Name:  "subfolder1",
				},
			},
		},
		{
			name:       "folder does not exist in the specified organization",
			orgID:      uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
			nameToFind: "non-existent-folder",
			folders: []folder.Folder{
				{
					OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
					Paths: "creative-scalphunter",
					Name:  "creative-scalphunter",
				},
				{
					OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths: "noble-vixen",
					Name:  "noble-vixen",
				},
			},
			wantErr: true,
		},
		{
			name:       "get child folders by empty orgID",
			orgID:      uuid.Nil,
			nameToFind: "creative-scalphunter",
			folders: []folder.Folder{
				{
					OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
					Paths: "creative-scalphunter",
					Name:  "creative-scalphunter",
				},
				{
					OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths: "noble-vixen",
					Name:  "noble-vixen",
				},
			},
			wantErr: true,
		},
		{
			name:       "get child folders by empty nameToFind",
			orgID:      uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
			nameToFind: "",
			folders: []folder.Folder{
				{
					OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
					Paths: "creative-scalphunter",
					Name:  "creative-scalphunter",
				},
				{
					OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths: "noble-vixen",
					Name:  "noble-vixen",
				},
			},
			wantErr: true,
		},
		{
			name:       "get child folders by non-existent orgID",
			orgID:      uuid.FromStringOrNil("non-existent-org-id"),
			nameToFind: "creative-scalphunter",
			folders: []folder.Folder{
				{
					OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
					Paths: "creative-scalphunter",
					Name:  "creative-scalphunter",
				},
				{
					OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths: "noble-vixen",
					Name:  "noble-vixen",
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			get, err := f.GetAllChildFolders(tt.orgID, tt.nameToFind)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllChildFolders() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err.Error() != "Folder does not exist in the specified organization" {
					t.Errorf("Expected error message = %q, got %q", "Folder does not exist in the specified organization", err.Error())
				}
			} else {
				assert.ElementsMatch(t, tt.want, get)
			}
		})
	}
}
