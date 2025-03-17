package main
/**
	✅ 推荐方式
	核心思想
	Functional Options 通过 定义一组 Option 函数，
	每个函数都修改一个配置项。最终，构造函数 NewXXX 接收

	优点
	✅ 可读性高，调用 NewServer 时，参数一目了然。 
	✅ 灵活扩展，可以随时添加新 Option，不影响现有代码。
 	✅ 避免默认值问题，未指定的参数保持默认值。
*/

//server 结构体
type Server struct {
	Host 	string
	Port 	int
	TLS 	bool
}

// Option 函数模式
type Option func(*Server)

// WithHost 设置 Host
func WithHost(host string) Option {
    return func(s *Server) {
        s.Host = host
    }
}

// WithPort 设置端口
func WithPort(port int) Option {
    return func(s *Server) {
        s.Port = port
    }
}

// WithTLS 启用 TLS
func WithTLS() Option {
    return func(s *Server) {
        s.TLS = true
    }
}

// NewServer 构造函数
func NewServerFunctional(opts ...Option) *Server {
    s := &Server{
        Host: "localhost", // 默认值
        Port: 8080,        // 默认端口
        TLS:  false,       // 默认不启用 TLS
    }
    for _, opt := range opts {
        opt(s)
    }
    return s
}