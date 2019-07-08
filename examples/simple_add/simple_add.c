#include <stdio.h>
int main() {
char t[30000] = {0};
char *p = t;
	*p+=2;
	p+=1;
	*p+=5;
	while (*p) {
		p-=1;
		*p+=1;
		p+=1;
		*p-=1;
	}
	*p+=8;
	while (*p) {
		p-=1;
		*p+=6;
		p+=1;
		*p-=1;
	}
	p-=1;
	putchar(*p);
}
