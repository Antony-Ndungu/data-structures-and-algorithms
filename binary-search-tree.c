#include <stdio.h>
#include <stdlib.h>

struct node {
    int value;
    struct node *leftChildPtr;
    struct node *rightChildPtr;
};

typedef struct node node_t;

struct binarySearchTree {
    node_t *rootPtr;
};

typedef struct binarySearchTree binarySearchTree_t;

binarySearchTree_t * create();

int main(void){
    binarySearchTree_t *btsPtr = create();
    free(btsPtr);
    return 0;
}

binarySearchTree_t * create(){
    binarySearchTree_t *ptr = (binarySearchTree_t *) malloc(sizeof(binarySearchTree_t));
    if(ptr == NULL)
        return NULL;
    ptr->rootPtr = NULL;
    return ptr;
}