package models

type Master struct {
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name"`
}

type Category struct {
	Id           int64  `json:"id,omitempty"`
	MasterId     int64  `json:"master_id"`
	NameCategory string `json:"name_category"`
	Bahan        Bahan  `json:"bahan"`
}

type Bahan struct {
	Id         int64  `json:"id,omitempty"`
	MasterId   int64  `json:"master_id"`
	CategoryId int64  `json:"category_id"`
	NameBahan  string `json:"name_bahan"`
}

type Request struct {
	Name         string   `json:"name"`
	NameCategory string   `json:"name_category"`
	NameBahan    []string `json:"name_bahan"`
}

type Response struct {
	ID      int64       `json:"id,omitempty"`
	Status  int64       `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type ResponseResult struct {
	Name         string   `json:"name"`
	NameCategory string   `json:"name_category"`
	NameBahan    []string `json:"name_bahan"`
}
