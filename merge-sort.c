#include <stdio.h>

void merge(int *nums, size_t start, size_t end);
void mergeSort(int *nums, size_t start, size_t end);

int main(void){
    int nums[] = {27, 4, 7, 12, 9};
    size_t numsSize = sizeof(nums)/sizeof(int);
    mergeSort(nums, 0, numsSize -1);
    for(size_t i = 0; i < numsSize; i++)
        printf("%d\t", nums[i]);
    printf("\n");
    return 0;
}

void merge(int *nums, size_t start, size_t end){
    size_t mid = (start + end) / 2;
    int tempArr[end-start+1];
    int i = start, j = mid + 1, k = 0;

    while(i <= mid && j <= end){
        if(nums[i] < nums[j])
            tempArr[k++] = nums[i++];
        else 
            tempArr[k++] = nums[j++];
    }

    while(i <= mid)
        tempArr[k++] = nums[i++];

    while(j <= end)
        tempArr[k++] = nums[j++];

    for(size_t i = 0; i < end + 1; i++)
        nums[start+i] = tempArr[i];
}

void mergeSort(int *nums, size_t start, size_t end){
    if(start == end)
        return;
    size_t mid = (start + end) / 2;
    mergeSort(nums, start, mid);
    mergeSort(nums, mid + 1, end);
    merge(nums, start, end);
}