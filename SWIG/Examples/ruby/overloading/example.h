#ifndef EXAMPLE_H
#define EXAMPLE_H

class Bar {
public:
    Bar();
    Bar(const Bar&);
    Bar(double);
    Bar(double, char *);
    Bar(int, int);
    Bar(char *);
    Bar(long);
    Bar(int);
    Bar(Bar *);

    void foo(const Bar&);
    void foo(double);
    void foo(double, char *);
    void foo(int, int);
    void foo(char *);
    void foo(long);
    void foo(int);
    void foo(Bar *);
};

void foo(const Bar&);
void foo(double);
void foo(double, char *);
void foo(int, int);
void foo(char *);
void foo(int);
void foo(long);
void foo(Bar *);

#endif
