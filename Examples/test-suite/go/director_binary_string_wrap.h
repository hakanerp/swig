/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 3.0.3
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: ./../director_binary_string.i

#ifndef SWIG_director_binary_string_WRAP_H_
#define SWIG_director_binary_string_WRAP_H_

class SwigDirector_Callback : public Callback
{
 public:
  SwigDirector_Callback(void *swig_p);
  virtual ~SwigDirector_Callback();
  void _swig_upcall_run(char *dataBufferAA, int sizeAA, char *dataBufferBB, int sizeBB) {
    Callback::run(dataBufferAA,sizeAA,dataBufferBB,sizeBB);
  }
  virtual void run(char *dataBufferAA, int sizeAA, char *dataBufferBB, int sizeBB);
 private:
  void *go_val;
};

#endif
