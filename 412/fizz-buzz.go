package leetcode

import (
	"strconv"
)

/*
https://leetcode.com/problems/fizz-buzz/

412. Fizz Buzz
Easy

493

733

Favorite

Share
Write a program that outputs the string representation of numbers from 1 to n.

But for multiples of three it should output “Fizz” instead of the number and for the multiples of five output “Buzz”. For numbers which are multiples of both three and five output “FizzBuzz”.

Example:

n = 15,

Return:
[
    "1",
    "2",
    "Fizz",
    "4",
    "Buzz",
    "Fizz",
    "7",
    "8",
    "Fizz",
    "Buzz",
    "11",
    "Fizz",
    "13",
    "14",
    "FizzBuzz"
]
 */

func fizzBuzz(n int) []string {
	rt := make([]string, n)
	var i = 1
	for i <= n {
		var t string
		if i%15 == 0 {
			t = "FizzBuzz"
		} else if i%5 == 0 {
			t = "Buzz"
		} else if i%3 == 0 {
			t = "Fizz"
		}else{
			t = strconv.Itoa(i)
		}
		rt[i-1] = t
		i++
	}
	return rt
}