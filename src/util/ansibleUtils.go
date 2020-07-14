package util

/**
 * 用于远程执行Docker命令
 */
var hosts = "./hosts"

const (
	ansible         = "ansible"
	ansiblePlaybook = "ansible-playbook"
)

func InstallDocker(yaml string) error {
	return RunCmd(ansiblePlaybook, "-i", hosts, yaml)
}

func DockerCmd(cmd string, node string) error {
	return RunCmd(ansible, "-i", hosts, node, "-m", "shell", "-a", cmd)
}
