package main

import (
	"log"

	"github.com/radeksimko/orphan-remover/aws"
)

func main() {
	// TODO: Create CLI iface "$0 <aws|google|azure|...> <waf|s3|...> [prefix]"
	a, err := aws.New("us-west-2")
	if err != nil {
		log.Fatal(err)
	}

	err = a.RemoveWaf()
	if err != nil {
		log.Fatal(err)
	}

	err = a.RemoveS3Buckets("tf-test-bucket")
	if err != nil {
		log.Fatal(err)
	}

	err = a.RemoveAPIGatewayRESTAPIs("test")
	if err != nil {
		log.Fatal(err)
	}

	err = a.RemoveAPIGatewayAPIKeys("foo")
	if err != nil {
		log.Fatal(err)
	}
}
