/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 3.0.3
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: ./../inherit_target_language.i

package inherit_target_language

import "unsafe"
import _ "runtime/cgo"

var _cgo_runtime_cgocall func(unsafe.Pointer, uintptr)



type _ unsafe.Pointer



type _swig_fnptr *byte
type _swig_memberptr *byte

type SwigcptrBase2 uintptr

func (p SwigcptrBase2) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrBase2) SwigIsBase2() {
}

var _wrap_delete_Base2 unsafe.Pointer

func _swig_wrap_delete_Base2(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_Base2, _swig_p)
	return
}

func DeleteBase2(arg1 Base2) {
	_swig_wrap_delete_Base2(arg1.Swigcptr())
}

var _wrap_new_Base2 unsafe.Pointer

func _swig_wrap_new_Base2() (base SwigcptrBase2) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_Base2, _swig_p)
	return
}

func NewBase2() (_swig_ret Base2) {
	return _swig_wrap_new_Base2()
}

type Base2 interface {
	Swigcptr() uintptr
	SwigIsBase2()
}

type SwigcptrDerived1 uintptr

func (p SwigcptrDerived1) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrDerived1) SwigIsDerived1() {
}

var _wrap_new_Derived1 unsafe.Pointer

func _swig_wrap_new_Derived1() (base SwigcptrDerived1) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_Derived1, _swig_p)
	return
}

func NewDerived1() (_swig_ret Derived1) {
	return _swig_wrap_new_Derived1()
}

var _wrap_delete_Derived1 unsafe.Pointer

func _swig_wrap_delete_Derived1(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_Derived1, _swig_p)
	return
}

func DeleteDerived1(arg1 Derived1) {
	_swig_wrap_delete_Derived1(arg1.Swigcptr())
}

type Derived1 interface {
	Swigcptr() uintptr
	SwigIsDerived1()
}

type SwigcptrDerived2 uintptr

func (p SwigcptrDerived2) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrDerived2) SwigIsDerived2() {
}

var _wrap_new_Derived2 unsafe.Pointer

func _swig_wrap_new_Derived2() (base SwigcptrDerived2) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_Derived2, _swig_p)
	return
}

func NewDerived2() (_swig_ret Derived2) {
	return _swig_wrap_new_Derived2()
}

var _wrap_delete_Derived2 unsafe.Pointer

func _swig_wrap_delete_Derived2(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_Derived2, _swig_p)
	return
}

func DeleteDerived2(arg1 Derived2) {
	_swig_wrap_delete_Derived2(arg1.Swigcptr())
}

func (p SwigcptrDerived2) SwigIsBase2() {
}

func (p SwigcptrDerived2) SwigGetBase2() Base2 {
	return SwigcptrBase2(p.Swigcptr())
}

type Derived2 interface {
	Swigcptr() uintptr
	SwigIsDerived2()
	SwigIsBase2()
	SwigGetBase2() Base2
}

type SwigcptrMBase2a uintptr

func (p SwigcptrMBase2a) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrMBase2a) SwigIsMBase2a() {
}

var _wrap_delete_MBase2a unsafe.Pointer

func _swig_wrap_delete_MBase2a(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_MBase2a, _swig_p)
	return
}

func DeleteMBase2a(arg1 MBase2a) {
	_swig_wrap_delete_MBase2a(arg1.Swigcptr())
}

var _wrap_MBase2a_c unsafe.Pointer

func _swig_wrap_MBase2a_c(base SwigcptrMBase2a) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_MBase2a_c, _swig_p)
	return
}

func (arg1 SwigcptrMBase2a) C() {
	_swig_wrap_MBase2a_c(arg1)
}

var _wrap_new_MBase2a unsafe.Pointer

func _swig_wrap_new_MBase2a() (base SwigcptrMBase2a) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_MBase2a, _swig_p)
	return
}

func NewMBase2a() (_swig_ret MBase2a) {
	return _swig_wrap_new_MBase2a()
}

type MBase2a interface {
	Swigcptr() uintptr
	SwigIsMBase2a()
	C()
}

type SwigcptrMBase2b uintptr

func (p SwigcptrMBase2b) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrMBase2b) SwigIsMBase2b() {
}

var _wrap_delete_MBase2b unsafe.Pointer

func _swig_wrap_delete_MBase2b(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_MBase2b, _swig_p)
	return
}

func DeleteMBase2b(arg1 MBase2b) {
	_swig_wrap_delete_MBase2b(arg1.Swigcptr())
}

var _wrap_MBase2b_d unsafe.Pointer

func _swig_wrap_MBase2b_d(base SwigcptrMBase2b) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_MBase2b_d, _swig_p)
	return
}

func (arg1 SwigcptrMBase2b) D() {
	_swig_wrap_MBase2b_d(arg1)
}

var _wrap_new_MBase2b unsafe.Pointer

func _swig_wrap_new_MBase2b() (base SwigcptrMBase2b) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_MBase2b, _swig_p)
	return
}

func NewMBase2b() (_swig_ret MBase2b) {
	return _swig_wrap_new_MBase2b()
}

type MBase2b interface {
	Swigcptr() uintptr
	SwigIsMBase2b()
	D()
}

type SwigcptrMultipleDerived1 uintptr

func (p SwigcptrMultipleDerived1) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrMultipleDerived1) SwigIsMultipleDerived1() {
}

var _wrap_new_MultipleDerived1 unsafe.Pointer

func _swig_wrap_new_MultipleDerived1() (base SwigcptrMultipleDerived1) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_MultipleDerived1, _swig_p)
	return
}

func NewMultipleDerived1() (_swig_ret MultipleDerived1) {
	return _swig_wrap_new_MultipleDerived1()
}

var _wrap_delete_MultipleDerived1 unsafe.Pointer

func _swig_wrap_delete_MultipleDerived1(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_MultipleDerived1, _swig_p)
	return
}

func DeleteMultipleDerived1(arg1 MultipleDerived1) {
	_swig_wrap_delete_MultipleDerived1(arg1.Swigcptr())
}

type MultipleDerived1 interface {
	Swigcptr() uintptr
	SwigIsMultipleDerived1()
}

type SwigcptrMultipleDerived2 uintptr

func (p SwigcptrMultipleDerived2) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrMultipleDerived2) SwigIsMultipleDerived2() {
}

var _wrap_new_MultipleDerived2 unsafe.Pointer

func _swig_wrap_new_MultipleDerived2() (base SwigcptrMultipleDerived2) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_MultipleDerived2, _swig_p)
	return
}

func NewMultipleDerived2() (_swig_ret MultipleDerived2) {
	return _swig_wrap_new_MultipleDerived2()
}

var _wrap_delete_MultipleDerived2 unsafe.Pointer

func _swig_wrap_delete_MultipleDerived2(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_MultipleDerived2, _swig_p)
	return
}

func DeleteMultipleDerived2(arg1 MultipleDerived2) {
	_swig_wrap_delete_MultipleDerived2(arg1.Swigcptr())
}

var _wrap_MultipleDerived2_d unsafe.Pointer

func _swig_wrap_MultipleDerived2_d(base SwigcptrMultipleDerived2) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_MultipleDerived2_d, _swig_p)
	return
}

func (_swig_base SwigcptrMultipleDerived2) D() {
	_swig_wrap_MultipleDerived2_d(_swig_base)
}

var _wrap_MultipleDerived2_SwigGetMBase2b unsafe.Pointer

func _swig_wrap_MultipleDerived2_SwigGetMBase2b(base SwigcptrMultipleDerived2) (_ SwigcptrMBase2b) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_MultipleDerived2_SwigGetMBase2b, _swig_p)
	return
}

func (arg1 SwigcptrMultipleDerived2) SwigGetMBase2b() (_swig_ret MBase2b) {
	return _swig_wrap_MultipleDerived2_SwigGetMBase2b(arg1)
}

type MultipleDerived2 interface {
	Swigcptr() uintptr
	SwigIsMultipleDerived2()
	D()
	SwigGetMBase2b() (_swig_ret MBase2b)
}

type SwigcptrMBase3b uintptr

func (p SwigcptrMBase3b) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrMBase3b) SwigIsMBase3b() {
}

var _wrap_delete_MBase3b unsafe.Pointer

func _swig_wrap_delete_MBase3b(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_MBase3b, _swig_p)
	return
}

func DeleteMBase3b(arg1 MBase3b) {
	_swig_wrap_delete_MBase3b(arg1.Swigcptr())
}

var _wrap_MBase3b_f unsafe.Pointer

func _swig_wrap_MBase3b_f(base SwigcptrMBase3b) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_MBase3b_f, _swig_p)
	return
}

func (arg1 SwigcptrMBase3b) F() {
	_swig_wrap_MBase3b_f(arg1)
}

var _wrap_new_MBase3b unsafe.Pointer

func _swig_wrap_new_MBase3b() (base SwigcptrMBase3b) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_MBase3b, _swig_p)
	return
}

func NewMBase3b() (_swig_ret MBase3b) {
	return _swig_wrap_new_MBase3b()
}

type MBase3b interface {
	Swigcptr() uintptr
	SwigIsMBase3b()
	F()
}

type SwigcptrMBase4a uintptr

func (p SwigcptrMBase4a) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrMBase4a) SwigIsMBase4a() {
}

var _wrap_delete_MBase4a unsafe.Pointer

func _swig_wrap_delete_MBase4a(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_MBase4a, _swig_p)
	return
}

func DeleteMBase4a(arg1 MBase4a) {
	_swig_wrap_delete_MBase4a(arg1.Swigcptr())
}

var _wrap_MBase4a_g unsafe.Pointer

func _swig_wrap_MBase4a_g(base SwigcptrMBase4a) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_MBase4a_g, _swig_p)
	return
}

func (arg1 SwigcptrMBase4a) G() {
	_swig_wrap_MBase4a_g(arg1)
}

var _wrap_new_MBase4a unsafe.Pointer

func _swig_wrap_new_MBase4a() (base SwigcptrMBase4a) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_MBase4a, _swig_p)
	return
}

func NewMBase4a() (_swig_ret MBase4a) {
	return _swig_wrap_new_MBase4a()
}

type MBase4a interface {
	Swigcptr() uintptr
	SwigIsMBase4a()
	G()
}

type SwigcptrMultipleDerived3 uintptr

func (p SwigcptrMultipleDerived3) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrMultipleDerived3) SwigIsMultipleDerived3() {
}

var _wrap_new_MultipleDerived3 unsafe.Pointer

func _swig_wrap_new_MultipleDerived3() (base SwigcptrMultipleDerived3) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_MultipleDerived3, _swig_p)
	return
}

func NewMultipleDerived3() (_swig_ret MultipleDerived3) {
	return _swig_wrap_new_MultipleDerived3()
}

var _wrap_delete_MultipleDerived3 unsafe.Pointer

func _swig_wrap_delete_MultipleDerived3(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_MultipleDerived3, _swig_p)
	return
}

func DeleteMultipleDerived3(arg1 MultipleDerived3) {
	_swig_wrap_delete_MultipleDerived3(arg1.Swigcptr())
}

var _wrap_MultipleDerived3_f unsafe.Pointer

func _swig_wrap_MultipleDerived3_f(base SwigcptrMultipleDerived3) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_MultipleDerived3_f, _swig_p)
	return
}

func (_swig_base SwigcptrMultipleDerived3) F() {
	_swig_wrap_MultipleDerived3_f(_swig_base)
}

var _wrap_MultipleDerived3_SwigGetMBase3b unsafe.Pointer

func _swig_wrap_MultipleDerived3_SwigGetMBase3b(base SwigcptrMultipleDerived3) (_ SwigcptrMBase3b) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_MultipleDerived3_SwigGetMBase3b, _swig_p)
	return
}

func (arg1 SwigcptrMultipleDerived3) SwigGetMBase3b() (_swig_ret MBase3b) {
	return _swig_wrap_MultipleDerived3_SwigGetMBase3b(arg1)
}

type MultipleDerived3 interface {
	Swigcptr() uintptr
	SwigIsMultipleDerived3()
	F()
	SwigGetMBase3b() (_swig_ret MBase3b)
}

type SwigcptrMultipleDerived4 uintptr

func (p SwigcptrMultipleDerived4) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrMultipleDerived4) SwigIsMultipleDerived4() {
}

var _wrap_new_MultipleDerived4 unsafe.Pointer

func _swig_wrap_new_MultipleDerived4() (base SwigcptrMultipleDerived4) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_MultipleDerived4, _swig_p)
	return
}

func NewMultipleDerived4() (_swig_ret MultipleDerived4) {
	return _swig_wrap_new_MultipleDerived4()
}

var _wrap_delete_MultipleDerived4 unsafe.Pointer

func _swig_wrap_delete_MultipleDerived4(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_MultipleDerived4, _swig_p)
	return
}

func DeleteMultipleDerived4(arg1 MultipleDerived4) {
	_swig_wrap_delete_MultipleDerived4(arg1.Swigcptr())
}

var _wrap_MultipleDerived4_g unsafe.Pointer

func _swig_wrap_MultipleDerived4_g(base SwigcptrMultipleDerived4) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_MultipleDerived4_g, _swig_p)
	return
}

func (_swig_base SwigcptrMultipleDerived4) G() {
	_swig_wrap_MultipleDerived4_g(_swig_base)
}

func (p SwigcptrMultipleDerived4) SwigIsMBase4a() {
}

func (p SwigcptrMultipleDerived4) SwigGetMBase4a() MBase4a {
	return SwigcptrMBase4a(p.Swigcptr())
}

type MultipleDerived4 interface {
	Swigcptr() uintptr
	SwigIsMultipleDerived4()
	G()
	SwigIsMBase4a()
	SwigGetMBase4a() MBase4a
}

type SwigcptrBaseX uintptr

func (p SwigcptrBaseX) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrBaseX) SwigIsBaseX() {
}

var _wrap_delete_BaseX unsafe.Pointer

func _swig_wrap_delete_BaseX(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_BaseX, _swig_p)
	return
}

func DeleteBaseX(arg1 BaseX) {
	_swig_wrap_delete_BaseX(arg1.Swigcptr())
}

var _wrap_BaseX_basex unsafe.Pointer

func _swig_wrap_BaseX_basex(base SwigcptrBaseX) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_BaseX_basex, _swig_p)
	return
}

func (arg1 SwigcptrBaseX) Basex() {
	_swig_wrap_BaseX_basex(arg1)
}

var _wrap_new_BaseX unsafe.Pointer

func _swig_wrap_new_BaseX() (base SwigcptrBaseX) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_BaseX, _swig_p)
	return
}

func NewBaseX() (_swig_ret BaseX) {
	return _swig_wrap_new_BaseX()
}

type BaseX interface {
	Swigcptr() uintptr
	SwigIsBaseX()
	Basex()
}

type SwigcptrDerivedX uintptr

func (p SwigcptrDerivedX) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrDerivedX) SwigIsDerivedX() {
}

var _wrap_DerivedX_derivedx unsafe.Pointer

func _swig_wrap_DerivedX_derivedx(base SwigcptrDerivedX) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_DerivedX_derivedx, _swig_p)
	return
}

func (arg1 SwigcptrDerivedX) Derivedx() {
	_swig_wrap_DerivedX_derivedx(arg1)
}

var _wrap_new_DerivedX unsafe.Pointer

func _swig_wrap_new_DerivedX() (base SwigcptrDerivedX) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_DerivedX, _swig_p)
	return
}

func NewDerivedX() (_swig_ret DerivedX) {
	return _swig_wrap_new_DerivedX()
}

var _wrap_delete_DerivedX unsafe.Pointer

func _swig_wrap_delete_DerivedX(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_DerivedX, _swig_p)
	return
}

func DeleteDerivedX(arg1 DerivedX) {
	_swig_wrap_delete_DerivedX(arg1.Swigcptr())
}

var _wrap_DerivedX_basex unsafe.Pointer

func _swig_wrap_DerivedX_basex(base SwigcptrDerivedX) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_DerivedX_basex, _swig_p)
	return
}

func (_swig_base SwigcptrDerivedX) Basex() {
	_swig_wrap_DerivedX_basex(_swig_base)
}

func (p SwigcptrDerivedX) SwigIsBaseX() {
}

func (p SwigcptrDerivedX) SwigGetBaseX() BaseX {
	return SwigcptrBaseX(p.Swigcptr())
}

type DerivedX interface {
	Swigcptr() uintptr
	SwigIsDerivedX()
	Derivedx()
	Basex()
	SwigIsBaseX()
	SwigGetBaseX() BaseX
}


type SwigcptrVoid uintptr
type Void interface {
	Swigcptr() uintptr;
}
func (p SwigcptrVoid) Swigcptr() uintptr {
	return uintptr(p)
}

