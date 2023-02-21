package main

import (
	"bytes"
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"testing"
)

func PreStrBuilder(n int, str string) string {
	var builder strings.Builder
	builder.Grow(n)
	for i := 0; i < n; i++ {
		builder.WriteString(str)
	}
	return builder.String()
}

func PreByteBuffer(n int, str string) string {
	buf := new(bytes.Buffer)
	buf.Grow(n * len(str))

	for i := 0; i < n; i++ {
		buf.WriteString(str)
	}
	return buf.String()
}

func TestName1(t *testing.T) {
	r := "10109"
	c := int(r[len(r)-1] - '0')
	fmt.Println(c)
}

func TestName(t *testing.T) {
	a := rand.Perm(100)

	//a = a[:copy(a, a[2:])]
	fmt.Println(a)
}

func BenchmarkTestCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := rand.Perm(100)
		fmt.Println(a)
	}

}

func TestBav(t *testing.T) {
	v := "0011001000111100001010110011100001001010010010100100100100101010010000000100101101000100001111100101001101100001010001010110001001011010011000110010010101101000"
	s, _ := strconv.Atoi(v)
	fmt.Println(s)

}
func addBinary(a string, b string) string {
	ans := ""
	carry := 0
	lenA, lenB := len(a), len(b)
	n := max(lenA, lenB)

	for i := 0; i < n; i++ {
		if i < lenA {
			carry += int(a[lenA-i-1] - '0')
		}
		if i < lenB {
			carry += int(b[lenB-i-1] - '0')
		}
		ans = strconv.Itoa(carry%2) + ans
		carry /= 2
	}
	if carry > 0 {
		ans = "1" + ans
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 快速乘
// x 和 y 是负数，z 是正数
// 判断 z * y >= x 是否成立
func quickAdd(y, z, x int) bool {
	for result, add := 0, y; z > 0; z >>= 1 { // 不能使用除法
		if z&1 > 0 {
			// 需要保证 result + add >= x
			if result < x-add {
				return false
			}
			result += add
		}
		if z != 1 {
			// 需要保证 add + add >= x
			if add < x-add {
				return false
			}
			add += add
		}
	}
	return true
}

func divide(a, b int) int {
	if a == math.MinInt32 { // 考虑被除数为最小值的情况
		if b == 1 {
			return math.MinInt32
		}
		if b == -1 {
			return math.MaxInt32
		}
	}
	if b == math.MinInt32 { // 考虑除数为最小值的情况
		if a == math.MinInt32 {
			return 1
		}
		return 0
	}
	if a == 0 { // 考虑被除数为 0 的情况
		return 0
	}

	// 一般情况，使用二分查找
	// 将所有的正数取相反数，这样就只需要考虑一种情况
	rev := false
	if a > 0 {
		a = -a
		rev = !rev
	}
	if b > 0 {
		b = -b
		rev = !rev
	}

	ans := 0
	left, right := 1, math.MaxInt32
	for left <= right {
		mid := left + (right-left)>>1 // 注意溢出，并且不能使用除法
		if quickAdd(b, mid, a) {
			ans = mid
			if mid == math.MaxInt32 { // 注意溢出
				break
			}
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if rev {
		return -ans
	}
	return ans
}
