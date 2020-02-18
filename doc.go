/*
	Package slowseeder implements a drop-in replacement for a rand source
	intended for cryptographic key generation.

	It has been designed to be simple to comprehend. Generation is
	deterministic from a seed, uses multiple layered hashing functions,
	and is parameterized to easily extend the time spent during each
	iteration, making brute force and pre-computation more difficult.
*/
package slowseeder
