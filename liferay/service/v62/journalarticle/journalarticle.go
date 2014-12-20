// Copyright (c) 2014 Andrea Di Giorgi. All rights reserved.
//
// This library is free software; you can redistribute it and/or modify it under
// the terms of the GNU Lesser General Public License as published by the Free
// Software Foundation; either version 2.1 of the License, or (at your option)
// any later version.
//
// This library is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS
// FOR A PARTICULAR PURPOSE. See the GNU Lesser General Public License for more
// details.

package journalarticle

import (
	"github.com/ithildir/liferay-sdk-go/liferay"
)

type Service struct {
	session liferay.Session
}

func NewService(s liferay.Session) *Service {
	return &Service{s}
}

func (s *Service) AddArticle(groupId int64, folderId int64, classNameId int64, classPK int64, articleId string, autoArticleId bool, titleMap map[string]interface{}, descriptionMap map[string]interface{}, content string, type string, ddmStructureKey string, ddmTemplateKey string, layoutUuid string, displayDateMonth int, displayDateDay int, displayDateYear int, displayDateHour int, displayDateMinute int, expirationDateMonth int, expirationDateDay int, expirationDateYear int, expirationDateHour int, expirationDateMinute int, neverExpire bool, reviewDateMonth int, reviewDateDay int, reviewDateYear int, reviewDateHour int, reviewDateMinute int, neverReview bool, indexable bool, articleURL string, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["folderId"] = folderId
	_params["classNameId"] = classNameId
	_params["classPK"] = classPK
	_params["articleId"] = articleId
	_params["autoArticleId"] = autoArticleId
	_params["titleMap"] = titleMap
	_params["descriptionMap"] = descriptionMap
	_params["content"] = content
	_params["type"] = type
	_params["ddmStructureKey"] = ddmStructureKey
	_params["ddmTemplateKey"] = ddmTemplateKey
	_params["layoutUuid"] = layoutUuid
	_params["displayDateMonth"] = displayDateMonth
	_params["displayDateDay"] = displayDateDay
	_params["displayDateYear"] = displayDateYear
	_params["displayDateHour"] = displayDateHour
	_params["displayDateMinute"] = displayDateMinute
	_params["expirationDateMonth"] = expirationDateMonth
	_params["expirationDateDay"] = expirationDateDay
	_params["expirationDateYear"] = expirationDateYear
	_params["expirationDateHour"] = expirationDateHour
	_params["expirationDateMinute"] = expirationDateMinute
	_params["neverExpire"] = neverExpire
	_params["reviewDateMonth"] = reviewDateMonth
	_params["reviewDateDay"] = reviewDateDay
	_params["reviewDateYear"] = reviewDateYear
	_params["reviewDateHour"] = reviewDateHour
	_params["reviewDateMinute"] = reviewDateMinute
	_params["neverReview"] = neverReview
	_params["indexable"] = indexable
	_params["articleURL"] = articleURL
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/journalarticle/add-article": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) AddArticle2(groupId int64, folderId int64, classNameId int64, classPK int64, articleId string, autoArticleId bool, titleMap map[string]interface{}, descriptionMap map[string]interface{}, content string, type string, ddmStructureKey string, ddmTemplateKey string, layoutUuid string, displayDateMonth int, displayDateDay int, displayDateYear int, displayDateHour int, displayDateMinute int, expirationDateMonth int, expirationDateDay int, expirationDateYear int, expirationDateHour int, expirationDateMinute int, neverExpire bool, reviewDateMonth int, reviewDateDay int, reviewDateYear int, reviewDateHour int, reviewDateMinute int, neverReview bool, indexable bool, smallImage bool, smallImageURL string, smallFile io.Reader, images map[string]interface{}, articleURL string, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["folderId"] = folderId
	_params["classNameId"] = classNameId
	_params["classPK"] = classPK
	_params["articleId"] = articleId
	_params["autoArticleId"] = autoArticleId
	_params["titleMap"] = titleMap
	_params["descriptionMap"] = descriptionMap
	_params["content"] = content
	_params["type"] = type
	_params["ddmStructureKey"] = ddmStructureKey
	_params["ddmTemplateKey"] = ddmTemplateKey
	_params["layoutUuid"] = layoutUuid
	_params["displayDateMonth"] = displayDateMonth
	_params["displayDateDay"] = displayDateDay
	_params["displayDateYear"] = displayDateYear
	_params["displayDateHour"] = displayDateHour
	_params["displayDateMinute"] = displayDateMinute
	_params["expirationDateMonth"] = expirationDateMonth
	_params["expirationDateDay"] = expirationDateDay
	_params["expirationDateYear"] = expirationDateYear
	_params["expirationDateHour"] = expirationDateHour
	_params["expirationDateMinute"] = expirationDateMinute
	_params["neverExpire"] = neverExpire
	_params["reviewDateMonth"] = reviewDateMonth
	_params["reviewDateDay"] = reviewDateDay
	_params["reviewDateYear"] = reviewDateYear
	_params["reviewDateHour"] = reviewDateHour
	_params["reviewDateMinute"] = reviewDateMinute
	_params["neverReview"] = neverReview
	_params["indexable"] = indexable
	_params["smallImage"] = smallImage
	_params["smallImageURL"] = smallImageURL
	_params["smallFile"] = smallFile
	_params["images"] = images
	_params["articleURL"] = articleURL
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/journalarticle/add-article": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) CopyArticle(groupId int64, oldArticleId string, newArticleId string, autoArticleId bool, version float64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["oldArticleId"] = oldArticleId
	_params["newArticleId"] = newArticleId
	_params["autoArticleId"] = autoArticleId
	_params["version"] = version

	_cmd := map[string]interface{}{
		"/journalarticle/copy-article": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) DeleteArticle(groupId int64, articleId string, articleURL string, serviceContext *liferay.ObjectWrapper) error {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["articleId"] = articleId
	_params["articleURL"] = articleURL
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/journalarticle/delete-article": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) DeleteArticle2(groupId int64, articleId string, version float64, articleURL string, serviceContext *liferay.ObjectWrapper) error {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["articleId"] = articleId
	_params["version"] = version
	_params["articleURL"] = articleURL
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/journalarticle/delete-article": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) ExpireArticle(groupId int64, articleId string, articleURL string, serviceContext *liferay.ObjectWrapper) error {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["articleId"] = articleId
	_params["articleURL"] = articleURL
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/journalarticle/expire-article": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) ExpireArticle2(groupId int64, articleId string, version float64, articleURL string, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["articleId"] = articleId
	_params["version"] = version
	_params["articleURL"] = articleURL
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/journalarticle/expire-article": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetArticle(id int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["id"] = id

	_cmd := map[string]interface{}{
		"/journalarticle/get-article": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetArticle2(groupId int64, articleId string) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["articleId"] = articleId

	_cmd := map[string]interface{}{
		"/journalarticle/get-article": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetArticle3(groupId int64, articleId string, version float64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["articleId"] = articleId
	_params["version"] = version

	_cmd := map[string]interface{}{
		"/journalarticle/get-article": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetArticle4(groupId int64, className string, classPK int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["className"] = className
	_params["classPK"] = classPK

	_cmd := map[string]interface{}{
		"/journalarticle/get-article": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetArticleByUrlTitle(groupId int64, urlTitle string) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["urlTitle"] = urlTitle

	_cmd := map[string]interface{}{
		"/journalarticle/get-article-by-url-title": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetArticleContent(groupId int64, articleId string, languageId string, themeDisplay *liferay.ObjectWrapper) (string, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["articleId"] = articleId
	_params["languageId"] = languageId
	liferay.MangleObjectWrapper(_params, "themeDisplay", "com.liferay.portal.theme.ThemeDisplay", themeDisplay)

	_cmd := map[string]interface{}{
		"/journalarticle/get-article-content": _params,
	}

	var v string

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(string)
	}

	return v, err
}

func (s *Service) GetArticleContent2(groupId int64, articleId string, version float64, languageId string, themeDisplay *liferay.ObjectWrapper) (string, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["articleId"] = articleId
	_params["version"] = version
	_params["languageId"] = languageId
	liferay.MangleObjectWrapper(_params, "themeDisplay", "com.liferay.portal.theme.ThemeDisplay", themeDisplay)

	_cmd := map[string]interface{}{
		"/journalarticle/get-article-content": _params,
	}

	var v string

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(string)
	}

	return v, err
}

func (s *Service) GetArticles(groupId int64, folderId int64) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["folderId"] = folderId

	_cmd := map[string]interface{}{
		"/journalarticle/get-articles": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetArticles2(groupId int64, folderId int64, start int, end int, obc *liferay.ObjectWrapper) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["folderId"] = folderId
	_params["start"] = start
	_params["end"] = end
	liferay.MangleObjectWrapper(_params, "obc", "com.liferay.portal.kernel.util.OrderByComparator", obc)

	_cmd := map[string]interface{}{
		"/journalarticle/get-articles": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetArticlesByArticleId(groupId int64, articleId string, start int, end int, obc *liferay.ObjectWrapper) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["articleId"] = articleId
	_params["start"] = start
	_params["end"] = end
	liferay.MangleObjectWrapper(_params, "obc", "com.liferay.portal.kernel.util.OrderByComparator", obc)

	_cmd := map[string]interface{}{
		"/journalarticle/get-articles-by-article-id": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetArticlesByLayoutUuid(groupId int64, layoutUuid string) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["layoutUuid"] = layoutUuid

	_cmd := map[string]interface{}{
		"/journalarticle/get-articles-by-layout-uuid": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetArticlesByStructureId(groupId int64, ddmStructureKey string, start int, end int, obc *liferay.ObjectWrapper) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["ddmStructureKey"] = ddmStructureKey
	_params["start"] = start
	_params["end"] = end
	liferay.MangleObjectWrapper(_params, "obc", "com.liferay.portal.kernel.util.OrderByComparator", obc)

	_cmd := map[string]interface{}{
		"/journalarticle/get-articles-by-structure-id": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetArticlesByStructureId2(groupId int64, classNameId int64, ddmStructureKey string, status int, start int, end int, obc *liferay.ObjectWrapper) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["classNameId"] = classNameId
	_params["ddmStructureKey"] = ddmStructureKey
	_params["status"] = status
	_params["start"] = start
	_params["end"] = end
	liferay.MangleObjectWrapper(_params, "obc", "com.liferay.portal.kernel.util.OrderByComparator", obc)

	_cmd := map[string]interface{}{
		"/journalarticle/get-articles-by-structure-id": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetArticlesCount(groupId int64, folderId int64) (int, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["folderId"] = folderId

	_cmd := map[string]interface{}{
		"/journalarticle/get-articles-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64))
	}

	return v, err
}

func (s *Service) GetArticlesCount2(groupId int64, folderId int64, status int) (int, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["folderId"] = folderId
	_params["status"] = status

	_cmd := map[string]interface{}{
		"/journalarticle/get-articles-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64))
	}

	return v, err
}

func (s *Service) GetArticlesCountByArticleId(groupId int64, articleId string) (int, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["articleId"] = articleId

	_cmd := map[string]interface{}{
		"/journalarticle/get-articles-count-by-article-id": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64))
	}

	return v, err
}

func (s *Service) GetArticlesCountByStructureId(groupId int64, ddmStructureKey string) (int, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["ddmStructureKey"] = ddmStructureKey

	_cmd := map[string]interface{}{
		"/journalarticle/get-articles-count-by-structure-id": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64))
	}

	return v, err
}

func (s *Service) GetArticlesCountByStructureId2(groupId int64, classNameId int64, ddmStructureKey string, status int) (int, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["classNameId"] = classNameId
	_params["ddmStructureKey"] = ddmStructureKey
	_params["status"] = status

	_cmd := map[string]interface{}{
		"/journalarticle/get-articles-count-by-structure-id": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64))
	}

	return v, err
}

func (s *Service) GetDisplayArticleByUrlTitle(groupId int64, urlTitle string) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["urlTitle"] = urlTitle

	_cmd := map[string]interface{}{
		"/journalarticle/get-display-article-by-url-title": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetFoldersAndArticlesCount(groupId int64, folderIds []interface{}) (int, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["folderIds"] = folderIds

	_cmd := map[string]interface{}{
		"/journalarticle/get-folders-and-articles-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64))
	}

	return v, err
}

func (s *Service) GetGroupArticles(groupId int64, userId int64, rootFolderId int64, start int, end int, orderByComparator *liferay.ObjectWrapper) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["userId"] = userId
	_params["rootFolderId"] = rootFolderId
	_params["start"] = start
	_params["end"] = end
	liferay.MangleObjectWrapper(_params, "orderByComparator", "com.liferay.portal.kernel.util.OrderByComparator", orderByComparator)

	_cmd := map[string]interface{}{
		"/journalarticle/get-group-articles": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetGroupArticles2(groupId int64, userId int64, rootFolderId int64, status int, start int, end int, orderByComparator *liferay.ObjectWrapper) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["userId"] = userId
	_params["rootFolderId"] = rootFolderId
	_params["status"] = status
	_params["start"] = start
	_params["end"] = end
	liferay.MangleObjectWrapper(_params, "orderByComparator", "com.liferay.portal.kernel.util.OrderByComparator", orderByComparator)

	_cmd := map[string]interface{}{
		"/journalarticle/get-group-articles": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) GetGroupArticlesCount(groupId int64, userId int64, rootFolderId int64) (int, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["userId"] = userId
	_params["rootFolderId"] = rootFolderId

	_cmd := map[string]interface{}{
		"/journalarticle/get-group-articles-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64))
	}

	return v, err
}

func (s *Service) GetGroupArticlesCount2(groupId int64, userId int64, rootFolderId int64, status int) (int, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["userId"] = userId
	_params["rootFolderId"] = rootFolderId
	_params["status"] = status

	_cmd := map[string]interface{}{
		"/journalarticle/get-group-articles-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64))
	}

	return v, err
}

func (s *Service) GetLatestArticle(resourcePrimKey int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["resourcePrimKey"] = resourcePrimKey

	_cmd := map[string]interface{}{
		"/journalarticle/get-latest-article": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetLatestArticle2(groupId int64, articleId string, status int) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["articleId"] = articleId
	_params["status"] = status

	_cmd := map[string]interface{}{
		"/journalarticle/get-latest-article": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) GetLatestArticle3(groupId int64, className string, classPK int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["className"] = className
	_params["classPK"] = classPK

	_cmd := map[string]interface{}{
		"/journalarticle/get-latest-article": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) MoveArticle(groupId int64, articleId string, newFolderId int64) error {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["articleId"] = articleId
	_params["newFolderId"] = newFolderId

	_cmd := map[string]interface{}{
		"/journalarticle/move-article": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) MoveArticleFromTrash(groupId int64, articleId string, newFolderId int64, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["articleId"] = articleId
	_params["newFolderId"] = newFolderId
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/journalarticle/move-article-from-trash": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) MoveArticleFromTrash2(groupId int64, resourcePrimKey int64, newFolderId int64, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["resourcePrimKey"] = resourcePrimKey
	_params["newFolderId"] = newFolderId
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/journalarticle/move-article-from-trash": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) MoveArticleToTrash(groupId int64, articleId string) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["articleId"] = articleId

	_cmd := map[string]interface{}{
		"/journalarticle/move-article-to-trash": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) RemoveArticleLocale(companyId int64, languageId string) error {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId
	_params["languageId"] = languageId

	_cmd := map[string]interface{}{
		"/journalarticle/remove-article-locale": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) RemoveArticleLocale2(groupId int64, articleId string, version float64, languageId string) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["articleId"] = articleId
	_params["version"] = version
	_params["languageId"] = languageId

	_cmd := map[string]interface{}{
		"/journalarticle/remove-article-locale": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) RestoreArticleFromTrash(resourcePrimKey int64) error {
	_params := make(map[string]interface{})

	_params["resourcePrimKey"] = resourcePrimKey

	_cmd := map[string]interface{}{
		"/journalarticle/restore-article-from-trash": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) RestoreArticleFromTrash2(groupId int64, articleId string) error {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["articleId"] = articleId

	_cmd := map[string]interface{}{
		"/journalarticle/restore-article-from-trash": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) Search(companyId int64, groupId int64, folderIds []interface{}, classNameId int64, keywords string, version *liferay.ObjectWrapper, type string, ddmStructureKey string, ddmTemplateKey string, displayDateGT int64, displayDateLT int64, status int, reviewDate int64, start int, end int, obc *liferay.ObjectWrapper) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId
	_params["groupId"] = groupId
	_params["folderIds"] = folderIds
	_params["classNameId"] = classNameId
	_params["keywords"] = keywords
	liferay.MangleObjectWrapper(_params, "version", "java.lang.Double", version)
	_params["type"] = type
	_params["ddmStructureKey"] = ddmStructureKey
	_params["ddmTemplateKey"] = ddmTemplateKey
	_params["displayDateGT"] = displayDateGT
	_params["displayDateLT"] = displayDateLT
	_params["status"] = status
	_params["reviewDate"] = reviewDate
	_params["start"] = start
	_params["end"] = end
	liferay.MangleObjectWrapper(_params, "obc", "com.liferay.portal.kernel.util.OrderByComparator", obc)

	_cmd := map[string]interface{}{
		"/journalarticle/search": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) Search2(companyId int64, groupId int64, folderIds []interface{}, classNameId int64, articleId string, version *liferay.ObjectWrapper, title string, description string, content string, type string, ddmStructureKey string, ddmTemplateKey string, displayDateGT int64, displayDateLT int64, status int, reviewDate int64, andOperator bool, start int, end int, obc *liferay.ObjectWrapper) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId
	_params["groupId"] = groupId
	_params["folderIds"] = folderIds
	_params["classNameId"] = classNameId
	_params["articleId"] = articleId
	liferay.MangleObjectWrapper(_params, "version", "java.lang.Double", version)
	_params["title"] = title
	_params["description"] = description
	_params["content"] = content
	_params["type"] = type
	_params["ddmStructureKey"] = ddmStructureKey
	_params["ddmTemplateKey"] = ddmTemplateKey
	_params["displayDateGT"] = displayDateGT
	_params["displayDateLT"] = displayDateLT
	_params["status"] = status
	_params["reviewDate"] = reviewDate
	_params["andOperator"] = andOperator
	_params["start"] = start
	_params["end"] = end
	liferay.MangleObjectWrapper(_params, "obc", "com.liferay.portal.kernel.util.OrderByComparator", obc)

	_cmd := map[string]interface{}{
		"/journalarticle/search": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) Search3(companyId int64, groupId int64, folderIds []interface{}, classNameId int64, articleId string, version *liferay.ObjectWrapper, title string, description string, content string, type string, ddmStructureKeys []interface{}, ddmTemplateKeys []interface{}, displayDateGT int64, displayDateLT int64, status int, reviewDate int64, andOperator bool, start int, end int, obc *liferay.ObjectWrapper) ([]interface{}, error) {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId
	_params["groupId"] = groupId
	_params["folderIds"] = folderIds
	_params["classNameId"] = classNameId
	_params["articleId"] = articleId
	liferay.MangleObjectWrapper(_params, "version", "java.lang.Double", version)
	_params["title"] = title
	_params["description"] = description
	_params["content"] = content
	_params["type"] = type
	_params["ddmStructureKeys"] = ddmStructureKeys
	_params["ddmTemplateKeys"] = ddmTemplateKeys
	_params["displayDateGT"] = displayDateGT
	_params["displayDateLT"] = displayDateLT
	_params["status"] = status
	_params["reviewDate"] = reviewDate
	_params["andOperator"] = andOperator
	_params["start"] = start
	_params["end"] = end
	liferay.MangleObjectWrapper(_params, "obc", "com.liferay.portal.kernel.util.OrderByComparator", obc)

	_cmd := map[string]interface{}{
		"/journalarticle/search": _params,
	}

	var v []interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.([]interface{})
	}

	return v, err
}

func (s *Service) Search4(groupId int64, creatorUserId int64, status int, start int, end int) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["creatorUserId"] = creatorUserId
	_params["status"] = status
	_params["start"] = start
	_params["end"] = end

	_cmd := map[string]interface{}{
		"/journalarticle/search": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) SearchCount(companyId int64, groupId int64, folderIds []interface{}, classNameId int64, keywords string, version *liferay.ObjectWrapper, type string, ddmStructureKey string, ddmTemplateKey string, displayDateGT int64, displayDateLT int64, status int, reviewDate int64) (int, error) {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId
	_params["groupId"] = groupId
	_params["folderIds"] = folderIds
	_params["classNameId"] = classNameId
	_params["keywords"] = keywords
	liferay.MangleObjectWrapper(_params, "version", "java.lang.Double", version)
	_params["type"] = type
	_params["ddmStructureKey"] = ddmStructureKey
	_params["ddmTemplateKey"] = ddmTemplateKey
	_params["displayDateGT"] = displayDateGT
	_params["displayDateLT"] = displayDateLT
	_params["status"] = status
	_params["reviewDate"] = reviewDate

	_cmd := map[string]interface{}{
		"/journalarticle/search-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64))
	}

	return v, err
}

func (s *Service) SearchCount2(companyId int64, groupId int64, folderIds []interface{}, classNameId int64, articleId string, version *liferay.ObjectWrapper, title string, description string, content string, type string, ddmStructureKey string, ddmTemplateKey string, displayDateGT int64, displayDateLT int64, status int, reviewDate int64, andOperator bool) (int, error) {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId
	_params["groupId"] = groupId
	_params["folderIds"] = folderIds
	_params["classNameId"] = classNameId
	_params["articleId"] = articleId
	liferay.MangleObjectWrapper(_params, "version", "java.lang.Double", version)
	_params["title"] = title
	_params["description"] = description
	_params["content"] = content
	_params["type"] = type
	_params["ddmStructureKey"] = ddmStructureKey
	_params["ddmTemplateKey"] = ddmTemplateKey
	_params["displayDateGT"] = displayDateGT
	_params["displayDateLT"] = displayDateLT
	_params["status"] = status
	_params["reviewDate"] = reviewDate
	_params["andOperator"] = andOperator

	_cmd := map[string]interface{}{
		"/journalarticle/search-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64))
	}

	return v, err
}

func (s *Service) SearchCount3(companyId int64, groupId int64, folderIds []interface{}, classNameId int64, articleId string, version *liferay.ObjectWrapper, title string, description string, content string, type string, ddmStructureKeys []interface{}, ddmTemplateKeys []interface{}, displayDateGT int64, displayDateLT int64, status int, reviewDate int64, andOperator bool) (int, error) {
	_params := make(map[string]interface{})

	_params["companyId"] = companyId
	_params["groupId"] = groupId
	_params["folderIds"] = folderIds
	_params["classNameId"] = classNameId
	_params["articleId"] = articleId
	liferay.MangleObjectWrapper(_params, "version", "java.lang.Double", version)
	_params["title"] = title
	_params["description"] = description
	_params["content"] = content
	_params["type"] = type
	_params["ddmStructureKeys"] = ddmStructureKeys
	_params["ddmTemplateKeys"] = ddmTemplateKeys
	_params["displayDateGT"] = displayDateGT
	_params["displayDateLT"] = displayDateLT
	_params["status"] = status
	_params["reviewDate"] = reviewDate
	_params["andOperator"] = andOperator

	_cmd := map[string]interface{}{
		"/journalarticle/search-count": _params,
	}

	var v int

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = int(res.(float64))
	}

	return v, err
}

func (s *Service) Subscribe(groupId int64) error {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId

	_cmd := map[string]interface{}{
		"/journalarticle/subscribe": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) Unsubscribe(groupId int64) error {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId

	_cmd := map[string]interface{}{
		"/journalarticle/unsubscribe": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) UpdateArticle(userId int64, groupId int64, folderId int64, articleId string, version float64, titleMap map[string]interface{}, descriptionMap map[string]interface{}, content string, layoutUuid string, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["userId"] = userId
	_params["groupId"] = groupId
	_params["folderId"] = folderId
	_params["articleId"] = articleId
	_params["version"] = version
	_params["titleMap"] = titleMap
	_params["descriptionMap"] = descriptionMap
	_params["content"] = content
	_params["layoutUuid"] = layoutUuid
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/journalarticle/update-article": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) UpdateArticle2(groupId int64, folderId int64, articleId string, version float64, titleMap map[string]interface{}, descriptionMap map[string]interface{}, content string, type string, ddmStructureKey string, ddmTemplateKey string, layoutUuid string, displayDateMonth int, displayDateDay int, displayDateYear int, displayDateHour int, displayDateMinute int, expirationDateMonth int, expirationDateDay int, expirationDateYear int, expirationDateHour int, expirationDateMinute int, neverExpire bool, reviewDateMonth int, reviewDateDay int, reviewDateYear int, reviewDateHour int, reviewDateMinute int, neverReview bool, indexable bool, smallImage bool, smallImageURL string, smallFile io.Reader, images map[string]interface{}, articleURL string, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["folderId"] = folderId
	_params["articleId"] = articleId
	_params["version"] = version
	_params["titleMap"] = titleMap
	_params["descriptionMap"] = descriptionMap
	_params["content"] = content
	_params["type"] = type
	_params["ddmStructureKey"] = ddmStructureKey
	_params["ddmTemplateKey"] = ddmTemplateKey
	_params["layoutUuid"] = layoutUuid
	_params["displayDateMonth"] = displayDateMonth
	_params["displayDateDay"] = displayDateDay
	_params["displayDateYear"] = displayDateYear
	_params["displayDateHour"] = displayDateHour
	_params["displayDateMinute"] = displayDateMinute
	_params["expirationDateMonth"] = expirationDateMonth
	_params["expirationDateDay"] = expirationDateDay
	_params["expirationDateYear"] = expirationDateYear
	_params["expirationDateHour"] = expirationDateHour
	_params["expirationDateMinute"] = expirationDateMinute
	_params["neverExpire"] = neverExpire
	_params["reviewDateMonth"] = reviewDateMonth
	_params["reviewDateDay"] = reviewDateDay
	_params["reviewDateYear"] = reviewDateYear
	_params["reviewDateHour"] = reviewDateHour
	_params["reviewDateMinute"] = reviewDateMinute
	_params["neverReview"] = neverReview
	_params["indexable"] = indexable
	_params["smallImage"] = smallImage
	_params["smallImageURL"] = smallImageURL
	_params["smallFile"] = smallFile
	_params["images"] = images
	_params["articleURL"] = articleURL
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/journalarticle/update-article": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) UpdateArticle3(groupId int64, folderId int64, articleId string, version float64, content string, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["folderId"] = folderId
	_params["articleId"] = articleId
	_params["version"] = version
	_params["content"] = content
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/journalarticle/update-article": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) UpdateArticleTranslation(groupId int64, articleId string, version float64, locale string, title string, description string, content string, images map[string]interface{}) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["articleId"] = articleId
	_params["version"] = version
	_params["locale"] = locale
	_params["title"] = title
	_params["description"] = description
	_params["content"] = content
	_params["images"] = images

	_cmd := map[string]interface{}{
		"/journalarticle/update-article-translation": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) UpdateArticleTranslation2(groupId int64, articleId string, version float64, locale string, title string, description string, content string, images map[string]interface{}, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["articleId"] = articleId
	_params["version"] = version
	_params["locale"] = locale
	_params["title"] = title
	_params["description"] = description
	_params["content"] = content
	_params["images"] = images
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/journalarticle/update-article-translation": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) UpdateContent(groupId int64, articleId string, version float64, content string) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["articleId"] = articleId
	_params["version"] = version
	_params["content"] = content

	_cmd := map[string]interface{}{
		"/journalarticle/update-content": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) UpdateStatus(groupId int64, articleId string, version float64, status int, articleURL string, serviceContext *liferay.ObjectWrapper) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["articleId"] = articleId
	_params["version"] = version
	_params["status"] = status
	_params["articleURL"] = articleURL
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/journalarticle/update-status": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

