package xing_http

import (
	"errors"
	"github.com/ghts/ghts/lib"
	"strings"
)

func New질의(값 lib.I질의값, ch질의 chan *S질의) *S질의 {
	s := new(S질의)
	s.M값 = 값
	s.Ch응답 = make(chan *S응답, 1)

	ch질의 <- s

	return s
}

type S질의 struct {
	M값   lib.I질의값
	Ch응답 chan *S응답
}

func New응답(값 interface{}) *S응답 {
	switch 값.(type) {
	case error:
		return &S응답{V: nil, E: 값.(error)}
	default:
		return &S응답{V: 값, E: nil}
	}
}

type S응답 struct {
	V interface{}
	E error
}

func (s S응답) G값() interface{} {
	if s.E == nil {
		return s.V
	} else {
		return s.E
	}
}

func New질의_JSON(값 lib.I질의값, ch질의 chan *S질의_JSON) *S질의_JSON {
	s := new(S질의_JSON)
	s.M값 = 값
	s.Ch응답 = make(chan *S응답_JSON, 1)

	ch질의 <- s

	return s
}

type S질의_JSON struct {
	M값   lib.I질의값
	Ch응답 chan *S응답_JSON
}

func New응답_JSON(값 interface{}) *S응답_JSON {
	switch 값.(type) {
	case error:
		return &S응답_JSON{V: nil, E: 값.(error).Error()}
	default:
		return &S응답_JSON{V: 값, E: ""}
	}
}

type S응답_JSON struct {
	V interface{}
	E string
}

func (s S응답_JSON) Error() error {
	if strings.TrimSpace(s.E) == "" {
		return nil
	} else {
		return errors.New(s.E)
	}
}
