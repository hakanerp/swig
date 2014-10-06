/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 3.0.3
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: ./../arrays.i

package arrays

import "unsafe"
import _ "runtime/cgo"

var _cgo_runtime_cgocall func(unsafe.Pointer, uintptr)



type _ unsafe.Pointer



type _swig_fnptr *byte
type _swig_memberptr *byte

const ARRAY_LEN int = 2
type Finger int
var _wrap_One unsafe.Pointer

func _swig_getOne() (_swig_ret int) {
	_swig_p := uintptr(unsafe.Pointer(&_swig_ret))
	_cgo_runtime_cgocall(_wrap_One, _swig_p)
	return
}
var One int = _swig_getOne()
var _wrap_Two unsafe.Pointer

func _swig_getTwo() (_swig_ret int) {
	_swig_p := uintptr(unsafe.Pointer(&_swig_ret))
	_cgo_runtime_cgocall(_wrap_Two, _swig_p)
	return
}
var Two int = _swig_getTwo()
var _wrap_Three unsafe.Pointer

func _swig_getThree() (_swig_ret int) {
	_swig_p := uintptr(unsafe.Pointer(&_swig_ret))
	_cgo_runtime_cgocall(_wrap_Three, _swig_p)
	return
}
var Three int = _swig_getThree()
var _wrap_Four unsafe.Pointer

func _swig_getFour() (_swig_ret int) {
	_swig_p := uintptr(unsafe.Pointer(&_swig_ret))
	_cgo_runtime_cgocall(_wrap_Four, _swig_p)
	return
}
var Four int = _swig_getFour()
var _wrap_Five unsafe.Pointer

func _swig_getFive() (_swig_ret int) {
	_swig_p := uintptr(unsafe.Pointer(&_swig_ret))
	_cgo_runtime_cgocall(_wrap_Five, _swig_p)
	return
}
var Five int = _swig_getFive()
type SwigcptrSimpleStruct uintptr

func (p SwigcptrSimpleStruct) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrSimpleStruct) SwigIsSimpleStruct() {
}

var _wrap_SimpleStruct_double_field_set unsafe.Pointer

func _swig_wrap_SimpleStruct_double_field_set(base SwigcptrSimpleStruct, _ float64) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_SimpleStruct_double_field_set, _swig_p)
	return
}

func (arg1 SwigcptrSimpleStruct) SetDouble_field(arg2 float64) {
	_swig_wrap_SimpleStruct_double_field_set(arg1, arg2)
}

var _wrap_SimpleStruct_double_field_get unsafe.Pointer

func _swig_wrap_SimpleStruct_double_field_get(base SwigcptrSimpleStruct) (_ float64) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_SimpleStruct_double_field_get, _swig_p)
	return
}

func (arg1 SwigcptrSimpleStruct) GetDouble_field() (_swig_ret float64) {
	return _swig_wrap_SimpleStruct_double_field_get(arg1)
}

var _wrap_new_SimpleStruct unsafe.Pointer

func _swig_wrap_new_SimpleStruct() (base SwigcptrSimpleStruct) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_SimpleStruct, _swig_p)
	return
}

func NewSimpleStruct() (_swig_ret SimpleStruct) {
	return _swig_wrap_new_SimpleStruct()
}

var _wrap_delete_SimpleStruct unsafe.Pointer

func _swig_wrap_delete_SimpleStruct(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_SimpleStruct, _swig_p)
	return
}

func DeleteSimpleStruct(arg1 SimpleStruct) {
	_swig_wrap_delete_SimpleStruct(arg1.Swigcptr())
}

type SimpleStruct interface {
	Swigcptr() uintptr
	SwigIsSimpleStruct()
	SetDouble_field(arg2 float64)
	GetDouble_field() (_swig_ret float64)
}

type SwigcptrArrayStruct uintptr

func (p SwigcptrArrayStruct) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrArrayStruct) SwigIsArrayStruct() {
}

var _wrap_ArrayStruct_array_c_set unsafe.Pointer

func _swig_wrap_ArrayStruct_array_c_set(base SwigcptrArrayStruct, _ string) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_ArrayStruct_array_c_set, _swig_p)
	return
}

func (arg1 SwigcptrArrayStruct) SetArray_c(arg2 string) {
	_swig_wrap_ArrayStruct_array_c_set(arg1, arg2)
}

var _wrap_ArrayStruct_array_c_get unsafe.Pointer

func _swig_wrap_ArrayStruct_array_c_get(base SwigcptrArrayStruct) (_ string) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_ArrayStruct_array_c_get, _swig_p)
	return
}

func (arg1 SwigcptrArrayStruct) GetArray_c() (_swig_ret string) {
	return _swig_wrap_ArrayStruct_array_c_get(arg1)
}

var _wrap_ArrayStruct_array_sc_set unsafe.Pointer

func _swig_wrap_ArrayStruct_array_sc_set(base SwigcptrArrayStruct, _ string) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_ArrayStruct_array_sc_set, _swig_p)
	return
}

func (arg1 SwigcptrArrayStruct) SetArray_sc(arg2 string) {
	_swig_wrap_ArrayStruct_array_sc_set(arg1, arg2)
}

var _wrap_ArrayStruct_array_sc_get unsafe.Pointer

func _swig_wrap_ArrayStruct_array_sc_get(base SwigcptrArrayStruct) (_ string) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_ArrayStruct_array_sc_get, _swig_p)
	return
}

func (arg1 SwigcptrArrayStruct) GetArray_sc() (_swig_ret string) {
	return _swig_wrap_ArrayStruct_array_sc_get(arg1)
}

var _wrap_ArrayStruct_array_uc_set unsafe.Pointer

func _swig_wrap_ArrayStruct_array_uc_set(base SwigcptrArrayStruct, _ string) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_ArrayStruct_array_uc_set, _swig_p)
	return
}

func (arg1 SwigcptrArrayStruct) SetArray_uc(arg2 string) {
	_swig_wrap_ArrayStruct_array_uc_set(arg1, arg2)
}

var _wrap_ArrayStruct_array_uc_get unsafe.Pointer

func _swig_wrap_ArrayStruct_array_uc_get(base SwigcptrArrayStruct) (_ string) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_ArrayStruct_array_uc_get, _swig_p)
	return
}

func (arg1 SwigcptrArrayStruct) GetArray_uc() (_swig_ret string) {
	return _swig_wrap_ArrayStruct_array_uc_get(arg1)
}

var _wrap_ArrayStruct_array_s_set unsafe.Pointer

func _swig_wrap_ArrayStruct_array_s_set(base SwigcptrArrayStruct, _ *int16) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_ArrayStruct_array_s_set, _swig_p)
	return
}

func (arg1 SwigcptrArrayStruct) SetArray_s(arg2 *int16) {
	_swig_wrap_ArrayStruct_array_s_set(arg1, arg2)
}

var _wrap_ArrayStruct_array_s_get unsafe.Pointer

func _swig_wrap_ArrayStruct_array_s_get(base SwigcptrArrayStruct) (_ *int16) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_ArrayStruct_array_s_get, _swig_p)
	return
}

func (arg1 SwigcptrArrayStruct) GetArray_s() (_swig_ret *int16) {
	return _swig_wrap_ArrayStruct_array_s_get(arg1)
}

var _wrap_ArrayStruct_array_us_set unsafe.Pointer

func _swig_wrap_ArrayStruct_array_us_set(base SwigcptrArrayStruct, _ *uint16) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_ArrayStruct_array_us_set, _swig_p)
	return
}

func (arg1 SwigcptrArrayStruct) SetArray_us(arg2 *uint16) {
	_swig_wrap_ArrayStruct_array_us_set(arg1, arg2)
}

var _wrap_ArrayStruct_array_us_get unsafe.Pointer

func _swig_wrap_ArrayStruct_array_us_get(base SwigcptrArrayStruct) (_ *uint16) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_ArrayStruct_array_us_get, _swig_p)
	return
}

func (arg1 SwigcptrArrayStruct) GetArray_us() (_swig_ret *uint16) {
	return _swig_wrap_ArrayStruct_array_us_get(arg1)
}

var _wrap_ArrayStruct_array_i_set unsafe.Pointer

func _swig_wrap_ArrayStruct_array_i_set(base SwigcptrArrayStruct, _ *int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_ArrayStruct_array_i_set, _swig_p)
	return
}

func (arg1 SwigcptrArrayStruct) SetArray_i(arg2 *int) {
	_swig_wrap_ArrayStruct_array_i_set(arg1, arg2)
}

var _wrap_ArrayStruct_array_i_get unsafe.Pointer

func _swig_wrap_ArrayStruct_array_i_get(base SwigcptrArrayStruct) (_ *int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_ArrayStruct_array_i_get, _swig_p)
	return
}

func (arg1 SwigcptrArrayStruct) GetArray_i() (_swig_ret *int) {
	return _swig_wrap_ArrayStruct_array_i_get(arg1)
}

var _wrap_ArrayStruct_array_ui_set unsafe.Pointer

func _swig_wrap_ArrayStruct_array_ui_set(base SwigcptrArrayStruct, _ *uint) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_ArrayStruct_array_ui_set, _swig_p)
	return
}

func (arg1 SwigcptrArrayStruct) SetArray_ui(arg2 *uint) {
	_swig_wrap_ArrayStruct_array_ui_set(arg1, arg2)
}

var _wrap_ArrayStruct_array_ui_get unsafe.Pointer

func _swig_wrap_ArrayStruct_array_ui_get(base SwigcptrArrayStruct) (_ *uint) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_ArrayStruct_array_ui_get, _swig_p)
	return
}

func (arg1 SwigcptrArrayStruct) GetArray_ui() (_swig_ret *uint) {
	return _swig_wrap_ArrayStruct_array_ui_get(arg1)
}

var _wrap_ArrayStruct_array_l_set unsafe.Pointer

func _swig_wrap_ArrayStruct_array_l_set(base SwigcptrArrayStruct, _ *int64) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_ArrayStruct_array_l_set, _swig_p)
	return
}

func (arg1 SwigcptrArrayStruct) SetArray_l(arg2 *int64) {
	_swig_wrap_ArrayStruct_array_l_set(arg1, arg2)
}

var _wrap_ArrayStruct_array_l_get unsafe.Pointer

func _swig_wrap_ArrayStruct_array_l_get(base SwigcptrArrayStruct) (_ *int64) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_ArrayStruct_array_l_get, _swig_p)
	return
}

func (arg1 SwigcptrArrayStruct) GetArray_l() (_swig_ret *int64) {
	return _swig_wrap_ArrayStruct_array_l_get(arg1)
}

var _wrap_ArrayStruct_array_ul_set unsafe.Pointer

func _swig_wrap_ArrayStruct_array_ul_set(base SwigcptrArrayStruct, _ *uint64) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_ArrayStruct_array_ul_set, _swig_p)
	return
}

func (arg1 SwigcptrArrayStruct) SetArray_ul(arg2 *uint64) {
	_swig_wrap_ArrayStruct_array_ul_set(arg1, arg2)
}

var _wrap_ArrayStruct_array_ul_get unsafe.Pointer

func _swig_wrap_ArrayStruct_array_ul_get(base SwigcptrArrayStruct) (_ *uint64) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_ArrayStruct_array_ul_get, _swig_p)
	return
}

func (arg1 SwigcptrArrayStruct) GetArray_ul() (_swig_ret *uint64) {
	return _swig_wrap_ArrayStruct_array_ul_get(arg1)
}

var _wrap_ArrayStruct_array_ll_set unsafe.Pointer

func _swig_wrap_ArrayStruct_array_ll_set(base SwigcptrArrayStruct, _ *int64) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_ArrayStruct_array_ll_set, _swig_p)
	return
}

func (arg1 SwigcptrArrayStruct) SetArray_ll(arg2 *int64) {
	_swig_wrap_ArrayStruct_array_ll_set(arg1, arg2)
}

var _wrap_ArrayStruct_array_ll_get unsafe.Pointer

func _swig_wrap_ArrayStruct_array_ll_get(base SwigcptrArrayStruct) (_ *int64) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_ArrayStruct_array_ll_get, _swig_p)
	return
}

func (arg1 SwigcptrArrayStruct) GetArray_ll() (_swig_ret *int64) {
	return _swig_wrap_ArrayStruct_array_ll_get(arg1)
}

var _wrap_ArrayStruct_array_f_set unsafe.Pointer

func _swig_wrap_ArrayStruct_array_f_set(base SwigcptrArrayStruct, _ *float32) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_ArrayStruct_array_f_set, _swig_p)
	return
}

func (arg1 SwigcptrArrayStruct) SetArray_f(arg2 *float32) {
	_swig_wrap_ArrayStruct_array_f_set(arg1, arg2)
}

var _wrap_ArrayStruct_array_f_get unsafe.Pointer

func _swig_wrap_ArrayStruct_array_f_get(base SwigcptrArrayStruct) (_ *float32) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_ArrayStruct_array_f_get, _swig_p)
	return
}

func (arg1 SwigcptrArrayStruct) GetArray_f() (_swig_ret *float32) {
	return _swig_wrap_ArrayStruct_array_f_get(arg1)
}

var _wrap_ArrayStruct_array_d_set unsafe.Pointer

func _swig_wrap_ArrayStruct_array_d_set(base SwigcptrArrayStruct, _ *float64) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_ArrayStruct_array_d_set, _swig_p)
	return
}

func (arg1 SwigcptrArrayStruct) SetArray_d(arg2 *float64) {
	_swig_wrap_ArrayStruct_array_d_set(arg1, arg2)
}

var _wrap_ArrayStruct_array_d_get unsafe.Pointer

func _swig_wrap_ArrayStruct_array_d_get(base SwigcptrArrayStruct) (_ *float64) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_ArrayStruct_array_d_get, _swig_p)
	return
}

func (arg1 SwigcptrArrayStruct) GetArray_d() (_swig_ret *float64) {
	return _swig_wrap_ArrayStruct_array_d_get(arg1)
}

var _wrap_ArrayStruct_array_struct_set unsafe.Pointer

func _swig_wrap_ArrayStruct_array_struct_set(base SwigcptrArrayStruct, _ uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_ArrayStruct_array_struct_set, _swig_p)
	return
}

func (arg1 SwigcptrArrayStruct) SetArray_struct(arg2 SimpleStruct) {
	_swig_wrap_ArrayStruct_array_struct_set(arg1, arg2.Swigcptr())
}

var _wrap_ArrayStruct_array_struct_get unsafe.Pointer

func _swig_wrap_ArrayStruct_array_struct_get(base SwigcptrArrayStruct) (_ SwigcptrSimpleStruct) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_ArrayStruct_array_struct_get, _swig_p)
	return
}

func (arg1 SwigcptrArrayStruct) GetArray_struct() (_swig_ret SimpleStruct) {
	return _swig_wrap_ArrayStruct_array_struct_get(arg1)
}

var _wrap_ArrayStruct_array_structpointers_set unsafe.Pointer

func _swig_wrap_ArrayStruct_array_structpointers_set(base SwigcptrArrayStruct, _ uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_ArrayStruct_array_structpointers_set, _swig_p)
	return
}

func (arg1 SwigcptrArrayStruct) SetArray_structpointers(arg2 SimpleStruct) {
	_swig_wrap_ArrayStruct_array_structpointers_set(arg1, arg2.Swigcptr())
}

var _wrap_ArrayStruct_array_structpointers_get unsafe.Pointer

func _swig_wrap_ArrayStruct_array_structpointers_get(base SwigcptrArrayStruct) (_ SwigcptrSimpleStruct) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_ArrayStruct_array_structpointers_get, _swig_p)
	return
}

func (arg1 SwigcptrArrayStruct) GetArray_structpointers() (_swig_ret SimpleStruct) {
	return _swig_wrap_ArrayStruct_array_structpointers_get(arg1)
}

var _wrap_ArrayStruct_array_ipointers_set unsafe.Pointer

func _swig_wrap_ArrayStruct_array_ipointers_set(base SwigcptrArrayStruct, _ **int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_ArrayStruct_array_ipointers_set, _swig_p)
	return
}

func (arg1 SwigcptrArrayStruct) SetArray_ipointers(arg2 **int) {
	_swig_wrap_ArrayStruct_array_ipointers_set(arg1, arg2)
}

var _wrap_ArrayStruct_array_ipointers_get unsafe.Pointer

func _swig_wrap_ArrayStruct_array_ipointers_get(base SwigcptrArrayStruct) (_ **int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_ArrayStruct_array_ipointers_get, _swig_p)
	return
}

func (arg1 SwigcptrArrayStruct) GetArray_ipointers() (_swig_ret **int) {
	return _swig_wrap_ArrayStruct_array_ipointers_get(arg1)
}

var _wrap_ArrayStruct_array_enum_set unsafe.Pointer

func _swig_wrap_ArrayStruct_array_enum_set(base SwigcptrArrayStruct, _ *Finger) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_ArrayStruct_array_enum_set, _swig_p)
	return
}

func (arg1 SwigcptrArrayStruct) SetArray_enum(arg2 *Finger) {
	_swig_wrap_ArrayStruct_array_enum_set(arg1, arg2)
}

var _wrap_ArrayStruct_array_enum_get unsafe.Pointer

func _swig_wrap_ArrayStruct_array_enum_get(base SwigcptrArrayStruct) (_ *Finger) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_ArrayStruct_array_enum_get, _swig_p)
	return
}

func (arg1 SwigcptrArrayStruct) GetArray_enum() (_swig_ret *Finger) {
	return _swig_wrap_ArrayStruct_array_enum_get(arg1)
}

var _wrap_ArrayStruct_array_enumpointers_set unsafe.Pointer

func _swig_wrap_ArrayStruct_array_enumpointers_set(base SwigcptrArrayStruct, _ **Finger) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_ArrayStruct_array_enumpointers_set, _swig_p)
	return
}

func (arg1 SwigcptrArrayStruct) SetArray_enumpointers(arg2 **Finger) {
	_swig_wrap_ArrayStruct_array_enumpointers_set(arg1, arg2)
}

var _wrap_ArrayStruct_array_enumpointers_get unsafe.Pointer

func _swig_wrap_ArrayStruct_array_enumpointers_get(base SwigcptrArrayStruct) (_ **Finger) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_ArrayStruct_array_enumpointers_get, _swig_p)
	return
}

func (arg1 SwigcptrArrayStruct) GetArray_enumpointers() (_swig_ret **Finger) {
	return _swig_wrap_ArrayStruct_array_enumpointers_get(arg1)
}

var _wrap_ArrayStruct_array_const_i_get unsafe.Pointer

func _swig_wrap_ArrayStruct_array_const_i_get(base SwigcptrArrayStruct) (_ *int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_ArrayStruct_array_const_i_get, _swig_p)
	return
}

func (arg1 SwigcptrArrayStruct) GetArray_const_i() (_swig_ret *int) {
	return _swig_wrap_ArrayStruct_array_const_i_get(arg1)
}

var _wrap_new_ArrayStruct unsafe.Pointer

func _swig_wrap_new_ArrayStruct() (base SwigcptrArrayStruct) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_ArrayStruct, _swig_p)
	return
}

func NewArrayStruct() (_swig_ret ArrayStruct) {
	return _swig_wrap_new_ArrayStruct()
}

var _wrap_delete_ArrayStruct unsafe.Pointer

func _swig_wrap_delete_ArrayStruct(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_ArrayStruct, _swig_p)
	return
}

func DeleteArrayStruct(arg1 ArrayStruct) {
	_swig_wrap_delete_ArrayStruct(arg1.Swigcptr())
}

type ArrayStruct interface {
	Swigcptr() uintptr
	SwigIsArrayStruct()
	SetArray_c(arg2 string)
	GetArray_c() (_swig_ret string)
	SetArray_sc(arg2 string)
	GetArray_sc() (_swig_ret string)
	SetArray_uc(arg2 string)
	GetArray_uc() (_swig_ret string)
	SetArray_s(arg2 *int16)
	GetArray_s() (_swig_ret *int16)
	SetArray_us(arg2 *uint16)
	GetArray_us() (_swig_ret *uint16)
	SetArray_i(arg2 *int)
	GetArray_i() (_swig_ret *int)
	SetArray_ui(arg2 *uint)
	GetArray_ui() (_swig_ret *uint)
	SetArray_l(arg2 *int64)
	GetArray_l() (_swig_ret *int64)
	SetArray_ul(arg2 *uint64)
	GetArray_ul() (_swig_ret *uint64)
	SetArray_ll(arg2 *int64)
	GetArray_ll() (_swig_ret *int64)
	SetArray_f(arg2 *float32)
	GetArray_f() (_swig_ret *float32)
	SetArray_d(arg2 *float64)
	GetArray_d() (_swig_ret *float64)
	SetArray_struct(arg2 SimpleStruct)
	GetArray_struct() (_swig_ret SimpleStruct)
	SetArray_structpointers(arg2 SimpleStruct)
	GetArray_structpointers() (_swig_ret SimpleStruct)
	SetArray_ipointers(arg2 **int)
	GetArray_ipointers() (_swig_ret **int)
	SetArray_enum(arg2 *Finger)
	GetArray_enum() (_swig_ret *Finger)
	SetArray_enumpointers(arg2 **Finger)
	GetArray_enumpointers() (_swig_ret **Finger)
	GetArray_const_i() (_swig_ret *int)
}

var _wrap_fn_taking_arrays unsafe.Pointer

func _swig_wrap_fn_taking_arrays(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_fn_taking_arrays, _swig_p)
	return
}

func Fn_taking_arrays(arg1 SimpleStruct) {
	_swig_wrap_fn_taking_arrays(arg1.Swigcptr())
}

var _wrap_newintpointer unsafe.Pointer

func Newintpointer() (_swig_ret *int) {
	_swig_p := uintptr(unsafe.Pointer(&_swig_ret))
	_cgo_runtime_cgocall(_wrap_newintpointer, _swig_p)
	return
}
var _wrap_setintfrompointer unsafe.Pointer

func _swig_wrap_setintfrompointer(base *int, _ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_setintfrompointer, _swig_p)
	return
}

func Setintfrompointer(arg1 *int, arg2 int) {
	_swig_wrap_setintfrompointer(arg1, arg2)
}

var _wrap_getintfrompointer unsafe.Pointer

func Getintfrompointer(arg1 *int) (_swig_ret int) {
	_swig_p := uintptr(unsafe.Pointer(&arg1))
	_cgo_runtime_cgocall(_wrap_getintfrompointer, _swig_p)
	return
}
var _wrap_array_pointer_func unsafe.Pointer

func _swig_wrap_array_pointer_func(base **int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_array_pointer_func, _swig_p)
	return
}

func Array_pointer_func(arg1 **int) {
	_swig_wrap_array_pointer_func(arg1)
}

type SwigcptrCartPoseData_t uintptr

func (p SwigcptrCartPoseData_t) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrCartPoseData_t) SwigIsCartPoseData_t() {
}

var _wrap_CartPoseData_t_p_set unsafe.Pointer

func _swig_wrap_CartPoseData_t_p_set(base SwigcptrCartPoseData_t, _ *float32) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_CartPoseData_t_p_set, _swig_p)
	return
}

func (arg1 SwigcptrCartPoseData_t) SetP(arg2 *float32) {
	_swig_wrap_CartPoseData_t_p_set(arg1, arg2)
}

var _wrap_CartPoseData_t_p_get unsafe.Pointer

func _swig_wrap_CartPoseData_t_p_get(base SwigcptrCartPoseData_t) (_ *float32) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_CartPoseData_t_p_get, _swig_p)
	return
}

func (arg1 SwigcptrCartPoseData_t) GetP() (_swig_ret *float32) {
	return _swig_wrap_CartPoseData_t_p_get(arg1)
}

var _wrap_new_CartPoseData_t unsafe.Pointer

func _swig_wrap_new_CartPoseData_t() (base SwigcptrCartPoseData_t) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_CartPoseData_t, _swig_p)
	return
}

func NewCartPoseData_t() (_swig_ret CartPoseData_t) {
	return _swig_wrap_new_CartPoseData_t()
}

var _wrap_delete_CartPoseData_t unsafe.Pointer

func _swig_wrap_delete_CartPoseData_t(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_CartPoseData_t, _swig_p)
	return
}

func DeleteCartPoseData_t(arg1 CartPoseData_t) {
	_swig_wrap_delete_CartPoseData_t(arg1.Swigcptr())
}

type CartPoseData_t interface {
	Swigcptr() uintptr
	SwigIsCartPoseData_t()
	SetP(arg2 *float32)
	GetP() (_swig_ret *float32)
}


type SwigcptrVoid uintptr
type Void interface {
	Swigcptr() uintptr;
}
func (p SwigcptrVoid) Swigcptr() uintptr {
	return uintptr(p)
}

