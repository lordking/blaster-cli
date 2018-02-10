package blastercli

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/lordking/blaster/common"
)

type (

	//Repo 脚手架代码存放仓库管理器
	Repo struct {
		Config *Config
	}
)

//Install 安装代码仓库
func (t *Repo) Install() error {
	var err error

	if err = t.CreateLocalRepo(); err != nil {
		return err
	}

	if err = t.AddPullingDest(defaultSpecsFileName); err != nil {
		return err
	}

	if err = t.Pull(); err != nil {
		return err
	}

	return nil
}

//CreateLocalRepo 创建本地仓库
func (t *Repo) CreateLocalRepo() error {
	var err error

	if _, err = ExecCommand("mkdir", "-p", t.Config.RepoPath); err != nil {
		return err
	}

	if _, err = ExecCommand("git", "init", t.Config.RepoPath); err != nil {
		return err
	}

	if _, err = ExecCommand("git", "-C", t.Config.RepoPath, "remote", "add", "origin", t.Config.RepoURL); err != nil {
		return err
	}

	return nil
}

//AddPullingDest  添加需要拉取的远程目标，文件或者文件夹。
func (t *Repo) AddPullingDest(dest string) error {

	infofile := fmt.Sprintf("%s/%s", t.Config.RepoPath, ".git/info/sparse-checkout")

	_, err := os.Stat(infofile)
	isExist := err == nil || os.IsExist(err)

	var data []byte
	if isExist {
		f, err := os.OpenFile(infofile, os.O_RDONLY, 0660)
		if err != nil {
			return err
		}

		r := bufio.NewReader(f)

		for {
			line, _, err := r.ReadLine()
			line = bytes.TrimSpace(line)

			if err == io.EOF {
				break
			}

			if dest == string(line) {
				return nil
			}

		}

	}

	data = []byte(dest + "\n")

	return common.AppendIntoFile(data, infofile)
}

//Pull 从远程仓库中拉取
func (t *Repo) Pull() error {
	var err error

	if _, err = ExecCommand("git", "-C", t.Config.RepoPath, "reset", "--hard"); err != nil {
		return err
	}

	if _, err = ExecCommand("git", "-C", t.Config.RepoPath, "pull", "origin", "master"); err != nil {
		return err
	}

	return nil
}

//NewRepo 对象实例化
func NewRepo(config *Config) Repo {
	return Repo{
		Config: config,
	}
}
