xlAide.factory('ServerService', function ($http, xlAideHost) {

    return {
        'import': function(params) {
            var defaults  = {"headers": {"Accept": 'application/json'}};
            return $http.get(xlAideHost + 'import', {
                params: params
            }, defaults);
        },
        'export': function(params) {
            var defaults = {"headers": {"Accept": 'application/json', "Content-Type": 'application/json'}};
            return $http.get(xlAideHost + 'export', {
                params: params
            }, defaults);
        }
    };

});