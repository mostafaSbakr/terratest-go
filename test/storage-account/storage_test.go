package storage_account

import (
	"github.com/gruntwork-io/terratest/modules/azure"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	"testing"
)

const subscriptionID string = "6b651eb2-7854-43ce-89b8-5490b1a07783"

func TestStorageAccount(t *testing.T) {
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../../terraform/module",
	})

	//defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)

	exists := azure.ResourceGroupExists(t, "terratesting", subscriptionID)
	assert.True(t, exists)
}
