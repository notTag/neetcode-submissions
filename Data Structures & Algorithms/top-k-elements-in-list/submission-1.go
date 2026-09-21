func topKFrequent(nums []int, k int) []int {
    //build frequency map
    m := make(map[int]int)
    for _, v := range nums {
        m[v]++
    }
    
    //build frequency buckets
    freqValSlice := make ([][]int, len(nums)+1)
    for curKey, curFreq := range m {
        freqValSlice[curFreq] = append(freqValSlice[curFreq], curKey)
    }

    //pull k highest elements (rely on golangs ordered slices)
    freqSlice := []int{}
    for i := len(nums); i > 0 && len(freqSlice) < k; {
        if len(freqValSlice[i]) > 0 {
            kRemaining := k - len(freqSlice)
            //build slice of k highest elements
            for j := 0; j < kRemaining && j < len(freqValSlice[i]); j++ {
                freqSlice = append(freqSlice, freqValSlice[i][j])
            }
        }
        i--
    }

    return freqSlice
}