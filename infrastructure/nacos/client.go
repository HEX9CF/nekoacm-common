package nacos

import (
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"nekoacm-common/pkg/config"
)

type NacosClient struct {
	nacosConfig   *config.NacosConf
	clientConfig  constant.ClientConfig
	serverConfigs []constant.ServerConfig
	namingClient  naming_client.INamingClient
	configClient  config_client.IConfigClient
}

func NewNacosClient(nacosConfig *config.NacosConf) *NacosClient {
	return &NacosClient{
		nacosConfig: nacosConfig,
	}
}

func (nc *NacosClient) Init() error {
	// 初始化 Nacos 配置
	nc.initNacosConfig()

	// 初始化客户端
	if err := nc.initClient(); err != nil {
		return err
	}

	return nil
}

func (nc *NacosClient) initNacosConfig() {
	cConf := nc.nacosConfig.Client
	sConfs := nc.nacosConfig.Server

	// 创建 clientConfig
	nc.clientConfig = *constant.NewClientConfig(
		constant.WithNamespaceId(cConf.NamespaceId),
		constant.WithTimeoutMs(cConf.TimeoutMs),
		constant.WithNotLoadCacheAtStart(cConf.NotLoadCacheAtStart),
		constant.WithLogDir(cConf.LogDir),
		constant.WithCacheDir(cConf.CacheDir),
		constant.WithLogLevel(cConf.LogLevel),
	)

	// 创建 serverConfig
	nc.serverConfigs = make([]constant.ServerConfig, 0, len(sConfs))
	for _, conf := range sConfs {
		nc.serverConfigs = append(nc.serverConfigs, *constant.NewServerConfig(
			conf.IpAddr,
			conf.Port,
			constant.WithScheme(conf.Scheme),
			constant.WithContextPath(conf.ContextPath),
		))
	}
}

func (nc *NacosClient) initClient() error {
	var err error

	// 创建服务发现客户端
	nc.namingClient, err = clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &nc.clientConfig,
			ServerConfigs: nc.serverConfigs,
		},
	)
	if err != nil {
		return err
	}

	// 创建动态配置客户端
	nc.configClient, err = clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &nc.clientConfig,
			ServerConfigs: nc.serverConfigs,
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func (nc *NacosClient) GetNamingClient() naming_client.INamingClient {
	return nc.namingClient
}

func (nc *NacosClient) GetConfigClient() config_client.IConfigClient {
	return nc.configClient
}
