#include <stdio.h>
int main() {
char t[30000] = {0};
char *p = t;
	*p+=11;
	p+=1;
	*p+=1;
	p+=4;
	*p+=44;
	p+=1;
	*p+=32;
	p-=6;
	while (*p) {
		p+=1;
		while (*p) {
			p+=6;
			*p+=1;
			p+=1;
			*p+=1;
			p-=7;
			*p-=1;
		}
		p+=7;
		while (*p) {
			p-=7;
			*p+=1;
			p+=7;
			*p-=1;
		}
		p-=1;
		while (*p) {
			p+=1;
			*p+=10;
			while (*p) {
				*p-=1;
				p-=1;
				*p-=1;
				while (*p) {
					p+=2;
					*p+=1;
					p+=1;
					*p+=1;
					p-=3;
					*p-=1;
				}
				p+=3;
				while (*p) {
					p-=3;
					*p+=1;
					p+=3;
					*p-=1;
				}
				*p+=1;
				p-=1;
				while (*p) {
					p+=1;
					while (*p) {
						*p-=1;
					}
					p-=1;
					while (*p) {
						*p-=1;
					}
				}
				p+=1;
				while (*p) {
					p-=2;
					while (*p) {
						p+=3;
						*p+=1;
						p-=3;
						*p-=1;
					}
					p+=2;
					while (*p) {
						*p-=1;
					}
				}
				p-=2;
			}
			p+=3;
			while (*p) {
				p+=2;
				*p+=1;
				p+=1;
				*p+=1;
				p-=3;
				*p-=1;
			}
			p+=3;
			while (*p) {
				p-=3;
				*p+=1;
				p+=3;
				*p-=1;
			}
			*p+=1;
			p-=1;
			while (*p) {
				p+=1;
				while (*p) {
					*p-=1;
				}
				p-=1;
				while (*p) {
					*p-=1;
				}
			}
			p+=1;
			while (*p) {
				p-=2;
				*p+=1;
				p+=2;
				while (*p) {
					*p-=1;
				}
			}
			p-=7;
		}
		p+=5;
		while (*p) {
			*p+=48;
			putchar(*p);
			while (*p) {
				*p-=1;
			}
		}
		*p+=10;
		p-=1;
		while (*p) {
			*p-=1;
			p+=1;
			*p-=1;
			p-=1;
		}
		p+=1;
		*p+=48;
		putchar(*p);
		while (*p) {
			*p-=1;
		}
		p-=12;
		while (*p) {
			p+=3;
			*p+=1;
			p+=1;
			*p+=1;
			p-=4;
			*p-=1;
		}
		p+=4;
		while (*p) {
			p-=4;
			*p+=1;
			p+=4;
			*p-=1;
		}
		p-=1;
		*p-=1;
		while (*p) {
			p+=2;
			putchar(*p);
			p+=1;
			putchar(*p);
			p-=3;
			while (*p) {
				*p-=1;
			}
		}
		p-=2;
		while (*p) {
			p+=2;
			*p+=1;
			p+=1;
			*p+=1;
			p-=3;
			*p-=1;
		}
		p+=3;
		while (*p) {
			p-=3;
			*p+=1;
			p+=3;
			*p-=1;
		}
		p-=2;
		while (*p) {
			p-=1;
			*p+=1;
			p+=1;
			*p-=1;
		}
		p+=1;
		while (*p) {
			p-=1;
			*p+=1;
			p+=1;
			*p-=1;
		}
		p-=3;
		*p-=1;
	}
}
