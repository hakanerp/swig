/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 3.0.3
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: ./../global_scope_types.i

package global_scope_types

import "unsafe"
import _ "runtime/cgo"

var _cgo_runtime_cgocall func(unsafe.Pointer, uintptr)



type _ unsafe.Pointer



type _swig_fnptr *byte
type _swig_memberptr *byte

type SwigcptrDingaling uintptr

func (p SwigcptrDingaling) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrDingaling) SwigIsDingaling() {
}

type Dingaling interface {
	Swigcptr() uintptr
	SwigIsDingaling()
}

type SwigcptrTing uintptr

func (p SwigcptrTing) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrTing) SwigIsTing() {
}

var _wrap_new_Ting unsafe.Pointer

func _swig_wrap_new_Ting() (base SwigcptrTing) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_Ting, _swig_p)
	return
}

func NewTing() (_swig_ret Ting) {
	return _swig_wrap_new_Ting()
}

var _wrap_delete_Ting unsafe.Pointer

func _swig_wrap_delete_Ting(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_Ting, _swig_p)
	return
}

func DeleteTing(arg1 Ting) {
	_swig_wrap_delete_Ting(arg1.Swigcptr())
}

type Ting interface {
	Swigcptr() uintptr
	SwigIsTing()
}

type SwigcptrTest uintptr

func (p SwigcptrTest) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrTest) SwigIsTest() {
}

var _wrap_Test_something unsafe.Pointer

func _swig_wrap_Test_something(base SwigcptrTest, _ uintptr, _ uintptr, _ uintptr, _ uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_Test_something, _swig_p)
	return
}

func (arg1 SwigcptrTest) Something(arg2 Dingaling, arg3 Dingaling, arg4 Dingaling, arg5 Dingaling) {
	_swig_wrap_Test_something(arg1, arg2.Swigcptr(), arg3.Swigcptr(), arg4.Swigcptr(), arg5.Swigcptr())
}

var _wrap_Test_tsomething unsafe.Pointer

func _swig_wrap_Test_tsomething(base SwigcptrTest, _ uintptr, _ uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_Test_tsomething, _swig_p)
	return
}

func (arg1 SwigcptrTest) Tsomething(arg2 MyTemplate_Sl__Dingaling_Sg_, arg3 MyTemplate_Sl_Dingaling_SS_const_Sm__Sg_) {
	_swig_wrap_Test_tsomething(arg1, arg2.Swigcptr(), arg3.Swigcptr())
}

var _wrap_Test_nothing unsafe.Pointer

func _swig_wrap_Test_nothing(base SwigcptrTest, _ uintptr, _ uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_Test_nothing, _swig_p)
	return
}

func (arg1 SwigcptrTest) Nothing(arg2 Ting, arg3 Ting) {
	_swig_wrap_Test_nothing(arg1, arg2.Swigcptr(), arg3.Swigcptr())
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
	Something(arg2 Dingaling, arg3 Dingaling, arg4 Dingaling, arg5 Dingaling)
	Tsomething(arg2 MyTemplate_Sl__Dingaling_Sg_, arg3 MyTemplate_Sl_Dingaling_SS_const_Sm__Sg_)
	Nothing(arg2 Ting, arg3 Ting)
}

var _wrap_funcptrtest unsafe.Pointer

func _swig_wrap_funcptrtest(base _swig_fnptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_funcptrtest, _swig_p)
	return
}

func Funcptrtest(arg1 _swig_fnptr) {
	_swig_wrap_funcptrtest(arg1)
}


type SwigcptrVoid uintptr
type Void interface {
	Swigcptr() uintptr;
}
func (p SwigcptrVoid) Swigcptr() uintptr {
	return uintptr(p)
}

type SwigcptrMyTemplate_Sl__Dingaling_Sg_ uintptr
type MyTemplate_Sl__Dingaling_Sg_ interface {
	Swigcptr() uintptr;
}
func (p SwigcptrMyTemplate_Sl__Dingaling_Sg_) Swigcptr() uintptr {
	return uintptr(p)
}

type SwigcptrMyTemplate_Sl_Dingaling_SS_const_Sm__Sg_ uintptr
type MyTemplate_Sl_Dingaling_SS_const_Sm__Sg_ interface {
	Swigcptr() uintptr;
}
func (p SwigcptrMyTemplate_Sl_Dingaling_SS_const_Sm__Sg_) Swigcptr() uintptr {
	return uintptr(p)
}

