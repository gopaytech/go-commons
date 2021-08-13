package git

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"github.com/gopaytech/go-commons/pkg/dir"
	"github.com/gopaytech/go-commons/pkg/file"
	cryptoSsh "golang.org/x/crypto/ssh"
)

// DefaultSshAuth this func will fetch default private key ~/.ssh/id_rsa
// default private key can be override with SSH_DEFAULT_PRIVATE_KEY
// default private key passphrase can be override with SSH_DEFAULT_PRIVATE_KEY_PASSPHRASE
func DefaultSshAuth() (transport.AuthMethod, error) {
	defaultKeyLocation := os.Getenv("SSH_DEFAULT_PRIVATE_KEY")
	defaultKeyPassphrase := os.Getenv("SSH_DEFAULT_PRIVATE_KEY_PASSPHRASE")

	if defaultKeyLocation == "" {
		defaultKeyLocation = dir.Home(".ssh", "id_rsa")
	}

	if !file.FileExists(defaultKeyLocation) {
		return nil, fmt.Errorf(
			"default key %s is not exists. "+
				"specify it on SSH_DEFAULT_PRIVATE_KEY env variable. "+
				"if your key has passhrase you can specify it on SSH_DEFAULT_PRIVATE_KEY_PASSPHRASE",
			defaultKeyLocation,
		)
	}

	return SshAuth(defaultKeyLocation, defaultKeyPassphrase)
}

func BasicAuth(username string, password string) transport.AuthMethod {
	return &http.BasicAuth{
		Username: username,
		Password: password,
	}
}

func SshAuth(keyFile string, passphrase string) (transport.AuthMethod, error) {
	bytes, err := ioutil.ReadFile(keyFile)
	if err != nil {
		return nil, err
	}

	return SshAuthByte(bytes, passphrase)
}

func SshAuthByte(byteFile []byte, passphrase string) (transport.AuthMethod, error) {
	var signer cryptoSsh.Signer
	if len(passphrase) > 0 {
		signerKey, err := cryptoSsh.ParsePrivateKeyWithPassphrase(byteFile, []byte(passphrase))
		if err != nil {
			return nil, err
		}
		signer = signerKey
	} else {
		signerKey, err := cryptoSsh.ParsePrivateKey(byteFile)
		if err != nil {
			return nil, err
		}
		signer = signerKey
	}

	return &ssh.PublicKeys{
		User:   "git",
		Signer: signer,
		HostKeyCallbackHelper: ssh.HostKeyCallbackHelper{
			HostKeyCallback: cryptoSsh.InsecureIgnoreHostKey(),
		},
	}, nil
}

type CloneOrOpenPublicFunc func(repositoryUrl string, destination string) (Repository, error)

// CloneOrOpenPublic fetch repository without authentication
// for `git:` url format, this func will fetch default private key ~/.ssh/id_rsa
// default private key can be override with SSH_DEFAULT_PRIVATE_KEY
// default private key passphrase can be override with SSH_DEFAULT_PRIVATE_KEY_PASSPHRASE
func CloneOrOpenPublic(repositoryUrl string, destination string) (Repository, error) {
	var option *git.CloneOptions
	if strings.HasPrefix(repositoryUrl, "git@") {
		defaultSshAuth, err := DefaultSshAuth()
		if err != nil {
			return nil, err
		}

		option = &git.CloneOptions{
			URL:  repositoryUrl,
			Auth: defaultSshAuth,
		}

	} else {
		option = &git.CloneOptions{
			URL: repositoryUrl,
		}
	}

	return CloneOrOpenWithOption(destination, option)
}

type CloneOrOpenFunc func(url string, auth transport.AuthMethod, destination string) (Repository, error)

func CloneOrOpen(url string, auth transport.AuthMethod, destination string) (Repository, error) {
	option := &git.CloneOptions{
		Auth: auth,
		URL:  url,
	}

	return CloneOrOpenWithOption(destination, option)
}

type CloneOrOpenWithOptionFunc func(destination string, option *git.CloneOptions) (Repository, error)

func CloneOrOpenWithOption(destination string, option *git.CloneOptions) (Repository, error) {
	gitRepository, err := git.PlainClone(destination, false, option)

	if err != nil {
		if errors.Is(err, git.ErrRepositoryAlreadyExists) {
			gitRepository, err = git.PlainOpen(destination)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	result := &repository{
		Repository: gitRepository,
		option:     option,
	}
	return result, nil
}

type ClonePublicFunc func(url string, destination string) (Repository, error)

// ClonePublic fetch repository without authentication
// for `git:` url format, this func will fetch default private key ~/.ssh/id_rsa
// default private key can be override with SSH_DEFAULT_PRIVATE_KEY
// default private key passphrase can be override with SSH_DEFAULT_PRIVATE_KEY_PASSPHRASE
func ClonePublic(url string, destination string) (Repository, error) {
	var option *git.CloneOptions
	if strings.HasPrefix(url, "git@") {
		defaultSshAuth, err := DefaultSshAuth()
		if err != nil {
			return nil, err
		}

		option = &git.CloneOptions{
			URL:  url,
			Auth: defaultSshAuth,
		}

	} else {
		option = &git.CloneOptions{
			URL: url,
		}
	}

	return CloneWithOption(destination, option)
}

type CloneFunc func(url string, destination string, auth transport.AuthMethod) (Repository, error)

func Clone(url string, destination string, auth transport.AuthMethod) (Repository, error) {
	option := &git.CloneOptions{
		Auth: auth,
		URL:  url,
	}

	return CloneWithOption(destination, option)
}

type CloneBranchFunc func(url string, destination string, auth transport.AuthMethod, branch string) (Repository, error)

func CloneBranch(url string, destination string, auth transport.AuthMethod, branch string) (Repository, error) {
	option := &git.CloneOptions{
		Auth:          auth,
		URL:           url,
		ReferenceName: plumbing.NewBranchReferenceName(branch),
	}

	return CloneWithOption(destination, option)
}

type CloneTagFunc func(url string, destination string, auth transport.AuthMethod, tag string) (Repository, error)

func CloneTag(url string, destination string, auth transport.AuthMethod, tag string) (Repository, error) {
	option := &git.CloneOptions{
		Auth:          auth,
		URL:           url,
		ReferenceName: plumbing.NewTagReferenceName(tag),
	}

	return CloneWithOption(destination, option)
}

type CloneWithOptionFunc func(destination string, option *git.CloneOptions) (Repository, error)

func CloneWithOption(destination string, option *git.CloneOptions) (Repository, error) {
	if !dir.IsEmpty(destination) {
		return nil, fmt.Errorf("destination %s is not empty, not exists, or not accessible", destination)
	}

	gitRepository, err := git.PlainClone(destination, false, option)
	if err != nil {
		return nil, err
	}

	result := &repository{
		Repository: gitRepository,
		option:     option,
	}
	return result, nil
}
