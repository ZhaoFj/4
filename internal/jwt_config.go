package internal

type JWTConfig struct {
	SingingKey string `mapstructure:"singing_key" json:"key"`
}
