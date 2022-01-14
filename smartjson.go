/*  Smart JSON functions - Helper functions to work with JSON
    (C) 2021-2022 Péter Deák (hyper80@gmail.com)
    License: Apache 2.0
*/

// smartjson is a go package to handle parsed json files more confortable.
// It provides helper functions to parse, query and covert the json data
package smartjson

import (
	"encoding/json"

	"github.com/hyper-prog/smartjsonyamlstub"
)

// SmartJSON holds the parsed data. You can call the smartjson functions on this structure
type SmartJSON struct {
	smartjsonyamlstub.SmartJsonYamlBase
}

// ParseJSON parse the raw json data (read from file) and returns a SmartJSON structure
func ParseJSON(rawdata []byte) (SmartJSON, error) {
	s := SmartJSON{}
	s.Config.InitConfig()
	err := json.Unmarshal(rawdata, &s.ParsedData)
	s.ParsedFrom = "json"
	return s, err
}

// String generate an indented json string
func (smartJson SmartJSON) String() string {
	return smartJson.JsonIndented()
}

// GetSubjsonByPath returns an another SmartJSON struct which holds a part of the parsed json specified by path
func (smartJson SmartJSON) GetSubjsonByPath(path string) (SmartJSON, string) {
	cd, str := smartJson.GetSubtreeByPath(path)
	rd := SmartJSON{}
	rd.Config = cd.Config
	rd.ParsedData = cd.ParsedData
	return rd, str
}
