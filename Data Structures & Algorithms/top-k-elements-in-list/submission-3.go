func topKFrequent(nums []int, k int) []int {
    m := make(map[int]int)
    for _, v := range nums {
        m[v]++
    }
    
    fb := make ([][]int, len(nums)+1)
    for kk, ff := range m {
        fb[ff] = append(fb[ff], kk)
    }

    rs := []int{}
    for i := len(nums); i > 0 && len(rs) < k; i--{
        if len(fb[i]) > 0 {
            kr := k - len(rs)
            for j := 0; j < kr && j < len(fb[i]); j++ {
                rs = append(rs, fb[i][j])
            }
        }
    }

    return rs
}