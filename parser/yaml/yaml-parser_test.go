package yaml

import (
	"testing"
	"strings"
)

var filePath = "../../data/sample.yml"
var original = Create(filePath)

func TestGetNode(t *testing.T) {
	root := Create(filePath)

	result := root.GetNode("services.dictionary-service.external_links")
	stringValue := result.String()
	expected := "- mysql"
	if strings.TrimSpace(stringValue)!= expected{
		t.Errorf("GetNode: Expected mydql, got %s", stringValue)
	}
}

func TestGetNodeArrayElement(t *testing.T) {
	root := Create(filePath)

	result := root.GetNode("services.dictionary-service.external_links[0]")
	stringValue := result.String()
	expected := "mysql"
	if strings.TrimSpace(stringValue)!= expected{
		t.Errorf("Expected %s, got %s",expected,  stringValue)
	}
}

func TestAddScalarToList(t *testing.T) {
	root := Create(filePath)

	links := root.GetNode("services.dictionary-service.external_links")
	links.AddScalar("db2")

	mysql := root.GetNode("services.dictionary-service.external_links[0]").String()
	db2 := root.GetNode("services.dictionary-service.external_links[1]").String()

	if strings.TrimSpace(mysql)!= "mysql"{
		t.Errorf("Expected mysql, got %s", mysql)
	}
	if strings.TrimSpace(db2)!= "db2"{
		t.Errorf("Expected db2, got %s", db2)
	}
}
