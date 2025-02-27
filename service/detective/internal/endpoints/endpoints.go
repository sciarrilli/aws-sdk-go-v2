// Code generated by smithy-go-codegen DO NOT EDIT.

package endpoints

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/endpoints"
	"regexp"
)

// Options is the endpoint resolver configuration options
type Options struct {
	DisableHTTPS bool
}

// Resolver Detective endpoint resolver
type Resolver struct {
	partitions endpoints.Partitions
}

// ResolveEndpoint resolves the service endpoint for the given region and options
func (r *Resolver) ResolveEndpoint(region string, options Options) (endpoint aws.Endpoint, err error) {
	if len(region) == 0 {
		return endpoint, &aws.MissingRegionError{}
	}

	opt := endpoints.Options{
		DisableHTTPS: options.DisableHTTPS,
	}
	return r.partitions.ResolveEndpoint(region, opt)
}

// New returns a new Resolver
func New() *Resolver {
	return &Resolver{
		partitions: defaultPartitions,
	}
}

var defaultPartitions = endpoints.Partitions{
	{
		ID: "aws",
		Defaults: endpoints.Endpoint{
			Hostname:          "api.detective.{region}.amazonaws.com",
			Protocols:         []string{"https"},
			SignatureVersions: []string{"v4"},
		},
		RegionRegex:    regexp.MustCompile("^(us|eu|ap|sa|ca|me|af)\\-\\w+\\-\\d+$"),
		IsRegionalized: true,
		Endpoints: endpoints.Endpoints{
			"af-south-1":     endpoints.Endpoint{},
			"ap-east-1":      endpoints.Endpoint{},
			"ap-northeast-1": endpoints.Endpoint{},
			"ap-northeast-2": endpoints.Endpoint{},
			"ap-south-1":     endpoints.Endpoint{},
			"ap-southeast-1": endpoints.Endpoint{},
			"ap-southeast-2": endpoints.Endpoint{},
			"ca-central-1":   endpoints.Endpoint{},
			"eu-central-1":   endpoints.Endpoint{},
			"eu-north-1":     endpoints.Endpoint{},
			"eu-south-1":     endpoints.Endpoint{},
			"eu-west-1":      endpoints.Endpoint{},
			"eu-west-2":      endpoints.Endpoint{},
			"eu-west-3":      endpoints.Endpoint{},
			"me-south-1":     endpoints.Endpoint{},
			"sa-east-1":      endpoints.Endpoint{},
			"us-east-1":      endpoints.Endpoint{},
			"us-east-1-fips": endpoints.Endpoint{
				Hostname: "api.detective-fips.us-east-1.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "us-east-1",
				},
			},
			"us-east-2": endpoints.Endpoint{},
			"us-east-2-fips": endpoints.Endpoint{
				Hostname: "api.detective-fips.us-east-2.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "us-east-2",
				},
			},
			"us-west-1": endpoints.Endpoint{},
			"us-west-1-fips": endpoints.Endpoint{
				Hostname: "api.detective-fips.us-west-1.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "us-west-1",
				},
			},
			"us-west-2": endpoints.Endpoint{},
			"us-west-2-fips": endpoints.Endpoint{
				Hostname: "api.detective-fips.us-west-2.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "us-west-2",
				},
			},
		},
	},
	{
		ID: "aws-cn",
		Defaults: endpoints.Endpoint{
			Hostname:          "api.detective.{region}.amazonaws.com.cn",
			Protocols:         []string{"https"},
			SignatureVersions: []string{"v4"},
		},
		RegionRegex:    regexp.MustCompile("^cn\\-\\w+\\-\\d+$"),
		IsRegionalized: true,
	},
	{
		ID: "aws-iso",
		Defaults: endpoints.Endpoint{
			Hostname:          "api.detective.{region}.c2s.ic.gov",
			Protocols:         []string{"https"},
			SignatureVersions: []string{"v4"},
		},
		RegionRegex:    regexp.MustCompile("^us\\-iso\\-\\w+\\-\\d+$"),
		IsRegionalized: true,
	},
	{
		ID: "aws-iso-b",
		Defaults: endpoints.Endpoint{
			Hostname:          "api.detective.{region}.sc2s.sgov.gov",
			Protocols:         []string{"https"},
			SignatureVersions: []string{"v4"},
		},
		RegionRegex:    regexp.MustCompile("^us\\-isob\\-\\w+\\-\\d+$"),
		IsRegionalized: true,
	},
	{
		ID: "aws-us-gov",
		Defaults: endpoints.Endpoint{
			Hostname:          "api.detective.{region}.amazonaws.com",
			Protocols:         []string{"https"},
			SignatureVersions: []string{"v4"},
		},
		RegionRegex:    regexp.MustCompile("^us\\-gov\\-\\w+\\-\\d+$"),
		IsRegionalized: true,
	},
}
