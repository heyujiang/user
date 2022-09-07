package config

import (
	"sync"
	"sync/atomic"

	"github.com/fsnotify/fsnotify"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

var (
	confValue       *atomic.Value
	watchChangeOnce sync.Once
)

type Config struct {
	DataBase DataBase `toml:"DataBase"`
	Redis    Redis    `toml:"Redis"`
}

type DataBase struct {
	Debug bool   `toml:"debug"`
	Type  string `toml:"type"`
	MySql MySql  `toml:"MySql"`
}

type MySql struct {
	Version  string `toml:"version"`
	UserName string `toml:"user_name"`
	PassWord string `toml:"pass_word"`
	Address  string `toml:"address"`
	Port     int64  `toml:"port"`
	DataBase string `toml:"data_base"`
}

type Redis struct {
	Url      string `toml:"url"`
	PassWord string `toml:"pass_word"`
	Db       int    `toml:"db"`
}

func init() {
	confValue = &atomic.Value{}
}

func GetConfig() Config {
	return *confValue.Load().(*Config)
}

func InitConfig(cfgFile string) {
	viper.AddConfigPath("./")
	viper.SetConfigName(cfgFile)
	viper.SetConfigType("toml")

	loadConfig()
	watchChangeOnce.Do(viper.WatchConfig)
}

func loadConfig() {
	conf := &Config{}

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(conf, func(config *mapstructure.DecoderConfig) {
		config.TagName = "toml"
	}); err != nil {
		panic(err)
	}
	confValue.Store(conf)
}

func update(action func()) {
	viper.OnConfigChange(func(in fsnotify.Event) {
		loadConfig()
		action()
	})
}
