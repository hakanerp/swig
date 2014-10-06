/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 3.0.3
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: ./../preproc_defined.i

package preproc_defined

import "unsafe"
import _ "runtime/cgo"

var _cgo_runtime_cgocall func(unsafe.Pointer, uintptr)



type _ unsafe.Pointer



type _swig_fnptr *byte
type _swig_memberptr *byte

var _wrap_call_checking unsafe.Pointer

func Call_checking() (_swig_ret int) {
	_swig_p := uintptr(unsafe.Pointer(&_swig_ret))
	_cgo_runtime_cgocall(_wrap_call_checking, _swig_p)
	return
}
var _wrap_thing unsafe.Pointer

func _swig_wrap_thing(base int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_thing, _swig_p)
	return
}

func Thing(arg1 int) {
	_swig_wrap_thing(arg1)
}

var _wrap_stuff unsafe.Pointer

func _swig_wrap_stuff(base int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_stuff, _swig_p)
	return
}

func Stuff(arg1 int) {
	_swig_wrap_stuff(arg1)
}

type SwigcptrDefined uintptr

func (p SwigcptrDefined) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrDefined) SwigIsDefined() {
}

var _wrap_Defined_defined_set unsafe.Pointer

func _swig_wrap_Defined_defined_set(base SwigcptrDefined, _ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_Defined_defined_set, _swig_p)
	return
}

func (arg1 SwigcptrDefined) SetDefined(arg2 int) {
	_swig_wrap_Defined_defined_set(arg1, arg2)
}

var _wrap_Defined_defined_get unsafe.Pointer

func _swig_wrap_Defined_defined_get(base SwigcptrDefined) (_ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_Defined_defined_get, _swig_p)
	return
}

func (arg1 SwigcptrDefined) GetDefined() (_swig_ret int) {
	return _swig_wrap_Defined_defined_get(arg1)
}

var _wrap_new_Defined unsafe.Pointer

func _swig_wrap_new_Defined() (base SwigcptrDefined) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_Defined, _swig_p)
	return
}

func NewDefined() (_swig_ret Defined) {
	return _swig_wrap_new_Defined()
}

var _wrap_delete_Defined unsafe.Pointer

func _swig_wrap_delete_Defined(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_Defined, _swig_p)
	return
}

func DeleteDefined(arg1 Defined) {
	_swig_wrap_delete_Defined(arg1.Swigcptr())
}

type Defined interface {
	Swigcptr() uintptr
	SwigIsDefined()
	SetDefined(arg2 int)
	GetDefined() (_swig_ret int)
}

var _wrap_bumpf unsafe.Pointer

func _swig_wrap_bumpf(base int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_bumpf, _swig_p)
	return
}

func Bumpf(arg1 int) {
	_swig_wrap_bumpf(arg1)
}


type SwigcptrVoid uintptr
type Void interface {
	Swigcptr() uintptr;
}
func (p SwigcptrVoid) Swigcptr() uintptr {
	return uintptr(p)
}

