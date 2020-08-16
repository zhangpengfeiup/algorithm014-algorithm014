func moveZeroes(nums []int)  {
     // Count zeros num
     num := 0
     for i := 0;i < len(nums);i++ {
         if nums[i] == 0 {
             num++;
         }
     }
     
     arr := make([]int, 0)
     // 开辟一个新数组，然后将不为0的先放入，放入以后，最后将为0的再放进去
     for i := 0;i < len(nums);i++ {
         if (nums[i] != 0) {
             arr = append(arr, nums[i])
         }
     }

     for i := 0;i < num;i++ {
         arr = append(arr, 0)
     }
     
     for i := 0;i < len(nums);i++ {
         nums[i] = arr[i]
     }
}
