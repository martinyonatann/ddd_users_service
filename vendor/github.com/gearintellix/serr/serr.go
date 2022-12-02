package serr

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"

	gerr "github.com/go-errors/errors"
)

type (
	// SErr interface
	SErr interface {
		Error() string
		Cause() error

		Level() ErrLevel
		Code() int
		Key() string
		Comments() string
		CommentStack() []string
		Payload() ErrPayload

		Callers() []uintptr
		StackFrames() []gerr.StackFrame

		Type() string
		File() string
		Line() int
		FN() string
		Package() string

		String() string
		ColoredString() string

		SetKey(key string)
		SetCode(code int)
		SetLevel(lvl ErrLevel)
		AddComment(msg string)
		AddCommentf(msg string, opts ...interface{})
		ApplyPayload(payload ErrPayload)
		SetPayload(key string, value interface{})
	}

	// ErrPayload type
	ErrPayload map[string]interface{}

	// ErrLevel type
	ErrLevel string
)

type serr struct {
	level    ErrLevel
	err      error
	key      string
	code     int
	comments []string
	payload  ErrPayload
	frames   []gerr.StackFrame
	stack    []uintptr
}

const (
	// ErrLevelFatal constant for fatal error level
	ErrLevelFatal ErrLevel = "fatal"

	// ErrLevelWarn constant for warning error level
	ErrLevelWarn ErrLevel = "warn"

	// ErrLevelInfo constant for info error level
	ErrLevelInfo ErrLevel = "info"
)

const (
	// ErrKeyNothing constant for empty error key
	ErrKeyNothing string = "-"

	// ErrKeyUnexpected constant for unexpected error key
	ErrKeyUnexpected string = "unexpected"

	// ErrKeyExpected constant for expected error key
	ErrKeyExpected string = "expected"

	// ErrCodeNothing constant for empty error code
	ErrCodeNothing int = 0
)

func construct(level ErrLevel, code int, key string, err error, skip int) *serr {
	stack := make([]uintptr, 50)
	length := runtime.Callers(3+skip, stack[:])

	res := &serr{
		level:    level,
		err:      err,
		key:      key,
		code:     code,
		comments: []string{},
		payload:  make(ErrPayload),
		stack:    stack[:length],
	}
	return res
}

// Error to get error message
func (ox serr) Error() string {
	return fmt.Sprintf("%+v", ox.err)
}

// Cause to get original error
func (ox serr) Cause() error {
	return ox.err
}

// Level to get error level
func (ox serr) Level() ErrLevel {
	return ox.level
}

// Code to get error code
func (ox serr) Code() int {
	return ox.code
}

// Key to get error key
func (ox serr) Key() string {
	return ox.key
}

// Comments to get error comments
func (ox serr) Comments() string {
	return strings.Join(ox.comments, ", ")
}

// CommentStack to get error comment stack
func (ox serr) CommentStack() []string {
	return ox.comments
}

// Payload to get error payload
func (ox serr) Payload() ErrPayload {
	return ox.payload
}

// Callers to get error callers stack
func (ox serr) Callers() []uintptr {
	return ox.stack
}

// StackFrames to get error stack frames
func (ox *serr) StackFrames() []gerr.StackFrame {
	if ox.frames == nil {
		ox.frames = make([]gerr.StackFrame, len(ox.stack))

		for i, pc := range ox.stack {
			ox.frames[i] = gerr.NewStackFrame(pc)
		}
	}

	return ox.frames
}

// Type get error type
func (ox serr) Type() string {
	return reflect.TypeOf(ox.err).String()
}

// File get error file path
func (ox serr) File() string {
	frames := ox.StackFrames()
	if len(frames) > 0 {
		return frames[0].File
	}
	return ""
}

// Line get error line
func (ox serr) Line() int {
	frames := ox.StackFrames()
	if len(frames) > 0 {
		return frames[0].LineNumber
	}
	return 0
}

// FN to get error function name
func (ox serr) FN() string {
	frames := ox.StackFrames()
	if len(frames) > 0 {
		return frames[0].Name
	}
	return ""
}

// Package to get error package name
func (ox serr) Package() string {
	frames := ox.StackFrames()
	if len(frames) > 0 {
		return frames[0].Package
	}
	return ""
}

// String to get formated error message
func (ox serr) String() string {
	comments := ""

	if ox.Code() != 0 {
		comments += fmt.Sprintf(" <code: %d>", ox.Code())
	}

	if isExists(ox.Key(), []string{"-", "!"}) {
		comments += fmt.Sprintf(" <key: %s>", ox.Key())
	}

	if len(ox.comments) > 0 {
		comments += fmt.Sprintf(" <comments: %s>", ox.Comments())
	}

	return fmt.Sprintf(
		StandardFormat(),
		ox.FN(),
		ox.File(),
		ox.Line(),
		ox.Error(),
		comments,
	)
}

// ColoredString to get formated error message with color (cli color code)
func (ox serr) ColoredString() string {
	comments := ""

	if ox.Code() != 0 {
		comments += fmt.Sprintf(" <code: %d>", ox.Code())
	}

	if !isExists(ox.Key(), []string{"", "-", "!"}) {
		comments += fmt.Sprintf(" <key: %s>", ox.Key())
	}

	if len(ox.comments) > 0 {
		comments += fmt.Sprintf(" <comments: %s>", ox.Comments())
	}

	return fmt.Sprintf(
		StandardColorFormat(),
		ox.FN(),
		ox.File(),
		ox.Line(),
		ox.Error(),
		comments,
	)
}

// SetKey to set error key
func (ox *serr) SetKey(key string) {
	ox.key = key
}

// SetCode to set error code
func (ox *serr) SetCode(code int) {
	ox.code = code
}

// SetLevel to set error level
func (ox *serr) SetLevel(lvl ErrLevel) {
	ox.level = lvl
}

// AddComment to add error comment
func (ox *serr) AddComment(msg string) {
	ox.comments = append(ox.comments, msg)
}

// AddCommentf to add error comment with string binding
func (ox *serr) AddCommentf(msg string, opts ...interface{}) {
	ox.comments = append(ox.comments, fmt.Sprintf(msg, opts...))
}

// ApplyPayload to apply error payload
func (ox *serr) ApplyPayload(load ErrPayload) {
	if ox.payload == nil {
		ox.payload = make(ErrPayload)
	}

	if load != nil {
		for k, v := range load {
			ox.payload[k] = v
		}
	}
}

// SetPayload to set error payload
func (ox *serr) SetPayload(key string, value interface{}) {
	if ox.payload == nil {
		ox.payload = make(ErrPayload)
	}

	ox.payload[key] = value
}
