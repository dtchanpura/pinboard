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
        service.GetAllBoards = GetAllBoards;
        service.AddBlock = AddBlock;
        service.UpdateBoard = UpdateBoard;
        service.UpdateBlock = UpdateBlock;


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

        function GetAllBoards() {
            return $http.get(window.location.origin + '/board', {
                // headers
            }).then(handleSuccess, handleError)
        }

        function UpdateBoard(boardID, board) {
            return $http.put(window.location.origin + '/board/' + boardID,
                board
            ).then(handleSuccess, handleError)
        }

        function UpdateBlock(boardID, blockID, block) {
            return $http.put(window.location.origin + '/board/' + boardID + '/block/' + blockID,
                block
            ).then(handleSuccess, handleError)
        }

        function AddBlock(boardID, block) {
            return $http.post(window.location.origin + '/board/' + boardID,
                block
            ).then(handleSuccess, handleError)
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
