package main

/**
	核心思想
	通过定义一个 Config 结构体，用户可填充所需字段，然后传入 NewXXX 构造函数。

	优点
	✅ 直观，所有参数集中在 Config 结构体中。 
	✅ 适合配置文件（如 JSON/YAML）。 
	❌ 可扩展性较差，添加新选项时，所有调用 NewServer 的代码都需要更新。
*/

//ServerConfig 结构体
type ServerConfig struct {
	Host 	string
	Port 	int
	TLS 	bool
}

// Server 结构体
type ServerStruct struct {
    config ServerConfig
}

// NewServer 构造函数
func NewServerStruct(config ServerConfig) *ServerStruct {
    if config.Host == "" {
        config.Host = "localhost"
    }
    if config.Port == 0 {
        config.Port = 8080
    }
    return &ServerStruct{config: config}
}