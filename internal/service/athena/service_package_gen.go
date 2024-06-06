// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package athena

import (
	"context"

	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	athena_sdkv2 "github.com/aws/aws-sdk-go-v2/service/athena"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  dataSourceNamedQuery,
			TypeName: "aws_athena_named_query",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceDataCatalog,
			TypeName: "aws_athena_data_catalog",
			Name:     "Data Catalog",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceDatabase,
			TypeName: "aws_athena_database",
		},
		{
			Factory:  resourceNamedQuery,
			TypeName: "aws_athena_named_query",
		},
		{
			Factory:  resourcePreparedStatement,
			TypeName: "aws_athena_prepared_statement",
			Name:     "Prepared Statement",
		},
		{
			Factory:  resourceWorkGroup,
			TypeName: "aws_athena_workgroup",
			Name:     "WorkGroup",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Athena
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*athena_sdkv2.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws_sdkv2.Config))

	return athena_sdkv2.NewFromConfig(cfg, func(o *athena_sdkv2.Options) {
		if endpoint := config[names.AttrEndpoint].(string); endpoint != "" {
			o.BaseEndpoint = aws_sdkv2.String(endpoint)
		}
	}), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
