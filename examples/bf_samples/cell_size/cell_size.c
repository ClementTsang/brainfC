#include <stdio.h>
int main() {
char t[30000] = {0};
char *p = t;
	*p+=8;
	while (*p) {
		p+=1;
		*p+=8;
		p-=1;
		*p-=1;
	}
	p+=1;
	while (*p) {
		p-=1;
		*p+=4;
		p+=1;
		*p-=1;
	}
	*p+=1;
	p-=1;
	while (*p) {
		p+=1;
		*p-=1;
		p-=1;
		while (*p) {
			p+=1;
			*p+=4;
			p-=1;
			*p-=1;
		}
		p+=1;
		while (*p) {
			p-=1;
			*p+=8;
			p+=1;
			*p-=1;
		}
		p-=1;
		while (*p) {
			p+=1;
			*p+=8;
			p-=1;
			*p-=1;
		}
		*p+=1;
		p+=1;
		while (*p) {
			p+=1;
			*p+=10;
			while (*p) {
				p+=1;
				*p+=5;
				p-=1;
				*p-=1;
			}
			p+=1;
			*p+=1;
			putchar(*p);
			*p-=1;
			putchar(*p);
			while (*p) {
				*p-=1;
			}
			p-=2;
			while (*p) {
				*p-=1;
			}
			p-=1;
			*p-=1;
			p+=1;
		}
		p-=1;
		while (*p) {
			p+=2;
			*p+=7;
			while (*p) {
				p+=1;
				*p+=7;
				p-=1;
				*p-=1;
			}
			p+=1;
			putchar(*p);
			*p+=5;
			putchar(*p);
			while (*p) {
				*p-=1;
			}
			p-=3;
			*p-=1;
		}
	}
	p+=1;
	while (*p) {
		p+=1;
		*p+=8;
		while (*p) {
			p+=1;
			*p+=7;
			p-=1;
			*p-=1;
		}
		p+=1;
		putchar(*p);
		while (*p) {
			*p-=1;
		}
		p-=2;
		*p-=1;
	}
	p-=1;
	*p+=11;
	while (*p) {
		p+=1;
		*p+=3;
		p+=1;
		*p+=9;
		p+=1;
		*p+=9;
		p+=1;
		*p+=1;
		p-=4;
		*p-=1;
	}
	p+=1;
	*p-=1;
	putchar(*p);
	p+=1;
	*p-=1;
	putchar(*p);
	*p+=7;
	putchar(*p);
	*p+=11;
	putchar(*p);
	p-=1;
	putchar(*p);
	p+=2;
	putchar(*p);
	*p+=2;
	putchar(*p);
	*p+=7;
	putchar(*p);
	putchar(*p);
	p-=1;
	*p-=1;
	putchar(*p);
	p+=2;
	*p-=1;
	putchar(*p);
	while (*p) {
		while (*p) {
			*p-=1;
		}
		p-=1;
	}
}
