// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package apigateway

import (
	"context"

	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	apigateway_sdkv2 "github.com/aws/aws-sdk-go-v2/service/apigateway"
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
			Factory:  dataSourceAPIKey,
			TypeName: "aws_api_gateway_api_key",
			Name:     "API Key",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  dataSourceAuthorizer,
			TypeName: "aws_api_gateway_authorizer",
			Name:     "Authorizer",
		},
		{
			Factory:  dataSourceAuthorizers,
			TypeName: "aws_api_gateway_authorizers",
			Name:     "Authorizers",
		},
		{
			Factory:  dataSourceDomainName,
			TypeName: "aws_api_gateway_domain_name",
			Name:     "Domain Name",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  dataSourceExport,
			TypeName: "aws_api_gateway_export",
			Name:     "Export",
		},
		{
			Factory:  dataSourceResource,
			TypeName: "aws_api_gateway_resource",
			Name:     "Resource",
		},
		{
			Factory:  dataSourceRestAPI,
			TypeName: "aws_api_gateway_rest_api",
			Name:     "REST API",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  dataSourceSDK,
			TypeName: "aws_api_gateway_sdk",
			Name:     "SDK",
		},
		{
			Factory:  DataSourceVPCLink,
			TypeName: "aws_api_gateway_vpc_link",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceAccount,
			TypeName: "aws_api_gateway_account",
			Name:     "Account",
		},
		{
			Factory:  resourceAPIKey,
			TypeName: "aws_api_gateway_api_key",
			Name:     "API Key",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  resourceAuthorizer,
			TypeName: "aws_api_gateway_authorizer",
			Name:     "Authorizer",
		},
		{
			Factory:  resourceBasePathMapping,
			TypeName: "aws_api_gateway_base_path_mapping",
			Name:     "Base Path Mapping",
		},
		{
			Factory:  resourceClientCertificate,
			TypeName: "aws_api_gateway_client_certificate",
			Name:     "Client Certificate",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  resourceDeployment,
			TypeName: "aws_api_gateway_deployment",
			Name:     "Deployment",
		},
		{
			Factory:  resourceDocumentationPart,
			TypeName: "aws_api_gateway_documentation_part",
			Name:     "Documentation Part",
		},
		{
			Factory:  resourceDocumentationVersion,
			TypeName: "aws_api_gateway_documentation_version",
			Name:     "Documentation Version",
		},
		{
			Factory:  resourceDomainName,
			TypeName: "aws_api_gateway_domain_name",
			Name:     "Domain Name",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  resourceGatewayResponse,
			TypeName: "aws_api_gateway_gateway_response",
			Name:     "Gateway Response",
		},
		{
			Factory:  resourceIntegration,
			TypeName: "aws_api_gateway_integration",
			Name:     "Integration",
		},
		{
			Factory:  resourceIntegrationResponse,
			TypeName: "aws_api_gateway_integration_response",
			Name:     "Integration Response",
		},
		{
			Factory:  resourceMethod,
			TypeName: "aws_api_gateway_method",
			Name:     "Method",
		},
		{
			Factory:  resourceMethodResponse,
			TypeName: "aws_api_gateway_method_response",
			Name:     "Method Response",
		},
		{
			Factory:  resourceMethodSettings,
			TypeName: "aws_api_gateway_method_settings",
			Name:     "Method Settings",
		},
		{
			Factory:  resourceModel,
			TypeName: "aws_api_gateway_model",
			Name:     "Model",
		},
		{
			Factory:  resourceRequestValidator,
			TypeName: "aws_api_gateway_request_validator",
			Name:     "Request Validator",
		},
		{
			Factory:  resourceResource,
			TypeName: "aws_api_gateway_resource",
			Name:     "Resource",
		},
		{
			Factory:  resourceRestAPI,
			TypeName: "aws_api_gateway_rest_api",
			Name:     "REST API",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  resourceRestAPIPolicy,
			TypeName: "aws_api_gateway_rest_api_policy",
			Name:     "REST API Policy",
		},
		{
			Factory:  ResourceStage,
			TypeName: "aws_api_gateway_stage",
			Name:     "Stage",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceUsagePlan,
			TypeName: "aws_api_gateway_usage_plan",
			Name:     "Usage Plan",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceUsagePlanKey,
			TypeName: "aws_api_gateway_usage_plan_key",
		},
		{
			Factory:  ResourceVPCLink,
			TypeName: "aws_api_gateway_vpc_link",
			Name:     "VPC Link",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.APIGateway
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*apigateway_sdkv2.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws_sdkv2.Config))

	return apigateway_sdkv2.NewFromConfig(cfg, func(o *apigateway_sdkv2.Options) {
		if endpoint := config["endpoint"].(string); endpoint != "" {
			o.BaseEndpoint = aws_sdkv2.String(endpoint)
		}
	}), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
