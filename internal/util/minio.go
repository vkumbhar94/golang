package util

import (
	"context"
	"fmt"

	awsv2 "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	credsv2 "github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Config struct {
	Bucket     string
	Region     string
	Endpoint   string
	AccessID   string
	AccessKey  string
	Encryption string
}

var MinioConfig = S3Config{
	Bucket:     "localdev",
	Region:     "us-west-1",
	Endpoint:   "http://localhost:9000",
	AccessID:   "minio",
	AccessKey:  "minio_key",
	Encryption: "none",
}

func (c S3Config) ToAwsV2() awsv2.Config {
	ac, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic(fmt.Errorf("fail to load config %w", err))
	}
	if c.Endpoint != "" {
		customResolver := awsv2.EndpointResolverWithOptionsFunc(func(service string, region string, options ...interface{}) (awsv2.Endpoint, error) {
			if service == s3.ServiceID {
				return awsv2.Endpoint{
					PartitionID:       "aws",
					URL:               c.Endpoint,
					SigningRegion:     c.Region,
					HostnameImmutable: true,
				}, nil
			}
			return awsv2.Endpoint{}, &awsv2.EndpointNotFoundError{}
		})

		ac, err = config.LoadDefaultConfig(context.TODO(), config.WithEndpointResolverWithOptions(customResolver))
		if err != nil {
			panic(fmt.Errorf("fail to load config %w", err))
		}
	}
	if c.Region != "" {
		ac.Region = c.Region
	}
	if c.AccessID != "" && c.AccessKey != "" {
		ac.Credentials = awsv2.NewCredentialsCache(credsv2.NewStaticCredentialsProvider(c.AccessID, c.AccessKey, ""))
	}
	return ac
}

func S3Client(cfg awsv2.Config) *s3.Client {
	return s3.NewFromConfig(cfg)
}
func MinIOS3Client() *s3.Client {
	return s3.NewFromConfig(MinioConfig.ToAwsV2())
}
