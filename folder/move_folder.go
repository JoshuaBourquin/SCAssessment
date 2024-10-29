package folder

import (
	//"github.com/gofrs/uuid"
	"strings"
	"fmt"
)

func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {

	// Check source and destination are different
	if name == dst {
		return []Folder{}, fmt.Errorf("Cannot move a folder to itself")
	}

	var sourceFolder, destFolder *Folder

	// Find source folder
	for _, folder := range f.folders {
		
		if folder.Name == name {
			sourceFolder = &folder
			break
		}

	}

	if sourceFolder == nil {
		return []Folder{}, fmt.Errorf("Could not find source folder with name %s", name)
	}

	// Find destination folder
	for _, folder := range f.folders {

		if folder.Name == dst {
			destFolder = &folder
			break
		}

	}

	if destFolder == nil {
		return []Folder{}, fmt.Errorf("Could not find destination folder with name %s", dst)
	}

	// Check that orgs match
	if sourceFolder.OrgId != destFolder.OrgId {
		return []Folder{}, fmt.Errorf("Cannot move a folder to a different organization")
	}

	// Create new path for source folder
	newPath := destFolder.Paths + "." + sourceFolder.Name

	// Check that new path does not contain the source folder name twice - if it does then it is moving into a child of itself
	if strings.Count(newPath, sourceFolder.Name) > 1 {
		return []Folder{}, fmt.Errorf("Cannot move a folder into a child of itself")
	}

	// Change source folder to new path and all child folders to new path
	for index, folder := range f.folders {
		f.folders[index].Paths = strings.ReplaceAll(folder.Paths, sourceFolder.Paths, newPath)
	}

	return f.folders, nil
}
