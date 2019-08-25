# brainfC

A Brainf\*\*k &lt;=&gt; C transpiler.

Note the code I write is probably not the prettiest or best practice Go... I'm quite new to this language, hopefully this will improve as this project progresses.

## Instructions to Install

Clone the repo into your Go path.  Go to the cloned repo directory and run the makefile.

## Features (so far)

* Convert from BF to somewhat optimized (but not really) C code
	* Optimizations include:
		* Peephole optimization
		* Duplicate instruction optimization
* Convert from (a subset of) C to BF
	* Current C syntax that is supported:

## Usage

To convert BF to C in command line, for example:
```bash
./brainfC -c ./examples/bf_samples/hello_world/hello_world.bf
```

Likewise, to convert C to BF in command line:
```bash
./brainfC -b ./examples/c_samples/hello_world/hello_world.c
```
If you want to interpret the resulting BF code, use something like [Beef](https://kiyuko.org/software/beef).


# Demo

*TODO: Insert demo using github pages*

## Contribution
If you spot a bug/typo/issue and/or want to contribute, let me know.

## Why...?

Why not?

It was mostly an interesting way to apply what I've learned in compiler design from class (a bunch of what I learned carried over), as well as try to use Go, albeit in a way that was neither really productive or practical.

## Credits/Thanks
Bunch of resources/projects I have to give thanks to.

* I referred to the source code of [c2bf](https://github.com/arthaud/c2bf) a lot for the C =&gt; component.
* Some BF code samples in the [examples](./examples) directory was sourced from the [Brainf\*\*k Wikipedia article](https://en.wikipedia.org/wiki/Brainfuck).
* Many of the BF algorithms used for converting from C to BF were sourced from the [Esolangs wiki for Brainf\*\*k](https://esolangs.org/wiki/Brainfuck_algorithms).
* yacc code for C to BF was based on [the yacc file from this website](http://www.quut.com/c/ANSI-C-grammar-y.html), and the lexer was based on [this page from the same website](http://www.quut.com/c/ANSI-C-grammar-l-2011.html). 
