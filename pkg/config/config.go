package config

import (
	"io/ioutil"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

type Config struct {
	ClusterConfig ClusterConfig `yaml:"cluster_config"`
}

type ClusterConfig struct {
	NumNamespaces   int `yaml:"num_namespaces"`
	NumNodes        int `yaml:"num_nodes"`
	NumPods         int `yaml:"num_pods"`
	NumDeployments  int `yaml:"num_deployments"`
	NumStatefulsets int `yaml:"num_statefulsets"`
}

var DefaultConfig = Config{
	ClusterConfig: defaultClusterConfig,
}

var defaultClusterConfig = ClusterConfig{
	NumNamespaces:   3,
	NumNodes:        3,
	NumPods:         3,
	NumDeployments:  3,
	NumStatefulsets: 3,
}

func NewConfig(filename string) (*Config, error) {
	cfg := &DefaultConfig
	if filename != "" {
		if err := LoadConfigFromFile(filename, cfg); err != nil {
			return nil, err
		}
	}
	return cfg, nil
}

func LoadConfigFromFile(filename string, cfg *Config) error {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return errors.Wrap(err, "Error reading config file")
	}

	err = yaml.Unmarshal(buf, cfg)
	if err != nil {
		return errors.Wrap(err, "Error parsing config file")
	}

	return nil
}
