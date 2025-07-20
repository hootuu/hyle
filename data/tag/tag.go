package tag

type Tag []string

func NewTag(values ...string) Tag {
	if len(values) > 0 {
		return Tag(values)
	}
	return make(Tag, 0)
}

func (tag *Tag) Exists(t string) bool {
	for _, v := range *tag {
		if v == t {
			return true
		}
	}
	return false
}

func (tag *Tag) Append(t ...string) *Tag {
	for _, v := range t {
		tag.doAppend(v)
	}
	return tag
}

func (tag *Tag) doAppend(t string) *Tag {
	if tag.Exists(t) {
		return tag
	}

	*tag = append(*tag, t)
	return tag
}

func (tag *Tag) Remove(t string) {
	for i, v := range *tag {
		if v == t {
			*tag = append((*tag)[:i], (*tag)[i+1:]...)
			return
		}
	}
}
