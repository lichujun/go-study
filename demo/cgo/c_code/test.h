#ifndef TEST_H
#define TEST_H

struct number {
    int value;
};

//int plus(int a, int b);

int plus(struct number *a, struct number *b);
void init_number(struct number *n, int value);

#endif