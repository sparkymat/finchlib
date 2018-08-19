package state

type Type string

const (
	Pending   = Type("pending")
	Approved  = Type("approved")
	Declined  = Type("declined")
	Expired   = Type("expired")
	Completed = Type("completed")
	Archived  = Type("archived")
	Fetched   = Type("fetched")
)

var Types = [...]Type{
	Pending,
	Approved,
	Declined,
	Expired,
	Completed,
	Archived,
	Fetched,
}
