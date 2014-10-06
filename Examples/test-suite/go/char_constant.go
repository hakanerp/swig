/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 3.0.3
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: ./../char_constant.i

package char_constant

import "unsafe"
import _ "runtime/cgo"

var _cgo_runtime_cgocall func(unsafe.Pointer, uintptr)



type _ unsafe.Pointer



type _swig_fnptr *byte
type _swig_memberptr *byte

const CHAR_CONSTANT byte = 'x'
const STRING_CONSTANT string = "xyzzy"
var _wrap_ESC_CONST unsafe.Pointer

func _swig_getESC_CONST() (_swig_ret byte) {
	_swig_p := uintptr(unsafe.Pointer(&_swig_ret))
	_cgo_runtime_cgocall(_wrap_ESC_CONST, _swig_p)
	return
}
var ESC_CONST byte = _swig_getESC_CONST()
var _wrap_NULL_CONST unsafe.Pointer

func _swig_getNULL_CONST() (_swig_ret byte) {
	_swig_p := uintptr(unsafe.Pointer(&_swig_ret))
	_cgo_runtime_cgocall(_wrap_NULL_CONST, _swig_p)
	return
}
var NULL_CONST byte = _swig_getNULL_CONST()
var _wrap_SPECIALCHAR unsafe.Pointer

func _swig_getSPECIALCHAR() (_swig_ret byte) {
	_swig_p := uintptr(unsafe.Pointer(&_swig_ret))
	_cgo_runtime_cgocall(_wrap_SPECIALCHAR, _swig_p)
	return
}
var SPECIALCHAR byte = _swig_getSPECIALCHAR()
var _wrap_ia_get unsafe.Pointer

func GetIa() (_swig_ret int) {
	_swig_p := uintptr(unsafe.Pointer(&_swig_ret))
	_cgo_runtime_cgocall(_wrap_ia_get, _swig_p)
	return
}
var _wrap_ib_get unsafe.Pointer

func GetIb() (_swig_ret int) {
	_swig_p := uintptr(unsafe.Pointer(&_swig_ret))
	_cgo_runtime_cgocall(_wrap_ib_get, _swig_p)
	return
}

