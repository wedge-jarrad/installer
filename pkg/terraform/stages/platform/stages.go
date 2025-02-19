package platform

import (
	"github.com/openshift/installer/pkg/terraform"
	"github.com/openshift/installer/pkg/terraform/stages/aws"
	"github.com/openshift/installer/pkg/terraform/stages/azure"
	"github.com/openshift/installer/pkg/terraform/stages/compat"
	"github.com/openshift/installer/pkg/terraform/stages/gcp"
	"github.com/openshift/installer/pkg/terraform/stages/ibmcloud"
	awstypes "github.com/openshift/installer/pkg/types/aws"
	azuretypes "github.com/openshift/installer/pkg/types/azure"
	gcptypes "github.com/openshift/installer/pkg/types/gcp"
	ibmcloudtypes "github.com/openshift/installer/pkg/types/ibmcloud"
)

// StagesForPlatform returns the terraform stages to run to provision the infrastructure for the specified platform.
func StagesForPlatform(platform string) []terraform.Stage {
	switch platform {
	case awstypes.Name:
		return aws.PlatformStages
	case azuretypes.Name:
		return azure.PlatformStages
	case azuretypes.StackCloud.Name():
		return azure.StackPlatformStages
	case gcptypes.Name:
		return gcp.PlatformStages
	case ibmcloudtypes.Name:
		return ibmcloud.PlatformStages
	default:
		return compat.PlatformStages(platform)
	}
}
