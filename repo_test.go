package blastercli

import (
	"path"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

var repo Repo

func init() {

	//获取配置作为测试数据
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath("$HOME/.blaster/")
	viper.AddConfigPath(".blaster/")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {

		form := ConfigForm{}
		if err := viper.Unmarshal(&form); err != nil {
			panic(err)
		}

		configDir := path.Dir(viper.ConfigFileUsed())
		config = GetConfig(form, configDir)

	} else {
		config, err = CreateDefaultConfig()
	}

	repo = NewRepo(&config)
}

func Test_Repo_Install(t *testing.T) {
	var err error
	var errorMessage string

	err = repo.Install()
	if err != nil {
		errorMessage = err.Error()
	}

	assert.Nil(t, err, errorMessage)
}

func Test_Repo_CreateLocalRepo(t *testing.T) {

	var err error
	var errorMessage string

	err = repo.CreateLocalRepo()
	if err != nil {
		errorMessage = err.Error()
	}

	assert.Nil(t, err, errorMessage)

}

//Test_AddPullingDest 启动该测试用例之前，先执行Test_CreateLocalRepo
func Test__Repo_AddPullingDest(t *testing.T) {

	var err error
	var errorMessage string

	err = repo.AddPullingDest(defaultSpecsFileName)
	if err != nil {
		errorMessage = err.Error()
	}

	assert.Nil(t, err, errorMessage)
}

//Test_Pull 启动该测试用例之前，先执行Test_AddPullingDest
func Test_Repo_Pull(t *testing.T) {
	var err error
	var errorMessage string

	err = repo.Pull()
	if err != nil {
		errorMessage = err.Error()
	}

	assert.Nil(t, err, errorMessage)
}
