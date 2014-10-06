/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 3.0.3
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: ./../li_std_vector_member_var.i

package li_std_vector_member_var

import "unsafe"
import _ "runtime/cgo"

var _cgo_runtime_cgocall func(unsafe.Pointer, uintptr)



type _ unsafe.Pointer



type _swig_fnptr *byte
type _swig_memberptr *byte

type SwigcptrVectorDbl uintptr

func (p SwigcptrVectorDbl) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrVectorDbl) SwigIsVectorDbl() {
}

var _wrap_new_vectorDbl__SWIG_0 unsafe.Pointer

func _swig_wrap_new_vectorDbl__SWIG_0() (base SwigcptrVectorDbl) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_vectorDbl__SWIG_0, _swig_p)
	return
}

func NewVectorDbl__SWIG_0() (_swig_ret VectorDbl) {
	return _swig_wrap_new_vectorDbl__SWIG_0()
}

var _wrap_new_vectorDbl__SWIG_1 unsafe.Pointer

func _swig_wrap_new_vectorDbl__SWIG_1(base int64) (_ SwigcptrVectorDbl) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_vectorDbl__SWIG_1, _swig_p)
	return
}

func NewVectorDbl__SWIG_1(arg1 int64) (_swig_ret VectorDbl) {
	return _swig_wrap_new_vectorDbl__SWIG_1(arg1)
}

func NewVectorDbl(a ...interface{}) VectorDbl {
	argc := len(a)
	if argc == 0 {
		return NewVectorDbl__SWIG_0()
	}
	if argc == 1 {
		return NewVectorDbl__SWIG_1(a[0].(int64))
	}
	panic("No match for overloaded function call")
}

var _wrap_vectorDbl_size unsafe.Pointer

func _swig_wrap_vectorDbl_size(base SwigcptrVectorDbl) (_ int64) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_vectorDbl_size, _swig_p)
	return
}

func (arg1 SwigcptrVectorDbl) Size() (_swig_ret int64) {
	return _swig_wrap_vectorDbl_size(arg1)
}

var _wrap_vectorDbl_capacity unsafe.Pointer

func _swig_wrap_vectorDbl_capacity(base SwigcptrVectorDbl) (_ int64) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_vectorDbl_capacity, _swig_p)
	return
}

func (arg1 SwigcptrVectorDbl) Capacity() (_swig_ret int64) {
	return _swig_wrap_vectorDbl_capacity(arg1)
}

var _wrap_vectorDbl_reserve unsafe.Pointer

func _swig_wrap_vectorDbl_reserve(base SwigcptrVectorDbl, _ int64) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_vectorDbl_reserve, _swig_p)
	return
}

func (arg1 SwigcptrVectorDbl) Reserve(arg2 int64) {
	_swig_wrap_vectorDbl_reserve(arg1, arg2)
}

var _wrap_vectorDbl_isEmpty unsafe.Pointer

func _swig_wrap_vectorDbl_isEmpty(base SwigcptrVectorDbl) (_ bool) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_vectorDbl_isEmpty, _swig_p)
	return
}

func (arg1 SwigcptrVectorDbl) IsEmpty() (_swig_ret bool) {
	return _swig_wrap_vectorDbl_isEmpty(arg1)
}

var _wrap_vectorDbl_clear unsafe.Pointer

func _swig_wrap_vectorDbl_clear(base SwigcptrVectorDbl) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_vectorDbl_clear, _swig_p)
	return
}

func (arg1 SwigcptrVectorDbl) Clear() {
	_swig_wrap_vectorDbl_clear(arg1)
}

var _wrap_vectorDbl_add unsafe.Pointer

func _swig_wrap_vectorDbl_add(base SwigcptrVectorDbl, _ float64) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_vectorDbl_add, _swig_p)
	return
}

func (arg1 SwigcptrVectorDbl) Add(arg2 float64) {
	_swig_wrap_vectorDbl_add(arg1, arg2)
}

var _wrap_vectorDbl_get unsafe.Pointer

func _swig_wrap_vectorDbl_get(base SwigcptrVectorDbl, _ int) (_ float64) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_vectorDbl_get, _swig_p)
	return
}

func (arg1 SwigcptrVectorDbl) Get(arg2 int) (_swig_ret float64) {
	return _swig_wrap_vectorDbl_get(arg1, arg2)
}

var _wrap_vectorDbl_set unsafe.Pointer

func _swig_wrap_vectorDbl_set(base SwigcptrVectorDbl, _ int, _ float64) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_vectorDbl_set, _swig_p)
	return
}

func (arg1 SwigcptrVectorDbl) Set(arg2 int, arg3 float64) {
	_swig_wrap_vectorDbl_set(arg1, arg2, arg3)
}

var _wrap_delete_vectorDbl unsafe.Pointer

func _swig_wrap_delete_vectorDbl(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_vectorDbl, _swig_p)
	return
}

func DeleteVectorDbl(arg1 VectorDbl) {
	_swig_wrap_delete_vectorDbl(arg1.Swigcptr())
}

type VectorDbl interface {
	Swigcptr() uintptr
	SwigIsVectorDbl()
	Size() (_swig_ret int64)
	Capacity() (_swig_ret int64)
	Reserve(arg2 int64)
	IsEmpty() (_swig_ret bool)
	Clear()
	Add(arg2 float64)
	Get(arg2 int) (_swig_ret float64)
	Set(arg2 int, arg3 float64)
}

type SwigcptrTest uintptr

func (p SwigcptrTest) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrTest) SwigIsTest() {
}

var _wrap_Test_v_set unsafe.Pointer

func _swig_wrap_Test_v_set(base SwigcptrTest, _ uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_Test_v_set, _swig_p)
	return
}

func (arg1 SwigcptrTest) SetV(arg2 VectorDbl) {
	_swig_wrap_Test_v_set(arg1, arg2.Swigcptr())
}

var _wrap_Test_v_get unsafe.Pointer

func _swig_wrap_Test_v_get(base SwigcptrTest) (_ SwigcptrVectorDbl) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_Test_v_get, _swig_p)
	return
}

func (arg1 SwigcptrTest) GetV() (_swig_ret VectorDbl) {
	return _swig_wrap_Test_v_get(arg1)
}

var _wrap_Test_x_set unsafe.Pointer

func _swig_wrap_Test_x_set(base SwigcptrTest, _ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_Test_x_set, _swig_p)
	return
}

func (arg1 SwigcptrTest) SetX(arg2 int) {
	_swig_wrap_Test_x_set(arg1, arg2)
}

var _wrap_Test_x_get unsafe.Pointer

func _swig_wrap_Test_x_get(base SwigcptrTest) (_ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_Test_x_get, _swig_p)
	return
}

func (arg1 SwigcptrTest) GetX() (_swig_ret int) {
	return _swig_wrap_Test_x_get(arg1)
}

var _wrap_new_Test unsafe.Pointer

func _swig_wrap_new_Test() (base SwigcptrTest) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_Test, _swig_p)
	return
}

func NewTest() (_swig_ret Test) {
	return _swig_wrap_new_Test()
}

var _wrap_Test_f unsafe.Pointer

func _swig_wrap_Test_f(base SwigcptrTest, _ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_Test_f, _swig_p)
	return
}

func (arg1 SwigcptrTest) F(arg2 int) {
	_swig_wrap_Test_f(arg1, arg2)
}

var _wrap_delete_Test unsafe.Pointer

func _swig_wrap_delete_Test(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_Test, _swig_p)
	return
}

func DeleteTest(arg1 Test) {
	_swig_wrap_delete_Test(arg1.Swigcptr())
}

type Test interface {
	Swigcptr() uintptr
	SwigIsTest()
	SetV(arg2 VectorDbl)
	GetV() (_swig_ret VectorDbl)
	SetX(arg2 int)
	GetX() (_swig_ret int)
	F(arg2 int)
}

type SwigcptrS uintptr

func (p SwigcptrS) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrS) SwigIsS() {
}

var _wrap_S_x_set unsafe.Pointer

func _swig_wrap_S_x_set(base SwigcptrS, _ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_S_x_set, _swig_p)
	return
}

func (arg1 SwigcptrS) SetX(arg2 int) {
	_swig_wrap_S_x_set(arg1, arg2)
}

var _wrap_S_x_get unsafe.Pointer

func _swig_wrap_S_x_get(base SwigcptrS) (_ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_S_x_get, _swig_p)
	return
}

func (arg1 SwigcptrS) GetX() (_swig_ret int) {
	return _swig_wrap_S_x_get(arg1)
}

var _wrap_new_S unsafe.Pointer

func _swig_wrap_new_S() (base SwigcptrS) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_S, _swig_p)
	return
}

func NewS() (_swig_ret S) {
	return _swig_wrap_new_S()
}

var _wrap_delete_S unsafe.Pointer

func _swig_wrap_delete_S(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_S, _swig_p)
	return
}

func DeleteS(arg1 S) {
	_swig_wrap_delete_S(arg1.Swigcptr())
}

type S interface {
	Swigcptr() uintptr
	SwigIsS()
	SetX(arg2 int)
	GetX() (_swig_ret int)
}

type SwigcptrT uintptr

func (p SwigcptrT) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrT) SwigIsT() {
}

var _wrap_T_start_t_set unsafe.Pointer

func _swig_wrap_T_start_t_set(base SwigcptrT, _ uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_T_start_t_set, _swig_p)
	return
}

func (arg1 SwigcptrT) SetStart_t(arg2 S) {
	_swig_wrap_T_start_t_set(arg1, arg2.Swigcptr())
}

var _wrap_T_start_t_get unsafe.Pointer

func _swig_wrap_T_start_t_get(base SwigcptrT) (_ SwigcptrS) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_T_start_t_get, _swig_p)
	return
}

func (arg1 SwigcptrT) GetStart_t() (_swig_ret S) {
	return _swig_wrap_T_start_t_get(arg1)
}

var _wrap_T_length_set unsafe.Pointer

func _swig_wrap_T_length_set(base SwigcptrT, _ uint) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_T_length_set, _swig_p)
	return
}

func (arg1 SwigcptrT) SetLength(arg2 uint) {
	_swig_wrap_T_length_set(arg1, arg2)
}

var _wrap_T_length_get unsafe.Pointer

func _swig_wrap_T_length_get(base SwigcptrT) (_ uint) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_T_length_get, _swig_p)
	return
}

func (arg1 SwigcptrT) GetLength() (_swig_ret uint) {
	return _swig_wrap_T_length_get(arg1)
}

var _wrap_new_T unsafe.Pointer

func _swig_wrap_new_T() (base SwigcptrT) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_T, _swig_p)
	return
}

func NewT() (_swig_ret T) {
	return _swig_wrap_new_T()
}

var _wrap_delete_T unsafe.Pointer

func _swig_wrap_delete_T(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_T, _swig_p)
	return
}

func DeleteT(arg1 T) {
	_swig_wrap_delete_T(arg1.Swigcptr())
}

type T interface {
	Swigcptr() uintptr
	SwigIsT()
	SetStart_t(arg2 S)
	GetStart_t() (_swig_ret S)
	SetLength(arg2 uint)
	GetLength() (_swig_ret uint)
}


type SwigcptrVoid uintptr
type Void interface {
	Swigcptr() uintptr;
}
func (p SwigcptrVoid) Swigcptr() uintptr {
	return uintptr(p)
}

