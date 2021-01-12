package baapi

import (
	"errors"
	"net/url"
	"strconv"

	jsoniter "github.com/json-iterator/go"
	"github.com/valyala/fasthttp"
)

// The constants below are set to simplify error handling.
// Please consult the appropriate page on the wiki to learn more about this topic.
const (
	FailedRequest            = "request failed"
	InvalidResponse          = "invalid response. Response is not JSON-decodable."
	UnexpectedError          = "unexpected error"
	MissingUsernameParameter = "username GET parameter required"
	MissingParameters        = "bot_id and user_id GET parameters required"
	BotNotFound              = "bot not found"
)

// GetBotInfo returns a struct containing informations about a bot present in BotsArchive.
// It might return one of the following errors: FailedRequest, InvalidResponse, UnexpectedError, MissingUsernameParameter, BotNotFound
func GetBotInfo(username string) (BotInfo, error) {

	apiResponse, err := callAPI("getBotID.php?username=" + url.QueryEscape(username))

	var response getBotResponseStruct
	err = jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(apiResponse, &response)

	if err != nil {
		return BotInfo{}, errors.New(InvalidResponse)
	}

	return populateBotInfoReturnStruct(response)

}

// GetUserRatingByBotID returns a boolean telling if the user has voted and an integer expressing the number of stars selected by the user.
// It might return one of the following errors: FailedRequest, InvalidResponse, UnexpectedError, MissingParameters, BotNotFound
func GetUserRatingByBotID(botid, userid int) (bool, int, error) {

	apiResponse, err := callAPI("getUserVote.php?user_id=" + strconv.Itoa(userid) + "&bot_id=" + strconv.Itoa(botid))

	var response getUserVoteResponseStruct
	err = jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(apiResponse, &response)

	if err != nil {
		return false, 0, errors.New(InvalidResponse)
	}

	return populateUserRatingReturnValues(response)

}

// GetUserRatingByBotUsername returns a boolean telling if the user has voted and an integer expressing the number of stars selected by the user.
// Please note: this method sends two API calls instead of one as GetUserRatingByBotID. If you have to get ratings for one bot only, it's better if you save the ID of the bot somewhere and you use GetUserRatingByBotID.
// It might return one of the following errors: FailedRequest, InvalidResponse, UnexpectedError, MissingParameters, BotNotFound
func GetUserRatingByBotUsername(username string, userid int) (bool, int, error) {

	botInfo, err := GetBotInfo(username)
	if err != nil {
		return false, 0, err
	}

	return GetUserRatingByBotID(botInfo.ID, userid)

}

func callAPI(endpoint string) ([]byte, error) {

	var body []byte
	statusCode, body, err := fasthttp.Get(body, "https://api.botsarchive.com/"+endpoint)

	if statusCode != fasthttp.StatusOK {
		return body, errors.New(FailedRequest)
	}

	return body, err

}
