/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 3.0.3
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: ./../argout.i

package argout

import "unsafe"
import _ "runtime/cgo"

var _cgo_runtime_cgocall func(unsafe.Pointer, uintptr)



type _ unsafe.Pointer



type _swig_fnptr *byte
type _swig_memberptr *byte

var _wrap_new_intp unsafe.Pointer

func New_intp() (_swig_ret *int) {
	_swig_p := uintptr(unsafe.Pointer(&_swig_ret))
	_cgo_runtime_cgocall(_wrap_new_intp, _swig_p)
	return
}
var _wrap_copy_intp unsafe.Pointer

func Copy_intp(arg1 int) (_swig_ret *int) {
	_swig_p := uintptr(unsafe.Pointer(&arg1))
	_cgo_runtime_cgocall(_wrap_copy_intp, _swig_p)
	return
}
var _wrap_delete_intp unsafe.Pointer

func _swig_wrap_delete_intp(base *int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_intp, _swig_p)
	return
}

func Delete_intp(arg1 *int) {
	_swig_wrap_delete_intp(arg1)
}

var _wrap_intp_assign unsafe.Pointer

func _swig_wrap_intp_assign(base *int, _ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_intp_assign, _swig_p)
	return
}

func Intp_assign(arg1 *int, arg2 int) {
	_swig_wrap_intp_assign(arg1, arg2)
}

var _wrap_intp_value unsafe.Pointer

func Intp_value(arg1 *int) (_swig_ret int) {
	_swig_p := uintptr(unsafe.Pointer(&arg1))
	_cgo_runtime_cgocall(_wrap_intp_value, _swig_p)
	return
}
var _wrap_incp unsafe.Pointer

func Incp(arg1 *int) (_swig_ret int) {
	_swig_p := uintptr(unsafe.Pointer(&arg1))
	_cgo_runtime_cgocall(_wrap_incp, _swig_p)
	return
}
var _wrap_incr unsafe.Pointer

func Incr(arg1 *int) (_swig_ret int) {
	_swig_p := uintptr(unsafe.Pointer(&arg1))
	_cgo_runtime_cgocall(_wrap_incr, _swig_p)
	return
}
var _wrap_inctr unsafe.Pointer

func Inctr(arg1 *int) (_swig_ret int) {
	_swig_p := uintptr(unsafe.Pointer(&arg1))
	_cgo_runtime_cgocall(_wrap_inctr, _swig_p)
	return
}
var _wrap_voidhandle unsafe.Pointer

func _swig_wrap_voidhandle(base *uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_voidhandle, _swig_p)
	return
}

func Voidhandle(arg1 *uintptr) {
	_swig_wrap_voidhandle(arg1)
}

var _wrap_handle unsafe.Pointer

func Handle(arg1 uintptr) (_swig_ret string) {
	_swig_p := uintptr(unsafe.Pointer(&arg1))
	_cgo_runtime_cgocall(_wrap_handle, _swig_p)
	return
}

type SwigcptrVoid uintptr
type Void interface {
	Swigcptr() uintptr;
}
func (p SwigcptrVoid) Swigcptr() uintptr {
	return uintptr(p)
}

