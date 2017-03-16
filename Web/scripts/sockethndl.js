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
    $('<li>').attr('id', msg.shelfName).html(
      "<div class='container' " + activeShelf + "> \
        <div class='col-sm-3'>" + msg.shelfName + "</div> \
        <div class='col-sm-3'><a class='btn''>Open</a></div> \
        <div class='col-sm-3'><a class='btn' data='"+ msg.shelfName +"' onclick='removeShelf(this)'>Remove</a></div> \
      </div>"))
  });
}

$(function() {
  getVersion();
  getShelfs();
  //getDebugMode();
})

function removeShelf(data) {
  socket.emit("removeShelf", data.getAttribute('data'), function(data) {
    var str = data
    if (data.length > 1) {
      document.getElementById(data).remove();
    } else {
      alert(data);
      }});
}

function openSettings() {
  document.getElementById("settingsBar").style.width = "30%";
}

function closeSettings() {
  document.getElementById("settingsBar").style.width = "0";
}
