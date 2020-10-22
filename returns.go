package baapi

import (
	"errors"
	"strconv"
	"strings"
)

func populateBotInfoReturnStruct(response getBotResponseStruct) (BotInfo, error) {

	var ret BotInfo
	var result = response.Result

	if response.Message != "" {
		return ret, errors.New(response.Message)
	}
	if result.Msg != "http://t.me/BotsArchive/0" {
		ret.ChannelLink = response.Result.Msg
	}

	if result.DeveloperID != "0" {
		devid, err := strconv.Atoi(result.DeveloperID)
		if err != nil {
			return ret, errors.New(UnexpectedError)
		}
		ret.DeveloperID = devid
	}

	ret.Infos.Languages = strings.Split(result.Languages, " ")
	ret.Infos.Tags = strings.Split(strings.ReplaceAll(result.Tags, "#", ""), " ")

	ret.Infos.Name, ret.Infos.Username, ret.Infos.Description = result.Name, result.Username, result.Description
	ret.Infos.HasPhoto, ret.Rating.TotalStars, ret.Rating.VotesCount = result.Photo, result.Stars, result.Votes
	ret.Rating.AverageRating, ret.ID = result.Vote, result.ID

	if result.Groups == 1 {
		ret.Infos.SupportsGroups = true
	}
	if result.Inline == 1 {
		ret.Infos.SupportsInline = true
	}
	if result.Warn != "" {
		ret.Infos.HasCopyrightAlert = true
	}

	return ret, nil

}

func populateUserRatingReturnValues(response getUserVoteResponseStruct) (bool, int, error) {

	var hasvoted bool

	if response.Message != "" {
		return false, 0, errors.New(response.Message)
	}

	if response.Result == "" {
		response.Result = "0"
	}

	vote, err := strconv.Atoi(response.Result)
	if err != nil {
		return false, 0, errors.New(UnexpectedError)
	}

	if vote != 0 {
		hasvoted = true
	}

	return hasvoted, vote, nil

}
