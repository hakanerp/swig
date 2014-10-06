/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 3.0.3
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: ./../typemap_variables.i

package typemap_variables

import "unsafe"
import _ "runtime/cgo"

var _cgo_runtime_cgocall func(unsafe.Pointer, uintptr)



type _ unsafe.Pointer



type _swig_fnptr *byte
type _swig_memberptr *byte

var _wrap_globul_set unsafe.Pointer

func _swig_wrap_globul_set(base int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_globul_set, _swig_p)
	return
}

func SetGlobul(arg1 int) {
	_swig_wrap_globul_set(arg1)
}

var _wrap_globul_get unsafe.Pointer

func GetGlobul() (_swig_ret int) {
	_swig_p := uintptr(unsafe.Pointer(&_swig_ret))
	_cgo_runtime_cgocall(_wrap_globul_get, _swig_p)
	return
}
var _wrap_nspace_set unsafe.Pointer

func _swig_wrap_nspace_set(base int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_nspace_set, _swig_p)
	return
}

func SetNspace(arg1 int) {
	_swig_wrap_nspace_set(arg1)
}

var _wrap_nspace_get unsafe.Pointer

func GetNspace() (_swig_ret int) {
	_swig_p := uintptr(unsafe.Pointer(&_swig_ret))
	_cgo_runtime_cgocall(_wrap_nspace_get, _swig_p)
	return
}
type SwigcptrStruct uintptr

func (p SwigcptrStruct) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrStruct) SwigIsStruct() {
}

var _wrap_Struct_member_set unsafe.Pointer

func _swig_wrap_Struct_member_set(base SwigcptrStruct, _ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_Struct_member_set, _swig_p)
	return
}

func (arg1 SwigcptrStruct) SetMember(arg2 int) {
	_swig_wrap_Struct_member_set(arg1, arg2)
}

var _wrap_Struct_member_get unsafe.Pointer

func _swig_wrap_Struct_member_get(base SwigcptrStruct) (_ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_Struct_member_get, _swig_p)
	return
}

func (arg1 SwigcptrStruct) GetMember() (_swig_ret int) {
	return _swig_wrap_Struct_member_get(arg1)
}

var _wrap_Struct_smember_set unsafe.Pointer

func _swig_wrap_Struct_smember_set(base int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_Struct_smember_set, _swig_p)
	return
}

func SetStructSmember(arg1 int) {
	_swig_wrap_Struct_smember_set(arg1)
}

var _wrap_Struct_smember_get unsafe.Pointer

func GetStructSmember() (_swig_ret int) {
	_swig_p := uintptr(unsafe.Pointer(&_swig_ret))
	_cgo_runtime_cgocall(_wrap_Struct_smember_get, _swig_p)
	return
}
var _wrap_new_Struct unsafe.Pointer

func _swig_wrap_new_Struct() (base SwigcptrStruct) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_Struct, _swig_p)
	return
}

func NewStruct() (_swig_ret Struct) {
	return _swig_wrap_new_Struct()
}

var _wrap_delete_Struct unsafe.Pointer

func _swig_wrap_delete_Struct(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_Struct, _swig_p)
	return
}

func DeleteStruct(arg1 Struct) {
	_swig_wrap_delete_Struct(arg1.Swigcptr())
}

type Struct interface {
	Swigcptr() uintptr
	SwigIsStruct()
	SetMember(arg2 int)
	GetMember() (_swig_ret int)
}


type SwigcptrVoid uintptr
type Void interface {
	Swigcptr() uintptr;
}
func (p SwigcptrVoid) Swigcptr() uintptr {
	return uintptr(p)
}

