package commands

import (
 	"fmt"
 	// "net/url"

	"github.com/urfave/cli/v2"

 	"github.com/thash/asana/api"
// 	"github.com/thash/asana/utils"
)

// const (
// 	CacheDuration = "5m"
// )

func MyTasks(c *cli.Context) {
	MyTasksFromAPI()
}

func MyTasksFromAPI() {
 	sections_and_tasks := api.MyTasks()
	for section_name, tasks := range sections_and_tasks {
		fmt.Println(section_name)
		for i, t := range tasks {
			fmt.Printf("%2d [ %10s ] %s\n", i, t.Due_on, t.Name)
		}
	}
}

func MyTasksGid(c *cli.Context) {
	fmt.Println(api.GetMyUserTaskListGid())
}
