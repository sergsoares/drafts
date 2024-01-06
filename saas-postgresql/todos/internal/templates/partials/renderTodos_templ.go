// Code generated by templ@v0.2.282 DO NOT EDIT.

package partials

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

// GoExpression
import (
	"github.com/stackus/todos/internal/domain"
)

func RenderTodos(todos []*domain.Todo) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_1 := templ.GetChildren(ctx)
		if var_1 == nil {
			var_1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		// Element (standard)
		_, err = templBuffer.WriteString("<form")
		if err != nil {
			return err
		}
		// Element Attributes
		_, err = templBuffer.WriteString(" hx-post=\"/todos/sort\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(" hx-trigger=\"end\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(" class=\"block p-0 mb-2 text-lg\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(">")
		if err != nil {
			return err
		}
		// Element (standard)
		_, err = templBuffer.WriteString("<div")
		if err != nil {
			return err
		}
		// Element Attributes
		_, err = templBuffer.WriteString(" id=\"todos\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(" class=\" sortable\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(">")
		if err != nil {
			return err
		}
		// For
		for _, todo := range todos {
			// TemplElement
			err = RenderTodo(todo).Render(ctx, templBuffer)
			if err != nil {
				return err
			}
		}
		// Element (standard)
		_, err = templBuffer.WriteString("<div")
		if err != nil {
			return err
		}
		// Element Attributes
		_, err = templBuffer.WriteString(" id=\"no-todos\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(" class=\"hidden first:block first:pb-2 first:pt-3\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(">")
		if err != nil {
			return err
		}
		// Element (standard)
		_, err = templBuffer.WriteString("<p>")
		if err != nil {
			return err
		}
		// Text
		var_2 := `Congrats, you have no todos! Or... do you? 😰`
		_, err = templBuffer.WriteString(var_2)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</p>")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</div>")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</div>")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</form>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = io.Copy(w, templBuffer)
		}
		return err
	})
}
