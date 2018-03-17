(function() {
    'use strict';

    angular
        .module('app')
        .controller('EditController', EditController);

    EditController.$inject = ['$location', '$cookies', 'ConnectionService'];

    function EditController($location, $cookies, ConnectionService) {
        var vm = this;
        vm.date = (new Date()).toDateString();
        vm.title = "";
        vm.updateBoardTitle = updateBoardTitle;
        vm.updateBlock = updateBlock;
        vm.deleteBlock = deleteBlock;

        (function init() {
            vm.boardID = vm.boardID === undefined ? $location.search().boardID : vm.boardID;
            initControllers();
        })();

        function initControllers() {
            console.log(vm.title);

            if (vm.boardID !== undefined) {
                ConnectionService.GetBoard(vm.boardID).then(function(response) {
                    vm.data = response.data;
                    vm.title = response.data.title;
                    vm.dataRight = {blocks: []};
                    vm.dataLeft = {blocks: []};
                    for (var i = 0; i < response.data.blocks.length; i++) {
                        if (response.data.blocks[i].side == "right") {
                            vm.dataRight.blocks.push(response.data.blocks[i]);
                        } else {
                            vm.dataLeft.blocks.push(response.data.blocks[i]);
                        }
                    }
                });
            }
            if ($location.search().boardID === undefined && vm.boardID !== undefined) {
                location.href = '#/edit?boardID=' + vm.boardID;
            }
            console.log("OK! for", vm.boardID);
        }

        function updateBoardTitle(boardID, boardTitle) {
            if (boardID !== undefined) {
                ConnectionService.UpdateBoard(boardID, {
                    board: {
                        title: boardTitle
                    }
                }).then(function(data) {
                    console.log(data)
                });
            } else {
                ConnectionService.AddBoard({
                    board: {
                        title: boardTitle
                    }
                }).then(function(response) {
                    console.log(response.data);
                    vm.boardID = response.data.id;
                    console.log(vm.boardID);
                    initControllers();
                });
            }
        }


        function updateBlock(boardID, blockID, post) {
            if (blockID !== undefined) {
                ConnectionService.UpdateBlock(boardID, blockID, {
                    block: post
                }).then(function(data) {
                    console.log(data);
                })
            } else {
                addBlock(boardID, post);
            }
        }

        function addBlock(boardID, block) {
            ConnectionService.AddBlock(boardID, block).then(function(data) {
                console.log(data);
                initControllers();
            })
        }

        function deleteBlock(boardID, blockID) {
            ConnectionService.DeleteBlock(boardID, blockID).then(function(data) {
                console.log(data);
                initControllers();
            })

        }

    }
})();
