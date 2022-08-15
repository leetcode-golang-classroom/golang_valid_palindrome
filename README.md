# golang_valid_palindrome

A phrase is a **palindrome** if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

Given a string `s`, return `true` *if it is a **palindrome**, or* `false` *otherwise*.

## Examples

**Example 1:**

```
Input: s = "A man, a plan, a canal: Panama"
Output: true
Explanation: "amanaplanacanalpanama" is a palindrome.

```

**Example 2:**

```
Input: s = "race a car"
Output: false
Explanation: "raceacar" is not a palindrome.

```

**Example 3:**

```
Input: s = " "
Output: true
Explanation: s is an empty string "" after removing non-alphanumeric characters.
Since an empty string reads the same forward and backward, it is a palindrome.
```

**Constraints:**

- `1 <= s.length <= 2 * 105`
- `s` consists only of printable ASCII characters.

## 解析

給定一個字串 s ,

假設經過一個運算只看是英文字母的字元，且把所有大寫英文字元轉成小寫後變成 s’

要求寫一個演算法判斷 s 在經過以上轉換後是否是回文字串

一個字串 s 如果是回文， 代表 i = 0..(n/2-1)

s[i] = s[n-1-i], where n = len(s)

由以上定義可以發現透過給定 2 個指標 一個從字串最前面 一個從字串最後面

當兩個指標位置不同時做以下比較

每次當遇到兩個字元相同時，兩個指標共同推進一個單位往下一個可以比較的字元去比較

遇到兩個字元不同時，則代表不是回文 回傳 false

特別要注意的是當其中一個指標遇到不是非英文字元，則需要直接把指標往下一個位置移動 

然後往一下個開始指標開始比較

假設直到最後一個指標都相同則代表是回文 所以回傳 true

演算法詳細如下圖

![](https://i.imgur.com/BxTUXpO.png)

## 程式碼
```go
package sol

func isPalindrome(s string) bool {
	lp, rp := 0, len(s)-1
	var isAlphaBet = func(c byte) bool {
		if ('a' <= c && 'z' >= c) || ('A' <= c && 'Z' >= c) {
			return true
		}
		return false
	}
	var toLower = func(c byte) byte {
		if 'A' <= c && 'Z' >= c {
			return (c - 'A') + 'a'
		}
		return c
	}
	for lp <= rp {
		if !isAlphaBet(s[lp]) {
			lp++
			continue
		}
		if !isAlphaBet(s[rp]) {
			rp--
			continue
		}
		if toLower(s[rp]) != toLower(s[lp]) {
			return false
		}
		rp--
		lp++
	}
	return true
}
```

## 困難點

1. 從 palindrome 定義思考出透過左右指標同時逼近來檢查

## Solve Point

- [x]  初始化 lp = 0, rp = len(s) -1
- [x]  for lp ≤ rp 做以下運算
- [x]  當 s[lp] 不在 ‘a’,’z’ 之間 且不在 ‘A’ ,‘Z’之間, 更新 lp += 1 continue
- [x]  當 s[rp] 不在 ‘a’,’z’ 之間 且不在 ‘A’ ,‘Z’之間, 更新 rp -= 1 continue
- [x]  當 toLower(s[rp]) ≠ toLower(s[lp]) , 回傳 false
- [x]  rp—; lp++
- [x]  回傳 true