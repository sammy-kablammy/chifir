# Chifir VM

A virtual machine based on the specification in
["The Cuneiform Tablets of 2015" by Long Tien Nguyen and Alan Kay](https://archive.org/details/tr2015004_cuneiform).

They don't seem to actually provide the Smalltalk-72 disk image, so I don't have
a good way of testing this.

They also don't specify how arithmetic and logical operations should work. I
assumed 2's complement for arithmetic and C-style logical operations, i.e. zero
is falsy and nonzero is truthy. Maybe it's fair to assume future
cyberarchaeologists would know this, maybe not.

Instruction 14 ("refresh the screen") and 15 ("get one character from the
keyboard and store it into M[A]") aren't supported.
