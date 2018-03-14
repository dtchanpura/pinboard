(function() {
    'use strict';

    angular
        .module('app')
        .controller('MainController', MainController);

    MainController.$inject = ['$location', '$cookies'];

    function MainController($location, $cookies) {
        var vm = this;
        vm.date = (new Date()).toDateString();
        // vm.data contains all details related to blocks
        // vm.data has array object with blocks
        vm.data = {
            blocks: [
                {
                    "title": "Notice 1",
                    "paragraphs": [
                        "This is a sample notice for testing."
                    ],
                    "side": "left"
                }
            ]
        };
        // vm.boardId should be the id for given board.

        (function init(){
            // if($cookies.get('token') !== undefined) {
            //     $location.path('/media');
            // } else {
            //     $location.path('/login');
            // }
            initControllers();
        })();

        function initControllers() {
            vm.boardId = $location.search().boardId;
            // var element = document.getElementById("status")
            console.log("OK! for", vm.boardId);
        }
    }
})();
