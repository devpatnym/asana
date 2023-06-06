package commands

import (
	"os/exec"
	"strconv"

	"github.com/urfave/cli/v2"

	"github.com/thash/asana/api"
	"github.com/thash/asana/config"
	"github.com/thash/asana/utils"
)

func Browse(c *cli.Context) {
	taskGid := api.FindTaskGid(c.Args().First(), true)
	url := "https://app.asana.com/0/" + strconv.Itoa(config.Load().Workspace) + "/" + taskGid
	launcher, err := utils.BrowserLauncher()
	utils.Check(err)
	cmd := exec.Command(launcher, url)
	err = cmd.Start()
	utils.Check(err)
}
