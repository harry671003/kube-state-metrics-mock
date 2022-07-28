package config

type Config struct {
	NumNamespaces   int `yaml:"num_namespaces"`
	NumNodes        int `yaml:"num_nodes"`
	NumPods         int `yaml:"num_pods"`
	NumDeployments  int `yaml:"num_deployments"`
	NumStatefulsets int `yaml:"num_statefulsets"`
}

var DefaultConfig = Config{
	NumNamespaces:   3,
	NumNodes:        3,
	NumPods:         3,
	NumDeployments:  3,
	NumStatefulsets: 3,
}
