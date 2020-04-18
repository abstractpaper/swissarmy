package json

import (
	"testing"
)

func TestPrint(t *testing.T) {
	obj := map[string]interface{}{
		"a": "1",
		"c": map[string]interface{}{
			"d": 3,
			"e": map[string]interface{}{
				"f": true,
				"g": nil,
			},
			"i": 3.6,
		},
	}

	opt := PrintOptions{
		Prepend: ">",
	}
	Print(obj, opt)
}

func TestPrintSkipEmptyStrings(t *testing.T) {
	obj := map[string]interface{}{
		"a": "1",
		"b": "",
	}

	opt := PrintOptions{
		SkipEmptyStrings: true,
	}
	Print(obj, opt)
}

func TestPrintStripStrings(t *testing.T) {
	obj := map[string]interface{}{
		"a": " 1",
		"b": "2\n",
	}

	opt := PrintOptions{
		TrimStrings: true,
	}
	Print(obj, opt)
}
