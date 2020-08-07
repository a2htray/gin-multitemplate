package gin_multitemplate

import (
	"github.com/gin-gonic/gin/render"
	"html/template"
	"path"
	"path/filepath"
	"strings"
)

const (
	defaultIncludeDir = "./templates"
	defaultLayoutDir = "./templates/layouts"
)

// dirFile: return the file paths in a given directory
func dirFiles(dir string, extension string) ([]string, error) {
	files, err := filepath.Glob(path.Join(dir, "*." + extension))
	return files, err
}

// NameFunc is a callback function to point a generation of a template name
func NamedFunc(t *TemplateInfo, file string) string {
	return strings.TrimSuffix(filepath.Base(file), "." + t.Extension)
}

// TemplateInfo is the appointed template information
// The Render struct will parse the templates corresponding to this information
type TemplateInfo struct {
	LayoutDir string
	Extension string
	IncludeDir string
	isParsed bool
	layouts map[string]string
	includes map[string]string
}

// Parse: parse the layout directory and include directory
func (t *TemplateInfo) Parse() *TemplateInfo {
	if t.layouts == nil {
		t.layouts = make(map[string]string)
	}
	if t.includes == nil {
		t.includes = make(map[string]string)
	}
	if layouts, err := dirFiles(t.LayoutDir, t.Extension); err != nil {
		panic(err)
	} else {
		for _, layout := range layouts {
			name := NamedFunc(t, layout)
			t.layouts[name] = layout
		}
	}

	if includes, err := dirFiles(t.IncludeDir, t.Extension); err != nil {
		panic(err)
	} else {
		for _, include := range includes {
			name := NamedFunc(t, include)
			t.includes[name] = include
		}
	}
	t.isParsed = true
	return t
}

// DefaultTemplateInfo ...
func DefaultTemplateInfo() *TemplateInfo {
	t := (&TemplateInfo{
		LayoutDir: defaultLayoutDir,
		IncludeDir: defaultIncludeDir,
		Extension: "html",
		layouts: make(map[string]string),
		includes: make(map[string]string),
	}).Parse()

	return t
}

// Render implements the render.HTMLRender interface
// engine := gin.Default()
// render := NewRender()
// engine.HTMLRender = render
type Render struct {
	tf        *TemplateInfo
	templates map[string]*template.Template
}

func (r *Render) LoadLayoutAndInclude() {
	if !r.tf.isParsed {
		r.tf.Parse()
	}

	layouts := make([]string, 0, len(r.tf.layouts))
	for _, layout := range r.tf.layouts {
		layouts = append(layouts, layout)
	}

	var err error
	for name, include := range r.tf.includes {
		layoutsCopy := make([]string, len(layouts))
		copy(layoutsCopy, layouts)
		layoutsCopy = append(layoutsCopy, include)
		r.templates[name], err = template.ParseFiles(layoutsCopy...)
		if err != nil {
			panic(err)
		}
	}
}

// NewRender ...
func NewRender(templateInfo *TemplateInfo) *Render {
	r := &Render{
		tf: templateInfo,
		templates: make(map[string]*template.Template),
	}
	r.LoadLayoutAndInclude()
	return r
}

// Instance: implements the render.HTMLRender's Instance method
func (r *Render) Instance(name string, data interface{}) render.Render {
	return render.HTML{
		Template: r.templates[name],
		Data: data,
	}
}




