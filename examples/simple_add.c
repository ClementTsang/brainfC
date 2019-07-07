#include <stdio.h>
int main() {
char array[5000000] = {0};
char *ptr = array;
++*ptr;
++*ptr;
++ptr;
++*ptr;
++*ptr;
++*ptr;
++*ptr;
++*ptr;
while (*ptr) {
--ptr;
++*ptr;
++ptr;
--*ptr;
}
++*ptr;
++*ptr;
++*ptr;
++*ptr;
++*ptr;
++*ptr;
++*ptr;
++*ptr;
while (*ptr) {
--ptr;
++*ptr;
++*ptr;
++*ptr;
++*ptr;
++*ptr;
++*ptr;
++ptr;
--*ptr;
}
--ptr;
putchar(*ptr);
}
