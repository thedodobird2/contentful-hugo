package contentful

const (
	spaceUrl = "https://cdn.contentful.com/spaces/"
)

type Space struct {
	Id, Token string
}

func (s *Space) Get() []byte {
	return contentfulRequest(spaceUrl+s.Id, s.Token)
}

func (s *Space) GetContentModel() []byte {
	return contentfulRequest(spaceUrl+s.Id+"/content_types", s.Token)
}

func (s *Space) GetContentType(id string) []byte {
	return contentfulRequest(spaceUrl+s.Id+"/content_types/"+id, s.Token)
}

func (s *Space) GetEntries() []byte {
	return contentfulRequest(spaceUrl+s.Id+"/entries", s.Token)
}

func (s *Space) GetEntry(id string) []byte {
	return contentfulRequest(spaceUrl+s.Id+"/entries/"+id, s.Token)
}

func (s *Space) GetAssets() []byte {
	return contentfulRequest(spaceUrl+s.Id+"/assets", s.Token)
}

func (s *Space) GetAsset(id string) []byte {
	return contentfulRequest(spaceUrl+s.Id+"/assets/"+id, s.Token)
}
