package poll

import (
	"math/rand"
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/matthewkappus/rosterUpdate/src/types"
)

// Ballot stores questions and students. Result URL set by google service
type Ballot struct {
	ID            string
	Expires       time.Time
	Question      string
	CorrectAnswer string
	// getCorrectAnswerIndex() for res ID
	ResultURL string
	Answers   []string
	Class     *types.Class
}

// New takes a class and returns a Ballot
func New(c *types.Class) *Ballot {
	return &Ballot{
		ID: uuid.Must(uuid.NewV4()).String(),
		// Expire in an hour
		Expires: time.Now().Add(time.Hour),
		Class:   c,
	}
}

// AddAnswers sets ballot answers and shuffles their order
func (b *Ballot) AddAnswers(correct string, incorrect ...string) *Ballot {
	b.CorrectAnswer = correct
	b.Answers = make([]string, len(incorrect)+1)
	for i := range incorrect {
		b.Answers[i] = incorrect[i]
	}
	b.Answers[len(b.Answers)-1] = correct

	rand.Shuffle(len(b.Answers), func(i, j int) {
		b.Answers[i], b.Answers[j] = b.Answers[j], b.Answers[i]
	})

	return b
}

// todo Service.NewBallot(class)
