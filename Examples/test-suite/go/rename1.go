/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 3.0.3
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: ./../rename1.i

package rename1

import "unsafe"
import _ "runtime/cgo"

var _cgo_runtime_cgocall func(unsafe.Pointer, uintptr)



type _ unsafe.Pointer



type _swig_fnptr *byte
type _swig_memberptr *byte

type SwigcptrKlass uintptr

func (p SwigcptrKlass) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrKlass) SwigIsKlass() {
}

var _wrap_new_Klass__SWIG_0 unsafe.Pointer

func _swig_wrap_new_Klass__SWIG_0(base int) (_ SwigcptrKlass) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_Klass__SWIG_0, _swig_p)
	return
}

func NewKlass__SWIG_0(arg1 int) (_swig_ret Klass) {
	return _swig_wrap_new_Klass__SWIG_0(arg1)
}

var _wrap_new_Klass__SWIG_1 unsafe.Pointer

func _swig_wrap_new_Klass__SWIG_1() (base SwigcptrKlass) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_Klass__SWIG_1, _swig_p)
	return
}

func NewKlass__SWIG_1() (_swig_ret Klass) {
	return _swig_wrap_new_Klass__SWIG_1()
}

func NewKlass(a ...interface{}) Klass {
	argc := len(a)
	if argc == 0 {
		return NewKlass__SWIG_1()
	}
	if argc == 1 {
		return NewKlass__SWIG_0(a[0].(int))
	}
	panic("No match for overloaded function call")
}

var _wrap_delete_Klass unsafe.Pointer

func _swig_wrap_delete_Klass(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_Klass, _swig_p)
	return
}

func DeleteKlass(arg1 Klass) {
	_swig_wrap_delete_Klass(arg1.Swigcptr())
}

type Klass interface {
	Swigcptr() uintptr
	SwigIsKlass()
}

type SwigcptrAnother uintptr

func (p SwigcptrAnother) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrAnother) SwigIsAnother() {
}

var _wrap_new_Another unsafe.Pointer

func _swig_wrap_new_Another() (base SwigcptrAnother) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_Another, _swig_p)
	return
}

func NewAnother() (_swig_ret Another) {
	return _swig_wrap_new_Another()
}

var _wrap_delete_Another unsafe.Pointer

func _swig_wrap_delete_Another(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_Another, _swig_p)
	return
}

func DeleteAnother(arg1 Another) {
	_swig_wrap_delete_Another(arg1.Swigcptr())
}

type Another interface {
	Swigcptr() uintptr
	SwigIsAnother()
}

type SpaceEnu int
var _wrap_En1 unsafe.Pointer

func _swig_getEn1() (_swig_ret SpaceEnu) {
	_swig_p := uintptr(unsafe.Pointer(&_swig_ret))
	_cgo_runtime_cgocall(_wrap_En1, _swig_p)
	return
}
var En1 SpaceEnu = _swig_getEn1()
var _wrap_En2 unsafe.Pointer

func _swig_getEn2() (_swig_ret SpaceEnu) {
	_swig_p := uintptr(unsafe.Pointer(&_swig_ret))
	_cgo_runtime_cgocall(_wrap_En2, _swig_p)
	return
}
var En2 SpaceEnu = _swig_getEn2()
var _wrap_En3 unsafe.Pointer

func _swig_getEn3() (_swig_ret SpaceEnu) {
	_swig_p := uintptr(unsafe.Pointer(&_swig_ret))
	_cgo_runtime_cgocall(_wrap_En3, _swig_p)
	return
}
var En3 SpaceEnu = _swig_getEn3()
type SwigcptrABC uintptr

func (p SwigcptrABC) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrABC) SwigIsABC() {
}

var _wrap_ABC_methodABC unsafe.Pointer

func _swig_wrap_ABC_methodABC(base SwigcptrABC, _ uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_ABC_methodABC, _swig_p)
	return
}

func (arg1 SwigcptrABC) MethodABC(arg2 ABC) {
	_swig_wrap_ABC_methodABC(arg1, arg2.Swigcptr())
}

var _wrap_ABC_methodKlass unsafe.Pointer

func _swig_wrap_ABC_methodKlass(base SwigcptrABC, _ uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_ABC_methodKlass, _swig_p)
	return
}

func (arg1 SwigcptrABC) MethodKlass(arg2 Klass) {
	_swig_wrap_ABC_methodKlass(arg1, arg2.Swigcptr())
}

var _wrap_ABC_opABC unsafe.Pointer

func _swig_wrap_ABC_opABC(base SwigcptrABC) (_ SwigcptrABC) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_ABC_opABC, _swig_p)
	return
}

func (arg1 SwigcptrABC) OpABC() (_swig_ret ABC) {
	return _swig_wrap_ABC_opABC(arg1)
}

var _wrap_ABC_opKlass unsafe.Pointer

func _swig_wrap_ABC_opKlass(base SwigcptrABC) (_ SwigcptrKlass) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_ABC_opKlass, _swig_p)
	return
}

func (arg1 SwigcptrABC) OpKlass() (_swig_ret Klass) {
	return _swig_wrap_ABC_opKlass(arg1)
}

var _wrap_new_ABC unsafe.Pointer

func _swig_wrap_new_ABC() (base SwigcptrABC) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_ABC, _swig_p)
	return
}

func NewABC() (_swig_ret ABC) {
	return _swig_wrap_new_ABC()
}

var _wrap_delete_ABC unsafe.Pointer

func _swig_wrap_delete_ABC(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_ABC, _swig_p)
	return
}

func DeleteABC(arg1 ABC) {
	_swig_wrap_delete_ABC(arg1.Swigcptr())
}

type ABC interface {
	Swigcptr() uintptr
	SwigIsABC()
	MethodABC(arg2 ABC)
	MethodKlass(arg2 Klass)
	OpABC() (_swig_ret ABC)
	OpKlass() (_swig_ret Klass)
}

type SwigcptrXYZInt uintptr

func (p SwigcptrXYZInt) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrXYZInt) SwigIsXYZInt() {
}

var _wrap_XYZInt_opIntPtrA unsafe.Pointer

func _swig_wrap_XYZInt_opIntPtrA(base SwigcptrXYZInt) (_ SwigcptrNotXYZInt) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_XYZInt_opIntPtrA, _swig_p)
	return
}

func (arg1 SwigcptrXYZInt) OpIntPtrA() (_swig_ret NotXYZInt) {
	return _swig_wrap_XYZInt_opIntPtrA(arg1)
}

var _wrap_XYZInt_opIntPtrB unsafe.Pointer

func _swig_wrap_XYZInt_opIntPtrB(base SwigcptrXYZInt) (_ SwigcptrXYZInt) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_XYZInt_opIntPtrB, _swig_p)
	return
}

func (arg1 SwigcptrXYZInt) OpIntPtrB() (_swig_ret XYZInt) {
	return _swig_wrap_XYZInt_opIntPtrB(arg1)
}

var _wrap_XYZInt_opAnother2 unsafe.Pointer

func _swig_wrap_XYZInt_opAnother2(base SwigcptrXYZInt) (_ SwigcptrAnother) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_XYZInt_opAnother2, _swig_p)
	return
}

func (arg1 SwigcptrXYZInt) OpAnother2() (_swig_ret Another) {
	return _swig_wrap_XYZInt_opAnother2(arg1)
}

var _wrap_XYZInt_tMethod2 unsafe.Pointer

func _swig_wrap_XYZInt_tMethod2(base SwigcptrXYZInt, _ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_XYZInt_tMethod2, _swig_p)
	return
}

func (arg1 SwigcptrXYZInt) TMethod2(arg2 int) {
	_swig_wrap_XYZInt_tMethod2(arg1, arg2)
}

var _wrap_XYZInt_tMethodNotXYZ2 unsafe.Pointer

func _swig_wrap_XYZInt_tMethodNotXYZ2(base SwigcptrXYZInt, _ uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_XYZInt_tMethodNotXYZ2, _swig_p)
	return
}

func (arg1 SwigcptrXYZInt) TMethodNotXYZ2(arg2 NotXYZInt) {
	_swig_wrap_XYZInt_tMethodNotXYZ2(arg1, arg2.Swigcptr())
}

var _wrap_XYZInt_tMethodXYZ2 unsafe.Pointer

func _swig_wrap_XYZInt_tMethodXYZ2(base SwigcptrXYZInt, _ uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_XYZInt_tMethodXYZ2, _swig_p)
	return
}

func (arg1 SwigcptrXYZInt) TMethodXYZ2(arg2 XYZInt) {
	_swig_wrap_XYZInt_tMethodXYZ2(arg1, arg2.Swigcptr())
}

var _wrap_XYZInt_opT2 unsafe.Pointer

func _swig_wrap_XYZInt_opT2(base SwigcptrXYZInt) (_ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_XYZInt_opT2, _swig_p)
	return
}

func (arg1 SwigcptrXYZInt) OpT2() (_swig_ret int) {
	return _swig_wrap_XYZInt_opT2(arg1)
}

var _wrap_XYZInt_opNotXYZ2 unsafe.Pointer

func _swig_wrap_XYZInt_opNotXYZ2(base SwigcptrXYZInt) (_ SwigcptrNotXYZInt) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_XYZInt_opNotXYZ2, _swig_p)
	return
}

func (arg1 SwigcptrXYZInt) OpNotXYZ2() (_swig_ret NotXYZInt) {
	return _swig_wrap_XYZInt_opNotXYZ2(arg1)
}

var _wrap_XYZInt_opXYZ2 unsafe.Pointer

func _swig_wrap_XYZInt_opXYZ2(base SwigcptrXYZInt) (_ SwigcptrXYZInt) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_XYZInt_opXYZ2, _swig_p)
	return
}

func (arg1 SwigcptrXYZInt) OpXYZ2() (_swig_ret XYZInt) {
	return _swig_wrap_XYZInt_opXYZ2(arg1)
}

var _wrap_new_XYZInt unsafe.Pointer

func _swig_wrap_new_XYZInt() (base SwigcptrXYZInt) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_XYZInt, _swig_p)
	return
}

func NewXYZInt() (_swig_ret XYZInt) {
	return _swig_wrap_new_XYZInt()
}

var _wrap_delete_XYZInt unsafe.Pointer

func _swig_wrap_delete_XYZInt(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_XYZInt, _swig_p)
	return
}

func DeleteXYZInt(arg1 XYZInt) {
	_swig_wrap_delete_XYZInt(arg1.Swigcptr())
}

type XYZInt interface {
	Swigcptr() uintptr
	SwigIsXYZInt()
	OpIntPtrA() (_swig_ret NotXYZInt)
	OpIntPtrB() (_swig_ret XYZInt)
	OpAnother2() (_swig_ret Another)
	TMethod2(arg2 int)
	TMethodNotXYZ2(arg2 NotXYZInt)
	TMethodXYZ2(arg2 XYZInt)
	OpT2() (_swig_ret int)
	OpNotXYZ2() (_swig_ret NotXYZInt)
	OpXYZ2() (_swig_ret XYZInt)
}

type SwigcptrXYZDouble uintptr

func (p SwigcptrXYZDouble) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrXYZDouble) SwigIsXYZDouble() {
}

var _wrap_XYZDouble_opIntPtrA unsafe.Pointer

func _swig_wrap_XYZDouble_opIntPtrA(base SwigcptrXYZDouble) (_ SwigcptrNotXYZInt) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_XYZDouble_opIntPtrA, _swig_p)
	return
}

func (arg1 SwigcptrXYZDouble) OpIntPtrA() (_swig_ret NotXYZInt) {
	return _swig_wrap_XYZDouble_opIntPtrA(arg1)
}

var _wrap_XYZDouble_opIntPtrB unsafe.Pointer

func _swig_wrap_XYZDouble_opIntPtrB(base SwigcptrXYZDouble) (_ SwigcptrXYZInt) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_XYZDouble_opIntPtrB, _swig_p)
	return
}

func (arg1 SwigcptrXYZDouble) OpIntPtrB() (_swig_ret XYZInt) {
	return _swig_wrap_XYZDouble_opIntPtrB(arg1)
}

var _wrap_XYZDouble_opAnother1 unsafe.Pointer

func _swig_wrap_XYZDouble_opAnother1(base SwigcptrXYZDouble) (_ SwigcptrAnother) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_XYZDouble_opAnother1, _swig_p)
	return
}

func (arg1 SwigcptrXYZDouble) OpAnother1() (_swig_ret Another) {
	return _swig_wrap_XYZDouble_opAnother1(arg1)
}

var _wrap_XYZDouble_tMethod1 unsafe.Pointer

func _swig_wrap_XYZDouble_tMethod1(base SwigcptrXYZDouble, _ float64) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_XYZDouble_tMethod1, _swig_p)
	return
}

func (arg1 SwigcptrXYZDouble) TMethod1(arg2 float64) {
	_swig_wrap_XYZDouble_tMethod1(arg1, arg2)
}

var _wrap_XYZDouble_tMethodNotXYZ1 unsafe.Pointer

func _swig_wrap_XYZDouble_tMethodNotXYZ1(base SwigcptrXYZDouble, _ uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_XYZDouble_tMethodNotXYZ1, _swig_p)
	return
}

func (arg1 SwigcptrXYZDouble) TMethodNotXYZ1(arg2 NotXYZDouble) {
	_swig_wrap_XYZDouble_tMethodNotXYZ1(arg1, arg2.Swigcptr())
}

var _wrap_XYZDouble_tMethodXYZ1 unsafe.Pointer

func _swig_wrap_XYZDouble_tMethodXYZ1(base SwigcptrXYZDouble, _ uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_XYZDouble_tMethodXYZ1, _swig_p)
	return
}

func (arg1 SwigcptrXYZDouble) TMethodXYZ1(arg2 XYZDouble) {
	_swig_wrap_XYZDouble_tMethodXYZ1(arg1, arg2.Swigcptr())
}

var _wrap_XYZDouble_opT1 unsafe.Pointer

func _swig_wrap_XYZDouble_opT1(base SwigcptrXYZDouble) (_ float64) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_XYZDouble_opT1, _swig_p)
	return
}

func (arg1 SwigcptrXYZDouble) OpT1() (_swig_ret float64) {
	return _swig_wrap_XYZDouble_opT1(arg1)
}

var _wrap_XYZDouble_opNotXYZ1 unsafe.Pointer

func _swig_wrap_XYZDouble_opNotXYZ1(base SwigcptrXYZDouble) (_ SwigcptrNotXYZDouble) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_XYZDouble_opNotXYZ1, _swig_p)
	return
}

func (arg1 SwigcptrXYZDouble) OpNotXYZ1() (_swig_ret NotXYZDouble) {
	return _swig_wrap_XYZDouble_opNotXYZ1(arg1)
}

var _wrap_XYZDouble_opXYZ1 unsafe.Pointer

func _swig_wrap_XYZDouble_opXYZ1(base SwigcptrXYZDouble) (_ SwigcptrXYZDouble) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_XYZDouble_opXYZ1, _swig_p)
	return
}

func (arg1 SwigcptrXYZDouble) OpXYZ1() (_swig_ret XYZDouble) {
	return _swig_wrap_XYZDouble_opXYZ1(arg1)
}

var _wrap_new_XYZDouble unsafe.Pointer

func _swig_wrap_new_XYZDouble() (base SwigcptrXYZDouble) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_XYZDouble, _swig_p)
	return
}

func NewXYZDouble() (_swig_ret XYZDouble) {
	return _swig_wrap_new_XYZDouble()
}

var _wrap_delete_XYZDouble unsafe.Pointer

func _swig_wrap_delete_XYZDouble(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_XYZDouble, _swig_p)
	return
}

func DeleteXYZDouble(arg1 XYZDouble) {
	_swig_wrap_delete_XYZDouble(arg1.Swigcptr())
}

type XYZDouble interface {
	Swigcptr() uintptr
	SwigIsXYZDouble()
	OpIntPtrA() (_swig_ret NotXYZInt)
	OpIntPtrB() (_swig_ret XYZInt)
	OpAnother1() (_swig_ret Another)
	TMethod1(arg2 float64)
	TMethodNotXYZ1(arg2 NotXYZDouble)
	TMethodXYZ1(arg2 XYZDouble)
	OpT1() (_swig_ret float64)
	OpNotXYZ1() (_swig_ret NotXYZDouble)
	OpXYZ1() (_swig_ret XYZDouble)
}

type SwigcptrXYZKlass uintptr

func (p SwigcptrXYZKlass) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrXYZKlass) SwigIsXYZKlass() {
}

var _wrap_XYZKlass_opIntPtrA unsafe.Pointer

func _swig_wrap_XYZKlass_opIntPtrA(base SwigcptrXYZKlass) (_ SwigcptrNotXYZInt) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_XYZKlass_opIntPtrA, _swig_p)
	return
}

func (arg1 SwigcptrXYZKlass) OpIntPtrA() (_swig_ret NotXYZInt) {
	return _swig_wrap_XYZKlass_opIntPtrA(arg1)
}

var _wrap_XYZKlass_opIntPtrB unsafe.Pointer

func _swig_wrap_XYZKlass_opIntPtrB(base SwigcptrXYZKlass) (_ SwigcptrXYZInt) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_XYZKlass_opIntPtrB, _swig_p)
	return
}

func (arg1 SwigcptrXYZKlass) OpIntPtrB() (_swig_ret XYZInt) {
	return _swig_wrap_XYZKlass_opIntPtrB(arg1)
}

var _wrap_XYZKlass_opAnother3 unsafe.Pointer

func _swig_wrap_XYZKlass_opAnother3(base SwigcptrXYZKlass) (_ SwigcptrAnother) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_XYZKlass_opAnother3, _swig_p)
	return
}

func (arg1 SwigcptrXYZKlass) OpAnother3() (_swig_ret Another) {
	return _swig_wrap_XYZKlass_opAnother3(arg1)
}

var _wrap_XYZKlass_tMethod3 unsafe.Pointer

func _swig_wrap_XYZKlass_tMethod3(base SwigcptrXYZKlass, _ uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_XYZKlass_tMethod3, _swig_p)
	return
}

func (arg1 SwigcptrXYZKlass) TMethod3(arg2 Klass) {
	_swig_wrap_XYZKlass_tMethod3(arg1, arg2.Swigcptr())
}

var _wrap_XYZKlass_tMethodNotXYZ3 unsafe.Pointer

func _swig_wrap_XYZKlass_tMethodNotXYZ3(base SwigcptrXYZKlass, _ uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_XYZKlass_tMethodNotXYZ3, _swig_p)
	return
}

func (arg1 SwigcptrXYZKlass) TMethodNotXYZ3(arg2 NotXYZKlass) {
	_swig_wrap_XYZKlass_tMethodNotXYZ3(arg1, arg2.Swigcptr())
}

var _wrap_XYZKlass_tMethodXYZ3 unsafe.Pointer

func _swig_wrap_XYZKlass_tMethodXYZ3(base SwigcptrXYZKlass, _ uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_XYZKlass_tMethodXYZ3, _swig_p)
	return
}

func (arg1 SwigcptrXYZKlass) TMethodXYZ3(arg2 XYZKlass) {
	_swig_wrap_XYZKlass_tMethodXYZ3(arg1, arg2.Swigcptr())
}

var _wrap_XYZKlass_opT3 unsafe.Pointer

func _swig_wrap_XYZKlass_opT3(base SwigcptrXYZKlass) (_ SwigcptrKlass) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_XYZKlass_opT3, _swig_p)
	return
}

func (arg1 SwigcptrXYZKlass) OpT3() (_swig_ret Klass) {
	return _swig_wrap_XYZKlass_opT3(arg1)
}

var _wrap_XYZKlass_opNotXYZ3 unsafe.Pointer

func _swig_wrap_XYZKlass_opNotXYZ3(base SwigcptrXYZKlass) (_ SwigcptrNotXYZKlass) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_XYZKlass_opNotXYZ3, _swig_p)
	return
}

func (arg1 SwigcptrXYZKlass) OpNotXYZ3() (_swig_ret NotXYZKlass) {
	return _swig_wrap_XYZKlass_opNotXYZ3(arg1)
}

var _wrap_XYZKlass_opXYZ3 unsafe.Pointer

func _swig_wrap_XYZKlass_opXYZ3(base SwigcptrXYZKlass) (_ SwigcptrXYZKlass) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_XYZKlass_opXYZ3, _swig_p)
	return
}

func (arg1 SwigcptrXYZKlass) OpXYZ3() (_swig_ret XYZKlass) {
	return _swig_wrap_XYZKlass_opXYZ3(arg1)
}

var _wrap_new_XYZKlass unsafe.Pointer

func _swig_wrap_new_XYZKlass() (base SwigcptrXYZKlass) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_XYZKlass, _swig_p)
	return
}

func NewXYZKlass() (_swig_ret XYZKlass) {
	return _swig_wrap_new_XYZKlass()
}

var _wrap_delete_XYZKlass unsafe.Pointer

func _swig_wrap_delete_XYZKlass(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_XYZKlass, _swig_p)
	return
}

func DeleteXYZKlass(arg1 XYZKlass) {
	_swig_wrap_delete_XYZKlass(arg1.Swigcptr())
}

type XYZKlass interface {
	Swigcptr() uintptr
	SwigIsXYZKlass()
	OpIntPtrA() (_swig_ret NotXYZInt)
	OpIntPtrB() (_swig_ret XYZInt)
	OpAnother3() (_swig_ret Another)
	TMethod3(arg2 Klass)
	TMethodNotXYZ3(arg2 NotXYZKlass)
	TMethodXYZ3(arg2 XYZKlass)
	OpT3() (_swig_ret Klass)
	OpNotXYZ3() (_swig_ret NotXYZKlass)
	OpXYZ3() (_swig_ret XYZKlass)
}

type SwigcptrXYZEnu uintptr

func (p SwigcptrXYZEnu) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrXYZEnu) SwigIsXYZEnu() {
}

var _wrap_XYZEnu_opIntPtrA unsafe.Pointer

func _swig_wrap_XYZEnu_opIntPtrA(base SwigcptrXYZEnu) (_ SwigcptrNotXYZInt) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_XYZEnu_opIntPtrA, _swig_p)
	return
}

func (arg1 SwigcptrXYZEnu) OpIntPtrA() (_swig_ret NotXYZInt) {
	return _swig_wrap_XYZEnu_opIntPtrA(arg1)
}

var _wrap_XYZEnu_opIntPtrB unsafe.Pointer

func _swig_wrap_XYZEnu_opIntPtrB(base SwigcptrXYZEnu) (_ SwigcptrXYZInt) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_XYZEnu_opIntPtrB, _swig_p)
	return
}

func (arg1 SwigcptrXYZEnu) OpIntPtrB() (_swig_ret XYZInt) {
	return _swig_wrap_XYZEnu_opIntPtrB(arg1)
}

var _wrap_XYZEnu_opAnother4 unsafe.Pointer

func _swig_wrap_XYZEnu_opAnother4(base SwigcptrXYZEnu) (_ SwigcptrAnother) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_XYZEnu_opAnother4, _swig_p)
	return
}

func (arg1 SwigcptrXYZEnu) OpAnother4() (_swig_ret Another) {
	return _swig_wrap_XYZEnu_opAnother4(arg1)
}

var _wrap_XYZEnu_tMethod4 unsafe.Pointer

func _swig_wrap_XYZEnu_tMethod4(base SwigcptrXYZEnu, _ SpaceEnu) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_XYZEnu_tMethod4, _swig_p)
	return
}

func (arg1 SwigcptrXYZEnu) TMethod4(arg2 SpaceEnu) {
	_swig_wrap_XYZEnu_tMethod4(arg1, arg2)
}

var _wrap_XYZEnu_tMethodNotXYZ4 unsafe.Pointer

func _swig_wrap_XYZEnu_tMethodNotXYZ4(base SwigcptrXYZEnu, _ uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_XYZEnu_tMethodNotXYZ4, _swig_p)
	return
}

func (arg1 SwigcptrXYZEnu) TMethodNotXYZ4(arg2 NotXYZEnu) {
	_swig_wrap_XYZEnu_tMethodNotXYZ4(arg1, arg2.Swigcptr())
}

var _wrap_XYZEnu_tMethodXYZ4 unsafe.Pointer

func _swig_wrap_XYZEnu_tMethodXYZ4(base SwigcptrXYZEnu, _ uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_XYZEnu_tMethodXYZ4, _swig_p)
	return
}

func (arg1 SwigcptrXYZEnu) TMethodXYZ4(arg2 XYZEnu) {
	_swig_wrap_XYZEnu_tMethodXYZ4(arg1, arg2.Swigcptr())
}

var _wrap_XYZEnu_opT4 unsafe.Pointer

func _swig_wrap_XYZEnu_opT4(base SwigcptrXYZEnu) (_ SpaceEnu) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_XYZEnu_opT4, _swig_p)
	return
}

func (arg1 SwigcptrXYZEnu) OpT4() (_swig_ret SpaceEnu) {
	return _swig_wrap_XYZEnu_opT4(arg1)
}

var _wrap_XYZEnu_opNotXYZ4 unsafe.Pointer

func _swig_wrap_XYZEnu_opNotXYZ4(base SwigcptrXYZEnu) (_ SwigcptrNotXYZEnu) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_XYZEnu_opNotXYZ4, _swig_p)
	return
}

func (arg1 SwigcptrXYZEnu) OpNotXYZ4() (_swig_ret NotXYZEnu) {
	return _swig_wrap_XYZEnu_opNotXYZ4(arg1)
}

var _wrap_XYZEnu_opXYZ4 unsafe.Pointer

func _swig_wrap_XYZEnu_opXYZ4(base SwigcptrXYZEnu) (_ SwigcptrXYZEnu) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_XYZEnu_opXYZ4, _swig_p)
	return
}

func (arg1 SwigcptrXYZEnu) OpXYZ4() (_swig_ret XYZEnu) {
	return _swig_wrap_XYZEnu_opXYZ4(arg1)
}

var _wrap_new_XYZEnu unsafe.Pointer

func _swig_wrap_new_XYZEnu() (base SwigcptrXYZEnu) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_XYZEnu, _swig_p)
	return
}

func NewXYZEnu() (_swig_ret XYZEnu) {
	return _swig_wrap_new_XYZEnu()
}

var _wrap_delete_XYZEnu unsafe.Pointer

func _swig_wrap_delete_XYZEnu(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_XYZEnu, _swig_p)
	return
}

func DeleteXYZEnu(arg1 XYZEnu) {
	_swig_wrap_delete_XYZEnu(arg1.Swigcptr())
}

type XYZEnu interface {
	Swigcptr() uintptr
	SwigIsXYZEnu()
	OpIntPtrA() (_swig_ret NotXYZInt)
	OpIntPtrB() (_swig_ret XYZInt)
	OpAnother4() (_swig_ret Another)
	TMethod4(arg2 SpaceEnu)
	TMethodNotXYZ4(arg2 NotXYZEnu)
	TMethodXYZ4(arg2 XYZEnu)
	OpT4() (_swig_ret SpaceEnu)
	OpNotXYZ4() (_swig_ret NotXYZEnu)
	OpXYZ4() (_swig_ret XYZEnu)
}

type SwigcptrNotXYZInt uintptr

func (p SwigcptrNotXYZInt) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrNotXYZInt) SwigIsNotXYZInt() {
}

var _wrap_new_NotXYZInt unsafe.Pointer

func _swig_wrap_new_NotXYZInt() (base SwigcptrNotXYZInt) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_NotXYZInt, _swig_p)
	return
}

func NewNotXYZInt() (_swig_ret NotXYZInt) {
	return _swig_wrap_new_NotXYZInt()
}

var _wrap_delete_NotXYZInt unsafe.Pointer

func _swig_wrap_delete_NotXYZInt(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_NotXYZInt, _swig_p)
	return
}

func DeleteNotXYZInt(arg1 NotXYZInt) {
	_swig_wrap_delete_NotXYZInt(arg1.Swigcptr())
}

type NotXYZInt interface {
	Swigcptr() uintptr
	SwigIsNotXYZInt()
}

type SwigcptrNotXYZDouble uintptr

func (p SwigcptrNotXYZDouble) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrNotXYZDouble) SwigIsNotXYZDouble() {
}

var _wrap_new_NotXYZDouble unsafe.Pointer

func _swig_wrap_new_NotXYZDouble() (base SwigcptrNotXYZDouble) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_NotXYZDouble, _swig_p)
	return
}

func NewNotXYZDouble() (_swig_ret NotXYZDouble) {
	return _swig_wrap_new_NotXYZDouble()
}

var _wrap_delete_NotXYZDouble unsafe.Pointer

func _swig_wrap_delete_NotXYZDouble(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_NotXYZDouble, _swig_p)
	return
}

func DeleteNotXYZDouble(arg1 NotXYZDouble) {
	_swig_wrap_delete_NotXYZDouble(arg1.Swigcptr())
}

type NotXYZDouble interface {
	Swigcptr() uintptr
	SwigIsNotXYZDouble()
}

type SwigcptrNotXYZKlass uintptr

func (p SwigcptrNotXYZKlass) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrNotXYZKlass) SwigIsNotXYZKlass() {
}

var _wrap_new_NotXYZKlass unsafe.Pointer

func _swig_wrap_new_NotXYZKlass() (base SwigcptrNotXYZKlass) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_NotXYZKlass, _swig_p)
	return
}

func NewNotXYZKlass() (_swig_ret NotXYZKlass) {
	return _swig_wrap_new_NotXYZKlass()
}

var _wrap_delete_NotXYZKlass unsafe.Pointer

func _swig_wrap_delete_NotXYZKlass(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_NotXYZKlass, _swig_p)
	return
}

func DeleteNotXYZKlass(arg1 NotXYZKlass) {
	_swig_wrap_delete_NotXYZKlass(arg1.Swigcptr())
}

type NotXYZKlass interface {
	Swigcptr() uintptr
	SwigIsNotXYZKlass()
}

type SwigcptrNotXYZEnu uintptr

func (p SwigcptrNotXYZEnu) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrNotXYZEnu) SwigIsNotXYZEnu() {
}

var _wrap_new_NotXYZEnu unsafe.Pointer

func _swig_wrap_new_NotXYZEnu() (base SwigcptrNotXYZEnu) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_NotXYZEnu, _swig_p)
	return
}

func NewNotXYZEnu() (_swig_ret NotXYZEnu) {
	return _swig_wrap_new_NotXYZEnu()
}

var _wrap_delete_NotXYZEnu unsafe.Pointer

func _swig_wrap_delete_NotXYZEnu(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_NotXYZEnu, _swig_p)
	return
}

func DeleteNotXYZEnu(arg1 NotXYZEnu) {
	_swig_wrap_delete_NotXYZEnu(arg1.Swigcptr())
}

type NotXYZEnu interface {
	Swigcptr() uintptr
	SwigIsNotXYZEnu()
}


type SwigcptrVoid uintptr
type Void interface {
	Swigcptr() uintptr;
}
func (p SwigcptrVoid) Swigcptr() uintptr {
	return uintptr(p)
}

