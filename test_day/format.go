package fmt

import "unicode/utf8"

const (
	ldigit = "0123456789abcdefx"
	udigit = "0123456789ABCDEFX"
)

const (
	signed   = true
	unsigned = false
)

type fmtFlags struct {
	widPresent  bool
	precPresent bool
	minus       bool
	plus        bool
	sharp       bool
	space       bool
	zero        bool

	plusV  bool
	sharpV bool
}

type fmt struct {
	buf *buffer

	fmtFlags

	wid  int
	prec int

	intbuf [68]byte
}

func (f *fmt) clearflags() {
	f.fmtFlags = fmtFlags{}
}

func (f *fmt) init(buf *buffer) {
	f.buf = buf
	f.clearflags()
}

func (f *fmt) writePadding(n int) {
	if n <= 0 {
		return
	}
	buf := *f.buf
	oldLen := len(buf)
	newLen := oldLen + n
	if newLen > cap(buf) {
		buf = make(buffer, cap(buf)*2+n)
		copy(buf, *f.buf)
	}

	padByte := byte(' ')
	if f.zero {
		padByte = byte('0')
	}
	padding := buf[oldLen:newLen]
	for i := range padding {
		padding[i] = padByte
	}
	*f.buf = buf[:newLen]
}

func (f *fmt) pad(b []byte) {
	if !f.widPresent || f.wid == 0 {
		f.buf.Write(b)
		return
	}
	width := f.wid - utf8.RuneCount(b)
	if !f.minus {
		f.writePadding(width)
		f.buf.Write(b)
	} else {
		f.buf.Write(b)
		f.writePadding(width)
	}
}
