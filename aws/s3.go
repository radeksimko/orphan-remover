package aws

import (
	"log"
	"strings"

	awsSDK "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/s3"
)

func (a *AWS) RemoveS3Buckets(prefix string) error {
	out, err := a.s3Conn.ListBuckets(&s3.ListBucketsInput{})
	if err != nil {
		return err
	}

	log.Printf("[INFO] Found %d S3 buckets", len(out.Buckets))

	for _, b := range out.Buckets {
		if !strings.HasPrefix(*b.Name, prefix) {
			continue
		}

		// Empty bucket first
		a.s3Conn.DeleteObject(&s3.DeleteObjectInput{})
		_, err = a.s3Conn.DeleteObject(&s3.DeleteObjectInput{
			Bucket: b.Name,
			Key:    awsSDK.String("/"),
		})
		if err != nil {
			if awsErr, ok := err.(awserr.Error); ok && awsErr.Code() == "BucketRegionError" {
				log.Printf("[WARN] Skipping %q: %s", *b.Name, awsErr.Message())
				continue
			}
			return err
		}
		log.Printf("[INFO] Deleted bucket: %q", *b.Name)
	}

	return nil
}
