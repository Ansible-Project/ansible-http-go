package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

func newServer(config *Config) {
	e := echo.New()
	e.GET("/version", getVersion)
	e.GET("/ansible/version", GetAnsibleVersion)
	e.POST("/ansible/playbook/run", GetAnsiblePlaybook)
	e.Run(standard.New(fmt.Sprintf(":%d", config.Port)))
}

func getVersion(c echo.Context) error {
	return c.String(http.StatusOK, version)
}
