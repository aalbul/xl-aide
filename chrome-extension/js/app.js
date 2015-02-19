'use strict';

var xlAide = angular.module('xl-aide', []);

xlAide.config(
    function ($httpProvider, $locationProvider) {
        $locationProvider.html5Mode(true);
    })
    .constant('xlAideHost', 'http://localhost:3000/');