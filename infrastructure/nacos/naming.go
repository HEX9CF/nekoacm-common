package nacos

import (
	"errors"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"log"
	"strconv"
)

func (nc *NacosClient) Register() error {
	conf := nc.nacosConfig.Register

	log.Println("注册 " + conf.ServiceName + " 到 Nacos：" + conf.Ip + ":" + strconv.FormatUint(conf.Port, 10))
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
