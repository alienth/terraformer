// Copyright 2019 The Terraformer Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package azure

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2019-08-01/network"
	"github.com/Azure/go-autorest/autorest"
	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/hashicorp/go-azure-helpers/authentication"
)

type NetworkInterfaceGenerator struct {
	AzureService
}

func (g NetworkInterfaceGenerator) createResources(interfaceListResultPage network.InterfaceListResultPage) []terraform_utils.Resource {
	var resources []terraform_utils.Resource
	for interfaceListResultPage.NotDone() {
		networkInterfaces := interfaceListResultPage.Values()
		for _, networkInterface := range networkInterfaces {
			resources = append(resources, terraform_utils.NewSimpleResource(
				*networkInterface.ID,
				*networkInterface.Name,
				"azurerm_network_interface",
				"azurerm",
				[]string{}))
		}
		if err := interfaceListResultPage.Next(); err != nil {
			log.Println(err)
			break
		}
	}
	return resources
}

func (g *NetworkInterfaceGenerator) InitResources() error {
	ctx := context.Background()
	interfacesClient := network.NewInterfacesClient(g.Args["config"].(authentication.Config).SubscriptionID)

	interfacesClient.Authorizer = g.Args["authorizer"].(autorest.Authorizer)
	output, err := interfacesClient.ListAll(ctx)
	if err != nil {
		return err
	}
	g.Resources = g.createResources(output)
	return nil
}
