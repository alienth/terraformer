package wavefront

import (
	"errors"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/GoogleCloudPlatform/terraformer/terraform_utils/provider_wrapper"
)

type WavefrontProvider struct {
	terraform_utils.Provider
}

func (p *WavefrontProvider) Init(args []string) error {
	return nil
}

func (p *WavefrontProvider) GetName() string {
	return "wavefront"
}

func (p *WavefrontProvider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{
		"provider": map[string]interface{}{
			"wavefront": map[string]interface{}{
				"version": provider_wrapper.GetProviderVersion(p.GetName()),
			},
		},
	}
}

func (WavefrontProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{}
}

func (p *WavefrontProvider) GetSupportedService() map[string]terraform_utils.ServiceGenerator {
	return map[string]terraform_utils.ServiceGenerator{
		"alert": &AlertGenerator{},
	}
}

func (p *WavefrontProvider) InitService(serviceName string, verbose bool) error {
	var isSupported bool
	if _, isSupported = p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New("wavefront: " + serviceName + " not supported service")
	}
	p.Service = p.GetSupportedService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetVerbose(verbose)
	p.Service.SetProviderName(p.GetName())

	return nil
}
