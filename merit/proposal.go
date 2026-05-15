package merit

// Proposal is also known as Candidate or Option.  It received grades from voters.
type Proposal struct {
	// Name of the Proposal.  An empty string is allowed.
	Name string
	// Tally of the grades received by this Proposal, from "worst" grade to "best" grade.
	// An empty list is not allowed, and tallies across proposals must be:
	// 1. Consistent: have the same length, that is they should represent the same amount of grades
	// 2. Balanced: their sum must be the same, that is they should hold the same amount of judgments
	Tally []uint64
}
