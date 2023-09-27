package static

import (
	"os"

	"github.com/Yeuoly/billboards/internal/types"
	"gopkg.in/yaml.v3"
)

var bocchiGlobalConfigurations types.BillboardsGlobalConfigurations

func InitConfig(path string) error {
	bocchiGlobalConfigurations = types.BillboardsGlobalConfigurations{}

	// read config file
	configFile, err := os.Open(path)
	if err != nil {
		return err
	}

	defer configFile.Close()

	// parse config file
	decoder := yaml.NewDecoder(configFile)
	err = decoder.Decode(&bocchiGlobalConfigurations)
	if err != nil {
		return err
	}

	return nil
}

// avoid global modification, use value copy instead
func GetBocchiGlobalConfigurations() types.BillboardsGlobalConfigurations {
	return bocchiGlobalConfigurations
}
