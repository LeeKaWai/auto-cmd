package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"

	"autocmd/messages"
	"gopkg.in/yaml.v3"
)

const configFileName = ".autocmd.yaml"

// Command 结构体存储单个命令的信息
type Command struct {
	Alias       string `yaml:"alias"`
	Command     string `yaml:"command"`
	Description string `yaml:"description,omitempty"` // omitempty 表示如果为空则不写入 YAML
}

// Config 结构体存储所有命令
type Config struct {
	Commands map[string]Command `yaml:"commands"`
}

// GetConfigPath 获取配置文件的完整路径 (例如 ~/.autocmd.yaml)
func GetConfigPath() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", fmt.Errorf(messages.ConfigUserHomeError, err)
	}
	return filepath.Join(usr.HomeDir, configFileName), nil
}

// LoadConfig 从文件加载配置
func LoadConfig() (*Config, error) {
	configPath, err := GetConfigPath()
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		Commands: make(map[string]Command),
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		// 文件不存在，返回一个空的配置，但不是错误
		return cfg, nil
	} else if err != nil {
		return nil, fmt.Errorf(messages.ConfigAccessError, configPath, err)
	}

	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf(messages.ConfigReadFileError, configPath, err)
	}

	err = yaml.Unmarshal(data, cfg)
	if err != nil {
		return nil, fmt.Errorf(messages.ConfigParseError, configPath, err)
	}

	return cfg, nil
}

// SaveConfig 将配置保存到文件
func SaveConfig(cfg *Config) error {
	configPath, err := GetConfigPath()
	if err != nil {
		return err
	}

	data, err := yaml.Marshal(cfg)
	if err != nil {
		return fmt.Errorf(messages.ConfigMarshalError, err)
	}

	// 写入文件，权限 0600 (只读写给文件所有者)
	err = ioutil.WriteFile(configPath, data, 0600)
	if err != nil {
		return fmt.Errorf(messages.ConfigWriteFileError, configPath, err)
	}

	return nil
}

