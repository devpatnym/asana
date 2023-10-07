package commands

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"regexp"
	"strconv"

	"github.com/urfave/cli/v2"

	"github.com/thash/asana/api"
	"github.com/thash/asana/utils"
)

const (
	CacheDuration = "5m"
)

func Tasks(c *cli.Context) {
	if c.Bool("no-cache") {
		fromAPI(false)
	} else {
		if utils.Older(CacheDuration, utils.CacheFile()) || c.Bool("refresh") {
			fromAPI(true)
		} else {
			txt, err := ioutil.ReadFile(utils.CacheFile())
			if err == nil {
				lines := regexp.MustCompile("\n").Split(string(txt), -1)
				for _, line := range lines {
					if len(line) < 1 {
						continue
					}
					format(line)
				}
			} else {
				fromAPI(true)
			}
		}
	}
}

func fromAPI(saveCache bool) {
	tasks := api.Tasks(url.Values{}, false)
	if saveCache {
		cache(tasks)
	}
	for i, t := range tasks {
		memberships := membershipsToSectionNames(t.Memberships)
		printfFromFields(i, t.Name, t.Due_on, memberships)
	}
}

func membershipsToSectionNames(memberships []api.Membership_t) string {
	outstr := ""
	for _, m := range memberships {
		outstr = outstr + m.Section.Name
	}
	return outstr
}

func printfFromFields(index int, name, due_on, memberships string) {
		if (regexp.MustCompile("Deployed|release").MatchString(memberships)) {
			fmt.Println("-")
		} else {
			fmt.Printf("%2d %s [%s] {%s}\n", index, name, due_on, memberships)
		}
}

func cache(tasks []api.Task_t) {
	f, _ := os.Create(utils.CacheFile())
	defer f.Close()
	for i, t := range tasks {
		f.WriteString(strconv.Itoa(i) + ":")
		f.WriteString(t.Gid + ":")
		f.WriteString(membershipsToSectionNames(t.Memberships) + ":")
		f.WriteString(t.Due_on + ":")
		f.WriteString(t.Name + "\n")
	}
}

func format(line string) {
	dateRegexp := "[0-9]{4}-[0-9]{2}-[0-9]{2}"

	index := regexp.MustCompile("^[0-9]*").FindString(line)
	index2, _ := strconv.Atoi(index)
	line = regexp.MustCompile("^[0-9]*:").ReplaceAllString(line, "") // remove index
	line = regexp.MustCompile("^[0-9]*:").ReplaceAllString(line, "") // remove task_id
	mems := regexp.MustCompile("^[^:]+").FindString(line)
	line = regexp.MustCompile("^[^:]+:").ReplaceAllString(line, "") // remove memberships
	date := regexp.MustCompile("^" + dateRegexp).FindString(line)
	line = regexp.MustCompile("^("+dateRegexp+")?:").ReplaceAllString(line, "") // remove date
	printfFromFields(index2, line, date, mems)
}
