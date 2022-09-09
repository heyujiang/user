package service

import (
	"fmt"
	consulApi "github.com/hashicorp/consul/api"
	"github.com/heyujiang/user/config"
)

func Register(conf config.RegisterCenter) {
	switch conf.Type {
	case "consul":
		registerConsul(conf.Consul)
	}
}

func registerConsul(conf config.Consul) {
	consulConf := &consulApi.Config{Address: conf.Address}
	client, err := consulApi.NewClient(consulConf)
	if err != nil {
		fmt.Println("")
	}

	registration := &consulApi.AgentServiceRegistration{
		ID:      "100",
		Name:    "UserServer100",
		Tags:    nil,
		Port:    81,
		Address: "127.0.0.1",
		//Check: &consulApi.AgentServiceCheck{
		//	Interval:                       "5s",
		//	Timeout:                        "5s",
		//	TLSSkipVerify:                  false,
		//	HTTP:                           "",
		//	DeregisterCriticalServiceAfter: "20s",
		//},
	}

	if err := client.Agent().ServiceRegister(registration); err != nil {
		fmt.Println("")
	}
}
