package remotesteps

import "fmt"

type DockerConfigParameterIn struct {
	ServerAddr     string `json:"serverAddr" description:"the vps address" required:"true"`
	ServerPort     int    `json:"serverPort" description:"the vps port" required:"true"`
	ServerUser     string `json:"serverUser" description:"the vps user" required:"true"`
	ServerPassword string `json:"serverPassword" description:"the vps password" required:"true"`
}

func (d DockerConfigParameterIn) Validate() error {
	if d.ServerAddr == "" {
		return fmt.Errorf("serverAddr is required")
	}
	if d.ServerPort == 0 {
		return fmt.Errorf("serverPort is required")
	}
	if d.ServerUser == "" {
		return fmt.Errorf("serverUser is required")
	}
	if d.ServerPassword == "" {
		return fmt.Errorf("serverPassword is required")
	}
	return nil
}
