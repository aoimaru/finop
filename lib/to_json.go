package lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type ToJ struct {
	Name       string   `json:"name"`
	Extentions []string `json:"extentions"`
}

func ToOrigin(filename string, extentions []string) ToJ {
	data := ToJ{
		Name:       filename,
		Extentions: extentions,
	}

	return data

}

func ToJson(objs *[]ToJ) ([]byte, error) {
	jsonBlob, err := json.Marshal(objs)
	if err != nil {
		return nil, err
	}
	return jsonBlob, nil
}

func ToClean(jsonBlob []uint8) {
	var buf bytes.Buffer
	err := json.Indent(&buf, []byte(jsonBlob), "", "  ")
	if err != nil {
		panic(err)
	}
	indentJson := buf.String()
	fmt.Println(indentJson)
}

func ToFile(directory string, datas []ToJ) {
	toDicFile := directory + ".json"
	file, _ := json.MarshalIndent(datas, "", " ")
	_ = ioutil.WriteFile(toDicFile, file, 0644)

}
