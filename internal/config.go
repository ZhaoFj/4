package internal

type ShopCartSrvConfig struct {
	SrvName string   `mapstructure:"srvName" json:"srvName"`
	Host    string   `mapstructure:"host" json:"host"`
	Port    int      `mapstructure:"port" json:"port"`
	Tags    []string `mapstructure:"tags" json:"tags"`
}

type ShopCartWebConfig struct {
	SrvName string   `mapstructure:"webName" json:"webName"`
	Host    string   `mapstructure:"host" json:"host"`
	Port    int      `mapstructure:"port" json:"port"`
	Tags    []string `mapstructure:"tags" json:"tags"`
}

type DBConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	DB       string `mapstructure:"db"`
	UserName string `mapstructure:"userName"`
	Password string `mapstructure:"password"`
}

type AppConfig struct {
	ConsulConfig      ConsulConfig      `mapstructure:"consul" json:"consul"`
	ShopCartSrvConfig ShopCartSrvConfig `mapstructure:"shopcart_srv" json:"shopcart_srv"`
	ShopCartWebConfig ShopCartWebConfig `mapstructure:"shopcart_web" json:"shopcart_web"`
	DBConfig          DBConfig          `mapstructure:"db" json:"db"`
	JWTConf           JWTConfig         `mapstructure:"jwt" json:"jwt"`
	Debug             bool              `mapstructure:"debug" json:"debug"`
}
