/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 3.0.3
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: ./../director_nspace_director_name_collision.i

#ifndef SWIG_director_nspace_director_name_collision_WRAP_H_
#define SWIG_director_nspace_director_name_collision_WRAP_H_

class SwigDirector_Foo : public TopLevel::A::Foo
{
 public:
  SwigDirector_Foo(void *swig_p);
  virtual ~SwigDirector_Foo();
  std::string _swig_upcall_ping() {
    return TopLevel::A::Foo::ping();
  }
  virtual std::string ping();
 private:
  void *go_val;
};

#endif
