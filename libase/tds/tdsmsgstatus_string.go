// Code generated by "stringer -type=TDSMsgStatus"; DO NOT EDIT.

package tds

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[TDS_MSG_HASNOARGS-0]
	_ = x[TDS_MSG_HASARGS-1]
}

const _TDSMsgStatus_name = "TDS_MSG_HASNOARGSTDS_MSG_HASARGS"

var _TDSMsgStatus_index = [...]uint8{0, 17, 32}

func (i TDSMsgStatus) String() string {
	if i >= TDSMsgStatus(len(_TDSMsgStatus_index)-1) {
		return "TDSMsgStatus(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _TDSMsgStatus_name[_TDSMsgStatus_index[i]:_TDSMsgStatus_index[i+1]]
}