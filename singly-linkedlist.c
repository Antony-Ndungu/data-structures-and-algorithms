#include <stdio.h>
#include <stdlib.h>

struct node{
    int value;
    struct node *next;
};

typedef struct node node_t;

struct list{
    node_t *head;
};

typedef struct list list_t;

list_t *create();
void append(list_t *listPtr, int value);
void print(list_t *listPtr);
void delete(list_t *listPtr, int value);
void destroy(list_t *listPtr);

int main(void){
    list_t *myListPtr = create();
    if(myListPtr == NULL){
        printf("Failed to create list\n");
        return 1;
    }

    append(myListPtr, 1);
    append(myListPtr, 2);
    append(myListPtr, 3);
    append(myListPtr, 4);
    append(myListPtr, 5);
    print(myListPtr);

    delete(myListPtr, 1);
    print(myListPtr);

    delete(myListPtr, 3);
    print(myListPtr);

    delete(myListPtr, 5);
    print(myListPtr);

    destroy(myListPtr);
    return 0;
}

list_t *create(){
    list_t *newListPtr = (list_t *) malloc(sizeof(list_t));
    if(newListPtr == NULL){
        return NULL;
    }

    newListPtr->head = NULL;
    return newListPtr;
}

void append(list_t *listPtr, int value){
    if(listPtr == NULL){
        printf("Failed to append the value because list pointer is NULL.\n");
        return;
    }

    node_t *newNodePtr = (node_t *) malloc(sizeof(node_t));
    newNodePtr->value = value;
    newNodePtr->next = NULL;

    if(listPtr->head == NULL){
        listPtr->head = newNodePtr;
        return;
    }

    node_t *currentNodePtr = listPtr->head;
    while(1){
        if(currentNodePtr->next == NULL){
            currentNodePtr->next = newNodePtr;
            return;
        }
        currentNodePtr = currentNodePtr->next;
    }
}

void print(list_t *listPtr){
    if(listPtr == NULL)
        return;
    node_t *currentNodePtr = listPtr->head;
    while(currentNodePtr != NULL){
        printf("%d => ", currentNodePtr->value);
        currentNodePtr = currentNodePtr->next;
    }
    printf("\n");
}

// Removes the first instance of the passed integer value in the list.
// It does nothing if the list does not contain integer value passed.
void delete(list_t *listPtr, int value){
    if(listPtr == NULL)
        return;
    node_t *currentNode = listPtr->head, *previousNode = NULL;
    while(currentNode != NULL){
        if(currentNode->value == value){
            if(currentNode == listPtr->head){
                listPtr->head = currentNode->next;
            }else{
                previousNode->next = currentNode->next;
            }
            free(currentNode);
            return;
        }
        previousNode = currentNode;
        currentNode = currentNode->next;
    }
}

void destroy(list_t *listPtr){
    if(listPtr == NULL)
        return;
    node_t *currentNodePtr = listPtr->head, *holdPtr = NULL;
    while(currentNodePtr != NULL){
        holdPtr = currentNodePtr;
        currentNodePtr = currentNodePtr->next;
        free(holdPtr);
    }
    free(listPtr);
}