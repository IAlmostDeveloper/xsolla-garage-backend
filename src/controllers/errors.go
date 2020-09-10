package controllers

import "errors"

var errNoTask = errors.New("task not found")
var errNoChanges = errors.New("no changes")
