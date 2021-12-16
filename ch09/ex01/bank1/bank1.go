package bank1

var deposits = make(chan int)
var balances = make(chan int)
var withdraws = make(chan int)
var canWithdraws = make(chan bool)

func Deposit(amount int) {
	deposits <- amount
}

func Balance() int {
	return <-balances
}

func Withdraw(amount int) bool {
	withdraws <- amount
	return <-canWithdraws
}

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case amount := <-withdraws:
			if balance >= amount {
				balance -= amount
				canWithdraws <- true
			} else {
				canWithdraws <- false
			}
		}
	}
}

func init() {
	go teller()
}
