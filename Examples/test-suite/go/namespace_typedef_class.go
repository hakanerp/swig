/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 3.0.3
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: ./../namespace_typedef_class.i

package namespace_typedef_class

import "unsafe"
import _ "runtime/cgo"

var _cgo_runtime_cgocall func(unsafe.Pointer, uintptr)



type _ unsafe.Pointer



type _swig_fnptr *byte
type _swig_memberptr *byte

type SwigcptrS1 uintptr

func (p SwigcptrS1) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrS1) SwigIsS1() {
}

var _wrap_S1_n_set unsafe.Pointer

func _swig_wrap_S1_n_set(base SwigcptrS1, _ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_S1_n_set, _swig_p)
	return
}

func (arg1 SwigcptrS1) SetN(arg2 int) {
	_swig_wrap_S1_n_set(arg1, arg2)
}

var _wrap_S1_n_get unsafe.Pointer

func _swig_wrap_S1_n_get(base SwigcptrS1) (_ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_S1_n_get, _swig_p)
	return
}

func (arg1 SwigcptrS1) GetN() (_swig_ret int) {
	return _swig_wrap_S1_n_get(arg1)
}

var _wrap_new_S1 unsafe.Pointer

func _swig_wrap_new_S1() (base SwigcptrS1) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_S1, _swig_p)
	return
}

func NewS1() (_swig_ret S1) {
	return _swig_wrap_new_S1()
}

var _wrap_delete_S1 unsafe.Pointer

func _swig_wrap_delete_S1(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_S1, _swig_p)
	return
}

func DeleteS1(arg1 S1) {
	_swig_wrap_delete_S1(arg1.Swigcptr())
}

type S1 interface {
	Swigcptr() uintptr
	SwigIsS1()
	SetN(arg2 int)
	GetN() (_swig_ret int)
}

type SwigcptrS2 uintptr

func (p SwigcptrS2) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrS2) SwigIsS2() {
}

var _wrap_S2_n_set unsafe.Pointer

func _swig_wrap_S2_n_set(base SwigcptrS2, _ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_S2_n_set, _swig_p)
	return
}

func (arg1 SwigcptrS2) SetN(arg2 int) {
	_swig_wrap_S2_n_set(arg1, arg2)
}

var _wrap_S2_n_get unsafe.Pointer

func _swig_wrap_S2_n_get(base SwigcptrS2) (_ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_S2_n_get, _swig_p)
	return
}

func (arg1 SwigcptrS2) GetN() (_swig_ret int) {
	return _swig_wrap_S2_n_get(arg1)
}

var _wrap_new_S2 unsafe.Pointer

func _swig_wrap_new_S2() (base SwigcptrS2) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_S2, _swig_p)
	return
}

func NewS2() (_swig_ret S2) {
	return _swig_wrap_new_S2()
}

var _wrap_delete_S2 unsafe.Pointer

func _swig_wrap_delete_S2(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_S2, _swig_p)
	return
}

func DeleteS2(arg1 S2) {
	_swig_wrap_delete_S2(arg1.Swigcptr())
}

type S2 interface {
	Swigcptr() uintptr
	SwigIsS2()
	SetN(arg2 int)
	GetN() (_swig_ret int)
}


type SwigcptrVoid uintptr
type Void interface {
	Swigcptr() uintptr;
}
func (p SwigcptrVoid) Swigcptr() uintptr {
	return uintptr(p)
}

