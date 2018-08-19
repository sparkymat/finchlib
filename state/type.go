package state

type Type string

const (
	Accepted  = Type("accepted")
	Approved  = Type("approved")
	Archived  = Type("archived")
	Completed = Type("completed")
	Declined  = Type("declined")
	Expired   = Type("expired")
	Fetched   = Type("fetched")
	Pending   = Type("pending")
)

var Types = [...]Type{
	Accepted,
	Approved,
	Archived,
	Completed,
	Declined,
	Expired,
	Fetched,
	Pending,
}
