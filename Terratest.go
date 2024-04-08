//Below is a simple example of a Terraform test using Terratest for the previously defined Terraform configuration:

package test

import (
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformAwsInstanceExample(t *testing.T) {
	terraformOptions := &terraform.Options{
		// The path to where your Terraform code is located
		TerraformDir: "../terraform/aws_instance_example",
	}

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// Run `terraform init` and `terraform apply`. Fail the test if there are any errors.
	terraform.InitAndApply(t, terraformOptions)

	// Get the instance ID from the output
	instanceID := terraform.Output(t, terraformOptions, "instance_id")

	// Verify that the instance exists
	instance := aws.GetInstanceState(t, instanceID, "us-west-2")

	// Check that the instance is in the running state
	assert.Equal(t, "running", instance.State.Name)

	// Check that the instance has the correct tags
	expectedTags := map[string]string{
		"Name": "ExampleInstance",
	}
	for key, value := range expectedTags {
		assert.Equal(t, value, instance.Tags[key])
	}
}




//Make sure to replace "../terraform/aws_instance_example" with the correct path to your Terraform configuration directory.

//This test does the following:

//Initializes and applies the Terraform configuration.
//Retrieves the instance ID from the Terraform output.
//Checks that the instance is in the running state.
//Verifies that the instance has the correct tags.
//To use this test, you'll need to have Terratest installed (go get -u github.com/gruntwork-io/terratest/modules/...) and configured properly. Additionally, ensure that you have AWS credentials configured and that the necessary permissions are granted to create and manage EC2 instances.