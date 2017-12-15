// this file was generated by gomacro command: import _b "bufio"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"bufio"
)

// reflection: allow interpreted code to import "bufio"
func init() {
	Packages["bufio"] = Package{
	Binds: map[string]Value{
		"ErrAdvanceTooFar":	ValueOf(&bufio.ErrAdvanceTooFar).Elem(),
		"ErrBufferFull":	ValueOf(&bufio.ErrBufferFull).Elem(),
		"ErrFinalToken":	ValueOf(&bufio.ErrFinalToken).Elem(),
		"ErrInvalidUnreadByte":	ValueOf(&bufio.ErrInvalidUnreadByte).Elem(),
		"ErrInvalidUnreadRune":	ValueOf(&bufio.ErrInvalidUnreadRune).Elem(),
		"ErrNegativeAdvance":	ValueOf(&bufio.ErrNegativeAdvance).Elem(),
		"ErrNegativeCount":	ValueOf(&bufio.ErrNegativeCount).Elem(),
		"ErrTooLong":	ValueOf(&bufio.ErrTooLong).Elem(),
		"MaxScanTokenSize":	ValueOf(bufio.MaxScanTokenSize),
		"NewReadWriter":	ValueOf(bufio.NewReadWriter),
		"NewReader":	ValueOf(bufio.NewReader),
		"NewReaderSize":	ValueOf(bufio.NewReaderSize),
		"NewScanner":	ValueOf(bufio.NewScanner),
		"NewWriter":	ValueOf(bufio.NewWriter),
		"NewWriterSize":	ValueOf(bufio.NewWriterSize),
		"ScanBytes":	ValueOf(bufio.ScanBytes),
		"ScanLines":	ValueOf(bufio.ScanLines),
		"ScanRunes":	ValueOf(bufio.ScanRunes),
		"ScanWords":	ValueOf(bufio.ScanWords),
	},Types: map[string]Type{
		"ReadWriter":	TypeOf((*bufio.ReadWriter)(nil)).Elem(),
		"Reader":	TypeOf((*bufio.Reader)(nil)).Elem(),
		"Scanner":	TypeOf((*bufio.Scanner)(nil)).Elem(),
		"SplitFunc":	TypeOf((*bufio.SplitFunc)(nil)).Elem(),
		"Writer":	TypeOf((*bufio.Writer)(nil)).Elem(),
	},Untypeds: map[string]string{
		"MaxScanTokenSize":	"int:65536",
	},Wrappers: map[string][]string{
		"ReadWriter":	[]string{"Available","Buffered","Discard","Flush","Peek","Read","ReadByte","ReadBytes","ReadFrom","ReadLine","ReadRune","ReadSlice","ReadString","Reset","UnreadByte","UnreadRune","Write","WriteByte","WriteRune","WriteString","WriteTo",},
	},
	}
}