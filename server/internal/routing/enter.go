package routing

type group struct {
	Cinema    cinema
	Movie     movie
	Plan      plan
	Email     email
	Comment   comment
	User      user
	Upload    upload
	Seats     seats
	Tags      tags
	UserMovie userMovie
	Ticket    ticket
	Count     count
	Order     order
}

var Group = new(group)
