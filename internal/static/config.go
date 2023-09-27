package static

import (
	"os"

	"github.com/Yeuoly/billboards/internal/types"
	"gopkg.in/yaml.v3"
)

var billboardsGlobalConfigurations types.BillboardsGlobalConfigurations

func InitConfig(path string) error {
	billboardsGlobalConfigurations = types.BillboardsGlobalConfigurations{}

	// read config file
	configFile, err := os.Open(path)
	if err != nil {
		return err
	}

	defer configFile.Close()

	// parse config file
	decoder := yaml.NewDecoder(configFile)
	err = decoder.Decode(&billboardsGlobalConfigurations)
	if err != nil {
		return err
	}

	return nil
}

// avoid global modification, use value copy instead
func GetBillboardsGlobalConfigurations() types.BillboardsGlobalConfigurations {
	return billboardsGlobalConfigurations
}
