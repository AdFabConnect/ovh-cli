package cmd

import (
	"fmt"
	"strconv"

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
		data := [][]string{}

		err := client.Get(applicationURI, &applicationsID)
		if err != nil {
			log.Errorln("Unable to list applications:", err)
			return
		}

		for _, applicationID := range applicationsID {
			if isQuiet {
				fmt.Println(applicationID)
			} else {
				application := fetchApplication(client, applicationID)
				data = append(data, []string{strconv.Itoa(applicationID), application.ApplicationKey, application.Name, application.Description, application.Status})
			}
		}
		if !isQuiet {
			table := utils.GetTable()
			table.SetHeader([]string{"Id", "Key", "Name", "Description", "Status"})

			for _, v := range data {
				table.Append(v)
			}
			table.Render()
		}
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
