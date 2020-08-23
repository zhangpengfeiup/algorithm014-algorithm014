// 思路：两边hash 表
// 先将数据放入到 hash 表中  value -> index  ?
// 循环遍历 nums, 查看 [target - nums] 在不在hash 表单中
func twoSum(nums []int, target int) []int {
    ret := make([]int, 0)
    m := make(map[int]int)

    for i := 0;i < len(nums);i++ {
        complement := target - nums[i]

        v,ok := m[complement]

        if ok {
            ret = append(ret, i)
            ret = append(ret, v)
        
            return ret
        }

        m[nums[i]] = i
    }

    return ret
}
