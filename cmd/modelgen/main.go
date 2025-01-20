package main

// This tool generates the code needed for the request object.
//
// To call it, just call "go generate ./..." in the root folder of the repository

import (
	"bytes"
	_ "embed"
	"fmt"
	"go/format"
	"io"
	"log"
	"os"
	"sort"
	"strings"
	"text/template"
	"unicode"

	"github.com/OpenSlides/openslides-projector-service/pkg/modelsymlparse"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

//go:generate sh -c "go run main.go"

//go:embed model_struct.go.tmpl
var modelStructTmpl string

func main() {
	if err := run(os.Stdout); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func run(w io.Writer) error {
	if err := genStrutcts(); err != nil {
		return fmt.Errorf("generate field methods: %w", err)
	}

	return nil
}

func genStrutcts() error {
	r, err := openModelYML()
	if err != nil {
		log.Fatalf("Can not load models defition: %v", err)
	}
	defer r.Close()

	collections, err := parse(r)
	if err != nil {
		log.Fatalf("Can not parse model definition: %v", err)
	}

	tmpl, err := template.New("model.go").Parse(modelStructTmpl)
	if err != nil {
		return fmt.Errorf("parsing template: %w", err)
	}

	for _, coll := range collections {
		buf := new(bytes.Buffer)

		if err := tmpl.Execute(buf, coll); err != nil {
			return fmt.Errorf("executing template: %w", err)
		}

		formated, err := format.Source(buf.Bytes())
		if err != nil {
			return fmt.Errorf("formating code: %w", err)
		}

		if err := os.WriteFile("../../pkg/models/"+coll.CollectionName+".go", formated, 0o644); err != nil {
			return fmt.Errorf("writing output: %w", err)
		}
	}

	return nil
}

func openModelYML() (io.ReadCloser, error) {
	return os.Open("../../meta/models.yml")
}

// parse returns all fields from the models.yml with their go-type as string.
func parse(r io.Reader) ([]collection, error) {
	inData, err := modelsymlparse.Unmarshal(r)
	if err != nil {
		return nil, fmt.Errorf("unmarshalling models.yml: %w", err)
	}

	var collections []collection
	for collectionName, modelCollection := range inData {
		var fields []field
		var relations []relation
		for fieldName, modelField := range modelCollection.Fields {
			f := field{}
			f.GoName = goName(fieldName)
			f.ValueType = valueType(modelField.Type, modelField.Required)
			f.FieldName = fieldName
			f.Required = modelField.Required

			if modelField.Type == "relation" || modelField.Type == "generic-relation" {
				f.SingleRelation = true
			}

			if modelField.Type == "relation" || modelField.Type == "relation-list" {
				var propName string
				if modelField.Type == "relation" {
					propName = strings.Replace(fieldName+"$", "_id$", "", 1)
				} else {
					propName = strings.Replace(fieldName+"$", "_ids$", "s", 1)
				}
				propName = goName(propName)
				propNameLc := []rune(propName)
				propNameLc[0] = unicode.ToLower(propNameLc[0])

				relCollection := strings.Split(modelField.To, "/")[0]
				relCollectionName := goName(relCollection)
				relations = append(relations, relation{
					GoName:         goName(collectionName),
					PropName:       propName,
					PropNameLc:     string(propNameLc),
					RelCollection:  relCollection,
					CollectionName: relCollectionName,
					IdField:        fieldName,
					IdFieldGo:      goName(fieldName),
					Required:       modelField.Required,
					List:           modelField.Type == "relation-list",
				})
			}

			if modelField.Type == "generic-relation" {
				propName := goName(strings.Replace(fieldName+"$", "_id$", "", 1))
				propNameLc := []rune(propName)
				propNameLc[0] = unicode.ToLower(propNameLc[0])

				toCollections := map[string]string{}
				for _, c := range modelField.Relation().ToCollections() {
					toCollections[c.Collection] = goName(c.Collection)
				}

				relations = append(relations, relation{
					GoName:         goName(collectionName),
					PropName:       propName,
					PropNameLc:     string(propNameLc),
					RelCollection:  "IBaseModel",
					CollectionName: "IBaseModel",
					IdField:        fieldName,
					IdFieldGo:      goName(fieldName),
					Required:       modelField.Required,
					List:           false,
					Generic:        true,
					ToCollections:  toCollections,
				})
			}

			fields = append(fields, f)
		}

		sort.Slice(relations, func(i, j int) bool {
			return relations[i].PropName < relations[j].PropName
		})

		sort.Slice(fields, func(i, j int) bool {
			return fields[i].GoName < fields[j].GoName
		})

		collections = append(collections, collection{
			GoName:         goName(collectionName),
			CollectionName: collectionName,
			Fields:         fields,
			Relations:      relations,
			HasRelations:   len(relations) > 0,
		})
	}

	return collections, nil
}

type collection struct {
	GoName         string
	CollectionName string
	Fields         []field
	Relations      []relation
	HasRelations   bool
}

type field struct {
	GoName         string
	ValueType      string
	Collection     string
	CollectionName string
	FieldName      string
	Required       bool
	SingleRelation bool
}

type relation struct {
	GoName         string
	PropName       string
	PropNameLc     string
	RelCollection  string
	CollectionName string
	ToCollections  map[string]string
	IdField        string
	IdFieldGo      string
	Required       bool
	List           bool
	Generic        bool
}

func goName(name string) string {
	if name == "id" {
		return "ID"
	}

	name = strings.ReplaceAll(name, "_$", "")

	parts := strings.Split(name, "_")
	caser := cases.Title(language.English)
	for i := range parts {
		parts[i] = caser.String(parts[i])
	}
	name = strings.Join(parts, "")

	name = strings.ReplaceAll(name, "Id", "ID")
	return name
}

func valueType(modelsType string, required bool) string {
	if required {
		switch modelsType {
		case "number", "relation", "timestamp":
			return "int"

		case "string", "text", "HTMLStrict", "color", "HTMLPermissive", "generic-relation", "template", "decimal(6)":
			return "string"

		case "boolean":
			return "bool"

		case "float":
			return "float32"

		case "relation-list", "number[]":
			return "[]int"

		case "JSON":
			return "json.RawMessage"

		case "string[]", "generic-relation-list":
			return "[]string"

		default:
			panic(fmt.Sprintf("Unknown type %q", modelsType))
		}
	}

	switch modelsType {
	case "number", "relation", "timestamp":
		return "*int"

	case "string", "text", "HTMLStrict", "color", "HTMLPermissive", "generic-relation", "template", "decimal(6)":
		return "*string"

	case "boolean":
		return "*bool"

	case "float":
		return "*float32"

	case "relation-list", "number[]":
		return "[]int"

	case "JSON":
		return "json.RawMessage"

	case "string[]", "generic-relation-list":
		return "[]string"

	default:
		panic(fmt.Sprintf("Unknown type %q", modelsType))
	}
}
