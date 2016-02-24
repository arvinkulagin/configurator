package git

import (
	"archive/tar"
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"net/url"
	"os/exec"
)

const (
	_GIT     = "git"
	_ARCHIVE = "archive"
	_REMOTE  = "--remote"
	_HEAD    = "HEAD"
)

var (
	ErrNotFound = errors.New("File not found in repository")
)

func GetFile(repo, file string) ([]byte, error) {
	if _, err := url.Parse(repo); err != nil {
		return nil, err
	}
	cmd := exec.Command(_GIT, _ARCHIVE, _REMOTE, repo, _HEAD, file)
	var dataBuffer bytes.Buffer
	var errBuffer bytes.Buffer
	cmd.Stdout = &dataBuffer
	cmd.Stderr = &errBuffer
	err := cmd.Run()
	if err != nil {
		return nil, errors.New("Git exit with " + errBuffer.String())
	}
	reader := tar.NewReader(&dataBuffer)
	for {
		header, err := reader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if header.Name == file {
			return ioutil.ReadAll(reader)
		}
	}
	return nil, ErrNotFound
}
