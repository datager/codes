#include <iostream>
#include "find.h"

int main() {
//    int i = 0;
//    int arr[3] = {0};
//    for (; i <= 3; i++) {
//        arr[i] = 0;
//        std::cout << arr[i] << "hello world\n" << std::endl;
//    }

    char arr1[6] = {4, 2, 3, 5, 9, 6};
    char arr2[6] = {4, 2, 3, 5, 9, 6};

    std::cout << find(arr1, sizeof(arr1) / sizeof(arr1[0]), 7) << std::endl;
    std::cout << find(arr2, sizeof(arr2) / sizeof(arr2[0]), 6) << std::endl;

    std::cout << sizeof(arr1)/sizeof(arr1[0]) << std::endl;
    return 0;

}