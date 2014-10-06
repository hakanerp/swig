/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 3.0.3
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: ./../array_member.i

package array_member

import "unsafe"
import _ "runtime/cgo"

var _cgo_runtime_cgocall func(unsafe.Pointer, uintptr)



type _ unsafe.Pointer



type _swig_fnptr *byte
type _swig_memberptr *byte

type SwigcptrFoo uintptr

func (p SwigcptrFoo) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrFoo) SwigIsFoo() {
}

var _wrap_Foo_text_set unsafe.Pointer

func _swig_wrap_Foo_text_set(base SwigcptrFoo, _ string) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_Foo_text_set, _swig_p)
	return
}

func (arg1 SwigcptrFoo) SetText(arg2 string) {
	_swig_wrap_Foo_text_set(arg1, arg2)
}

var _wrap_Foo_text_get unsafe.Pointer

func _swig_wrap_Foo_text_get(base SwigcptrFoo) (_ string) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_Foo_text_get, _swig_p)
	return
}

func (arg1 SwigcptrFoo) GetText() (_swig_ret string) {
	return _swig_wrap_Foo_text_get(arg1)
}

var _wrap_Foo_data_set unsafe.Pointer

func _swig_wrap_Foo_data_set(base SwigcptrFoo, _ *int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_Foo_data_set, _swig_p)
	return
}

func (arg1 SwigcptrFoo) SetData(arg2 *int) {
	_swig_wrap_Foo_data_set(arg1, arg2)
}

var _wrap_Foo_data_get unsafe.Pointer

func _swig_wrap_Foo_data_get(base SwigcptrFoo) (_ *int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_Foo_data_get, _swig_p)
	return
}

func (arg1 SwigcptrFoo) GetData() (_swig_ret *int) {
	return _swig_wrap_Foo_data_get(arg1)
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

var _wrap_delete_Foo unsafe.Pointer

func _swig_wrap_delete_Foo(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_Foo, _swig_p)
	return
}

func DeleteFoo(arg1 Foo) {
	_swig_wrap_delete_Foo(arg1.Swigcptr())
}

type Foo interface {
	Swigcptr() uintptr
	SwigIsFoo()
	SetText(arg2 string)
	GetText() (_swig_ret string)
	SetData(arg2 *int)
	GetData() (_swig_ret *int)
}

var _wrap_global_data_set unsafe.Pointer

func _swig_wrap_global_data_set(base *int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_global_data_set, _swig_p)
	return
}

func SetGlobal_data(arg1 *int) {
	_swig_wrap_global_data_set(arg1)
}

var _wrap_global_data_get unsafe.Pointer

func GetGlobal_data() (_swig_ret *int) {
	_swig_p := uintptr(unsafe.Pointer(&_swig_ret))
	_cgo_runtime_cgocall(_wrap_global_data_get, _swig_p)
	return
}
var _wrap_set_value unsafe.Pointer

func _swig_wrap_set_value(base *int, _ int, _ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_set_value, _swig_p)
	return
}

func Set_value(arg1 *int, arg2 int, arg3 int) {
	_swig_wrap_set_value(arg1, arg2, arg3)
}

var _wrap_get_value unsafe.Pointer

func Get_value(arg1 *int, arg2 int) (_swig_ret int) {
	_swig_p := uintptr(unsafe.Pointer(&arg1))
	_cgo_runtime_cgocall(_wrap_get_value, _swig_p)
	return
}
type SwigcptrMaterial uintptr

func (p SwigcptrMaterial) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrMaterial) SwigIsMaterial() {
}

var _wrap_new_Material unsafe.Pointer

func _swig_wrap_new_Material() (base SwigcptrMaterial) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_Material, _swig_p)
	return
}

func NewMaterial() (_swig_ret Material) {
	return _swig_wrap_new_Material()
}

var _wrap_delete_Material unsafe.Pointer

func _swig_wrap_delete_Material(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_Material, _swig_p)
	return
}

func DeleteMaterial(arg1 Material) {
	_swig_wrap_delete_Material(arg1.Swigcptr())
}

type Material interface {
	Swigcptr() uintptr
	SwigIsMaterial()
}

type SwigcptrRayPacketData uintptr

func (p SwigcptrRayPacketData) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrRayPacketData) SwigIsRayPacketData() {
}

var _wrap_Size_RayPacketData unsafe.Pointer

func _swig_getRayPacketData_Size_RayPacketData() (_swig_ret int) {
	_swig_p := uintptr(unsafe.Pointer(&_swig_ret))
	_cgo_runtime_cgocall(_wrap_Size_RayPacketData, _swig_p)
	return
}
var RayPacketDataSize int = _swig_getRayPacketData_Size_RayPacketData()
var _wrap_RayPacketData_chitMat_set unsafe.Pointer

func _swig_wrap_RayPacketData_chitMat_set(base SwigcptrRayPacketData, _ uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_RayPacketData_chitMat_set, _swig_p)
	return
}

func (arg1 SwigcptrRayPacketData) SetChitMat(arg2 Material) {
	_swig_wrap_RayPacketData_chitMat_set(arg1, arg2.Swigcptr())
}

var _wrap_RayPacketData_chitMat_get unsafe.Pointer

func _swig_wrap_RayPacketData_chitMat_get(base SwigcptrRayPacketData) (_ SwigcptrMaterial) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_RayPacketData_chitMat_get, _swig_p)
	return
}

func (arg1 SwigcptrRayPacketData) GetChitMat() (_swig_ret Material) {
	return _swig_wrap_RayPacketData_chitMat_get(arg1)
}

var _wrap_RayPacketData_hitMat_val_set unsafe.Pointer

func _swig_wrap_RayPacketData_hitMat_val_set(base SwigcptrRayPacketData, _ uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_RayPacketData_hitMat_val_set, _swig_p)
	return
}

func (arg1 SwigcptrRayPacketData) SetHitMat_val(arg2 Material) {
	_swig_wrap_RayPacketData_hitMat_val_set(arg1, arg2.Swigcptr())
}

var _wrap_RayPacketData_hitMat_val_get unsafe.Pointer

func _swig_wrap_RayPacketData_hitMat_val_get(base SwigcptrRayPacketData) (_ SwigcptrMaterial) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_RayPacketData_hitMat_val_get, _swig_p)
	return
}

func (arg1 SwigcptrRayPacketData) GetHitMat_val() (_swig_ret Material) {
	return _swig_wrap_RayPacketData_hitMat_val_get(arg1)
}

var _wrap_RayPacketData_hitMat_set unsafe.Pointer

func _swig_wrap_RayPacketData_hitMat_set(base SwigcptrRayPacketData, _ uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_RayPacketData_hitMat_set, _swig_p)
	return
}

func (arg1 SwigcptrRayPacketData) SetHitMat(arg2 Material) {
	_swig_wrap_RayPacketData_hitMat_set(arg1, arg2.Swigcptr())
}

var _wrap_RayPacketData_hitMat_get unsafe.Pointer

func _swig_wrap_RayPacketData_hitMat_get(base SwigcptrRayPacketData) (_ SwigcptrMaterial) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_RayPacketData_hitMat_get, _swig_p)
	return
}

func (arg1 SwigcptrRayPacketData) GetHitMat() (_swig_ret Material) {
	return _swig_wrap_RayPacketData_hitMat_get(arg1)
}

var _wrap_RayPacketData_chitMat2_set unsafe.Pointer

func _swig_wrap_RayPacketData_chitMat2_set(base SwigcptrRayPacketData, _ uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_RayPacketData_chitMat2_set, _swig_p)
	return
}

func (arg1 SwigcptrRayPacketData) SetChitMat2(arg2 Material) {
	_swig_wrap_RayPacketData_chitMat2_set(arg1, arg2.Swigcptr())
}

var _wrap_RayPacketData_chitMat2_get unsafe.Pointer

func _swig_wrap_RayPacketData_chitMat2_get(base SwigcptrRayPacketData) (_ SwigcptrMaterial) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_RayPacketData_chitMat2_get, _swig_p)
	return
}

func (arg1 SwigcptrRayPacketData) GetChitMat2() (_swig_ret Material) {
	return _swig_wrap_RayPacketData_chitMat2_get(arg1)
}

var _wrap_RayPacketData_hitMat_val2_set unsafe.Pointer

func _swig_wrap_RayPacketData_hitMat_val2_set(base SwigcptrRayPacketData, _ uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_RayPacketData_hitMat_val2_set, _swig_p)
	return
}

func (arg1 SwigcptrRayPacketData) SetHitMat_val2(arg2 Material) {
	_swig_wrap_RayPacketData_hitMat_val2_set(arg1, arg2.Swigcptr())
}

var _wrap_RayPacketData_hitMat_val2_get unsafe.Pointer

func _swig_wrap_RayPacketData_hitMat_val2_get(base SwigcptrRayPacketData) (_ SwigcptrMaterial) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_RayPacketData_hitMat_val2_get, _swig_p)
	return
}

func (arg1 SwigcptrRayPacketData) GetHitMat_val2() (_swig_ret Material) {
	return _swig_wrap_RayPacketData_hitMat_val2_get(arg1)
}

var _wrap_RayPacketData_hitMat2_set unsafe.Pointer

func _swig_wrap_RayPacketData_hitMat2_set(base SwigcptrRayPacketData, _ uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_RayPacketData_hitMat2_set, _swig_p)
	return
}

func (arg1 SwigcptrRayPacketData) SetHitMat2(arg2 Material) {
	_swig_wrap_RayPacketData_hitMat2_set(arg1, arg2.Swigcptr())
}

var _wrap_RayPacketData_hitMat2_get unsafe.Pointer

func _swig_wrap_RayPacketData_hitMat2_get(base SwigcptrRayPacketData) (_ SwigcptrMaterial) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_RayPacketData_hitMat2_get, _swig_p)
	return
}

func (arg1 SwigcptrRayPacketData) GetHitMat2() (_swig_ret Material) {
	return _swig_wrap_RayPacketData_hitMat2_get(arg1)
}

var _wrap_new_RayPacketData unsafe.Pointer

func _swig_wrap_new_RayPacketData() (base SwigcptrRayPacketData) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_RayPacketData, _swig_p)
	return
}

func NewRayPacketData() (_swig_ret RayPacketData) {
	return _swig_wrap_new_RayPacketData()
}

var _wrap_delete_RayPacketData unsafe.Pointer

func _swig_wrap_delete_RayPacketData(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_RayPacketData, _swig_p)
	return
}

func DeleteRayPacketData(arg1 RayPacketData) {
	_swig_wrap_delete_RayPacketData(arg1.Swigcptr())
}

type RayPacketData interface {
	Swigcptr() uintptr
	SwigIsRayPacketData()
	SetChitMat(arg2 Material)
	GetChitMat() (_swig_ret Material)
	SetHitMat_val(arg2 Material)
	GetHitMat_val() (_swig_ret Material)
	SetHitMat(arg2 Material)
	GetHitMat() (_swig_ret Material)
	SetChitMat2(arg2 Material)
	GetChitMat2() (_swig_ret Material)
	SetHitMat_val2(arg2 Material)
	GetHitMat_val2() (_swig_ret Material)
	SetHitMat2(arg2 Material)
	GetHitMat2() (_swig_ret Material)
}

const BUFF_LEN int = 12
type SwigcptrMyBuff uintptr

func (p SwigcptrMyBuff) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrMyBuff) SwigIsMyBuff() {
}

var _wrap_MyBuff_i_set unsafe.Pointer

func _swig_wrap_MyBuff_i_set(base SwigcptrMyBuff, _ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_MyBuff_i_set, _swig_p)
	return
}

func (arg1 SwigcptrMyBuff) SetI(arg2 int) {
	_swig_wrap_MyBuff_i_set(arg1, arg2)
}

var _wrap_MyBuff_i_get unsafe.Pointer

func _swig_wrap_MyBuff_i_get(base SwigcptrMyBuff) (_ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_MyBuff_i_get, _swig_p)
	return
}

func (arg1 SwigcptrMyBuff) GetI() (_swig_ret int) {
	return _swig_wrap_MyBuff_i_get(arg1)
}

var _wrap_MyBuff_x_set unsafe.Pointer

func _swig_wrap_MyBuff_x_set(base SwigcptrMyBuff, _ string) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_MyBuff_x_set, _swig_p)
	return
}

func (arg1 SwigcptrMyBuff) SetX(arg2 string) {
	_swig_wrap_MyBuff_x_set(arg1, arg2)
}

var _wrap_MyBuff_x_get unsafe.Pointer

func _swig_wrap_MyBuff_x_get(base SwigcptrMyBuff) (_ string) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_MyBuff_x_get, _swig_p)
	return
}

func (arg1 SwigcptrMyBuff) GetX() (_swig_ret string) {
	return _swig_wrap_MyBuff_x_get(arg1)
}

var _wrap_new_MyBuff unsafe.Pointer

func _swig_wrap_new_MyBuff() (base SwigcptrMyBuff) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_MyBuff, _swig_p)
	return
}

func NewMyBuff() (_swig_ret MyBuff) {
	return _swig_wrap_new_MyBuff()
}

var _wrap_delete_MyBuff unsafe.Pointer

func _swig_wrap_delete_MyBuff(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_MyBuff, _swig_p)
	return
}

func DeleteMyBuff(arg1 MyBuff) {
	_swig_wrap_delete_MyBuff(arg1.Swigcptr())
}

type MyBuff interface {
	Swigcptr() uintptr
	SwigIsMyBuff()
	SetI(arg2 int)
	GetI() (_swig_ret int)
	SetX(arg2 string)
	GetX() (_swig_ret string)
}

type SwigcptrMySBuff uintptr

func (p SwigcptrMySBuff) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrMySBuff) SwigIsMySBuff() {
}

var _wrap_MySBuff_i_set unsafe.Pointer

func _swig_wrap_MySBuff_i_set(base SwigcptrMySBuff, _ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_MySBuff_i_set, _swig_p)
	return
}

func (arg1 SwigcptrMySBuff) SetI(arg2 int) {
	_swig_wrap_MySBuff_i_set(arg1, arg2)
}

var _wrap_MySBuff_i_get unsafe.Pointer

func _swig_wrap_MySBuff_i_get(base SwigcptrMySBuff) (_ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_MySBuff_i_get, _swig_p)
	return
}

func (arg1 SwigcptrMySBuff) GetI() (_swig_ret int) {
	return _swig_wrap_MySBuff_i_get(arg1)
}

var _wrap_MySBuff_x_set unsafe.Pointer

func _swig_wrap_MySBuff_x_set(base SwigcptrMySBuff, _ string) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_MySBuff_x_set, _swig_p)
	return
}

func (arg1 SwigcptrMySBuff) SetX(arg2 string) {
	_swig_wrap_MySBuff_x_set(arg1, arg2)
}

var _wrap_MySBuff_x_get unsafe.Pointer

func _swig_wrap_MySBuff_x_get(base SwigcptrMySBuff) (_ string) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_MySBuff_x_get, _swig_p)
	return
}

func (arg1 SwigcptrMySBuff) GetX() (_swig_ret string) {
	return _swig_wrap_MySBuff_x_get(arg1)
}

var _wrap_new_MySBuff unsafe.Pointer

func _swig_wrap_new_MySBuff() (base SwigcptrMySBuff) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_MySBuff, _swig_p)
	return
}

func NewMySBuff() (_swig_ret MySBuff) {
	return _swig_wrap_new_MySBuff()
}

var _wrap_delete_MySBuff unsafe.Pointer

func _swig_wrap_delete_MySBuff(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_MySBuff, _swig_p)
	return
}

func DeleteMySBuff(arg1 MySBuff) {
	_swig_wrap_delete_MySBuff(arg1.Swigcptr())
}

type MySBuff interface {
	Swigcptr() uintptr
	SwigIsMySBuff()
	SetI(arg2 int)
	GetI() (_swig_ret int)
	SetX(arg2 string)
	GetX() (_swig_ret string)
}


type SwigcptrVoid uintptr
type Void interface {
	Swigcptr() uintptr;
}
func (p SwigcptrVoid) Swigcptr() uintptr {
	return uintptr(p)
}

