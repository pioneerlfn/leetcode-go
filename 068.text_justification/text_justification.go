package main

import "fmt"

// Input:
// words = ["This", "is", "an", "example", "of", "text", "justification."]
// maxWidth = 16
// Output:
// [
//   "This    is    an",
//     3        4   5
//    index = 3, last = 6.
//   "example  of text",
//   "justification.  "
// ]
func fullJustify(words []string, maxWidth int) []string {
	var res []string
	totalWords := len(words)
	var index, last int
	for index < totalWords && last < totalWords {
		index = last
		count := len(words[index])
		last = index
		for {
			last++
			if last == totalWords || count+1+len(words[last]) > maxWidth {
				break
			}
			count += 1 + len(words[last])
		}

		// fmt.Printf("\nlast: %d, count: %d, index: %d\n", last, count, index)

		var line string
		spaceSlot := last - index - 1
		if last != totalWords && spaceSlot != 0 {
			// lines above the last.
			var charCount int
			for i := index; i < last; i++ {
				charCount += len(words[i])
			}
			totalSpace := maxWidth - charCount
			base, remainder := totalSpace / spaceSlot, totalSpace % spaceSlot

			// fmt.Printf("slot: %d, totalSpace: %d, base: %d, remainder: %d\n", spaceSlot, totalSpace, base, remainder)

			for i := index; i < last-1; i ++ {
				line += words[i]
				for i := 0; i < base; i++ {
					line += " "
				}
				if remainder > 0 {
					line += " "
					remainder--
				}
			}
			line += words[last-1]
		} else {
			// last line.
			for i := index; i < last-1; i++ {
				line += words[i]
				line += " "
			}
			line += words[last-1]
			padding := maxWidth - len(line)
			for i := 0; i < padding; i++ {
				line += " "
			}

		}
		res = append(res, string(line))
	}

	return res
}

func main() {
	// words := []string{"This", "is", "an", "example", "of", "text", "justification."}
	// words := []string{"What","must","be","acknowledgment","shall","be"}
	// words := []string{"Science","is","what","we","understand","well","enough","to","explain",
		// "to","a","computer.","Art","is","everything","else","we","do"}
	words := []string{"ask","not","what","your","country","can","do","for","you","ask","what","you","can","do","for","your","country"}
	maxWidth := 16
	res := fullJustify(words, maxWidth)
	fmt.Printf("res: %+q\n", res)
}
