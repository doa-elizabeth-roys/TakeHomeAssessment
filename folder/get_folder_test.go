package folder_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

// feel free to change how the unit test is structured
func Test_folder_GetFoldersByOrgID(t *testing.T) {
	t.Parallel()
	tests := [...]struct {
		name    string
		orgID   uuid.UUID
		folders []folder.Folder
		want    []folder.Folder
	}{
		{
			name:  "test with matching orgID",
			orgID: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
			folders: []folder.Folder{
				{Name: "alpha", OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"), Paths: "alpha"},
				{Name: "bravo", OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"), Paths: "alpha.bravo"},
				{Name: "charlie", OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"), Paths: "alpha.bravo.charlie"},
				{Name: "delta", OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"), Paths: "alpha.delta"},
			},
			want: []folder.Folder{
				{Name: "alpha", OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"), Paths: "alpha"},
				{Name: "bravo", OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"), Paths: "alpha.bravo"},
				{Name: "charlie", OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"), Paths: "alpha.bravo.charlie"},
				{Name: "delta", OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"), Paths: "alpha.delta"},
			},
		},
		{
			name:  "test with no matching orgID",
			orgID: uuid.FromStringOrNil("invalid-uuid-1234"), // Invalid orgID
			folders: []folder.Folder{
				{Name: "alpha", OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"), Paths: "alpha"},
				{Name: "bravo", OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"), Paths: "alpha.bravo"},
			},
			want: []folder.Folder{}, // No folders should match
		},
		{
			name:  "test with mixed orgIDs",
			orgID: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"), // Matching orgID
			folders: []folder.Folder{
				{Name: "alpha", OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"), Paths: "alpha"},
				{Name: "bravo", OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"), Paths: "alpha.bravo"},
				{Name: "foxtrot", OrgId: uuid.FromStringOrNil("some-other-org-id"), Paths: "foxtrot"},
			},
			want: []folder.Folder{
				{Name: "alpha", OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"), Paths: "alpha"},
				{Name: "bravo", OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"), Paths: "alpha.bravo"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			get := f.GetFoldersByOrgID(tt.orgID)
			assert.ElementsMatch(t, tt.want, get, "Folders by orgID don't match")

		})
	}
}
