// Code generated by "stringer -type=Language"; DO NOT EDIT.

package cmd

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Any-0]
	_ = x[JSON-1]
	_ = x[YAML-2]
	_ = x[TOML-3]
}

const _Language_name = "AnyJSONYAMLTOML"

var _Language_index = [...]uint8{0, 3, 7, 11, 15}

func (i Language) String() string {
	if i < 0 || i >= Language(len(_Language_index)-1) {
		return "Language(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Language_name[_Language_index[i]:_Language_index[i+1]]
}
