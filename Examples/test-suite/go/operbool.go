/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 3.0.3
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: ./../operbool.i

package operbool

import "unsafe"
import _ "runtime/cgo"

var _cgo_runtime_cgocall func(unsafe.Pointer, uintptr)



type _ unsafe.Pointer



type _swig_fnptr *byte
type _swig_memberptr *byte

type SwigcptrTest uintptr

func (p SwigcptrTest) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrTest) SwigIsTest() {
}

var _wrap_Test_operator_bool unsafe.Pointer

func _swig_wrap_Test_operator_bool(base SwigcptrTest) (_ bool) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_Test_operator_bool, _swig_p)
	return
}

func (arg1 SwigcptrTest) Operator_bool() (_swig_ret bool) {
	return _swig_wrap_Test_operator_bool(arg1)
}

var _wrap_new_Test unsafe.Pointer

func _swig_wrap_new_Test() (base SwigcptrTest) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_Test, _swig_p)
	return
}

func NewTest() (_swig_ret Test) {
	return _swig_wrap_new_Test()
}

var _wrap_delete_Test unsafe.Pointer

func _swig_wrap_delete_Test(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_Test, _swig_p)
	return
}

func DeleteTest(arg1 Test) {
	_swig_wrap_delete_Test(arg1.Swigcptr())
}

type Test interface {
	Swigcptr() uintptr
	SwigIsTest()
	Operator_bool() (_swig_ret bool)
}


type SwigcptrVoid uintptr
type Void interface {
	Swigcptr() uintptr;
}
func (p SwigcptrVoid) Swigcptr() uintptr {
	return uintptr(p)
}

