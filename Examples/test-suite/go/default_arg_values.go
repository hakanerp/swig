/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 3.0.3
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: ./../default_arg_values.i

package default_arg_values

import "unsafe"
import _ "runtime/cgo"

var _cgo_runtime_cgocall func(unsafe.Pointer, uintptr)



type _ unsafe.Pointer



type _swig_fnptr *byte
type _swig_memberptr *byte

type SwigcptrDisplay uintptr

func (p SwigcptrDisplay) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrDisplay) SwigIsDisplay() {
}

var _wrap_Display_draw1__SWIG_0 unsafe.Pointer

func _swig_wrap_Display_draw1__SWIG_0(base SwigcptrDisplay, _ float32) (_ float32) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_Display_draw1__SWIG_0, _swig_p)
	return
}

func (arg1 SwigcptrDisplay) Draw1__SWIG_0(arg2 float32) (_swig_ret float32) {
	return _swig_wrap_Display_draw1__SWIG_0(arg1, arg2)
}

var _wrap_Display_draw1__SWIG_1 unsafe.Pointer

func _swig_wrap_Display_draw1__SWIG_1(base SwigcptrDisplay) (_ float32) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_Display_draw1__SWIG_1, _swig_p)
	return
}

func (arg1 SwigcptrDisplay) Draw1__SWIG_1() (_swig_ret float32) {
	return _swig_wrap_Display_draw1__SWIG_1(arg1)
}

func (p SwigcptrDisplay) Draw1(a ...interface{}) float32 {
	argc := len(a)
	if argc == 0 {
		return p.Draw1__SWIG_1()
	}
	if argc == 1 {
		return p.Draw1__SWIG_0(a[0].(float32))
	}
	panic("No match for overloaded function call")
}

var _wrap_Display_draw2__SWIG_0 unsafe.Pointer

func _swig_wrap_Display_draw2__SWIG_0(base SwigcptrDisplay, _ *float32) (_ float32) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_Display_draw2__SWIG_0, _swig_p)
	return
}

func (arg1 SwigcptrDisplay) Draw2__SWIG_0(arg2 *float32) (_swig_ret float32) {
	return _swig_wrap_Display_draw2__SWIG_0(arg1, arg2)
}

var _wrap_Display_draw2__SWIG_1 unsafe.Pointer

func _swig_wrap_Display_draw2__SWIG_1(base SwigcptrDisplay) (_ float32) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_Display_draw2__SWIG_1, _swig_p)
	return
}

func (arg1 SwigcptrDisplay) Draw2__SWIG_1() (_swig_ret float32) {
	return _swig_wrap_Display_draw2__SWIG_1(arg1)
}

func (p SwigcptrDisplay) Draw2(a ...interface{}) float32 {
	argc := len(a)
	if argc == 0 {
		return p.Draw2__SWIG_1()
	}
	if argc == 1 {
		return p.Draw2__SWIG_0(a[0].(*float32))
	}
	panic("No match for overloaded function call")
}

var _wrap_new_Display unsafe.Pointer

func _swig_wrap_new_Display() (base SwigcptrDisplay) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_Display, _swig_p)
	return
}

func NewDisplay() (_swig_ret Display) {
	return _swig_wrap_new_Display()
}

var _wrap_delete_Display unsafe.Pointer

func _swig_wrap_delete_Display(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_Display, _swig_p)
	return
}

func DeleteDisplay(arg1 Display) {
	_swig_wrap_delete_Display(arg1.Swigcptr())
}

type Display interface {
	Swigcptr() uintptr
	SwigIsDisplay()
	Draw1(a ...interface{}) float32
	Draw2(a ...interface{}) float32
}

var _wrap_createPtr unsafe.Pointer

func CreatePtr(arg1 float32) (_swig_ret *float32) {
	_swig_p := uintptr(unsafe.Pointer(&arg1))
	_cgo_runtime_cgocall(_wrap_createPtr, _swig_p)
	return
}

type SwigcptrVoid uintptr
type Void interface {
	Swigcptr() uintptr;
}
func (p SwigcptrVoid) Swigcptr() uintptr {
	return uintptr(p)
}

