
// This is the cpp_typedef runtime testcase. It checks that shadow classes are
// generated for typedef'd types.

import cpp_typedef.*;

public class cpp_typedef_runme {

  static {
    try {
	System.loadLibrary("cpp_typedef");
    } catch (UnsatisfiedLinkError e) {
      System.err.println("Native code library failed to load. See the chapter on Dynamic Linking Problems in the SWIG Java documentation for help.\n" + e);
      System.exit(1);
    }
  }

  public static void main(String argv[]) {

    Test test = new Test();
    UnnamedStruct unnamed = new UnnamedStruct();
    NamedStruct named = new NamedStruct();

    UnnamedStruct unnamed2 = test.test1(unnamed);
    NamedStruct named2 = test.test2(named);
    NamedStruct named3 = test.test3(named);
    NamedStruct named4 = test.test4(named);
  }
}

