package repository

import (
	"github.com/spf13/viper"
)

var SERVICE_NAME []string
var SERVICE_DATA = make(map[string]string)
var SERVICE_TAGS = make(map[string]string)

func SetupServiceData() {
	ServiceList := viper.Get("service").([]interface{})
	//log.Println(ServiceList)
	for x := range ServiceList {
		SERVICE_NAME = append(SERVICE_NAME, ServiceList[x].(map[string]interface{})["name"].(string))
		SERVICE_TAGS[ServiceList[x].(map[string]interface{})["name"].(string)] = ServiceList[x].(map[string]interface{})["service-tags"].(string)
		SERVICE_DATA[ServiceList[x].(map[string]interface{})["name"].(string)] = ServiceList[x].(map[string]interface{})["service-id"].(string)
	}
	//log.Println(reflect.TypeOf(SERVICE_NAME), SERVICE_DATA)
}
