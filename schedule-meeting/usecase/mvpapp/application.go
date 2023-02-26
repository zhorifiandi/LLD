package mvpapp

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/zhorifiandi/LLD/schedulemeeting/domain"
)

type ApplicationInputs struct{}

func NewApplication(inputs ApplicationInputs) (app *Application) {
	return &Application{
		Employees: map[string]domain.Employee{},
		Meetings:  map[string]domain.Meeting{},
	}
}

type Application struct {
	Employees map[string]domain.Employee
	Meetings  map[string]domain.Meeting
}

func (a *Application) CreateEmployee(
	name string,
) domain.Employee {
	employee := domain.Employee{
		ID:   (uuid.New()).String(),
		Name: name,
	}

	a.Employees[employee.ID] = employee

	return employee
}

func (a *Application) GetEmployeeByID(
	id string,
) domain.Employee {
	return a.Employees[id]
}

func (a *Application) ShowEmployees() {
	fmt.Printf("%+v\n", a.Employees)
}

// new 2 4
// new 6 7
//  1     5

func (a *Application) GetConflictedMeetings(
	EmployeeID string,
	StartTime domain.Time,
	EndTime domain.Time,
) (meetings []domain.Meeting) {
	for _, meeting := range a.Meetings {
		isTimeOverlap := !(StartTime.Before(meeting.StartTime) || EndTime.After(meeting.EndTime))

		isSameEmployee := false

		for _, emp := range meeting.Participants {
			if emp.ID == EmployeeID {
				isSameEmployee = true
				break
			}
		}

		log.Printf(">>> %+v %+v\n", isTimeOverlap, isSameEmployee)

		if isTimeOverlap && isSameEmployee {
			meetings = append(meetings, meeting)
		}
	}

	fmt.Printf("Conflicted Meetings: %+v\n", meetings)
	return
}

func (a *Application) IsEmployeeCanAttendMeeting(
	employeeID string,
	StartTime domain.Time,
	EndTime domain.Time,
) bool {
	conflictedMeetings := a.GetConflictedMeetings(
		employeeID,
		StartTime,
		EndTime,
	)

	return len(conflictedMeetings) == 0
}

func (a *Application) EmployeeCreateMeeting(
	CreatorID string,
	StartTime domain.Time,
	EndTime domain.Time,
	ParticipantEmployeeIDs []string,
) (meeting domain.Meeting) {
	participants := []domain.Employee{}
	for _, id := range ParticipantEmployeeIDs {
		employee := a.Employees[id]
		isCanAttendMeeting := a.IsEmployeeCanAttendMeeting(
			employee.ID,
			StartTime,
			EndTime,
		)

		if !isCanAttendMeeting {
			log.Printf("Error: Employee has another meeting\n")
			return
		}

		if employee.ID != "" && isCanAttendMeeting {
			participants = append(participants,
				employee,
			)
		}
	}

	meeting = domain.Meeting{
		ID:           (uuid.New()).String(),
		CreatorID:    CreatorID,
		StartTime:    StartTime,
		EndTime:      EndTime,
		Participants: participants,
	}

	a.Meetings[meeting.ID] = meeting
	return
}

func (a *Application) RemoveMeeting(
	meetingID string,
) {
	delete(a.Meetings, meetingID)
}

func (a *Application) ShowMeetings(
	StartTime domain.Time,
	EndTime domain.Time,
) (meetings []domain.Meeting) {
	for _, meeting := range a.Meetings {
		isWithinRange := (meeting.StartTime.After(StartTime) || meeting.StartTime.Equal(StartTime)) && (meeting.EndTime.Before(EndTime) || meeting.EndTime.Equal(EndTime))
		if isWithinRange {
			meetings = append(meetings, meeting)
		}
	}
	fmt.Printf("Meetings: %+v\n", meetings)
	return
}

func (a *Application) GetAvailableSlots(
	GivenDate domain.Time,
	ParticipantEmployeeIDs []string,
) (
	StartTime domain.Time,
	EndTime domain.Time,
) {
	return
}
