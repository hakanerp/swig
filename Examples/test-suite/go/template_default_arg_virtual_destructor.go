/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 3.0.3
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: ./../template_default_arg_virtual_destructor.i

package template_default_arg_virtual_destructor

import "unsafe"
import _ "runtime/cgo"

var _cgo_runtime_cgocall func(unsafe.Pointer, uintptr)



type _ unsafe.Pointer



type _swig_fnptr *byte
type _swig_memberptr *byte

type SwigcptrA uintptr

func (p SwigcptrA) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrA) SwigIsA() {
}

var _wrap_new_A unsafe.Pointer

func _swig_wrap_new_A() (base SwigcptrA) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_A, _swig_p)
	return
}

func NewA() (_swig_ret A) {
	return _swig_wrap_new_A()
}

var _wrap_delete_A unsafe.Pointer

func _swig_wrap_delete_A(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_A, _swig_p)
	return
}

func DeleteA(arg1 A) {
	_swig_wrap_delete_A(arg1.Swigcptr())
}

type A interface {
	Swigcptr() uintptr
	SwigIsA()
}

type SwigcptrB_AF uintptr

func (p SwigcptrB_AF) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrB_AF) SwigIsB_AF() {
}

var _wrap_new_B_AF unsafe.Pointer

func _swig_wrap_new_B_AF(base float32) (_ SwigcptrB_AF) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_B_AF, _swig_p)
	return
}

func NewB_AF(arg1 float32) (_swig_ret B_AF) {
	return _swig_wrap_new_B_AF(arg1)
}

var _wrap_delete_B_AF unsafe.Pointer

func _swig_wrap_delete_B_AF(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_B_AF, _swig_p)
	return
}

func DeleteB_AF(arg1 B_AF) {
	_swig_wrap_delete_B_AF(arg1.Swigcptr())
}

type B_AF interface {
	Swigcptr() uintptr
	SwigIsB_AF()
}

type SwigcptrB_A uintptr

func (p SwigcptrB_A) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrB_A) SwigIsB_A() {
}

var _wrap_new_B_A unsafe.Pointer

func _swig_wrap_new_B_A(base int, _ int) (_ SwigcptrB_A) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_B_A, _swig_p)
	return
}

func NewB_A(arg1 int, arg2 int) (_swig_ret B_A) {
	return _swig_wrap_new_B_A(arg1, arg2)
}

var _wrap_delete_B_A unsafe.Pointer

func _swig_wrap_delete_B_A(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_B_A, _swig_p)
	return
}

func DeleteB_A(arg1 B_A) {
	_swig_wrap_delete_B_A(arg1.Swigcptr())
}

type B_A interface {
	Swigcptr() uintptr
	SwigIsB_A()
}


type SwigcptrVoid uintptr
type Void interface {
	Swigcptr() uintptr;
}
func (p SwigcptrVoid) Swigcptr() uintptr {
	return uintptr(p)
}

