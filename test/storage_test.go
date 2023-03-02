package test

import (
	"github.com/gruntwork-io/terratest/modules/azure"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestStorageAccount(t *testing.T) {
	exists := azure.ResourceGroupExists(t, resourceGroupName, ARM_SUBSCRIPTION_ID)
	assert.True(t, exists)

	subscriptionId := os.Getenv(ARM_SUBSCRIPTION_ID)

	azure.StorageAccountExists(t, storageAccountName, resourceGroupName, subscriptionId)
	tier := azure.GetStorageAccountSkuTier(t, storageAccountName, resourceGroupName, ARM_SUBSCRIPTION_ID)
	assert.Equal(t, storageAccountTier, tier)
}
