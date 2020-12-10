package config

type Specification struct {
	ListenAddr        string `envconfig:"GEOCODER_LISTEN_ADDR" required:"false" default:"0.0.0.0:8080"`
	MetricsListenAddr string `envconfig:"GEOCODER_METRICS_LISTEN_ADDR" required:"false" default:"0.0.0.0:9092"`
	RedisHost         string `envconfig:"REDIS_HOST" required:"false" default:"127.0.0.1"`
	RedisDatabase     int    `envconfig:"REDIS_DATABASE" required:"false" default:"15"`
	RedisPort         string `envconfig:"REDIS_PORT" required:"false" default:"6379"`
	GoogleApiKey      string `envconfig:"GOOGLE_API_KEY_SERVER" required:"true"`
	OpencageApiKey    string `envconfig:"OPENCAGE_API_KEY" required:"true"`
}
