var configDir = ""

function updatePath() {
  document.getElementById("inputShelfDirectory").value = configDir + "/" + document.getElementById("inputShelfName").value + ".shelf";
}

function closeModalShelf() {
  document.getElementById("inputShelfDirectory").value = configDir + "/";
  document.getElementById("inputShelfName").value = "";
}

function addShelfItem() {
  var name = $('#inputShelfName').val()
  var path = $('#inputShelfDirectory').val()

  if (name === '') {
    alert("Empty name!")
  } else {
    var obj = new Object();
    obj.name = name;
    obj.path = path;

    addShelf(obj)

    $("#newShelf").modal("hide");
  }
}

function setDebug() {
  var val = $('#debugMode').text()
  var obj = new Object()

  if (val == "Enable") {
    obj.debugMode = "true";
  } else {
    obj.debugMode = "false";
  }

  setDebugMode(obj)
}
