(function() {
    'use strict';

    angular
        .module('app')
        .controller('EditController', EditController);

    EditController.$inject = ['$location', '$cookies', 'ConnectionService'];

    function EditController($location, $cookies, ConnectionService) {
        var vm = this;
        vm.date = (new Date()).toDateString();
        vm.updateBoardTitle = updateBoardTitle;
        vm.updateBlock = updateBlock;
        vm.addParagraph = addParagraph;
        // vm.data contains all details related to blocks
        // vm.data has array object with blocks
        // vm.boardId should be the id for given board.

        (function init() {
            // if($cookies.get('token') !== undefined) {
            //     $location.path('/media');
            // } else {
            //     $location.path('/login');
            // }
            initControllers();
        })();

        function initControllers() {
            vm.boardId = $location.search().boardId;

            ConnectionService.GetBoard(vm.boardId).then(function(response) {
                // console.log(response);
                vm.data = response.data;
                vm.title = response.data.title;
            });
            // var element = document.getElementById("status")
            console.log("OK! for", vm.boardId);
        }

        function updateBoardTitle(boardID, boardTitle) {
            ConnectionService.UpdateBoard(boardID, {
                board: {
                    title: boardTitle
                }
            }).then(function(data) {
                console.log(data)
            })
        }
        function updateBlock(boardID, blockID, post) {
            ConnectionService.UpdateBlock(boardID, blockID, {
                block: post
            }).then(function(data) {
                console.log(data)
            })
        }

        function addParagraph(post) {
            post.paragraphs.push("");
        }
        function addBlock(block) {
            ConnectionService.AddBlock(vm.boardID, block).then(function(data) {
                console.log(data);
            })
        }
    }
})();
