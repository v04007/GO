package main

type Page []struct {
	MenuID    int    `json:"menu_id"`
	Title     string `json:"title"`
	Path      string `json:"path"`
	Key       string `json:"key"`
	ParentKey string `json:"parentKey"`
	Icon      string `json:"icon"`
	KeepAlive string `json:"keepAlive"`
	Order     int    `json:"order"`
	Children  []struct {
		MenuID    int    `json:"menu_id"`
		Title     string `json:"title"`
		Path      string `json:"path"`
		Key       string `json:"key"`
		ParentKey string `json:"parentKey"`
		Icon      string `json:"icon"`
		KeepAlive string `json:"keepAlive"`
		Order     int    `json:"order"`
	} `json:"children,omitempty"`
}


"data": Page{{9, "列表页", "/list", "list", "", "icon_list", "false", 1,
[]struct {
MenuID    int    `json:"menu_id"`
Title     string `json:"title"`
Path      string `json:"path"`
Key       string `json:"key"`
ParentKey string `json:"parentKey"`
Icon      string `json:"icon"`
KeepAlive string `json:"keepAlive"`
Order     int    `json:"order"`
}{{10, "卡片列表", "/card", "listCard", "list", "", "false", 5485},
{11, "查询列表", "/search", "listSearch", "list", "", "false", 9588},
{7, "表单页", "/form", "from", "", "icon_form", "false", 3},
}},