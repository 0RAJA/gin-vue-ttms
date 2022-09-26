package v1

// 对外访问的接口
type group struct {
	Email     email
	Cinema    cinema
	Movie     movie
	Comment   comment
	Plan      plan
	Seats     seats
	Tags      tags
	User      user
	Upload    upload
	UserMovie userMovie
	Ticket    ticket
	Count     count
	Order     order
}

var Group = new(group)
