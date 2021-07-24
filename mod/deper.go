package mod

import (
	"fmt"
	"golang.org/x/mod/semver"
)

type Deponer struct {
	cfg        Cfg
	deprecates map[string]struct{}
	mods       map[string]string
}

func NewDeponer(cfg *Cfg) *Deponer {
	d := Deponer{}

	if len(cfg.Deprecates) > 0 {
		p := make(map[string]struct{})
		for _, v := range cfg.Deprecates {
			p[v] = struct{}{}
		}
		d.deprecates = p
	}
	if len(cfg.Mods) > 0 {
		t := make(map[string]string, 0)
		for _, v := range cfg.Mods {
			t[v.Name] = v.MinVer
		}
		d.mods = t
	}
	return &d
}

const (
	deprecatesStr     = "%s 已被弃用，请移除"
	blowMinimumVerStr = "%s - %s 低于最小设定版本 %s"
)

func (d *Deponer) IsValid(mod string, version string) error {
	if _, ok := d.deprecates[mod]; ok {
		return fmt.Errorf(deprecatesStr, mod)
	}
	if v, ok := d.mods[mod]; ok {
		r := semver.Compare(version, v)
		if r >= 0 {
			return nil
		} else {
			return fmt.Errorf(blowMinimumVerStr, mod, version, v)
		}
	}
	return nil
}

func (d *Deponer) IsDeprecate(mod string) bool {
	_, ok := d.deprecates[mod]
	return ok
}
