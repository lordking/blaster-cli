package blastercli

import (
	"path"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

var templates Templates

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
	templates = NewTemplates(&config, &repo)
}

//Test_template_CreateProject 测试之前，先查询模版，根据模版创建项目
func Test_template_CreateProject(t *testing.T) {

	var err error
	var errorMessage string

	err = templates.CreateProject("database-mango", "test")
	if err != nil {
		errorMessage = err.Error()
	}

	assert.Nil(t, err, errorMessage)
}
