'use strict';

xlAide.controller( 'ServerController', function ServerController( $scope ) {

    $scope.importSnapshot = function() {
        if ($scope.serverForm.$valid) {
            console.log('importSnapshot');
        }
    };

    $scope.exportSnapshot = function() {
        if ($scope.serverForm.$valid) {
            console.log('exportSnapshot');
        }
    };

    $scope.isImportDisabled = function() {
        return $scope.serverForm.$invalid;
    };

    $scope.isExportDisabled = function() {
        return $scope.serverForm.$invalid;
    };

});
