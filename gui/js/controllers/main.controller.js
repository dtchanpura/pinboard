(function() {
    'use strict';

    angular
        .module('app')
        .controller('MainController', MainController);

    MainController.$inject = ['$location', '$cookies', 'ConnectionService'];

    function MainController($location, $cookies, ConnectionService) {
        var vm = this;
        vm.date = (new Date()).toDateString();
        // vm.data contains all details related to blocks
        // vm.data has array object with blocks
        vm.data = {};
        // vm.boardID should be the id for given board.

        (function init() {
            // if($cookies.get('token') !== undefined) {
            //     $location.path('/media');
            // } else {
            //     $location.path('/login');
            // }
            initControllers();
            setInterval(function() {
                // location.reload();
                initControllers();
            }, 10000);

        })();

        function initControllers() {
            vm.boardID = $location.search().boardID;

            ConnectionService.GetBoard(vm.boardID).then(function(response) {
                if (response.reload) {
                    console.log('reloaded.');
                    location.reload(true);
                }
                vm.dataRight = {blocks: []};
                vm.dataLeft = {blocks: []};
                vm.data = response.data;
                vm.title = response.data.title;
                for (var i = 0; i < response.data.blocks.length; i++) {
                    if (response.data.blocks[i].side == "right") {
                        vm.dataRight.blocks.push(response.data.blocks[i]);
                    } else {
                        vm.dataLeft.blocks.push(response.data.blocks[i]);
                    }
                }
            });
            console.log("OK! for", vm.boardID);
        }
    }
})();
