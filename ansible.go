package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"

	"github.com/labstack/echo"
)

func GetAnsibleVersion(c echo.Context) error {
	fmt.Println("GetAnsibleVersion")
	out, _ := exec.Command("ansible", "--version").Output()
	text := fmt.Sprintf("%s", out)
	return c.String(http.StatusOK, text)
}

func GetAnsiblePlaybook(c echo.Context) error {
	// create uuid
	uuid := CreateUuid()
	// get parameters
	playbook := c.QueryParam("playbook")
	inventory := c.QueryParam("inventory")
	limit := c.QueryParam("limit")
	tags := c.QueryParam("tags")
	skiptags := c.QueryParam("skiptags")
	extravars := c.QueryParam("extravars")
	verbose := c.QueryParam("verbose")
	dir := c.QueryParam("dir")
	branchName := c.QueryParam("branch")
	// TODO: validate parameters
	// init parameters
	args := make([]string, 0)
	if playbook == "" {
		return c.String(http.StatusBadRequest, fmt.Sprintf("playbook is required"))
	}
	args = append(args, playbook)
	args = append(args, "-i")
	if inventory != "" {
		args = append(args, inventory)
	} else if config.DefaultInventory != "" {
		args = append(args, config.DefaultInventory)
	} else {
		return c.String(http.StatusBadRequest, fmt.Sprintf("inventory is required"))
	}
	if limit != "" {
		args = append(args, "-l")
		args = append(args, limit)
	}
	if tags != "" {
		args = append(args, "-t")
		args = append(args, tags)
	}
	if skiptags != "" {
		args = append(args, "--skip-tags")
		args = append(args, skiptags)
	}
	if extravars != "" {
		args = append(args, "--extra-vars")
		args = append(args, extravars)
	}
	if verbose != "" {
		args = append(args, verbose)
	}	else if config.DefaultVerbose != "" {
		args = append(args, config.DefaultVerbose)
	}
	if branchName == "" {
		branchName = config.DefaultBranch
	}

	// TODO: logging parameters
	fmt.Printf("GetAnsiblePlaybook [%s]: %s, dir=%s, branch=%s", uuid, args, dir, branchName)

	// prepare workspace
	fmt.Println("CreateWorkspace")
	workspaceDir, err := CreateWorkspace("ansible", uuid)
	defer func() {
		_ = os.RemoveAll(workspaceDir)
	}()
	if err != nil {
		fmt.Sprintf("workspace create error: %s\n", err)
		return c.String(http.StatusInternalServerError, fmt.Sprintf("workspace create error: %s", err))
	}
	fmt.Printf("GitClone %s to %s\n", branchName, workspaceDir)
	sourceDir, err := GitCloneWorkspace(workspaceDir, "", config.RepositoryUrl, branchName)
	if err != nil {
		// TODO: ssh err ->500, branch not found ->404
		fmt.Printf("git clone error: %s\n", err)
		return c.String(http.StatusInternalServerError, fmt.Sprintf("git clone error: %s", err))
	}

	// exec ansible
	fmt.Println("execute Ansible")
	cmd := exec.Command("ansible-playbook", args...)
	cmd.Dir = fmt.Sprintf("%s/%s", sourceDir, dir)
	out, err := cmd.CombinedOutput()
	// TODO: exit code != 0 => error handling
	text := fmt.Sprintf("%s", out)
	fmt.Printf("result: %s", out)
	// TODO: rm workdir <- all error
	if err != nil {
		fmt.Printf("ansible execution error: %s\n", err)
		return c.String(http.StatusInternalServerError, text)
	} else {
		fmt.Println("ansible execution complete")
		return c.String(http.StatusOK, text)
	}

}
