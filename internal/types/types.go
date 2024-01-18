package types

type Org struct {
	Login             string `json:"login,omitempty"`
	Id                int    `json:"id,omitempty"`
	NodeId            string `json:"node_id,omitempty"`
	Url               string `json:"url,omitempty"`
	ReposUrl          string `json:"repos_url,omitempty"`
	EventsUrl         string `json:"events_url,omitempty"`
	HooksUrl          string `json:"hooks_url,omitempty"`
	IssuesUrl         string `json:"issues_url,omitempty"`
	MembersUrl        string `json:"members_url,omitempty"`
	PublicMembers_url string `json:"public_members_url,omitempty"`
	AvatarUrl         string `json:"avatar_url,omitempty"`
	Description       string `json:"description,omitempty"`
}

type Repo struct {
    Name string `json:"name,required"`
    FullName string `json:"full_name,required"`
    Url string `json:"html_url,required"`
    IconPath string
}

func (r *Repo) String() string {
    return r.FullName
}

type AlfredItem struct {
    Title string `json:"title"`
    Subtitle string `json:"subtitle"`
    Arg string `json:"arg"`
    Autocomplete string `json:"autocomplete"`
    Icon AlfredIcon `json:"icon"`
}

type AlfredIcon struct {
    IconType string `json:"type"`
    Path string `json:"path"`
}
