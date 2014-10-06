/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 3.0.3
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: ./../catches.i

package catches

import "unsafe"
import _ "runtime/cgo"

var _cgo_runtime_cgocall func(unsafe.Pointer, uintptr)



type _ unsafe.Pointer



type _swig_fnptr *byte
type _swig_memberptr *byte

type SwigcptrThreeException uintptr

func (p SwigcptrThreeException) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrThreeException) SwigIsThreeException() {
}

var _wrap_new_ThreeException unsafe.Pointer

func _swig_wrap_new_ThreeException() (base SwigcptrThreeException) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_ThreeException, _swig_p)
	return
}

func NewThreeException() (_swig_ret ThreeException) {
	return _swig_wrap_new_ThreeException()
}

var _wrap_delete_ThreeException unsafe.Pointer

func _swig_wrap_delete_ThreeException(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_ThreeException, _swig_p)
	return
}

func DeleteThreeException(arg1 ThreeException) {
	_swig_wrap_delete_ThreeException(arg1.Swigcptr())
}

type ThreeException interface {
	Swigcptr() uintptr
	SwigIsThreeException()
}

var _wrap_test_catches unsafe.Pointer

func _swig_wrap_test_catches(base int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_test_catches, _swig_p)
	return
}

func Test_catches(arg1 int) {
	_swig_wrap_test_catches(arg1)
}

var _wrap_test_exception_specification unsafe.Pointer

func _swig_wrap_test_exception_specification(base int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_test_exception_specification, _swig_p)
	return
}

func Test_exception_specification(arg1 int) {
	_swig_wrap_test_exception_specification(arg1)
}

var _wrap_test_catches_all unsafe.Pointer

func _swig_wrap_test_catches_all(base int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_test_catches_all, _swig_p)
	return
}

func Test_catches_all(arg1 int) {
	_swig_wrap_test_catches_all(arg1)
}


type SwigcptrVoid uintptr
type Void interface {
	Swigcptr() uintptr;
}
func (p SwigcptrVoid) Swigcptr() uintptr {
	return uintptr(p)
}

