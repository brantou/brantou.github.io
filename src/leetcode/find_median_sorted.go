package main
import (
	"fmt"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
  sum_size := len(nums1) + len(nums2)
  nums1_start, nums1_end := 0, len(nums1)
  nums2_start, nums2_end := 0, len(nums2)
  for {
    fmt.Println(nums1_start, nums1_end, nums2_start, nums2_end)
    if nums1_start < nums1_end && nums2_start < nums2_end {
      nums1_mid := (nums1_start + nums1_end) / 2
      nums2_mid := (nums2_start + nums2_end) / 2

      if nums1[nums1_mid] < nums2[nums2_mid] {
        nums1_start = nums1_mid
        nums2_end = nums2_mid
      } else if nums1[nums1_mid] > nums2[nums2_mid] {
        nums1_end = nums1_mid
        nums2_start = nums2_mid
      } else if nums1[nums1_mid] == nums2[nums2_mid] {
        return float64(nums1[nums1_mid])
      }

      continue
    }

    if nums1_start < nums1_end {
      nums1_mid := (nums1_start + nums1_end) / 2

      if nums1[nums1_mid] < nums2[nums2_start] {
        nums1_start = nums1_mid
      } else if nums1[nums1_mid] > nums2[nums2_start] {
        nums1_end = nums1_mid
      } else if nums1[nums1_mid] == nums2[nums2_start] {
        return float64(nums2[nums2_start])
      }

      continue
    }

    if nums2_start < nums2_end {
      nums2_mid := (nums2_start + nums2_end) / 2

      if nums2[nums2_mid] < nums1[nums1_start] {
        nums2_start = nums2_mid
      } else if nums2[nums2_mid] > nums1[nums1_start] {
        nums2_end = nums2_mid
      } else if nums2[nums2_mid] == nums1[nums1_start] {
        return float64(nums1[nums1_start])
      }

      continue
    }

    if nums1_start >= nums1_end && nums2_start >= nums2_end {
      if nums1[nums1_start] < nums2[nums2_start] {
        if (nums1_start+1)*2 == sum_size {
          return float64(nums1[nums1_start]+nums2[nums2_start]) / float64(2)
        }

        return float64(nums1[nums1_start])
      } else {
        if (nums2_start+1)*2 == sum_size {
          return float64(nums1[nums1_start]+nums2[nums2_start]) / float64(2)
        }
        return float64(nums2[nums2_start])
      }
    }
  }
}
func main() {
  // fmt.Println(findMedianSortedArrays([]int{1,3}, []int{2}))
  // fmt.Println(findMedianSortedArrays([]int{1,2}, []int{3, 4}))
  fmt.Println(findMedianSortedArrays([]int{2,4,8}, []int{3,6,9}))
}
