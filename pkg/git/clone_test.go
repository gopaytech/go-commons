package git

import (
	"github.com/gopaytech/go-commons/pkg/strings"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestCloneOrOpenPublic(t *testing.T) {
	t.Run("clone public repository should success", func(t *testing.T) {
		repositoryUrl := "https://github.com/gopaytech/go-commons"

		destination, err := ioutil.TempDir("", strings.RandomAlphanumeric(5))
		assert.Nil(t, err)

		repository, ierr := CloneOrOpenPublic(repositoryUrl, destination)
		assert.Nil(t, ierr)
		assert.NotNil(t, repository)

		_ = os.Remove(destination)

		repositoryUrl = "git@github.com:gopaytech/go-commons.git"
		destination, err = ioutil.TempDir("", strings.RandomAlphanumeric(5))
		assert.Nil(t, err)

		repository, ierr = CloneOrOpenPublic(repositoryUrl, destination)
		t.Logf("error %s", ierr)
		assert.Nil(t, ierr)
		assert.NotNil(t, repository)

		_ = os.Remove(destination)
	})

	t.Run("open existing local repository should success", func(t *testing.T) {
		repositoryUrl := "https://github.com/gopaytech/go-commons"

		destination, err := ioutil.TempDir("", strings.RandomAlphanumeric(5))
		assert.Nil(t, err)

		// make sure clone success
		repository, ierr := CloneOrOpenPublic(repositoryUrl, destination)
		assert.Nil(t, ierr)
		assert.NotNil(t, repository)

		// make sure directory still exists
		stat, err := os.Stat(destination)
		assert.Nil(t, err)
		assert.True(t, stat.IsDir())

		// make sure able to open exist local repository
		repository, ierr = CloneOrOpenPublic(repositoryUrl, destination)
		assert.Nil(t, ierr)
		assert.NotNil(t, repository)
		_ = os.Remove(destination)
	})

	t.Run("clone non existence repository should failed ", func(t *testing.T) {
		repositoryUrl := "https://github.com/nonexistance/repo"

		destination, err := ioutil.TempDir("", strings.RandomAlphanumeric(5))
		assert.Nil(t, err)

		repository, ierr := CloneOrOpenPublic(repositoryUrl, destination)
		assert.NotNil(t, ierr)
		assert.Nil(t, repository)

		_ = os.Remove(destination)
	})

}
