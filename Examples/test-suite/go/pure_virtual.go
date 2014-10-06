/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 3.0.3
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: ./../pure_virtual.i

package pure_virtual

import "unsafe"
import _ "runtime/cgo"

var _cgo_runtime_cgocall func(unsafe.Pointer, uintptr)



type _ unsafe.Pointer



type _swig_fnptr *byte
type _swig_memberptr *byte

type SwigcptrA uintptr

func (p SwigcptrA) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrA) SwigIsA() {
}

var _wrap_delete_A unsafe.Pointer

func _swig_wrap_delete_A(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_A, _swig_p)
	return
}

func DeleteA(arg1 A) {
	_swig_wrap_delete_A(arg1.Swigcptr())
}

var _wrap_A_something unsafe.Pointer

func _swig_wrap_A_something(base SwigcptrA) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_A_something, _swig_p)
	return
}

func (arg1 SwigcptrA) Something() {
	_swig_wrap_A_something(arg1)
}

var _wrap_A_method unsafe.Pointer

func _swig_wrap_A_method(base SwigcptrA) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_A_method, _swig_p)
	return
}

func (arg1 SwigcptrA) Method() {
	_swig_wrap_A_method(arg1)
}

type A interface {
	Swigcptr() uintptr
	SwigIsA()
	Something()
	Method()
}

type SwigcptrB uintptr

func (p SwigcptrB) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrB) SwigIsB() {
}

var _wrap_new_B unsafe.Pointer

func _swig_wrap_new_B() (base SwigcptrB) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_B, _swig_p)
	return
}

func NewB() (_swig_ret B) {
	return _swig_wrap_new_B()
}

var _wrap_delete_B unsafe.Pointer

func _swig_wrap_delete_B(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_B, _swig_p)
	return
}

func DeleteB(arg1 B) {
	_swig_wrap_delete_B(arg1.Swigcptr())
}

var _wrap_B_something unsafe.Pointer

func _swig_wrap_B_something(base SwigcptrB) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_B_something, _swig_p)
	return
}

func (arg1 SwigcptrB) Something() {
	_swig_wrap_B_something(arg1)
}

var _wrap_B_method unsafe.Pointer

func _swig_wrap_B_method(base SwigcptrB) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_B_method, _swig_p)
	return
}

func (arg1 SwigcptrB) Method() {
	_swig_wrap_B_method(arg1)
}

func (p SwigcptrB) SwigIsA() {
}

func (p SwigcptrB) SwigGetA() A {
	return SwigcptrA(p.Swigcptr())
}

type B interface {
	Swigcptr() uintptr
	SwigIsB()
	Something()
	Method()
	SwigIsA()
	SwigGetA() A
}

type SwigcptrC uintptr

func (p SwigcptrC) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrC) SwigIsC() {
}

var _wrap_delete_C unsafe.Pointer

func _swig_wrap_delete_C(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_C, _swig_p)
	return
}

func DeleteC(arg1 C) {
	_swig_wrap_delete_C(arg1.Swigcptr())
}

var _wrap_C_method unsafe.Pointer

func _swig_wrap_C_method(base SwigcptrC) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_C_method, _swig_p)
	return
}

func (arg1 SwigcptrC) Method() {
	_swig_wrap_C_method(arg1)
}

var _wrap_C_something unsafe.Pointer

func _swig_wrap_C_something(base SwigcptrC) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_C_something, _swig_p)
	return
}

func (_swig_base SwigcptrC) Something() {
	_swig_wrap_C_something(_swig_base)
}

func (p SwigcptrC) SwigIsA() {
}

func (p SwigcptrC) SwigGetA() A {
	return SwigcptrA(p.Swigcptr())
}

type C interface {
	Swigcptr() uintptr
	SwigIsC()
	Method()
	Something()
	SwigIsA()
	SwigGetA() A
}

type SwigcptrD uintptr

func (p SwigcptrD) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrD) SwigIsD() {
}

var _wrap_delete_D unsafe.Pointer

func _swig_wrap_delete_D(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_D, _swig_p)
	return
}

func DeleteD(arg1 D) {
	_swig_wrap_delete_D(arg1.Swigcptr())
}

var _wrap_D_something unsafe.Pointer

func _swig_wrap_D_something(base SwigcptrD) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_D_something, _swig_p)
	return
}

func (arg1 SwigcptrD) Something() {
	_swig_wrap_D_something(arg1)
}

var _wrap_new_D unsafe.Pointer

func _swig_wrap_new_D() (base SwigcptrD) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_D, _swig_p)
	return
}

func NewD() (_swig_ret D) {
	return _swig_wrap_new_D()
}

var _wrap_D_method unsafe.Pointer

func _swig_wrap_D_method(base SwigcptrD) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_D_method, _swig_p)
	return
}

func (_swig_base SwigcptrD) Method() {
	_swig_wrap_D_method(_swig_base)
}

func (p SwigcptrD) SwigIsC() {
}

func (p SwigcptrD) SwigGetC() C {
	return SwigcptrC(p.Swigcptr())
}

func (p SwigcptrD) SwigIsA() {
}

func (p SwigcptrD) SwigGetA() A {
	return SwigcptrA(p.Swigcptr())
}

type D interface {
	Swigcptr() uintptr
	SwigIsD()
	Something()
	Method()
	SwigIsC()
	SwigGetC() C
	SwigIsA()
	SwigGetA() A
}

type SwigcptrAA uintptr

func (p SwigcptrAA) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrAA) SwigIsAA() {
}

var _wrap_delete_AA unsafe.Pointer

func _swig_wrap_delete_AA(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_AA, _swig_p)
	return
}

func DeleteAA(arg1 AA) {
	_swig_wrap_delete_AA(arg1.Swigcptr())
}

var _wrap_AA_method2 unsafe.Pointer

func _swig_wrap_AA_method2(base SwigcptrAA) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_AA_method2, _swig_p)
	return
}

func (arg1 SwigcptrAA) Method2() {
	_swig_wrap_AA_method2(arg1)
}

type AA interface {
	Swigcptr() uintptr
	SwigIsAA()
	Method2()
}

type SwigcptrE uintptr

func (p SwigcptrE) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrE) SwigIsE() {
}

var _wrap_E_something unsafe.Pointer

func _swig_wrap_E_something(base SwigcptrE) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_E_something, _swig_p)
	return
}

func (arg1 SwigcptrE) Something() {
	_swig_wrap_E_something(arg1)
}

var _wrap_E_method unsafe.Pointer

func _swig_wrap_E_method(base SwigcptrE) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_E_method, _swig_p)
	return
}

func (_swig_base SwigcptrE) Method() {
	_swig_wrap_E_method(_swig_base)
}

var _wrap_E_method2 unsafe.Pointer

func _swig_wrap_E_method2(base SwigcptrE) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_E_method2, _swig_p)
	return
}

func (_swig_base SwigcptrE) Method2() {
	_swig_wrap_E_method2(_swig_base)
}

func (p SwigcptrE) SwigIsC() {
}

func (p SwigcptrE) SwigGetC() C {
	return SwigcptrC(p.Swigcptr())
}

func (p SwigcptrE) SwigIsA() {
}

func (p SwigcptrE) SwigGetA() A {
	return SwigcptrA(p.Swigcptr())
}

var _wrap_E_SwigGetAA unsafe.Pointer

func _swig_wrap_E_SwigGetAA(base SwigcptrE) (_ SwigcptrAA) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_E_SwigGetAA, _swig_p)
	return
}

func (arg1 SwigcptrE) SwigGetAA() (_swig_ret AA) {
	return _swig_wrap_E_SwigGetAA(arg1)
}

type E interface {
	Swigcptr() uintptr
	SwigIsE()
	Something()
	Method()
	Method2()
	SwigIsC()
	SwigGetC() C
	SwigIsA()
	SwigGetA() A
	SwigGetAA() (_swig_ret AA)
}

type SwigcptrF uintptr

func (p SwigcptrF) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrF) SwigIsF() {
}

var _wrap_F_method2 unsafe.Pointer

func _swig_wrap_F_method2(base SwigcptrF) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_F_method2, _swig_p)
	return
}

func (arg1 SwigcptrF) Method2() {
	_swig_wrap_F_method2(arg1)
}

var _wrap_new_F unsafe.Pointer

func _swig_wrap_new_F() (base SwigcptrF) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_F, _swig_p)
	return
}

func NewF() (_swig_ret F) {
	return _swig_wrap_new_F()
}

var _wrap_delete_F unsafe.Pointer

func _swig_wrap_delete_F(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_F, _swig_p)
	return
}

func DeleteF(arg1 F) {
	_swig_wrap_delete_F(arg1.Swigcptr())
}

var _wrap_F_something unsafe.Pointer

func _swig_wrap_F_something(base SwigcptrF) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_F_something, _swig_p)
	return
}

func (_swig_base SwigcptrF) Something() {
	_swig_wrap_F_something(_swig_base)
}

var _wrap_F_method unsafe.Pointer

func _swig_wrap_F_method(base SwigcptrF) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_F_method, _swig_p)
	return
}

func (_swig_base SwigcptrF) Method() {
	_swig_wrap_F_method(_swig_base)
}

func (p SwigcptrF) SwigIsE() {
}

func (p SwigcptrF) SwigGetE() E {
	return SwigcptrE(p.Swigcptr())
}

func (p SwigcptrF) SwigIsC() {
}

func (p SwigcptrF) SwigGetC() C {
	return SwigcptrC(p.Swigcptr())
}

func (p SwigcptrF) SwigIsA() {
}

func (p SwigcptrF) SwigGetA() A {
	return SwigcptrA(p.Swigcptr())
}

func (p SwigcptrF) SwigGetAA() AA {
	return p.SwigGetE().SwigGetAA()
}

type F interface {
	Swigcptr() uintptr
	SwigIsF()
	Method2()
	Something()
	Method()
	SwigIsE()
	SwigGetE() E
	SwigIsC()
	SwigGetC() C
	SwigIsA()
	SwigGetA() A
	SwigGetAA() AA
}


type SwigcptrVoid uintptr
type Void interface {
	Swigcptr() uintptr;
}
func (p SwigcptrVoid) Swigcptr() uintptr {
	return uintptr(p)
}

