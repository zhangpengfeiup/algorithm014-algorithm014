学习笔记

1. 懵逼的时候怎么办呢？
   从最简单的情况开始分析 -> 范化 -> 最近重复子问题 -> 找重复性

   数学归纳法

   找重复性：计算机适合做重复的事情：if/else for while recursion
   

2. 数组遍历

嵌套循环，注意，第一个判断可以写成 len(height) - 1,height 为数组
```
  for i := 0;i < len(height) - 1;i++ {
         for j := i + 1;j < len(height);j++ {
             currentArea := getArea(i, j, height)
             ret = max(currentArea, ret)
         }
     }
```

two Point ： 左右夹逼

```
     i := 0
     j := len(height) - 1

    for i < j {
        currentArea := getArea(i, j, height)

        ret = max(currentArea, ret)

        if height[i] > height[j] {
            j--
        } else {
            i++
        } 
    }
```

3. 一点小想法

时间管理：分清优先级，时间是挤出来的

承认自己写的不好，自己不怎么会写代码

hash 自己原本以为很简单，思路也很简单，但是自己用起来，去使用的时候，发现自己不会使用，一直报错，所以，归根结底，还是基本功太差，需要多写多练，不要眼高手低。

错误的想法：题目只做一遍，死磕，不要死磕，给自己5～8分钟时间思考，如果不知道怎么解决,换一种方式思考，看题解，看别人怎么思考和解决的，承认自己不会，去学会，去动手
