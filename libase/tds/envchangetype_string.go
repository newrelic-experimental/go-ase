// Code generated by "stringer -type=EnvChangeType"; DO NOT EDIT.

package tds

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[TDS_ENV_DB-1]
	_ = x[TDS_ENV_LANG-2]
	_ = x[TDS_ENV_CHARSET-3]
	_ = x[TDS_ENV_PACKSIZE-4]
}

const _EnvChangeType_name = "TDS_ENV_DBTDS_ENV_LANGTDS_ENV_CHARSETTDS_ENV_PACKSIZE"

var _EnvChangeType_index = [...]uint8{0, 10, 22, 37, 53}

func (i EnvChangeType) String() string {
	i -= 1
	if i >= EnvChangeType(len(_EnvChangeType_index)-1) {
		return "EnvChangeType(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _EnvChangeType_name[_EnvChangeType_index[i]:_EnvChangeType_index[i+1]]
}