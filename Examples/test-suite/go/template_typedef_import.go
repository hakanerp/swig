/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 3.0.3
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: ./../template_typedef_import.i

package template_typedef_import

import "unsafe"
import _ "runtime/cgo"
import "template_typedef_cplx2"

var _cgo_runtime_cgocall func(unsafe.Pointer, uintptr)



type _ unsafe.Pointer



type _swig_fnptr *byte
type _swig_memberptr *byte

var _wrap_my_func_r unsafe.Pointer

func _swig_wrap_my_func_r(base uintptr) (_ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_my_func_r, _swig_p)
	return
}

func My_func_r(arg1 template_typedef_cplx2.ArithUnaryFunction_double_double) (_swig_ret int) {
	return _swig_wrap_my_func_r(arg1.Swigcptr())
}

var _wrap_my_func_c unsafe.Pointer

func _swig_wrap_my_func_c(base uintptr) (_ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_my_func_c, _swig_p)
	return
}

func My_func_c(arg1 template_typedef_cplx2.ArithUnaryFunction_complex_complex) (_swig_ret int) {
	return _swig_wrap_my_func_c(arg1.Swigcptr())
}

type SwigcptrSin uintptr

func (p SwigcptrSin) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrSin) SwigIsSin() {
}

var _wrap_new_Sin unsafe.Pointer

func _swig_wrap_new_Sin() (base SwigcptrSin) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_Sin, _swig_p)
	return
}

func NewSin() (_swig_ret Sin) {
	return _swig_wrap_new_Sin()
}

var _wrap_delete_Sin unsafe.Pointer

func _swig_wrap_delete_Sin(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_Sin, _swig_p)
	return
}

func DeleteSin(arg1 Sin) {
	_swig_wrap_delete_Sin(arg1.Swigcptr())
}

var _wrap_Sin_get_arith_value unsafe.Pointer

func _swig_wrap_Sin_get_arith_value(base SwigcptrSin) (_ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_Sin_get_arith_value, _swig_p)
	return
}

func (_swig_base SwigcptrSin) Get_arith_value() (_swig_ret int) {
	return _swig_wrap_Sin_get_arith_value(_swig_base)
}

var _wrap_Sin_get_value unsafe.Pointer

func _swig_wrap_Sin_get_value(base SwigcptrSin) (_ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_Sin_get_value, _swig_p)
	return
}

func (_swig_base SwigcptrSin) Get_value() (_swig_ret int) {
	return _swig_wrap_Sin_get_value(_swig_base)
}

var _wrap_Sin_get_base_value unsafe.Pointer

func _swig_wrap_Sin_get_base_value(base SwigcptrSin) (_ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_Sin_get_base_value, _swig_p)
	return
}

func (_swig_base SwigcptrSin) Get_base_value() (_swig_ret int) {
	return _swig_wrap_Sin_get_base_value(_swig_base)
}

func (p SwigcptrSin) SwigIsArithUnaryFunction_double_double() {
}

func (p SwigcptrSin) SwigGetArithUnaryFunction_double_double() template_typedef_cplx2.ArithUnaryFunction_double_double {
	return template_typedef_cplx2.SwigcptrArithUnaryFunction_double_double(p.Swigcptr())
}

func (p SwigcptrSin) SwigIsUnaryFunction_double_double() {
}

func (p SwigcptrSin) SwigGetUnaryFunction_double_double() template_typedef_cplx2.UnaryFunction_double_double {
	return template_typedef_cplx2.SwigcptrUnaryFunction_double_double(p.Swigcptr())
}

func (p SwigcptrSin) SwigIsUnaryFunctionBase() {
}

func (p SwigcptrSin) SwigGetUnaryFunctionBase() template_typedef_cplx2.UnaryFunctionBase {
	return template_typedef_cplx2.SwigcptrUnaryFunctionBase(p.Swigcptr())
}

type Sin interface {
	Swigcptr() uintptr
	SwigIsSin()
	Get_arith_value() (_swig_ret int)
	Get_value() (_swig_ret int)
	Get_base_value() (_swig_ret int)
	SwigIsArithUnaryFunction_double_double()
	SwigGetArithUnaryFunction_double_double() template_typedef_cplx2.ArithUnaryFunction_double_double
	SwigIsUnaryFunction_double_double()
	SwigGetUnaryFunction_double_double() template_typedef_cplx2.UnaryFunction_double_double
	SwigIsUnaryFunctionBase()
	SwigGetUnaryFunctionBase() template_typedef_cplx2.UnaryFunctionBase
}

type SwigcptrCSin uintptr

func (p SwigcptrCSin) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrCSin) SwigIsCSin() {
}

var _wrap_new_CSin unsafe.Pointer

func _swig_wrap_new_CSin() (base SwigcptrCSin) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_CSin, _swig_p)
	return
}

func NewCSin() (_swig_ret CSin) {
	return _swig_wrap_new_CSin()
}

var _wrap_delete_CSin unsafe.Pointer

func _swig_wrap_delete_CSin(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_CSin, _swig_p)
	return
}

func DeleteCSin(arg1 CSin) {
	_swig_wrap_delete_CSin(arg1.Swigcptr())
}

var _wrap_CSin_get_arith_value unsafe.Pointer

func _swig_wrap_CSin_get_arith_value(base SwigcptrCSin) (_ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_CSin_get_arith_value, _swig_p)
	return
}

func (_swig_base SwigcptrCSin) Get_arith_value() (_swig_ret int) {
	return _swig_wrap_CSin_get_arith_value(_swig_base)
}

var _wrap_CSin_get_value unsafe.Pointer

func _swig_wrap_CSin_get_value(base SwigcptrCSin) (_ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_CSin_get_value, _swig_p)
	return
}

func (_swig_base SwigcptrCSin) Get_value() (_swig_ret int) {
	return _swig_wrap_CSin_get_value(_swig_base)
}

var _wrap_CSin_get_base_value unsafe.Pointer

func _swig_wrap_CSin_get_base_value(base SwigcptrCSin) (_ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_CSin_get_base_value, _swig_p)
	return
}

func (_swig_base SwigcptrCSin) Get_base_value() (_swig_ret int) {
	return _swig_wrap_CSin_get_base_value(_swig_base)
}

func (p SwigcptrCSin) SwigIsArithUnaryFunction_complex_complex() {
}

func (p SwigcptrCSin) SwigGetArithUnaryFunction_complex_complex() template_typedef_cplx2.ArithUnaryFunction_complex_complex {
	return template_typedef_cplx2.SwigcptrArithUnaryFunction_complex_complex(p.Swigcptr())
}

func (p SwigcptrCSin) SwigIsUnaryFunction_complex_complex() {
}

func (p SwigcptrCSin) SwigGetUnaryFunction_complex_complex() template_typedef_cplx2.UnaryFunction_complex_complex {
	return template_typedef_cplx2.SwigcptrUnaryFunction_complex_complex(p.Swigcptr())
}

func (p SwigcptrCSin) SwigIsUnaryFunctionBase() {
}

func (p SwigcptrCSin) SwigGetUnaryFunctionBase() template_typedef_cplx2.UnaryFunctionBase {
	return template_typedef_cplx2.SwigcptrUnaryFunctionBase(p.Swigcptr())
}

type CSin interface {
	Swigcptr() uintptr
	SwigIsCSin()
	Get_arith_value() (_swig_ret int)
	Get_value() (_swig_ret int)
	Get_base_value() (_swig_ret int)
	SwigIsArithUnaryFunction_complex_complex()
	SwigGetArithUnaryFunction_complex_complex() template_typedef_cplx2.ArithUnaryFunction_complex_complex
	SwigIsUnaryFunction_complex_complex()
	SwigGetUnaryFunction_complex_complex() template_typedef_cplx2.UnaryFunction_complex_complex
	SwigIsUnaryFunctionBase()
	SwigGetUnaryFunctionBase() template_typedef_cplx2.UnaryFunctionBase
}


type SwigcptrVoid uintptr
type Void interface {
	Swigcptr() uintptr;
}
func (p SwigcptrVoid) Swigcptr() uintptr {
	return uintptr(p)
}

