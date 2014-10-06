/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 3.0.3
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: ./../template_tbase_template.i

package template_tbase_template

import "unsafe"
import _ "runtime/cgo"

var _cgo_runtime_cgocall func(unsafe.Pointer, uintptr)



type _ unsafe.Pointer



type _swig_fnptr *byte
type _swig_memberptr *byte

type SwigcptrTraits_dd uintptr

func (p SwigcptrTraits_dd) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrTraits_dd) SwigIsTraits_dd() {
}

var _wrap_new_traits_dd unsafe.Pointer

func _swig_wrap_new_traits_dd() (base SwigcptrTraits_dd) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_traits_dd, _swig_p)
	return
}

func NewTraits_dd() (_swig_ret Traits_dd) {
	return _swig_wrap_new_traits_dd()
}

var _wrap_delete_traits_dd unsafe.Pointer

func _swig_wrap_delete_traits_dd(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_traits_dd, _swig_p)
	return
}

func DeleteTraits_dd(arg1 Traits_dd) {
	_swig_wrap_delete_traits_dd(arg1.Swigcptr())
}

type Traits_dd interface {
	Swigcptr() uintptr
	SwigIsTraits_dd()
}

type SwigcptrFunktion_dd uintptr

func (p SwigcptrFunktion_dd) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrFunktion_dd) SwigIsFunktion_dd() {
}

var _wrap_Funktion_dd_test unsafe.Pointer

func _swig_wrap_Funktion_dd_test(base SwigcptrFunktion_dd) (_ string) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_Funktion_dd_test, _swig_p)
	return
}

func (arg1 SwigcptrFunktion_dd) Test() (_swig_ret string) {
	return _swig_wrap_Funktion_dd_test(arg1)
}

var _wrap_new_Funktion_dd unsafe.Pointer

func _swig_wrap_new_Funktion_dd() (base SwigcptrFunktion_dd) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_Funktion_dd, _swig_p)
	return
}

func NewFunktion_dd() (_swig_ret Funktion_dd) {
	return _swig_wrap_new_Funktion_dd()
}

var _wrap_delete_Funktion_dd unsafe.Pointer

func _swig_wrap_delete_Funktion_dd(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_Funktion_dd, _swig_p)
	return
}

func DeleteFunktion_dd(arg1 Funktion_dd) {
	_swig_wrap_delete_Funktion_dd(arg1.Swigcptr())
}

type Funktion_dd interface {
	Swigcptr() uintptr
	SwigIsFunktion_dd()
	Test() (_swig_ret string)
}

type SwigcptrClass_dd uintptr

func (p SwigcptrClass_dd) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrClass_dd) SwigIsClass_dd() {
}

var _wrap_new_Class_dd unsafe.Pointer

func _swig_wrap_new_Class_dd() (base SwigcptrClass_dd) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_Class_dd, _swig_p)
	return
}

func NewClass_dd() (_swig_ret Class_dd) {
	return _swig_wrap_new_Class_dd()
}

var _wrap_delete_Class_dd unsafe.Pointer

func _swig_wrap_delete_Class_dd(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_Class_dd, _swig_p)
	return
}

func DeleteClass_dd(arg1 Class_dd) {
	_swig_wrap_delete_Class_dd(arg1.Swigcptr())
}

var _wrap_Class_dd_test unsafe.Pointer

func _swig_wrap_Class_dd_test(base SwigcptrClass_dd) (_ string) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_Class_dd_test, _swig_p)
	return
}

func (_swig_base SwigcptrClass_dd) Test() (_swig_ret string) {
	return _swig_wrap_Class_dd_test(_swig_base)
}

func (p SwigcptrClass_dd) SwigIsFunktion_dd() {
}

func (p SwigcptrClass_dd) SwigGetFunktion_dd() Funktion_dd {
	return SwigcptrFunktion_dd(p.Swigcptr())
}

type Class_dd interface {
	Swigcptr() uintptr
	SwigIsClass_dd()
	Test() (_swig_ret string)
	SwigIsFunktion_dd()
	SwigGetFunktion_dd() Funktion_dd
}

var _wrap_make_Class_dd unsafe.Pointer

func _swig_wrap_make_Class_dd() (base SwigcptrFunktion_dd) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_make_Class_dd, _swig_p)
	return
}

func Make_Class_dd() (_swig_ret Funktion_dd) {
	return _swig_wrap_make_Class_dd()
}


type SwigcptrVoid uintptr
type Void interface {
	Swigcptr() uintptr;
}
func (p SwigcptrVoid) Swigcptr() uintptr {
	return uintptr(p)
}

