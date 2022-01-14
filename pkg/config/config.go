package config

type Release struct {
	URL         string  `json:"url"`
	ID          int64   `json:"id"`
	TagName     string  `json:"tag_name"`
	Name        string  `json:"name"`
	PreRelease  bool    `json:"prerelease"`
	CreatedAt   string  `json:"created_at"`
	PublishedAt string  `json:"published_at"`
	Assets      []Asset `json:"assets"`
}

type Asset struct {
	URL           string `json:"url"`
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	ContentType   string `json:"content_type"`
	Size          int64  `json:"size"`
	DownloadCount int64  `json:"download_count"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}

type Repository struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	FullName    string `json:"full_name"`
	Description string `json:"description"`
	Stars       int64  `json:"stargazers_count"`
	Forks       int64  `json:"forks"`
	Subscribers int64  `json:"subscribers_count"`
	Language    string `json:"language"`
}
