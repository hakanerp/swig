/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 3.0.3
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: ./../constructor_explicit.i

package constructor_explicit

import "unsafe"
import _ "runtime/cgo"

var _cgo_runtime_cgocall func(unsafe.Pointer, uintptr)



type _ unsafe.Pointer



type _swig_fnptr *byte
type _swig_memberptr *byte

type SwigcptrFoo uintptr

func (p SwigcptrFoo) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrFoo) SwigIsFoo() {
}

var _wrap_new_Foo__SWIG_0 unsafe.Pointer

func _swig_wrap_new_Foo__SWIG_0() (base SwigcptrFoo) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_Foo__SWIG_0, _swig_p)
	return
}

func NewFoo__SWIG_0() (_swig_ret Foo) {
	return _swig_wrap_new_Foo__SWIG_0()
}

var _wrap_new_Foo__SWIG_1 unsafe.Pointer

func _swig_wrap_new_Foo__SWIG_1(base int) (_ SwigcptrFoo) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_Foo__SWIG_1, _swig_p)
	return
}

func NewFoo__SWIG_1(arg1 int) (_swig_ret Foo) {
	return _swig_wrap_new_Foo__SWIG_1(arg1)
}

func NewFoo(a ...interface{}) Foo {
	argc := len(a)
	if argc == 0 {
		return NewFoo__SWIG_0()
	}
	if argc == 1 {
		return NewFoo__SWIG_1(a[0].(int))
	}
	panic("No match for overloaded function call")
}

var _wrap_delete_Foo unsafe.Pointer

func _swig_wrap_delete_Foo(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_Foo, _swig_p)
	return
}

func DeleteFoo(arg1 Foo) {
	_swig_wrap_delete_Foo(arg1.Swigcptr())
}

type Foo interface {
	Swigcptr() uintptr
	SwigIsFoo()
}

var _wrap_test unsafe.Pointer

func _swig_wrap_test(base uintptr) (_ SwigcptrFoo) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_test, _swig_p)
	return
}

func Test(arg1 Foo) (_swig_ret Foo) {
	return _swig_wrap_test(arg1.Swigcptr())
}


type SwigcptrVoid uintptr
type Void interface {
	Swigcptr() uintptr;
}
func (p SwigcptrVoid) Swigcptr() uintptr {
	return uintptr(p)
}

