func search(nums []int, target int) int {
    left,right := 0,len(nums) - 1

    for left <= right {
        //fmt.Println(left, right)
        mid := left + (right - left) / 2

        //fmt.Println(mid)

        if nums[mid] == target {
            return mid
        }

        // 左半边有序
        if nums[mid] >= nums[left] {
            if nums[left] <= target && target < nums[mid] {
                right = mid - 1
            }else {
                left = mid + 1
            }
        }else {
            if nums[mid] < target && target <= nums[right] {
                left = mid + 1
            }else {
                right = mid - 1;
            }
        }
    }

    return -1;
}
