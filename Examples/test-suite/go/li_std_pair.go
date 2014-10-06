/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 3.0.3
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: ./../li_std_pair.i

package li_std_pair

import "unsafe"
import _ "runtime/cgo"

var _cgo_runtime_cgocall func(unsafe.Pointer, uintptr)



type _ unsafe.Pointer



type _swig_fnptr *byte
type _swig_memberptr *byte

type SwigcptrIntPair uintptr

func (p SwigcptrIntPair) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrIntPair) SwigIsIntPair() {
}

var _wrap_new_IntPair__SWIG_0 unsafe.Pointer

func _swig_wrap_new_IntPair__SWIG_0() (base SwigcptrIntPair) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_IntPair__SWIG_0, _swig_p)
	return
}

func NewIntPair__SWIG_0() (_swig_ret IntPair) {
	return _swig_wrap_new_IntPair__SWIG_0()
}

var _wrap_new_IntPair__SWIG_1 unsafe.Pointer

func _swig_wrap_new_IntPair__SWIG_1(base int, _ int) (_ SwigcptrIntPair) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_IntPair__SWIG_1, _swig_p)
	return
}

func NewIntPair__SWIG_1(arg1 int, arg2 int) (_swig_ret IntPair) {
	return _swig_wrap_new_IntPair__SWIG_1(arg1, arg2)
}

var _wrap_new_IntPair__SWIG_2 unsafe.Pointer

func _swig_wrap_new_IntPair__SWIG_2(base uintptr) (_ SwigcptrIntPair) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_IntPair__SWIG_2, _swig_p)
	return
}

func NewIntPair__SWIG_2(arg1 IntPair) (_swig_ret IntPair) {
	return _swig_wrap_new_IntPair__SWIG_2(arg1.Swigcptr())
}

func NewIntPair(a ...interface{}) IntPair {
	argc := len(a)
	if argc == 0 {
		return NewIntPair__SWIG_0()
	}
	if argc == 1 {
		return NewIntPair__SWIG_2(a[0].(IntPair))
	}
	if argc == 2 {
		return NewIntPair__SWIG_1(a[0].(int), a[1].(int))
	}
	panic("No match for overloaded function call")
}

var _wrap_IntPair_first_set unsafe.Pointer

func _swig_wrap_IntPair_first_set(base SwigcptrIntPair, _ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_IntPair_first_set, _swig_p)
	return
}

func (arg1 SwigcptrIntPair) SetFirst(arg2 int) {
	_swig_wrap_IntPair_first_set(arg1, arg2)
}

var _wrap_IntPair_first_get unsafe.Pointer

func _swig_wrap_IntPair_first_get(base SwigcptrIntPair) (_ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_IntPair_first_get, _swig_p)
	return
}

func (arg1 SwigcptrIntPair) GetFirst() (_swig_ret int) {
	return _swig_wrap_IntPair_first_get(arg1)
}

var _wrap_IntPair_second_set unsafe.Pointer

func _swig_wrap_IntPair_second_set(base SwigcptrIntPair, _ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_IntPair_second_set, _swig_p)
	return
}

func (arg1 SwigcptrIntPair) SetSecond(arg2 int) {
	_swig_wrap_IntPair_second_set(arg1, arg2)
}

var _wrap_IntPair_second_get unsafe.Pointer

func _swig_wrap_IntPair_second_get(base SwigcptrIntPair) (_ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_IntPair_second_get, _swig_p)
	return
}

func (arg1 SwigcptrIntPair) GetSecond() (_swig_ret int) {
	return _swig_wrap_IntPair_second_get(arg1)
}

var _wrap_delete_IntPair unsafe.Pointer

func _swig_wrap_delete_IntPair(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_IntPair, _swig_p)
	return
}

func DeleteIntPair(arg1 IntPair) {
	_swig_wrap_delete_IntPair(arg1.Swigcptr())
}

type IntPair interface {
	Swigcptr() uintptr
	SwigIsIntPair()
	SetFirst(arg2 int)
	GetFirst() (_swig_ret int)
	SetSecond(arg2 int)
	GetSecond() (_swig_ret int)
}

var _wrap_makeIntPair unsafe.Pointer

func _swig_wrap_makeIntPair(base int, _ int) (_ SwigcptrIntPair) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_makeIntPair, _swig_p)
	return
}

func MakeIntPair(arg1 int, arg2 int) (_swig_ret IntPair) {
	return _swig_wrap_makeIntPair(arg1, arg2)
}

var _wrap_makeIntPairPtr unsafe.Pointer

func _swig_wrap_makeIntPairPtr(base int, _ int) (_ SwigcptrIntPair) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_makeIntPairPtr, _swig_p)
	return
}

func MakeIntPairPtr(arg1 int, arg2 int) (_swig_ret IntPair) {
	return _swig_wrap_makeIntPairPtr(arg1, arg2)
}

var _wrap_makeIntPairRef unsafe.Pointer

func _swig_wrap_makeIntPairRef(base int, _ int) (_ SwigcptrIntPair) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_makeIntPairRef, _swig_p)
	return
}

func MakeIntPairRef(arg1 int, arg2 int) (_swig_ret IntPair) {
	return _swig_wrap_makeIntPairRef(arg1, arg2)
}

var _wrap_makeIntPairConstRef unsafe.Pointer

func _swig_wrap_makeIntPairConstRef(base int, _ int) (_ SwigcptrIntPair) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_makeIntPairConstRef, _swig_p)
	return
}

func MakeIntPairConstRef(arg1 int, arg2 int) (_swig_ret IntPair) {
	return _swig_wrap_makeIntPairConstRef(arg1, arg2)
}

var _wrap_product1 unsafe.Pointer

func _swig_wrap_product1(base uintptr) (_ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_product1, _swig_p)
	return
}

func Product1(arg1 IntPair) (_swig_ret int) {
	return _swig_wrap_product1(arg1.Swigcptr())
}

var _wrap_product2 unsafe.Pointer

func _swig_wrap_product2(base uintptr) (_ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_product2, _swig_p)
	return
}

func Product2(arg1 IntPair) (_swig_ret int) {
	return _swig_wrap_product2(arg1.Swigcptr())
}

var _wrap_product3 unsafe.Pointer

func _swig_wrap_product3(base uintptr) (_ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_product3, _swig_p)
	return
}

func Product3(arg1 IntPair) (_swig_ret int) {
	return _swig_wrap_product3(arg1.Swigcptr())
}


type SwigcptrVoid uintptr
type Void interface {
	Swigcptr() uintptr;
}
func (p SwigcptrVoid) Swigcptr() uintptr {
	return uintptr(p)
}

