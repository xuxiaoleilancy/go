/*
 Copyright 2011 The Go Authors.  All rights reserved.
 Use of this source code is governed by a BSD-style
 license that can be found in the LICENSE file.
*/

function ProductCtrl($scope, $http) {
  $scope.Products = [];
  var logError = function(data, status) {
    console.log('code '+status+': '+data);
  };

  var refresh = function() {
    return $http.get('/rautomatic/product/').
      success(function(data) { $scope.Products = data.Products; }).
      error(logError);
  };

  $scope.addProduct = function() {
    $http.post('/rautomatic/product/', {Name: $scope.Name}).
      error(logError).
      success(function() {
        refresh().then(function() {
          	$scope.Name = '';
        })
      });
  };

  refresh().then(function() {  });
}