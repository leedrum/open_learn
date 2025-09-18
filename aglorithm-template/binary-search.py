def binary_search(list, low, high, x):
  while low < high:
    mid = low + (high - low) // 2
    if list[mid] == x:
      return mid + 1
    elif list[mid] < x:
      low = mid + 1
    else:
      high = mid - 1
  return -1

list = [1,2,4,5,7,9]
print(binary_search(list, list[0], list[len(list) - 1], 5)) # return 4
