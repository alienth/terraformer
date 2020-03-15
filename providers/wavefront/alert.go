package wavefront

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	wavefront "github.com/WavefrontHQ/go-wavefront-management-api"
)

type AlertGenerator struct {
	WavefrontService
}

func (g *AlertGenerator) createAlertResources(client *wavefront.Client) error {
	alertsClient := client.Alerts()
	alerts, err := alertsClient.Find(nil)
	if err != nil {
		return err
	}

	seen := make(map[string]bool)
	for _, alert := range alerts {
		if _, found := seen[*alert.ID]; found {
			continue
		}
		seen[*alert.ID] = true
		g.Resources = append(g.Resources, terraform_utils.NewSimpleResource(
			fmt.Sprintf("%s", *alert.ID),
			fmt.Sprintf("%s-%s", alert.Name, *alert.ID),
			"wavefront_alert",
			g.ProviderName,
			[]string{},
		))
	}

	return nil
}

func (g *AlertGenerator) InitResources() error {
	client, err := g.Client()
	if err != nil {
		return err
	}

	funcs := []func(*wavefront.Client) error{
		g.createAlertResources,
	}

	for _, f := range funcs {
		err := f(client)
		if err != nil {
			return err
		}
	}

	return nil
}

// func (g *AlertGenerator) PostConvertHook() error {
// 	for i, resource := range g.Resources {
// 		if resource.InstanceInfo.Type == "newrelic_alert_condition" {
// 			if resource.Item["violation_close_timer"] == "0" {
// 				delete(g.Resources[i].Item, "violation_close_timer")
// 			}
// 		}
// 	}

// 	return nil
// }
