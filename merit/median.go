package merit

// EvenMedianStrategy decides where we will draw the median line when the amount of judgments is even.
type EvenMedianStrategy int

const (
	// DeadCenter always draws the median line at 50%.
	DeadCenter EvenMedianStrategy = iota
	// FavorAdhesion draws the median line in the middle of the adjacent adhesion judgment.
	// It is akin to choosing the high median.
	FavorAdhesion
	// FavorContestation draws the median line in the middle of the adjacent contestation judgment.
	// It is akin to choosing the low median.  This is the default behavior in Majority Judgment.
	FavorContestation
)
