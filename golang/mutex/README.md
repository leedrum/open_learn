# Mutex and Confinement

## Mutex
- Many clients are trying to buy a ticket from a store. Without mutex, sometime the total number of tickets is not correct.
- Step to proceduce
1. comment `mutex.Lock()` and `mutex.Unlock()`
2. run the command

  ```
  for run in {1..10}; do go run mutex.go | grep "remain ticket" | wc -l; done
  ```
3. the restul will be similar to the log

```
204
200
200
201
200
201
202
200
202
200
```

4. To solve the problem. We undo the step 1 and re-run the step 2.


## Confinement
