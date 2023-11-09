package static

import (
	"os"

	"github.com/Yeuoly/coshub/internal/types"
	"gopkg.in/yaml.v3"
)

var coshubGlobalConfigurations types.CoshubGlobalConfigurations

func InitConfig(path string) error {
	coshubGlobalConfigurations = types.CoshubGlobalConfigurations{}

	// read config file
	configFile, err := os.Open(path)
	if err != nil {
		return err
	}

	defer configFile.Close()

	// parse config file
	decoder := yaml.NewDecoder(configFile)
	err = decoder.Decode(&coshubGlobalConfigurations)
	if err != nil {
		return err
	}

	return nil
}

// avoid global modification, use value copy instead
func GetCoshubGlobalConfigurations() types.CoshubGlobalConfigurations {
	return coshubGlobalConfigurations
}
