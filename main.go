package main

import (
	"log"
	"os"
	"fmt"
	"github.com/yaroslav-nikolaiko/docker-fugue/parser"
)

func init() {
// Change the device for logging to stdout.
	log.SetOutput(os.Stdout)
}

func main(){
	root := yaml.Create("data/sample.yml")
	r := root.Value()
	fmt.Println(r)
	node := root.GetNode("services.dictionary-service.external_links")
	node.AddScalar("sdfsdf")
	fmt.Println(node)
	//root.Write("output.yaml")

}