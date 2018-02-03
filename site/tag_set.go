package site

// TagSet is a set of tags
type TagSet map[string]struct{}

// Add adds tag to set if it not in there
func (ts TagSet) Add(tags []string) {
	for _, tag := range tags {
		ts[tag] = struct{}{}
	}
}

// All returns slice of all tags
func (ts TagSet) All() []string {
	var tags []string
	for tag := range ts {
		tags = append(tags, tag)
	}
	return tags
}
