最终使用 method1、method4、method5、method6 方法

### method1、method5、method6 都要掌握

1. method6 把两个有序切片合并成一个有序切片 代码片段 代码行数最少，最简洁
2. method1 直接返回中间数组。 method4 比 method1 多了一步把中间数组拷贝到arr的过程
3. method1 merge 使用 append，method4 merge 使用 index 操作数组
4. method5 对比 method4 ，method5 把 merge 拆成单独函数