#include <stdio.h>

int max(int *num, size_t n);

int main(void){
    int nums[] = {32, 2, 3, 102};
    printf("%d\n", max(nums, sizeof(nums)/sizeof(int)));
    return 0;
}

/*
Algorithm max(nums, n):
Input: An array nums storing n integers
Output: The maximum element in nums
currentMax <- nums[0]
for i <- 1 to n-1 do
    if currentMax < nums[i] then currentMax <- nums[i]
return currentMax
*/
int max(int *nums, size_t n){
    int currentMax = nums[0];
    for(size_t i = 1; i < n; i++){
        if(currentMax < nums[i])
            currentMax = nums[i];
    }
    return currentMax;
}