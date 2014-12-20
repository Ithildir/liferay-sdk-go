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

package shoppingorder

import (
	"github.com/ithildir/liferay-sdk-go/liferay"
)

type Service struct {
	session liferay.Session
}

func NewService(s liferay.Session) *Service {
	return &Service{s}
}

func (s *Service) CompleteOrder(groupId int64, number string, ppTxnId string, ppPaymentStatus string, ppPaymentGross float64, ppReceiverEmail string, ppPayerEmail string, serviceContext *liferay.ObjectWrapper) error {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["number"] = number
	_params["ppTxnId"] = ppTxnId
	_params["ppPaymentStatus"] = ppPaymentStatus
	_params["ppPaymentGross"] = ppPaymentGross
	_params["ppReceiverEmail"] = ppReceiverEmail
	_params["ppPayerEmail"] = ppPayerEmail
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/shoppingorder/complete-order": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) DeleteOrder(groupId int64, orderId int64) error {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["orderId"] = orderId

	_cmd := map[string]interface{}{
		"/shoppingorder/delete-order": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) GetOrder(groupId int64, orderId int64) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["orderId"] = orderId

	_cmd := map[string]interface{}{
		"/shoppingorder/get-order": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) SendEmail(groupId int64, orderId int64, emailType string, serviceContext *liferay.ObjectWrapper) error {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["orderId"] = orderId
	_params["emailType"] = emailType
	liferay.MangleObjectWrapper(_params, "serviceContext", "com.liferay.portal.service.ServiceContext", serviceContext)

	_cmd := map[string]interface{}{
		"/shoppingorder/send-email": _params,
	}

	_, err := s.session.Invoke(_cmd)

	return err
}

func (s *Service) UpdateOrder(groupId int64, orderId int64, billingFirstName string, billingLastName string, billingEmailAddress string, billingCompany string, billingStreet string, billingCity string, billingState string, billingZip string, billingCountry string, billingPhone string, shipToBilling bool, shippingFirstName string, shippingLastName string, shippingEmailAddress string, shippingCompany string, shippingStreet string, shippingCity string, shippingState string, shippingZip string, shippingCountry string, shippingPhone string, ccName string, ccType string, ccNumber string, ccExpMonth int, ccExpYear int, ccVerNumber string, comments string) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["orderId"] = orderId
	_params["billingFirstName"] = billingFirstName
	_params["billingLastName"] = billingLastName
	_params["billingEmailAddress"] = billingEmailAddress
	_params["billingCompany"] = billingCompany
	_params["billingStreet"] = billingStreet
	_params["billingCity"] = billingCity
	_params["billingState"] = billingState
	_params["billingZip"] = billingZip
	_params["billingCountry"] = billingCountry
	_params["billingPhone"] = billingPhone
	_params["shipToBilling"] = shipToBilling
	_params["shippingFirstName"] = shippingFirstName
	_params["shippingLastName"] = shippingLastName
	_params["shippingEmailAddress"] = shippingEmailAddress
	_params["shippingCompany"] = shippingCompany
	_params["shippingStreet"] = shippingStreet
	_params["shippingCity"] = shippingCity
	_params["shippingState"] = shippingState
	_params["shippingZip"] = shippingZip
	_params["shippingCountry"] = shippingCountry
	_params["shippingPhone"] = shippingPhone
	_params["ccName"] = ccName
	_params["ccType"] = ccType
	_params["ccNumber"] = ccNumber
	_params["ccExpMonth"] = ccExpMonth
	_params["ccExpYear"] = ccExpYear
	_params["ccVerNumber"] = ccVerNumber
	_params["comments"] = comments

	_cmd := map[string]interface{}{
		"/shoppingorder/update-order": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

func (s *Service) UpdateOrder2(groupId int64, orderId int64, ppTxnId string, ppPaymentStatus string, ppPaymentGross float64, ppReceiverEmail string, ppPayerEmail string) (map[string]interface{}, error) {
	_params := make(map[string]interface{})

	_params["groupId"] = groupId
	_params["orderId"] = orderId
	_params["ppTxnId"] = ppTxnId
	_params["ppPaymentStatus"] = ppPaymentStatus
	_params["ppPaymentGross"] = ppPaymentGross
	_params["ppReceiverEmail"] = ppReceiverEmail
	_params["ppPayerEmail"] = ppPayerEmail

	_cmd := map[string]interface{}{
		"/shoppingorder/update-order": _params,
	}

	var v map[string]interface{}

	res, err := s.session.Invoke(_cmd)

	if err == nil && res != nil {
		v = res.(map[string]interface{})
	}

	return v, err
}

