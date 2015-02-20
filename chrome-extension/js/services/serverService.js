xlAide.factory('HttpService', function ($http, xlAideHost) {

    return {
        'get': function(url, params) {
            var defaults  = {"headers": {"Accept": 'application/json'}};
            return $http.get(xlAideHost + url, {
                params: params
            }, defaults);
        }
    };

});