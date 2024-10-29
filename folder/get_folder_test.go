package folder_test

import (
	"testing"
	"reflect"
	"sc-takehome/folder"
	"github.com/gofrs/uuid"
	// "github.com/stretchr/testify/assert"
)

func Test_folder_GetFoldersByOrgID(t *testing.T) {
	t.Parallel()
	tests := [...]struct {
		testName    string
		orgID   uuid.UUID
		folders []folder.Folder
		want    []folder.Folder
	}{
		{
			testName: "Get only folders in org",
			orgID: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
			folders: []folder.Folder{
				{
					Name:	"creative-scalphunter",
					OrgId:	uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
					Paths:	"creative-scalphunter",
				},
				{
					Name:	"clear-arclight",
					OrgId:	uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
					Paths:	"creative-scalphunter.clear-arclight",
				},
				{
					Name:	"topical-micromax",
					OrgId:	uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
					Paths:	"creative-scalphunter.clear-arclight.topical-micromax",
				},
				{
					Name:	"noble-vixen",
					OrgId:	uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths:	"noble-vixen",
				},
				{
					Name:	"nearby-secret",
					OrgId:	uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths:	"noble-vixen.nearby-secret",
				},
			},
			want: []folder.Folder{
				{
					Name:	"creative-scalphunter",
					OrgId:	uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
					Paths:	"creative-scalphunter",
				},
				{
					Name:	"clear-arclight",
					OrgId:	uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
					Paths:	"creative-scalphunter.clear-arclight",
				},
				{
					Name:	"topical-micromax",
					OrgId:	uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
					Paths:	"creative-scalphunter.clear-arclight.topical-micromax",
				},
			},
		},
		{
			testName: "Get nothing for empty org",
			orgID: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
			folders: []folder.Folder{
				{
					Name:	"noble-vixen",
					OrgId:	uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths:	"noble-vixen",
				},
				{
					Name:	"nearby-secret",
					OrgId:	uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths:	"noble-vixen.nearby-secret",
				},
			},
			want: []folder.Folder{

			},
		},
	}
	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			
			folders := folder.NewDriver(test.folders)
			res := folders.GetFoldersByOrgID(test.orgID)

			if !reflect.DeepEqual(test.want, res) {
				t.FailNow()
			}

		})
	}
}

func Test_folder_GetAllChildFolders(t *testing.T) {
	t.Parallel()
	tests := [...]struct {
		testName    string
		orgID   	uuid.UUID
		fileName	string
		folders 	[]folder.Folder
		want    	[]folder.Folder
		err			bool
	}{
		{
			testName:	"Gets all child folders",
			orgID:		uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
			fileName:	"creative-scalphunter",
			folders: []folder.Folder{
				{
					Name:	"creative-scalphunter",
					OrgId:	uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
					Paths:	"creative-scalphunter",
				},
				{
					Name:	"clear-arclight",
					OrgId:	uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
					Paths:	"creative-scalphunter.clear-arclight",
				},
				{
					Name:	"topical-micromax",
					OrgId:	uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
					Paths:	"creative-scalphunter.clear-arclight.topical-micromax",
				},
				{
					Name:	"close-layla-miller",
					OrgId:	uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
					Paths:	"creative-scalphunter.close-layla-miller",
				},
				{
					Name:	"noble-vixen",
					OrgId:	uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths:	"noble-vixen",
				},
				{
					Name:	"nearby-secret",
					OrgId:	uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths:	"noble-vixen.nearby-secret",
				},
			},
			want: []folder.Folder{
				{
					Name:	"clear-arclight",
					OrgId:	uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
					Paths:	"creative-scalphunter.clear-arclight",
				},
				{
					Name:	"topical-micromax",
					OrgId:	uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
					Paths:	"creative-scalphunter.clear-arclight.topical-micromax",
				},
				{
					Name:	"close-layla-miller",
					OrgId:	uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
					Paths:	"creative-scalphunter.close-layla-miller",
				},
			},
			err:	false,
		},
		{
			testName:	"Error on missing folder",
			orgID:		uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
			fileName:	"evident-silver-centurion",
			folders: []folder.Folder{
				{
					Name:	"creative-scalphunter",
					OrgId:	uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
					Paths:	"creative-scalphunter",
				},
				{
					Name:	"clear-arclight",
					OrgId:	uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
					Paths:	"creative-scalphunter.clear-arclight",
				},
				{
					Name:	"topical-micromax",
					OrgId:	uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
					Paths:	"creative-scalphunter.clear-arclight.topical-micromax",
				},
				{
					Name:	"close-layla-miller",
					OrgId:	uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
					Paths:	"creative-scalphunter.close-layla-miller",
				},
				{
					Name:	"noble-vixen",
					OrgId:	uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths:	"noble-vixen",
				},
				{
					Name:	"nearby-secret",
					OrgId:	uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths:	"noble-vixen.nearby-secret",
				},
			},
			want: []folder.Folder{

			},
			err:	true,
		},
		{
			testName:	"Error on missing org",
			orgID:		uuid.FromStringOrNil("badUUID"),
			fileName:	"creative-scalphunter",
			folders: []folder.Folder{
				{
					Name:	"creative-scalphunter",
					OrgId:	uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
					Paths:	"creative-scalphunter",
				},
				{
					Name:	"clear-arclight",
					OrgId:	uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
					Paths:	"creative-scalphunter.clear-arclight",
				},
				{
					Name:	"topical-micromax",
					OrgId:	uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
					Paths:	"creative-scalphunter.clear-arclight.topical-micromax",
				},
				{
					Name:	"close-layla-miller",
					OrgId:	uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
					Paths:	"creative-scalphunter.close-layla-miller",
				},
				{
					Name:	"noble-vixen",
					OrgId:	uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths:	"noble-vixen",
				},
				{
					Name:	"nearby-secret",
					OrgId:	uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths:	"noble-vixen.nearby-secret",
				},
			},
			want: []folder.Folder{

			},
			err:	true,
		},

	}
	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			
			folders := folder.NewDriver(test.folders)
			res, err := folders.GetAllChildFolders(test.orgID, test.fileName)

			if !test.err && err != nil {
				t.FailNow()
			} else if test.err && err == nil {
				t.FailNow()
			}

			if !reflect.DeepEqual(test.want, res) {
				t.FailNow()
			}

		})
	}
}