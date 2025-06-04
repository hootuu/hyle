package hypes

type MediaType = int

const (
	MediaTypePhoto MediaType = 1
	MediaTypeVideo MediaType = 2
)

type Media struct {
	Type MediaType         `json:"type"`
	Link string            `json:"link"`
	Meta map[string]string `json:"meta,omitempty"`
}

func NewMedia(t MediaType, link string) *Media {
	return &Media{
		Type: t,
		Link: link,
		Meta: nil,
	}
}

func (m *Media) SetMeta(key string, val string) *Media {
	if m.Meta == nil {
		m.Meta = make(map[string]string)
	}
	m.Meta[key] = val
	return m
}

type MediaDict map[string][]*Media

func NewMediaDict() MediaDict {
	return make(MediaDict)
}

func (md MediaDict) PutAppend(key string, m *Media) MediaDict {
	arr, ok := md[key]
	if !ok {
		md[key] = []*Media{m}
		return md
	}
	arr = append(arr, m)
	return md
}
