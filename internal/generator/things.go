package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	wordwrap "github.com/mitchellh/go-wordwrap"
)

// go run internal/generator/things.go ~/Downloads/schema.jsonld things

type schemaFull struct {
	Graph []schemaFullNode `json:"@graph"`
}

type schemaFullNode struct {
	ID             string                       `json:"@id"`
	Type           schemaFullNodeType           `json:"@type"`
	DomainIncludes schemaFullNodeDomainIncludes `json:"http://schema.org/domainIncludes"`
	Comment        string                       `json:"rdfs:comment"`
}

type schemaFullNodeDomainIncludes []schemaFullRef

func (s *schemaFullNodeDomainIncludes) UnmarshalJSON(bytes []byte) error {
	var data []schemaFullRef

	if bytes[0] == '{' {
		bytes = append([]byte("["), bytes...)
		bytes = append(bytes, ']')
	}

	err := json.Unmarshal(bytes, &data)
	if err != nil {
		return err
	}

	*s = schemaFullNodeDomainIncludes(data)

	return nil
}

func (s schemaFullNodeDomainIncludes) Includes(value string) bool {
	for _, t := range s {
		if t.ID == value {
			return true
		}
	}

	return false
}

type schemaFullNodeType []string

func (s *schemaFullNodeType) UnmarshalJSON(bytes []byte) error {
	var data []string

	if bytes[0] == '"' {
		bytes = append([]byte("["), bytes...)
		bytes = append(bytes, ']')
	}

	err := json.Unmarshal(bytes, &data)
	if err != nil {
		return err
	}

	*s = schemaFullNodeType(data)

	return nil
}

func (s schemaFullNodeType) Includes(value string) bool {
	for _, t := range s {
		if t == value {
			return true
		}
	}

	return false
}

type schemaFullRef struct {
	ID string `json:"@id"`
}

func main() {
	bytes, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	var data schemaFull

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		panic(err)
	}

	for _, node := range data.Graph {
		if node.Type.Includes("http://schema.org/DataType") {
			continue
		} else if node.ID == "http://schema.org/URL" {
			continue
		} else if node.ID == "http://schema.org/Float" {
			continue
		} else if node.ID == "http://schema.org/Integer" {
			continue
		} else if node.ID == "http://schema.org/False" {
			continue
		} else if node.ID == "http://schema.org/True" {
			continue
		} else if !node.Type.Includes("rdfs:Class") {
			continue
		}

		typ := path.Base(node.ID)
		ctx := strings.TrimSuffix(node.ID, fmt.Sprintf("/%s", typ))
		typLower := strings.ToLower(typ)

		err = os.MkdirAll(filepath.Join(os.Args[2], typLower), 0755)
		if err != nil {
			panic(err)
		}

		err = ioutil.WriteFile(
			filepath.Join(os.Args[2], typLower, "type.go"),
			[]byte(
				fmt.Sprintf(
					`package %s

import "github.com/dpb587/go-schemaorg"

// %s
var Type = schemaorg.NewDataType("%s", "%s")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
`,
					typLower,
					"// "+strings.Replace(wordwrap.WrapString(node.Comment, 77), "\n", "\n// ", -1),
					ctx,
					typ,
				),
			),
			0644)
		if err != nil {
			panic(err)
		}

		properties := []schemaFullNode{}

		for _, subnode := range data.Graph {
			if !subnode.DomainIncludes.Includes(node.ID) {
				continue
			}

			properties = append(properties, subnode)
		}

		if len(properties) > 0 {
			first := true
			str := fmt.Sprintf(
				`package %s

import "github.com/dpb587/go-schemaorg"

var (`,
				typLower,
			)

			for _, property := range properties {
				pro := path.Base(property.ID)

				if !first {
					str = str + "\n"
				}

				str = str + fmt.Sprintf(
					`
	%s
	%s = schemaorg.NewProperty("%s")`,
					"// "+strings.Replace(wordwrap.WrapString(property.Comment, 77), "\n", "\n	// ", -1),
					strings.ToUpper(string(pro[0]))+pro[1:],
					pro,
				)

				first = false
			}

			str = str + `
)
`

			err = ioutil.WriteFile(
				filepath.Join(os.Args[2], typLower, "properties.go"),
				[]byte(str),
				0644,
			)
			if err != nil {
				panic(err)
			}
		}
	}
}
