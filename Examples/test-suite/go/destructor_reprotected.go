/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 3.0.3
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: ./../destructor_reprotected.i

package destructor_reprotected

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
}

type SwigcptrB uintptr

func (p SwigcptrB) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrB) SwigIsB() {
}

func (p SwigcptrB) SwigIsA() {
}

func (p SwigcptrB) SwigGetA() A {
	return SwigcptrA(p.Swigcptr())
}

type B interface {
	Swigcptr() uintptr
	SwigIsB()
	SwigIsA()
	SwigGetA() A
}

type SwigcptrC uintptr

func (p SwigcptrC) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrC) SwigIsC() {
}

var _wrap_new_C unsafe.Pointer

func _swig_wrap_new_C() (base SwigcptrC) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_C, _swig_p)
	return
}

func NewC() (_swig_ret C) {
	return _swig_wrap_new_C()
}

var _wrap_delete_C unsafe.Pointer

func _swig_wrap_delete_C(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_C, _swig_p)
	return
}

func DeleteC(arg1 C) {
	_swig_wrap_delete_C(arg1.Swigcptr())
}

func (p SwigcptrC) SwigIsB() {
}

func (p SwigcptrC) SwigGetB() B {
	return SwigcptrB(p.Swigcptr())
}

func (p SwigcptrC) SwigIsA() {
}

func (p SwigcptrC) SwigGetA() A {
	return SwigcptrA(p.Swigcptr())
}

type C interface {
	Swigcptr() uintptr
	SwigIsC()
	SwigIsB()
	SwigGetB() B
	SwigIsA()
	SwigGetA() A
}


type SwigcptrVoid uintptr
type Void interface {
	Swigcptr() uintptr;
}
func (p SwigcptrVoid) Swigcptr() uintptr {
	return uintptr(p)
}

