package aws

import (
	"log"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/apigateway"
	"github.com/hashicorp/terraform/helper/resource"
)

func (a *AWS) RemoveAPIGatewayRESTAPIs(prefix string) error {
	out, err := a.apigConn.GetRestApis(&apigateway.GetRestApisInput{})
	if err != nil {
		return err
	}

	log.Printf("[INFO] Found %d APIGateway REST APIs", len(out.Items))
	for _, i := range out.Items {
		if !strings.HasPrefix(*i.Name, prefix) {
			continue
		}

		err := resource.Retry(5*time.Minute, func() *resource.RetryError {
			_, err := a.apigConn.DeleteRestApi(&apigateway.DeleteRestApiInput{
				RestApiId: i.Id,
			})
			if err != nil {
				if awsErr, ok := err.(awserr.Error); ok && awsErr.Code() == "TooManyRequestsException" {
					return resource.RetryableError(err)
				}
				return resource.NonRetryableError(err)
			}

			log.Printf("[INFO] Deleted APIGateway REST API: %q", *i.Id)
			return nil
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *AWS) RemoveAPIGatewayAPIKeys(prefix string) error {
	out, err := a.apigConn.GetApiKeys(&apigateway.GetApiKeysInput{
		NameQuery: aws.String(prefix),
	})
	if err != nil {
		return err
	}

	log.Printf("[INFO] Found %d APIGateway API Keys", len(out.Items))
	for _, i := range out.Items {
		err := resource.Retry(5*time.Minute, func() *resource.RetryError {
			_, err := a.apigConn.DeleteApiKey(&apigateway.DeleteApiKeyInput{
				ApiKey: i.Id,
			})
			if err != nil {
				if awsErr, ok := err.(awserr.Error); ok && awsErr.Code() == "TooManyRequestsException" {
					return resource.RetryableError(err)
				}
				return resource.NonRetryableError(err)
			}

			log.Printf("[INFO] Deleted APIGateway API Key: %q", *i.Id)
			return nil
		})
		if err != nil {
			return err
		}
	}

	return nil
}
