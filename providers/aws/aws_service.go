// Copyright 2018 The Terraformer Authors.
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

package aws

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/aws/stscreds"
	"os"
	"regexp"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
)

type AWSService struct {
	terraform_utils.Service
}

var AWS_VARIABLE = regexp.MustCompile(`(\${[0-9A-Za-z:]+})`)

func (s *AWSService) generateConfig() (aws.Config, error) {
	config, e := s.buildBaseConfig()

	if e != nil {
		return config, e
	}
	if s.Verbose {
		config.LogLevel = aws.LogDebugWithHTTPBody
	}

	creds, e := config.Credentials.Retrieve(context.Background())

	if e != nil {
		return config, e
	}

	// terraform cannot ask for MFA token, so we need to pass STS session token, which might contain credentials with MFA requirement
	accessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	if accessKey == "" {
		os.Setenv("AWS_ACCESS_KEY_ID", creds.AccessKeyID)
		os.Setenv("AWS_SECRET_ACCESS_KEY", creds.SecretAccessKey)

		if creds.SessionToken != "" {
			os.Setenv("AWS_SESSION_TOKEN", creds.SessionToken)
		}
	}

	return config, nil
}

func (s *AWSService) buildBaseConfig() (aws.Config, error) {
	if s.GetArgs()["region"].(string) != "" {
		return external.LoadDefaultAWSConfig(
			external.WithRegion(s.GetArgs()["region"].(string)),
			external.WithMFATokenFunc(stscreds.StdinTokenProvider))
	} else {
		return external.LoadDefaultAWSConfig(external.WithMFATokenFunc(stscreds.StdinTokenProvider))
	}
}

// for CF interpolation and IAM Policy variables
func (_ *AWSService) escapeAwsInterpolation(str string) string {
	return AWS_VARIABLE.ReplaceAllString(str, "$$$1")
}
