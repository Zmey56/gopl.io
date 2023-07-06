// Exercise 9.1:
// Add a function Withdraw(amount int) bool to the
// gopl.io/ch9/bank1 program.
//
// The result should indicate whether the transaction succeeded or failed
// due to insufficient funds.
//
// The message sent to the monitor goroutine must contain both the amount
// to withdraw and a new channel over which the monitor goroutine can send the
// boolean result back to Withdraw.
package bank

var deposits = make(chan int)                  // send amount to deposit
var balances = make(chan int)                  // receive balance
var withdrawalRequests = make(chan Withdrawal) // receive withdrawal requests

func Deposit(amount int) {
	deposits <- amount
}

func Balance() int {
	return <-balances
}

func Withdraw(amount int) bool {
	resultChannel := make(chan bool)
	withdrawal := Withdrawal{amount, resultChannel}
	withdrawalRequests <- withdrawal
	return <-resultChannel
}

type Withdrawal struct {
	amount int
	result chan<- bool
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case withdrawal := <-withdrawalRequests:
			if withdrawal.amount > balance {
				withdrawal.result <- false
			} else {
				balance -= withdrawal.amount
				withdrawal.result <- true
			}
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}
