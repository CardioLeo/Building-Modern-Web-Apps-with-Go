package models

// this will never import other packages; it will only exist to be imported
// by packages other than itself.

// TemplateData holds data sent from handlers to templates
type TemplateData struct {
        StringMap map[string]string
        IntMap map[string]int
        FloatMap map[string]float32
        Data map[string]interface{} //special data
                                        // "and in Go, when you're not sure
                                        // what the data type is, we call it
                                        // an interface."
        CSRFtoken string
        Flash string
        Warning string
        Error string
}
