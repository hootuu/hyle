package media

type Type = int

const (
	ImageType Type = 1
	VideoType Type = 2
)

type Media struct {
	Type Type              `json:"type"`
	Link string            `json:"link"`
	Meta map[string]string `json:"meta,omitempty"`
}

type More []*Media

type Dict map[string]More

func New(t Type, link string) *Media {
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
