/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 3.0.3
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: ./../director_default.i

#ifndef SWIG_director_default_WRAP_H_
#define SWIG_director_default_WRAP_H_

class SwigDirector_Foo : public Foo
{
 public:
  SwigDirector_Foo(void *swig_p, int i);
  SwigDirector_Foo(void *swig_p);
  virtual ~SwigDirector_Foo();
  std::string _swig_upcall_Msg__SWIG_0(std::string msg) {
    return Foo::Msg(msg);
  }
  virtual std::string Msg(std::string msg);
  std::string _swig_upcall_Msg__SWIG_1() {
    return Foo::Msg();
  }
  virtual std::string Msg();
 private:
  void *go_val;
};

class SwigDirector_DefaultsBase : public DefaultsBase
{
 public:
  SwigDirector_DefaultsBase(void *swig_p);
  virtual IntegerPtr defaultargs(double d, int *a);
  virtual IntegerPtr defaultargs(double d);
  virtual ~SwigDirector_DefaultsBase();
 private:
  void *go_val;
};

class SwigDirector_DefaultsDerived : public DefaultsDerived
{
 public:
  SwigDirector_DefaultsDerived(void *swig_p);
  int *_swig_upcall_defaultargs__SWIG_0(Double d, IntegerPtr a) {
    return DefaultsDerived::defaultargs(d,a);
  }
  virtual int *defaultargs(Double d, IntegerPtr a);
  int *_swig_upcall_defaultargs__SWIG_1(Double d) {
    return DefaultsDerived::defaultargs(d);
  }
  virtual int *defaultargs(Double d);
  virtual ~SwigDirector_DefaultsDerived();
 private:
  void *go_val;
};

#endif
