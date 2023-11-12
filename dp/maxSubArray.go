package dp

func MaxSubArray(nums []int) int {
    var res, sum int 
    for i := 0; i < len(nums); i++{
        sum = Max(nums[i], sum+nums[i])
        res = Max(res, sum)
    }
    return res
}
/*
-2,1,-3,4,-1,2,1,-5,4
*/

func Max(a, b int) int{
    if a < b{
        return b
    }
    return a
}