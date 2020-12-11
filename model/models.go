package model

import "fmt"

type FizzbuzzParam struct {
	FstMultiple int    `form:"fstM" binding:"required,gt=0"`
	SecMultiple int    `form:"secM" binding:"required,gt=0"`
	Limit       int    `form:"limit" binding:"required,gt=0"`
	Label1      string `form:"label1" binding:"required,min=1"`
	Label2      string `form:"label2" binding:"required,min=1"`
}

type StatResponse struct {
	Hits            int    `json:"hits"`
	AssociateParams string `json:"associateParams"`
}

type FizzBuzzRresponse struct {
	Sequences []string `json:"sequence"`
}

func (f FizzbuzzParam) String() string {
	return fmt.Sprintf("%d, %d, %d, %s, %s", f.FstMultiple, f.SecMultiple, f.Limit, f.Label1, f.Label2)
}
