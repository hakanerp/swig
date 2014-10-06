/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 3.0.3
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: ./../special_variables.i

package special_variables

import "unsafe"
import _ "runtime/cgo"

var _cgo_runtime_cgocall func(unsafe.Pointer, uintptr)



type _ unsafe.Pointer



type _swig_fnptr *byte
type _swig_memberptr *byte

var _wrap_testmethod unsafe.Pointer

func _swig_wrap_testmethod(base int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_testmethod, _swig_p)
	return
}

func Testmethod(arg1 int) {
	_swig_wrap_testmethod(arg1)
}

type SwigcptrKKK uintptr

func (p SwigcptrKKK) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrKKK) SwigIsKKK() {
}

var _wrap_KKK_testmethod unsafe.Pointer

func _swig_wrap_KKK_testmethod(base SwigcptrKKK, _ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_KKK_testmethod, _swig_p)
	return
}

func (arg1 SwigcptrKKK) Testmethod(arg2 int) {
	_swig_wrap_KKK_testmethod(arg1, arg2)
}

var _wrap_KKK_teststaticmethod unsafe.Pointer

func _swig_wrap_KKK_teststaticmethod(base int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_KKK_teststaticmethod, _swig_p)
	return
}

func KKKTeststaticmethod(arg1 int) {
	_swig_wrap_KKK_teststaticmethod(arg1)
}

var _wrap_new_KKK unsafe.Pointer

func _swig_wrap_new_KKK() (base SwigcptrKKK) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_KKK, _swig_p)
	return
}

func NewKKK() (_swig_ret KKK) {
	return _swig_wrap_new_KKK()
}

var _wrap_delete_KKK unsafe.Pointer

func _swig_wrap_delete_KKK(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_KKK, _swig_p)
	return
}

func DeleteKKK(arg1 KKK) {
	_swig_wrap_delete_KKK(arg1.Swigcptr())
}

type KKK interface {
	Swigcptr() uintptr
	SwigIsKKK()
	Testmethod(arg2 int)
}

var _wrap_ExceptionVars unsafe.Pointer

func ExceptionVars(arg1 float64, arg2 float64) (_swig_ret string) {
	_swig_p := uintptr(unsafe.Pointer(&arg1))
	_cgo_runtime_cgocall(_wrap_ExceptionVars, _swig_p)
	return
}
var _wrap_overloadedmethod__SWIG_0 unsafe.Pointer

func Overloadedmethod__SWIG_0(arg1 float64) (_swig_ret string) {
	_swig_p := uintptr(unsafe.Pointer(&arg1))
	_cgo_runtime_cgocall(_wrap_overloadedmethod__SWIG_0, _swig_p)
	return
}
var _wrap_overloadedmethod__SWIG_1 unsafe.Pointer

func Overloadedmethod__SWIG_1() (_swig_ret string) {
	_swig_p := uintptr(unsafe.Pointer(&_swig_ret))
	_cgo_runtime_cgocall(_wrap_overloadedmethod__SWIG_1, _swig_p)
	return
}
func Overloadedmethod(a ...interface{}) string {
	argc := len(a)
	if argc == 0 {
		return Overloadedmethod__SWIG_1()
	}
	if argc == 1 {
		return Overloadedmethod__SWIG_0(a[0].(float64))
	}
	panic("No match for overloaded function call")
}

var _wrap_declaration_set unsafe.Pointer

func _swig_wrap_declaration_set(base string) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_declaration_set, _swig_p)
	return
}

func SetDeclaration(arg1 string) {
	_swig_wrap_declaration_set(arg1)
}

var _wrap_declaration_get unsafe.Pointer

func GetDeclaration() (_swig_ret string) {
	_swig_p := uintptr(unsafe.Pointer(&_swig_ret))
	_cgo_runtime_cgocall(_wrap_declaration_get, _swig_p)
	return
}
type SwigcptrABC uintptr

func (p SwigcptrABC) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrABC) SwigIsABC() {
}

var _wrap_new_ABC__SWIG_0 unsafe.Pointer

func _swig_wrap_new_ABC__SWIG_0(base int, _ float64) (_ SwigcptrABC) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_ABC__SWIG_0, _swig_p)
	return
}

func NewABC__SWIG_0(arg1 int, arg2 float64) (_swig_ret ABC) {
	return _swig_wrap_new_ABC__SWIG_0(arg1, arg2)
}

var _wrap_new_ABC__SWIG_1 unsafe.Pointer

func _swig_wrap_new_ABC__SWIG_1() (base SwigcptrABC) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_ABC__SWIG_1, _swig_p)
	return
}

func NewABC__SWIG_1() (_swig_ret ABC) {
	return _swig_wrap_new_ABC__SWIG_1()
}

func NewABC(a ...interface{}) ABC {
	argc := len(a)
	if argc == 0 {
		return NewABC__SWIG_1()
	}
	if argc == 2 {
		return NewABC__SWIG_0(a[0].(int), a[1].(float64))
	}
	panic("No match for overloaded function call")
}

var _wrap_ABC_staticmethod unsafe.Pointer

func ABCStaticmethod(arg1 int, arg2 bool) (_swig_ret *int16) {
	_swig_p := uintptr(unsafe.Pointer(&arg1))
	_cgo_runtime_cgocall(_wrap_ABC_staticmethod, _swig_p)
	return
}
var _wrap_ABC_instancemethod__SWIG_0 unsafe.Pointer

func _swig_wrap_ABC_instancemethod__SWIG_0(base SwigcptrABC, _ int, _ bool) (_ *int16) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_ABC_instancemethod__SWIG_0, _swig_p)
	return
}

func (arg1 SwigcptrABC) Instancemethod__SWIG_0(arg2 int, arg3 bool) (_swig_ret *int16) {
	return _swig_wrap_ABC_instancemethod__SWIG_0(arg1, arg2, arg3)
}

var _wrap_ABC_instancemethod__SWIG_1 unsafe.Pointer

func _swig_wrap_ABC_instancemethod__SWIG_1(base SwigcptrABC, _ int) (_ *int16) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_ABC_instancemethod__SWIG_1, _swig_p)
	return
}

func (arg1 SwigcptrABC) Instancemethod__SWIG_1(arg2 int) (_swig_ret *int16) {
	return _swig_wrap_ABC_instancemethod__SWIG_1(arg1, arg2)
}

func (p SwigcptrABC) Instancemethod(a ...interface{}) *int16 {
	argc := len(a)
	if argc == 1 {
		return p.Instancemethod__SWIG_1(a[0].(int))
	}
	if argc == 2 {
		return p.Instancemethod__SWIG_0(a[0].(int), a[1].(bool))
	}
	panic("No match for overloaded function call")
}

var _wrap_ABC_constmethod unsafe.Pointer

func _swig_wrap_ABC_constmethod(base SwigcptrABC, _ int) (_ *int16) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_ABC_constmethod, _swig_p)
	return
}

func (arg1 SwigcptrABC) Constmethod(arg2 int) (_swig_ret *int16) {
	return _swig_wrap_ABC_constmethod(arg1, arg2)
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
	Instancemethod(a ...interface{}) *int16
	Constmethod(arg2 int) (_swig_ret *int16)
}

var _wrap_globtemplate unsafe.Pointer

func _swig_wrap_globtemplate(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_globtemplate, _swig_p)
	return
}

func Globtemplate(arg1 TemplateABC) {
	_swig_wrap_globtemplate(arg1.Swigcptr())
}

type SwigcptrTemplateABC uintptr

func (p SwigcptrTemplateABC) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrTemplateABC) SwigIsTemplateABC() {
}

var _wrap_TemplateABC_tmethod unsafe.Pointer

func _swig_wrap_TemplateABC_tmethod(base SwigcptrTemplateABC, _ uintptr) (_ string) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_TemplateABC_tmethod, _swig_p)
	return
}

func (arg1 SwigcptrTemplateABC) Tmethod(arg2 ABC) (_swig_ret string) {
	return _swig_wrap_TemplateABC_tmethod(arg1, arg2.Swigcptr())
}

var _wrap_new_TemplateABC unsafe.Pointer

func _swig_wrap_new_TemplateABC() (base SwigcptrTemplateABC) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_TemplateABC, _swig_p)
	return
}

func NewTemplateABC() (_swig_ret TemplateABC) {
	return _swig_wrap_new_TemplateABC()
}

var _wrap_delete_TemplateABC unsafe.Pointer

func _swig_wrap_delete_TemplateABC(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_TemplateABC, _swig_p)
	return
}

func DeleteTemplateABC(arg1 TemplateABC) {
	_swig_wrap_delete_TemplateABC(arg1.Swigcptr())
}

type TemplateABC interface {
	Swigcptr() uintptr
	SwigIsTemplateABC()
	Tmethod(arg2 ABC) (_swig_ret string)
}

var _wrap_director_testmethod unsafe.Pointer

func _swig_wrap_director_testmethod(base int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_director_testmethod, _swig_p)
	return
}

func Director_testmethod(arg1 int) {
	_swig_wrap_director_testmethod(arg1)
}

type _swig_DirectorDirectorTest struct {
	SwigcptrDirectorTest
	v interface{}
}

func (p *_swig_DirectorDirectorTest) Swigcptr() uintptr {
	return p.SwigcptrDirectorTest.Swigcptr()
}

func (p *_swig_DirectorDirectorTest) SwigIsDirectorTest() {
}

func (p *_swig_DirectorDirectorTest) DirectorInterface() interface{} {
	return p.v
}

var _wrap__swig_NewDirectorDirectorTestDirectorTest unsafe.Pointer

func _swig_NewDirectorDirectorTestDirectorTest(_swig_director *_swig_DirectorDirectorTest) (_swig_ret SwigcptrDirectorTest) {
	_swig_p := uintptr(unsafe.Pointer(&_swig_director))
	_cgo_runtime_cgocall(_wrap__swig_NewDirectorDirectorTestDirectorTest, _swig_p)
	return
}

func NewDirectorDirectorTest(v interface{}) DirectorTest {
	p := &_swig_DirectorDirectorTest{0, v}
	p.SwigcptrDirectorTest = _swig_NewDirectorDirectorTestDirectorTest(p)
	return p
}

type _swig_DirectorInterfaceDirectorTestDirector_testmethod interface {
	Director_testmethod(int)
}

var _wrap__swig_DirectorDirectorTest_upcall_Director_testmethod unsafe.Pointer

func _swig_wrap__swig_DirectorDirectorTest_upcall_Director_testmethod(_swig_ptr SwigcptrDirectorTest, _ int) {
	_swig_p := uintptr(unsafe.Pointer(&_swig_ptr))
	_cgo_runtime_cgocall(_wrap__swig_DirectorDirectorTest_upcall_Director_testmethod, _swig_p)
	return
}

func (swig_p *_swig_DirectorDirectorTest) Director_testmethod(i int) {
	if swig_g, swig_ok := swig_p.v.(_swig_DirectorInterfaceDirectorTestDirector_testmethod); swig_ok {
		swig_g.Director_testmethod(i)
		return
	}
	_swig_wrap__swig_DirectorDirectorTest_upcall_Director_testmethod(swig_p.SwigcptrDirectorTest, i)
}

func DirectorDirectorTestDirector_testmethod(p DirectorTest, arg2 int) {
	_swig_wrap__swig_DirectorDirectorTest_upcall_Director_testmethod(p.(*_swig_DirectorDirectorTest).SwigcptrDirectorTest, arg2)
}

func Swig_DirectorDirectorTest_callback_director_testmethod(p *_swig_DirectorDirectorTest, arg2 int) {
	p.Director_testmethod(arg2)
}

var _wrap_DeleteDirectorDirectorTest unsafe.Pointer

func _swig_wrap_DeleteDirectorDirectorTest(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_DeleteDirectorDirectorTest, _swig_p)
	return
}

func DeleteDirectorDirectorTest(arg1 DirectorTest) {
	_swig_wrap_DeleteDirectorDirectorTest(arg1.Swigcptr())
}

func Swiggo_DeleteDirector_DirectorTest(p *_swig_DirectorDirectorTest) {
	p.SwigcptrDirectorTest = 0
}

type SwigcptrDirectorTest uintptr

func (p SwigcptrDirectorTest) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrDirectorTest) SwigIsDirectorTest() {
}

func (p SwigcptrDirectorTest) DirectorInterface() interface{} {
	return nil
}

var _wrap_DirectorTest_director_testmethod unsafe.Pointer

func _swig_wrap_DirectorTest_director_testmethod(base SwigcptrDirectorTest, _ int) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_DirectorTest_director_testmethod, _swig_p)
	return
}

func (arg1 SwigcptrDirectorTest) Director_testmethod(arg2 int) {
	_swig_wrap_DirectorTest_director_testmethod(arg1, arg2)
}

var _wrap_delete_DirectorTest unsafe.Pointer

func _swig_wrap_delete_DirectorTest(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_DirectorTest, _swig_p)
	return
}

func DeleteDirectorTest(arg1 DirectorTest) {
	_swig_wrap_delete_DirectorTest(arg1.Swigcptr())
}

var _wrap_new_DirectorTest unsafe.Pointer

func _swig_wrap_new_DirectorTest() (base SwigcptrDirectorTest) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_DirectorTest, _swig_p)
	return
}

func NewDirectorTest() (_swig_ret DirectorTest) {
	return _swig_wrap_new_DirectorTest()
}

type DirectorTest interface {
	Swigcptr() uintptr
	SwigIsDirectorTest()
	DirectorInterface() interface{}
	Director_testmethod(arg2 int)
}

type SwigcptrDEFNewName uintptr

func (p SwigcptrDEFNewName) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrDEFNewName) SwigIsDEFNewName() {
}

var _wrap_DEFNewName_instance_def unsafe.Pointer

func _swig_wrap_DEFNewName_instance_def(base SwigcptrDEFNewName) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_DEFNewName_instance_def, _swig_p)
	return
}

func (arg1 SwigcptrDEFNewName) Instance_def() {
	_swig_wrap_DEFNewName_instance_def(arg1)
}

var _wrap_DEFNewName_static_def unsafe.Pointer

func _swig_wrap_DEFNewName_static_def() {
	var _swig_p uintptr
	_cgo_runtime_cgocall(_wrap_DEFNewName_static_def, _swig_p)
	return
}

func DEFNewNameStatic_def() {
	_swig_wrap_DEFNewName_static_def()
}

var _wrap_new_DEFNewName unsafe.Pointer

func _swig_wrap_new_DEFNewName() (base SwigcptrDEFNewName) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_new_DEFNewName, _swig_p)
	return
}

func NewDEFNewName() (_swig_ret DEFNewName) {
	return _swig_wrap_new_DEFNewName()
}

var _wrap_delete_DEFNewName unsafe.Pointer

func _swig_wrap_delete_DEFNewName(base uintptr) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_delete_DEFNewName, _swig_p)
	return
}

func DeleteDEFNewName(arg1 DEFNewName) {
	_swig_wrap_delete_DEFNewName(arg1.Swigcptr())
}

var _wrap_DEFNewName_staticmethod unsafe.Pointer

func DEFNewNameStaticmethod(arg1 int, arg2 bool) (_swig_ret *int16) {
	_swig_p := uintptr(unsafe.Pointer(&arg1))
	_cgo_runtime_cgocall(_wrap_DEFNewName_staticmethod, _swig_p)
	return
}
var _wrap_DEFNewName_instancemethod__SWIG_0 unsafe.Pointer

func _swig_wrap_DEFNewName_instancemethod__SWIG_0(base SwigcptrDEFNewName, _ int, _ bool) (_ *int16) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_DEFNewName_instancemethod__SWIG_0, _swig_p)
	return
}

func (_swig_base SwigcptrDEFNewName) Instancemethod__SWIG_0(arg1 int, arg2 bool) (_swig_ret *int16) {
	return _swig_wrap_DEFNewName_instancemethod__SWIG_0(_swig_base, arg1, arg2)
}

var _wrap_DEFNewName_instancemethod__SWIG_1 unsafe.Pointer

func _swig_wrap_DEFNewName_instancemethod__SWIG_1(base SwigcptrDEFNewName, _ int) (_ *int16) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_DEFNewName_instancemethod__SWIG_1, _swig_p)
	return
}

func (_swig_base SwigcptrDEFNewName) Instancemethod__SWIG_1(arg1 int) (_swig_ret *int16) {
	return _swig_wrap_DEFNewName_instancemethod__SWIG_1(_swig_base, arg1)
}

func (p SwigcptrDEFNewName) Instancemethod(a ...interface{}) *int16 {
	argc := len(a)
	if argc == 1 {
		return p.Instancemethod__SWIG_1(a[0].(int))
	}
	if argc == 2 {
		return p.Instancemethod__SWIG_0(a[0].(int), a[1].(bool))
	}
	panic("No match for overloaded function call")
}

var _wrap_DEFNewName_constmethod unsafe.Pointer

func _swig_wrap_DEFNewName_constmethod(base SwigcptrDEFNewName, _ int) (_ *int16) {
	_swig_p := uintptr(unsafe.Pointer(&base))
	_cgo_runtime_cgocall(_wrap_DEFNewName_constmethod, _swig_p)
	return
}

func (_swig_base SwigcptrDEFNewName) Constmethod(arg1 int) (_swig_ret *int16) {
	return _swig_wrap_DEFNewName_constmethod(_swig_base, arg1)
}

func (p SwigcptrDEFNewName) SwigIsABC() {
}

func (p SwigcptrDEFNewName) SwigGetABC() ABC {
	return SwigcptrABC(p.Swigcptr())
}

type DEFNewName interface {
	Swigcptr() uintptr
	SwigIsDEFNewName()
	Instance_def()
	Instancemethod(a ...interface{}) *int16
	Constmethod(arg1 int) (_swig_ret *int16)
	SwigIsABC()
	SwigGetABC() ABC
}


type SwigcptrSwigDirector_DirectorTest uintptr
type SwigDirector_DirectorTest interface {
	Swigcptr() uintptr;
}
func (p SwigcptrSwigDirector_DirectorTest) Swigcptr() uintptr {
	return uintptr(p)
}

type SwigcptrVoid uintptr
type Void interface {
	Swigcptr() uintptr;
}
func (p SwigcptrVoid) Swigcptr() uintptr {
	return uintptr(p)
}

