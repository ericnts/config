package config

import "time"

type Project struct {
	Name        string        `yaml:"name"`
	HttpPrefix  string        `yaml:"httpPrefix"`  //请求路径前缀
	HttpPort    uint16        `yaml:"httpPort"`    // http端口
	RpcPort     uint16        `yaml:"rpcPort"`     // rpc端口
	TcpPort     uint16        `yaml:"tcpPort"`     // tcp端口
	LogResponse bool          `yaml:"logResponse"` // 是否打印http的response
	Swagger     bool          `yaml:"swagger"`     // 是否开启swagger
	Pprof       bool          `yaml:"pprof"`       // 是否开启性能调试
	S           time.Duration `yaml:"s"`
}
