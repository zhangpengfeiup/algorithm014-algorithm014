func threeSum(nums []int) [][]int {
    ret := make([][]int, 0)

    if len(nums) < 3 {
        return ret
    }

    sort.Ints(nums)

    for i := 0;i < len(nums) - 2;i++ {
        if nums[i] > 0 {
            return ret
        }

        if i > 0 && nums[i] == nums[i - 1] {
            continue
        }

        target := -nums[i];
        left := i + 1
        right := len(nums) - 1
        
        for left < right  {
            // fmt.Println(left, right)
            if nums[left] + nums[right] == target {
                tmp := make([]int, 0)
                tmp = append(tmp, nums[i])
                tmp = append(tmp, nums[left])
                tmp = append(tmp, nums[right])

                ret = append(ret, tmp)

                // 过滤相同元素，将right 向左移动
                for (left < right) && (nums[right] == nums[right - 1]) {
                    right--
                }

                // 过滤相同元素后，将left 向右移动
                for (left < right) && (nums[left] == nums[left + 1]) {
                    left++
                }

                // 可以借鉴和参考别人的程序和代码，但是还是要弄清楚你自己要做什么，和人生也一样，你要知道自己要什么，别人仅仅只是给你参考，真正怎么走还是要靠你自己和自己内心真正的想法
                right--;
                left++;
            }else if nums[left] + nums[right] > target {
                right--
            }else {
                left++
            }
        }
    }

    return ret
}
