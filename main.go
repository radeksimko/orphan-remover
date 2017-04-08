package main

import (
	"log"

	"github.com/radeksimko/orphan-remover/aws"
)

func main() {
	// TODO: Create CLI iface "$0 <aws|google|azure|...> <waf|s3|...>"
	a, err := aws.New("us-east-1")
	if err != nil {
		log.Fatal(err)
	}

	err = a.RemoveWaf()
	if err != nil {
		log.Fatal(err)
	}
}
