// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Nav() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<nav class=\"navbar navbar-expand-lg navbar-dark bg-dark\"><div class=\"container-fluid\"><a class=\"navbar-brand\" href=\"#\">Navbar</a> <button class=\"navbar-toggler\" type=\"button\" data-bs-toggle=\"collapse\" data-bs-target=\"#navbarSupportedContent\" aria-controls=\"navbarSupportedContent\" aria-expanded=\"false\" aria-label=\"Toggle navigation\"><span class=\"navbar-toggler-icon\"></span></button><div class=\"collapse navbar-collapse\" id=\"navbarSupportedContent\"><ul class=\"navbar-nav me-auto mb-2 mb-lg-0\"><li class=\"nav-item\"><a class=\"nav-link active\" aria-current=\"page\" href=\"/\">Home</a></li><li class=\"nav-item\"><a class=\"nav-link\" href=\"/about\">About</a></li><li class=\"nav-item dropdown\"><a class=\"nav-link dropdown-toggle\" href=\"#\" id=\"navbarDropdown\" role=\"button\" data-bs-toggle=\"dropdown\" aria-expanded=\"false\">Rooms</a><ul class=\"dropdown-menu\" aria-labelledby=\"navbarDropdown\"><li><a class=\"dropdown-item\" href=\"/general\">General's Quarters</a></li><li><hr class=\"dropdown-divider\"></li><li><a class=\"dropdown-item\" href=\"/major\">Major's suite</a></li></ul></li><li class=\"nav-item\"><a class=\"nav-link\" href=\"/reservation\">Book now</a></li><li class=\"nav-item\"><a class=\"nav-link\" href=\"/contact\">Contact</a></li></ul></div></div></nav>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}