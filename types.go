package baapi

type getBotResponseStruct struct {
	Ok      int    `json:"ok"`
	Message string `json:"message"`
	ID      int    `json:"id"`
	Result  struct {
		ID          int         `json:"id"`
		Name        string      `json:"name"`
		Username    string      `json:"username"`
		Description string      `json:"description"`
		Warn        string      `json:"warn"`
		Msg         string      `json:"msg"`
		Category    interface{} `json:"category"`
		Groups      int         `json:"groups"`
		Inline      int         `json:"inline"`
		DeveloperID interface{} `json:"developer_id"`
		Stars       int         `json:"stars"`
		Votes       int         `json:"votes"`
		Vote        float64     `json:"vote"`
		Tags        string      `json:"tags"`
		Languages   string      `json:"languages"`
		Photo       bool        `json:"photo"`
	} `json:"result"`
}

type getUserVoteResponseStruct struct {
	Ok      int    `json:"ok"`
	Result  string `json:"result"`
	Vote    bool   `json:"vote"`
	Message string `json:"message"`
}

// BotInfo contains bot informations.
// If field ChannelLink is == "", the bot hasn't been published yet.
// If DeveloperID == 0, the bot wasn't claimed by its developer.
type BotInfo struct {
	Rating struct {
		AverageRating          float64
		TotalStars, VotesCount int
	}
	Infos struct {
		Name, Username, Description                                 string
		Languages, Tags                                             []string
		HasPhoto, SupportsGroups, SupportsInline, HasCopyrightAlert bool
	}
	ID, DeveloperID int
	ChannelLink     string
}
