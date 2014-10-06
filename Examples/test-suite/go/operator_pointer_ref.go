/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 3.0.3
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: ./../operator_pointer_ref.i

package operator_pointer_ref

import "unsafe"
import _ "runtime/cgo"

var _cgo_runtime_cgocall func(unsafe.Pointer, uintptr)



type _ unsafe.Pointer



type _swig_fnptr *byte
type _swig_memberptr *byte

type SwigcptrMyClass uintptr

func (p SwigcptrMyClass) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrMyClass) SwigIsMyClass() {
}

var _wrap_new_MyClass__SWIG_0 unsafe.Pointer

func _swig_wrap_new_MyClass__SWIG_0(base string) (_ SwigcptrMyClass) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_MyClass__SWIG_0, _swig_p)
	return
}

func NewMyClass__SWIG_0(arg1 string) (_swig_ret MyClass) {
	return _swig_wrap_new_MyClass__SWIG_0(arg1)
}

var _wrap_new_MyClass__SWIG_1 unsafe.Pointer

func _swig_wrap_new_MyClass__SWIG_1() (base SwigcptrMyClass) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_MyClass__SWIG_1, _swig_p)
	return
}

func NewMyClass__SWIG_1() (_swig_ret MyClass) {
	return _swig_wrap_new_MyClass__SWIG_1()
}

func NewMyClass(a ...interface{}) MyClass {
	argc := len(a)
	if argc == 0 {
		return NewMyClass__SWIG_1()
	}
	if argc == 1 {
		return NewMyClass__SWIG_0(a[0].(string))
	}
	panic("No match for overloaded function call")
}

var _wrap_delete_MyClass unsafe.Pointer

func _swig_wrap_delete_MyClass(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_MyClass, _swig_p)
	return
}

func DeleteMyClass(arg1 MyClass) {
	_swig_wrap_delete_MyClass(arg1.Swigcptr())
}

var _wrap_MyClass_AsCharStarRef unsafe.Pointer

func _swig_wrap_MyClass_AsCharStarRef(base SwigcptrMyClass) (_ string) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_MyClass_AsCharStarRef, _swig_p)
	return
}

func (arg1 SwigcptrMyClass) AsCharStarRef() (_swig_ret string) {
	return _swig_wrap_MyClass_AsCharStarRef(arg1)
}

type MyClass interface {
	Swigcptr() uintptr
	SwigIsMyClass()
	AsCharStarRef() (_swig_ret string)
}


type SwigcptrVoid uintptr
type Void interface {
	Swigcptr() uintptr;
}
func (p SwigcptrVoid) Swigcptr() uintptr {
	return uintptr(p)
}

