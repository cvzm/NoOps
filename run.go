package main

import (
	"nopos/src/util"
)

func main() {
	//TODO List
	// 1. 调用ansible-playbook, check err
	// 		err：校验ssh是否推送
	// 			推送 or 未知异常
	// 2. 服务状态 check
	//      err: 输出异常，尝试解决
	// 3. go 监控模块

	util.DockerCmd("docker run --rm -d redis", "nodes")
}
