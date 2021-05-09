package config

import (
	"io/ioutil"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/mlycore/log"
	"gopkg.in/yaml.v2"
)

// Config defines original config
type Config struct {
	rwLock      *sync.RWMutex `json:"-"`
	Loglevel    string        `json:"loglevel,omitempty"`
	DebugCharts bool          `yaml:"debugCharts"`

	// MySQL config
	MySQL *MySQLConfig `json:"mysql,omitempty"`
	//Redis config
	Redis *RedisConfig `json:"redis,omitempty"`
}

// MySQLConfig contains database and host configuration of MySQL
type MySQLConfig struct {
	Database string `json:"database,omitempty"`
	Proxy    string `json:"proxy,omitempty"`
	User     string `json:"user,omitempty"`
	Pass     string `json:"pass,omitempty"`
	Port     string `json:"port"`
}

// RedisConfig contains redis configuration of Redis
type RedisConfig struct {
	Proxy       string `json:"proxy,omitempty"`
	MaxIdle     int    `json:"maxIdle,omitempty"`
	MaxActive   int    `json:"maxActive,omitempty"`
	IdleTimeout int    `json:"idleTimeout,omitempty"`
	Port        string `json:"port"`
}

// DefaultConfigPath defines config path
const DefaultConfigPath = "configs/default.yml"

var (
	once   sync.Once
	config = &Config{
		rwLock: &sync.RWMutex{},
	}
)

// NewServerConfigOrDie defines a Config
func NewServerConfigOrDie() *Config {
	once.Do(func() {
		cfg, err := retrieveConfigFromFile(DefaultConfigPath)
		if err != nil {
			log.Fatalf("read config error: %s", err)
		}
		config = cfg
		log.Infof("config loaded")

		watcher, err := fsnotify.NewWatcher()
		if err != nil {
			log.Fatalf("get file watcher error: %s", err)
		}
		// defer watcher.Close()

		go func() {
			for {
				select {
				case event, ok := <-watcher.Events:
					if !ok {
						return
					}
					if event.Op&fsnotify.Write == fsnotify.Write {
						cfg, err := retrieveConfigFromFile(DefaultConfigPath)
						if err != nil {
							log.Errorf("log reloaded error: %s", err)
						} else {
							config = cfg
						}
					}
				case err, ok := <-watcher.Errors:
					if !ok {
						return
					}
					log.Errorf("file watcher error: %s", err)
				}
			}
		}()

		err = watcher.Add(DefaultConfigPath)
		if err != nil {
			log.Fatalf("add file watcher error: %s", err)
		}
		log.Infof("config file has been watched")
	})

	return config
}

func retrieveConfigFromFile(file string) (*Config, error) {
	cfg := &Config{}
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return cfg, err
	}

	err = yaml.Unmarshal(data, cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}

// GlobalServerConfig return a global config instance
func GlobalServerConfig() Config {
	return *config
}
