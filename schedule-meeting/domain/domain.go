package domain

type Time = int

type GroupInput struct {
	ID        string
	MinQuorum int
}

type Group struct {
	ID      string
	Name    string
	Members []Employee
}

type Employee struct {
	ID   string
	Name string
}

type Meeting struct {
	ID           string
	CreatorID    string
	StartTime    Time
	EndTime      Time
	Participants []Employee
}

type GroupMeeting struct {
	Meeting
	GroupParticipants []Group
}
