package cmhtypes

type Source struct {
	Id              string   `json:"id" binding:"required"`
	Name            string   `json:"name" binding:"required"`
	AllowedMachines []string `json:"allowedMachines omitempty"`
	SourceTypeName  string   `json:"sourceTypeName" binding:"required"`
}

type Statistics map[string]string

type SourceType struct {
	Name  string     `json:"name" binding:"required"`
	Stats Statistics `json:"stats"`
}

type Fetcher struct {
	Id       string `json:"id" binding:"required"`
	ListName string `json:"listName" binding:"required"`
}
