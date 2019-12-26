package handlers

type link struct {
	Url string `json:"url"`
}

type links []link

func getLink(url string) link {
	return link{Url: url}
}

func getLinks(urls ...string) links {
	result := make(links, len(urls))
	for i, url := range urls {
		result[i] = getLink(url)
	}
	return result
}

type LinksMessaging struct {
	Links links `json:"links"`
}

func (m *LinksMessaging) getMessage() interface{} {
	return *m
}

func NewLinksMessaging(links ...string) *LinksMessaging {
	return &LinksMessaging{
		Links: getLinks(links...),
	}
}
