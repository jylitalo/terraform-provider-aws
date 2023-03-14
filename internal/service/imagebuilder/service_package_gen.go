// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package imagebuilder

import (
	"context"

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
			Factory:  DataSourceComponent,
			TypeName: "aws_imagebuilder_component",
		},
		{
			Factory:  DataSourceComponents,
			TypeName: "aws_imagebuilder_components",
		},
		{
			Factory:  DataSourceContainerRecipe,
			TypeName: "aws_imagebuilder_container_recipe",
		},
		{
			Factory:  DataSourceContainerRecipes,
			TypeName: "aws_imagebuilder_container_recipes",
		},
		{
			Factory:  DataSourceDistributionConfiguration,
			TypeName: "aws_imagebuilder_distribution_configuration",
		},
		{
			Factory:  DataSourceDistributionConfigurations,
			TypeName: "aws_imagebuilder_distribution_configurations",
		},
		{
			Factory:  DataSourceImage,
			TypeName: "aws_imagebuilder_image",
		},
		{
			Factory:  DataSourceImagePipeline,
			TypeName: "aws_imagebuilder_image_pipeline",
		},
		{
			Factory:  DataSourceImagePipelines,
			TypeName: "aws_imagebuilder_image_pipelines",
		},
		{
			Factory:  DataSourceImageRecipe,
			TypeName: "aws_imagebuilder_image_recipe",
		},
		{
			Factory:  DataSourceImageRecipes,
			TypeName: "aws_imagebuilder_image_recipes",
		},
		{
			Factory:  DataSourceInfrastructureConfiguration,
			TypeName: "aws_imagebuilder_infrastructure_configuration",
		},
		{
			Factory:  DataSourceInfrastructureConfigurations,
			TypeName: "aws_imagebuilder_infrastructure_configurations",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceComponent,
			TypeName: "aws_imagebuilder_component",
		},
		{
			Factory:  ResourceContainerRecipe,
			TypeName: "aws_imagebuilder_container_recipe",
		},
		{
			Factory:  ResourceDistributionConfiguration,
			TypeName: "aws_imagebuilder_distribution_configuration",
		},
		{
			Factory:  ResourceImage,
			TypeName: "aws_imagebuilder_image",
		},
		{
			Factory:  ResourceImagePipeline,
			TypeName: "aws_imagebuilder_image_pipeline",
		},
		{
			Factory:  ResourceImageRecipe,
			TypeName: "aws_imagebuilder_image_recipe",
		},
		{
			Factory:  ResourceInfrastructureConfiguration,
			TypeName: "aws_imagebuilder_infrastructure_configuration",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.ImageBuilder
}

var ServicePackage = &servicePackage{}
