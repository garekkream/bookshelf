var socket = io();

function getVersion() {
  socket.emit("getVersion")
  socket.on("setVersion", function(msg) {
    $("#version").text(msg);
  });
}

function getDebugMode() {
  socket.emit("getDebugMode")
  socket.on("setDebugMode", function(msg) {
    if (String(msg) === "true") {
      $("#debugMode").attr("checked", true);
    } else {
      $("#debugMode").attr("checked", false);
    }
  });
}

function getShelfs() {
  socket.emit("getShelfs");
  socket.on("setShelfs", function(msg) {

  var activeShelf = " "

  if(msg.shelfActive === true) {
    activeShelf = "id='activeShelf'"
  }

  $("#availableShelfs").append(
    $('<li>').html(
      "<div class='container' " + activeShelf + "><div class='col-sm-2'>" + msg.shelfName + "</div><div class='col-sm-2'>" + msg.shelfPath + "</div></div>"))
  });
}

$(function() {
  getVersion();
  getShelfs();
  getDebugMode();
})


function openSettings() {
  document.getElementById("settingsBar").style.width = "30%";
}

function closeSettings() {
  document.getElementById("settingsBar").style.width = "0";
}
