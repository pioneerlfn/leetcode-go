package _04_median_of_sorted_arrays

import "math"

// 讲解见： https://blog.csdn.net/chen_xinjia/article/details/69258706
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)
	if len1 > len2 {
		return findMedianSortedArrays(nums2, nums1)
	}
	// now we know that len1 <= len2.
	if len1 == 0 {
		if len2 > 1 {
			// include len2 is odd and even.
			return float64(nums2[len2/2] + nums2[(len2 - 1)/2]) / 2
		} else {
			return float64(nums2[0])
		}
	}

	size := len1 + len2
	// cutX(X=L,R,1,2) means items number after the cut.
	cutL, cutR := 0, len1
	var cut1, cut2 int
	
	var L1, R1, L2, R2 int
	
	for cut1 <= len1 {
		cut1 = (cutR - cutL) / 2 + cutL
		cut2 = size / 2 - cut1

//		L1, R1 = nums1[cut1-1], nums1[cut1]
//		L2, R2 = nums2[cut2-1], nums2[cut2]
		// because of cut1-1 and cut2 -1, our 2 lines code become nearly 20 lines! so ugly!
		if cut1 == 0 {
			L1 = math.MinInt64
		} else {
			L1 = nums1[cut1-1]
		}
		if cut2 == 0 {
			L2 = math.MinInt64
		} else {
			L2 = nums2[cut2-1]
		}
		if cut1 == len1 {
			R1 = math.MaxInt64
		} else {
			R1 = nums1[cut1]
		}
		if cut2 == len2 {
			R2 = math.MaxInt64
		} else {
			R2 = nums2[cut2]
		}

		if L1 > R2 {
			cutR = cut1 -1
		} else if L2 > R1 {
			cutL = cut1 + 1
		} else {
			if size % 2 == 0 {
				// note: the result of min and max should first converted to float64, then average.
				return (float64(max(L1, L2)) + float64(min(R1, R2))) / 2
			} else {
				return float64(min(R1, R2))
			}
		}
	}
	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
