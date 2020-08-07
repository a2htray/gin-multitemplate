### gin-multitemplate

A multiple template render package based on [gin](https://github.com/gin-gonic/gin). The
implementation is copy from [multitemplate](https://github.com/gin-contrib/multitemplate) to
my own go project.

The Usage is simple. The layout directory and the includes directory are passed to
the render system.

### Examples

```go
import (
    multitemplate "github.com/a2htray/gin-multitemplate"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    
    router.HTMLRender = multitemplate.NewRender(&multitemplate.TemplateInfo{
        LayoutDir: "./template/layout",
        IncludeDir: "./template",
        Extension: "html",
    })
    
    router.GET("/index", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index", gin.H{})
    })
    
    err = router.Run(":8080")
    
    if err != nil {
        log.Fatal("starting the website ... failed")
    }
}
```