func max(a int,b int)int {
    if a > b {
        return a
    }
    return b
}

func canJump(nums []int) bool {
    if len(nums) == 0 {
        return true
    }

    jumpMax := 0

    for i := 0;i < len(nums);i++ {
        if i <= jumpMax {
            jumpMax = max(jumpMax, i + nums[i])
        }
        if jumpMax >= len(nums) - 1 {
            return true
        }
    }
    return false
}
