/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 3.0.3
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: ./../template_typedef_cplx5.i

package template_typedef_cplx5

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

var _wrap_A_test1 unsafe.Pointer

func _swig_wrap_A_test1(base SwigcptrA) (_ SwigcptrStd_complex_Sl_double_Sg_) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_A_test1, _swig_p)
	return
}

func (arg1 SwigcptrA) Test1() (_swig_ret Std_complex_Sl_double_Sg_) {
	return _swig_wrap_A_test1(arg1)
}

var _wrap_A_test2 unsafe.Pointer

func _swig_wrap_A_test2(base SwigcptrA) (_ SwigcptrStd_complex_Sl_double_Sg_) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_A_test2, _swig_p)
	return
}

func (arg1 SwigcptrA) Test2() (_swig_ret Std_complex_Sl_double_Sg_) {
	return _swig_wrap_A_test2(arg1)
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
	Test1() (_swig_ret Std_complex_Sl_double_Sg_)
	Test2() (_swig_ret Std_complex_Sl_double_Sg_)
}


type SwigcptrStd_complex_Sl_double_Sg_ uintptr
type Std_complex_Sl_double_Sg_ interface {
	Swigcptr() uintptr;
}
func (p SwigcptrStd_complex_Sl_double_Sg_) Swigcptr() uintptr {
	return uintptr(p)
}

type SwigcptrVoid uintptr
type Void interface {
	Swigcptr() uintptr;
}
func (p SwigcptrVoid) Swigcptr() uintptr {
	return uintptr(p)
}

