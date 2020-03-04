package cctest

import (
	"fmt"
	"github.com/micro/go-micro/v2/config"
	"os"
	"testing"
)

type Host struct {
	Address string `json:"address"`
	Port int `json:"port"`
}

var host Host




func TestConfig(t *testing.T) {

	pwd, _ := os.Getwd()
	err := config.LoadFile(pwd + "/tmp/config.json")
	if err != nil {
		fmt.Println(err)
	}

	mp := config.Map()
	fmt.Println(mp["hosts"])

	config.Get("hosts", "database").Scan(&host)

	// 10.0.0.1 3306
	fmt.Println(host.Address, host.Port)

}
