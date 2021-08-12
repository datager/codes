//
// Created by 孙羽川 on 2020/10/7.
//

#include<stdio.h>

extern int count;

void write_extern(void) {
    printf("count is %d\n", count);
}