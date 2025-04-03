package config

import (
    "sync"

    "github.com/MafiaLogiki/common/logger"
    "github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
    IsDebug *bool `yaml:"is_debug"`
    Listen struct {
        BindIp string `yaml:"bind_ip"`
        Port string `yaml:"port"`
    } `yaml:"listen"`

    Postgres struct {
        Host string `yaml:"host"` 
        Port string `yaml:"port"`
        HostName string `yaml:"host_name"`
        Password string `yaml:"password"`
        DBName string `yaml:"db_name"`
    } `yaml:"postgres"`
}

var instance *Config
var once sync.Once

func GetConfig(l logger.Logger) *Config {
    once.Do(func() {
        l.Info("Read application configuration")
        instance = &Config{}
        
        if err := cleanenv.ReadConfig("config.yml", instance); err != nil {
            help, _ := cleanenv.GetDescription(instance, nil)
            l.Info(help)
            l.Fatal(err)
        }
    })

    return instance
}
