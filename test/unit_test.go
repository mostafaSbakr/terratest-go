package test

import (
	"github.com/gruntwork-io/terratest/modules/terraform"
	"os"
	"path"
	"testing"
)

func TestPlan(t *testing.T) {
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../terraform/resource/keyvault",
	})

	tfPlanOutput := "terraform.tfplan"
	terraform.Init(t, terraformOptions)
	terraform.RunTerraformCommand(t, terraformOptions, terraform.FormatArgs(terraformOptions, "plan", "-out="+tfPlanOutput)...)
	f, err := os.Open(path.Join(terraformOptions.TerraformDir, tfPlanOutput))

	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

}
