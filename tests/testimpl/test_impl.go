package testimpl

import (
	"context"
	"os"
	"strings"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/launchbynttdata/lcaf-component-terratest/types"
	"github.com/stretchr/testify/assert"
)

func TestComposableComplete(t *testing.T, ctx types.TestContext) {
	testApplicationInsightsComplete(t, ctx)
}

func TestApplicationInsightsComplete(t *testing.T, ctx types.TestContext) {
	testApplicationInsightsComplete(t, ctx)
}

func testApplicationInsightsComplete(t *testing.T, ctx types.TestContext) {
	subscriptionID := os.Getenv("ARM_SUBSCRIPTION_ID")
	if len(subscriptionID) == 0 {
		t.Fatal("ARM_SUBSCRIPTION_ID is not set in the environment variables ")
	}
	credential, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		t.Fatalf("Unable to get credentials: %v\n", err)
	}

	options := arm.ClientOptions{
		ClientOptions: azcore.ClientOptions{
			Cloud: cloud.AzurePublic,
		},
	}

	clientFactory, err := armapplicationinsights.NewClientFactory(subscriptionID, credential, &options)
	if err != nil {
		t.Fatalf("Unable to get clientFactory: %v\n", err)

	}

	componentsClient := clientFactory.NewComponentsClient()

	expectedRgName := terraform.OutputContext(t, context.Background(), ctx.TerratestTerraformOptions(), "resource_group_name")
	expectedAppInsightsName := terraform.OutputContext(t, context.Background(), ctx.TerratestTerraformOptions(), "app_insights_name")
	expectedAppInsightsId := terraform.OutputContext(t, context.Background(), ctx.TerratestTerraformOptions(), "app_insights_id")

	res, err := componentsClient.Get(context.Background(), expectedRgName, expectedAppInsightsName, nil)
	if err != nil {
		t.Fatalf("Error occurred while getting resource: %v\n", err)
	}

	t.Run("AppInsightsExists", func(t *testing.T) {
		assert.Equal(t, strings.ToLower(expectedAppInsightsId), strings.ToLower(*res.ID), "Ids must match")
	})
}
