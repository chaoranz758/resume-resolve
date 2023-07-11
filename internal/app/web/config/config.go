package config

import (
	"encoding/json"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/viper"
	"os"
)

const (
	filePath = "../../configs/web.json"
)

type Config struct {
	ConfigInNacos configInNacos
	Nacos         ConfigNacos
}

type configInNacos struct {
	ResumeResolvingService `json:"resumeResolvingService"`
	Log                    `json:"log"`
	EtcdServer             `json:"etcdServer"`
	Oss                    `json:"oss"`
	RequestGRPCServer      `json:"requestRpcServer"`
	Jwt                    `json:"jwt"`
}

type ResumeResolvingService struct {
	Name string `mapstructure:"name" json:"name"`
	Mode string `mapstructure:"mode" json:"mode"`
}

type Log struct {
	Level      string `mapstructure:"level" json:"level"`
	FilePath   string `mapstructure:"file_path" json:"file_path"`
	MaxSize    int    `mapstructure:"max_size" json:"max_size"`
	MaxAge     int    `mapstructure:"max_age" json:"max_age"`
	MaxBackups int    `mapstructure:"max_backups" json:"max_backups"`
}

type Jwt struct {
	Key        string `json:"key"`
	TimeOut    int    `json:"time_out"`
	MaxRefresh int    `json:"max_refresh"`
}

type EtcdServer struct {
	Ip   string `json:"ip"`
	Port int    `json:"port"`
}

type Oss struct {
	Endpoint        string `json:"endpoint"`
	AccessKeyId     string `json:"accessKeyId"`
	AccessKeySecret string `json:"accessKeySecret"`
	BucketName      string `json:"bucketName"`
}

type RequestGRPCServer struct {
	UserService struct {
		Name string `json:"name"`
	} `json:"user_service"`
	ResumeService struct {
		Name string `json:"name"`
	} `json:"resume_service"`
	PostService struct {
		Name string `json:"name"`
	} `json:"post_service"`
	PostCategoryService struct {
		Name string `json:"name"`
	} `json:"post_category_service"`
	DepartmentService struct {
		Name string `json:"name"`
	} `json:"department_service"`
}

type ConfigNacos struct {
	Nacos
}

type Nacos struct {
	IP                  string `mapstructure:"ip" json:"ip"`
	Port                int    `mapstructure:"port" json:"port"`
	Namespace           string `mapstructure:"namespace" json:"namespace"`
	TimeoutMs           int    `mapstructure:"timeoutMs" json:"timeoutMs"`
	NotLoadCacheAtStart bool   `mapstructure:"notLoadCacheAtStart" json:"notLoadCacheAtStart"`
	LogDir              string `mapstructure:"logDir" json:"logDir"`
	CacheDir            string `mapstructure:"cacheDir" json:"cacheDir"`
	LogLevel            string `mapstructure:"logLevel" json:"logLevel"`
	DataID              string `mapstructure:"dataID" json:"dataID"`
	Group               string `mapstructure:"group" json:"group"`
}

func NewConfig() *Config {
	return new(Config)
}

func (config *Config) Init() (err error) {
	//指定配置文件路径
	viper.SetConfigFile(filePath)
	//打开配置文件
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	//读取配置文件
	if err = viper.ReadConfig(file); err != nil {
		return
	}
	if err = viper.Unmarshal(&config.Nacos); err != nil {
		return
	}
	// 至少一个ServerConfig
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: config.Nacos.IP,
			Port:   uint64(config.Nacos.Port),
		},
	}
	// 创建clientConfig
	clientConfig := constant.ClientConfig{
		NamespaceId:         config.Nacos.Namespace, // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId。当namespace是public时，此处填空字符串。
		TimeoutMs:           uint64(config.Nacos.TimeoutMs),
		NotLoadCacheAtStart: config.Nacos.NotLoadCacheAtStart,
		LogDir:              config.Nacos.LogDir,
		CacheDir:            config.Nacos.CacheDir,
		LogLevel:            config.Nacos.LogLevel,
	}
	// 创建动态配置客户端的另一种方式 (推荐)
	configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		return
	}
	//获取配置信息
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: config.Nacos.DataID,
		Group:  config.Nacos.Group})
	if err != nil {
		return
	}
	if err = json.Unmarshal([]byte(content), &config.ConfigInNacos); err != nil {
		return
	}
	//监听配置
	err = configClient.ListenConfig(vo.ConfigParam{
		DataId: config.Nacos.DataID,
		Group:  config.Nacos.Group,
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("配置中心的配置文件修改了")
			fmt.Println("group:" + group + ", dataId:" + dataId)
		},
	})
	if err != nil {
		return
	}
	//监控配置文件变化
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		// 配置文件发生变更之后会调用的回调函数
		fmt.Println("本地配置文件修改了")
		if err = viper.Unmarshal(&config.ConfigInNacos); err != nil {
			return
		}
	})
	return
}
