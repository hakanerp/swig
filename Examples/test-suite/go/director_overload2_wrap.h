/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 3.0.3
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: ./../director_overload2.i

#ifndef SWIG_director_overload2_WRAP_H_
#define SWIG_director_overload2_WRAP_H_

class SwigDirector_OverloadBase : public OverloadBase
{
 public:
  SwigDirector_OverloadBase(void *swig_p);
  virtual ~SwigDirector_OverloadBase();
  void _swig_upcall_mmm() {
    OverloadBase::mmm();
  }
  virtual void mmm();
  void _swig_upcall_nnn__SWIG_0(int vvv) {
    OverloadBase::nnn(vvv);
  }
  virtual void nnn(int vvv);
  void _swig_upcall_nnn__SWIG_1() {
    OverloadBase::nnn();
  }
  virtual void nnn();
 private:
  void *go_val;
};

class SwigDirector_OverloadDerived1 : public OverloadDerived1
{
 public:
  SwigDirector_OverloadDerived1(void *swig_p);
  virtual ~SwigDirector_OverloadDerived1();
  void _swig_upcall_mmm() {
    OverloadBase::mmm();
  }
  virtual void mmm();
  void _swig_upcall_nnn(int vvv) {
    OverloadDerived1::nnn(vvv);
  }
  virtual void nnn(int vvv);
  void _swig_upcall_nnn__SWIG_1() {
    OverloadBase::nnn();
  }
  virtual void nnn();
 private:
  void *go_val;
};

class SwigDirector_OverloadDerived2 : public OverloadDerived2
{
 public:
  SwigDirector_OverloadDerived2(void *swig_p);
  virtual ~SwigDirector_OverloadDerived2();
  void _swig_upcall_mmm() {
    OverloadBase::mmm();
  }
  virtual void mmm();
  void _swig_upcall_nnn__SWIG_0(int vvv) {
    OverloadBase::nnn(vvv);
  }
  virtual void nnn(int vvv);
  void _swig_upcall_nnn() {
    OverloadDerived2::nnn();
  }
  virtual void nnn();
 private:
  void *go_val;
};

#endif
