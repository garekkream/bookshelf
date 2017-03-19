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
    activeShelf = "activeShelf"
  }

  $('#availableShelfs').append(
    $('<li>').attr('id', msg.shelfName).addClass("list-group-item " + activeShelf).html(
      msg.shelfName + "&emsp; <a><i class='fa fa-upload'></i></a> \
      &emsp; <a data='"+ msg.shelfName +"' onclick='removeShelf(this)'><i class='fa fa-times-rectangle'></i></a></div>"))
  });
}

$(function() {
  getVersion();
  getShelfs();
  //getDebugMode();

  socket.on("errorMsg", function(msg) {
    alert(msg);
  });
})

function removeShelf(data) {
  socket.emit("removeShelf", data.getAttribute('data'), function(data) {
    if (data.length > 1) {
      document.getElementById(data).remove();
    }
})}

function openSettings() {
  document.getElementById("settingsBar").style.width = "30%";
}

function closeSettings() {
  document.getElementById("settingsBar").style.width = "0";
}
