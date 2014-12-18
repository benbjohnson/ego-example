package main
import (
"fmt"
"io"
)
//line index.ego:1
 func RenderIndex(w io.Writer, widgets []*Widget) error  {
//line index.ego:2
_, _ = fmt.Fprintf(w, "\n\n<html>\n<body>\n  <h1>Widgets for Sale!</h1>\n\n  <ul>\n    ")
//line index.ego:8
 for _, widget := range widgets { 
//line index.ego:9
_, _ = fmt.Fprintf(w, "\n      <li>")
//line index.ego:9
_, _ = fmt.Fprintf(w, "%v",  widget.Name )
//line index.ego:9
_, _ = fmt.Fprintf(w, " for $")
//line index.ego:9
_, _ = fmt.Fprintf(w, "%v",  widget.Price )
//line index.ego:9
_, _ = fmt.Fprintf(w, "</li>\n    ")
//line index.ego:10
 } 
//line index.ego:11
_, _ = fmt.Fprintf(w, "\n  </ul>\n</body>\n</html>\n")
return nil
}
