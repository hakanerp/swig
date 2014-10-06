/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 3.0.3
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: ./../li_std_deque.i

package li_std_deque

import "unsafe"
import _ "runtime/cgo"

var _cgo_runtime_cgocall func(unsafe.Pointer, uintptr)



type _ unsafe.Pointer



type _swig_fnptr *byte
type _swig_memberptr *byte

type SwigcptrIntDeque uintptr

func (p SwigcptrIntDeque) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrIntDeque) SwigIsIntDeque() {
}

var _wrap_IntDeque_empty unsafe.Pointer

func _swig_wrap_IntDeque_empty(base SwigcptrIntDeque) (_ bool) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_IntDeque_empty, _swig_p)
	return
}

func (arg1 SwigcptrIntDeque) Empty() (_swig_ret bool) {
	return _swig_wrap_IntDeque_empty(arg1)
}

var _wrap_new_IntDeque__SWIG_0 unsafe.Pointer

func _swig_wrap_new_IntDeque__SWIG_0() (base SwigcptrIntDeque) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_IntDeque__SWIG_0, _swig_p)
	return
}

func NewIntDeque__SWIG_0() (_swig_ret IntDeque) {
	return _swig_wrap_new_IntDeque__SWIG_0()
}

var _wrap_new_IntDeque__SWIG_1 unsafe.Pointer

func _swig_wrap_new_IntDeque__SWIG_1(base uint, _ int) (_ SwigcptrIntDeque) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_IntDeque__SWIG_1, _swig_p)
	return
}

func NewIntDeque__SWIG_1(arg1 uint, arg2 int) (_swig_ret IntDeque) {
	return _swig_wrap_new_IntDeque__SWIG_1(arg1, arg2)
}

var _wrap_new_IntDeque__SWIG_2 unsafe.Pointer

func _swig_wrap_new_IntDeque__SWIG_2(base uint) (_ SwigcptrIntDeque) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_IntDeque__SWIG_2, _swig_p)
	return
}

func NewIntDeque__SWIG_2(arg1 uint) (_swig_ret IntDeque) {
	return _swig_wrap_new_IntDeque__SWIG_2(arg1)
}

var _wrap_new_IntDeque__SWIG_3 unsafe.Pointer

func _swig_wrap_new_IntDeque__SWIG_3(base uintptr) (_ SwigcptrIntDeque) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_IntDeque__SWIG_3, _swig_p)
	return
}

func NewIntDeque__SWIG_3(arg1 IntDeque) (_swig_ret IntDeque) {
	return _swig_wrap_new_IntDeque__SWIG_3(arg1.Swigcptr())
}

func NewIntDeque(a ...interface{}) IntDeque {
	argc := len(a)
	if argc == 0 {
		return NewIntDeque__SWIG_0()
	}
	if argc == 1 {
		if _, ok := a[0].(IntDeque); !ok {
			goto check_2
		}
		return NewIntDeque__SWIG_3(a[0].(IntDeque))
	}
check_2:
	if argc == 1 {
		return NewIntDeque__SWIG_2(a[0].(uint))
	}
	if argc == 2 {
		return NewIntDeque__SWIG_1(a[0].(uint), a[1].(int))
	}
	panic("No match for overloaded function call")
}

var _wrap_delete_IntDeque unsafe.Pointer

func _swig_wrap_delete_IntDeque(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_IntDeque, _swig_p)
	return
}

func DeleteIntDeque(arg1 IntDeque) {
	_swig_wrap_delete_IntDeque(arg1.Swigcptr())
}

var _wrap_IntDeque_assign unsafe.Pointer

func _swig_wrap_IntDeque_assign(base SwigcptrIntDeque, _ uint, _ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_IntDeque_assign, _swig_p)
	return
}

func (arg1 SwigcptrIntDeque) Assign(arg2 uint, arg3 int) {
	_swig_wrap_IntDeque_assign(arg1, arg2, arg3)
}

var _wrap_IntDeque_swap unsafe.Pointer

func _swig_wrap_IntDeque_swap(base SwigcptrIntDeque, _ uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_IntDeque_swap, _swig_p)
	return
}

func (arg1 SwigcptrIntDeque) Swap(arg2 IntDeque) {
	_swig_wrap_IntDeque_swap(arg1, arg2.Swigcptr())
}

var _wrap_IntDeque_size unsafe.Pointer

func _swig_wrap_IntDeque_size(base SwigcptrIntDeque) (_ uint) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_IntDeque_size, _swig_p)
	return
}

func (arg1 SwigcptrIntDeque) Size() (_swig_ret uint) {
	return _swig_wrap_IntDeque_size(arg1)
}

var _wrap_IntDeque_max_size unsafe.Pointer

func _swig_wrap_IntDeque_max_size(base SwigcptrIntDeque) (_ uint) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_IntDeque_max_size, _swig_p)
	return
}

func (arg1 SwigcptrIntDeque) Max_size() (_swig_ret uint) {
	return _swig_wrap_IntDeque_max_size(arg1)
}

var _wrap_IntDeque_resize__SWIG_0 unsafe.Pointer

func _swig_wrap_IntDeque_resize__SWIG_0(base SwigcptrIntDeque, _ uint, _ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_IntDeque_resize__SWIG_0, _swig_p)
	return
}

func (arg1 SwigcptrIntDeque) Resize__SWIG_0(arg2 uint, arg3 int) {
	_swig_wrap_IntDeque_resize__SWIG_0(arg1, arg2, arg3)
}

var _wrap_IntDeque_resize__SWIG_1 unsafe.Pointer

func _swig_wrap_IntDeque_resize__SWIG_1(base SwigcptrIntDeque, _ uint) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_IntDeque_resize__SWIG_1, _swig_p)
	return
}

func (arg1 SwigcptrIntDeque) Resize__SWIG_1(arg2 uint) {
	_swig_wrap_IntDeque_resize__SWIG_1(arg1, arg2)
}

func (p SwigcptrIntDeque) Resize(a ...interface{}) {
	argc := len(a)
	if argc == 1 {
		p.Resize__SWIG_1(a[0].(uint))
		return
	}
	if argc == 2 {
		p.Resize__SWIG_0(a[0].(uint), a[1].(int))
		return
	}
	panic("No match for overloaded function call")
}

var _wrap_IntDeque_front unsafe.Pointer

func _swig_wrap_IntDeque_front(base SwigcptrIntDeque) (_ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_IntDeque_front, _swig_p)
	return
}

func (arg1 SwigcptrIntDeque) Front() (_swig_ret int) {
	return _swig_wrap_IntDeque_front(arg1)
}

var _wrap_IntDeque_back unsafe.Pointer

func _swig_wrap_IntDeque_back(base SwigcptrIntDeque) (_ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_IntDeque_back, _swig_p)
	return
}

func (arg1 SwigcptrIntDeque) Back() (_swig_ret int) {
	return _swig_wrap_IntDeque_back(arg1)
}

var _wrap_IntDeque_push_front unsafe.Pointer

func _swig_wrap_IntDeque_push_front(base SwigcptrIntDeque, _ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_IntDeque_push_front, _swig_p)
	return
}

func (arg1 SwigcptrIntDeque) Push_front(arg2 int) {
	_swig_wrap_IntDeque_push_front(arg1, arg2)
}

var _wrap_IntDeque_push_back unsafe.Pointer

func _swig_wrap_IntDeque_push_back(base SwigcptrIntDeque, _ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_IntDeque_push_back, _swig_p)
	return
}

func (arg1 SwigcptrIntDeque) Push_back(arg2 int) {
	_swig_wrap_IntDeque_push_back(arg1, arg2)
}

var _wrap_IntDeque_pop_front unsafe.Pointer

func _swig_wrap_IntDeque_pop_front(base SwigcptrIntDeque) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_IntDeque_pop_front, _swig_p)
	return
}

func (arg1 SwigcptrIntDeque) Pop_front() {
	_swig_wrap_IntDeque_pop_front(arg1)
}

var _wrap_IntDeque_pop_back unsafe.Pointer

func _swig_wrap_IntDeque_pop_back(base SwigcptrIntDeque) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_IntDeque_pop_back, _swig_p)
	return
}

func (arg1 SwigcptrIntDeque) Pop_back() {
	_swig_wrap_IntDeque_pop_back(arg1)
}

var _wrap_IntDeque_clear unsafe.Pointer

func _swig_wrap_IntDeque_clear(base SwigcptrIntDeque) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_IntDeque_clear, _swig_p)
	return
}

func (arg1 SwigcptrIntDeque) Clear() {
	_swig_wrap_IntDeque_clear(arg1)
}

var _wrap_IntDeque_getitem unsafe.Pointer

func _swig_wrap_IntDeque_getitem(base SwigcptrIntDeque, _ int) (_ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_IntDeque_getitem, _swig_p)
	return
}

func (arg1 SwigcptrIntDeque) Getitem(arg2 int) (_swig_ret int) {
	return _swig_wrap_IntDeque_getitem(arg1, arg2)
}

var _wrap_IntDeque_setitem unsafe.Pointer

func _swig_wrap_IntDeque_setitem(base SwigcptrIntDeque, _ int, _ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_IntDeque_setitem, _swig_p)
	return
}

func (arg1 SwigcptrIntDeque) Setitem(arg2 int, arg3 int) {
	_swig_wrap_IntDeque_setitem(arg1, arg2, arg3)
}

var _wrap_IntDeque_delitem unsafe.Pointer

func _swig_wrap_IntDeque_delitem(base SwigcptrIntDeque, _ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_IntDeque_delitem, _swig_p)
	return
}

func (arg1 SwigcptrIntDeque) Delitem(arg2 int) {
	_swig_wrap_IntDeque_delitem(arg1, arg2)
}

var _wrap_IntDeque_getslice unsafe.Pointer

func _swig_wrap_IntDeque_getslice(base SwigcptrIntDeque, _ int, _ int) (_ SwigcptrIntDeque) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_IntDeque_getslice, _swig_p)
	return
}

func (arg1 SwigcptrIntDeque) Getslice(arg2 int, arg3 int) (_swig_ret IntDeque) {
	return _swig_wrap_IntDeque_getslice(arg1, arg2, arg3)
}

var _wrap_IntDeque_setslice unsafe.Pointer

func _swig_wrap_IntDeque_setslice(base SwigcptrIntDeque, _ int, _ int, _ uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_IntDeque_setslice, _swig_p)
	return
}

func (arg1 SwigcptrIntDeque) Setslice(arg2 int, arg3 int, arg4 IntDeque) {
	_swig_wrap_IntDeque_setslice(arg1, arg2, arg3, arg4.Swigcptr())
}

var _wrap_IntDeque_delslice unsafe.Pointer

func _swig_wrap_IntDeque_delslice(base SwigcptrIntDeque, _ int, _ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_IntDeque_delslice, _swig_p)
	return
}

func (arg1 SwigcptrIntDeque) Delslice(arg2 int, arg3 int) {
	_swig_wrap_IntDeque_delslice(arg1, arg2, arg3)
}

type IntDeque interface {
	Swigcptr() uintptr
	SwigIsIntDeque()
	Empty() (_swig_ret bool)
	Assign(arg2 uint, arg3 int)
	Swap(arg2 IntDeque)
	Size() (_swig_ret uint)
	Max_size() (_swig_ret uint)
	Resize(a ...interface{})
	Front() (_swig_ret int)
	Back() (_swig_ret int)
	Push_front(arg2 int)
	Push_back(arg2 int)
	Pop_front()
	Pop_back()
	Clear()
	Getitem(arg2 int) (_swig_ret int)
	Setitem(arg2 int, arg3 int)
	Delitem(arg2 int)
	Getslice(arg2 int, arg3 int) (_swig_ret IntDeque)
	Setslice(arg2 int, arg3 int, arg4 IntDeque)
	Delslice(arg2 int, arg3 int)
}

type SwigcptrDoubleDeque uintptr

func (p SwigcptrDoubleDeque) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrDoubleDeque) SwigIsDoubleDeque() {
}

var _wrap_DoubleDeque_empty unsafe.Pointer

func _swig_wrap_DoubleDeque_empty(base SwigcptrDoubleDeque) (_ bool) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_DoubleDeque_empty, _swig_p)
	return
}

func (arg1 SwigcptrDoubleDeque) Empty() (_swig_ret bool) {
	return _swig_wrap_DoubleDeque_empty(arg1)
}

var _wrap_new_DoubleDeque__SWIG_0 unsafe.Pointer

func _swig_wrap_new_DoubleDeque__SWIG_0() (base SwigcptrDoubleDeque) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_DoubleDeque__SWIG_0, _swig_p)
	return
}

func NewDoubleDeque__SWIG_0() (_swig_ret DoubleDeque) {
	return _swig_wrap_new_DoubleDeque__SWIG_0()
}

var _wrap_new_DoubleDeque__SWIG_1 unsafe.Pointer

func _swig_wrap_new_DoubleDeque__SWIG_1(base uint, _ float64) (_ SwigcptrDoubleDeque) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_DoubleDeque__SWIG_1, _swig_p)
	return
}

func NewDoubleDeque__SWIG_1(arg1 uint, arg2 float64) (_swig_ret DoubleDeque) {
	return _swig_wrap_new_DoubleDeque__SWIG_1(arg1, arg2)
}

var _wrap_new_DoubleDeque__SWIG_2 unsafe.Pointer

func _swig_wrap_new_DoubleDeque__SWIG_2(base uint) (_ SwigcptrDoubleDeque) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_DoubleDeque__SWIG_2, _swig_p)
	return
}

func NewDoubleDeque__SWIG_2(arg1 uint) (_swig_ret DoubleDeque) {
	return _swig_wrap_new_DoubleDeque__SWIG_2(arg1)
}

var _wrap_new_DoubleDeque__SWIG_3 unsafe.Pointer

func _swig_wrap_new_DoubleDeque__SWIG_3(base uintptr) (_ SwigcptrDoubleDeque) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_DoubleDeque__SWIG_3, _swig_p)
	return
}

func NewDoubleDeque__SWIG_3(arg1 DoubleDeque) (_swig_ret DoubleDeque) {
	return _swig_wrap_new_DoubleDeque__SWIG_3(arg1.Swigcptr())
}

func NewDoubleDeque(a ...interface{}) DoubleDeque {
	argc := len(a)
	if argc == 0 {
		return NewDoubleDeque__SWIG_0()
	}
	if argc == 1 {
		if _, ok := a[0].(DoubleDeque); !ok {
			goto check_2
		}
		return NewDoubleDeque__SWIG_3(a[0].(DoubleDeque))
	}
check_2:
	if argc == 1 {
		return NewDoubleDeque__SWIG_2(a[0].(uint))
	}
	if argc == 2 {
		return NewDoubleDeque__SWIG_1(a[0].(uint), a[1].(float64))
	}
	panic("No match for overloaded function call")
}

var _wrap_delete_DoubleDeque unsafe.Pointer

func _swig_wrap_delete_DoubleDeque(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_DoubleDeque, _swig_p)
	return
}

func DeleteDoubleDeque(arg1 DoubleDeque) {
	_swig_wrap_delete_DoubleDeque(arg1.Swigcptr())
}

var _wrap_DoubleDeque_assign unsafe.Pointer

func _swig_wrap_DoubleDeque_assign(base SwigcptrDoubleDeque, _ uint, _ float64) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_DoubleDeque_assign, _swig_p)
	return
}

func (arg1 SwigcptrDoubleDeque) Assign(arg2 uint, arg3 float64) {
	_swig_wrap_DoubleDeque_assign(arg1, arg2, arg3)
}

var _wrap_DoubleDeque_swap unsafe.Pointer

func _swig_wrap_DoubleDeque_swap(base SwigcptrDoubleDeque, _ uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_DoubleDeque_swap, _swig_p)
	return
}

func (arg1 SwigcptrDoubleDeque) Swap(arg2 DoubleDeque) {
	_swig_wrap_DoubleDeque_swap(arg1, arg2.Swigcptr())
}

var _wrap_DoubleDeque_size unsafe.Pointer

func _swig_wrap_DoubleDeque_size(base SwigcptrDoubleDeque) (_ uint) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_DoubleDeque_size, _swig_p)
	return
}

func (arg1 SwigcptrDoubleDeque) Size() (_swig_ret uint) {
	return _swig_wrap_DoubleDeque_size(arg1)
}

var _wrap_DoubleDeque_max_size unsafe.Pointer

func _swig_wrap_DoubleDeque_max_size(base SwigcptrDoubleDeque) (_ uint) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_DoubleDeque_max_size, _swig_p)
	return
}

func (arg1 SwigcptrDoubleDeque) Max_size() (_swig_ret uint) {
	return _swig_wrap_DoubleDeque_max_size(arg1)
}

var _wrap_DoubleDeque_resize__SWIG_0 unsafe.Pointer

func _swig_wrap_DoubleDeque_resize__SWIG_0(base SwigcptrDoubleDeque, _ uint, _ float64) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_DoubleDeque_resize__SWIG_0, _swig_p)
	return
}

func (arg1 SwigcptrDoubleDeque) Resize__SWIG_0(arg2 uint, arg3 float64) {
	_swig_wrap_DoubleDeque_resize__SWIG_0(arg1, arg2, arg3)
}

var _wrap_DoubleDeque_resize__SWIG_1 unsafe.Pointer

func _swig_wrap_DoubleDeque_resize__SWIG_1(base SwigcptrDoubleDeque, _ uint) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_DoubleDeque_resize__SWIG_1, _swig_p)
	return
}

func (arg1 SwigcptrDoubleDeque) Resize__SWIG_1(arg2 uint) {
	_swig_wrap_DoubleDeque_resize__SWIG_1(arg1, arg2)
}

func (p SwigcptrDoubleDeque) Resize(a ...interface{}) {
	argc := len(a)
	if argc == 1 {
		p.Resize__SWIG_1(a[0].(uint))
		return
	}
	if argc == 2 {
		p.Resize__SWIG_0(a[0].(uint), a[1].(float64))
		return
	}
	panic("No match for overloaded function call")
}

var _wrap_DoubleDeque_front unsafe.Pointer

func _swig_wrap_DoubleDeque_front(base SwigcptrDoubleDeque) (_ float64) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_DoubleDeque_front, _swig_p)
	return
}

func (arg1 SwigcptrDoubleDeque) Front() (_swig_ret float64) {
	return _swig_wrap_DoubleDeque_front(arg1)
}

var _wrap_DoubleDeque_back unsafe.Pointer

func _swig_wrap_DoubleDeque_back(base SwigcptrDoubleDeque) (_ float64) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_DoubleDeque_back, _swig_p)
	return
}

func (arg1 SwigcptrDoubleDeque) Back() (_swig_ret float64) {
	return _swig_wrap_DoubleDeque_back(arg1)
}

var _wrap_DoubleDeque_push_front unsafe.Pointer

func _swig_wrap_DoubleDeque_push_front(base SwigcptrDoubleDeque, _ float64) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_DoubleDeque_push_front, _swig_p)
	return
}

func (arg1 SwigcptrDoubleDeque) Push_front(arg2 float64) {
	_swig_wrap_DoubleDeque_push_front(arg1, arg2)
}

var _wrap_DoubleDeque_push_back unsafe.Pointer

func _swig_wrap_DoubleDeque_push_back(base SwigcptrDoubleDeque, _ float64) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_DoubleDeque_push_back, _swig_p)
	return
}

func (arg1 SwigcptrDoubleDeque) Push_back(arg2 float64) {
	_swig_wrap_DoubleDeque_push_back(arg1, arg2)
}

var _wrap_DoubleDeque_pop_front unsafe.Pointer

func _swig_wrap_DoubleDeque_pop_front(base SwigcptrDoubleDeque) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_DoubleDeque_pop_front, _swig_p)
	return
}

func (arg1 SwigcptrDoubleDeque) Pop_front() {
	_swig_wrap_DoubleDeque_pop_front(arg1)
}

var _wrap_DoubleDeque_pop_back unsafe.Pointer

func _swig_wrap_DoubleDeque_pop_back(base SwigcptrDoubleDeque) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_DoubleDeque_pop_back, _swig_p)
	return
}

func (arg1 SwigcptrDoubleDeque) Pop_back() {
	_swig_wrap_DoubleDeque_pop_back(arg1)
}

var _wrap_DoubleDeque_clear unsafe.Pointer

func _swig_wrap_DoubleDeque_clear(base SwigcptrDoubleDeque) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_DoubleDeque_clear, _swig_p)
	return
}

func (arg1 SwigcptrDoubleDeque) Clear() {
	_swig_wrap_DoubleDeque_clear(arg1)
}

var _wrap_DoubleDeque_getitem unsafe.Pointer

func _swig_wrap_DoubleDeque_getitem(base SwigcptrDoubleDeque, _ int) (_ float64) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_DoubleDeque_getitem, _swig_p)
	return
}

func (arg1 SwigcptrDoubleDeque) Getitem(arg2 int) (_swig_ret float64) {
	return _swig_wrap_DoubleDeque_getitem(arg1, arg2)
}

var _wrap_DoubleDeque_setitem unsafe.Pointer

func _swig_wrap_DoubleDeque_setitem(base SwigcptrDoubleDeque, _ int, _ float64) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_DoubleDeque_setitem, _swig_p)
	return
}

func (arg1 SwigcptrDoubleDeque) Setitem(arg2 int, arg3 float64) {
	_swig_wrap_DoubleDeque_setitem(arg1, arg2, arg3)
}

var _wrap_DoubleDeque_delitem unsafe.Pointer

func _swig_wrap_DoubleDeque_delitem(base SwigcptrDoubleDeque, _ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_DoubleDeque_delitem, _swig_p)
	return
}

func (arg1 SwigcptrDoubleDeque) Delitem(arg2 int) {
	_swig_wrap_DoubleDeque_delitem(arg1, arg2)
}

var _wrap_DoubleDeque_getslice unsafe.Pointer

func _swig_wrap_DoubleDeque_getslice(base SwigcptrDoubleDeque, _ int, _ int) (_ SwigcptrDoubleDeque) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_DoubleDeque_getslice, _swig_p)
	return
}

func (arg1 SwigcptrDoubleDeque) Getslice(arg2 int, arg3 int) (_swig_ret DoubleDeque) {
	return _swig_wrap_DoubleDeque_getslice(arg1, arg2, arg3)
}

var _wrap_DoubleDeque_setslice unsafe.Pointer

func _swig_wrap_DoubleDeque_setslice(base SwigcptrDoubleDeque, _ int, _ int, _ uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_DoubleDeque_setslice, _swig_p)
	return
}

func (arg1 SwigcptrDoubleDeque) Setslice(arg2 int, arg3 int, arg4 DoubleDeque) {
	_swig_wrap_DoubleDeque_setslice(arg1, arg2, arg3, arg4.Swigcptr())
}

var _wrap_DoubleDeque_delslice unsafe.Pointer

func _swig_wrap_DoubleDeque_delslice(base SwigcptrDoubleDeque, _ int, _ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_DoubleDeque_delslice, _swig_p)
	return
}

func (arg1 SwigcptrDoubleDeque) Delslice(arg2 int, arg3 int) {
	_swig_wrap_DoubleDeque_delslice(arg1, arg2, arg3)
}

type DoubleDeque interface {
	Swigcptr() uintptr
	SwigIsDoubleDeque()
	Empty() (_swig_ret bool)
	Assign(arg2 uint, arg3 float64)
	Swap(arg2 DoubleDeque)
	Size() (_swig_ret uint)
	Max_size() (_swig_ret uint)
	Resize(a ...interface{})
	Front() (_swig_ret float64)
	Back() (_swig_ret float64)
	Push_front(arg2 float64)
	Push_back(arg2 float64)
	Pop_front()
	Pop_back()
	Clear()
	Getitem(arg2 int) (_swig_ret float64)
	Setitem(arg2 int, arg3 float64)
	Delitem(arg2 int)
	Getslice(arg2 int, arg3 int) (_swig_ret DoubleDeque)
	Setslice(arg2 int, arg3 int, arg4 DoubleDeque)
	Delslice(arg2 int, arg3 int)
}

type SwigcptrRealDeque uintptr

func (p SwigcptrRealDeque) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrRealDeque) SwigIsRealDeque() {
}

var _wrap_RealDeque_empty unsafe.Pointer

func _swig_wrap_RealDeque_empty(base SwigcptrRealDeque) (_ bool) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_RealDeque_empty, _swig_p)
	return
}

func (arg1 SwigcptrRealDeque) Empty() (_swig_ret bool) {
	return _swig_wrap_RealDeque_empty(arg1)
}

var _wrap_new_RealDeque__SWIG_0 unsafe.Pointer

func _swig_wrap_new_RealDeque__SWIG_0() (base SwigcptrRealDeque) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_RealDeque__SWIG_0, _swig_p)
	return
}

func NewRealDeque__SWIG_0() (_swig_ret RealDeque) {
	return _swig_wrap_new_RealDeque__SWIG_0()
}

var _wrap_new_RealDeque__SWIG_1 unsafe.Pointer

func _swig_wrap_new_RealDeque__SWIG_1(base uint, _ float32) (_ SwigcptrRealDeque) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_RealDeque__SWIG_1, _swig_p)
	return
}

func NewRealDeque__SWIG_1(arg1 uint, arg2 float32) (_swig_ret RealDeque) {
	return _swig_wrap_new_RealDeque__SWIG_1(arg1, arg2)
}

var _wrap_new_RealDeque__SWIG_2 unsafe.Pointer

func _swig_wrap_new_RealDeque__SWIG_2(base uint) (_ SwigcptrRealDeque) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_RealDeque__SWIG_2, _swig_p)
	return
}

func NewRealDeque__SWIG_2(arg1 uint) (_swig_ret RealDeque) {
	return _swig_wrap_new_RealDeque__SWIG_2(arg1)
}

var _wrap_new_RealDeque__SWIG_3 unsafe.Pointer

func _swig_wrap_new_RealDeque__SWIG_3(base uintptr) (_ SwigcptrRealDeque) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_RealDeque__SWIG_3, _swig_p)
	return
}

func NewRealDeque__SWIG_3(arg1 RealDeque) (_swig_ret RealDeque) {
	return _swig_wrap_new_RealDeque__SWIG_3(arg1.Swigcptr())
}

func NewRealDeque(a ...interface{}) RealDeque {
	argc := len(a)
	if argc == 0 {
		return NewRealDeque__SWIG_0()
	}
	if argc == 1 {
		if _, ok := a[0].(RealDeque); !ok {
			goto check_2
		}
		return NewRealDeque__SWIG_3(a[0].(RealDeque))
	}
check_2:
	if argc == 1 {
		return NewRealDeque__SWIG_2(a[0].(uint))
	}
	if argc == 2 {
		return NewRealDeque__SWIG_1(a[0].(uint), a[1].(float32))
	}
	panic("No match for overloaded function call")
}

var _wrap_delete_RealDeque unsafe.Pointer

func _swig_wrap_delete_RealDeque(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_RealDeque, _swig_p)
	return
}

func DeleteRealDeque(arg1 RealDeque) {
	_swig_wrap_delete_RealDeque(arg1.Swigcptr())
}

var _wrap_RealDeque_assign unsafe.Pointer

func _swig_wrap_RealDeque_assign(base SwigcptrRealDeque, _ uint, _ float32) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_RealDeque_assign, _swig_p)
	return
}

func (arg1 SwigcptrRealDeque) Assign(arg2 uint, arg3 float32) {
	_swig_wrap_RealDeque_assign(arg1, arg2, arg3)
}

var _wrap_RealDeque_swap unsafe.Pointer

func _swig_wrap_RealDeque_swap(base SwigcptrRealDeque, _ uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_RealDeque_swap, _swig_p)
	return
}

func (arg1 SwigcptrRealDeque) Swap(arg2 RealDeque) {
	_swig_wrap_RealDeque_swap(arg1, arg2.Swigcptr())
}

var _wrap_RealDeque_size unsafe.Pointer

func _swig_wrap_RealDeque_size(base SwigcptrRealDeque) (_ uint) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_RealDeque_size, _swig_p)
	return
}

func (arg1 SwigcptrRealDeque) Size() (_swig_ret uint) {
	return _swig_wrap_RealDeque_size(arg1)
}

var _wrap_RealDeque_max_size unsafe.Pointer

func _swig_wrap_RealDeque_max_size(base SwigcptrRealDeque) (_ uint) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_RealDeque_max_size, _swig_p)
	return
}

func (arg1 SwigcptrRealDeque) Max_size() (_swig_ret uint) {
	return _swig_wrap_RealDeque_max_size(arg1)
}

var _wrap_RealDeque_resize__SWIG_0 unsafe.Pointer

func _swig_wrap_RealDeque_resize__SWIG_0(base SwigcptrRealDeque, _ uint, _ float32) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_RealDeque_resize__SWIG_0, _swig_p)
	return
}

func (arg1 SwigcptrRealDeque) Resize__SWIG_0(arg2 uint, arg3 float32) {
	_swig_wrap_RealDeque_resize__SWIG_0(arg1, arg2, arg3)
}

var _wrap_RealDeque_resize__SWIG_1 unsafe.Pointer

func _swig_wrap_RealDeque_resize__SWIG_1(base SwigcptrRealDeque, _ uint) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_RealDeque_resize__SWIG_1, _swig_p)
	return
}

func (arg1 SwigcptrRealDeque) Resize__SWIG_1(arg2 uint) {
	_swig_wrap_RealDeque_resize__SWIG_1(arg1, arg2)
}

func (p SwigcptrRealDeque) Resize(a ...interface{}) {
	argc := len(a)
	if argc == 1 {
		p.Resize__SWIG_1(a[0].(uint))
		return
	}
	if argc == 2 {
		p.Resize__SWIG_0(a[0].(uint), a[1].(float32))
		return
	}
	panic("No match for overloaded function call")
}

var _wrap_RealDeque_front unsafe.Pointer

func _swig_wrap_RealDeque_front(base SwigcptrRealDeque) (_ float32) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_RealDeque_front, _swig_p)
	return
}

func (arg1 SwigcptrRealDeque) Front() (_swig_ret float32) {
	return _swig_wrap_RealDeque_front(arg1)
}

var _wrap_RealDeque_back unsafe.Pointer

func _swig_wrap_RealDeque_back(base SwigcptrRealDeque) (_ float32) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_RealDeque_back, _swig_p)
	return
}

func (arg1 SwigcptrRealDeque) Back() (_swig_ret float32) {
	return _swig_wrap_RealDeque_back(arg1)
}

var _wrap_RealDeque_push_front unsafe.Pointer

func _swig_wrap_RealDeque_push_front(base SwigcptrRealDeque, _ float32) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_RealDeque_push_front, _swig_p)
	return
}

func (arg1 SwigcptrRealDeque) Push_front(arg2 float32) {
	_swig_wrap_RealDeque_push_front(arg1, arg2)
}

var _wrap_RealDeque_push_back unsafe.Pointer

func _swig_wrap_RealDeque_push_back(base SwigcptrRealDeque, _ float32) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_RealDeque_push_back, _swig_p)
	return
}

func (arg1 SwigcptrRealDeque) Push_back(arg2 float32) {
	_swig_wrap_RealDeque_push_back(arg1, arg2)
}

var _wrap_RealDeque_pop_front unsafe.Pointer

func _swig_wrap_RealDeque_pop_front(base SwigcptrRealDeque) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_RealDeque_pop_front, _swig_p)
	return
}

func (arg1 SwigcptrRealDeque) Pop_front() {
	_swig_wrap_RealDeque_pop_front(arg1)
}

var _wrap_RealDeque_pop_back unsafe.Pointer

func _swig_wrap_RealDeque_pop_back(base SwigcptrRealDeque) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_RealDeque_pop_back, _swig_p)
	return
}

func (arg1 SwigcptrRealDeque) Pop_back() {
	_swig_wrap_RealDeque_pop_back(arg1)
}

var _wrap_RealDeque_clear unsafe.Pointer

func _swig_wrap_RealDeque_clear(base SwigcptrRealDeque) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_RealDeque_clear, _swig_p)
	return
}

func (arg1 SwigcptrRealDeque) Clear() {
	_swig_wrap_RealDeque_clear(arg1)
}

var _wrap_RealDeque_getitem unsafe.Pointer

func _swig_wrap_RealDeque_getitem(base SwigcptrRealDeque, _ int) (_ float32) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_RealDeque_getitem, _swig_p)
	return
}

func (arg1 SwigcptrRealDeque) Getitem(arg2 int) (_swig_ret float32) {
	return _swig_wrap_RealDeque_getitem(arg1, arg2)
}

var _wrap_RealDeque_setitem unsafe.Pointer

func _swig_wrap_RealDeque_setitem(base SwigcptrRealDeque, _ int, _ float32) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_RealDeque_setitem, _swig_p)
	return
}

func (arg1 SwigcptrRealDeque) Setitem(arg2 int, arg3 float32) {
	_swig_wrap_RealDeque_setitem(arg1, arg2, arg3)
}

var _wrap_RealDeque_delitem unsafe.Pointer

func _swig_wrap_RealDeque_delitem(base SwigcptrRealDeque, _ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_RealDeque_delitem, _swig_p)
	return
}

func (arg1 SwigcptrRealDeque) Delitem(arg2 int) {
	_swig_wrap_RealDeque_delitem(arg1, arg2)
}

var _wrap_RealDeque_getslice unsafe.Pointer

func _swig_wrap_RealDeque_getslice(base SwigcptrRealDeque, _ int, _ int) (_ SwigcptrRealDeque) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_RealDeque_getslice, _swig_p)
	return
}

func (arg1 SwigcptrRealDeque) Getslice(arg2 int, arg3 int) (_swig_ret RealDeque) {
	return _swig_wrap_RealDeque_getslice(arg1, arg2, arg3)
}

var _wrap_RealDeque_setslice unsafe.Pointer

func _swig_wrap_RealDeque_setslice(base SwigcptrRealDeque, _ int, _ int, _ uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_RealDeque_setslice, _swig_p)
	return
}

func (arg1 SwigcptrRealDeque) Setslice(arg2 int, arg3 int, arg4 RealDeque) {
	_swig_wrap_RealDeque_setslice(arg1, arg2, arg3, arg4.Swigcptr())
}

var _wrap_RealDeque_delslice unsafe.Pointer

func _swig_wrap_RealDeque_delslice(base SwigcptrRealDeque, _ int, _ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_RealDeque_delslice, _swig_p)
	return
}

func (arg1 SwigcptrRealDeque) Delslice(arg2 int, arg3 int) {
	_swig_wrap_RealDeque_delslice(arg1, arg2, arg3)
}

type RealDeque interface {
	Swigcptr() uintptr
	SwigIsRealDeque()
	Empty() (_swig_ret bool)
	Assign(arg2 uint, arg3 float32)
	Swap(arg2 RealDeque)
	Size() (_swig_ret uint)
	Max_size() (_swig_ret uint)
	Resize(a ...interface{})
	Front() (_swig_ret float32)
	Back() (_swig_ret float32)
	Push_front(arg2 float32)
	Push_back(arg2 float32)
	Pop_front()
	Pop_back()
	Clear()
	Getitem(arg2 int) (_swig_ret float32)
	Setitem(arg2 int, arg3 float32)
	Delitem(arg2 int)
	Getslice(arg2 int, arg3 int) (_swig_ret RealDeque)
	Setslice(arg2 int, arg3 int, arg4 RealDeque)
	Delslice(arg2 int, arg3 int)
}

var _wrap_average unsafe.Pointer

func _swig_wrap_average(base uintptr) (_ float64) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_average, _swig_p)
	return
}

func Average(arg1 IntDeque) (_swig_ret float64) {
	return _swig_wrap_average(arg1.Swigcptr())
}

var _wrap_half unsafe.Pointer

func _swig_wrap_half(base uintptr) (_ SwigcptrRealDeque) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_half, _swig_p)
	return
}

func Half(arg1 RealDeque) (_swig_ret RealDeque) {
	return _swig_wrap_half(arg1.Swigcptr())
}

var _wrap_halve_in_place unsafe.Pointer

func _swig_wrap_halve_in_place(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_halve_in_place, _swig_p)
	return
}

func Halve_in_place(arg1 DoubleDeque) {
	_swig_wrap_halve_in_place(arg1.Swigcptr())
}


type SwigcptrVoid uintptr
type Void interface {
	Swigcptr() uintptr;
}
func (p SwigcptrVoid) Swigcptr() uintptr {
	return uintptr(p)
}

