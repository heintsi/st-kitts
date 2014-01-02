'use strict';
require([], function () {
  console.log('loaded main script');

  var serverUrl = 'http://localhost:8080/';

  function success() { // data
    console.log('we have success');
  }

  var sampleTurnJSON = {
    player: {
      playerID: 'pid01'
    },
    action: 'someActionString'
  };

  $('#endTurn').click(function() {
    console.log('Doing request');
    $.post(
      serverUrl + 'submit/',
      sampleTurnJSON,
      success
    );
  });

});