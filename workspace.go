package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
)

func CreateWorkspace(name string, uuid uuid.UUID) (string, error) {
	dirname := fmt.Sprintf("%s/%s-%s", config.WorkDir, name, uuid)
	if err := os.MkdirAll(dirname, 0777); err != nil {
		return "", err
	}
	return dirname, nil
}

func GitCloneWorkspace(dirname string, destpath string, repositoryAddr string, branchName string) (string, error) {
	// create clone dest directory
	sourceDir := path.Clean(fmt.Sprintf("%s/%s", dirname, destpath))
	if err := os.MkdirAll(sourceDir, 0777); err != nil {
		return "", err
	}
	// git clone
	out, err := exec.Command("git", "clone", "--depth", "1", "-b", branchName, repositoryAddr, sourceDir).CombinedOutput()
	if err != nil {
		return "", errors.Wrap(err, string(out))
	}

	fmt.Printf("[success] git clone --depth 1 -b %s %s %s\n", branchName, repositoryAddr, sourceDir)
	return sourceDir, nil
}

func CreateUuid() uuid.UUID {
	return uuid.NewV4()
}
