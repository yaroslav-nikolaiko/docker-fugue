package yaml

import (
	"fmt"
	"github.com/kylelemons/go-gypsy/yaml"
	"os"
	"strings"
	"io/ioutil"
)

type YamlNode struct {
	root yaml.Node
}

func Create(filePath string) *YamlNode{
	file, err := yaml.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return &YamlNode{file.Root}
}

func (p YamlNode) Value() yaml.Node{
/*	n := nodeToMap(p.root)
	n["qwe"] = yaml.Scalar("asdasd")*/
	return p.root
}

func (p YamlNode) GetNode(path string) YamlNode{
	split := strings.Split(path, ".")
	return getNode(split, p.root)
}

func (p YamlNode) Write(path string){
	ioutil.WriteFile(path, []byte(yaml.Render(p.root)), 0644)
}

func (p YamlNode) AddScalar(element string) yaml.List{
	return append(nodeToList(p.root), yaml.Scalar(element))
}


func getNode(path []string, node yaml.Node) YamlNode{
	if len(path) > 1{
		return getNode(path[1:], nodeToMap(node)[path[0]])
	}
	return YamlNode{nodeToMap(node)[path[0]]}
}

func nodeToMap(node yaml.Node) (yaml.Map) {
	m, ok := node.(yaml.Map)
	if !ok {
		panic(fmt.Sprintf("%v is not of type map", node))
	}
	return m
}

func nodeToList(node yaml.Node) (yaml.List) {
	m, ok := node.(yaml.List)
	if !ok {
		panic(fmt.Sprintf("%v is not of type list", node))
	}
	return m
}


