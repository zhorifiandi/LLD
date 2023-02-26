package groupmeeting

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/zhorifiandi/LLD/schedulemeeting/domain"
	"github.com/zhorifiandi/LLD/schedulemeeting/usecase/mvpapp"
)

type BaseApp = mvpapp.Application

type Application struct {
	BaseApp
	GroupMeetings map[string]domain.GroupMeeting
	Groups        map[string]domain.Group
}

func NewApplication() (app *Application) {
	baseApp := mvpapp.NewApplication(
		mvpapp.ApplicationInputs{},
	)
	return &Application{
		BaseApp:       *baseApp,
		GroupMeetings: map[string]domain.GroupMeeting{},
		Groups:        map[string]domain.Group{},
	}
}

func (a *Application) CreateGroup(
	name string,
) (group domain.Group) {
	id := (uuid.New()).String()
	group = domain.Group{
		ID:   id,
		Name: name,
	}
	a.Groups[id] = group

	return
}

func (a *Application) GetGroupByID(
	groupID string,
) (group domain.Group) {
	group = a.Groups[groupID]
	return
}

func (a *Application) ShowGroups() {
	fmt.Printf("Groups: %+v\n", a.Groups)
}

func (a *Application) AddEmployeeToGroup(employeeID string, groupID string) {
	employee := a.GetEmployeeByID(employeeID)
	group := a.GetGroupByID(groupID)
	group.Members = append(group.Members, employee)
	a.Groups[groupID] = group
}

func (a *Application) EmployeeCreateMeetingWithGroup(
	CreatorID string,
	StartTime domain.Time,
	EndTime domain.Time,
	ParticipantEmployeeIDs []string,
	ParticipantGroups []domain.GroupInput,
) (meeting domain.Meeting) {
	meeting = a.EmployeeCreateMeeting(
		CreatorID,
		StartTime,
		EndTime,
		ParticipantEmployeeIDs,
	)
	if meeting.ID == "" {
		return
	}

	groups := []domain.Group{}
	IsGroupInputValid := true
	for _, input := range ParticipantGroups {
		group := a.GetGroupByID(input.ID)
		participantCount := 0
		for _, employee := range group.Members {
			isCanAttendMeeting := a.IsEmployeeCanAttendMeeting(
				employee.ID,
				StartTime,
				EndTime,
			)

			if isCanAttendMeeting {
				participantCount += 1
			}

		}

		if participantCount < input.MinQuorum {
			log.Printf("Error: Not quorum!!! \n")
			IsGroupInputValid = false
			break
		}

		groups = append(groups, group)
	}

	if IsGroupInputValid {
		a.GroupMeetings[meeting.ID] = domain.GroupMeeting{
			Meeting:           meeting,
			GroupParticipants: groups,
		}
	} else {
		a.RemoveMeeting(meeting.ID)
	}

	return
}

func (a *Application) ShowGroupMeetings(
	StartTime domain.Time,
	EndTime domain.Time,
) (meetings []domain.GroupMeeting) {
	for _, meeting := range a.GroupMeetings {
		isWithinRange := meeting.StartTime >= StartTime && meeting.EndTime <= EndTime
		if isWithinRange {
			meetings = append(meetings, meeting)
		}
	}
	fmt.Printf("Group Meetings: %+v\n", meetings)
	return
}
