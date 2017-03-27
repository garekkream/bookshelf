var url_host = "http://localhost:1234/"

function getSettings() {
  console.log("getSettings")
  var urll = url_host + "settings"

  $.ajax({
    type: "GET",
    url: urll,
    success: function(data) {
      configDir = data['configPath'];
      document.getElementById("inputShelfDirectory").value = configDir
    },
    error: function(xhdr, data, err) {
      console.log(err);
    }
  })
}

function getVersion() {
  console.log("getVersion")
  var urll = url_host + "version"

  $.ajax({
    type: "GET",
    url : urll,
    success : function(data) {
      var version = data['version'];

      $("#version").text(version);
    },
    error : function(xhdr, data, err) {
      console.log(err);
    }
  })
}

function addShelf(formData) {
  console.log("addShelf")
  var urll = url_host + "shelfs"
  var status = false

  console.log(JSON.stringify(formData));

  $.ajax({
    type: "POST",
    url: urll,
    data: JSON.stringify(formData),
    dataType: 'json',
    contentType: "application/json",
    success: function(data) {
      console.log(data)

      $('#availableShelfs').append(
        $('<li>').attr('id', formData['name']).addClass("list-group-item ").html(
        formData['name'] + "&emsp; <a><i class='fa fa-upload'></i></a> \
        &emsp; <a data='"+ formData['name'] +"' onclick='removeShelf(this)'><i class='fa fa-times-rectangle'></i></a></div>"))

      document.getElementById("inputShelfDirectory").value = configDir + "/";
      document.getElementById("inputShelfName").value = "";
    },
    error: function(xhdr, data, err) {
      alert(err)
    }
  })
}

function getShelfs() {
  console.log("getShelfs")
  var urll = url_host + "shelfs"

  $.ajax({
    type: "GET",
    url : urll,
    success : function(data) {
      jQuery.each(data, function(index) {
        var name = data[index]['shelfName']
        var active = data[index]['shelfActive']
        var activeShelf = " "

        if(active === true) {
          activeShelf = "&emsp; <i class='fa fa-check'></i>"
        }

        $('#availableShelfs').append(
          $('<li>').attr('id', name).addClass("list-group-item").html(
            name + "&emsp; <a><i class='fa fa-upload'></i></a> \
            &emsp; <a data='"+ name +"' onclick='removeShelf(this)'><i class='fa fa-times-rectangle'></i></a>" + activeShelf +"</div>"));
      });
    },
    error : function(xhdr, data, err) {
      console.log(err)
    }
  })
}

$(document).ready(function() {
  console.log("Website loaded");

  getVersion();
  getSettings();
  getShelfs();
})
