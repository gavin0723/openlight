// Author: lipixun
// Created Time : 二 10/18 21:51:04 2016
//
// File Name: main.go
// Description:
//	The openlight cli main entry
package main

import (
	"fmt"
	"github.com/ops-openlight/openlight/cli/build"
	"github.com/ops-openlight/openlight/cli/runner"
	"gopkg.in/urfave/cli.v1"
	"os"
)

var (
	buildBranch string
	buildCommit string
	buildTime   string
	buildTag    string
	buildGraph  string
	// The version string
	Version = fmt.Sprintf("Branch [%s] Commit [%s] Build Time [%s] Tag [%s]", buildBranch, buildCommit, buildTime, buildTag)
)

func main() {
	// The main entry
	// Create cli application
	app := cli.NewApp()
	app.Name = "op"
	app.Usage = "Openlight CLI"
	app.Version = Version
	app.EnableBashCompletion = false
	// Global flags
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "verbose",
			Usage: "Show verbose log (debug log)",
		},
		cli.StringFlag{
			Name:  "workdir",
			Usage: "The openlight workdir",
		},
	}
	// Add commands from modules
	for _, cmd := range build.GetCommand() {
		app.Commands = append(app.Commands, cmd)
	}
	for _, cmd := range runner.GetCommand() {
		app.Commands = append(app.Commands, cmd)
	}
	// Run it
	app.Run(os.Args)
}
