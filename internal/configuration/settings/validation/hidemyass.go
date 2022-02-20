package validation

import (
	"github.com/qdm12/gluetun/internal/models"
)

func HideMyAssCountryChoices(servers []models.HideMyAssServer) (choices []string) {
	choices = make([]string, len(servers))
	for i := range servers {
		choices[i] = servers[i].Country
	}
	return makeUnique(choices)
}

func HideMyAssRegionChoices(servers []models.HideMyAssServer) (choices []string) {
	choices = make([]string, len(servers))
	for i := range servers {
		choices[i] = servers[i].Region
	}
	return makeUnique(choices)
}

func HideMyAssCityChoices(servers []models.HideMyAssServer) (choices []string) {
	choices = make([]string, len(servers))
	for i := range servers {
		choices[i] = servers[i].City
	}
	return makeUnique(choices)
}

func HideMyAssHostnameChoices(servers []models.HideMyAssServer) (choices []string) {
	choices = make([]string, len(servers))
	for i := range servers {
		choices[i] = servers[i].Hostname
	}
	return makeUnique(choices)
}