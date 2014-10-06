/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 3.0.3
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: ./../director_ignore.i

#ifndef SWIG_director_ignore_WRAP_H_
#define SWIG_director_ignore_WRAP_H_

class SwigDirector_DIgnores : public DIgnores
{
 public:
  SwigDirector_DIgnores(void *swig_p);
  virtual ~SwigDirector_DIgnores();
  void _swig_upcall_OverloadedMethod(bool b) {
    DIgnores::OverloadedMethod(b);
  }
  virtual void OverloadedMethod(bool b);
  int _swig_upcall_Triple(int n) {
    return DIgnores::Triple(n);
  }
  virtual int Triple(int n);
  virtual double PublicPureVirtualMethod1();
  virtual void PublicPureVirtualMethod2();
  virtual void TempMethod();
  void _swig_upcall_OverloadedProtectedMethod() {
    DIgnores::OverloadedProtectedMethod();
  }
  virtual void OverloadedProtectedMethod();
  virtual double ProtectedPureVirtualMethod1();
  virtual void ProtectedPureVirtualMethod2();
 private:
  void *go_val;
};

class SwigDirector_DAbstractIgnores : public DAbstractIgnores
{
 public:
  SwigDirector_DAbstractIgnores(void *swig_p);
  virtual ~SwigDirector_DAbstractIgnores();
  virtual double OverloadedMethod(int n, int xoffset, int yoffset);
  virtual double OverloadedMethod(int n, int xoffset);
  virtual double OverloadedMethod(int n);
  virtual double OverloadedMethod(bool b);
  int _swig_upcall_Quadruple(int n) {
    return DAbstractIgnores::Quadruple(n);
  }
  virtual int Quadruple(int n);
  virtual double OverloadedProtectedMethod(int n, int xoffset, int yoffset);
  virtual double OverloadedProtectedMethod(int n, int xoffset);
  virtual double OverloadedProtectedMethod(int n);
  virtual double OverloadedProtectedMethod();
 private:
  void *go_val;
};

class SwigDirector_DTemplateAbstractIgnoresInt : public DTemplateAbstractIgnores< int >
{
 public:
  SwigDirector_DTemplateAbstractIgnoresInt(void *swig_p);
  virtual ~SwigDirector_DTemplateAbstractIgnoresInt();
  virtual double OverloadedMethod(int n, int xoffset, int yoffset);
  virtual double OverloadedMethod(int n, int xoffset);
  virtual double OverloadedMethod(int n);
  virtual double OverloadedMethod(bool b);
  int _swig_upcall_Quadruple(int n) {
    return DTemplateAbstractIgnores< int >::Quadruple(n);
  }
  virtual int Quadruple(int n);
  virtual double OverloadedProtectedMethod(int n, int xoffset, int yoffset);
  virtual double OverloadedProtectedMethod(int n, int xoffset);
  virtual double OverloadedProtectedMethod(int n);
  virtual double OverloadedProtectedMethod();
 private:
  void *go_val;
};

class SwigDirector_DIgnoreConstructor : public DIgnoreConstructor
{
 public:
  SwigDirector_DIgnoreConstructor(void *swig_p, std::string s, int i);
  virtual ~SwigDirector_DIgnoreConstructor();
 private:
  void *go_val;
};

class SwigDirector_DIgnoreOnlyConstructor : public DIgnoreOnlyConstructor
{
 public:
  virtual ~SwigDirector_DIgnoreOnlyConstructor();
 private:
  void *go_val;
};

#endif
