/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 3.0.3
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: ./../voidtest.i

package voidtest

import "unsafe"
import _ "runtime/cgo"

var _cgo_runtime_cgocall func(unsafe.Pointer, uintptr)



type _ unsafe.Pointer



type _swig_fnptr *byte
type _swig_memberptr *byte

var _wrap_globalfunc unsafe.Pointer

func _swig_wrap_globalfunc() {
	var _swig_p uintptr
	_cgo_runtime_cgocall(_wrap_globalfunc, _swig_p)
	return
}

func Globalfunc() {
	_swig_wrap_globalfunc()
}

type SwigcptrFoo uintptr

func (p SwigcptrFoo) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrFoo) SwigIsFoo() {
}

var _wrap_new_Foo unsafe.Pointer

func _swig_wrap_new_Foo() (base SwigcptrFoo) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_Foo, _swig_p)
	return
}

func NewFoo() (_swig_ret Foo) {
	return _swig_wrap_new_Foo()
}

var _wrap_Foo_memberfunc unsafe.Pointer

func _swig_wrap_Foo_memberfunc(base SwigcptrFoo) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_Foo_memberfunc, _swig_p)
	return
}

func (arg1 SwigcptrFoo) Memberfunc() {
	_swig_wrap_Foo_memberfunc(arg1)
}

var _wrap_Foo_staticmemberfunc unsafe.Pointer

func _swig_wrap_Foo_staticmemberfunc() {
	var _swig_p uintptr
	_cgo_runtime_cgocall(_wrap_Foo_staticmemberfunc, _swig_p)
	return
}

func FooStaticmemberfunc() {
	_swig_wrap_Foo_staticmemberfunc()
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
	Memberfunc()
}

var _wrap_vfunc1 unsafe.Pointer

func Vfunc1(arg1 uintptr) (_swig_ret uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&arg1))
	_cgo_runtime_cgocall(_wrap_vfunc1, _swig_p)
	return
}
var _wrap_vfunc2 unsafe.Pointer

func _swig_wrap_vfunc2(base uintptr) (_ uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_vfunc2, _swig_p)
	return
}

func Vfunc2(arg1 Foo) (_swig_ret uintptr) {
	return _swig_wrap_vfunc2(arg1.Swigcptr())
}

var _wrap_vfunc3 unsafe.Pointer

func _swig_wrap_vfunc3(base uintptr) (_ SwigcptrFoo) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_vfunc3, _swig_p)
	return
}

func Vfunc3(arg1 uintptr) (_swig_ret Foo) {
	return _swig_wrap_vfunc3(arg1)
}

var _wrap_vfunc4 unsafe.Pointer

func _swig_wrap_vfunc4(base uintptr) (_ SwigcptrFoo) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_vfunc4, _swig_p)
	return
}

func Vfunc4(arg1 Foo) (_swig_ret Foo) {
	return _swig_wrap_vfunc4(arg1.Swigcptr())
}


type SwigcptrVoid uintptr
type Void interface {
	Swigcptr() uintptr;
}
func (p SwigcptrVoid) Swigcptr() uintptr {
	return uintptr(p)
}

