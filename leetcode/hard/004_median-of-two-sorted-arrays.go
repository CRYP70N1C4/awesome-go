package main

import (
	"common"
	"fmt"
)

func findMedianSortedArrays(nums1, nums2 []int) float64 {
	iMin, iMax := 0, len(nums2)
	for iMin < iMax {

		i := (iMin + iMax) / 2
		j := (len(nums1)+len(nums2)+1)/2 - i
		fmt.Println(iMin, iMax, i, j)

		if nums1[i-1] <= nums2[j] && nums2[j-1] <= nums1[i] {
			tmp := common.Max(nums1[i-1], nums2[j-1]) + common.Min(nums1[i], nums2[j])
			return float64(tmp) / 2.0
		} else if nums1[i-1] > nums2[j] {
			iMax = i
		} else {
			iMin = j
		}

	}
	return -999
}
