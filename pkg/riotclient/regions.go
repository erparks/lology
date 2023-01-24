package riotclient

type Region string

const (
	BR1  Region = "br1"
	EUN1 Region = "eun1"
	EUW1 Region = "euw1"
	JP1  Region = "jp1"
	KR   Region = "kr"
	LA1  Region = "la1"
	LA2  Region = "la2"
	NA1  Region = "na1"
	OC1  Region = "oc1"
	TR1  Region = "tr1"
	RU   Region = "ru"
	PH2  Region = "ph1"
	SG2  Region = "sg1"
	TH2  Region = "th1"
	TW2  Region = "tw1"
	VN2  Region = "vn1"
)

func (r Region) String() string {
	return string(r)
}

func (r Region) Routing() string {
	return map[Region]string{
		BR1:  "americas",
		EUN1: "europe",
		EUW1: "europe",
		JP1:  "asia",
		KR:   "asia",
		LA1:  "americas",
		LA2:  "americas",
		NA1:  "americas",
		OC1:  "sea",
		TR1:  "europe",
		RU:   "europe",
		PH2:  "sea",
		SG2:  "sea",
		TH2:  "sea",
		TW2:  "sea",
		VN2:  "sea",
	}[r]
}
