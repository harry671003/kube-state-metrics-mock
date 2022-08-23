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
	// Cluster level
	NumNamespaces int `yaml:"num_namespaces"`
	NumNodes      int `yaml:"num_nodes"`

	// Namespace level
	NumDeploymentsPerNamespace  int `yaml:"num_deployments_per_namepace"`
	NumStatefulsetsPerNamespace int `yaml:"num_statefulsets_per_namespace"`
	NumPVCsPerNamespace         int `yaml:"num_pvcs_per_namespace"`

	// Deployment level
	NumReplicasPerDeployment  int `yaml:"num_replicas_per_deployments"`
	NumReplicasPerStatefulset int `yaml:"num_replicas_per_statefulset"`

	// Pod level
	NumContainersPerPod int `yaml:"num_containers_per_pod"`
}

var DefaultConfig = Config{
	ClusterConfig: defaultClusterConfig,
}

var defaultClusterConfig = ClusterConfig{
	NumNamespaces: 3,
	NumNodes:      3,

	NumDeploymentsPerNamespace:  3,
	NumStatefulsetsPerNamespace: 3,
	NumPVCsPerNamespace:         3,

	NumReplicasPerDeployment:  3,
	NumReplicasPerStatefulset: 3,

	NumContainersPerPod: 3,
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
