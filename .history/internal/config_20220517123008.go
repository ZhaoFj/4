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

type ProductSrvConfig struct {
	SrvName string   `mapstructure:"srvName" json:"srvName"`
	Host    string   `mapstructure:"host" json:"host"`
	Port    int      `mapstructure:"port" json:"port"`
	Tags    []string `mapstructure:"tags" json:"tags"`
}

type ProductWebConfig struct {
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
	DBConfig          DBConfig          `mapstructure:"db" json:"db"`
	JaegerConfig      JaegerConfig      `mapstructure:"jaeger" json:"jaeger"`
	ShopCartSrvConfig ShopCartSrvConfig `mapstructure:"shopcart_srv" json:"shopcart_srv"`
	ShopCartWebConfig ShopCartWebConfig `mapstructure:"shopcart_web" json:"shopcart_web"`
	ProductSrvConfig  ProductSrvConfig  `mapstructure:"product_srv" json:"product_srv"`
	ProductWebConfig  ProductWebConfig  `mapstructure:"product_web" json:"product_web"`
	JWTConf           JWTConfig         `mapstructure:"jwt" json:"jwt"`
	Debug             bool              `mapstructure:"debug" json:"debug"`
}

type JaegerConfig struct {
	AgentHost string `mapstructure:"agentHost" json:"agentHost"`
	AgentProt int    `mapstructure:"agentPort" json:"agentPort"`
}
