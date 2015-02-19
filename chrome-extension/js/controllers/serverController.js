'use strict';

xlAide.controller('ServerController', function ServerController($scope, ServerService) {

    $scope.importSnapshot = function () {
        ServerService.import(
            {
                jiraIssue: $scope.jiraIssue,
                restartServerAfterImport: $scope.restartServerAfterImport
            }
        );
    };

    $scope.exportSnapshot = function () {
        ServerService.export(
            {
                jiraIssue: $scope.jiraIssue,
                overwriteAlreadyExported: $scope.overwriteAlreadyExported
            }
        );
    };

    $scope.isImportDisabled = function () {
        return $scope.serverForm.$invalid;
    };

    $scope.isExportDisabled = function () {
        return $scope.serverForm.$invalid;
    };

});
