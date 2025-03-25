package models

// TemplateData moved from handlers package in video 37

// "This will never import another package; but will only exist to be imported by other packages" -- to avoid import cycle problem

// new type for any kind of data I might want to pass, video 37
// TemplateData holds data from handlers to templates
type TemplateData struct {
        //
        StringMap map[string]string
        IntMap map[string]int
        FloatMap map[string]float32
        // special type for anything
        // "and in Go, when you're not sure what the datatype is, you call it an interface"
        Data map[string]interface{}
        // the following is added just for now, so it doesn't have to be added later
        // Cross-Site-Certificate-Forgery...
        CSRFToken string
        // for success message or whatever:
        Flash string
        Warning string
        Error string
}
