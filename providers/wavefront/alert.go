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

	for _, alert := range alerts {
		g.Resources = append(g.Resources, terraform_utils.NewSimpleResource(
			fmt.Sprintf("%d", alert.ID),
			fmt.Sprintf("%s", alert.Name),
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
