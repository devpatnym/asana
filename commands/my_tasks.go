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
	for j, section_and_tasks := range sections_and_tasks {
		if section_and_tasks.Name == "Done" {
			continue
		}
		fmt.Println(section_and_tasks.Name)
		tasks := section_and_tasks.Tasks
		if (len(tasks) == 0) {
			fmt.Println(" -  - none")
		}
		for i, t := range tasks {
			fmt.Printf("%2d %2d %.80s %s \n", j, i, t.Name, t.Due_on)
			//fmt.Printf("%2d [ %10s ] %s\n", i, t.Due_on, t.Name)
		}
	}
}

func MyTasksGid(c *cli.Context) {
	fmt.Println(api.GetMyUserTaskListGid())
}
