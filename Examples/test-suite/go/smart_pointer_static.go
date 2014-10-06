/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 3.0.3
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: ./../smart_pointer_static.i

package smart_pointer_static

import "unsafe"
import _ "runtime/cgo"

var _cgo_runtime_cgocall func(unsafe.Pointer, uintptr)



type _ unsafe.Pointer



type _swig_fnptr *byte
type _swig_memberptr *byte

type SwigcptrMyHandle_Foo2 uintptr

func (p SwigcptrMyHandle_Foo2) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrMyHandle_Foo2) SwigIsMyHandle_Foo2() {
}

var _wrap_MyHandle_Foo2___deref__ unsafe.Pointer

func _swig_wrap_MyHandle_Foo2___deref__(base SwigcptrMyHandle_Foo2) (_ SwigcptrFoo2) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_MyHandle_Foo2___deref__, _swig_p)
	return
}

func (arg1 SwigcptrMyHandle_Foo2) X__deref__() (_swig_ret Foo2) {
	return _swig_wrap_MyHandle_Foo2___deref__(arg1)
}

var _wrap_new_MyHandle_Foo2 unsafe.Pointer

func _swig_wrap_new_MyHandle_Foo2() (base SwigcptrMyHandle_Foo2) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_MyHandle_Foo2, _swig_p)
	return
}

func NewMyHandle_Foo2() (_swig_ret MyHandle_Foo2) {
	return _swig_wrap_new_MyHandle_Foo2()
}

var _wrap_delete_MyHandle_Foo2 unsafe.Pointer

func _swig_wrap_delete_MyHandle_Foo2(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_MyHandle_Foo2, _swig_p)
	return
}

func DeleteMyHandle_Foo2(arg1 MyHandle_Foo2) {
	_swig_wrap_delete_MyHandle_Foo2(arg1.Swigcptr())
}

var _wrap_MyHandle_Foo2_sum__SWIG_0 unsafe.Pointer

func _swig_wrap_MyHandle_Foo2_sum__SWIG_0(base SwigcptrMyHandle_Foo2, _ int, _ int) (_ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_MyHandle_Foo2_sum__SWIG_0, _swig_p)
	return
}

func (arg1 SwigcptrMyHandle_Foo2) Sum__SWIG_0(arg2 int, arg3 int) (_swig_ret int) {
	return _swig_wrap_MyHandle_Foo2_sum__SWIG_0(arg1, arg2, arg3)
}

var _wrap_MyHandle_Foo2_sum__SWIG_1 unsafe.Pointer

func _swig_wrap_MyHandle_Foo2_sum__SWIG_1(base SwigcptrMyHandle_Foo2, _ int, _ int, _ int) (_ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_MyHandle_Foo2_sum__SWIG_1, _swig_p)
	return
}

func (arg1 SwigcptrMyHandle_Foo2) Sum__SWIG_1(arg2 int, arg3 int, arg4 int) (_swig_ret int) {
	return _swig_wrap_MyHandle_Foo2_sum__SWIG_1(arg1, arg2, arg3, arg4)
}

func (p SwigcptrMyHandle_Foo2) Sum(a ...interface{}) int {
	argc := len(a)
	if argc == 2 {
		return p.Sum__SWIG_0(a[0].(int), a[1].(int))
	}
	if argc == 3 {
		return p.Sum__SWIG_1(a[0].(int), a[1].(int), a[2].(int))
	}
	panic("No match for overloaded function call")
}

type MyHandle_Foo2 interface {
	Swigcptr() uintptr
	SwigIsMyHandle_Foo2()
	X__deref__() (_swig_ret Foo2)
	Sum(a ...interface{}) int
}

type SwigcptrFoo2 uintptr

func (p SwigcptrFoo2) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrFoo2) SwigIsFoo2() {
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

var _wrap_Foo2_sum__SWIG_0 unsafe.Pointer

func _swig_wrap_Foo2_sum__SWIG_0(base SwigcptrFoo2, _ int, _ int) (_ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_Foo2_sum__SWIG_0, _swig_p)
	return
}

func (arg1 SwigcptrFoo2) Sum__SWIG_0(arg2 int, arg3 int) (_swig_ret int) {
	return _swig_wrap_Foo2_sum__SWIG_0(arg1, arg2, arg3)
}

var _wrap_Foo2_sum__SWIG_1 unsafe.Pointer

func Foo2Sum__SWIG_1(arg1 int, arg2 int, arg3 int) (_swig_ret int) {
	_swig_p := uintptr(unsafe.Pointer(&arg1))
	_cgo_runtime_cgocall(_wrap_Foo2_sum__SWIG_1, _swig_p)
	return
}
func Foo2Sum(a ...interface{}) int {
	argc := len(a)
	if argc == 3 {
		if _, ok := a[0].(Foo2); !ok {
			goto check_1
		}
		return a[0].(SwigcptrFoo2).Sum__SWIG_0(a[1].(int), a[2].(int))
	}
check_1:
	if argc == 3 {
		return Foo2Sum__SWIG_1(a[0].(int), a[1].(int), a[2].(int))
	}
	panic("No match for overloaded function call")
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

type Foo2 interface {
	Swigcptr() uintptr
	SwigIsFoo2()
}


type SwigcptrVoid uintptr
type Void interface {
	Swigcptr() uintptr;
}
func (p SwigcptrVoid) Swigcptr() uintptr {
	return uintptr(p)
}

