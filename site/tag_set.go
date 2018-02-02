package site

type TagSet map[string]struct{}

func (ts TagSet) Add(tags []string) {
	for _, tag := range tags {
		ts[tag] = struct{}{}
	}
}

func (ts TagSet) All() []string {
	var tags []string
	for tag := range ts {
		tags = append(tags, tag)
	}
	return tags
}
