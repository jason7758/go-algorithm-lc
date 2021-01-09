package main

import "fmt"

func multiply(num1 string, num2 string) string {
	if num1 =="0" || num2 == "0" {
		return "0"
	}
	//string 转换为[]byte， 容易取相应位上的值
	bsi := []byte(num1)
	bsj := []byte(num2)
	// temp 的长度只可能为 len(num1) + len(num2) 或 len(num1)+len(num2) -1
	//取最大长度，免得不够
	temp := make([]int, len(num1) + len(num2))
	for i := len(bsi) -1; i >= 0; i-- {
		for j := len(bsj) - 1; j >= 0; j-- {
			mul := int(bsi[i] - '0') * int(bsj[j] - '0')
			p1 := i + j //当前数乘位所在位数
			p2 := i + j +1 //当前乘数的进位所在位
			sum := temp[p2] + mul //当前乘积的数+进位过来数，就是本次的总数
			//fmt.Printf("i=%d,j=%d,tmp=%v",i, j, temp)
			//fmt.Printf("bsi[i]=%d,bsj[j]=%d,temp[p1]=%d,temp[p2]=%d,mul=%d,sum=%d,\n",bsi[i]-'0',bsj[j]-'0', temp[p1], temp[p2],mul, sum)
			temp[p1] +=  sum / 10 //当前位的数
			temp[p2] =  sum % 10 //要进位的数
			//fmt.Printf("tmp=%v\n",temp)
		}
	}
	if temp[0] == 0 { //去除数字前导的0
		temp = temp[1:]
	}
	//转换结果
	//temp 选用[]int, 而不是[]byte,是因为
	//在Go中， byte的基础结构是uint8， 最大值是255
	//不考虑进位的位， temp会溢出
	res := make([]byte, len(temp))
	for i := 0; i < len(temp); i++ {
		res[i] = '0' + byte(temp[i])
	}

	return string(res)
}
func main() {
	num1 := "123"
	num2 := "456"
	mul := multiply(num1, num2)
	fmt.Println(mul)
}
