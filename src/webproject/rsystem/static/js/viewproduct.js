/*
 Copyright 2011 The Go Authors.  All rights reserved.
 Use of this source code is governed by a BSD-style
 license that can be found in the LICENSE file.
*/

function ViewProductCtrl($scope, $http) {
  $scope.ViewProducts = [];
  var logError = function(data, status) {
    console.log('code '+status+': '+data);
  };

  var refresh = function() {
    return $http.get('/rautomatic/viewproduct/').
      success(function(data) { $scope.ViewProducts = data.ViewProducts; }).
      error(logError);
  };

  refresh().then(function() {  });
}