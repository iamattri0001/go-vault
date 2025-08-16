package config

type Config struct {
	ServiceConfig ServiceConfig `json:"service"`
	MongoConfig   MongoConfig   `json:"mongo"`
}

type ServiceConfig struct {
	Name string `json:"name"`
	Port int    `json:"port"`
}

type MongoConfig struct {
	URI                    string `json:"uri"`
	Database               string `json:"database"`
	ConnectionTimeout      uint64 `json:"connection_timeout"`
	Timeout                uint64 `json:"timeout"`
	MaxPoolSize            uint64 `json:"max_pool_size"`
	MinPoolSize            uint64 `json:"min_pool_size"`
	IdleConnTimeout        int64  `json:"idle_conn_timeout"`
	ServerSelectionTimeout uint64 `json:"server_selection_timeout"`
	ConnectTimeout         uint64 `json:"connect_timeout"`
	QueryTimeout           uint64 `json:"query_timeout"`
	SocketTimeout          uint64 `json:"socket_timeout"`
	RetryWrites            bool   `json:"retry_writes"`
	RetryReads             bool   `json:"retry_reads"`
}
