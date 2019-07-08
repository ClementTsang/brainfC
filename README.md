# brainfC
A Brainf\*\*k &lt;=&gt; C compiler.  Useful?  No, not at all (unless you really want to convert C to BF or vice versa).  Educational?  At least for me, yes - it was an interesting way to use what I've learned in compiler design in class, albeit in a less... productive way.

Note the code I write is probably not the prettiest or best practice Go... I'm quite new to this language, hopefully this will improve as this project progresses.

## Instructions to Install
Clone the repo into your Go path and install.  *TODO: Provide more detailed instructions later*.

## Features (so far)
* Convert from BF to somewhat optimized C code
	* Optimizations include:
		* Peephole optimization
		* Duplicate instruction optimization
* Convert from (a subset of) C to BF
	* Current C syntax that is supported:

## Demo
To convert BF to C in command line, for example:
```bash
./brainfC -c ./examples/simple_add.bf
gcc ./examples/simple_add.c -o ./examples/simple_add.c
./examples/simple_add
```

Likewise, to convert C to BF in command line:
```bash
```
If you want to interpret the resulting BF code, use something like [Beef](https://kiyuko.org/software/beef).

*TODO: Insert demo using github pages*

## Credits/Thanks
I referred to the source code of [c2bf](https://github.com/arthaud/c2bf) a lot for the C =&gt; component.  Some BF code samples in the [examples](./examples) directory was sourced from the [Brainf\*\*k Wikipedia article](https://en.wikipedia.org/wiki/Brainfuck).  Many of the BF algorithms used for converting from C were sourced from the [Esolangs wiki for Brainf\*\*k](https://esolangs.org/wiki/Brainfuck_algorithms).