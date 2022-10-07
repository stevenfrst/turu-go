package gitlab

import (
	"encoding/json"
	"fmt"
	"gitag-request/config"
	"gitag-request/repository"
	"github.com/manifoldco/promptui"
	"io/ioutil"
	"net/http"
)

type Tags struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

var TagsName []string

func GetAPI(projectID string, tagsName string) {
	client := &http.Client{}
	req, err := http.NewRequest(config.GITLAB_METHOD, config.GITLAB_URL+projectID+"/repository/tags", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", config.API_KEY)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(string(body))
	tags := []Tags{}
	json.Unmarshal([]byte(body), &tags)
	fmt.Println("Available Tags :")
	for _, tag := range tags {
		//fmt.Println(tag.Name)
		TagsName = append(TagsName, tag.Name)
	}

	fmt.Println("Select Tags : ")

	prompt := promptui.Select{
		Label: "Tags List",
		Items: TagsName,
	}
	_, result, err := prompt.Run()
	repository.ReadEnv(result, tagsName)
	//log.Println("Selected Tags : ", result)
}
