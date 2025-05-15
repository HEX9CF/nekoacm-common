package nacos

import (
	"errors"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

func (nc *NacosClient) Register() error {
	conf := nc.nacosConfig.Register

	success, err := nc.namingClient.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          conf.Ip,
		Port:        conf.Port,
		ServiceName: conf.ServiceName,
		Weight:      conf.Weight,
		Enable:      conf.Enable,
		Healthy:     conf.Healthy,
	})
	if err != nil {
		return err
	}
	if !success {
		return errors.New("Nacos 注册实例失败")
	}

	return nil
}
