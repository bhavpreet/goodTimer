// Code generated by goctl. DO NOT EDIT.
package types

type AddRoundRequest struct {
	Name string `json:"name"`
}

type Round struct {
	ID        string `json:"id" storm:"id,unique"`
	Name      string `json:"name" storm:"unique"`
	Status    string `json:"status,options=RUNNING|FINISHED"`
	ISCurrent bool   `json:"is_current"`
}

type Bib struct {
	ID             string   `json:"id" storm:"id,unique"`
	No             string   `json:"no" storm:"unique"`
	ParentRoundID  string   `json:"-" storm:"index"`
	Round          Round    `json:"round"`
	Status         string   `json:"status,options=RUNNING|FINISHED"`
	ISCurrentStart bool     `json:"is_current_start"`
	ISCurrentEnd   bool     `json:"is_current_end"`
	StartTimes     []string `json:"start_times"`
	StartTime      string   `json:"start_time"`
	ST             string   `json:"st"`
	EndTimes       []string `json:"end_times"`
	EndTime        string   `json:"end_time"`
	Duration       string   `json:"duration"`
}

type Current struct {
	ID              string `json:"id"`
	CurrentRound    string `json:"current_round"`
	CurrentStartBib string `json:"current_start_bib"`
	CurrentEndBib   string `json:"current_end_bib"`
}

type ListRoundsResp struct {
	Rounds []Round `json:"rounds"`
}

type GetRoundRequest struct {
	ID string `path:"id"`
}

type UpdateRoundRequest struct {
	ID   string `path:"round"`
	Name string `json:"name"`
}

type UpdateBibRequest struct {
	Round     string `path:"round"`
	ID        string `path:"id"`
	BibNo     string `json:"bib_no"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

type DeleteRoundRequest struct {
	Round string `path:"round"`
}

type AddBibRequest struct {
	Round string `path:"round"`
	Bib   string `json:"bib"`
}

type GetBibRequest struct {
	Round string `path:"round"`
	Bib   string `path:"bib"`
}

type DeleteBibRequest struct {
	Round string `path:"round"`
	ID    string `path:"id"`
}

type SetRoundCurrentRequest struct {
	Round string `path:"round"`
}

type SetBibStatusRequest struct {
	Round  string `path:"round"`
	Status string `path:"status" json:"status,options=RUNNING|FINISHED"`
}

type SetBibCurrentReq struct {
	Round string `path:"round"`
	Bib   string `path:"bib"`
}

type ListBibsReq struct {
	Round string `path:"round"`
}

type ListBibsResp struct {
	Bibs []Bib `json:"bibs"`
}

type GetCurrentRoundResp struct {
	CurrentRound string `json:"current_round"`
}
