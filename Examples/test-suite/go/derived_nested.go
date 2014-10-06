/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 3.0.3
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: ./../derived_nested.i

package derived_nested

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

var _wrap_A_x_set unsafe.Pointer

func _swig_wrap_A_x_set(base SwigcptrA, _ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_A_x_set, _swig_p)
	return
}

func (arg1 SwigcptrA) SetX(arg2 int) {
	_swig_wrap_A_x_set(arg1, arg2)
}

var _wrap_A_x_get unsafe.Pointer

func _swig_wrap_A_x_get(base SwigcptrA) (_ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_A_x_get, _swig_p)
	return
}

func (arg1 SwigcptrA) GetX() (_swig_ret int) {
	return _swig_wrap_A_x_get(arg1)
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
	SetX(arg2 int)
	GetX() (_swig_ret int)
}

type SwigcptrB uintptr

func (p SwigcptrB) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrB) SwigIsB() {
}

var _wrap_new_B unsafe.Pointer

func _swig_wrap_new_B() (base SwigcptrB) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_B, _swig_p)
	return
}

func NewB() (_swig_ret B) {
	return _swig_wrap_new_B()
}

var _wrap_delete_B unsafe.Pointer

func _swig_wrap_delete_B(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_B, _swig_p)
	return
}

func DeleteB(arg1 B) {
	_swig_wrap_delete_B(arg1.Swigcptr())
}

type B interface {
	Swigcptr() uintptr
	SwigIsB()
}

type SwigcptrBB uintptr

func (p SwigcptrBB) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrBB) SwigIsBB() {
}

var _wrap_BB_ff_instance_set unsafe.Pointer

func _swig_wrap_BB_ff_instance_set(base SwigcptrBB, _ uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_BB_ff_instance_set, _swig_p)
	return
}

func (arg1 SwigcptrBB) SetFf_instance(arg2 BB_FF) {
	_swig_wrap_BB_ff_instance_set(arg1, arg2.Swigcptr())
}

var _wrap_BB_ff_instance_get unsafe.Pointer

func _swig_wrap_BB_ff_instance_get(base SwigcptrBB) (_ SwigcptrBB_FF) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_BB_ff_instance_get, _swig_p)
	return
}

func (arg1 SwigcptrBB) GetFf_instance() (_swig_ret BB_FF) {
	return _swig_wrap_BB_ff_instance_get(arg1)
}

var _wrap_BB_useEE unsafe.Pointer

func _swig_wrap_BB_useEE(base SwigcptrBB, _ uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_BB_useEE, _swig_p)
	return
}

func (arg1 SwigcptrBB) UseEE(arg2 BB_EE) {
	_swig_wrap_BB_useEE(arg1, arg2.Swigcptr())
}

var _wrap_new_BB unsafe.Pointer

func _swig_wrap_new_BB() (base SwigcptrBB) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_BB, _swig_p)
	return
}

func NewBB() (_swig_ret BB) {
	return _swig_wrap_new_BB()
}

var _wrap_delete_BB unsafe.Pointer

func _swig_wrap_delete_BB(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_BB, _swig_p)
	return
}

func DeleteBB(arg1 BB) {
	_swig_wrap_delete_BB(arg1.Swigcptr())
}

type BB interface {
	Swigcptr() uintptr
	SwigIsBB()
	SetFf_instance(arg2 BB_FF)
	GetFf_instance() (_swig_ret BB_FF)
	UseEE(arg2 BB_EE)
}


type SwigcptrBB_FF uintptr
type BB_FF interface {
	Swigcptr() uintptr;
}
func (p SwigcptrBB_FF) Swigcptr() uintptr {
	return uintptr(p)
}

type SwigcptrBB_EE uintptr
type BB_EE interface {
	Swigcptr() uintptr;
}
func (p SwigcptrBB_EE) Swigcptr() uintptr {
	return uintptr(p)
}

type SwigcptrVoid uintptr
type Void interface {
	Swigcptr() uintptr;
}
func (p SwigcptrVoid) Swigcptr() uintptr {
	return uintptr(p)
}

