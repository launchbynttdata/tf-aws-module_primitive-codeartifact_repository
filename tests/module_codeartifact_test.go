package test

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/codeartifact"
	"github.com/gruntwork-io/terratest/modules/logger"
	"github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
	"github.com/stretchr/testify/assert"
)

const (
	base            = "../examples/"
	testVarFileName = "/test.tfvars"
)

var standardTags = map[string]string{
	"provisioner": "Terraform",
}

func TestCodeArtifact(t *testing.T) {
	t.Parallel()
	stage := test_structure.RunTestStage

	files, err := os.ReadDir(base)
	assert.NoError(t, err)
	basePath, _ := filepath.Abs(base)
	for _, file := range files {
		dir := filepath.Join(basePath, file.Name()) //base + file.Name()
		if file.IsDir() {
			defer stage(t, "teardown_codeartifact", func() { tearDownCodeArtifact(t, dir) })
			stage(t, "setup_codeartifact", func() { setupCodeArtifactTest(t, dir) })
			stage(t, "test_codeartifact", func() { testCodeArtifact(t, dir) })
		}
	}
}

func setupCodeArtifactTest(t *testing.T, dir string) {

	terraformOptions := &terraform.Options{
		TerraformDir: dir,
		VarFiles:     []string{dir + testVarFileName},
		NoColor:      true,
		Logger:       logger.Discard,
	}

	test_structure.SaveTerraformOptions(t, dir, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

}

func testCodeArtifact(t *testing.T, dir string) {
	terraformOptions := test_structure.LoadTerraformOptions(t, dir)
	terraformOptions.Logger = logger.Discard

	expectedPatternARN := "^arn:aws:codeartifact:[a-z0-9-]+:[0-9]{12}:[a-z0-9-]+/.+$"

	actualID := terraform.Output(t, terraformOptions, "id")
	assert.NotEmpty(t, actualID, "ARN ID is empty")
	assert.Regexp(t, expectedPatternARN, actualID, "ID does not match expected pattern")
	actualARN := terraform.Output(t, terraformOptions, "arn")
	assert.NotEmpty(t, actualARN, "ARN is empty")
	assert.Regexp(t, expectedPatternARN, actualARN, "ARN does not match expected pattern")

	sess, err := session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	})
	assert.NoError(t, err, "can't connect to aws")
	client := codeartifact.New(sess)
	tfvarsFullPath := dir + testVarFileName

	domainName, err := terraform.GetVariableAsStringFromVarFileE(t, tfvarsFullPath, "domain")
	assert.NoError(t, err)
	repositoryName, err := terraform.GetVariableAsStringFromVarFileE(t, tfvarsFullPath, "repository")
	assert.NoError(t, err)

	input := &codeartifact.DescribeRepositoryInput{
		Domain:     aws.String(domainName),
		Repository: aws.String(repositoryName),
	}

	result, err := client.DescribeRepository(input)
	assert.NoError(t, err, "The expected code artifact was not found")

	repository := result.Repository

	expectedName, err := terraform.GetVariableAsStringFromVarFileE(t, tfvarsFullPath, "repository")
	assert.NoError(t, err)
	actualName := repository.Name
	assert.Equal(t, expectedName, *actualName, "Repository Name does not match")
	checkTagsMatch(t, tfvarsFullPath, actualARN, client)
}

func checkTagsMatch(t *testing.T, tfvarsFullPath string, actualARN string, client *codeartifact.CodeArtifact) {
	expectedTags, err := terraform.GetVariableAsMapFromVarFileE(t, tfvarsFullPath, "tags")
	assert.NoError(t, err)

	input := &codeartifact.ListTagsForResourceInput{
		ResourceArn: aws.String(actualARN),
	}
	result, err := client.ListTagsForResource(input) //client.ListTags(context.TODO(), &acmpca.ListTagsInput{CertificateAuthorityArn: aws.String(actualARN)})
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

func tearDownCodeArtifact(t *testing.T, dir string) {
	terraformOptions := test_structure.LoadTerraformOptions(t, dir)
	terraformOptions.Logger = logger.Discard
	terraform.Destroy(t, terraformOptions)
}
