/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 3.0.3
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: ./../contract.i

package contract

import "unsafe"
import _ "runtime/cgo"

var _cgo_runtime_cgocall func(unsafe.Pointer, uintptr)



type _ unsafe.Pointer



type _swig_fnptr *byte
type _swig_memberptr *byte

var _wrap_test_preassert unsafe.Pointer

func Test_preassert(arg1 int, arg2 int) (_swig_ret int) {
	_swig_p := uintptr(unsafe.Pointer(&arg1))
	_cgo_runtime_cgocall(_wrap_test_preassert, _swig_p)
	return
}
var _wrap_test_postassert unsafe.Pointer

func Test_postassert(arg1 int) (_swig_ret int) {
	_swig_p := uintptr(unsafe.Pointer(&arg1))
	_cgo_runtime_cgocall(_wrap_test_postassert, _swig_p)
	return
}
var _wrap_test_prepost unsafe.Pointer

func Test_prepost(arg1 int, arg2 int) (_swig_ret int) {
	_swig_p := uintptr(unsafe.Pointer(&arg1))
	_cgo_runtime_cgocall(_wrap_test_prepost, _swig_p)
	return
}
type SwigcptrFoo uintptr

func (p SwigcptrFoo) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrFoo) SwigIsFoo() {
}

var _wrap_delete_Foo unsafe.Pointer

func _swig_wrap_delete_Foo(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_Foo, _swig_p)
	return
}

func DeleteFoo(arg1 Foo) {
	_swig_wrap_delete_Foo(arg1.Swigcptr())
}

var _wrap_Foo_test_preassert unsafe.Pointer

func _swig_wrap_Foo_test_preassert(base SwigcptrFoo, _ int, _ int) (_ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_Foo_test_preassert, _swig_p)
	return
}

func (arg1 SwigcptrFoo) Test_preassert(arg2 int, arg3 int) (_swig_ret int) {
	return _swig_wrap_Foo_test_preassert(arg1, arg2, arg3)
}

var _wrap_Foo_test_postassert unsafe.Pointer

func _swig_wrap_Foo_test_postassert(base SwigcptrFoo, _ int) (_ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_Foo_test_postassert, _swig_p)
	return
}

func (arg1 SwigcptrFoo) Test_postassert(arg2 int) (_swig_ret int) {
	return _swig_wrap_Foo_test_postassert(arg1, arg2)
}

var _wrap_Foo_test_prepost unsafe.Pointer

func _swig_wrap_Foo_test_prepost(base SwigcptrFoo, _ int, _ int) (_ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_Foo_test_prepost, _swig_p)
	return
}

func (arg1 SwigcptrFoo) Test_prepost(arg2 int, arg3 int) (_swig_ret int) {
	return _swig_wrap_Foo_test_prepost(arg1, arg2, arg3)
}

var _wrap_Foo_stest_prepost unsafe.Pointer

func FooStest_prepost(arg1 int, arg2 int) (_swig_ret int) {
	_swig_p := uintptr(unsafe.Pointer(&arg1))
	_cgo_runtime_cgocall(_wrap_Foo_stest_prepost, _swig_p)
	return
}
var _wrap_new_Foo unsafe.Pointer

func _swig_wrap_new_Foo() (base SwigcptrFoo) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_Foo, _swig_p)
	return
}

func NewFoo() (_swig_ret Foo) {
	return _swig_wrap_new_Foo()
}

type Foo interface {
	Swigcptr() uintptr
	SwigIsFoo()
	Test_preassert(arg2 int, arg3 int) (_swig_ret int)
	Test_postassert(arg2 int) (_swig_ret int)
	Test_prepost(arg2 int, arg3 int) (_swig_ret int)
}

type SwigcptrBar uintptr

func (p SwigcptrBar) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrBar) SwigIsBar() {
}

var _wrap_Bar_test_prepost unsafe.Pointer

func _swig_wrap_Bar_test_prepost(base SwigcptrBar, _ int, _ int) (_ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_Bar_test_prepost, _swig_p)
	return
}

func (arg1 SwigcptrBar) Test_prepost(arg2 int, arg3 int) (_swig_ret int) {
	return _swig_wrap_Bar_test_prepost(arg1, arg2, arg3)
}

var _wrap_new_Bar unsafe.Pointer

func _swig_wrap_new_Bar() (base SwigcptrBar) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_Bar, _swig_p)
	return
}

func NewBar() (_swig_ret Bar) {
	return _swig_wrap_new_Bar()
}

var _wrap_delete_Bar unsafe.Pointer

func _swig_wrap_delete_Bar(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_Bar, _swig_p)
	return
}

func DeleteBar(arg1 Bar) {
	_swig_wrap_delete_Bar(arg1.Swigcptr())
}

var _wrap_Bar_test_preassert unsafe.Pointer

func _swig_wrap_Bar_test_preassert(base SwigcptrBar, _ int, _ int) (_ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_Bar_test_preassert, _swig_p)
	return
}

func (_swig_base SwigcptrBar) Test_preassert(arg1 int, arg2 int) (_swig_ret int) {
	return _swig_wrap_Bar_test_preassert(_swig_base, arg1, arg2)
}

var _wrap_Bar_test_postassert unsafe.Pointer

func _swig_wrap_Bar_test_postassert(base SwigcptrBar, _ int) (_ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_Bar_test_postassert, _swig_p)
	return
}

func (_swig_base SwigcptrBar) Test_postassert(arg1 int) (_swig_ret int) {
	return _swig_wrap_Bar_test_postassert(_swig_base, arg1)
}

var _wrap_Bar_stest_prepost unsafe.Pointer

func BarStest_prepost(arg1 int, arg2 int) (_swig_ret int) {
	_swig_p := uintptr(unsafe.Pointer(&arg1))
	_cgo_runtime_cgocall(_wrap_Bar_stest_prepost, _swig_p)
	return
}
func (p SwigcptrBar) SwigIsFoo() {
}

func (p SwigcptrBar) SwigGetFoo() Foo {
	return SwigcptrFoo(p.Swigcptr())
}

type Bar interface {
	Swigcptr() uintptr
	SwigIsBar()
	Test_prepost(arg2 int, arg3 int) (_swig_ret int)
	Test_preassert(arg1 int, arg2 int) (_swig_ret int)
	Test_postassert(arg1 int) (_swig_ret int)
	SwigIsFoo()
	SwigGetFoo() Foo
}

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

var _wrap_A_foo unsafe.Pointer

func _swig_wrap_A_foo(base SwigcptrA, _ int, _ int, _ int, _ int, _ int) (_ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_A_foo, _swig_p)
	return
}

func (arg1 SwigcptrA) Foo(arg2 int, arg3 int, arg4 int, arg5 int, arg6 int) (_swig_ret int) {
	return _swig_wrap_A_foo(arg1, arg2, arg3, arg4, arg5, arg6)
}

var _wrap_new_A unsafe.Pointer

func _swig_wrap_new_A() (base SwigcptrA) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_A, _swig_p)
	return
}

func NewA() (_swig_ret A) {
	return _swig_wrap_new_A()
}

type A interface {
	Swigcptr() uintptr
	SwigIsA()
	Foo(arg2 int, arg3 int, arg4 int, arg5 int, arg6 int) (_swig_ret int)
}

type SwigcptrB uintptr

func (p SwigcptrB) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrB) SwigIsB() {
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

var _wrap_B_bar unsafe.Pointer

func _swig_wrap_B_bar(base SwigcptrB, _ int, _ int, _ int, _ int, _ int) (_ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_B_bar, _swig_p)
	return
}

func (arg1 SwigcptrB) Bar(arg2 int, arg3 int, arg4 int, arg5 int, arg6 int) (_swig_ret int) {
	return _swig_wrap_B_bar(arg1, arg2, arg3, arg4, arg5, arg6)
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

type B interface {
	Swigcptr() uintptr
	SwigIsB()
	Bar(arg2 int, arg3 int, arg4 int, arg5 int, arg6 int) (_swig_ret int)
}

type SwigcptrC uintptr

func (p SwigcptrC) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrC) SwigIsC() {
}

var _wrap_C_foo unsafe.Pointer

func _swig_wrap_C_foo(base SwigcptrC, _ int, _ int, _ int, _ int, _ int) (_ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_C_foo, _swig_p)
	return
}

func (arg1 SwigcptrC) Foo(arg2 int, arg3 int, arg4 int, arg5 int, arg6 int) (_swig_ret int) {
	return _swig_wrap_C_foo(arg1, arg2, arg3, arg4, arg5, arg6)
}

var _wrap_C_bar unsafe.Pointer

func _swig_wrap_C_bar(base SwigcptrC, _ int, _ int, _ int, _ int, _ int) (_ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_C_bar, _swig_p)
	return
}

func (arg1 SwigcptrC) Bar(arg2 int, arg3 int, arg4 int, arg5 int, arg6 int) (_swig_ret int) {
	return _swig_wrap_C_bar(arg1, arg2, arg3, arg4, arg5, arg6)
}

var _wrap_new_C unsafe.Pointer

func _swig_wrap_new_C() (base SwigcptrC) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_C, _swig_p)
	return
}

func NewC() (_swig_ret C) {
	return _swig_wrap_new_C()
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

func (p SwigcptrC) SwigIsA() {
}

func (p SwigcptrC) SwigGetA() A {
	return SwigcptrA(p.Swigcptr())
}

var _wrap_C_SwigGetB unsafe.Pointer

func _swig_wrap_C_SwigGetB(base SwigcptrC) (_ SwigcptrB) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_C_SwigGetB, _swig_p)
	return
}

func (arg1 SwigcptrC) SwigGetB() (_swig_ret B) {
	return _swig_wrap_C_SwigGetB(arg1)
}

type C interface {
	Swigcptr() uintptr
	SwigIsC()
	Foo(arg2 int, arg3 int, arg4 int, arg5 int, arg6 int) (_swig_ret int)
	Bar(arg2 int, arg3 int, arg4 int, arg5 int, arg6 int) (_swig_ret int)
	SwigIsA()
	SwigGetA() A
	SwigGetB() (_swig_ret B)
}

type SwigcptrD uintptr

func (p SwigcptrD) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrD) SwigIsD() {
}

var _wrap_D_foo unsafe.Pointer

func _swig_wrap_D_foo(base SwigcptrD, _ int, _ int, _ int, _ int, _ int) (_ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_D_foo, _swig_p)
	return
}

func (arg1 SwigcptrD) Foo(arg2 int, arg3 int, arg4 int, arg5 int, arg6 int) (_swig_ret int) {
	return _swig_wrap_D_foo(arg1, arg2, arg3, arg4, arg5, arg6)
}

var _wrap_D_bar unsafe.Pointer

func _swig_wrap_D_bar(base SwigcptrD, _ int, _ int, _ int, _ int, _ int) (_ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_D_bar, _swig_p)
	return
}

func (arg1 SwigcptrD) Bar(arg2 int, arg3 int, arg4 int, arg5 int, arg6 int) (_swig_ret int) {
	return _swig_wrap_D_bar(arg1, arg2, arg3, arg4, arg5, arg6)
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

var _wrap_delete_D unsafe.Pointer

func _swig_wrap_delete_D(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_D, _swig_p)
	return
}

func DeleteD(arg1 D) {
	_swig_wrap_delete_D(arg1.Swigcptr())
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

func (p SwigcptrD) SwigGetB() B {
	return p.SwigGetC().SwigGetB()
}

type D interface {
	Swigcptr() uintptr
	SwigIsD()
	Foo(arg2 int, arg3 int, arg4 int, arg5 int, arg6 int) (_swig_ret int)
	Bar(arg2 int, arg3 int, arg4 int, arg5 int, arg6 int) (_swig_ret int)
	SwigIsC()
	SwigGetC() C
	SwigIsA()
	SwigGetA() A
	SwigGetB() B
}

type SwigcptrE uintptr

func (p SwigcptrE) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrE) SwigIsE() {
}

var _wrap_E_m_i_set unsafe.Pointer

func _swig_wrap_E_m_i_set(base SwigcptrE, _ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_E_m_i_set, _swig_p)
	return
}

func (arg1 SwigcptrE) SetM_i(arg2 int) {
	_swig_wrap_E_m_i_set(arg1, arg2)
}

var _wrap_E_m_i_get unsafe.Pointer

func _swig_wrap_E_m_i_get(base SwigcptrE) (_ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_E_m_i_get, _swig_p)
	return
}

func (arg1 SwigcptrE) GetM_i() (_swig_ret int) {
	return _swig_wrap_E_m_i_get(arg1)
}

var _wrap_E_manipulate_i unsafe.Pointer

func _swig_wrap_E_manipulate_i(base SwigcptrE, _ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_E_manipulate_i, _swig_p)
	return
}

func (arg1 SwigcptrE) Manipulate_i(arg2 int) {
	_swig_wrap_E_manipulate_i(arg1, arg2)
}

var _wrap_new_E unsafe.Pointer

func _swig_wrap_new_E() (base SwigcptrE) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_E, _swig_p)
	return
}

func NewE() (_swig_ret E) {
	return _swig_wrap_new_E()
}

var _wrap_delete_E unsafe.Pointer

func _swig_wrap_delete_E(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_E, _swig_p)
	return
}

func DeleteE(arg1 E) {
	_swig_wrap_delete_E(arg1.Swigcptr())
}

type E interface {
	Swigcptr() uintptr
	SwigIsE()
	SetM_i(arg2 int)
	GetM_i() (_swig_ret int)
	Manipulate_i(arg2 int)
}

type SwigcptrMyClass uintptr

func (p SwigcptrMyClass) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrMyClass) SwigIsMyClass() {
}

var _wrap_new_myClass unsafe.Pointer

func _swig_wrap_new_myClass(base int) (_ SwigcptrMyClass) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_myClass, _swig_p)
	return
}

func NewMyClass(arg1 int) (_swig_ret MyClass) {
	return _swig_wrap_new_myClass(arg1)
}

var _wrap_delete_myClass unsafe.Pointer

func _swig_wrap_delete_myClass(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_myClass, _swig_p)
	return
}

func DeleteMyClass(arg1 MyClass) {
	_swig_wrap_delete_myClass(arg1.Swigcptr())
}

type MyClass interface {
	Swigcptr() uintptr
	SwigIsMyClass()
}


type SwigcptrVoid uintptr
type Void interface {
	Swigcptr() uintptr;
}
func (p SwigcptrVoid) Swigcptr() uintptr {
	return uintptr(p)
}

