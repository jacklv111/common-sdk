/*
 * Created on Tue Jul 11 2023
 *
 * Copyright (c) 2023 Company-placeholder. All rights reserved.
 *
 * Author Yubinlv.
 */

package annotation

import aifsclientgo "github.com/jacklv111/aifs-client-go"

type AnnotationTemplate struct {
	Name   string  `json:"name"`
	Labels []Label `json:"labels"`
}

type Label struct {
	Name  string `json:"name"`
	Color int32  `json:"color"`
}

type AnnotationTemplateDetails struct {
	aifsclientgo.AnnotationTemplateDetails
	labelNameIdMap map[string]string
}

func NewAnnotationTemplateDetails(details aifsclientgo.AnnotationTemplateDetails) AnnotationTemplateDetails {
	labelNameIdMap := make(map[string]string)
	for _, label := range details.Labels {
		labelNameIdMap[label.Name] = *label.Id
	}
	return AnnotationTemplateDetails{AnnotationTemplateDetails: details, labelNameIdMap: labelNameIdMap}
}

func (data AnnotationTemplateDetails) GetAnnoTempId() string {
	return *data.Id
}

func (data AnnotationTemplateDetails) GetIdByName(name string) string {
	return data.labelNameIdMap[name]
}
