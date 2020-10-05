package main

/*
Mark and Jane are very happy after having their first child. Their son loves toys, so Mark wants to buy some. There are a number of different toys lying in front of him, tagged with their prices. Mark has only a certain amount to spend, and he wants to maximize the number of toys he buys with this money.

Given a list of prices and an amount to spend, what is the maximum number of toys Mark can buy? For example, if  and Mark has  to spend, he can buy items  for , or  for  units of currency. He would choose the first group of  items.

Function Description

Complete the function maximumToys in the editor below. It should return an integer representing the maximum number of toys Mark can purchase.

maximumToys has the following parameter(s):

prices: an array of integers representing toy prices
k: an integer, Mark's budget
Input Format

The first line contains two integers,  and , the number of priced toys and the amount Mark has to spend.
The next line contains  space-separated integers 

Constraints

A toy can't be bought multiple times.

Output Format

An integer that denotes the maximum number of toys Mark can buy for his son.

Sample Input

7 50
1 12 5 111 200 1000 10
Sample Output

4
Explanation

He can buy only  toys at most. These toys have the following prices: .
*/

// Complete the maximumToys function below.
func maximumToys(prices []int32, k int32) int32 {
	n := len(prices)
	numberOfItemsToBuy := int32(0)
	sumOfToysSoFar := int32(0)

	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if prices[j] < prices[j+1] {
				tmp := prices[j]
				prices[j] = prices[j+1]
				prices[j+1] = tmp
			}
		}
		sumOfToysSoFar += prices[n-i-1]
		if sumOfToysSoFar <= k {
			numberOfItemsToBuy++
			continue
		}
		break
	}
	return numberOfItemsToBuy
}
