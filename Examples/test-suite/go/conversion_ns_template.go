/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 3.0.3
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: ./../conversion_ns_template.i

package conversion_ns_template

import "unsafe"
import _ "runtime/cgo"

var _cgo_runtime_cgocall func(unsafe.Pointer, uintptr)



type _ unsafe.Pointer



type _swig_fnptr *byte
type _swig_memberptr *byte

type OssTest int
var _wrap_One unsafe.Pointer

func _swig_getOne() (_swig_ret OssTest) {
	_swig_p := uintptr(unsafe.Pointer(&_swig_ret))
	_cgo_runtime_cgocall(_wrap_One, _swig_p)
	return
}
var One OssTest = _swig_getOne()
var _wrap_Two unsafe.Pointer

func _swig_getTwo() (_swig_ret OssTest) {
	_swig_p := uintptr(unsafe.Pointer(&_swig_ret))
	_cgo_runtime_cgocall(_wrap_Two, _swig_p)
	return
}
var Two OssTest = _swig_getTwo()
type SwigcptrHi uintptr

func (p SwigcptrHi) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrHi) SwigIsHi() {
}

var _wrap_new_create unsafe.Pointer

func _swig_wrap_new_create(base int) (_ SwigcptrHi) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_create, _swig_p)
	return
}

func NewCreate(arg1 int) (_swig_ret Hi) {
	return _swig_wrap_new_create(arg1)
}

var _wrap_delete_Hi unsafe.Pointer

func _swig_wrap_delete_Hi(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_Hi, _swig_p)
	return
}

func DeleteHi(arg1 Hi) {
	_swig_wrap_delete_Hi(arg1.Swigcptr())
}

type Hi interface {
	Swigcptr() uintptr
	SwigIsHi()
}

type SwigcptrFoo_One uintptr

func (p SwigcptrFoo_One) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrFoo_One) SwigIsFoo_One() {
}

var _wrap_new_Foo_One unsafe.Pointer

func _swig_wrap_new_Foo_One() (base SwigcptrFoo_One) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_Foo_One, _swig_p)
	return
}

func NewFoo_One() (_swig_ret Foo_One) {
	return _swig_wrap_new_Foo_One()
}

var _wrap_delete_Foo_One unsafe.Pointer

func _swig_wrap_delete_Foo_One(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_Foo_One, _swig_p)
	return
}

func DeleteFoo_One(arg1 Foo_One) {
	_swig_wrap_delete_Foo_One(arg1.Swigcptr())
}

type Foo_One interface {
	Swigcptr() uintptr
	SwigIsFoo_One()
}

type SwigcptrBar_One uintptr

func (p SwigcptrBar_One) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrBar_One) SwigIsBar_One() {
}

var _wrap_new_Bar_create unsafe.Pointer

func _swig_wrap_new_Bar_create(base int) (_ SwigcptrBar_One) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_Bar_create, _swig_p)
	return
}

func NewBar_create(arg1 int) (_swig_ret Bar_One) {
	return _swig_wrap_new_Bar_create(arg1)
}

var _wrap_Bar_One_hello1 unsafe.Pointer

func _swig_wrap_Bar_One_hello1(base SwigcptrBar_One) (_ *int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_Bar_One_hello1, _swig_p)
	return
}

func (arg1 SwigcptrBar_One) Hello1() (_swig_ret *int) {
	return _swig_wrap_Bar_One_hello1(arg1)
}

var _wrap_Bar_One_hello2 unsafe.Pointer

func _swig_wrap_Bar_One_hello2(base SwigcptrBar_One) (_ SwigcptrFoo_One) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_Bar_One_hello2, _swig_p)
	return
}

func (arg1 SwigcptrBar_One) Hello2() (_swig_ret Foo_One) {
	return _swig_wrap_Bar_One_hello2(arg1)
}

var _wrap_delete_Bar_One unsafe.Pointer

func _swig_wrap_delete_Bar_One(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_Bar_One, _swig_p)
	return
}

func DeleteBar_One(arg1 Bar_One) {
	_swig_wrap_delete_Bar_One(arg1.Swigcptr())
}

type Bar_One interface {
	Swigcptr() uintptr
	SwigIsBar_One()
	Hello1() (_swig_ret *int)
	Hello2() (_swig_ret Foo_One)
}


type SwigcptrVoid uintptr
type Void interface {
	Swigcptr() uintptr;
}
func (p SwigcptrVoid) Swigcptr() uintptr {
	return uintptr(p)
}

