#include <stdio.h>

void swap(int *a, int *b);
void selectionSort(int *nums, size_t n);

int main(void){
    int nums[] = {27, 4, 7, 12, 9};
    size_t numsSize = sizeof(nums)/sizeof(int);
    selectionSort(nums, numsSize);
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

//O(n^2) time | O(1) Space
void selectionSort(int *nums, size_t n){
    int indexOfMinNum;
    for(size_t i = 0; i < n - 1; i++){
        indexOfMinNum = i;
        for(size_t j = i + 1; j < n; j++){
            if(nums[j] < nums[indexOfMinNum])
                indexOfMinNum = j;
        }
        swap(&nums[i], &nums[indexOfMinNum]);
    }
}