#include <stdio.h>

int x = 1;
int y = 2;

int addtwonum();

#define LENGTH 10
#define WIDTH 5
#define NEWLINE '\n'

void func1(void);

//static int count = 10;
int count;

extern void write_extern();


int main(int argc, char *argv[]) {
    count = 5;
    write_extern();

//    while (count--) {
//        func1();
//    }

//    printf("%lu\n", sizeof(int));

//    int result;
//    result = addtwonum();
//    printf("result: %d", result);
//
//    int area;
//    area = LENGTH * WIDTH;
//    printf("value of area: %d", area);
//    printf("%c", NEWLINE);

    return 0;
}

void func1(void) {
    static int thingy = 5;
    thingy++;
    printf("%d %d\n", thingy, count);
}