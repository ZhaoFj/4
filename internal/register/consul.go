package register

import (
	"fmt"
	"micro-trainning-part4/internal"

	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
)

type IRegister interface {
	Register(name, id string, port int, tags []string) error
	DeRegister(serviceId string) error
}

type ConsulRegistry struct {
	Host string
	Port int
}

func NewConsulRegistry(host string, port int) ConsulRegistry {
	return ConsulRegistry{
		Host: host,
		Port: port,
	}
}

func (cr ConsulRegistry) Register(name, id, serviceIp string, port int, tags []string) error {
	defaultConfig := api.DefaultConfig()
	h := internal.AppConf.ConsulConfig.Host
	p := internal.AppConf.ConsulConfig.Port
	defaultConfig.Address = fmt.Sprintf("%s:%d", h, p)
	client, err := api.NewClient(defaultConfig)
	if err != nil {
		zap.S().Error(err)
		return err
	}
	agentServiceReg := new(api.AgentServiceRegistration)
	agentServiceReg.Address = defaultConfig.Address
	agentServiceReg.Port = port
	agentServiceReg.ID = id
	agentServiceReg.Name = name
	agentServiceReg.Tags = tags
	serverAddr := fmt.Sprintf("http://%s:%d/health", serviceIp, port)
	check := api.AgentServiceCheck{
		HTTP:                           serverAddr,
		Timeout:                        "3s",
		Interval:                       "1s",
		DeregisterCriticalServiceAfter: "5s",
	}
	agentServiceReg.Check = &check
	err = client.Agent().ServiceRegister(agentServiceReg)
	if err != nil {
		zap.S().Error(err)
		return err
	}
	return nil

}

func (cr ConsulRegistry) DeRegister(serviceId string) error {
	defaultConfig := api.DefaultConfig()
	h := internal.AppConf.ConsulConfig.Host
	p := internal.AppConf.ConsulConfig.Port
	defaultConfig.Address = fmt.Sprintf("%s:%d", h, p)
	client, err := api.NewClient(defaultConfig)
	if err != nil {
		zap.S().Error(err)
		return err
	}
	err = client.Agent().ServiceDeregister(serviceId)
	if err != nil {
		zap.S().Error(err)
		return err
	}
	return nil
}
