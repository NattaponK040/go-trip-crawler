package config

import (
	"github.com/spf13/viper"
	"log"
	"reflect"
	"strings"
)

type ServiceConfig struct {
	Server    Server  `mapstructure:"server"`
	Mongoinfo Mongodb `mapstructure:"mongodb"`
}

type Server struct {
	MailSender string `mapstructure:"mailSender"`
	PassSender string `mapstructure:"passSender"`
}

type Mongodb struct {
	Url                string `mapstrcuture:"uri"`
	DBname             string `mapstrcuture:"dbName"`
	WonnaiTravelReview string `mapstrcuture:"wongnai-travel-review"`
	WonnaiTravelData   string `mapstrcuture:"wongnai-travel-data"`
	WonnaiHotelReview  string `mapstrcuture:"wongnai-hotel-review"`
	WonnaiHotelData    string `mapstrcuture:"wongnai-hotel-data"`
	WonnaiShopReview   string `mapstrcuture:"wongnai-shop-review"`
	WonnaiShopData     string `mapstrcuture:"wongnai-shop-data"`

	TripadvisorTravelReview string `mapstrcuture:"tripadvior-travel-review"`
	TripadvisorTravelData   string `mapstrcuture:"tripadvior-travel-data"`
	TripadvisorHotelReview  string `mapstrcuture:"tripadvior-hotel-review"`
	TripadvisorHotelData    string `mapstrcuture:"tripadvior-hotel-data"`
	TripadvisorShopReview   string `mapstrcuture:"tripadvior-shop-review"`
	TripadvisorShopData     string `mapstrcuture:"tripadvior-shop-data"`

	MaxPool uint64 `mapstrcuture:"maxpool"`
	MinPool uint64 `mapstrcuture:"minpool"`
}

type Config struct {
	parenPath   string
	resource    string
	env         string
	application string
}

// Load function will read config from environment or config file.
func LoadConfig(parentPath string, container string, env string, fileNames ...string) ServiceConfig {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))
	viper.SetConfigType("yaml")

	for _, fileName := range fileNames {
		viper.SetConfigName(fileName)
	}

	viper.AddConfigPath("./" + container + "/")
	if len(parentPath) > 0 {
		viper.AddConfigPath("./" + parentPath + "/" + container + "/")
	}

	if err := viper.ReadInConfig(); err != nil {
		switch err.(type) {
		case viper.ConfigFileNotFoundError:
			log.Println("config file not found")
		default:
			panic(err)
		}
	}
	if len(env) > 0 {
		env2 := strings.ToLower(env)
		for _, fileName2 := range fileNames {
			name := fileName2 + "-" + env2
			viper.SetConfigName(name)
			viper.MergeInConfig()
		}
	}
	var c ServiceConfig
	bindEnvs(c)
	viper.Unmarshal(&c)
	return c
}

// bindEnvs function will bind ymal file to struc model
func bindEnvs(iface interface{}, parts ...string) {
	ifv := reflect.ValueOf(iface)
	ift := reflect.TypeOf(iface)
	for i := 0; i < ift.NumField(); i++ {
		v := ifv.Field(i)
		t := ift.Field(i)
		tv, ok := t.Tag.Lookup("mapstructure")
		if !ok {
			continue
		}
		switch v.Kind() {
		case reflect.Struct:
			bindEnvs(v.Interface(), append(parts, tv)...)
		default:
			viper.BindEnv(strings.Join(append(parts, tv), "."))
		}
	}
}
