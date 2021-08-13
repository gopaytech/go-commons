package gitlab_file

import (
	"sync"

	"github.com/gopaytech/go-commons/pkg/gitlab"
	gl "github.com/xanzy/go-gitlab"
)

type ProjectFiles struct {
	Project gl.Project
	Files   []gl.File
}

type GroupFileOps interface {
	ListProjectWithFileMetadata(id gitlab.NameOrId, ref string, fileNames ...string) (<-chan ProjectFiles, error)
	ListProjectWithFiles(id gitlab.NameOrId, ref string, fileNames ...string) (<-chan ProjectFiles, error)
	SearchProjectByFileContent(id gitlab.NameOrId, ref string, searchFunc SearchFunc, fileNames ...string) (<-chan ProjectFiles, error)
}

type groupFileOps struct {
	client  *gl.Client
	project ProjectFileOps
	group   gitlab.Group
}

func (g *groupFileOps) SearchProjectByFileContent(id gitlab.NameOrId, ref string, searchFunc SearchFunc, fileNames ...string) (<-chan ProjectFiles, error) {
	projectFiles := make(chan ProjectFiles)

	projectChan, err := g.group.ListAllProjects(id)
	if err != nil {
		return projectFiles, err
	}

	var wg sync.WaitGroup
	for p := range projectChan {
		wg.Add(1)
		go func(gp gl.Project) {
			files, _ := g.project.SearchFileWithContent(gitlab.NewId(gp.ID), ref, searchFunc, fileNames...)

			if len(files) > 0 {
				projectFile := ProjectFiles{
					Project: gp,
					Files:   files,
				}

				projectFiles <- projectFile
			}
			wg.Done()
		}(p)
	}

	go func() {
		wg.Wait()
		close(projectFiles)
	}()

	return projectFiles, nil
}

func (g *groupFileOps) ListProjectWithFiles(id gitlab.NameOrId, ref string, fileNames ...string) (<-chan ProjectFiles, error) {
	projectFiles := make(chan ProjectFiles)

	projectChan, err := g.group.ListAllProjects(id)
	if err != nil {
		return projectFiles, err
	}

	var wg sync.WaitGroup
	for p := range projectChan {
		wg.Add(1)
		go func(gp gl.Project) {
			files, _ := g.project.SearchFiles(gitlab.NewId(gp.ID), ref, fileNames...)

			if len(files) > 0 {
				projectFile := ProjectFiles{
					Project: gp,
					Files:   files,
				}

				projectFiles <- projectFile
			}
			wg.Done()
		}(p)
	}

	go func() {
		wg.Wait()
		close(projectFiles)
	}()

	return projectFiles, nil
}

func (g *groupFileOps) ListProjectWithFileMetadata(id gitlab.NameOrId, ref string, fileNames ...string) (<-chan ProjectFiles, error) {
	projectFiles := make(chan ProjectFiles)

	projectChan, err := g.group.ListAllProjects(id)
	if err != nil {
		return projectFiles, err
	}

	var wg sync.WaitGroup
	for p := range projectChan {
		wg.Add(1)
		go func(gp gl.Project) {
			files, _ := g.project.SearchFilesMetadata(gitlab.NewId(gp.ID), ref, fileNames...)

			if len(files) > 0 {
				projectFile := ProjectFiles{
					Project: gp,
					Files:   files,
				}

				projectFiles <- projectFile
			}
			wg.Done()
		}(p)
	}

	go func() {
		wg.Wait()
		close(projectFiles)
	}()

	return projectFiles, nil
}

func NewGroupFileOps(client *gl.Client, project ProjectFileOps) GroupFileOps {
	return &groupFileOps{client: client, project: project}
}
