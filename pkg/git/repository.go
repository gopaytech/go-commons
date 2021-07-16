package git

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
)

type Repository interface {
	CreateBranch(name string) error
	CreateTag(name string) error

	RemoveBranch(name string) error
	RemoveTag(name string) error

	CheckoutBranch(name string) error
	CheckoutTag(name string) error

	AddAllAndCommit(message string, author *object.Signature) (string, error)
	Add(path string) (string, error)

	Commit(message string, author *object.Signature) (string, error)
	PushDefault() error

	GitRepository() *git.Repository
	Option() *git.CloneOptions
}

type repository struct {
	*git.Repository
	option *git.CloneOptions
}

func (r *repository) CheckoutBranch(name string) error {
	w, err := r.Worktree()
	if err != nil {
		return err
	}

	return w.Checkout(&git.CheckoutOptions{
		Branch: plumbing.NewBranchReferenceName(name),
	})
}

func (r *repository) CheckoutTag(name string) error {
	w, err := r.Worktree()
	if err != nil {
		return err
	}

	return w.Checkout(&git.CheckoutOptions{
		Branch: plumbing.NewTagReferenceName(name),
	})
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

func (r *repository) AddAllAndCommit(message string, author *object.Signature) (string, error) {
	_, err := r.Add(".")
	if err != nil {
		return "", err
	}

	return r.Commit(message, author)
}

func (r *repository) Add(path string) (string, error) {
	w, err := r.Worktree()
	if err != nil {
		return "", err
	}

	hash, err := w.Add(path)
	if err != nil {
		return "", err
	}

	result := ""
	if !hash.IsZero() {
		result = hash.String()
	}

	return result, err

}

func (r *repository) Commit(message string, author *object.Signature) (string, error) {
	w, err := r.Worktree()
	if err != nil {
		return "", err
	}

	hash, err := w.Commit(message, &git.CommitOptions{
		Author: author,
	})

	if err != nil {
		return "", err
	}

	result := ""
	if !hash.IsZero() {
		result = hash.String()
	}

	return result, err
}

func (r *repository) PushDefault() error {
	return r.Push(&git.PushOptions{
		Auth: r.option.Auth,
	})
}
