package log

// 定义默认的常量
const (
	defaultBaseDirectoryName  = "logs"  // 日志根目录
	defaultInfoDirectoryName  = "info"  // info日志目录
	defaultWarnDirectoryName  = "warn"  // warn日志目录
	defaultErrorDirectoryName = "error" // error日志目录
	defaultInfoFileName       = "info"  // info日志文件
	defaultWarnFileName       = "warn"  // warn日志文件
	defaultErrorFileName      = "error" // error日志文件
	defaultLogFileMaxSize     = 128     // 日志文件大小，单位：MB
	defaultLogFileMaxBackups  = 180     // 日志文件保留个数 多于180个文件后，清理比价旧的日志
	defaultLogFileMaxAge      = 1       // 日志文件一天一切隔
	defaultLogFileCompress    = false   // 日志文件是否压缩
	defaultLogPrintTag        = false   // true:在终端和文件同时输出日志; false:只在文件输出日志

)

// Config 配置文件结构体
type Config struct {
	BaseDirectoryName  string
	InfoDirectoryName  string
	WarnDirectoryName  string
	ErrorDirectoryName string
	InfoFileName       string
	WarnFileName       string
	ErrorFileName      string
	LogFileMaxSize     int
	LogFileMaxBackups  int
	LogFileMaxAge      int
	LogFileCompress    bool
	LogPrintTag        bool
}

// Option 定义配置选项函数
type Option func(*Config)

// WithBaseDirectoryName 自定义日志根目录
func WithBaseDirectoryName(name string) Option {
	return func(c *Config) {
		c.BaseDirectoryName = name
	}
}

// WithInfoDirectoryName 自定义info日志目录
func WithInfoDirectoryName(name string) Option {
	return func(c *Config) {
		c.InfoDirectoryName = name
	}
}

// WithWarnDirectoryName 自定义warn日志目录
func WithWarnDirectoryName(name string) Option {
	return func(c *Config) {
		c.WarnDirectoryName = name
	}
}

// WithErrorDirectoryName 自定义error日志目录
func WithErrorDirectoryName(name string) Option {
	return func(c *Config) {
		c.ErrorDirectoryName = name
	}
}

// WithInfoFileName 自定义info文件名
func WithInfoFileName(name string) Option {
	return func(c *Config) {
		c.InfoFileName = name
	}
}

// WithWarnFileName 自定义warn文件名
func WithWarnFileName(name string) Option {
	return func(c *Config) {
		c.WarnFileName = name
	}
}

// WithErrorFileName 自定义error文件名
func WithErrorFileName(name string) Option {
	return func(c *Config) {
		c.ErrorFileName = name
	}
}

// WithLogFileMaxSize 自定义日志文件大小
func WithLogFileMaxSize(size int) Option {
	return func(c *Config) {
		c.LogFileMaxSize = size
	}
}

// WithLogFileMaxBackups 自定义日志文件保留个数
func WithLogFileMaxBackups(size int) Option {
	return func(c *Config) {
		c.LogFileMaxBackups = size
	}
}

// WithLogFileMaxAge 自定义日志文件切隔间隔
func WithLogFileMaxAge(size int) Option {
	return func(c *Config) {
		c.LogFileMaxAge = size
	}
}

// WithLogFileCompress 自定义日志文件是否压缩
func WithLogFileCompress(compress bool) Option {
	return func(c *Config) {
		c.LogFileCompress = compress
	}
}

// WithLogPrintTag 自定义日志输出标记位 true:在终端和文件同时输出日志; false:只在文件输出日志
func WithLogPrintTag(tag bool) Option {
	return func(c *Config) {
		c.LogPrintTag = tag
	}
}

// NewConfig 应用函数选项配置
func NewConfig(opts ...Option) Config {
	// 初始化默认值
	defaultConfig := Config{
		BaseDirectoryName:  defaultBaseDirectoryName,
		InfoDirectoryName:  defaultInfoDirectoryName,
		WarnDirectoryName:  defaultWarnDirectoryName,
		ErrorDirectoryName: defaultErrorDirectoryName,
		InfoFileName:       defaultInfoFileName,
		WarnFileName:       defaultWarnFileName,
		ErrorFileName:      defaultErrorFileName,
		LogFileMaxSize:     defaultLogFileMaxSize,
		LogFileMaxBackups:  defaultLogFileMaxBackups,
		LogFileMaxAge:      defaultLogFileMaxAge,
		LogFileCompress:    defaultLogFileCompress,
		LogPrintTag:        defaultLogPrintTag,
	}

	// 依次调用opts函数列表中的函数，为结构体成员赋值
	for _, opt := range opts {
		opt(&defaultConfig)
	}

	return defaultConfig
}
