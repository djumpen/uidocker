package surv

import survey "gopkg.in/AlecAivazis/survey.v1"

type question struct {
	Message string
	Options []string
	Result  string
}

type prompt struct {
	Message string
}

func NewQuestion(m string, o []string) *question {
	return &question{
		Message: m,
		Options: o,
	}
}

func (q *question) Ask() (*question, error) {
	name := "answer"
	var qs = []*survey.Question{
		{
			Name: name,
			Prompt: &survey.Select{
				Message: q.Message,
				Options: q.Options,
			},
		},
	}

	var ans = struct {
		Answer string
	}{}

	err := survey.Ask(qs, &ans)
	if err != nil {
		return nil, err
	}

	q.Result = ans.Answer

	return q, nil
}

func (q *question) Answer() string {
	return q.Result
}
