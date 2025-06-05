# OpenSlides Projector Service

WIP

## Slides

To create new slides certain steps need to be done. 
This section will give a brief overview on how slides are structured in the projector service.

The following files are relevant for a slide of type `slide_name`:

```
pkg/
├─ projector/
│  ├─ slide/
│  │  ├─ <slide_name>.go
│  │  ├─ slide.go
templates/
├─ slides/
│  ├─ <slide_name>.html
web/
├─ src/
│  ├─ slide/
│  │  ├─ <slide_name>.css
│  │  ├─ <slide_name>.js
│  ├─ projector.js

```

### Slide handler

The projector decides on which slide to display by the field `type` of a projection.

A projection type needs to be registered in `slide.go` in the function `New` to be handled by the projector service. 
Projections not registered there will be ignored. 

```go
	routes["projection_type"] = ProjectionTypeSlideHandler
```

Projection handlers are functions that provide the data needed for a projection. 
Handlers should be created in an individual file per projection type in `pkg/projector/slide/<projection_type>.go`.

A handler receives a `projectionRequest` struct which contains the projection and other relevant data to create the projection. 
The return value of the handler is a `map[string]any` which contains the data that will be passed to the template. 
When returning a `nil` map the slide will not be rendered. 
The function signature of a handler needs to be as following: 

```go
func ProjectionTypeSlideHandler(ctx context.Context, req *projectionRequest) (map[string]any, error)
```

### Template

After a handler is executed the provided data will be automatically passed to the template in `templates/slides/<projection_type>.html`.
A custom template can be used by passing a `_template` field in the map returned by the handler.
In that case `templates/slides/<_template>.html` is used. 

The templates are parsed with Go `html/template` library.

### (optional) Add stylesheets and scripts

Stylesheets and JavaScript can be added in `web/src/slide/`.

When adding stylesheets the filename should contain the projection type. 
The created stylesheet should be added in the projection template at the top as following:

```html
<link rel="stylesheet" type="text/css" href="/system/projector/static/slide/<projection_type>.css" />
```

The preferred way of adding JavaScript to a slide is by creating a custom html element. 
For this web components should be used. 
A guide for this can be found [here](https://developer.mozilla.org/de/docs/Web/API/Web_components/Using_custom_elements).
The web component needs to be registered in `web/src/projector.js`.

After changing stylesheets or JavaScript they need to be bundled.
The make targets `build-web-assets` and `build-watch-web-assets` can be used for this.
