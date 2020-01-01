/*
 Copyright 2011 The Go Authors.  All rights reserved.
 Use of this source code is governed by a BSD-style
 license that can be found in the LICENSE file.
*/

function TaskCtrl($scope, $http) {
  $scope.tasks = [];
  $scope.working = false;

  var logError = function(data, status) {
    console.log('code '+status+': '+data);
    $scope.working = false;
  };

  var refresh = function() {
    return $http.get('/todo/task/').
      success(function(data) { $scope.tasks = data.Tasks; }).
      error(logError);
  };

  $scope.addTodo = function() {
    $scope.working = true;
    $http.post('/todo/task/', {Title: $scope.todoText}).
      error(logError).
      success(function() {
        refresh().then(function() {
          $scope.working = false;
          $scope.todoText = '';
        })
      });
  };

  $scope.toggleDone = function(task) {
    data = {ID: task.Id, Title: task.Title, Done: !task.Done,Des:task.Des,Bgntime:task.Bgntime,Endtime:task.Endtime}
    $http.put('/todo/task/'+task.Id, data).
      error(logError).
      success(function() { task.Done = !task.Done });
  };

  refresh().then(function() { $scope.working = false; });
}