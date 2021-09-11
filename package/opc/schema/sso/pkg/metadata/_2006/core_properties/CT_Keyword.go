package core_properties

type CT_Keyword struct {
	LangAttr *string
	Content  string
}

func NewCT_Keyword() *CT_Keyword {
	return &CT_Keyword{}
}
