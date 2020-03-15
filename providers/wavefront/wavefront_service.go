package wavefront

import (
	"errors"
	"fmt"
	"os"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	wavefront "github.com/WavefrontHQ/go-wavefront-management-api"
)

type WavefrontService struct {
	terraform_utils.Service
}

func (s *WavefrontService) Client() (*wavefront.Client, error) {
	token := os.Getenv("WAVEFRONT_TOKEN")

	if token == "" {
		err := errors.New("No WAVEFRONT_TOKEN environment set")
		return nil, err
	}

	config := &wavefront.Config{
		Address: os.Getenv("WAVEFRONT_ADDRESS"),
		Token:   token,
	}

	client, err := wavefront.NewClient(config)
	if err != nil {
		return nil, fmt.Errorf("failed to configure Wavefront Client %s", err)
	}
	return client, nil
}
