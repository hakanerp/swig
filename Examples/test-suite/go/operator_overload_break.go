/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 3.0.3
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: ./../operator_overload_break.i

package operator_overload_break

import "unsafe"
import _ "runtime/cgo"

var _cgo_runtime_cgocall func(unsafe.Pointer, uintptr)



type _ unsafe.Pointer



type _swig_fnptr *byte
type _swig_memberptr *byte

type SwigcptrOp uintptr

func (p SwigcptrOp) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrOp) SwigIsOp() {
}

var _wrap_new_Op__SWIG_0 unsafe.Pointer

func _swig_wrap_new_Op__SWIG_0(base int) (_ SwigcptrOp) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_Op__SWIG_0, _swig_p)
	return
}

func NewOp__SWIG_0(arg1 int) (_swig_ret Op) {
	return _swig_wrap_new_Op__SWIG_0(arg1)
}

var _wrap_new_Op__SWIG_1 unsafe.Pointer

func _swig_wrap_new_Op__SWIG_1(base uintptr) (_ SwigcptrOp) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_Op__SWIG_1, _swig_p)
	return
}

func NewOp__SWIG_1(arg1 Op) (_swig_ret Op) {
	return _swig_wrap_new_Op__SWIG_1(arg1.Swigcptr())
}

func NewOp(a ...interface{}) Op {
	argc := len(a)
	if argc == 1 {
		if _, ok := a[0].(Op); !ok {
			goto check_1
		}
		return NewOp__SWIG_1(a[0].(Op))
	}
check_1:
	if argc == 1 {
		return NewOp__SWIG_0(a[0].(int))
	}
	panic("No match for overloaded function call")
}

var _wrap_Op_EqualEqual__SWIG_0 unsafe.Pointer

func _swig_wrap_Op_EqualEqual__SWIG_0(base SwigcptrOp, _ uintptr) (_ bool) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_Op_EqualEqual__SWIG_0, _swig_p)
	return
}

func (arg1 SwigcptrOp) EqualEqual__SWIG_0(arg2 Op) (_swig_ret bool) {
	return _swig_wrap_Op_EqualEqual__SWIG_0(arg1, arg2.Swigcptr())
}

var _wrap_Op_EqualEqual__SWIG_1 unsafe.Pointer

func _swig_wrap_Op_EqualEqual__SWIG_1(base SwigcptrOp, _ int) (_ bool) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_Op_EqualEqual__SWIG_1, _swig_p)
	return
}

func (arg1 SwigcptrOp) EqualEqual__SWIG_1(arg2 int) (_swig_ret bool) {
	return _swig_wrap_Op_EqualEqual__SWIG_1(arg1, arg2)
}

func (p SwigcptrOp) EqualEqual(a ...interface{}) bool {
	argc := len(a)
	if argc == 1 {
		if _, ok := a[0].(Op); !ok {
			goto check_1
		}
		return p.EqualEqual__SWIG_0(a[0].(Op))
	}
check_1:
	if argc == 1 {
		return p.EqualEqual__SWIG_1(a[0].(int))
	}
	panic("No match for overloaded function call")
}

var _wrap_Op_Plus__SWIG_0 unsafe.Pointer

func _swig_wrap_Op_Plus__SWIG_0(base SwigcptrOp, _ uintptr) (_ SwigcptrOp) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_Op_Plus__SWIG_0, _swig_p)
	return
}

func (arg1 SwigcptrOp) Plus__SWIG_0(arg2 Op) (_swig_ret Op) {
	return _swig_wrap_Op_Plus__SWIG_0(arg1, arg2.Swigcptr())
}

var _wrap_Op_Plus__SWIG_1 unsafe.Pointer

func _swig_wrap_Op_Plus__SWIG_1(base SwigcptrOp, _ int) (_ SwigcptrOp) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_Op_Plus__SWIG_1, _swig_p)
	return
}

func (arg1 SwigcptrOp) Plus__SWIG_1(arg2 int) (_swig_ret Op) {
	return _swig_wrap_Op_Plus__SWIG_1(arg1, arg2)
}

func (p SwigcptrOp) Plus(a ...interface{}) Op {
	argc := len(a)
	if argc == 1 {
		if _, ok := a[0].(Op); !ok {
			goto check_1
		}
		return p.Plus__SWIG_0(a[0].(Op))
	}
check_1:
	if argc == 1 {
		return p.Plus__SWIG_1(a[0].(int))
	}
	panic("No match for overloaded function call")
}

var _wrap_Op_Minus__SWIG_0 unsafe.Pointer

func _swig_wrap_Op_Minus__SWIG_0(base SwigcptrOp, _ uintptr) (_ SwigcptrOp) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_Op_Minus__SWIG_0, _swig_p)
	return
}

func (arg1 SwigcptrOp) Minus__SWIG_0(arg2 Op) (_swig_ret Op) {
	return _swig_wrap_Op_Minus__SWIG_0(arg1, arg2.Swigcptr())
}

var _wrap_Op_Minus__SWIG_1 unsafe.Pointer

func _swig_wrap_Op_Minus__SWIG_1(base SwigcptrOp, _ int) (_ SwigcptrOp) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_Op_Minus__SWIG_1, _swig_p)
	return
}

func (arg1 SwigcptrOp) Minus__SWIG_1(arg2 int) (_swig_ret Op) {
	return _swig_wrap_Op_Minus__SWIG_1(arg1, arg2)
}

func (p SwigcptrOp) Minus(a ...interface{}) Op {
	argc := len(a)
	if argc == 1 {
		if _, ok := a[0].(Op); !ok {
			goto check_1
		}
		return p.Minus__SWIG_0(a[0].(Op))
	}
check_1:
	if argc == 1 {
		return p.Minus__SWIG_1(a[0].(int))
	}
	panic("No match for overloaded function call")
}

var _wrap_Op___rsub__ unsafe.Pointer

func _swig_wrap_Op___rsub__(base SwigcptrOp, _ int) (_ SwigcptrOp) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_Op___rsub__, _swig_p)
	return
}

func (arg1 SwigcptrOp) X__rsub__(arg2 int) (_swig_ret Op) {
	return _swig_wrap_Op___rsub__(arg1, arg2)
}

var _wrap_Op_PlusPlusPrefix unsafe.Pointer

func _swig_wrap_Op_PlusPlusPrefix(base SwigcptrOp) (_ SwigcptrOp) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_Op_PlusPlusPrefix, _swig_p)
	return
}

func (arg1 SwigcptrOp) PlusPlusPrefix() (_swig_ret Op) {
	return _swig_wrap_Op_PlusPlusPrefix(arg1)
}

var _wrap_Op_PrintK unsafe.Pointer

func _swig_wrap_Op_PrintK(base SwigcptrOp) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_Op_PrintK, _swig_p)
	return
}

func (arg1 SwigcptrOp) PrintK() {
	_swig_wrap_Op_PrintK(arg1)
}

var _wrap_Op_k_set unsafe.Pointer

func _swig_wrap_Op_k_set(base SwigcptrOp, _ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_Op_k_set, _swig_p)
	return
}

func (arg1 SwigcptrOp) SetK(arg2 int) {
	_swig_wrap_Op_k_set(arg1, arg2)
}

var _wrap_Op_k_get unsafe.Pointer

func _swig_wrap_Op_k_get(base SwigcptrOp) (_ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_Op_k_get, _swig_p)
	return
}

func (arg1 SwigcptrOp) GetK() (_swig_ret int) {
	return _swig_wrap_Op_k_get(arg1)
}

var _wrap_delete_Op unsafe.Pointer

func _swig_wrap_delete_Op(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_Op, _swig_p)
	return
}

func DeleteOp(arg1 Op) {
	_swig_wrap_delete_Op(arg1.Swigcptr())
}

type Op interface {
	Swigcptr() uintptr
	SwigIsOp()
	EqualEqual(a ...interface{}) bool
	Plus(a ...interface{}) Op
	Minus(a ...interface{}) Op
	X__rsub__(arg2 int) (_swig_ret Op)
	PlusPlusPrefix() (_swig_ret Op)
	PrintK()
	SetK(arg2 int)
	GetK() (_swig_ret int)
}


type SwigcptrVoid uintptr
type Void interface {
	Swigcptr() uintptr;
}
func (p SwigcptrVoid) Swigcptr() uintptr {
	return uintptr(p)
}

