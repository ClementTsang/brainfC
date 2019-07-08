#include <stdio.h>
int main() {
char t[30000] = {0};
char *p = t;
while (*p) {
}
(*p)+=1;
p+=2;
(*p)-=3;
p-=4;
(*p)+=3;
}
