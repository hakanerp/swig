/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 3.0.3
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: ./../template_whitespace.i

package template_whitespace

import "unsafe"
import _ "runtime/cgo"

var _cgo_runtime_cgocall func(unsafe.Pointer, uintptr)



type _ unsafe.Pointer



type _swig_fnptr *byte
type _swig_memberptr *byte

var _wrap_foo unsafe.Pointer

func _swig_wrap_foo(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_foo, _swig_p)
	return
}

func Foo(arg1 Vector_Sl_int_Sg_) {
	_swig_wrap_foo(arg1.Swigcptr())
}

var _wrap_bar unsafe.Pointer

func _swig_wrap_bar(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_bar, _swig_p)
	return
}

func Bar(arg1 Vector_Sl_unsigned_SS_int_Sg_) {
	_swig_wrap_bar(arg1.Swigcptr())
}

var _wrap_baz unsafe.Pointer

func _swig_wrap_baz(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_baz, _swig_p)
	return
}

func Baz(arg1 Map_Sl_int_Sc_int_Sg_) {
	_swig_wrap_baz(arg1.Swigcptr())
}


type SwigcptrVector_Sl_unsigned_SS_int_Sg_ uintptr
type Vector_Sl_unsigned_SS_int_Sg_ interface {
	Swigcptr() uintptr;
}
func (p SwigcptrVector_Sl_unsigned_SS_int_Sg_) Swigcptr() uintptr {
	return uintptr(p)
}

type SwigcptrVector_Sl_int_Sg_ uintptr
type Vector_Sl_int_Sg_ interface {
	Swigcptr() uintptr;
}
func (p SwigcptrVector_Sl_int_Sg_) Swigcptr() uintptr {
	return uintptr(p)
}

type SwigcptrMap_Sl_int_Sc_int_Sg_ uintptr
type Map_Sl_int_Sc_int_Sg_ interface {
	Swigcptr() uintptr;
}
func (p SwigcptrMap_Sl_int_Sc_int_Sg_) Swigcptr() uintptr {
	return uintptr(p)
}

type SwigcptrVoid uintptr
type Void interface {
	Swigcptr() uintptr;
}
func (p SwigcptrVoid) Swigcptr() uintptr {
	return uintptr(p)
}

