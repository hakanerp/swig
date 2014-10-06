/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 3.0.3
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: ./../namespace_extend.i

package namespace_extend

import "unsafe"
import _ "runtime/cgo"

var _cgo_runtime_cgocall func(unsafe.Pointer, uintptr)



type _ unsafe.Pointer



type _swig_fnptr *byte
type _swig_memberptr *byte

type SwigcptrBar uintptr

func (p SwigcptrBar) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrBar) SwigIsBar() {
}

var _wrap_new_bar unsafe.Pointer

func _swig_wrap_new_bar() (base SwigcptrBar) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_bar, _swig_p)
	return
}

func NewBar() (_swig_ret Bar) {
	return _swig_wrap_new_bar()
}

var _wrap_delete_bar unsafe.Pointer

func _swig_wrap_delete_bar(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_bar, _swig_p)
	return
}

func DeleteBar(arg1 Bar) {
	_swig_wrap_delete_bar(arg1.Swigcptr())
}

var _wrap_bar_blah unsafe.Pointer

func _swig_wrap_bar_blah(base SwigcptrBar, _ int) (_ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_bar_blah, _swig_p)
	return
}

func (arg1 SwigcptrBar) Blah(arg2 int) (_swig_ret int) {
	return _swig_wrap_bar_blah(arg1, arg2)
}

type Bar interface {
	Swigcptr() uintptr
	SwigIsBar()
	Blah(arg2 int) (_swig_ret int)
}


type SwigcptrVoid uintptr
type Void interface {
	Swigcptr() uintptr;
}
func (p SwigcptrVoid) Swigcptr() uintptr {
	return uintptr(p)
}

