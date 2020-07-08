package formula

type Tag struct {
	Name   string `json:"name"`
	ZipUrl string `json:"zipball_url"`
}

type Tags map[string]string

type Repo struct {
	Name     string `json:"name"`
	ZipUrl   string `json:"zipUrl"`
	Version  string `json:"version"`
	Token    string `json:"token,omitempty"`
	Priority int    `json:"priority"`
}

type Repos []Repo

func (r Repos) Len() int {
	return len(r)
}

func (r Repos) Less(i, j int) bool {
	return r[i].Priority < r[j].Priority
}

func (r Repos) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

type RepositoryAdder interface {
	Add(d Repo) error
}

type RepositoryLister interface {
	List() (Repos, error)
}

type RepositoryUpdater interface {
	Update() error
}

type RepositoryDeleter interface {
	Delete(name string) error
}

type RepositoryCleaner interface {
	Clean(name string) error
}

type RepositoryLoader interface {
	Load() error
}

type RepositoryAddLister interface {
	RepositoryAdder
	RepositoryLister
}

type RepositoryDelLister interface {
	RepositoryDeleter
	RepositoryLister
}
