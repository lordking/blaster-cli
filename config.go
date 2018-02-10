package blastercli

import (
	"fmt"
	"log"
	"os"

	"github.com/lordking/blaster/common"

	yaml "gopkg.in/yaml.v2"
)

type (

	//Config 系统配置定义
	Config struct {
		Base     string
		RepoPath string
		RepoURL  string
		Specs    string
	}

	//ConfigForm 从配置文件读取配置的表单
	ConfigForm struct {
		Repo  string `yaml:"repo"`
		Specs string `yaml:"specs"`
	}
)

var (
	defaultConfigDir     = ".blaster"
	defaultRepoPath      = "repo"
	defaultRepoURL       = "https://github.com/lordking/blaster-seed.git"
	defaultSpecsFileName = "specs.yaml"
	defaultSeedPath      = "github.com/lordking/blaster-seed" //模板种子的$GOPATH包位置
)

//CreateDefaultConfig 创建缺省的配置目录和配置文件
func CreateDefaultConfig() (Config, error) {
	var err error

	defaultConfigDir = fmt.Sprintf("%s/%s", os.Getenv("HOME"), defaultConfigDir)

	_, err = os.Stat(defaultConfigDir)
	if !os.IsNotExist(err) {
		return Config{}, fmt.Errorf("The Config directory exists: %s", defaultConfigDir)
	}

	log.Printf("It will create the config directory: %s\n", defaultConfigDir)
	if _, err = ExecCommand("mkdir", "-p", defaultConfigDir); err != nil {
		return Config{}, err
	}

	form := ConfigForm{
		Repo:  defaultRepoURL,
		Specs: defaultSpecsFileName,
	}
	var data []byte
	if data, err = yaml.Marshal(&form); err != nil {
		return Config{}, err
	}

	configFile := fmt.Sprintf("%s/%s", defaultConfigDir, "config.yaml")
	common.AppendIntoFile(data, configFile)

	return GetConfig(form, defaultConfigDir), nil
}

//GetConfig 从表单的配置获取系统配置
func GetConfig(form ConfigForm, base string) Config {
	return Config{
		Base:     base,
		RepoPath: fmt.Sprintf("%s/%s", base, defaultRepoPath),
		RepoURL:  form.Repo,
		Specs:    fmt.Sprintf("%s/%s/%s", base, defaultRepoPath, defaultSpecsFileName),
	}
}
