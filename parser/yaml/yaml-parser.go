package yaml

import (
	"fmt"
	"github.com/kylelemons/go-gypsy/yaml"
	"os"
	"io/ioutil"
	"strings"
)

type YamlNode struct {
	root yaml.Node
	node yaml.Node
	path string
}

func Create(filePath string) YamlNode{
	file, err := yaml.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return YamlNode{file.Root,file.Root, "."}
}

func (p YamlNode) String() string{
	return yaml.Render(p.node)
}

func (p YamlNode) GetNode(path string) YamlNode{
	path = p.path + path
	node, _ := yaml.Child(p.node, path)
	return YamlNode{p.root, node, path}
}

func (p YamlNode) Write(path string){
	ioutil.WriteFile(path, []byte(yaml.Render(p.root)), 0644)
}

func (p YamlNode) AddScalar(element string) YamlNode{
	list, _ := p.node.(yaml.List)
	pathSlice := strings.Split(p.path, ".")
	currentKey := pathSlice[len(pathSlice) - 1 ]
	parent, _ := yaml.Child(p.root, strings.Join(pathSlice[:len(pathSlice)-1], "."))
	parentMap, _ := parent.(yaml.Map)

	list = append(list, yaml.Scalar(element))
	parentMap[currentKey] = list
	p.node = list

	return p
}



