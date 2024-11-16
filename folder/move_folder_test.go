package folder_test

import (
	"testing"

	"encoding/json"
	"os"

	"github.com/georgechieng-sc/interns-2022/folder"
)

func Test_folder_MoveFolder(t *testing.T) {
	data, err := os.ReadFile("sample.json")
	if err != nil {
		t.Fatal(err)
	}

	var folders []folder.Folder
	err = json.Unmarshal(data, &folders)
	if err != nil {
		t.Fatal(err)
	}

	driver := folder.NewDriver(folders)

	tests := []struct {
		name       string
		srcName    string
		dstName    string
		wantErr    bool
		wantFolder *folder.Folder
	}{
		{
			name:       "move folder within same organization",
			srcName:    "clear-arclight",
			dstName:    "topical-micromax",
			wantErr:    false,
			wantFolder: &folder.Folder{Name: "clear-arclight", Paths: "creative-scalphunter.topical-micromax.clear-arclight"},
		},
		{
			name:       "move folder to different organization",
			srcName:    "clear-arclight",
			dstName:    "noble-vixen",
			wantErr:    true,
			wantFolder: nil,
		},
		{
			name:       "move folder to itself",
			srcName:    "clear-arclight",
			dstName:    "clear-arclight",
			wantErr:    true,
			wantFolder: nil,
		},
		{
			name:       "move folder to child",
			srcName:    "topical-micromax",
			dstName:    "bursting-lionheart",
			wantErr:    true,
			wantFolder: nil,
		},
		{
			name:       "move non-existent folder",
			srcName:    "non-existent",
			dstName:    "topical-micromax",
			wantErr:    true,
			wantFolder: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			folders, err := driver.MoveFolder(tt.srcName, tt.dstName)
			if (err != nil) != tt.wantErr {
				t.Errorf("MoveFolder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				for _, folder := range folders {
					if folder.Name == tt.wantFolder.Name && folder.Paths == tt.wantFolder.Paths {
						return
					}
				}
				t.Errorf("MoveFolder() did not update folder correctly")
			}
		})
	}
}
