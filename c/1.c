#include <stdio.h>
#include <math.h>

struct list {
    int c;
    int d;
    struct list *next;
};

struct list *map()

int main() {
    int n = 1000;
    for(int a = 1; a <= 1000; a++) {
        for(int b = 1; b <= 1000; b++) {
            for(int c = 1; c <= 1000; c++) {
                int d = pow( pow(a,3) + pow(b,3) - pow(c,3), 1/3);
                if(pow(a, 3) + pow(b,3) == pow(c, 3) + pow(d, 3)){
                    printf("%d %d %d %d\n", a, b, c, d);
                    break;
                }

            }
        }
    }
    return 0;
}
