package internal

import (
	"context"
	"fmt"
	"html/template"
	"net/url"
	"path"
	"strings"

	"github.com/elazarl/go-bindata-assetfs"
	"github.com/kihamo/shadow"
	"github.com/kihamo/shadow/components/dashboard"
)

const (
	TemplatePostfix       = ".html"
	TemplateLayoutsDir    = "templates/layouts"
	TemplateViewsDir      = "templates/views"
	TemplateDefaultLayout = "base"
)

type Renderer struct {
	baseLayouts map[string]string
	globals     map[string]interface{}
	views       map[string]map[string]*template.Template
	funcs       template.FuncMap
}

func NewRenderer() *Renderer {
	r := &Renderer{
		baseLayouts: map[string]string{},
		globals:     map[string]interface{}{},
		views:       map[string]map[string]*template.Template{},
	}

	r.funcs = template.FuncMap{
		"raw":        r.funcRaw,
		"add":        r.funcAdd,
		"mod":        r.funcMod,
		"replace":    r.funcReplace,
		"staticHTML": r.funcStaticHTML,
		"date_since": shadow.DateSinceAsMessage,
	}

	return r
}

func (r *Renderer) AddFunc(name string, f interface{}) {
	r.funcs[name] = f
}

func (r *Renderer) AddBaseLayouts(f *assetfs.AssetFS) error {
	files, err := r.getTemplateFiles(TemplateLayoutsDir, f)
	if err != nil {
		return err
	}

	for name, content := range files {
		r.baseLayouts[strings.TrimSuffix(name, TemplatePostfix)] = content
	}

	return nil
}

func (r *Renderer) AddGlobalVar(n string, v interface{}) {
	r.globals[n] = v
}

func (r *Renderer) AddComponents(c string, f *assetfs.AssetFS) error {
	baseComponent := template.New("_component").Funcs(r.funcs)

	// layouts
	for name, content := range r.baseLayouts {
		baseComponent.New(name).Parse(content)
	}

	if files, err := r.getTemplateFiles(TemplateLayoutsDir, f); err == nil {
		for name, content := range files {
			tplName := strings.TrimSuffix(name, TemplatePostfix)

			tpl := baseComponent.Lookup(tplName)
			if tpl == nil {
				tpl.New(tplName)
			}

			tpl.Parse(content)
		}
	}

	// views
	files, err := r.getTemplateFiles(TemplateViewsDir, f)
	if err != nil {
		return nil
	}

	views := map[string]*template.Template{}
	for name, content := range files {
		view, err := baseComponent.Clone()
		if err != nil {
			return err
		}

		if view, err = view.Parse(content); err != nil {
			return err
		}

		views[name] = view
	}

	r.views[c] = views

	return nil
}

func (r *Renderer) Render(ctx context.Context, c, v string, d map[string]interface{}) error {
	return r.RenderLayout(ctx, c, v, TemplateDefaultLayout, d)
}

func (r *Renderer) RenderLayout(ctx context.Context, c, v, l string, d map[string]interface{}) error {
	component, ok := r.views[c]
	if !ok {
		return fmt.Errorf("Templates for component \"%s\" not found", c)
	}

	view, ok := component[v+TemplatePostfix]
	if !ok {
		return fmt.Errorf("Template \"%s\" for component \"%s\" not found", v, c)
	}

	data := r.getContextVariables(ctx)

	for i := range r.globals {
		data[i] = r.globals[i]
	}

	if d != nil {
		for i := range d {
			data[i] = d[i]
		}
	}

	return view.ExecuteTemplate(dashboard.ResponseFromContext(ctx), l, data)
}

func (r *Renderer) getContextVariables(ctx context.Context) map[string]interface{} {
	request := dashboard.RequestFromContext(ctx)

	return map[string]interface{}{
		"Request": request,
		"User":    request.User(),
	}
}

func (r *Renderer) getTemplateFiles(d string, f *assetfs.AssetFS) (map[string]string, error) {
	files, err := f.AssetDir(d)
	if err != nil {
		return nil, err
	}

	templates := make(map[string]string, 0)

	for _, file := range files {
		if !strings.HasSuffix(file, TemplatePostfix) {
			continue
		}

		content, err := f.Asset(d + "/" + file)
		if err != nil {
			continue
		}

		templates[file] = string(content)
	}

	return templates, nil
}

func (r *Renderer) funcRaw(x string) template.HTML {
	return template.HTML(x)
}

func (r *Renderer) funcAdd(x, y int) (interface{}, error) {
	return x + y, nil
}

func (r *Renderer) funcMod(x, y int) (bool, error) {
	return x%y == 0, nil
}

func (r *Renderer) funcReplace(input, from, to string) string {
	return strings.Replace(input, from, to, -1)
}

func (r *Renderer) funcStaticHTML(file string) template.HTML {
	if file == "" {
		return template.HTML(file)
	}

	u, err := url.Parse(file)
	if err != nil {
		return template.HTML(file)
	}

	ext := path.Ext(u.Path)

	switch strings.ToLower(ext) {
	case ".css":
		return template.HTML("<link href=\"" + file + "\" rel=\"stylesheet\">")
	case ".js":
		return template.HTML("<script src=\"" + file + "\"></script>")
	}

	return template.HTML(file)
}
