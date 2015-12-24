package configs

import (
	"github.com/vivowares/octopus/Godeps/_workspace/src/github.com/spf13/viper"
	"path"
	"time"
)

var Config *Conf

func InitializeConfig(filename string) error {
	ext := path.Ext(filename)
	filePath := path.Dir(filename)
	filename = path.Base(filename[0 : len(filename)-len(ext)])

	viper.SetConfigName(filename)
	viper.AddConfigPath(filePath)
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	serviceConfig := &ServiceConf{
		Host:     viper.GetString("service.host"),
		HttpPort: viper.GetInt("service.http_port"),
		WsPort:   viper.GetInt("service.ws_port"),
		PidFile:  viper.GetString("service.pid_file"),
	}

	dbConfig := &DbConf{
		DbType: viper.GetString("database.db_type"),
		DbFile: viper.GetString("database.db_file"),
	}

	indexConfig := &IndexConf{
		Host:             viper.GetString("indices.host"),
		Port:             viper.GetInt("indices.port"),
		NumberOfShards:   viper.GetInt("indices.number_of_shards"),
		NumberOfReplicas: viper.GetInt("indices.number_of_replicas"),
		TTLEnabled:       viper.GetBool("indices.ttl_enabled"),
		TTL:              viper.GetDuration("indices.ttl"),
	}

	connConfig := &ConnectionConf{
		Registry:         viper.GetString("connections.registry"),
		NShards:          viper.GetInt("connections.nshards"),
		InitShardSize:    viper.GetInt("connections.init_shard_size"),
		RequestQueueSize: viper.GetInt("connections.request_queue_size"),
		Expiry:           viper.GetDuration("connections.expiry"),
		Timeouts: &ConnectionTimeoutConf{
			Write:    viper.GetDuration("connections.timeouts.write"),
			Read:     viper.GetDuration("connections.timeouts.read"),
			Request:  viper.GetDuration("connections.timeouts.request"),
			Response: viper.GetDuration("connections.timeouts.response"),
		},
		BufferSizes: &ConnectionBufferSizeConf{
			Write: viper.GetInt("connections.buffer_sizes.write"),
			Read:  viper.GetInt("connections.buffer_sizes.read"),
		},
	}

	Config = &Conf{
		Service:     serviceConfig,
		Connections: connConfig,
		Indices:     indexConfig,
		Database:    dbConfig,
	}

	return nil
}

type Conf struct {
	Service     *ServiceConf
	Connections *ConnectionConf
	Indices     *IndexConf
	Database    *DbConf
}

type DbConf struct {
	DbType string
	DbFile string
}

type IndexConf struct {
	Host             string
	Port             int
	NumberOfShards   int
	NumberOfReplicas int
	TTLEnabled       bool
	TTL              time.Duration
}

type ServiceConf struct {
	Host     string
	HttpPort int
	WsPort   int
	PidFile  string
}

type ConnectionConf struct {
	Registry         string
	NShards          int
	InitShardSize    int
	RequestQueueSize int
	Expiry           time.Duration
	Timeouts         *ConnectionTimeoutConf
	BufferSizes      *ConnectionBufferSizeConf
}

type ConnectionTimeoutConf struct {
	Write    time.Duration
	Read     time.Duration
	Request  time.Duration
	Response time.Duration
}

type ConnectionBufferSizeConf struct {
	Write int
	Read  int
}
