package main

import (
	"log"
	"os"
	"fmt"
	"github.com/yaroslav-nikolaiko/docker-fugue/parser/yaml"
)

func init() {
// Change the device for logging to stdout.
	log.SetOutput(os.Stdout)
}

func main(){
	root := yaml.Create("data/sample.yml")
	node := root.GetNode("services.dictionary-service.external_links")
	node = node.AddScalar("sdfsdfnb")
	fmt.Println(node)
	fmt.Println(root)
	//root.Write("output.yaml")

}