package benchmarks

import (
	h "github.com/go-andiamo/aitch"
	"github.com/go-andiamo/aitch/context"
	. "github.com/go-andiamo/aitch/html"
	"io"
	"testing"
)

var htmlBasic = Document("en",
	[]any{
		Meta(Charset("utf-8")),
		TitleElement(Text("My document")),
	},
	[]any{
		Div(
			P(Class("foo"), Text("bar")),
		),
	})

var htmlWithMergedAttrs = Document("en",
	[]any{
		Meta(Charset("utf-8")),
		TitleElement(Text("My document")),
	},
	[]any{
		Div(
			P(Class("foo"), Text("bar"), Class("bar")),
		),
	})

var htmlWithConditionalAttrs = Document("en",
	[]any{
		Meta(Charset("utf-8")),
		TitleElement(Text("My document")),
	},
	[]any{
		Div(
			P(Class("foo"), Text("bar"), h.When("yes", Class("bar"), Id(1))),
		),
	})

func BenchmarkRendering(b *testing.B) {
	b.ReportAllocs()
	// use null writer so that write output doesn't skew allocs
	ctx := context.New(&nullWriter{}, nil, nil)
	for i := 0; i < b.N; i++ {
		_ = htmlBasic.Render(ctx)
	}
}

func BenchmarkRendering_MergedAttrs(b *testing.B) {
	b.ReportAllocs()
	// use null writer so that write output doesn't skew allocs
	ctx := context.New(&nullWriter{}, nil, nil)
	for i := 0; i < b.N; i++ {
		_ = htmlWithMergedAttrs.Render(ctx)
	}
}

func BenchmarkRendering_ConditionalAttrs(b *testing.B) {
	b.ReportAllocs()
	// use null writer so that write output doesn't skew allocs
	ctx := context.New(&nullWriter{}, map[string]any{"yes": true}, nil)
	for i := 0; i < b.N; i++ {
		_ = htmlWithConditionalAttrs.Render(ctx)
	}
}

type nullWriter struct{}

var _ io.Writer = &nullWriter{}

func (nw *nullWriter) Write(p []byte) (n int, err error) {
	//do nothing
	return len(p), nil
}
