package mod

import (
	"encoding/json"
	"io/ioutil"
)

/**
{
	"mods":[
		{
"name":"xxxx",
"min_ver":"v1.1.0",
"deprecated": true
}

	],
"deprecated_mod":{

}
}
*/
type Cfg struct {
	Mods       []*CfgMod `json:"mods"`
	Deprecates []string  `json:"deprecates"`
}

type CfgMod struct {
	Name       string `json:"name"`
	MinVer     string `json:"min_ver"`
	Deprecated *bool  `json:"deprecated"`
}

func Config(path string) (*Cfg, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg = &Cfg{}
	if err := json.Unmarshal(bytes, cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
