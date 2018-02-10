package blastercli

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var config Config

//Test_Config_CreateDefaultConfig 测试之前先删除默认的.blaster
func Test_Config_CreateDefaultConfig(t *testing.T) {

	var err error
	var errorMessage string
	base := fmt.Sprintf("%s/%s", os.Getenv("HOME"), defaultConfigDir)
	expectedConfig := Config{
		Base:     base,
		RepoPath: fmt.Sprintf("%s/%s", base, defaultRepoPath),
		RepoURL:  defaultRepoURL,
		Specs:    fmt.Sprintf("%s/%s/%s", base, defaultRepoPath, defaultSpecsFileName),
	}

	result, err := CreateDefaultConfig()
	if err != nil {
		errorMessage = err.Error()
	}

	assert.Nil(t, err, errorMessage)
	assert.Equal(t, expectedConfig, result, "Result not match expected object.")

	if err == nil {
		ExecCommand("rm", "-rf", base)
		t.Logf("result: %s", result)
	}

}
