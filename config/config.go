package config

import (
	"errors"
	"fmt"
	"github.com/VSRestia/VSRestia-Client/utils"
	"io/ioutil"
	"os"
	"syscall"
)

var errNotExist = errors.New("config not exist")

var config map[string]interface{}

func CheckConfig() (map[string]interface{}, error) {
	var err error
	if !ConfigIsExist() {
		err = errNotExist
		return nil, err
	} else {
		configFile, err := os.OpenFile("config.json", syscall.O_RDONLY, 0666)
		if err != nil {
			return nil, err
		}
		ctx, _ := ioutil.ReadAll(configFile)
		config = utils.JsonLoads(string(ctx))
		defer func() {
			_ = configFile.Close()
		}()
	}
	return config, nil
}

//Checking config file if exist.
//If not, create config file.
func ConfigIsExist() bool {
	if utils.IsExist("config.json") {
		fmt.Println("Config file detected")
		return true
	} else {
		fmt.Println("Cannot find config file")
		configFile, err1 := os.OpenFile("config.json", syscall.O_CREAT|syscall.O_RDWR, 0666)
		if err1 != nil {
			fmt.Println(err1)
		}
		_, _ = configFile.WriteString("{\n  \"LocalProxyPort\": 25565,\n  \"ServerPort\": 25565,\n  \"ServerIP\": \"xxx.xxx.xxx.xxx\",\n  \"SecretKey\": \"xxxxxxxxxxxxxxxx\",\n  \"EncryptType\": \"AES\",\n  \"EncryptMode\": \"CBC\"\n}")
		defer func() {
			err2 := configFile.Close()
			if err2 != nil {
				fmt.Println(err2)
			}
		}()
		fmt.Println("Config file created.")
		return false
	}
}
