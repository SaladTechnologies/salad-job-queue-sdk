package workers

import (
	"github.com/saladtechnologies/salad-cloud-imds-sdk-go/pkg/saladcloudimdssdk"
	"github.com/saladtechnologies/salad-cloud-imds-sdk-go/pkg/saladcloudimdssdkconfig"
)

// Creates a SaladCloud Instance Metadata Service (IMDS) client.
func newMetadataClient(addr string) *saladcloudimdssdk.SaladCloudImdsSdk {
	config := saladcloudimdssdkconfig.NewConfig()
	config.SetBaseUrl(addr)
	return saladcloudimdssdk.NewSaladCloudImdsSdk(config)
}
