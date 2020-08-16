func plusOne(digits []int) []int {
    len := len(digits)

    for i := len -1;i >= 0;i-- {
        digits[i]++;
        digits[i] %= 10

        if digits[i] !=0 {
            return digits
        }
    }
    // 刚开始一直很疑惑，走到这里，不就是所有的数组前面都要加一个1了么，还是漏掉了上面的return, 如果进位数字在中间或者是后面的话，也就是最高位没有进1的时候直接返回了
    // 还有处理很巧妙，看了很多代码，最后处理的很巧妙，将首位置为1，然后从最后将一个0 push进去
    digits[0] = 1
	digits = append(digits, 0)
    return digits
}
