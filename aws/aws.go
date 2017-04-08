package aws

import (
	"fmt"
	"log"

	awsSDK "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/waf"
)

func New(region string) (*AWS, error) {
	providers := []credentials.Provider{
		&credentials.EnvProvider{},
		&credentials.SharedCredentialsProvider{},
	}

	creds := credentials.NewChainCredentials(providers)
	cp, err := creds.Get()
	if err != nil {
		return nil, fmt.Errorf("Error loading credentials: %s", err)
	}

	log.Printf("[INFO] AWS Auth provider used: %q", cp.ProviderName)

	awsConfig := &awsSDK.Config{
		Credentials: creds,
		Region:      awsSDK.String(region),
	}

	return &AWS{
		wafConn: waf.New(session.New(awsConfig)),
		region:  region,
	}, nil
}

type AWS struct {
	wafConn *waf.WAF
	region  string
}
