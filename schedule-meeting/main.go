package main

import (
	"log"

	"github.com/zhorifiandi/LLD/schedulemeeting/domain"
	"github.com/zhorifiandi/LLD/schedulemeeting/usecase/groupmeeting"
)

type IApplication interface {
	CreateEmployee(name string) domain.Employee
	GetEmployeeByID(empID string) domain.Employee
	ShowEmployees()

	EmployeeCreateMeeting(
		CreatorID string,
		StartTime domain.Time,
		EndTime domain.Time,
		ParticipantEmployeeIDs []string,
	) domain.Meeting
	ShowMeetings(
		StartTime domain.Time,
		EndTime domain.Time,
	) []domain.Meeting
	GetAvailableSlots(
		GivenDate domain.Time,
		ParticipantEmployeeIDs []string,
	) (
		StartTime domain.Time,
		EndTime domain.Time,
	)

	// Req 2
	CreateGroup(name string) domain.Group
	GetGroupByID(groupID string) domain.Group
	ShowGroups()
	AddEmployeeToGroup(employeeID string, groupID string)
	EmployeeCreateMeetingWithGroup(
		CreatorID string,
		StartTime domain.Time,
		EndTime domain.Time,
		ParticipantEmployeeIDs []string,
		ParticipantGroups []domain.GroupInput,
	) domain.Meeting

	ShowGroupMeetings(
		StartTime domain.Time,
		EndTime domain.Time,
	) []domain.GroupMeeting
}

func main() {
	var app IApplication = groupmeeting.NewApplication()
	log.Printf("App is running..... %+v\n", app)

	a := app.CreateEmployee("a")
	b := app.CreateEmployee("b")
	c := app.CreateEmployee("c")

	d1 := app.CreateGroup("d1")
	app.AddEmployeeToGroup(b.ID, d1.ID)
	app.AddEmployeeToGroup(c.ID, d1.ID)

	app.EmployeeCreateMeetingWithGroup(
		a.ID,
		0,
		1,
		[]string{a.ID},
		[]domain.GroupInput{
			{
				ID:        d1.ID,
				MinQuorum: 2,
			},
		},
	)

	app.ShowGroupMeetings(0, 1)
}
