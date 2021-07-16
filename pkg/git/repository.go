package git

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

type Repository interface {
	CreateBranch(name string) error
	CreateTag(name string) error
	RemoveBranch(name string) error
	RemoveTag(name string) error
	GitRepository() *git.Repository
	Option() *git.CloneOptions
}

type repository struct {
	*git.Repository
	option *git.CloneOptions
}

func (r *repository) GitRepository() *git.Repository {
	return r.Repository
}

func (r *repository) Option() *git.CloneOptions {
	return r.option
}

func (r *repository) CreateBranch(name string) error {
	headRef, err := r.Head()
	if err != nil {
		return err
	}

	ref := plumbing.NewHashReference(plumbing.NewBranchReferenceName(name), headRef.Hash())

	return r.Storer.SetReference(ref)
}

func (r *repository) CreateTag(name string) error {
	headRef, err := r.Head()
	if err != nil {
		return err
	}

	ref := plumbing.NewHashReference(plumbing.NewTagReferenceName(name), headRef.Hash())

	return r.Storer.SetReference(ref)
}

func (r *repository) RemoveBranch(name string) error {
	headRef, err := r.Head()
	if err != nil {
		return err
	}

	ref := plumbing.NewHashReference(plumbing.NewBranchReferenceName(name), headRef.Hash())

	return r.Storer.RemoveReference(ref.Name())
}

func (r *repository) RemoveTag(name string) error {
	headRef, err := r.Head()
	if err != nil {
		return err
	}

	ref := plumbing.NewHashReference(plumbing.NewTagReferenceName(name), headRef.Hash())

	return r.Storer.RemoveReference(ref.Name())
}
