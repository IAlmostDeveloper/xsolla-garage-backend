package controllers

import "errors"

var errNoTask = errors.New("task not found")
var errNoChanges = errors.New("no changes")
var errJsonDecode = errors.New("cannot decode json body")
var errNotFound = errors.New("item not exist")
var errValidateTask = errors.New("error validating task")
var errNoAccess = errors.New("insufficient privileges to access this resource")
