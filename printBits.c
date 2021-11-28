#include <stdio.h>

void printBits(int num);

int main(void){
    int num = 8;
    printBits(num);
    return 0;
}

void printBits(int num){
    unsigned int checkBit = 1 << (sizeof(int) * 8 - 1);

    while(checkBit != 0){
        int result = num & checkBit;

        if(result == checkBit)
            printf("%d", 1);
        else 
            printf("%d", 0);
        
        checkBit = checkBit >> 1;
    }

    printf("\n");
}