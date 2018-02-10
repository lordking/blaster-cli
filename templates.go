package blastercli

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/lordking/blaster/common"
	yaml "gopkg.in/yaml.v2"
)

type (

	//TemplateVO 单个模版对象
	TemplateVO struct {
		Name string
		Path string
	}

	//Templates 模版管理器
	Templates struct {
		Config *Config
		Repo   *Repo
	}
)

//CreateProject 创建项目
func (t *Templates) CreateProject(templateName, projectDirectory string) error {
	var err error

	//检测GOPATH是否存在
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		return fmt.Errorf("Not found GOPATH")
	}
	destPath := fmt.Sprintf("%s/src/%s", gopath, projectDirectory)

	//获取模板
	m, err := t.List()
	if err != nil {
		return err
	}

	tmpl := m[templateName]
	if tmpl.Name == "" {
		return fmt.Errorf("Not found template name: %s", templateName)
	}

	//下载模板
	log.Printf("It will create a new project in `%s` by `%s` ", destPath, templateName)
	err = t.Repo.AddPullingDest(tmpl.Path)
	if err != nil {
		return err
	}

	t.Repo.Pull()

	//复制模板
	srcPath := fmt.Sprintf("%s/%s", t.Repo.Config.RepoPath, tmpl.Path)
	if ExecCommand("cp", "-rf", srcPath, destPath); err != nil {
		return err
	}

	//更新模板内的import
	common.ReplaceInDirectory(destPath, ".go", fmt.Sprintf("%s/%s", defaultSeedPath, tmpl.Path), projectDirectory)

	return err
}

//List 列出所有模版的清单
func (t *Templates) List() (map[string]TemplateVO, error) {
	var err error
	var result map[string]TemplateVO

	b, err := ioutil.ReadFile(t.Config.Specs)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	err = yaml.Unmarshal(b, &m)
	if err != nil {
		return nil, err
	}

	result = make(map[string]TemplateVO)
	for k, v := range m {
		t := TemplateVO{
			Name: k,
			Path: v,
		}

		result[k] = t
	}

	return result, err
}

//NewTemplates 实例化
func NewTemplates(config *Config, repo *Repo) Templates {

	return Templates{
		Config: config,
		Repo:   repo,
	}
}
