package structs

import (
	"encoding/json"
	"io/ioutil"
)

const FILE_NAME = "storage/workspace.json"

type Workspace struct {
	Id   int
	Name string
	Desc string
}

/**
Create new workspace in []*Workspace and stored it in the storage
Then return new &workspace
*/
func CreateWorkspace(name string, desc string) *Workspace {
	wss := GetWorkspaces()
	newWs := Workspace{Id: len(wss), Name: name, Desc: desc}

	wss = append(wss, &newWs)

	file, _ := json.MarshalIndent(wss, "", "")
	_ = ioutil.WriteFile(FILE_NAME, file, 0777)

	return &newWs
}

/**
Get workspaces from storage and return as slice of *Workspace
*/
func GetWorkspaces() []*Workspace {
	file, _ := ioutil.ReadFile(FILE_NAME)
	wss := []*Workspace{}
	_ = json.Unmarshal([]byte(file), &wss)

	return wss
}
