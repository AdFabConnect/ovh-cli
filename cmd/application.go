package cmd

import (
	"sort"
	"strconv"
	"sync"

	ovhClient "github.com/AdFabConnect/ovh-cli/ovh"
	"github.com/AdFabConnect/ovh-cli/utils"
	"github.com/ovh/go-ovh/ovh"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const applicationURI string = "/me/api/application/"

type PartialApplication struct {
	Name           string `json:"name"`
	ApplicationKey string `json:"applicationKey"`
	ApplicationID  int    `json:"applicationId"`
	Description    string `json:"description"`
	Status         string `json:"status"`
}

func getHeaderToDisplay() []string {
	return []string{"Id", "Key", "Name", "Description", "Status"}
}

func (a *PartialApplication) toArrow() []string {
	return []string{strconv.Itoa(a.ApplicationID), a.ApplicationKey, a.Name, a.Description, a.Status}
}

func fetchApplication(client *ovh.Client, applicationID int) PartialApplication {
	var application PartialApplication
	err := client.Get(applicationURI+strconv.Itoa(applicationID), &application)
	if err != nil {
		log.Errorln("Unable to get application details:", err)
	}
	return application
}

var listApplicationCmd = &cobra.Command{
	Use:   "list",
	Short: "list applications",
	Run: func(cmd *cobra.Command, args []string) {
		applicationsID := []int{}
		client := ovhClient.GetOvhClient()
		isQuiet, _ := cmd.Flags().GetBool("quiet")

		err := client.Get(applicationURI, &applicationsID)
		if err != nil {
			log.Errorln("Unable to list applications:", err)
			return
		}
		table := utils.GetTable()

		sort.Ints(applicationsID)
		if isQuiet {
			for _, applicationID := range applicationsID {
				table.Append([]string{strconv.Itoa(applicationID)})
			}
			table.SetHeaderLine(false)
		} else {
			applications := make([]PartialApplication, len(applicationsID))
			var wg sync.WaitGroup
			for index, applicationID := range applicationsID {
				wg.Add(1)
				go func(index int, applicationID int) {
					defer wg.Done()
					client := ovhClient.GetOvhClient()
					application := fetchApplication(client, applicationID)
					applications[index] = application
				}(index, applicationID)
			}
			wg.Wait()
			table.SetHeader(getHeaderToDisplay())
			for _, application := range applications {
				table.Append(application.toArrow())
			}
		}
		table.Render()
	},
}

var applicationCmd = &cobra.Command{
	Use:   "application",
	Short: "Manipulate applications",
}

func init() {
	rootCmd.AddCommand(applicationCmd)
	applicationCmd.AddCommand(listApplicationCmd)
	listApplicationCmd.Flags().BoolP("quiet", "q", false, "Only display application id")
}
