/*
1. 暴力法
2. 排序后，二分法 这样没有办法保证下标，所以没有办法使用二分
*/
func twoSum(nums []int, target int) []int {
      ret := make([]int, 0)
      sort.Ints(nums)
      
      left := 0
      right := len(nums)  - 1

      for left < right {
          mid := left + (right - left) / 2

          if (nums[left] + nums[right] == target) {
              ret = append(ret, left)
              ret = append(ret, right)
              return ret
          }else if (nums[left] + nums[right] > target) {
              right = mid;
          }else {
              // nums[left] + nums[right] < target
              left = mid
          }
      }

      return ret  
}
