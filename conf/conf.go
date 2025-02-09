package conf

import (
	"embed"
	_ "embed"
	"os"
	"path/filepath"
	"sync"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gopkg.in/yaml.v2"
)

//go:embed dev/config.yaml

//go:embed prod/config.yaml

var f embed.FS

var (
	conf *Config
	once sync.Once
)

type Config struct {
	Env string

	Hertz    *Hertz      `yaml:"hertz"`
	MySQL    []*MySQL    `yaml:"mysql"`
	Redis    []*Redis    `yaml:"redis"`
	Metrics  *Metrics    `yaml:"metrics"`
	Sentry   *Sentry     `yaml:"sentry"`
	Contract []*Contract `yaml:"contract"`
	Kafka    *Kafka      `yaml:"kafka"`
}

type MySQL struct {
	Name            string `yaml:"name"`
	DSN             string `yaml:"dsn"`
	MaxOpenConns    int    `yaml:"max_open_conns"`
	MaxIdleConns    int    `yaml:"max_idle_conns"`
	ConnMaxIdleTime int    `yaml:"conn_max_idle_time"`
	ConnMaxLifetime int    `yaml:"conn_max_life_time"`
}

type Redis struct {
	Name     string `yaml:"name"`
	Address  string `yaml:"address"`
	Password string `yaml:"password"`
	Username string `yaml:"username"`
	DB       int    `yaml:"db"`
	PoolSize int    `yaml:"pool_size"`
}

type Hertz struct {
	Address         string `yaml:"address"`
	HttpsAddress    string `yaml:"https_address"`
	EnableHttps     bool   `yaml:"enable_https"`
	EnablePprof     bool   `yaml:"enable_pprof"`
	EnableGzip      bool   `yaml:"enable_gzip"`
	EnableAccessLog bool   `yaml:"enable_access_log"`
	LogLevel        string `yaml:"log_level"`
	LogFileName     string `yaml:"log_file_name"`
}

type Metrics struct {
	Address string `yaml:"address"`
}

type Sentry struct {
	Dsn string `yaml:"dsn"`
}

type Contract struct {
	Name         string   `yaml:"name"`
	Address      string   `yaml:"address"`
	RpcUrl       string   `yaml:"rpc_url"`
	BackupRpcUrl []string `yaml:"backup_rpc_url"`
	Type         int      `yaml:"type"`
	MoveAccount  string   `yaml:"move_account"`
	MoveModule   string   `yaml:"move_module"`
}

type Kafka struct {
	ServerAddr      string `yaml:"server_addr"`
	CheckInTopic    string `yaml:"check_in_topic"`
	PredictionTopic string `yaml:"prediction_topic"`
}

// GetConf gets configuration instance
func GetConf(name string) *Config {
	once.Do(func() {
		initConf(name)
	})
	return conf
}

func initConf(name string) {
	confFileRelPath := filepath.Join(filepath.Join(GetEnv(), name+".yaml"))
	content, err := f.ReadFile(confFileRelPath)
	if err != nil {
		content, err = f.ReadFile("../" + confFileRelPath)
		if err != nil {
			panic(err)
		}
	}

	conf = new(Config)
	err = yaml.Unmarshal(content, conf)
	if err != nil {
		hlog.Error("parse yaml error - %v", err)
		panic(err)
	}

	conf.Env = GetEnv()
}

func GetEnv() string {
	e := os.Getenv("GO_ENV")
	if len(e) == 0 {
		return "dev"
	}
	return e
}

func LogLevel() hlog.Level {
	if conf == nil {
		return hlog.LevelDebug
	}
	level := conf.Hertz.LogLevel
	switch level {
	case "trace":
		return hlog.LevelTrace
	case "debug":
		return hlog.LevelDebug
	case "info":
		return hlog.LevelInfo
	case "notice":
		return hlog.LevelNotice
	case "warn":
		return hlog.LevelWarn
	case "error":
		return hlog.LevelError
	case "fatal":
		return hlog.LevelFatal
	default:
		return hlog.LevelInfo
	}
}
