package aws

import (
	"fmt"
	"log"

	awsSDK "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/aws/aws-sdk-go/service/apigateway"
	"github.com/aws/aws-sdk-go/service/s3"
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
		apigConn: apigateway.New(session.New(awsConfig)),
		s3Conn:   s3.New(session.New(awsConfig)),
		wafConn:  waf.New(session.New(awsConfig)),
		region:   region,
	}, nil
}

type AWS struct {
	apigConn *apigateway.APIGateway
	wafConn  *waf.WAF
	s3Conn   *s3.S3
	region   string
}
