package main

import (
	"fmt"
)

/*
	Given three strings - A, B, C - determine if C is a valid interleaved
	combination of A and B. The characters within A and B should appear in their
	original order, but they can be interleaved arbitrarily (e.g. all of A's
	characters can come first, or we can alternate, all of B's can come first,
	etc.).

	A = 'abc',
	B = 'def',

	abcdef => true
	defabc => true
	daebfc => true
	cabfed => false
	daecfb => false (b should come before c)

	A = ab
	B = ac

	acba => false
*/

func Interleave(a string, b string, c string) bool {
	pA := 0
	// pB := 0

	for i, s := range c {
		fmt.Println(i, c[i], s)
		if a[pA] == c[i] {
			pA++
		}

		// if b[pB] == c[i] {
		// 	pB++
		// }
	}

	// if len(a) == pA && len(b) == pB {
	// 	return true
	// }

	return false
}

// process.stdin.resume();
// process.stdin.setEncoding("ascii");

// function badPseudoCheckInterleavedStrings(a, b, c) {
//     /*
//         a = 'abc'
//         a = 'aab'
//         b = 'baa'

//         let a = [
//             // a,
//             // a,
//             // b,
//         ]

//         let b = [
//         b,
//         a,
//         a,
//         ]

//         let t = [
//             a
//         ]

//         let pointer = {
//             a: 0
//             b: 0
//             c: 0
//         }

//         // a = 'ab', b = 'ac', c = 'acab'
//         // c = 'aabbaa'
//         // a = ab, b = cd, c = abce
//         //
//         for _, s := range c {
//             if (a[pointer.a] === s) {
//                 pointer.a += 1
//             }

//             if (b[pointer.b] === s) {
//                 pointer.b += 1
//             }

//             pointer.c += 1
//         }

//         if (pointer.a === a.length && pointer.b === b.length) {
//             return true
//         } else {
//             return false
//         }

//     */
// }

// // console.log(checkInterleavedStrings("ab", "ac", "acba"));
// function checkInterleavedStrings(a, b, c) {
//     const pointer = {
//         a: 0,
//         b: 0,
//     };

//     for (i = 0; i < c.length; i += 1) {
//         if (a.charAt(pointer.a) === c.charAt(i)) {
//             pointer.a += 1;
//         }

//         if (b.charAt(pointer.b) === c.charAt(i)) {
//             pointer.b += 1;
//         }
//     }

//     if (pointer.a === a.length && pointer.b === b.length) {
//         return true
//     } else {
//         return false
//     }
// }

// console.log("Valid cases:");
// console.log(checkInterleavedStrings("abc", "def", "adbecf"));
// console.log(checkInterleavedStrings("abc", "def", "abcdef"));
// console.log(checkInterleavedStrings("abc", "def", "abdefc"));

// // console.log(checkInterleavedStrings("aab", "baa", "abdefc"));
// console.log(checkInterleavedStrings("ab", "ac", "acab"));

// console.log("Invalid cases:");
// console.log(checkInterleavedStrings("abc", "def", "adcebf"));
// console.log(checkInterleavedStrings("ab", "ac", "acba"));

// //

// var input = "";
// process.stdin.on("data", function (chunk) {
//     input += chunk;
// });
// process.stdin.on("end", function () {
//     // now we can read/parse input
// });

// /*
// Given three strings - A, B, C - determine if C is a valid interleaved combination of A and B. The characters within A and B should appear in their original order, but they can be interleaved arbitrarily (e.g. all of A's characters can come first, or we can alternate, all of B's can come first, etc.).
// A = 'abc', B = 'def', C = 'abcdef' (chars in A comes first), 'defabc' (B comes first), 'daebfc', 'deabfc' => true

// abcdef => true
// defabc => true
// daebfc => true
// cabfed => false
// daecfb => false (b should come before c)
// */
