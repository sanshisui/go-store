package main

/**
	核心思想
	通过定义一个 Config 结构体，用户可填充所需字段，然后传入 NewXXX 构造函数。

	优点
	✅ 可读性高，调用 NewServer 时，参数一目了然。 
	✅ 灵活扩展，可以随时添加新 Option，不影响现有代码。
 	✅ 避免默认值问题，未指定的参数保持默认值。
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