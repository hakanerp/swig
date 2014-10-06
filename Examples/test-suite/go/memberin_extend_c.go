/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 3.0.3
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: ./../memberin_extend_c.i

package memberin_extend_c

import "unsafe"
import _ "runtime/cgo"

var _cgo_runtime_cgocall func(unsafe.Pointer, uintptr)



type _ unsafe.Pointer



type _swig_fnptr *byte
type _swig_memberptr *byte

type SwigcptrPerson uintptr

func (p SwigcptrPerson) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrPerson) SwigIsPerson() {
}

var _wrap_Person_name_set unsafe.Pointer

func _swig_wrap_Person_name_set(base SwigcptrPerson, _ string) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_Person_name_set, _swig_p)
	return
}

func (arg1 SwigcptrPerson) SetName(arg2 string) {
	_swig_wrap_Person_name_set(arg1, arg2)
}

var _wrap_Person_name_get unsafe.Pointer

func _swig_wrap_Person_name_get(base SwigcptrPerson) (_ string) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_Person_name_get, _swig_p)
	return
}

func (arg1 SwigcptrPerson) GetName() (_swig_ret string) {
	return _swig_wrap_Person_name_get(arg1)
}

var _wrap_new_Person unsafe.Pointer

func _swig_wrap_new_Person() (base SwigcptrPerson) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_Person, _swig_p)
	return
}

func NewPerson() (_swig_ret Person) {
	return _swig_wrap_new_Person()
}

var _wrap_delete_Person unsafe.Pointer

func _swig_wrap_delete_Person(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_Person, _swig_p)
	return
}

func DeletePerson(arg1 Person) {
	_swig_wrap_delete_Person(arg1.Swigcptr())
}

type Person interface {
	Swigcptr() uintptr
	SwigIsPerson()
	SetName(arg2 string)
	GetName() (_swig_ret string)
}


type SwigcptrVoid uintptr
type Void interface {
	Swigcptr() uintptr;
}
func (p SwigcptrVoid) Swigcptr() uintptr {
	return uintptr(p)
}

