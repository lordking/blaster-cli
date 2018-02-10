package blastercli

import (
	"bufio"
	"io"
	"os/exec"
)

//ExecCommand 执行系统指令
func ExecCommand(commandName string, arg ...string) ([]string, error) {
	cmd := exec.Command(commandName, arg...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	cmd.Start()

	reader := bufio.NewReader(stdout)

	lines := make([]string, 0)
	for {
		line, err := reader.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		lines = append(lines, line)
	}

	cmd.Wait()

	return lines, nil
}
