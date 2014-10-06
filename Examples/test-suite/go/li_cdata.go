/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 3.0.3
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: ./../li_cdata.i

package li_cdata

import "unsafe"
import _ "runtime/cgo"

var _cgo_runtime_cgocall func(unsafe.Pointer, uintptr)



type _ unsafe.Pointer



type _swig_fnptr *byte
type _swig_memberptr *byte

var _wrap_cdata unsafe.Pointer

func _swig_wrap_cdata(base int, _ uintptr, _ int) (_  []byte ) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_cdata, _swig_p)
	return
}

func Cdata(arg1 uintptr, _swig_args ...interface{}) (_swig_ret  []byte ) {
	var arg2 int
	if len(_swig_args) > 0 {
		arg2 = _swig_args[0].(int)
	}
	return _swig_wrap_cdata(len(_swig_args), arg1, arg2)
}

var _wrap_memmove unsafe.Pointer

func _swig_wrap_memmove(base uintptr, _ string, _ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_memmove, _swig_p)
	return
}

func Memmove(arg1 uintptr, arg2 string, arg3 int) {
	_swig_wrap_memmove(arg1, arg2, arg3)
}

var _wrap_cdata_int unsafe.Pointer

func _swig_wrap_cdata_int(base int, _ *int, _ int) (_  []byte ) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_cdata_int, _swig_p)
	return
}

func Cdata_int(arg1 *int, _swig_args ...interface{}) (_swig_ret  []byte ) {
	var arg2 int
	if len(_swig_args) > 0 {
		arg2 = _swig_args[0].(int)
	}
	return _swig_wrap_cdata_int(len(_swig_args), arg1, arg2)
}

var _wrap_cdata_double unsafe.Pointer

func _swig_wrap_cdata_double(base int, _ *float64, _ int) (_  []byte ) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_cdata_double, _swig_p)
	return
}

func Cdata_double(arg1 *float64, _swig_args ...interface{}) (_swig_ret  []byte ) {
	var arg2 int
	if len(_swig_args) > 0 {
		arg2 = _swig_args[0].(int)
	}
	return _swig_wrap_cdata_double(len(_swig_args), arg1, arg2)
}

var _wrap_malloc unsafe.Pointer

func Malloc(arg1 int64) (_swig_ret uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&arg1))
	_cgo_runtime_cgocall(_wrap_malloc, _swig_p)
	return
}

type SwigcptrVoid uintptr
type Void interface {
	Swigcptr() uintptr;
}
func (p SwigcptrVoid) Swigcptr() uintptr {
	return uintptr(p)
}

