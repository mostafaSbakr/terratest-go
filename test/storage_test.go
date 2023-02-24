package test

import (
	"github.com/gruntwork-io/terratest/modules/azure"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStorageAccount(t *testing.T) {
	exists := azure.ResourceGroupExists(t, resourceGroupName, subscriptionID)
	assert.True(t, exists)

	azure.StorageAccountExists(t, storageAccountName, resourceGroupName, subscriptionID)
	tier := azure.GetStorageAccountSkuTier(t, storageAccountName, resourceGroupName, subscriptionID)
	assert.Equal(t, storageAccountTier, tier)
}
