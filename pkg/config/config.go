package config

type CommonConfig struct {
	Nacos NacosConf `yaml:"nacos" json:"nacos"`
}

func (c *CommonConfig) Default() {
	c.Nacos.Default()
}
