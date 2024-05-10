// Code generated by "stringer -type=TypeKind"; DO NOT EDIT.

package go_inspect

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[KNone-0]
	_ = x[KTFunc-1]
	_ = x[KTType-2]
	_ = x[KTConst-3]
	_ = x[KTVar-4]
	_ = x[KVar-5]
	_ = x[KFunc-6]
	_ = x[KNamed-7]
	_ = x[KStruct-8]
	_ = x[KMap-9]
	_ = x[KPointer-10]
	_ = x[KArray-11]
	_ = x[KSlice-12]
	_ = x[KInterface-13]
	_ = x[KChan-14]
	_ = x[KBasic-15]
	_ = x[KTypeParam-16]
	_ = x[KSignature-17]
	_ = x[KUnion-18]
	_ = x[KTuple-19]
	_ = x[KTerm-20]
	_ = x[KMField-21]
	_ = x[KMMapKey-22]
	_ = x[KMMapValue-23]
	_ = x[KMMethod-24]
	_ = x[KMParam-25]
	_ = x[KMResult-26]
	_ = x[KMEmbedded-27]
}

const _TypeKind_name = "KNoneKTFuncKTTypeKTConstKTVarKVarKFuncKNamedKStructKMapKPointerKArrayKSliceKInterfaceKChanKBasicKTypeParamKSignatureKUnionKTupleKTermKMFieldKMMapKeyKMMapValueKMMethodKMParamKMResultKMEmbedded"

var _TypeKind_index = [...]uint8{0, 5, 11, 17, 24, 29, 33, 38, 44, 51, 55, 63, 69, 75, 85, 90, 96, 106, 116, 122, 128, 133, 140, 148, 158, 166, 173, 181, 191}

func (i TypeKind) String() string {
	if i < 0 || i >= TypeKind(len(_TypeKind_index)-1) {
		return "TypeKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _TypeKind_name[_TypeKind_index[i]:_TypeKind_index[i+1]]
}