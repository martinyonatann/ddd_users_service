package serr

import (
	"errors"
	"fmt"
)

// New serr
func New(msg string) SErr {
	return construct(ErrLevelFatal, ErrCodeNothing, ErrKeyNothing, errors.New(msg), 0)
}

// Newf serr with message binding
func Newf(frmt string, args ...interface{}) SErr {
	return construct(ErrLevelFatal, ErrCodeNothing, ErrKeyNothing, fmt.Errorf(frmt, args...), 0)
}

// News serr from stack skip
func News(skip int, msg string) SErr {
	return construct(ErrLevelFatal, ErrCodeNothing, ErrKeyNothing, errors.New(msg), skip)
}

// Newsl serr from stack skip and error level
func Newsl(skip int, lvl ErrLevel, msg string) SErr {
	return construct(lvl, ErrCodeNothing, ErrKeyNothing, errors.New(msg), skip)
}

// Newsli serr from stack skip, error level and error code
func Newsli(skip int, lvl ErrLevel, code int, msg string) SErr {
	return construct(lvl, code, ErrKeyNothing, errors.New(msg), skip)
}

// Newslik serr from stack skip, error level, error code and error key
func Newslik(skip int, lvl ErrLevel, code int, key string, msg string) SErr {
	return construct(lvl, code, key, errors.New(msg), skip)
}

// Newslikc serr from stack skip, error level, error code, error key and comments
func Newslikc(skip int, lvl ErrLevel, code int, key string, msg string, comment string) SErr {
	errx := construct(lvl, code, key, errors.New(msg), skip)

	if comment == "@" {
		comment = msg
	}

	errx.AddComment(comment)
	return errx
}

// Newsi serr from stack skip and error code
func Newsi(skip int, code int, msg string) SErr {
	return construct(ErrLevelFatal, code, ErrKeyNothing, errors.New(msg), skip)
}

// Newsik serr from stack skip, error code and error key
func Newsik(skip int, code int, key string, msg string) SErr {
	return construct(ErrLevelFatal, code, key, errors.New(msg), skip)
}

// Newsikc serr from stack skip, error code, error key and comment
func Newsikc(skip int, code int, key string, msg string, comment string) SErr {
	errx := construct(ErrLevelFatal, code, key, errors.New(msg), skip)

	if comment == "@" {
		comment = msg
	}

	errx.AddComment(comment)
	return errx
}

// Newsk serr from stack skip and error key
func Newsk(skip int, key string, msg string) SErr {
	return construct(ErrLevelFatal, ErrCodeNothing, key, errors.New(msg), skip)
}

// Newskc serr from stack skip, error key and comment
func Newskc(skip int, key string, msg string, comment string) SErr {
	errx := construct(ErrLevelFatal, ErrCodeNothing, key, errors.New(msg), skip)

	if comment == "@" {
		comment = msg
	}

	errx.AddComment(comment)
	return errx
}

// Newsf serr from stack skip with message binding
func Newsf(skip int, msg string, args ...interface{}) SErr {
	return construct(ErrLevelFatal, ErrCodeNothing, ErrKeyNothing, fmt.Errorf(msg, args...), skip)
}

// Newsc serr from stack skip and comment
func Newsc(skip int, msg string, comment string) SErr {
	errx := construct(ErrLevelFatal, ErrCodeNothing, ErrKeyNothing, errors.New(msg), skip)

	if comment == "@" {
		comment = msg
	}

	errx.AddComment(comment)
	return errx
}

// Newl serr from error level
func Newl(lvl ErrLevel, msg string) SErr {
	return construct(lvl, ErrCodeNothing, ErrKeyNothing, errors.New(msg), 0)
}

// Newlf serr from error level with message binding
func Newlf(lvl ErrLevel, msg string, args ...interface{}) SErr {
	return construct(lvl, ErrCodeNothing, ErrKeyNothing, fmt.Errorf(msg, args...), 0)
}

// Newli serr from error level and error code
func Newli(lvl ErrLevel, code int, msg string) SErr {
	return construct(lvl, code, ErrKeyNothing, errors.New(msg), 0)
}

// Newlik serr from error level, error code and error key
func Newlik(lvl ErrLevel, code int, key string, msg string) SErr {
	return construct(lvl, code, key, errors.New(msg), 0)
}

// Newlikc serr from error level, error code, error key and comment
func Newlikc(lvl ErrLevel, code int, key string, msg string, comment string) SErr {
	errx := construct(lvl, code, key, errors.New(msg), 0)

	if comment == "@" {
		comment = msg
	}

	errx.AddComment(comment)
	return errx
}

// Newi serr from error code
func Newi(code int, msg string) SErr {
	return construct(ErrLevelFatal, code, ErrKeyNothing, errors.New(msg), 0)
}

// Newif serr from error code with message binding
func Newif(code int, frmt string, args ...interface{}) SErr {
	return construct(ErrLevelFatal, code, ErrKeyNothing, fmt.Errorf(frmt, args...), 0)
}

// Newik serr from error code and error key
func Newik(code int, key string, msg string) SErr {
	return construct(ErrLevelFatal, code, key, errors.New(msg), 0)
}

// Newikc serr from error code, error key and comment
func Newikc(code int, key string, msg string, comment string) SErr {
	errx := construct(ErrLevelFatal, code, key, errors.New(msg), 0)

	if comment == "@" {
		comment = msg
	}

	errx.AddComment(comment)
	return errx
}

// Newk serr from error key
func Newk(key string, msg string) SErr {
	return construct(ErrLevelFatal, ErrCodeNothing, key, errors.New(msg), 0)
}

// Newkf serr from error key with message binding
func Newkf(key string, frmt string, args ...interface{}) SErr {
	return construct(ErrLevelFatal, ErrCodeNothing, key, fmt.Errorf(frmt, args...), 0)
}

// Newkc serr from error key and comment
func Newkc(key string, msg string, comment string) SErr {
	errx := construct(ErrLevelFatal, ErrCodeNothing, key, errors.New(msg), 0)

	if comment == "@" {
		comment = msg
	}

	errx.AddComment(comment)
	return errx
}

// Newc serr from comment
func Newc(msg string, comment string) SErr {
	errx := construct(ErrLevelFatal, ErrCodeNothing, ErrKeyNothing, errors.New(msg), 0)

	if comment == "@" {
		comment = msg
	}

	errx.AddComment(comment)
	return errx
}

// Newslic serr from stack skip, error level, error code and comment
func Newslic(skip int, lvl ErrLevel, code int, msg string, comment string) SErr {
	errx := construct(lvl, code, ErrKeyNothing, errors.New(msg), skip)

	if comment == "@" {
		comment = msg
	}

	errx.AddComment(comment)
	return errx
}

// Newslk serr from stack skip, error level and error key
func Newslk(skip int, lvl ErrLevel, key string, msg string) SErr {
	return construct(lvl, ErrCodeNothing, key, errors.New(msg), 0)
}

// Newsic serr from stack skip, error code and comment
func Newsic(skip int, code int, msg string, comment string) SErr {
	errx := construct(ErrLevelFatal, code, ErrKeyNothing, errors.New(msg), skip)

	if comment == "@" {
		comment = msg
	}

	errx.AddComment(comment)
	return errx
}

// Newslc serr from stack skip, error level and comment
func Newslc(skip int, lvl ErrLevel, msg string, comment string) SErr {
	errx := construct(lvl, ErrCodeNothing, ErrKeyNothing, errors.New(msg), skip)

	if comment == "@" {
		comment = msg
	}

	errx.AddComment(comment)
	return errx
}

// Newlkc serr from error level, error key and comment
func Newlkc(lvl ErrLevel, key string, msg string, comment string) SErr {
	errx := construct(lvl, ErrCodeNothing, key, errors.New(msg), 0)

	if comment == "@" {
		comment = msg
	}

	errx.AddComment(comment)
	return errx
}

// Newlic serr from error level, error code and comment
func Newlic(lvl ErrLevel, code int, msg string, comment string) SErr {
	errx := construct(lvl, code, ErrKeyNothing, errors.New(msg), 0)

	if comment == "@" {
		comment = msg
	}

	errx.AddComment(comment)
	return errx
}

// Newlk serr from error level and error key
func Newlk(lvl ErrLevel, key string, msg string) SErr {
	return construct(lvl, ErrCodeNothing, key, errors.New(msg), 0)
}

// Newlc serr from error level and comment
func Newlc(lvl ErrLevel, msg string, comment string) SErr {
	errx := construct(lvl, ErrCodeNothing, ErrKeyNothing, errors.New(msg), 0)

	if comment == "@" {
		comment = msg
	}

	errx.AddComment(comment)
	return errx
}

// NewFromError serr from error
func NewFromError(err error) SErr {
	return construct(ErrLevelFatal, ErrCodeNothing, ErrKeyNothing, err, 0)
}

// NewFromErrors serr from error and stack skip
func NewFromErrors(skip int, err error) SErr {
	return construct(ErrLevelFatal, ErrCodeNothing, ErrKeyNothing, err, skip)
}

// NewFromErrorl serr from error and error level
func NewFromErrorl(lvl ErrLevel, err error) SErr {
	return construct(lvl, ErrCodeNothing, ErrKeyNothing, err, 0)
}

// NewFromErrori serr from error and error code
func NewFromErrori(code int, err error) SErr {
	return construct(ErrLevelFatal, code, ErrKeyNothing, err, 0)
}

// NewFromErrork serr from error and error key
func NewFromErrork(key string, err error) SErr {
	return construct(ErrLevelFatal, ErrCodeNothing, key, err, 0)
}

// NewFromErrorc serr from error and comment
func NewFromErrorc(err error, comment string) SErr {
	errx := construct(ErrLevelFatal, ErrCodeNothing, ErrKeyNothing, err, 0)

	if comment == "@" {
		comment = err.Error()
	}

	errx.AddComment(comment)
	return errx
}
