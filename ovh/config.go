package ovh

import (
	"os"
	"sort"
	"strings"

	"github.com/AdFabConnect/ovh-cli/utils"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/viper"
)

type apiConfig struct {
	Profile           string
	Endpoint          string
	ApplicationKey    string
	ApplicationSecret string
	ConsumerKey       string
}

func getHeaderToDisplay() []string {
	return []string{"Profile", "Application key", "Application secret", "Consumer secret", "Endpoint"}
}

func (c *apiConfig) toArrow() []string {
	return []string{c.Profile, c.ApplicationKey, c.ApplicationSecret, c.ConsumerKey, c.Endpoint}
}

func getAPIConfig(profile string) apiConfig {
	specificConfig := viper.Sub(profile)
	if specificConfig != nil {
		return apiConfig{
			Profile:           profile,
			Endpoint:          specificConfig.GetString("endpoint"),
			ApplicationKey:    specificConfig.GetString("application-key"),
			ApplicationSecret: specificConfig.GetString("application-secret"),
			ConsumerKey:       specificConfig.GetString("consumer-key"),
		}
	}
	return apiConfig{Profile: profile}
}

// DisplayCurrentOvhConfig display current OVH configuration used to communicate with OVH API
func DisplayCurrentOvhConfig() {
	profile := viper.GetString("profile")
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(getHeaderToDisplay())
	currentConfig := getAPIConfig(profile)
	table.Append(currentConfig.toArrow())
	table.Render()
}

// DisplayOvhConfig display all OVH configuration available to communicate with OVH API
func DisplayOvhConfig() {
	keys := viper.AllKeys()
	var profiles []string
	for _, key := range keys {
		items := strings.Split(key, ".")
		if len(items) > 1 {
			if !utils.StringSliceContains(profiles, items[0]) {
				profiles = append(profiles, items[0])
			}
		}
	}
	sort.Strings(profiles)
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(getHeaderToDisplay())
	for _, profile := range profiles {
		currentConfig := getAPIConfig(profile)
		table.Append(currentConfig.toArrow())
	}
	table.Render()
}
