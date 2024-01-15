// Code generated by qtc from "UUID.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/components/view/UUID.html:2
package view

//line views/components/view/UUID.html:2
import "github.com/google/uuid"

//line views/components/view/UUID.html:4
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/components/view/UUID.html:4
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/components/view/UUID.html:4
func StreamUUID(qw422016 *qt422016.Writer, value *uuid.UUID) {
//line views/components/view/UUID.html:5
	if value != nil {
//line views/components/view/UUID.html:6
		qw422016.E().S(value.String())
//line views/components/view/UUID.html:7
	}
//line views/components/view/UUID.html:8
}

//line views/components/view/UUID.html:8
func WriteUUID(qq422016 qtio422016.Writer, value *uuid.UUID) {
//line views/components/view/UUID.html:8
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/view/UUID.html:8
	StreamUUID(qw422016, value)
//line views/components/view/UUID.html:8
	qt422016.ReleaseWriter(qw422016)
//line views/components/view/UUID.html:8
}

//line views/components/view/UUID.html:8
func UUID(value *uuid.UUID) string {
//line views/components/view/UUID.html:8
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/view/UUID.html:8
	WriteUUID(qb422016, value)
//line views/components/view/UUID.html:8
	qs422016 := string(qb422016.B)
//line views/components/view/UUID.html:8
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/view/UUID.html:8
	return qs422016
//line views/components/view/UUID.html:8
}