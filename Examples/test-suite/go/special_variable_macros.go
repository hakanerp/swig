/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 3.0.3
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: ./../special_variable_macros.i

package special_variable_macros

import "unsafe"
import _ "runtime/cgo"

var _cgo_runtime_cgocall func(unsafe.Pointer, uintptr)



type _ unsafe.Pointer



type _swig_fnptr *byte
type _swig_memberptr *byte

type SwigcptrName uintptr

func (p SwigcptrName) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrName) SwigIsName() {
}

var _wrap_new_Name__SWIG_0 unsafe.Pointer

func _swig_wrap_new_Name__SWIG_0(base string) (_ SwigcptrName) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_Name__SWIG_0, _swig_p)
	return
}

func NewName__SWIG_0(arg1 string) (_swig_ret Name) {
	return _swig_wrap_new_Name__SWIG_0(arg1)
}

var _wrap_new_Name__SWIG_1 unsafe.Pointer

func _swig_wrap_new_Name__SWIG_1() (base SwigcptrName) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_Name__SWIG_1, _swig_p)
	return
}

func NewName__SWIG_1() (_swig_ret Name) {
	return _swig_wrap_new_Name__SWIG_1()
}

var _wrap_new_Name__SWIG_2 unsafe.Pointer

func _swig_wrap_new_Name__SWIG_2(base uintptr) (_ SwigcptrName) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_Name__SWIG_2, _swig_p)
	return
}

func NewName__SWIG_2(arg1 Name) (_swig_ret Name) {
	return _swig_wrap_new_Name__SWIG_2(arg1.Swigcptr())
}

func NewName(a ...interface{}) Name {
	argc := len(a)
	if argc == 0 {
		return NewName__SWIG_1()
	}
	if argc == 1 {
		if _, ok := a[0].(Name); !ok {
			goto check_2
		}
		return NewName__SWIG_2(a[0].(Name))
	}
check_2:
	if argc == 1 {
		return NewName__SWIG_0(a[0].(string))
	}
	panic("No match for overloaded function call")
}

var _wrap_delete_Name unsafe.Pointer

func _swig_wrap_delete_Name(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_Name, _swig_p)
	return
}

func DeleteName(arg1 Name) {
	_swig_wrap_delete_Name(arg1.Swigcptr())
}

var _wrap_Name_getName unsafe.Pointer

func _swig_wrap_Name_getName(base SwigcptrName) (_ string) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_Name_getName, _swig_p)
	return
}

func (arg1 SwigcptrName) GetName() (_swig_ret string) {
	return _swig_wrap_Name_getName(arg1)
}

var _wrap_Name_getNamePtr unsafe.Pointer

func _swig_wrap_Name_getNamePtr(base SwigcptrName) (_ SwigcptrName) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_Name_getNamePtr, _swig_p)
	return
}

func (arg1 SwigcptrName) GetNamePtr() (_swig_ret Name) {
	return _swig_wrap_Name_getNamePtr(arg1)
}

type Name interface {
	Swigcptr() uintptr
	SwigIsName()
	GetName() (_swig_ret string)
	GetNamePtr() (_swig_ret Name)
}

type SwigcptrNameWrap uintptr

func (p SwigcptrNameWrap) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrNameWrap) SwigIsNameWrap() {
}

var _wrap_new_NameWrap__SWIG_0 unsafe.Pointer

func _swig_wrap_new_NameWrap__SWIG_0(base string) (_ SwigcptrNameWrap) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_NameWrap__SWIG_0, _swig_p)
	return
}

func NewNameWrap__SWIG_0(arg1 string) (_swig_ret NameWrap) {
	return _swig_wrap_new_NameWrap__SWIG_0(arg1)
}

var _wrap_new_NameWrap__SWIG_1 unsafe.Pointer

func _swig_wrap_new_NameWrap__SWIG_1() (base SwigcptrNameWrap) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_NameWrap__SWIG_1, _swig_p)
	return
}

func NewNameWrap__SWIG_1() (_swig_ret NameWrap) {
	return _swig_wrap_new_NameWrap__SWIG_1()
}

func NewNameWrap(a ...interface{}) NameWrap {
	argc := len(a)
	if argc == 0 {
		return NewNameWrap__SWIG_1()
	}
	if argc == 1 {
		return NewNameWrap__SWIG_0(a[0].(string))
	}
	panic("No match for overloaded function call")
}

var _wrap_NameWrap_getNamePtr unsafe.Pointer

func _swig_wrap_NameWrap_getNamePtr(base SwigcptrNameWrap) (_ SwigcptrName) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_NameWrap_getNamePtr, _swig_p)
	return
}

func (arg1 SwigcptrNameWrap) GetNamePtr() (_swig_ret Name) {
	return _swig_wrap_NameWrap_getNamePtr(arg1)
}

var _wrap_delete_NameWrap unsafe.Pointer

func _swig_wrap_delete_NameWrap(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_NameWrap, _swig_p)
	return
}

func DeleteNameWrap(arg1 NameWrap) {
	_swig_wrap_delete_NameWrap(arg1.Swigcptr())
}

type NameWrap interface {
	Swigcptr() uintptr
	SwigIsNameWrap()
	GetNamePtr() (_swig_ret Name)
}

var _wrap_testFred unsafe.Pointer

func _swig_wrap_testFred(base uintptr) (_ string) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_testFred, _swig_p)
	return
}

func TestFred(arg1 Name) (_swig_ret string) {
	return _swig_wrap_testFred(arg1.Swigcptr())
}

var _wrap_testJack unsafe.Pointer

func _swig_wrap_testJack(base uintptr) (_ string) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_testJack, _swig_p)
	return
}

func TestJack(arg1 Name) (_swig_ret string) {
	return _swig_wrap_testJack(arg1.Swigcptr())
}

var _wrap_testJill unsafe.Pointer

func _swig_wrap_testJill(base uintptr) (_ string) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_testJill, _swig_p)
	return
}

func TestJill(arg1 Name) (_swig_ret string) {
	return _swig_wrap_testJill(arg1.Swigcptr())
}

var _wrap_testMary unsafe.Pointer

func _swig_wrap_testMary(base uintptr) (_ string) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_testMary, _swig_p)
	return
}

func TestMary(arg1 Name) (_swig_ret string) {
	return _swig_wrap_testMary(arg1.Swigcptr())
}

var _wrap_testJames unsafe.Pointer

func _swig_wrap_testJames(base uintptr) (_ string) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_testJames, _swig_p)
	return
}

func TestJames(arg1 Name) (_swig_ret string) {
	return _swig_wrap_testJames(arg1.Swigcptr())
}

var _wrap_testJim unsafe.Pointer

func _swig_wrap_testJim(base uintptr) (_ string) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_testJim, _swig_p)
	return
}

func TestJim(arg1 Name) (_swig_ret string) {
	return _swig_wrap_testJim(arg1.Swigcptr())
}

var _wrap_testJohn unsafe.Pointer

func _swig_wrap_testJohn(base uintptr) (_ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_testJohn, _swig_p)
	return
}

func TestJohn(arg1 PairIntBool) (_swig_ret int) {
	return _swig_wrap_testJohn(arg1.Swigcptr())
}

type SwigcptrPairIntBool uintptr

func (p SwigcptrPairIntBool) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrPairIntBool) SwigIsPairIntBool() {
}

var _wrap_new_PairIntBool__SWIG_0 unsafe.Pointer

func _swig_wrap_new_PairIntBool__SWIG_0(base int, _ bool) (_ SwigcptrPairIntBool) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_PairIntBool__SWIG_0, _swig_p)
	return
}

func NewPairIntBool__SWIG_0(arg1 int, arg2 bool) (_swig_ret PairIntBool) {
	return _swig_wrap_new_PairIntBool__SWIG_0(arg1, arg2)
}

var _wrap_new_PairIntBool__SWIG_1 unsafe.Pointer

func _swig_wrap_new_PairIntBool__SWIG_1() (base SwigcptrPairIntBool) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_PairIntBool__SWIG_1, _swig_p)
	return
}

func NewPairIntBool__SWIG_1() (_swig_ret PairIntBool) {
	return _swig_wrap_new_PairIntBool__SWIG_1()
}

func NewPairIntBool(a ...interface{}) PairIntBool {
	argc := len(a)
	if argc == 0 {
		return NewPairIntBool__SWIG_1()
	}
	if argc == 2 {
		return NewPairIntBool__SWIG_0(a[0].(int), a[1].(bool))
	}
	panic("No match for overloaded function call")
}

var _wrap_PairIntBool_first_set unsafe.Pointer

func _swig_wrap_PairIntBool_first_set(base SwigcptrPairIntBool, _ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_PairIntBool_first_set, _swig_p)
	return
}

func (arg1 SwigcptrPairIntBool) SetFirst(arg2 int) {
	_swig_wrap_PairIntBool_first_set(arg1, arg2)
}

var _wrap_PairIntBool_first_get unsafe.Pointer

func _swig_wrap_PairIntBool_first_get(base SwigcptrPairIntBool) (_ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_PairIntBool_first_get, _swig_p)
	return
}

func (arg1 SwigcptrPairIntBool) GetFirst() (_swig_ret int) {
	return _swig_wrap_PairIntBool_first_get(arg1)
}

var _wrap_PairIntBool_second_set unsafe.Pointer

func _swig_wrap_PairIntBool_second_set(base SwigcptrPairIntBool, _ bool) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_PairIntBool_second_set, _swig_p)
	return
}

func (arg1 SwigcptrPairIntBool) SetSecond(arg2 bool) {
	_swig_wrap_PairIntBool_second_set(arg1, arg2)
}

var _wrap_PairIntBool_second_get unsafe.Pointer

func _swig_wrap_PairIntBool_second_get(base SwigcptrPairIntBool) (_ bool) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_PairIntBool_second_get, _swig_p)
	return
}

func (arg1 SwigcptrPairIntBool) GetSecond() (_swig_ret bool) {
	return _swig_wrap_PairIntBool_second_get(arg1)
}

var _wrap_delete_PairIntBool unsafe.Pointer

func _swig_wrap_delete_PairIntBool(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_PairIntBool, _swig_p)
	return
}

func DeletePairIntBool(arg1 PairIntBool) {
	_swig_wrap_delete_PairIntBool(arg1.Swigcptr())
}

type PairIntBool interface {
	Swigcptr() uintptr
	SwigIsPairIntBool()
	SetFirst(arg2 int)
	GetFirst() (_swig_ret int)
	SetSecond(arg2 bool)
	GetSecond() (_swig_ret bool)
}


type SwigcptrVoid uintptr
type Void interface {
	Swigcptr() uintptr;
}
func (p SwigcptrVoid) Swigcptr() uintptr {
	return uintptr(p)
}

