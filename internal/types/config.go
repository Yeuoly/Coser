package types

type BillboardsGlobalConfigurations struct {
	App struct {
		Port  int  `yaml:"port"`
		Debug bool `yaml:"debug"`
	} `yaml:"app"`
	DB struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
		User string `yaml:"user"`
		Pass string `yaml:"pass"`
	} `yaml:"db"`
	Redis struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
		Pass string `yaml:"pass"`
	} `yaml:"redis"`
	Minio struct {
		Host   string `yaml:"host"`
		Port   int    `yaml:"port"`
		AK     string `yaml:"ak"`
		SK     string `yaml:"sk"`
		Bucket string `yaml:"bucket"`
	} `yaml:"minio"`
	Email struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
		User string `yaml:"user"`
		Pass string `yaml:"pass"`
	}
	Cors struct {
		AllowOrigins []string `yaml:"allow_origins"`
		AllowMethods []string `yaml:"allow_methods"`
		AllowHeaders []string `yaml:"allow_headers"`
	} `yaml:"cors"`
	Github struct {
		ClientID     string `yaml:"client_id"`
		ClientSecret string `yaml:"client_secret"`
		RedirectURI  string `yaml:"redirect_uri"`
	} `yaml:"github"`
	Proxy struct {
		Address  string `yaml:"address"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"proxy"`
}
