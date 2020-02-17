#include <string.h>
#include <stdlib.h>
#include <stdio.h>

#define LEN(x) (sizeof(x) / sizeof((x)[0]))
/*
#define name7(NAME, INT) (NAME)[s##INT]->v=s##INT \
					(NAME)[s##INT]->v=s##INT \
					(NAME)[s##INT]->v=s##INT \
					(NAME)[s##INT]->v=s##INT \
					(NAME)[s##INT]->v=s##INT \
					(NAME)[s##INT]->v=s##INT \
					(NAME)[s##INT]->v=s##INT

#define loop7(NAME, INT, NAME2) (NAME)[s##INT]=(NAME2)() \
								(NAME)[s##INT]=(NAME2)() \
								(NAME)[s##INT]=(NAME2)() \
								(NAME)[s##INT]=(NAME2)() \
								(NAME)[s##INT]=(NAME2)() \
								(NAME)[s##INT]=(NAME2)() \
								(NAME)[s##INT]=(NAME2)()
*/
int s1[4] = {2, 7, 11, 15}; 
int s2[4] = {2, 7, 11, 2};
int s3[7] = {11, 1, 7, 10, 9, 12, 4};
int s4[7] = {5, 6, 7, 1, 8, 0, 7};
int s5[12] = {11, 5, 50, 11, 12, 40, 6, 3, 88, 8, 100, 4};
int s6[8] = {1, 1, 1, 6, 10, 2, 4, 1};
int s7[7] = {1, 5, 10, 3, 2, 23, 1};

struct node_of_samples { 
	int *v;
	struct samples *next;
};

struct list {
	struct list *head;
};

typedef struct node_of_samples node;

int goals[] = {4, 4, 23, 50, 77, 12, 2}

struct list *create_list() {
	struct list *list = (struct list*) malloc(sizeof(struct list));
	list->head = NULL;
	return list;
}

struct node *create_node(char *s) {
	struct node *node = (struct node*) malloc(sizeof(struct node *));
	node->v = s;
	node->next = NULL;
	return node;
}

struct list *add_to_list() {

	struct list *list = create_list();
	struct node *node1 = create_node();

	list->head = node1;

	node1->v = s1;
	node1->next = create_node();

	struct node *node2 = node1->next;
	node2->v = s2;
	node2->next = create_node();

	struct node *node3 = node2->next;
	node3->v = s3;
	node3->next = create_node();

	struct node *node4 = node3->next;
	node4->v = s4;
	node4->next = create_node();

	struct node *node5 = node4->next;
	node5->v = s5;
	node5->next = create_node();

	struct node *node6 = node5->next;
	node6->v = s6;
	node6->next = create_node();

	struct node *node7 = node6->next;
	node7->v = s7;
	node7->next = create_node();

	return list;
}

int* twoSum(int* nums, int numsSize, int target, int* returnSize){
    int i, j;
    int current;
    int *found = (int *) malloc(2*sizeof(int));
    int pool[numsSize];
    int memsize;
	for(i = numsSize - 1, j = 0; i >= 0; i--, j++) {
    	memsize = (numsSize - j)*(sizeof(int));
        memcpy(pool, nums, memsize);
        for(int k = 0; k < (numsSize - j); k++) {
        	current = nums[i] + pool[k];

	        if(current == target) {
	            found[1] = i;
	            found[0] = k;
                returnSize[0] = nums[k];
                returnSize[1] = nums[i];
	            return found;
	        } else {
                current = 0;
            }
	    }
        current = 0;
        memset(pool, 0x00, numsSize*sizeof(int));
    }
    return NULL;
}

int main() {
	srand((unsigned int)time(NULL));

	int numsSize, target, include;
	int *nums, *ranNums, *found;

	int returnSize[2];

	struct list *test_arrs_list = add_to_list();
	struct node *iter;

	for(iter = list->head; iter != NULL; iter = iter->next) {
		
	}

	nums = (int *) malloc(numsSize * sizeof(int));

	for(int i = 0; i < numsSize; i++) {
		printf("%d ", nums[i]);
		if(i % 10 == 0)
			printf("\n");
	}
	printf("\n");
	
	found = twoSum(nums, numsSize, target, returnSize);

	if(found) {
		printf("%d %d\n", found[0], found[1]);
		printf("%d %d\n", returnSize[0], returnSize[1]);
	} else
		printf("NULL\n");

	return 0;
}
