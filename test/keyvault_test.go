package test

import (
	"fmt"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKeyVaultSecret(t *testing.T) {
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../terraform/resource/keyvault",
	})

	terraform.InitAndApply(t, terraformOptions)
	value := terraform.Output(t, terraformOptions, "value")
	version := terraform.Output(t, terraformOptions, "version")

	fmt.Println("Secret", value)
	fmt.Println("Secret", version)
	assert.Equal(t, "<go4U+2QgRw2yRZ5)f-97O]Pzvy[]C", value)
}
