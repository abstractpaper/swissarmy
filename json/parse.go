package json

import (
	"reflect"
	"strings"

	log "github.com/sirupsen/logrus"
)

// PrintOptions for Print function
type PrintOptions struct {
	Prepend          string
	TrimStrings      bool
	SkipEmptyStrings bool
	prependText      string
}

// Print is a recursive function that parses a nested
// json object all the way.
func Print(obj map[string]interface{}, opt PrintOptions) {
	if opt.prependText == "" {
		opt.prependText = opt.Prepend
	}
	opt.prependText = opt.prependText + " "

	for _, v := range obj {
		t := reflect.TypeOf(v)
		if t == nil {
			log.Info(opt.prependText, nil)
		} else {
			switch t.Kind() {
			case reflect.String:
				text := v.(string)
				if opt.TrimStrings {
					text = strings.TrimSpace(text)
				}
				// either text is not empty
				// or
				// SkipEmptyStrings is false
				if text != "" || opt.SkipEmptyStrings == false {
					log.Info(opt.prependText, text)
				}
			case reflect.Int:
				log.Info(opt.prependText, v.(int))
			case reflect.Bool:
				log.Info(opt.prependText, v.(bool))
			case reflect.Float64:
				log.Info(opt.prependText, v.(float64))
			default:
				o := opt
				o.prependText = o.prependText + o.Prepend
				Print(v.(map[string]interface{}), o)
			}
		}
	}
}
