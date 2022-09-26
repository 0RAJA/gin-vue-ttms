package logic

type group struct {
	Cinema    cinema
	Movie     movie
	Auto      auto
	Plan      plan
	Email     email
	User      user
	Comment   comment
	Upload    up
	Seats     seats
	UserMovie userMovie
	Tags      tags
	Ticket    ticket
	Count     count
	Order     order
}

var Group = new(group)
