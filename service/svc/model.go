package svc

type AdminKnowledge struct {
	Reliability  float64           `form:"reliability" json:"reliability,omitempty" yaml:"reliability"`
	Province     *DefaultKnowledge `form:"province" json:"province,omitempty" yaml:"province"`
	City         *DefaultKnowledge `form:"city" json:"city,omitempty" yaml:"city"`
	District     *DefaultKnowledge `form:"district" json:"district,omitempty" yaml:"district"`
	Township     *DefaultKnowledge `form:"township" json:"township,omitempty" yaml:"township"`
	Neighborhood *DefaultKnowledge `form:"neighborhood" json:"neighborhood,omitempty" yaml:"neighborhood"`
}

type Basics struct {
	Name          string               `form:"name" json:"name,omitempty" yaml:"name"`
	Id            string               `form:"id" json:"id,omitempty" yaml:"id"`
	Category      string               `form:"category" json:"category,omitempty" yaml:"category"`
	Reliabilities DefaultKnowledgeList `form:"reliabilities" json:"Reliabilities,omitempty" yaml:"reliabilities"`
}

type DefaultKnowledge struct {
	Name         string   `form:"name" json:"name,omitempty" yaml:"name"`
	Id           string   `form:"id" json:"id,omitempty" yaml:"id"`
	Labels       []string `form:"labels" json:"labels,omitempty" yaml:"labels"`
	Category     string   `form:"category" json:"category,omitempty" yaml:"category"`
	Relationship string   `form:"relationship" json:"relationship,omitempty" yaml:"relationship"`
	Reliability  float64  `form:"reliability" json:"reliability,omitempty" yaml:"reliability"`
}

type DefaultKnowledgeList []*DefaultKnowledge

func (s DefaultKnowledgeList) Len() int {
	return len(s)
}

func (s DefaultKnowledgeList) Less(i, j int) bool {
	if s[i].Reliability > s[j].Reliability {
		return true
	} else {
		return false
	}
}

func (s DefaultKnowledgeList) Swap(i, j int) {
	temp := s[i]
	s[i] = s[j]
	s[j] = temp
}
