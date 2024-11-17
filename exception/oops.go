package exception

import (
	"errors"
	"runtime"
)

type Oops struct {
	code         int
	parent       error
	calledAtLine int
	calledAtFile string
	status       string
	message      string
	info         string
	data         interface{}
}

func New(code int, status string, message string) Oops {
	_, file, line, ok := runtime.Caller(1)

	if !ok {
		file = "???"
		line = 0
	}

	return Oops{
		code:         code,
		status:       status,
		calledAtFile: file,
		calledAtLine: line,
		message:      message,
		info:         message,
	}
}

func (b Oops) Is(err error) bool {
	target, ok := err.(Oops)
	return ok && b.status == target.status
}

func (b Oops) Error() string {
	return b.message
}

func (b Oops) SetInfo(info string) Oops {
	b.info = info
	return b
}

func (b Oops) SetData(data interface{}) Oops {
	b.data = data
	return b
}

func (b Oops) Here() Oops {
	_, file, line, ok := runtime.Caller(1)

	if !ok {
		file = "???"
		line = 0
	}

	b.calledAtFile = file
	b.calledAtLine = line

	return b
}

func (b *Oops) fillDataInfoFromErr(err error) {
	var o Oops

	if errors.As(err, &o) {
		b.info = o.info
		b.data = o.data

		return
	}
}

func (b Oops) Wrap(err error) Oops {
	b.parent = err
	b.fillDataInfoFromErr(err)

	return b
}

func (b Oops) HereWrap(err error) Oops {
	_, file, line, ok := runtime.Caller(1)

	if !ok {
		file = "???"
		line = 0
	}

	b.calledAtFile = file
	b.calledAtLine = line
	b.parent = err
	b.fillDataInfoFromErr(err)

	return b
}

func (b Oops) Unwrap() error {
	return b.parent
}

func (b Oops) SetFromError(err error) Oops {
	var o Oops

	if errors.As(err, &o) {
		b.info = o.info
		b.data = o.data

		return b
	}

	b.info = err.Error()

	return b
}

func (b Oops) CalledAtFile() string {
	return b.calledAtFile
}

func (b Oops) CalledAtLine() int {
	return b.calledAtLine
}

func NewNotFoundError(message string) Oops {
	return New(404, "NOT_FOUND", message)
}

func NewBadRequestError(message string) Oops {
	return New(400, "BAD_REQUEST", message)
}

func NewDuplicatedData(message string) Oops {
	return New(409, "DUPLICATED_DATA", message)
}

func NewUnauthenticatedError(message string) Oops {
	return New(401, "UNAUTHORIZED", message)
}
