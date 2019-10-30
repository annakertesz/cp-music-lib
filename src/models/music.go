/* © Copyright IBM Corp. 2018
All Rights Reserved.
This software is the confidential and proprietary information
of the IBM Corporation. (‘Confidential Information’). Redistribution
of the source code or binary form is not permitted without prior authorization
from the IBM Corporation.
*/

package models

type Song struct{
	title  string
	artist Artist
	album Album
	url string
	tagList []string
}

type Artist struct {
	name string
}

type Album struct {
	name string
	cover string
}