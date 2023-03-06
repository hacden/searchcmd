package pkg

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

type SearchDataStruct struct {
	Name string `yaml:"name"`
	Tags []string `yaml:"tags"`
	Data string `yaml:"data"`
}


func Search(file string) (string, string, string){
	Data := new(SearchDataStruct)
	yamlFile, err := os.ReadFile(file)
	CheckErrorOnExit(err)
	err = yaml.Unmarshal(yamlFile, Data)
	CheckErrorOnExit(err)
	return Data.Name,fmt.Sprintf("%s",Data.Tags),Data.Data
}

func ShowHighLightData(data string){
	lines := strings.Split(data,"\n")
	printColor := false
	for _,line := range lines{
		if strings.HasPrefix(line,"```"){
			if printColor {
				printColor = false
				continue
			}else{
				fmt.Println()
				printColor = true
				continue
			}
		}
		if printColor {
			GreenF.Println(line)
			continue
		}
		fmt.Println(line)
	}
}

func CheckErrorOnExit(err error) {
	if err != nil{
		panic(err)
	}
}