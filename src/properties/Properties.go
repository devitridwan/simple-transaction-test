package properties

import (
	"fmt"
)

type DatabaseProperties struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	SllMode  string `json:"sslmode"`
}

type RouterEndPointProperties struct {
	RestBaseUrl        string                                          `json:"rest_base_url" mapstructure:"rest_base_url"`
	EnvirontmentStatus string                                          `json:"environtment_status" mapstructure:"environtment_status"`
	Channel            map[string]map[string]ChannelEndpointProperties `json:"channel" `
}

func (this *RouterEndPointProperties) GetChannelRest(name string, types string) string {
	return fmt.Sprintf("%s/%s/%s", this.RestBaseUrl, this.EnvirontmentStatus, this.Channel[types][name].RestPath)
}

type ChannelProperties struct {
	Path        string `json:"path" mapstructure:"path"`
	QueueName   string `json:"queue_name" mapstructure:"queue_name"`
	DurableName string `json:"durable_name" mapstructure:"durable_name"`
}

type ChannelEndpointProperties struct {
	RestPath string `json:"rest_path" mapstructure:"rest_path"`
}

type EndpointProperties struct {
	Topic      RouterEndPointProperties `json:"topic" mapstructure:"topic"`
	TargetPath string                   `json:"target_path" mapstructure:"target_path"`
	Port       int                      `json:"port" mapstructure:"port"`
	Database   DatabaseProperties       `json:"database"  mapstructure:"database"`
}
