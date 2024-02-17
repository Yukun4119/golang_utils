package util

type Config struct {
	TestConfig TestConfig `yaml:"testConfig"`
}

type TestConfig struct {
	Key   string `yaml:"key"`
	Value string `yaml:"value"`
}
