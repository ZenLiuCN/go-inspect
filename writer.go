package inspect

import (
	"bytes"
	"fmt"
	"log"
	"strings"
	"sync"
	"unicode"
)

var (
	Debug = false
	Trace = false
)

func debugf(format string, args ...any) {
	if Debug {
		_ = log.Output(2, "[DEBUG] "+fmt.Sprintf(format, args...))
	}
}
func tracef(n int, format string, args ...any) {
	if Trace {
		_ = log.Output(n+2, "[TRACE] "+fmt.Sprintf(format, args...))
	}
}

var (
	p = sync.Pool{New: func() any { return new(writer).init() }}
)

type (
	Writer interface {
		Buffer() *bytes.Buffer
		Reset()
		UnWrite(n int)
		Format(format string, args ...any) Writer
		Append(w Writer) Writer
		Release()
		NotEmpty() bool
		LF() Writer    //write line-feed and with Indent if Indent not zero
		IndLF() Writer //write line-feed and with Indent if Indent not zero
		Indent(n int) Writer
		IndInc(n int) Writer
		IndDec(n int) Writer
		Indents() int
		IndentString() string
		Bytes() []byte
		fmt.Stringer
	}
	writer struct {
		i    string
		buf  *bytes.Buffer
		free bool
	}
)

func (s *writer) Bytes() []byte {
	return s.buf.Bytes()
}
func (s *writer) String() string {
	return s.buf.String()
}
func (s *writer) NotEmpty() bool {
	return s.buf.Len() > 0
}
func (s *writer) LF() Writer {
	if s.free {
		panic("writer released")
	}
	s.buf.WriteByte('\n')
	return s
}
func (s *writer) IndLF() Writer {
	if s.free {
		panic("writer released")
	}
	s.buf.WriteByte('\n')
	s.buf.WriteString(s.i)
	return s
}
func (s *writer) Indent(n int) Writer {
	if s.free {
		panic("writer released")
	}
	s.i = strings.Repeat("\t", n)
	return s
}
func (s *writer) IndDec(n int) Writer {
	if s.free {
		panic("writer released")
	}
	s.i = strings.Repeat("\t", s.Indents()-n)
	return s
}
func (s *writer) IndInc(n int) Writer {
	if s.free {
		panic("writer released")
	}
	s.i = strings.Repeat("\t", s.Indents()+n)
	return s
}
func (s *writer) Indents() int {
	if s.free {
		panic("writer released")
	}
	return len(s.i)
}
func (s *writer) IndentString() string {
	if s.free {
		panic("writer released")
	}
	return s.i
}

func (s *writer) Buffer() *bytes.Buffer {
	if s.free {
		panic("writer released")
	}
	return s.buf
}

func (s *writer) Format(format string, args ...any) Writer {
	if s.free {
		panic("writer released")
	}
	_, _ = fmt.Fprintf(s.buf, format, args...)
	return s
}
func (s *writer) Reset() {
	if s.free {
		panic("writer released")
	}
	s.i = ""
	s.buf.Reset()
}
func (s *writer) UnWrite(n int) {
	if s.free {
		panic("writer released")
	}
	if n > s.buf.Len() {
		return
	}
	if n == s.buf.Len() {
		s.buf.Reset()
		return
	}
	b := s.buf.Bytes()
	s.buf.Reset()
	s.buf.Write(b[:len(b)-n])
}

func (s *writer) Append(w Writer) Writer {
	if s.free {
		panic("writer released")
	}
	_, _ = w.Buffer().WriteTo(s.buf)
	return s
}

func GetWriter() Writer {
	w := p.Get().(*writer)
	w.free = false
	return w
}
func (s *writer) init() *writer {
	if s.buf == nil {
		s.buf = bytes.NewBuffer(make([]byte, 0, 255))
	} else {
		s.buf.Reset()
	}
	return s
}
func (s *writer) Release() {
	if s.free {
		return
	}
	s.free = true
	s.i = ""
	s.buf.Reset()
	p.Put(s)
}
func CamelCase(s string) string {
	for i := 0; i < len(s); i++ {
		if i > 0 && !unicode.IsUpper(rune(s[i])) {
			if i == 1 {
				return strings.ToLower(s[:1]) + s[1:]
			} else {
				return strings.ToLower(s[:i-1]) + s[i-1:]
			}
		}
	}
	return strings.ToLower(s)
}
func PascalCase(s string) string {
	for i := 0; i < len(s); i++ {
		if i > 0 && unicode.IsLower(rune(s[i])) {
			if i == 1 {
				return strings.ToUpper(s[:1]) + s[1:]
			} else {
				return strings.ToUpper(s[:i-1]) + s[i-1:]
			}
		}
	}
	return strings.ToUpper(s[0:0]) + s[1:]
}
