/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 3.0.3
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: ./../template_base_template.i

package template_base_template

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
}

type SwigcptrKlass_dd uintptr

func (p SwigcptrKlass_dd) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrKlass_dd) SwigIsKlass_dd() {
}

var _wrap_new_Klass_dd unsafe.Pointer

func _swig_wrap_new_Klass_dd() (base SwigcptrKlass_dd) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_Klass_dd, _swig_p)
	return
}

func NewKlass_dd() (_swig_ret Klass_dd) {
	return _swig_wrap_new_Klass_dd()
}

var _wrap_delete_Klass_dd unsafe.Pointer

func _swig_wrap_delete_Klass_dd(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_Klass_dd, _swig_p)
	return
}

func DeleteKlass_dd(arg1 Klass_dd) {
	_swig_wrap_delete_Klass_dd(arg1.Swigcptr())
}

func (p SwigcptrKlass_dd) SwigIsFunktion_dd() {
}

func (p SwigcptrKlass_dd) SwigGetFunktion_dd() Funktion_dd {
	return SwigcptrFunktion_dd(p.Swigcptr())
}

type Klass_dd interface {
	Swigcptr() uintptr
	SwigIsKlass_dd()
	SwigIsFunktion_dd()
	SwigGetFunktion_dd() Funktion_dd
}


type SwigcptrVoid uintptr
type Void interface {
	Swigcptr() uintptr;
}
func (p SwigcptrVoid) Swigcptr() uintptr {
	return uintptr(p)
}

