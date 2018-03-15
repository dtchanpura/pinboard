(function() {
    'use strict';

    angular
        .module('app')
        .factory('ConnectionService', ConnectionService);

    ConnectionService.$inject = ['$http'];

    function ConnectionService($http) {
        var service = {};

        service.Connect = Connect;
        service.GetBoard = GetBoard;

        return service;

        function Connect() {
            return $http.get(window.location.origin + '/api', {
                // headers: {
                //     "Authorization": "Bearer " + token
                // }
            }).then(handleSuccess, handleError);
        }

        function GetBoard(id) {
            return $http.get(window.location.origin + '/board/' + id, {
                // headers
            }).then(handleSuccess, handleError)
        }

        function handleSuccess(res) {
            return res.data;
        }

        function handleError(error) {
            return function() {
                console.log('error');
                console.log(error);
                return {
                    ok: false,
                    message: error
                };
            };
        }
    }

})();
