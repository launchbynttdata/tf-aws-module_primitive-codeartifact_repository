package common

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/codeartifact"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/launchbynttdata/lcaf-component-terratest/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var standardTags = map[string]string{
	"provisioner": "Terraform",
}

func TestCodeArtifact(t *testing.T, ctx types.TestContext) {

	t.Run("TestARNAndIDPatternMatches", func(t *testing.T) {
		checkARNFormat(t, ctx)
		checkARNIDFormat(t, ctx)
	})

	t.Run("TestingExpectedArtifactAndRepo", func(t *testing.T) {
		testCodeArtifact(t, ctx)
	})

	t.Run("TestingValidTags", func(t *testing.T) {
		checkTagsMatch(t, ctx)
	})
}

func checkARNIDFormat(t *testing.T, ctx types.TestContext) {
	expectedPatternARN := "^arn:aws:codeartifact:[a-z0-9-]+:[0-9]{12}:[a-z0-9-]+/.+$"

	actualID := terraform.Output(t, ctx.TerratestTerraformOptions(), "id")
	assert.NotEmpty(t, actualID, "ARN ID is empty")
	assert.Regexp(t, expectedPatternARN, actualID, "ID does not match expected pattern")
}

func checkARNFormat(t *testing.T, ctx types.TestContext) {
	expectedPatternARN := "^arn:aws:codeartifact:[a-z0-9-]+:[0-9]{12}:[a-z0-9-]+/.+$"
	actualARN := terraform.Output(t, ctx.TerratestTerraformOptions(), "arn")
	assert.NotEmpty(t, actualARN, "ARN is empty")
	assert.Regexp(t, expectedPatternARN, actualARN, "ARN does not match expected pattern")
}

func testCodeArtifact(t *testing.T, ctx types.TestContext) {
	input := &codeartifact.DescribeRepositoryInput{
		Domain:     aws.String(terraform.Output(t, ctx.TerratestTerraformOptions(), "domain")),
		Repository: aws.String(terraform.Output(t, ctx.TerratestTerraformOptions(), "repository")),
	}

	client := GetCodeArtifactClient(t)

	result, err := client.DescribeRepository(context.TODO(), input)
	assert.NoError(t, err, "The expected code artifact was not found")

	repository := result.Repository

	expectedName := terraform.Output(t, ctx.TerratestTerraformOptions(), "repository")
	assert.NoError(t, err)
	actualName := repository.Name
	assert.Equal(t, expectedName, *actualName, "Repository Name does not match")
}

func checkTagsMatch(t *testing.T, ctx types.TestContext) {
	expectedTags := terraform.OutputMap(t, ctx.TerratestTerraformOptions(), "tags_all")
	actualARN := terraform.Output(t, ctx.TerratestTerraformOptions(), "arn")
	client := GetCodeArtifactClient(t)

	input := &codeartifact.ListTagsForResourceInput{
		ResourceArn: aws.String(actualARN),
	}
	result, err := client.ListTagsForResource(context.TODO(), input) //client.ListTags(context.TODO(), &acmpca.ListTagsInput{CertificateAuthorityArn: aws.String(actualARN)})
	assert.NoError(t, err, "Failed to retrieve tags from AWS")
	// convert AWS Tag[] to map so we can compare
	actualTags := map[string]string{}
	for _, tag := range result.Tags {
		actualTags[*tag.Key] = *tag.Value
	}

	// add the standard tags and the resource_name tag to the expected tags
	for k, v := range standardTags {
		expectedTags[k] = v
	}
	// expectedTags["resource_name"] = "g" //actualTags["resource_name"]
	// assert.Equal(t, expectedTags["provisione"], actualTags["provisioner"], "Artifact Name does not match")
	assert.True(t, reflect.DeepEqual(actualTags, expectedTags), fmt.Sprintf("tags did not match, expected: %v\nactual: %v", expectedTags, actualTags))
}

func GetAWSConfig(t *testing.T) (cfg aws.Config) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	require.NoErrorf(t, err, "unable to load SDK config, %v", err)
	return cfg
}

func GetCodeArtifactClient(t *testing.T) *codeartifact.Client {
	codeartifactClient := codeartifact.NewFromConfig(GetAWSConfig(t))
	return codeartifactClient
}
