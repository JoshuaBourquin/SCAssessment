package main

import (
	"fmt"
	"sc-takehome/folder"
	"github.com/gofrs/uuid"
)

func main() {

	orgID := uuid.FromStringOrNil(folder.DefaultOrgID)

	driver := folder.NewDriver(folder.GetAllFolders())

	res, err := driver.GetAllChildFolders(orgID, "stunning-horridus")

	if err != nil {
		panic(err)
	}

	folder.PrettyPrint(res)

	/*
	orgID := uuid.FromStringOrNil(folder.DefaultOrgID)

	res := folder.GetAllFolders()

	// example usage
	folderDriver := folder.NewDriver(res)
	orgFolder := folderDriver.GetFoldersByOrgID(orgID)

	folder.PrettyPrint(res)
	fmt.Printf("\n Folders for orgID: %s", orgID)
	folder.PrettyPrint(orgFolder)
	*/
}
