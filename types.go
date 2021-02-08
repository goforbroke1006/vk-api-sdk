package vksdk

import (
	"errors"
	"strconv"
	"strings"
)

type Sex uint8

const (
	SexFemale = Sex(1)
	SexMale   = Sex(2)
)

type Country uint8

const (
	Russia  = Country(1)
	Ukraine = Country(2)
)

type Relation uint8

const (
	NoMarried      = Relation(1)
	HasFriend      = Relation(2)
	Engaged        = Relation(3)
	Married        = Relation(4)
	HardUnderstand = Relation(5)
	ActiveSearch   = Relation(6)
	FallInLove     = Relation(7)
	CivilMarriage  = Relation(8)
)

type NameCase string

const (
	Nom = NameCase("nom")
	Gen = NameCase("gen")
	Dat = NameCase("dat")
	Acc = NameCase("acc")
	Ins = NameCase("ins")
	Abl = NameCase("abl")
)

var allowedProfileFields = []string{
	"photo_id",
	"verified",
	"sex",
	"bdate",
	"city",
	"countryhome_town",
	"has_photo",
	"photo_50",
	"photo_100",
	"photo_200_orig",
	"photo_200",
	"photo_400_orig",
	"photo_max",
	"photo_max_orig",
	"online",
	"domain",
	"has_mobile",
	"contacts",
	"site",
	"education",
	"universities",
	"schools",
	"status",
	"last_seen",
	"followers_count",
	"common_count",
	"occupation",
	"nickname",
	"relatives",
	"relation",
	"personal",
	"connections",
	"exports",
	"activities",
	"interests",
	"music",
	"movies",
	"tv",
	"books",
	"games",
	"about",
	"quotes",
	"can_post",
	"can_see_all_posts",
	"can_see_audio",
	"can_write_private_message",
	"can_send_friend_request",
	"is_favorite",
	"is_hidden_from_feed",
	"timezone",
	"screen_name",
	"maiden_name",
	"crop_photo",
	"is_friend",
	"friend_status",
	"career",
	"military",
	"blacklisted",
	"blacklisted_by_me",
	"can_be_invited_group",
}

type UsersGetResponse struct {
	Response []User `json:"response"`
}

type User struct {
	ID        int64     `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Sex       Sex       `json:"sex"`
	BDate     BirthDate `json:"bdate"`
	//
	Status   string   `json:"status"`
	Relation Relation `json:"relation"`
}

type BirthDate string

func (bd BirthDate) Day() (int64, error) {
	parts := strings.Split(string(bd), ".")
	day, err := strconv.ParseInt(parts[0], 10, 64)
	return day, err
}
func (bd BirthDate) Month() (int64, error) {
	parts := strings.Split(string(bd), ".")
	if len(parts) < 2 {
		return 0, errors.New("month is not specified")
	}
	day, err := strconv.ParseInt(parts[1], 10, 64)
	return day, err
}
func (bd BirthDate) Year() (int64, error) {
	parts := strings.Split(string(bd), ".")
	if len(parts) < 3 {
		return 0, errors.New("year is not specified")
	}
	year, err := strconv.ParseInt(parts[2], 10, 64)
	return year, err
}

type ResponseError struct {
	Code          int    `json:"error_code"`
	Message       string `json:"error_msg"`
	RequestParams []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"request_params"`
}

type PhotosGetAlbumsResponse struct {
	Response struct {
		Count uint64  `json:"count"`
		Items []Album `json:"items"`
	}
	Error *ResponseError `json:"error"`
}

type Album struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

const (
	PhotosSystemAlbumOnThePage   = -6
	PhotosSystemAlbumOnTheWall   = -7
	PhotosSystemAlbumSavedPhotos = -15
)

type PhotosGetResponse struct {
	Response struct {
		Count uint64  `json:"count"`
		Items []Photo `json:"items"`
	}
	Error *ResponseError `json:"error"`
}

type Photo struct {
	ID      int64  `json:"id"`
	AlbumID int64  `json:"album_id"`
	Title   string `json:"title"`
	Text    string `json:"text"`
	Date    int64  `json:"date"`
	Sizes   []struct {
		Width  int    `json:"width"`
		Height int    `json:"height"`
		Url    string `json:"url"`
		Type   string `json:"type"`
	} `json:"sizes"`
}

type VideoGetResponse struct {
	Response struct {
		Count uint64  `json:"count"`
		Items []Video `json:"items"`
	}
	Error *ResponseError `json:"error"`
}

type Video struct {
	ID          int64          `json:"id"`
	Title       string         `json:"title"`
	OwnerID     int64          `json:"owner_id"`
	Width       int            `json:"width"`
	Height      int            `json:"height"`
	Comments    uint64         `json:"comments"`
	AddingDate  int64          `json:"adding_date"`
	Date        int64          `json:"date"`
	Description string         `json:"description"`
	Duration    uint64         `json:"duration"`
	Views       uint64         `json:"views"`
	Image       []VideoPreview `json:"image"`
	Player      string         `json:"player"`
	Files       struct {
		MP4Quality240 string `json:"mp4_240"`
		MP4Quality360 string `json:"mp4_360"`
		MP4Quality480 string `json:"mp4_480"`
		MP4Quality720 string `json:"mp4_720"`
	} `json:"files"`
	Likes struct {
		Count     uint64 `json:"count"`
		UserLikes uint64 `json:"user_likes"`
	} `json:"likes"`
	Reposts struct {
		Count        uint64 `json:"count"`
		UserReposted uint64 `json:"user_reposted"`
	} `json:"reposts"`
}

type VideoPreview struct {
	Width       int    `json:"width"`
	Height      int    `json:"height"`
	Url         string `json:"url"`
	WithPadding int    `json:"with_padding"`
}

type VideoFileQuality string

const (
	VideoFileQuality240 = VideoFileQuality("mp4_240")
	VideoFileQuality360 = VideoFileQuality("mp4_360")
	VideoFileQuality480 = VideoFileQuality("mp4_480")
	VideoFileQuality720 = VideoFileQuality("mp4_720")
)

type VideoGetAlbumsResponse struct {
	Response struct {
		Count uint64       `json:"count"`
		Items []VideoAlbum `json:"items"`
	}
	Error *ResponseError `json:"error"`
}

type VideoAlbum struct {
	ID          int64  `json:"id"`
	OwnerID     int64  `json:"owner_id"`
	Title       string `json:"title"`
	Count       uint   `json:"count"`
	Photo320    string `json:"photo_320"`
	Photo160    string `json:"photo_160"`
	UpdatedTime int64  `json:"updated_time"`
}

const (
	VideoSystemAlbumUploaded = -1
	VideoSystemAlbumAdded    = -2
)
