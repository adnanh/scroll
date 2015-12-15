# scroll
continuously scroll a given string...

```
Usage:

  scroll [OPTIONS] input string

Options:
  -delay=600: delay in miliseconds between line outputs
  -size=80: number of characters to display
  -start=0: character position to start from
  -step=1: number of characters to step by
```

Example:
```
➜  ~  scroll -size 15 hello there, this message should scroll...
hello there, th
ello there, thi
llo there, this
lo there, this 
o there, this m
 there, this me
there, this mes
here, this mess
ere, this messa
re, this messag
e, this message
, this message 
 this message s
this message sh
his message sho
is message shou
s message shoul
 message should
message should 

.

.

.
```
