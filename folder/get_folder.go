package folder

import (
	"github.com/gofrs/uuid"
	"strings"
	"fmt"
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

func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) ([]Folder, error) {

	// Get folders within org
	orgFolders := f.GetFoldersByOrgID(orgID)

	// Check that org has folders
	if len(orgFolders) == 0 {
		return []Folder{}, fmt.Errorf("Org %s contains no folders", orgID)
	}

	// Check that folder with name exists
	for _, folder := range orgFolders {
		if folder.Name == name {
			break
		}
		return []Folder{}, fmt.Errorf("Could not find folder with name %s", name)
	}

	// Get folders with paths that contain the parent folder's name
	res := []Folder{}
	for _, folder := range orgFolders {

		if strings.Contains(folder.Paths, name) && folder.Name != name {

			// Add folder to results
			res = append(res, folder)

		}
	}
	
	return res, nil
}
