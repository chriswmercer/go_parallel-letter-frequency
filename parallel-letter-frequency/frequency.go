package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

//ConcurrentFrequency counts the frequency of each rune for a given list of items
func ConcurrentFrequency(items []string) FreqMap {
	var results FreqMap = make(map[rune]int)
	var channel = make(chan FreqMap)

	for _, e := range items {
		go func(data string) {
			result := Frequency(data)
			channel <- result
		}(e)
	}

	for range items {
		result := <-channel
		for k, v := range result {
			results[k] += v
		}
	}
	return results
}