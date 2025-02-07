package models

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
