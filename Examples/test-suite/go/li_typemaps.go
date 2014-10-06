/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 3.0.3
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: ./../li_typemaps.i

package li_typemaps

import "unsafe"
import _ "runtime/cgo"

var _cgo_runtime_cgocall func(unsafe.Pointer, uintptr)



type _ unsafe.Pointer



type _swig_fnptr *byte
type _swig_memberptr *byte

type SwigcptrFoo uintptr

func (p SwigcptrFoo) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrFoo) SwigIsFoo() {
}

var _wrap_Foo_a_set unsafe.Pointer

func _swig_wrap_Foo_a_set(base SwigcptrFoo, _ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_Foo_a_set, _swig_p)
	return
}

func (arg1 SwigcptrFoo) SetA(arg2 int) {
	_swig_wrap_Foo_a_set(arg1, arg2)
}

var _wrap_Foo_a_get unsafe.Pointer

func _swig_wrap_Foo_a_get(base SwigcptrFoo) (_ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_Foo_a_get, _swig_p)
	return
}

func (arg1 SwigcptrFoo) GetA() (_swig_ret int) {
	return _swig_wrap_Foo_a_get(arg1)
}

var _wrap_new_Foo unsafe.Pointer

func _swig_wrap_new_Foo() (base SwigcptrFoo) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_Foo, _swig_p)
	return
}

func NewFoo() (_swig_ret Foo) {
	return _swig_wrap_new_Foo()
}

var _wrap_delete_Foo unsafe.Pointer

func _swig_wrap_delete_Foo(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_Foo, _swig_p)
	return
}

func DeleteFoo(arg1 Foo) {
	_swig_wrap_delete_Foo(arg1.Swigcptr())
}

type Foo interface {
	Swigcptr() uintptr
	SwigIsFoo()
	SetA(arg2 int)
	GetA() (_swig_ret int)
}

var _wrap_in_bool unsafe.Pointer

func In_bool(arg1 bool) (_swig_ret bool) {
	_swig_p := uintptr(unsafe.Pointer(&arg1))
	_cgo_runtime_cgocall(_wrap_in_bool, _swig_p)
	return
}
var _wrap_in_int unsafe.Pointer

func In_int(arg1 int) (_swig_ret int) {
	_swig_p := uintptr(unsafe.Pointer(&arg1))
	_cgo_runtime_cgocall(_wrap_in_int, _swig_p)
	return
}
var _wrap_in_long unsafe.Pointer

func In_long(arg1 int64) (_swig_ret int64) {
	_swig_p := uintptr(unsafe.Pointer(&arg1))
	_cgo_runtime_cgocall(_wrap_in_long, _swig_p)
	return
}
var _wrap_in_short unsafe.Pointer

func In_short(arg1 int16) (_swig_ret int16) {
	_swig_p := uintptr(unsafe.Pointer(&arg1))
	_cgo_runtime_cgocall(_wrap_in_short, _swig_p)
	return
}
var _wrap_in_uint unsafe.Pointer

func In_uint(arg1 uint) (_swig_ret uint) {
	_swig_p := uintptr(unsafe.Pointer(&arg1))
	_cgo_runtime_cgocall(_wrap_in_uint, _swig_p)
	return
}
var _wrap_in_ushort unsafe.Pointer

func In_ushort(arg1 uint16) (_swig_ret uint16) {
	_swig_p := uintptr(unsafe.Pointer(&arg1))
	_cgo_runtime_cgocall(_wrap_in_ushort, _swig_p)
	return
}
var _wrap_in_ulong unsafe.Pointer

func In_ulong(arg1 uint64) (_swig_ret uint64) {
	_swig_p := uintptr(unsafe.Pointer(&arg1))
	_cgo_runtime_cgocall(_wrap_in_ulong, _swig_p)
	return
}
var _wrap_in_uchar unsafe.Pointer

func In_uchar(arg1 byte) (_swig_ret byte) {
	_swig_p := uintptr(unsafe.Pointer(&arg1))
	_cgo_runtime_cgocall(_wrap_in_uchar, _swig_p)
	return
}
var _wrap_in_schar unsafe.Pointer

func In_schar(arg1 int8) (_swig_ret int8) {
	_swig_p := uintptr(unsafe.Pointer(&arg1))
	_cgo_runtime_cgocall(_wrap_in_schar, _swig_p)
	return
}
var _wrap_in_float unsafe.Pointer

func In_float(arg1 float32) (_swig_ret float32) {
	_swig_p := uintptr(unsafe.Pointer(&arg1))
	_cgo_runtime_cgocall(_wrap_in_float, _swig_p)
	return
}
var _wrap_in_double unsafe.Pointer

func In_double(arg1 float64) (_swig_ret float64) {
	_swig_p := uintptr(unsafe.Pointer(&arg1))
	_cgo_runtime_cgocall(_wrap_in_double, _swig_p)
	return
}
var _wrap_in_longlong unsafe.Pointer

func In_longlong(arg1 int64) (_swig_ret int64) {
	_swig_p := uintptr(unsafe.Pointer(&arg1))
	_cgo_runtime_cgocall(_wrap_in_longlong, _swig_p)
	return
}
var _wrap_in_ulonglong unsafe.Pointer

func In_ulonglong(arg1 uint64) (_swig_ret uint64) {
	_swig_p := uintptr(unsafe.Pointer(&arg1))
	_cgo_runtime_cgocall(_wrap_in_ulonglong, _swig_p)
	return
}
var _wrap_inr_bool unsafe.Pointer

func Inr_bool(arg1 bool) (_swig_ret bool) {
	_swig_p := uintptr(unsafe.Pointer(&arg1))
	_cgo_runtime_cgocall(_wrap_inr_bool, _swig_p)
	return
}
var _wrap_inr_int unsafe.Pointer

func Inr_int(arg1 int) (_swig_ret int) {
	_swig_p := uintptr(unsafe.Pointer(&arg1))
	_cgo_runtime_cgocall(_wrap_inr_int, _swig_p)
	return
}
var _wrap_inr_long unsafe.Pointer

func Inr_long(arg1 int64) (_swig_ret int64) {
	_swig_p := uintptr(unsafe.Pointer(&arg1))
	_cgo_runtime_cgocall(_wrap_inr_long, _swig_p)
	return
}
var _wrap_inr_short unsafe.Pointer

func Inr_short(arg1 int16) (_swig_ret int16) {
	_swig_p := uintptr(unsafe.Pointer(&arg1))
	_cgo_runtime_cgocall(_wrap_inr_short, _swig_p)
	return
}
var _wrap_inr_uint unsafe.Pointer

func Inr_uint(arg1 uint) (_swig_ret uint) {
	_swig_p := uintptr(unsafe.Pointer(&arg1))
	_cgo_runtime_cgocall(_wrap_inr_uint, _swig_p)
	return
}
var _wrap_inr_ushort unsafe.Pointer

func Inr_ushort(arg1 uint16) (_swig_ret uint16) {
	_swig_p := uintptr(unsafe.Pointer(&arg1))
	_cgo_runtime_cgocall(_wrap_inr_ushort, _swig_p)
	return
}
var _wrap_inr_ulong unsafe.Pointer

func Inr_ulong(arg1 uint64) (_swig_ret uint64) {
	_swig_p := uintptr(unsafe.Pointer(&arg1))
	_cgo_runtime_cgocall(_wrap_inr_ulong, _swig_p)
	return
}
var _wrap_inr_uchar unsafe.Pointer

func Inr_uchar(arg1 byte) (_swig_ret byte) {
	_swig_p := uintptr(unsafe.Pointer(&arg1))
	_cgo_runtime_cgocall(_wrap_inr_uchar, _swig_p)
	return
}
var _wrap_inr_schar unsafe.Pointer

func Inr_schar(arg1 int8) (_swig_ret int8) {
	_swig_p := uintptr(unsafe.Pointer(&arg1))
	_cgo_runtime_cgocall(_wrap_inr_schar, _swig_p)
	return
}
var _wrap_inr_float unsafe.Pointer

func Inr_float(arg1 float32) (_swig_ret float32) {
	_swig_p := uintptr(unsafe.Pointer(&arg1))
	_cgo_runtime_cgocall(_wrap_inr_float, _swig_p)
	return
}
var _wrap_inr_double unsafe.Pointer

func Inr_double(arg1 float64) (_swig_ret float64) {
	_swig_p := uintptr(unsafe.Pointer(&arg1))
	_cgo_runtime_cgocall(_wrap_inr_double, _swig_p)
	return
}
var _wrap_inr_longlong unsafe.Pointer

func Inr_longlong(arg1 int64) (_swig_ret int64) {
	_swig_p := uintptr(unsafe.Pointer(&arg1))
	_cgo_runtime_cgocall(_wrap_inr_longlong, _swig_p)
	return
}
var _wrap_inr_ulonglong unsafe.Pointer

func Inr_ulonglong(arg1 uint64) (_swig_ret uint64) {
	_swig_p := uintptr(unsafe.Pointer(&arg1))
	_cgo_runtime_cgocall(_wrap_inr_ulonglong, _swig_p)
	return
}
var _wrap_out_bool unsafe.Pointer

func _swig_wrap_out_bool(base bool, _ []bool) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_out_bool, _swig_p)
	return
}

func Out_bool(arg1 bool, arg2 []bool) {
	_swig_wrap_out_bool(arg1, arg2)
}

var _wrap_out_int unsafe.Pointer

func _swig_wrap_out_int(base int, _ []int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_out_int, _swig_p)
	return
}

func Out_int(arg1 int, arg2 []int) {
	_swig_wrap_out_int(arg1, arg2)
}

var _wrap_out_short unsafe.Pointer

func _swig_wrap_out_short(base int16, _ []int16) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_out_short, _swig_p)
	return
}

func Out_short(arg1 int16, arg2 []int16) {
	_swig_wrap_out_short(arg1, arg2)
}

var _wrap_out_long unsafe.Pointer

func _swig_wrap_out_long(base int64, _ []int64) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_out_long, _swig_p)
	return
}

func Out_long(arg1 int64, arg2 []int64) {
	_swig_wrap_out_long(arg1, arg2)
}

var _wrap_out_uint unsafe.Pointer

func _swig_wrap_out_uint(base uint, _ []uint) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_out_uint, _swig_p)
	return
}

func Out_uint(arg1 uint, arg2 []uint) {
	_swig_wrap_out_uint(arg1, arg2)
}

var _wrap_out_ushort unsafe.Pointer

func _swig_wrap_out_ushort(base uint16, _ []uint16) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_out_ushort, _swig_p)
	return
}

func Out_ushort(arg1 uint16, arg2 []uint16) {
	_swig_wrap_out_ushort(arg1, arg2)
}

var _wrap_out_ulong unsafe.Pointer

func _swig_wrap_out_ulong(base uint64, _ []uint64) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_out_ulong, _swig_p)
	return
}

func Out_ulong(arg1 uint64, arg2 []uint64) {
	_swig_wrap_out_ulong(arg1, arg2)
}

var _wrap_out_uchar unsafe.Pointer

func _swig_wrap_out_uchar(base byte, _ []byte) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_out_uchar, _swig_p)
	return
}

func Out_uchar(arg1 byte, arg2 []byte) {
	_swig_wrap_out_uchar(arg1, arg2)
}

var _wrap_out_schar unsafe.Pointer

func _swig_wrap_out_schar(base int8, _ []int8) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_out_schar, _swig_p)
	return
}

func Out_schar(arg1 int8, arg2 []int8) {
	_swig_wrap_out_schar(arg1, arg2)
}

var _wrap_out_float unsafe.Pointer

func _swig_wrap_out_float(base float32, _ []float32) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_out_float, _swig_p)
	return
}

func Out_float(arg1 float32, arg2 []float32) {
	_swig_wrap_out_float(arg1, arg2)
}

var _wrap_out_double unsafe.Pointer

func _swig_wrap_out_double(base float64, _ []float64) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_out_double, _swig_p)
	return
}

func Out_double(arg1 float64, arg2 []float64) {
	_swig_wrap_out_double(arg1, arg2)
}

var _wrap_out_longlong unsafe.Pointer

func _swig_wrap_out_longlong(base int64, _ []int64) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_out_longlong, _swig_p)
	return
}

func Out_longlong(arg1 int64, arg2 []int64) {
	_swig_wrap_out_longlong(arg1, arg2)
}

var _wrap_out_ulonglong unsafe.Pointer

func _swig_wrap_out_ulonglong(base uint64, _ []uint64) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_out_ulonglong, _swig_p)
	return
}

func Out_ulonglong(arg1 uint64, arg2 []uint64) {
	_swig_wrap_out_ulonglong(arg1, arg2)
}

var _wrap_out_foo unsafe.Pointer

func _swig_wrap_out_foo(base int, _ []int) (_ SwigcptrFoo) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_out_foo, _swig_p)
	return
}

func Out_foo(arg1 int, arg2 []int) (_swig_ret Foo) {
	return _swig_wrap_out_foo(arg1, arg2)
}

var _wrap_outr_bool unsafe.Pointer

func _swig_wrap_outr_bool(base bool, _ []bool) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_outr_bool, _swig_p)
	return
}

func Outr_bool(arg1 bool, arg2 []bool) {
	_swig_wrap_outr_bool(arg1, arg2)
}

var _wrap_outr_int unsafe.Pointer

func _swig_wrap_outr_int(base int, _ []int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_outr_int, _swig_p)
	return
}

func Outr_int(arg1 int, arg2 []int) {
	_swig_wrap_outr_int(arg1, arg2)
}

var _wrap_outr_short unsafe.Pointer

func _swig_wrap_outr_short(base int16, _ []int16) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_outr_short, _swig_p)
	return
}

func Outr_short(arg1 int16, arg2 []int16) {
	_swig_wrap_outr_short(arg1, arg2)
}

var _wrap_outr_long unsafe.Pointer

func _swig_wrap_outr_long(base int64, _ []int64) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_outr_long, _swig_p)
	return
}

func Outr_long(arg1 int64, arg2 []int64) {
	_swig_wrap_outr_long(arg1, arg2)
}

var _wrap_outr_uint unsafe.Pointer

func _swig_wrap_outr_uint(base uint, _ []uint) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_outr_uint, _swig_p)
	return
}

func Outr_uint(arg1 uint, arg2 []uint) {
	_swig_wrap_outr_uint(arg1, arg2)
}

var _wrap_outr_ushort unsafe.Pointer

func _swig_wrap_outr_ushort(base uint16, _ []uint16) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_outr_ushort, _swig_p)
	return
}

func Outr_ushort(arg1 uint16, arg2 []uint16) {
	_swig_wrap_outr_ushort(arg1, arg2)
}

var _wrap_outr_ulong unsafe.Pointer

func _swig_wrap_outr_ulong(base uint64, _ []uint64) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_outr_ulong, _swig_p)
	return
}

func Outr_ulong(arg1 uint64, arg2 []uint64) {
	_swig_wrap_outr_ulong(arg1, arg2)
}

var _wrap_outr_uchar unsafe.Pointer

func _swig_wrap_outr_uchar(base byte, _ []byte) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_outr_uchar, _swig_p)
	return
}

func Outr_uchar(arg1 byte, arg2 []byte) {
	_swig_wrap_outr_uchar(arg1, arg2)
}

var _wrap_outr_schar unsafe.Pointer

func _swig_wrap_outr_schar(base int8, _ []int8) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_outr_schar, _swig_p)
	return
}

func Outr_schar(arg1 int8, arg2 []int8) {
	_swig_wrap_outr_schar(arg1, arg2)
}

var _wrap_outr_float unsafe.Pointer

func _swig_wrap_outr_float(base float32, _ []float32) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_outr_float, _swig_p)
	return
}

func Outr_float(arg1 float32, arg2 []float32) {
	_swig_wrap_outr_float(arg1, arg2)
}

var _wrap_outr_double unsafe.Pointer

func _swig_wrap_outr_double(base float64, _ []float64) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_outr_double, _swig_p)
	return
}

func Outr_double(arg1 float64, arg2 []float64) {
	_swig_wrap_outr_double(arg1, arg2)
}

var _wrap_outr_longlong unsafe.Pointer

func _swig_wrap_outr_longlong(base int64, _ []int64) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_outr_longlong, _swig_p)
	return
}

func Outr_longlong(arg1 int64, arg2 []int64) {
	_swig_wrap_outr_longlong(arg1, arg2)
}

var _wrap_outr_ulonglong unsafe.Pointer

func _swig_wrap_outr_ulonglong(base uint64, _ []uint64) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_outr_ulonglong, _swig_p)
	return
}

func Outr_ulonglong(arg1 uint64, arg2 []uint64) {
	_swig_wrap_outr_ulonglong(arg1, arg2)
}

var _wrap_inout_bool unsafe.Pointer

func _swig_wrap_inout_bool(base []bool) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_inout_bool, _swig_p)
	return
}

func Inout_bool(arg1 []bool) {
	_swig_wrap_inout_bool(arg1)
}

var _wrap_inout_int unsafe.Pointer

func _swig_wrap_inout_int(base []int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_inout_int, _swig_p)
	return
}

func Inout_int(arg1 []int) {
	_swig_wrap_inout_int(arg1)
}

var _wrap_inout_short unsafe.Pointer

func _swig_wrap_inout_short(base []int16) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_inout_short, _swig_p)
	return
}

func Inout_short(arg1 []int16) {
	_swig_wrap_inout_short(arg1)
}

var _wrap_inout_long unsafe.Pointer

func _swig_wrap_inout_long(base []int64) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_inout_long, _swig_p)
	return
}

func Inout_long(arg1 []int64) {
	_swig_wrap_inout_long(arg1)
}

var _wrap_inout_uint unsafe.Pointer

func _swig_wrap_inout_uint(base []uint) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_inout_uint, _swig_p)
	return
}

func Inout_uint(arg1 []uint) {
	_swig_wrap_inout_uint(arg1)
}

var _wrap_inout_ushort unsafe.Pointer

func _swig_wrap_inout_ushort(base []uint16) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_inout_ushort, _swig_p)
	return
}

func Inout_ushort(arg1 []uint16) {
	_swig_wrap_inout_ushort(arg1)
}

var _wrap_inout_ulong unsafe.Pointer

func _swig_wrap_inout_ulong(base []uint64) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_inout_ulong, _swig_p)
	return
}

func Inout_ulong(arg1 []uint64) {
	_swig_wrap_inout_ulong(arg1)
}

var _wrap_inout_uchar unsafe.Pointer

func _swig_wrap_inout_uchar(base []byte) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_inout_uchar, _swig_p)
	return
}

func Inout_uchar(arg1 []byte) {
	_swig_wrap_inout_uchar(arg1)
}

var _wrap_inout_schar unsafe.Pointer

func _swig_wrap_inout_schar(base []int8) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_inout_schar, _swig_p)
	return
}

func Inout_schar(arg1 []int8) {
	_swig_wrap_inout_schar(arg1)
}

var _wrap_inout_float unsafe.Pointer

func _swig_wrap_inout_float(base []float32) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_inout_float, _swig_p)
	return
}

func Inout_float(arg1 []float32) {
	_swig_wrap_inout_float(arg1)
}

var _wrap_inout_double unsafe.Pointer

func _swig_wrap_inout_double(base []float64) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_inout_double, _swig_p)
	return
}

func Inout_double(arg1 []float64) {
	_swig_wrap_inout_double(arg1)
}

var _wrap_inout_longlong unsafe.Pointer

func _swig_wrap_inout_longlong(base []int64) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_inout_longlong, _swig_p)
	return
}

func Inout_longlong(arg1 []int64) {
	_swig_wrap_inout_longlong(arg1)
}

var _wrap_inout_ulonglong unsafe.Pointer

func _swig_wrap_inout_ulonglong(base []uint64) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_inout_ulonglong, _swig_p)
	return
}

func Inout_ulonglong(arg1 []uint64) {
	_swig_wrap_inout_ulonglong(arg1)
}

var _wrap_inoutr_bool unsafe.Pointer

func _swig_wrap_inoutr_bool(base []bool) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_inoutr_bool, _swig_p)
	return
}

func Inoutr_bool(arg1 []bool) {
	_swig_wrap_inoutr_bool(arg1)
}

var _wrap_inoutr_int unsafe.Pointer

func _swig_wrap_inoutr_int(base []int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_inoutr_int, _swig_p)
	return
}

func Inoutr_int(arg1 []int) {
	_swig_wrap_inoutr_int(arg1)
}

var _wrap_inoutr_short unsafe.Pointer

func _swig_wrap_inoutr_short(base []int16) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_inoutr_short, _swig_p)
	return
}

func Inoutr_short(arg1 []int16) {
	_swig_wrap_inoutr_short(arg1)
}

var _wrap_inoutr_long unsafe.Pointer

func _swig_wrap_inoutr_long(base []int64) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_inoutr_long, _swig_p)
	return
}

func Inoutr_long(arg1 []int64) {
	_swig_wrap_inoutr_long(arg1)
}

var _wrap_inoutr_uint unsafe.Pointer

func _swig_wrap_inoutr_uint(base []uint) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_inoutr_uint, _swig_p)
	return
}

func Inoutr_uint(arg1 []uint) {
	_swig_wrap_inoutr_uint(arg1)
}

var _wrap_inoutr_ushort unsafe.Pointer

func _swig_wrap_inoutr_ushort(base []uint16) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_inoutr_ushort, _swig_p)
	return
}

func Inoutr_ushort(arg1 []uint16) {
	_swig_wrap_inoutr_ushort(arg1)
}

var _wrap_inoutr_ulong unsafe.Pointer

func _swig_wrap_inoutr_ulong(base []uint64) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_inoutr_ulong, _swig_p)
	return
}

func Inoutr_ulong(arg1 []uint64) {
	_swig_wrap_inoutr_ulong(arg1)
}

var _wrap_inoutr_uchar unsafe.Pointer

func _swig_wrap_inoutr_uchar(base []byte) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_inoutr_uchar, _swig_p)
	return
}

func Inoutr_uchar(arg1 []byte) {
	_swig_wrap_inoutr_uchar(arg1)
}

var _wrap_inoutr_schar unsafe.Pointer

func _swig_wrap_inoutr_schar(base []int8) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_inoutr_schar, _swig_p)
	return
}

func Inoutr_schar(arg1 []int8) {
	_swig_wrap_inoutr_schar(arg1)
}

var _wrap_inoutr_float unsafe.Pointer

func _swig_wrap_inoutr_float(base []float32) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_inoutr_float, _swig_p)
	return
}

func Inoutr_float(arg1 []float32) {
	_swig_wrap_inoutr_float(arg1)
}

var _wrap_inoutr_double unsafe.Pointer

func _swig_wrap_inoutr_double(base []float64) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_inoutr_double, _swig_p)
	return
}

func Inoutr_double(arg1 []float64) {
	_swig_wrap_inoutr_double(arg1)
}

var _wrap_inoutr_longlong unsafe.Pointer

func _swig_wrap_inoutr_longlong(base []int64) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_inoutr_longlong, _swig_p)
	return
}

func Inoutr_longlong(arg1 []int64) {
	_swig_wrap_inoutr_longlong(arg1)
}

var _wrap_inoutr_ulonglong unsafe.Pointer

func _swig_wrap_inoutr_ulonglong(base []uint64) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_inoutr_ulonglong, _swig_p)
	return
}

func Inoutr_ulonglong(arg1 []uint64) {
	_swig_wrap_inoutr_ulonglong(arg1)
}

var _wrap_inoutr_int2 unsafe.Pointer

func _swig_wrap_inoutr_int2(base []int, _ []int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_inoutr_int2, _swig_p)
	return
}

func Inoutr_int2(arg1 []int, arg2 []int) {
	_swig_wrap_inoutr_int2(arg1, arg2)
}


type SwigcptrVoid uintptr
type Void interface {
	Swigcptr() uintptr;
}
func (p SwigcptrVoid) Swigcptr() uintptr {
	return uintptr(p)
}

