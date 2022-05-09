package main

import "fmt"

//import "github.com/willf/bitset"
import "github.com/bits-and-blooms/bitset"

/*

Given a complete binary tree with nodes of values of either 1 or 0, the following rules always hold:

(1) a node's value is 1 if and only if all its subtree nodes' values are 1
(2) a leaf node can have value either 1 or 0

Implement the following 2 APIs:

set_bit(offset, length), set the bits at range from offset to offset+length-1
clear_bit(offset, length), clear the bits at range from offset to offset+length-1

i.e. The tree is like:
             0
          /     \
         0        1
       /  \      /  \
      1    0    1    1
     /\   / \   /
    1  1 1   0 1
Since it's complete binary tree, the nodes can be stored in an array:
[0,0,1,1,0,1,1,1,1,1,0,1]
*/

type bitarray struct {
	arr bitset.Bitset
	len int
}

func (*bitarray) clearBit(offset, length, start, end int) {

}

func (*bitarray) setBit(offset, length, start, end int) {

}

func (*bitarray) setBit(offset, length int) {
	// bounds checking
}

func (*bitarray) clearBit(offset, length int) {

}

func main() {
	var ba bitarray

	ba.setBit(0, 1)
	ba.setBit(8, 2)

	fmt.Println("len", b.Len())
	fmt.Println("count", b.Count())
	fmt.Println(b)
	fmt.Println(b.Test(8), b.Test(0), b.Test(10), b.Test(1024))

	b.Set(1023)

	fmt.Println(b, b.Len(), b.Len()/64, b.Count())
	fmt.Println(b.Test(8), b.Test(0), b.Test(10), b.Test(1024))
}

/*
			0				2^0
		  /    \
		0        0			2^1
	  /  \      /  \
	 1    1    0    0		2^2  == #leafs, 2^h = 4, 2^h = 2^2 , h = sqrt(4)
  /
0							2^3  = 8......h = 3,

8 page allocator = h, array size 2^(h+1)-1 , 15 bits 0-indexed


4 pages ==> ?? height  2^h = 4 == 2h ......2^(2+1)-1 == 7
4 pagesva
  - height = sqrt(4) = 2
  - #nodes = 2^(h+1)-1 = 2^(2+1)-1 = 7
2^0 + 2^1 + …. 2^h = 2^(h+1)-1 , h == 2

4GB ->

total memory = 4 pages
minimum allocation = 4


     3 nodes
     -------
  0  1  2  3  4  5  6
[ 0, 0, 1, 1, 0, 1, 1]
  r  lr       rr

           0  1  2  3  0-indexed
           1  2  3  4  1-indexed
  0  1  2  3  4  5  6
[ 0, 0, 0, 1, 1, 0, 0]
  r  -  -
     lr rr ----  -----
            ll     rl

4 byte mem allocator. #Leafs = 4,
lr = 2i+1, 0 ->
rr = 2i+2

power of 2 allocator, full binary tree
< power of 2 allocator, 3pages,

	set(startLeaf, size)

				set(0,4)		    4
			   /	    \
   		set(1,2)		set(2,2)
      /      \			/		\
set(3,1)   set(4,1)  set(5,1)	set(6,1)


*/
