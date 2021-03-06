package validator

import (
	"time"

	"myecho/handler/errors"
	"myecho/handler/rtype"
)

func ValidateArticleRequest(articleRequest *rtype.ArticleRequest) error {
	if len(articleRequest.Title) == 0 {
		return errors.ErrTitleEmpty
	}
	if len(articleRequest.Content) == 0 {
		return errors.ErrContentEmpty
	}
	if articleRequest.PostTime.IsZero() {
		articleRequest.PostTime = time.Now()
	}
	if err := ValidateCategoryID(articleRequest.CategoryID); err != nil {
		return err
	}
	if err := ValidateTagIDs(articleRequest.TagIDs); err != nil {
		return err
	}
	return nil
}
