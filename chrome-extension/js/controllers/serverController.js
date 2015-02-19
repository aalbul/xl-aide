'use strict';

xlAide.controller('ServerController', function ServerController($scope, ServerService) {

    $scope.importSnapshot = function () {
        $scope.clear();

        ServerService.import(
            {
                jiraIssue: $scope.jiraIssue,
                restartServerAfterImport: $scope.restartServerAfterImport
            }
        ).success(function(data) {
            $scope.jiraIssue = '';
            $scope.successResult = data;
        }).error(function(data) {
            $scope.errorResult = data;
        });
    };

    $scope.exportSnapshot = function () {
        $scope.clear();

        ServerService.export(
            {
                jiraIssue: $scope.jiraIssue,
                overwriteAlreadyExported: $scope.overwriteAlreadyExported
            }
        ).success(function(data) {
            $scope.jiraIssue = '';
            $scope.successResult = data;
        }).error(function(data) {
            $scope.errorResult = data;
        });
    };

    $scope.clear = function() {
      $scope.errorResult = '';
      $scope.successResult = '';
    };

    $scope.isImportDisabled = function () {
        return $scope.serverForm.$invalid;
    };

    $scope.isExportDisabled = function () {
        return $scope.serverForm.$invalid;
    };

});
