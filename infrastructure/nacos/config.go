package nacos

import (
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"log"
	"nekoacm-common/pkg/utils"
)

func (nc *NacosClient) GetConfig(v interface{}, dataId string) error {
	conf := nc.nacosConfig.Config

	log.Println("获取配置:", conf.Group, dataId)
	content, err := nc.configClient.GetConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  conf.Group})
	if err != nil {
		return err
	}
	if content == "" {
		return nil
	}

	err = utils.UnmarshalYaml(v, content)
	if err != nil {
		return err
	}

	return nil
}
