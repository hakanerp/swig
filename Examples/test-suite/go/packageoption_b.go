/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 3.0.3
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: ./../packageoption_b.i

package packageoption_b

import "unsafe"
import _ "runtime/cgo"

var _cgo_runtime_cgocall func(unsafe.Pointer, uintptr)



type _ unsafe.Pointer



type _swig_fnptr *byte
type _swig_memberptr *byte

type SwigcptrB uintptr

func (p SwigcptrB) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrB) SwigIsB() {
}

var _wrap_B_testInt unsafe.Pointer

func _swig_wrap_B_testInt(base SwigcptrB) (_ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_B_testInt, _swig_p)
	return
}

func (arg1 SwigcptrB) TestInt() (_swig_ret int) {
	return _swig_wrap_B_testInt(arg1)
}

var _wrap_new_B unsafe.Pointer

func _swig_wrap_new_B() (base SwigcptrB) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_B, _swig_p)
	return
}

func NewB() (_swig_ret B) {
	return _swig_wrap_new_B()
}

var _wrap_delete_B unsafe.Pointer

func _swig_wrap_delete_B(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_B, _swig_p)
	return
}

func DeleteB(arg1 B) {
	_swig_wrap_delete_B(arg1.Swigcptr())
}

type B interface {
	Swigcptr() uintptr
	SwigIsB()
	TestInt() (_swig_ret int)
}


type SwigcptrVoid uintptr
type Void interface {
	Swigcptr() uintptr;
}
func (p SwigcptrVoid) Swigcptr() uintptr {
	return uintptr(p)
}

