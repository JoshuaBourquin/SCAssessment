package folder_test

import (
	"testing"
	"reflect"
	"sc-takehome/folder"
	"github.com/gofrs/uuid"
)

func Test_folder_MoveFolder(t *testing.T) {
	t.Parallel()
	tests := [...]struct {
		testName    string
		name		string
		dst			string
		folders 	[]folder.Folder
		want    	[]folder.Folder
		err			bool
	}{
		{
			testName:	"Leaf folder successfully moves to new destination",
			name:		"bursting-lionheart",
			dst:		"close-layla-miller",
			folders:	[]folder.Folder{
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
					Name:	"bursting-lionheart",
					OrgId:	uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
					Paths:	"creative-scalphunter.clear-arclight.topical-micromax.bursting-lionheart",
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
			want:		[]folder.Folder{
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
					Name:	"bursting-lionheart",
					OrgId:	uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
					Paths:	"creative-scalphunter.close-layla-miller.bursting-lionheart",
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
			err:		false,
		},
		{
			testName:	"Branch folder moves child folders when moved",
			name:		"topical-micromax",
			dst:		"close-layla-miller",
			folders:	[]folder.Folder{
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
					Name:	"bursting-lionheart",
					OrgId:	uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
					Paths:	"creative-scalphunter.clear-arclight.topical-micromax.bursting-lionheart",
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
			want:		[]folder.Folder{
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
					Paths:	"creative-scalphunter.close-layla-miller.topical-micromax",
				},
				{
					Name:	"bursting-lionheart",
					OrgId:	uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
					Paths:	"creative-scalphunter.close-layla-miller.topical-micromax.bursting-lionheart",
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
			err:		false,
		},
		{
			testName:	"Error when moving to different org",
			name:		"topical-micromax",
			dst:		"noble-vixen",
			folders:	[]folder.Folder{
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
					Name:	"bursting-lionheart",
					OrgId:	uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
					Paths:	"creative-scalphunter.clear-arclight.topical-micromax.bursting-lionheart",
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
			want:		[]folder.Folder{
				
			},
			err:		true,
		},
		{
			testName:	"Error when moving folder into itself",
			name:		"topical-micromax",
			dst:		"topical-micromax",
			folders:	[]folder.Folder{
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
					Name:	"bursting-lionheart",
					OrgId:	uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
					Paths:	"creative-scalphunter.clear-arclight.topical-micromax.bursting-lionheart",
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
			want:		[]folder.Folder{
				
			},
			err:		true,
		},
		{
			testName:	"Error when moving folder into child of itself",
			name:		"topical-micromax",
			dst:		"bursting-lionheart",
			folders:	[]folder.Folder{
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
					Name:	"bursting-lionheart",
					OrgId:	uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
					Paths:	"creative-scalphunter.clear-arclight.topical-micromax.bursting-lionheart",
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
			want:		[]folder.Folder{
				
			},
			err:		true,
		},
		{
			testName:	"Error if source folder does not exist",
			name:		"DOESNOTEXIST",
			dst:		"bursting-lionheart",
			folders:	[]folder.Folder{
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
					Name:	"bursting-lionheart",
					OrgId:	uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
					Paths:	"creative-scalphunter.clear-arclight.topical-micromax.bursting-lionheart",
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
			want:		[]folder.Folder{
				
			},
			err:		true,
		},
		{
			testName:	"Error if destination folder does not exist",
			name:		"topical-micromax",
			dst:		"DOESNOTEXIST",
			folders:	[]folder.Folder{
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
					Name:	"bursting-lionheart",
					OrgId:	uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
					Paths:	"creative-scalphunter.clear-arclight.topical-micromax.bursting-lionheart",
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
			want:		[]folder.Folder{
				
			},
			err:		true,
		},
	}
	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			
			folders := folder.NewDriver(test.folders)
			res, err := folders.MoveFolder(test.name, test.dst)

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