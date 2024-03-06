package models

type DirectoryEntry struct {
	Name     string `json:"name"`
	FullPath string `json:"full_path"`
	IsDir    bool   `json:"is_directory"`
	Size     int64  `json:"size"`
	Mode     string `json:"mode"`
	ModTime  string `json:"modified_time"`
}

type SuccessResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type ErrorResponse struct {
	Status string `json:"status"`
	Error  Error  `json:"error"`
}

type Error struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Details   string `json:"details,omitempty"`
	Timestamp string `json:"timestamp"`
}
