#include <stdio.h>
int main() {
char t[30000] = {0};
char *p = t;
	*p-=1;
	*p = getchar();
	*p+=1;
	while (*p) {
		*p-=1;
		while (*p) {
			p+=2;
			*p+=4;
			while (*p) {
				p+=1;
				*p+=8;
				p-=1;
				*p-=1;
			}
			p-=1;
			*p+=1;
			p-=1;
			*p-=1;
			while (*p) {
				p+=1;
				*p+=1;
				p+=1;
				*p+=1;
				p+=1;
				*p-=1;
				while (*p) {
					p+=3;
				}
				p-=1;
				while (*p) {
					while (*p) {
						p+=1;
						*p+=1;
						p-=1;
						*p-=1;
					}
					p+=2;
					*p+=1;
					p+=1;
				}
				p-=5;
				*p-=1;
			}
		}
		p+=3;
		while (*p) {
			*p-=1;
		}
		*p+=1;
		p+=1;
		*p-=2;
		while (*p) {
			*p-=1;
			while (*p) {
				p-=1;
				*p-=1;
				p+=1;
				*p+=3;
				while (*p) {
					*p-=1;
				}
			}
		}
		p-=1;
		while (*p) {
			*p+=12;
			p-=1;
			while (*p) {
				p+=1;
				*p-=1;
				while (*p) {
					p+=1;
					*p+=1;
					p+=2;
				}
				p+=1;
				while (*p) {
					*p+=1;
					while (*p) {
						p-=1;
						*p+=1;
						p+=1;
						*p-=1;
					}
					p+=1;
					*p+=1;
					p+=2;
				}
				p-=5;
				*p-=1;
			}
			p+=2;
			while (*p) {
				p-=1;
				*p+=1;
				p+=1;
				*p-=1;
			}
			p+=1;
			while (*p) {
				*p-=1;
				while (*p) {
					*p-=1;
					p-=2;
					while (*p) {
						*p-=1;
					}
					p+=2;
				}
				p-=2;
				while (*p) {
					p-=2;
					*p-=1;
					p+=2;
					*p-=1;
				}
				p+=2;
			}
			p-=2;
			while (*p) {
				p-=2;
				*p+=1;
				p+=2;
				*p-=1;
			}
		}
		p-=1;
		while (*p) {
			*p-=1;
		}
		p-=1;
		putchar(*p);
		while (*p) {
			*p-=1;
		}
		p-=1;
		*p-=1;
		*p = getchar();
		*p+=1;
	}
}
