package main

import (
	"fmt"
	"slices"
	"strings"
	"sync"
)

type Set = map[string]bool

/* Entities */
type User struct {
	id         int
	name       string
	Reputation *UserReputation
}

func (u *User) SetID(id int) {
	u.id = id
}

type QuestionHandler struct {
	mu           sync.Mutex
	questionList []*Question // for ease of development, so we can direct reference
}

func (qh *QuestionHandler) Save(q *Question) {
	qh.mu.Lock()
	defer qh.mu.Unlock()
	q.ID = len(qh.questionList)
	qh.questionList = append(qh.questionList, q)
}

func (qh *QuestionHandler) Find(query *SearchQuery) (ql []*Question) {
	qh.mu.Lock()
	defer qh.mu.Unlock()
	if query == nil {
		return qh.questionList
	}

	for _, question := range qh.questionList {
		isTitleMatchKeyword := query.keyword != "" && strings.Contains(
			strings.ToLower(question.Title),
			strings.ToLower(query.keyword),
		)

		isDescriptionMatchKeyword := query.keyword != "" && strings.Contains(
			strings.ToLower(question.Description),
			strings.ToLower(query.keyword),
		)

		tagStrings := []string{}
		for _, tag := range question.tags {
			tagStrings = append(tagStrings, strings.ToLower(tag.Name))
		}
		isTagMatched := query.tag != "" && slices.Contains(tagStrings, strings.ToLower(query.tag))

		isUserNameMatched := query.userName != "" && strings.EqualFold(question.Author.name, query.userName)

		if isTitleMatchKeyword || isDescriptionMatchKeyword || isTagMatched || isUserNameMatched {
			ql = append(ql, question)
			continue
		}
	}

	return ql
}

func (qh *QuestionHandler) ShowQuestions(ql []*Question) {
	fmt.Println("List of Questions:")
	for i, q := range ql {
		fmt.Printf("\n%d. ", i)
		q.Show()
	}
	fmt.Println("")
}

type ICommentable interface {
	SaveComment(author *User, c *Comment)
}

type Question struct {
	mu          sync.Mutex
	ID          int
	Title       string
	Description string
	Author      *User
	answers     []*Answer
	comments    []*Comment
	votes       map[*Vote]bool
	tags        []*Tag
}

func (q *Question) Show() {
	q.mu.Lock()
	defer q.mu.Unlock()
	fmt.Printf("[%s] %s (by %s)\n", q.Title, q.Description, q.Author.name)
	fmt.Printf("\tVotes: %d\n", len(q.votes))
	fmt.Printf("\tTags: %s\n", q.tags)
	fmt.Print("\tAnswers:\n")
	for _, answer := range q.answers {
		fmt.Printf("\t- ")
		answer.Show()
	}
	fmt.Print("\n\tComments:\n")
	for _, comment := range q.comments {
		fmt.Printf("\t- ")
		comment.Show()
	}
}

func (q *Question) AddAnswer(a *Answer) {
	q.mu.Lock()
	defer q.mu.Unlock()
	a.QuestionID = q.ID
	q.answers = append(q.answers, a)
}

func (q *Question) AddVotes(v *Vote) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.votes == nil {
		q.votes = map[*Vote]bool{}
	}

	v.QuestionID = q.ID
	q.votes[v] = true
}

func (q *Question) SaveComment(author *User, c *Comment) {
	q.mu.Lock()
	defer q.mu.Unlock()
	c.Author = author
	q.comments = append(q.comments, c)
}

type Answer struct {
	mu          sync.Mutex
	QuestionID  int
	Author      *User
	Title       string
	Description string
	comments    []*Comment
}

func (a *Answer) Show() {
	a.mu.Lock()
	defer a.mu.Unlock()
	fmt.Printf("[%s] %s (by %s)\n", a.Title, a.Description, a.Author.name)
	fmt.Printf("\t\tComments:\n")
	for _, comment := range a.comments {
		fmt.Printf("\t\t- ")
		comment.Show()
	}
}

func (a *Answer) SaveComment(author *User, c *Comment) {
	a.mu.Lock()
	defer a.mu.Unlock()
	c.Author = author
	a.comments = append(a.comments, c)
}

type Comment struct {
	Title       string
	Description string
	Author      *User
}

func (c *Comment) Show() {
	fmt.Printf("[%s] %s (by %s)\n", c.Title, c.Description, c.Author.name)
}

type Vote struct {
	QuestionID int
	Voter      *User
}

type Tag struct {
	Name string
}

type UserReputation struct {
	mu        sync.Mutex
	comments  []*Comment
	votes     []*Vote
	questions []*Question
	answers   []*Answer
}

func (ur *UserReputation) Store(t any) {
	ur.mu.Lock()
	defer ur.mu.Unlock()
	switch v := t.(type) {
	default:
		fmt.Printf("unexpected type %T", v)
	case *Comment:
		ur.comments = append(ur.comments, t.(*Comment))
	case *Vote:
		ur.votes = append(ur.votes, t.(*Vote))
	case *Question:
		ur.questions = append(ur.questions, t.(*Question))
	case *Answer:
		ur.answers = append(ur.answers, t.(*Answer))
	}
}

func (ur *UserReputation) GetScore() int {
	ur.mu.Lock()
	defer ur.mu.Unlock()
	score := 0

	score += len(ur.questions) * 2
	score += len(ur.answers) * 3
	score += len(ur.votes) * 1
	score += len(ur.comments) * 1

	return score
}

/* Use Cases */
type Application struct {
	users           []*User
	questionHandler QuestionHandler
}

// Base requirements
func (app *Application) RegisterUser(u *User) {
	u.id = len(app.users)
	u.Reputation = &UserReputation{}
	app.users = append(app.users, u)
}

func (app *Application) ListUsers() {
	fmt.Println("List of Users:")
	for _, u := range app.users {
		fmt.Printf("  %d. %+v (score: %d)\n", u.id, u.name, u.Reputation.GetScore())
	}
	fmt.Println("")
}

// Requirement 1. Users can post questions, answer questions, and comment on questions and answers.
// Requirement 3. Questions should have tags associated with them.
func (app *Application) PostQuestion(u *User, q *Question) {
	q.Author = u
	app.questionHandler.Save(q)
	u.Reputation.Store(q)
}
func (app *Application) AnswerQuestion(u *User, q *Question, answer *Answer) {
	answer.Author = u
	q.AddAnswer(answer)
	u.Reputation.Store(answer)
}
func (app *Application) Comment(u *User, p ICommentable, c *Comment) {
	p.SaveComment(u, c)
	u.Reputation.Store(c)
}

// Requirement 2. Users can vote on questions and answers.
func (app *Application) Vote(u *User, q *Question) {
	vote := Vote{
		QuestionID: q.ID,
		Voter:      u,
	}
	q.AddVotes(&vote)
	u.Reputation.Store(vote)
}

// Requirement 4. Users can search for questions based on keywords, tags, or user profiles.
type SearchQuery struct {
	keyword  string
	tag      string
	userName string
}

func (app *Application) SearchQuestions(q *SearchQuery) {
	fmt.Printf("Query: %+v\n", q)
	ql := app.questionHandler.Find(q)
	app.questionHandler.ShowQuestions(ql)
}

func main() {
	app := Application{}
	u1 := User{name: "budi"}
	app.RegisterUser(&u1)
	u2 := User{name: "acan"}
	app.RegisterUser(&u2)
	app.ListUsers()

	businessTag := &Tag{Name: "BUSINESS"}
	sweTag := &Tag{Name: "SOFTWARE ENGINEERING"}

	// Requirement 1 and 3
	q1 := Question{Title: "How to become RICH?", Description: "What's your strategy to become rich?", tags: []*Tag{businessTag}}
	app.PostQuestion(&u1, &q1)
	a1 := Answer{
		Title: "Hmm", Description: "I don't know....",
	}
	app.AnswerQuestion(&u2, &q1, &a1)
	app.Comment(&u1, &a1, &Comment{Title: "Well", Description: "Better to not reply an answer but a comment.."})
	app.Comment(&u1, &q1, &Comment{Title: "Ok...", Description: "Moved to this comment"})

	q2 := Question{Title: "How to crack system design interview?", Description: "I really want to join FAANG company...", tags: []*Tag{sweTag}}
	app.PostQuestion(&u2, &q2)

	// Requirement 2
	app.Vote(&u1, &q1)
	app.Vote(&u2, &q1)
	app.Vote(&u1, &q2)

	// Show all
	app.SearchQuestions(nil)

	// Requirement 4
	app.SearchQuestions(&SearchQuery{keyword: "rich"})
	app.SearchQuestions(&SearchQuery{keyword: "faang"})
	app.SearchQuestions(&SearchQuery{tag: "software engineering"})
	app.SearchQuestions(&SearchQuery{userName: "ACAN"})

	// Requirement 5
	app.ListUsers()

}
