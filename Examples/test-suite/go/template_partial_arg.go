/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 3.0.3
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: ./../template_partial_arg.i

package template_partial_arg

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

var _wrap_new_Bar unsafe.Pointer

func _swig_wrap_new_Bar() (base SwigcptrBar) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_Bar, _swig_p)
	return
}

func NewBar() (_swig_ret Bar) {
	return _swig_wrap_new_Bar()
}

var _wrap_delete_Bar unsafe.Pointer

func _swig_wrap_delete_Bar(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_Bar, _swig_p)
	return
}

func DeleteBar(arg1 Bar) {
	_swig_wrap_delete_Bar(arg1.Swigcptr())
}

type Bar interface {
	Swigcptr() uintptr
	SwigIsBar()
}

type SwigcptrFoo1 uintptr

func (p SwigcptrFoo1) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrFoo1) SwigIsFoo1() {
}

var _wrap_Foo1_bar unsafe.Pointer

func _swig_wrap_Foo1_bar(base SwigcptrFoo1) (_ SwigcptrBar) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_Foo1_bar, _swig_p)
	return
}

func (arg1 SwigcptrFoo1) Bar() (_swig_ret Bar) {
	return _swig_wrap_Foo1_bar(arg1)
}

var _wrap_Foo1_baz unsafe.Pointer

func _swig_wrap_Foo1_baz(base SwigcptrFoo1) (_ SwigcptrBar) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_Foo1_baz, _swig_p)
	return
}

func (arg1 SwigcptrFoo1) Baz() (_swig_ret Bar) {
	return _swig_wrap_Foo1_baz(arg1)
}

var _wrap_new_Foo1 unsafe.Pointer

func _swig_wrap_new_Foo1() (base SwigcptrFoo1) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_Foo1, _swig_p)
	return
}

func NewFoo1() (_swig_ret Foo1) {
	return _swig_wrap_new_Foo1()
}

var _wrap_delete_Foo1 unsafe.Pointer

func _swig_wrap_delete_Foo1(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_Foo1, _swig_p)
	return
}

func DeleteFoo1(arg1 Foo1) {
	_swig_wrap_delete_Foo1(arg1.Swigcptr())
}

type Foo1 interface {
	Swigcptr() uintptr
	SwigIsFoo1()
	Bar() (_swig_ret Bar)
	Baz() (_swig_ret Bar)
}

type SwigcptrFoo2 uintptr

func (p SwigcptrFoo2) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrFoo2) SwigIsFoo2() {
}

var _wrap_Foo2_bar unsafe.Pointer

func _swig_wrap_Foo2_bar(base SwigcptrFoo2) (_ SwigcptrBar) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_Foo2_bar, _swig_p)
	return
}

func (arg1 SwigcptrFoo2) Bar() (_swig_ret Bar) {
	return _swig_wrap_Foo2_bar(arg1)
}

var _wrap_Foo2_baz unsafe.Pointer

func _swig_wrap_Foo2_baz(base SwigcptrFoo2) (_ SwigcptrBar) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_Foo2_baz, _swig_p)
	return
}

func (arg1 SwigcptrFoo2) Baz() (_swig_ret Bar) {
	return _swig_wrap_Foo2_baz(arg1)
}

var _wrap_new_Foo2 unsafe.Pointer

func _swig_wrap_new_Foo2() (base SwigcptrFoo2) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_Foo2, _swig_p)
	return
}

func NewFoo2() (_swig_ret Foo2) {
	return _swig_wrap_new_Foo2()
}

var _wrap_delete_Foo2 unsafe.Pointer

func _swig_wrap_delete_Foo2(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_Foo2, _swig_p)
	return
}

func DeleteFoo2(arg1 Foo2) {
	_swig_wrap_delete_Foo2(arg1.Swigcptr())
}

type Foo2 interface {
	Swigcptr() uintptr
	SwigIsFoo2()
	Bar() (_swig_ret Bar)
	Baz() (_swig_ret Bar)
}


type SwigcptrVoid uintptr
type Void interface {
	Swigcptr() uintptr;
}
func (p SwigcptrVoid) Swigcptr() uintptr {
	return uintptr(p)
}

