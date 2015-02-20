'use strict';

xlAide.controller('ServerController', function ServerController($scope, HttpService) {

    $scope.importSnapshot = function () {
        $scope.clear();

        HttpService.get('import',
            {
                jiraIssue: $scope.jiraIssue,
                restartServerAfterImport: $scope.restartServerAfterImport
            }
        ).success(function(data) {
            $scope.jiraIssue = '';
            $scope.successResult = data;
        }).error(function(data, status) {
            $scope.errorResult = data;
            $scope.checkAndNotifyAboutServerConnection(status);
        });
    };

    $scope.exportSnapshot = function () {
        $scope.clear();

        HttpService.get('export',
            {
                jiraIssue: $scope.jiraIssue,
                overwriteAlreadyExported: $scope.overwriteAlreadyExported
            }
        ).success(function(data) {
            $scope.jiraIssue = '';
            $scope.successResult = data;
        }).error(function(data, status) {
            $scope.errorResult = data;
            $scope.checkAndNotifyAboutServerConnection(status);
        });
    };

    $scope.checkAndNotifyAboutServerConnection = function(status) {
        if (status == 0) {
            $scope.errorResult = 'Server is not reachable. Please check that XL-Aide server is up and running';
        }
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
