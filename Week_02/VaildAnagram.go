/*
    建立一个 hash 表，然后遍历 s 字符串，然后每个位置++
    然后遍历另一个字符串然后每个位置--
    最后比那里 map 看是否有为1的，有则 return false
    否则 return true
*/
func isAnagram(s string, t string) bool {
    if len(s) !=  len(t) {
        return false
    }

     m := make(map[byte]int)

     for i := 0;i < len(s);i++ {
         m[s[i] - 'a']++;
     }

     for i := 0;i < len(t);i++ {
         m[t[i] - 'a']--;
     }

     for _,v := range m {
         if v != 0 {
             return false
         }
     }

     return true
}
