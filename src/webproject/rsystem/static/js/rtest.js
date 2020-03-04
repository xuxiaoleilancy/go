/*
 Copyright 2011 The Go Authors.  All rights reserved.
 Use of this source code is governed by a BSD-style
 license that can be found in the LICENSE file.
*/

function RTestCtrl($scope, $http) {
  $scope.cmds = [];
  $scope.working = false;

  var logError = function(data, status) {
    console.log('code '+status+': '+data);
    $scope.working = false;
  };

  var refresh = function() {
    return $http.get('/rtest/cmds/').
      success(function(data) { $scope.cmds = data.Cmds; }).
      error(logError);
  };

  $scope.exeCmd = function() {
    $scope.working = true;
    $http.post('/rtest/cmds/', {Cmd: $scope.todoText}).
      error(logError).
      success(function() {
        refresh().then(function() {
          $scope.working = false;
          $scope.todoText = '';
        })
      });
  };

  $scope.toggleDone = function(cmd) {
    data = {ID: cmd.Id, Title: cmd.Title, Des:cmd.Des,Params:cmd.Params,Help:cmd.Help}
    $http.put('/rtest/cmd/'+task.Id, data).
      error(logError).
      success(function() {  });
  };

  refresh().then(function() { $scope.working = false; });
}