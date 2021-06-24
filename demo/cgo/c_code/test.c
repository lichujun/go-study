#include "test.h"

/*int plus(int a, int b) {
    return a + b;
}*/

void init_number(struct number *n, int value) {
    n->value = value;
}

int plus(struct number *a, struct number *b) {
    return a->value + b->value;
}