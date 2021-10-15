#include <stdio.h>

void swap(int *a, int *b);
void bubbleSort(int *nums, size_t n);

int main(void){
    int nums[] = {27, 4, 7, 12, 9};
    size_t numsSize = sizeof(nums)/sizeof(int);
    bubbleSort(nums, numsSize);
    for(size_t i = 0; i < numsSize; i++)
        printf("%d\t", nums[i]);
    printf("\n");
    return 0;
}

void swap(int *a, int *b){
    int temp = *a;
    *a = *b;
    *b = temp;
}

// O(n^2) time | O(1) Space
void bubbleSort(int *nums, size_t n){
    int swaps = 0;
    do {
        swaps = 0;
        for(size_t i = 0; i < n - 1; i++){
            if(nums[i] > nums[i+1]){
                swap(&nums[i], &nums[i+1]);
                swaps++;
            }
        }
    }while(swaps > 0);
}
