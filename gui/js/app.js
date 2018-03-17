(function () {
    'use strict';

    angular.module('app', ['ngRoute','ngCookies'])
  .config(config);
  //.run(run);

config.$inject = ['$routeProvider', '$locationProvider'];
function config($routeProvider, $locationProvider){
  $routeProvider
  .when('/', {
    templateUrl: 'templates/view.html',
    controller: 'MainController',
    controllerAs: 'vm'
  })
  .when('/edit', {
    templateUrl: 'templates/edit.html',
    controller: 'EditController',
    controllerAs: 'vm'
  })
  // // .when('/socket', {
  // //   templateUrl: 'templates/socket.html',
  // //   controller: 'SocketController',
  // //   controllerAs: 'vm'
  // // })
  // .when('/login', {
  //   templateUrl: 'templates/login.html',
  //   controller: 'LoginController',
  //   controllerAs: 'vm'
  // })
  // .when('/logout', {
  //   templateUrl: 'templates/login.html',
  //   controller: 'LogoutController',
  //   controllerAs: 'vm'
  // })
  .otherwise({ redirectTo: '/'});
}

})();
