/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 3.0.3
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: ./../apply_strings.i

#ifndef SWIG_apply_strings_WRAP_H_
#define SWIG_apply_strings_WRAP_H_

class SwigDirector_DirectorTest : public DirectorTest
{
 public:
  SwigDirector_DirectorTest(void *swig_p);
  UCharPtr _swig_upcall_UCharFunction(UCharPtr str) {
    return DirectorTest::UCharFunction(str);
  }
  virtual UCharPtr UCharFunction(UCharPtr str);
  SCharPtr _swig_upcall_SCharFunction(SCharPtr str) {
    return DirectorTest::SCharFunction(str);
  }
  virtual SCharPtr SCharFunction(SCharPtr str);
  CUCharPtr _swig_upcall_CUCharFunction(CUCharPtr str) {
    return DirectorTest::CUCharFunction(str);
  }
  virtual CUCharPtr CUCharFunction(CUCharPtr str);
  CSCharPtr _swig_upcall_CSCharFunction(CSCharPtr str) {
    return DirectorTest::CSCharFunction(str);
  }
  virtual CSCharPtr CSCharFunction(CSCharPtr str);
  CharPtr _swig_upcall_CharFunction(CharPtr buffer) {
    return DirectorTest::CharFunction(buffer);
  }
  virtual CharPtr CharFunction(CharPtr buffer);
  CCharPtr _swig_upcall_CCharFunction(CCharPtr buffer) {
    return DirectorTest::CCharFunction(buffer);
  }
  virtual CCharPtr CCharFunction(CCharPtr buffer);
  virtual ~SwigDirector_DirectorTest();
 private:
  void *go_val;
};

#endif
