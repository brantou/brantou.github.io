package main


import "fmt"

func parent(i int) int {
  return i / 2
}

func left(i int) int {
  return 2 * i
}

func right(i int) int {
  return 2*i + 1
}

func hsort(nums []int) {
  if len(nums) < 2 {
    return
  }

  p := len(nums) / 2
  for {
    if p < 1 {
      break
    }

    l := left(p)
    r := right(p)
    if l < len(nums) && nums[l-1] > nums[p-1] {
      nums[l-1], nums[p-1] = nums[p-1], nums[l-1]
    }

    if r < len(nums) && nums[r-1] > nums[p-1] {
      nums[r-1], nums[p-1] = nums[p-1], nums[r-1]
    }

    p -= 1
  }
	return

  size := len(nums)
  for {
    nums[0], nums[size-1] = nums[size-1], nums[0]
    size -= 1
    if size < 2 {
      break
    }

    p := 1
    for {
      if p > size/2 {
        break
      }
      if size == 6 && p == 3 {
          return
      }
      fmt.Println(size, p, nums)

      l := left(p)
      r := right(p)
      if l < size && r < size {
        if nums[l-1] > nums[r-1] {
          if nums[l-1] > nums[p-1] {
            nums[l-1], nums[p-1] = nums[p-1], nums[l-1]
            p = l
            continue
          } else {
            break
          }
        } else {
          if nums[r-1] > nums[p-1] {
            nums[r-1], nums[p-1] = nums[p-1], nums[r-1]
            p = r
            continue
          } else {
            break
          }
        }
      }

      if l < size {
        if nums[l-1] > nums[p-1] {
          nums[l-1], nums[p-1] = nums[p-1], nums[l-1]
          p = l
        }
        break
      }
    }
  }

  return
}

func main() {
  nums := []int{2, 4, 6, 3, 5, 7, 1, 10, 8, 9}
  hsort(nums)
  fmt.Println(nums)
}
