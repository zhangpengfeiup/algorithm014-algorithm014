func dfs(nums []int,pos int, list []int,visited []bool,result *[][]int) {
     // terminator
     if pos == len(nums) {
         ans := make([]int, len(list))
         copy(ans, list)
         *result = append(*result, ans)
         return
     }

     for i := 0;i < len(nums);i++ {
         if visited[i] {
             continue;
         }

         // process current logic
         list = append(list, nums[i])
         visited[i] = true

         // drill down
         dfs(nums, pos + 1, list, visited, result)
    
         list = list[0:len(list) - 1]
         visited[i] = false
     }   
}

func permute(nums []int) [][]int {
     result := make([][]int ,0)
     list := make([]int, 0)
     visited := make([]bool, len(nums))
     pos := 0

    dfs(nums, pos, list, visited, &result)

    return result
}
