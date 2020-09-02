package model

type Group struct {
	Name        string        `yaml:"name"`
	Description string        `yaml:"description"`
	Chairs      string        `yaml:"chairs"`
	Meetings    []Meeting     `yaml:"meetings"`
	Members     []GroupMember `yaml:"members"`
	Website     string        `yaml:"website"`
	Email       string        `yaml:"email"`
}

type Meeting struct {
	Topic           string `yaml:"topic"`
	MeetingTime     string `yaml:"meetingTime"`
	MeetingLocation string `yaml:"meetingLocation"`
}

type GroupMember struct {
	Username    string `yaml:"username"`
	Role        string `yaml:"role"`
	DisplayName string `yaml:"displayName"`
	Email       string `yaml:"email"`
}

const (
	GroupCommittees = "committees"
	GroupSIGs       = "sigs"

	GroupTop4      = "Top4"
	GroupAdmin     = "Admin"
	GroupCorporate = "Corporate"
)

var (
	GroupValidCommittees = []string{GroupTop4, GroupAdmin, GroupCorporate}
)
