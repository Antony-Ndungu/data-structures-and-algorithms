#include <stdio.h>

void insertionSort(int *nums, size_t n);

int main(void){
    int nums[] = {27, 4, 7, 12, 9};
    size_t numsSize = sizeof(nums)/sizeof(int);
    insertionSort(nums, numsSize);
    for(size_t i = 0; i < numsSize; i++)
        printf("%d\t", nums[i]);
    printf("\n");
    return 0;
}

// O(n^2) time | O(1) space
void insertionSort(int *nums, size_t n){
    size_t i, j;
    int tmp;
    for(i = 1; i < n; i++){
        j = i;
        tmp = nums[i];
        while(j > 0 && nums[j - 1] > tmp){
            nums[j] = nums[j -1];
            j--;
        }
        nums[j] = tmp;
    }
}