package bank3

var (
	mu sync.Mutex
	balance int
)

func Deposit(amount int){
	mu.Lock()
	balance = balance + amount
	mu.Unloclk()
}

func Balance() int{
	mu.Lock()
	b := balance
	mu.Unloclk()
	return b
}