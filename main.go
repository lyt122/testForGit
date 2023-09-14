package main

import "math"

type ListNode struct {
	Val  int
	Next *ListNode
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func maxSumDivThree(nums []int) int {
	// 创建一个数组用于保存累积的最大和，其中索引为0、1、2分别表示模数0、1、2
	dp := [3]int{0, math.MinInt32, math.MinInt32}
	for _, num := range nums {
		// 复制当前的dp状态，用于更新
		temp := dp
		for i := 0; i < 3; i++ {
			// 更新dp状态
			newIndex := (i + num) % 3
			dp[newIndex] = max(dp[newIndex], temp[i]+num)
		}
	}
	// 返回模数为0的累积最大和
	return dp[0]
}

func main() {
	nums := []int{2, 2, 2, 5}
	maxSumDivThree(nums)
	// make只传入一个参数指定长度，则容量和长度相等。以下输出："len: 10, cap: 10"
	/*s := make([]int, 10)
	fmt.Printf("len: %d, cap: %d\n", len(s), cap(s))

	// make 传入长度和容量。以下输出："len: 10, cap: 15"
	s = make([]int, 10, 15)
	fmt.Printf("len: %d, cap: %d\n", len(s), cap(s))

	s = []int{99: 0}
	fmt.Printf("len: %d, cap: %d\n", len(s), cap(s))
	//s1:=s[low:high:max]low指定开始元素下标，high指定结束元素下标，max指定切片能增长到的元素下标
	//这三个参数都可以省略，low省略默认从下标0开始，high省略默认为最后一个元素下标，max省略默认是底层数组或切片的容量（这里也要注意max不能小于high）。
	s1 := s
	fmt.Printf("len: %d, cap: %d\n", len(s1), cap(s1))
	s2 := s[23:67]
	fmt.Printf("len: %d, cap: %d\n", len(s2), cap(s2))*/
	/*s := []int{1, 2, 3, 4, 5}
	s1 := s[2:4]
	fmt.Printf("s: %v", s)

	s1 = append(s1, 5)
	fmt.Printf("s1: %v", s1)
	var fan map[string]string //创建集合,[string]是键，string是值
	fan = make(map[string]string)
	//map插入key-value对
	fan["One"] = "666"
	fan["Four"] = "999"
	//使用键输出
	for value := range fan {
		fmt.Println(value, "is", fan[value])
	}
	//查看元素在集合中是否存在
	val, ok := fan["Two"] //如果确定是真实的,则存在,否则不存在
	if ok {
		fmt.Println("fanTwo is", val)
	} else {
		fmt.Println("fanTwo not exist")
	}
	delete(fan, "One")*/

}
