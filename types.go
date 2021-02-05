package vk_api_sdk

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
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Sex       Sex    `json:"sex"`
	BDate     string `json:"bdate"`
	//
	Status   string   `json:"status"`
	Relation Relation `json:"relation"`
}
