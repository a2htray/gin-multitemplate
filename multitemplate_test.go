package gin_multitemplate

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNamedFunc(t *testing.T) {
	tpInfo := DefaultTemplateInfo()
	name := NamedFunc(tpInfo, "./a/b/c/d.html")
	assert.Equal(t, "d", name, "name is not equal to 'd'")

	tpInfo.Extension = "gohtml"
	name = NamedFunc(tpInfo, "./a/b/c/d.gohtml")
	assert.Equal(t, "d", name, "name is not equal to 'd'")
}

func TestTemplateInfo_Parse(t *testing.T) {
	tpInfo := &TemplateInfo{
		LayoutDir: "./tests/layouts",
		IncludeDir: "./tests",
		Extension: "html",
	}
	tpInfo.Parse()
	fmt.Println(tpInfo.LayoutDir)
	fmt.Println(tpInfo.IncludeDir)
	fmt.Println(tpInfo.Extension)
	fmt.Println(tpInfo.layouts)
	fmt.Println(tpInfo.includes)
}

func TestNewRender(t *testing.T) {
	tpInfo := &TemplateInfo{
		LayoutDir: "./tests/layouts",
		IncludeDir: "./tests",
		Extension: "html",
	}
	render := NewRender(tpInfo)
	fmt.Println(render.templates)
}