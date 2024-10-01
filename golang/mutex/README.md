# Mutex and Confinement
- Context: Many clients are trying to buy a ticket from a store. Without mutex, sometime the total number of tickets is not correct.

## Step to proceduce
1. comment `mutex.Lock()` and `mutex.Unlock()` in `mutex.go` and `usingConfinement(totalTicket)` in `main.go`
2. run the command

  ```
  for run in {1..10}; do go run *.go | grep "remain ticket" | wc -l; done
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
- We can see, sometime the total sold tickets are incorrect
