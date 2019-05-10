package ovh

import (
	"sort"
	"strings"

	"github.com/AdFabConnect/ovh-cli/utils"
	"github.com/spf13/viper"
)

type APIConfig struct {
	Profile           string
	Endpoint          string
	ApplicationKey    string
	ApplicationSecret string
	ConsumerKey       string
}

func getHeaderToDisplay() []string {
	return []string{"Profile", "Application key", "Application secret", "Consumer secret", "Endpoint"}
}

func (c *APIConfig) toArrow() []string {
	return []string{c.Profile, c.ApplicationKey, c.ApplicationSecret, c.ConsumerKey, c.Endpoint}
}

func GetCurrentAPIConfig() APIConfig {
	profile := viper.GetString("profile")
	return getAPIConfig(profile)
}

func getAPIConfig(profile string) APIConfig {
	specificConfig := viper.Sub(profile)
	if specificConfig != nil {
		return APIConfig{
			Profile:           profile,
			Endpoint:          specificConfig.GetString("endpoint"),
			ApplicationKey:    specificConfig.GetString("application-key"),
			ApplicationSecret: specificConfig.GetString("application-secret"),
			ConsumerKey:       specificConfig.GetString("consumer-key"),
		}
	}
	return APIConfig{Profile: profile}
}

// DisplayCurrentOvhConfig display current OVH configuration used to communicate with OVH API
func DisplayCurrentOvhConfig() {
	profile := viper.GetString("profile")
	table := utils.GetTable()
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
	table := utils.GetTable()
	table.SetHeader(getHeaderToDisplay())
	for _, profile := range profiles {
		currentConfig := getAPIConfig(profile)
		table.Append(currentConfig.toArrow())
	}
	table.Render()
}
