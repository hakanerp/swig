/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 3.0.3
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: ./../namespace_template.i

package namespace_template

import "unsafe"
import _ "runtime/cgo"

var _cgo_runtime_cgocall func(unsafe.Pointer, uintptr)



type _ unsafe.Pointer



type _swig_fnptr *byte
type _swig_memberptr *byte

var _wrap_maxint unsafe.Pointer

func Maxint(arg1 int, arg2 int) (_swig_ret int) {
	_swig_p := uintptr(unsafe.Pointer(&arg1))
	_cgo_runtime_cgocall(_wrap_maxint, _swig_p)
	return
}
type SwigcptrVectorint uintptr

func (p SwigcptrVectorint) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrVectorint) SwigIsVectorint() {
}

var _wrap_new_vectorint unsafe.Pointer

func _swig_wrap_new_vectorint() (base SwigcptrVectorint) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_vectorint, _swig_p)
	return
}

func NewVectorint() (_swig_ret Vectorint) {
	return _swig_wrap_new_vectorint()
}

var _wrap_delete_vectorint unsafe.Pointer

func _swig_wrap_delete_vectorint(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_vectorint, _swig_p)
	return
}

func DeleteVectorint(arg1 Vectorint) {
	_swig_wrap_delete_vectorint(arg1.Swigcptr())
}

var _wrap_vectorint_blah unsafe.Pointer

func _swig_wrap_vectorint_blah(base SwigcptrVectorint, _ int) (_ string) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_vectorint_blah, _swig_p)
	return
}

func (arg1 SwigcptrVectorint) Blah(arg2 int) (_swig_ret string) {
	return _swig_wrap_vectorint_blah(arg1, arg2)
}

type Vectorint interface {
	Swigcptr() uintptr
	SwigIsVectorint()
	Blah(arg2 int) (_swig_ret string)
}

var _wrap_maxshort unsafe.Pointer

func Maxshort(arg1 int16, arg2 int16) (_swig_ret int16) {
	_swig_p := uintptr(unsafe.Pointer(&arg1))
	_cgo_runtime_cgocall(_wrap_maxshort, _swig_p)
	return
}
type SwigcptrVectorshort uintptr

func (p SwigcptrVectorshort) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrVectorshort) SwigIsVectorshort() {
}

var _wrap_new_vectorshort unsafe.Pointer

func _swig_wrap_new_vectorshort() (base SwigcptrVectorshort) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_vectorshort, _swig_p)
	return
}

func NewVectorshort() (_swig_ret Vectorshort) {
	return _swig_wrap_new_vectorshort()
}

var _wrap_delete_vectorshort unsafe.Pointer

func _swig_wrap_delete_vectorshort(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_vectorshort, _swig_p)
	return
}

func DeleteVectorshort(arg1 Vectorshort) {
	_swig_wrap_delete_vectorshort(arg1.Swigcptr())
}

var _wrap_vectorshort_blah unsafe.Pointer

func _swig_wrap_vectorshort_blah(base SwigcptrVectorshort, _ int16) (_ string) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_vectorshort_blah, _swig_p)
	return
}

func (arg1 SwigcptrVectorshort) Blah(arg2 int16) (_swig_ret string) {
	return _swig_wrap_vectorshort_blah(arg1, arg2)
}

type Vectorshort interface {
	Swigcptr() uintptr
	SwigIsVectorshort()
	Blah(arg2 int16) (_swig_ret string)
}

var _wrap_maxlong unsafe.Pointer

func Maxlong(arg1 int64, arg2 int64) (_swig_ret int64) {
	_swig_p := uintptr(unsafe.Pointer(&arg1))
	_cgo_runtime_cgocall(_wrap_maxlong, _swig_p)
	return
}
type SwigcptrVectorlong uintptr

func (p SwigcptrVectorlong) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrVectorlong) SwigIsVectorlong() {
}

var _wrap_new_vectorlong unsafe.Pointer

func _swig_wrap_new_vectorlong() (base SwigcptrVectorlong) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_vectorlong, _swig_p)
	return
}

func NewVectorlong() (_swig_ret Vectorlong) {
	return _swig_wrap_new_vectorlong()
}

var _wrap_delete_vectorlong unsafe.Pointer

func _swig_wrap_delete_vectorlong(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_vectorlong, _swig_p)
	return
}

func DeleteVectorlong(arg1 Vectorlong) {
	_swig_wrap_delete_vectorlong(arg1.Swigcptr())
}

var _wrap_vectorlong_blah unsafe.Pointer

func _swig_wrap_vectorlong_blah(base SwigcptrVectorlong, _ int64) (_ string) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_vectorlong_blah, _swig_p)
	return
}

func (arg1 SwigcptrVectorlong) Blah(arg2 int64) (_swig_ret string) {
	return _swig_wrap_vectorlong_blah(arg1, arg2)
}

type Vectorlong interface {
	Swigcptr() uintptr
	SwigIsVectorlong()
	Blah(arg2 int64) (_swig_ret string)
}

var _wrap_maxInteger unsafe.Pointer

func MaxInteger(arg1 int, arg2 int) (_swig_ret int) {
	_swig_p := uintptr(unsafe.Pointer(&arg1))
	_cgo_runtime_cgocall(_wrap_maxInteger, _swig_p)
	return
}
type SwigcptrVectorInteger uintptr

func (p SwigcptrVectorInteger) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrVectorInteger) SwigIsVectorInteger() {
}

var _wrap_new_vectorInteger unsafe.Pointer

func _swig_wrap_new_vectorInteger() (base SwigcptrVectorInteger) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_vectorInteger, _swig_p)
	return
}

func NewVectorInteger() (_swig_ret VectorInteger) {
	return _swig_wrap_new_vectorInteger()
}

var _wrap_delete_vectorInteger unsafe.Pointer

func _swig_wrap_delete_vectorInteger(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_vectorInteger, _swig_p)
	return
}

func DeleteVectorInteger(arg1 VectorInteger) {
	_swig_wrap_delete_vectorInteger(arg1.Swigcptr())
}

var _wrap_vectorInteger_blah unsafe.Pointer

func _swig_wrap_vectorInteger_blah(base SwigcptrVectorInteger, _ int) (_ string) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_vectorInteger_blah, _swig_p)
	return
}

func (arg1 SwigcptrVectorInteger) Blah(arg2 int) (_swig_ret string) {
	return _swig_wrap_vectorInteger_blah(arg1, arg2)
}

type VectorInteger interface {
	Swigcptr() uintptr
	SwigIsVectorInteger()
	Blah(arg2 int) (_swig_ret string)
}


type SwigcptrVoid uintptr
type Void interface {
	Swigcptr() uintptr;
}
func (p SwigcptrVoid) Swigcptr() uintptr {
	return uintptr(p)
}

