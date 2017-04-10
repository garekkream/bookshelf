var url_host = "http://localhost:1234/"
var alert_delay = 2500

function showAlert(message, type, closeDelay) {
  if ($("#alerts-container").length == 0) {
    // alerts-container does not exist, create it
    $("body")
    .append( $('<div id="alerts-container" style="position: fixed; width: 50%; left: 25%; top: 10%;">') );
  }

  // default to alert-info; other options include success, warning, danger
  type = type || "info";

  // create the alert div
  var alert = $('<div class="alert alert-' + type + ' fade in">')
    .append(
      $('<button type="button" class="close" data-dismiss="alert">')
        .append("&times;")
    ).append(message);

  // add the alert div to top of alerts-container, use append() to add to bottom
  $("#alerts-container").prepend(alert);

  // if closeDelay was passed - set a timeout to close the alert
  if (closeDelay) {
    window.setTimeout(function() { alert.alert("close") }, closeDelay);
  }
}

function getSettings() {
  var urll = url_host + "settings"

  $.ajax({
    type: "GET",
    url: urll,
    success: function(data) {
      configDir = data['configPath'];
      document.getElementById("inputShelfDirectory").value = configDir
    },
    error: function(xhdr, data, err) {
      showAlert(xhdr.responseText, "danger", alert_delay);
    }
  })
}

function getVersion() {
  var urll = url_host + "version"

  $.ajax({
    type: "GET",
    url : urll,
    success : function(data) {
      var version = data['version'];

      $("#version").text(version);
    },
    error : function(xhdr, data, err) {
      showAlert(xhdr.responseText, "danger", alert_delay);
    }
  })
}

function addShelf(formData) {
  var urll = url_host + "shelfs"
  var status = false

  $.ajax({
    type: "POST",
    url: urll,
    data: JSON.stringify(formData),
    dataType: 'json',
    contentType: "application/json",
    success: function(data) {
      $('#availableShelfs').append(
        $('<li>').attr('id', data['id']).addClass("list-group-item ").html(
        formData['name'] + "&emsp; <a><i class='fa fa-upload'></i></a> \
        &emsp; <a data-id='"+ formData['id'] +"' onclick='removeShelf(this)'><i class='fa fa-times-rectangle'></i></a></div>"))

      document.getElementById("inputShelfDirectory").value = configDir + "/";
      document.getElementById("inputShelfName").value = "";
    },
    error: function(xhdr, data, err) {
      showAlert(xhdr.responseText, "danger", alert_delay);
    }
  })
}

function removeShelf(data) {
  var id = $(data).attr('data-id');
  var urll = url_host + "shelfs/" +id;

  $.ajax({
    type: "DELETE",
    url: urll,
    data: String(id),
    success: function(data) {
      document.getElementById(String(id)).remove()
    },
    error: function(xhdr, data, err) {
      showAlert(xhdr.responseText, "danger", alert_delay);
    }
  })
}

function getShelfs() {
  var urll = url_host + "shelfs"

  $.ajax({
    type: "GET",
    url : urll,
    success : function(data) {
      jQuery.each(data, function(index) {
        var id = data[index]['shelfId']
        var name = data[index]['shelfName']
        var active = data[index]['shelfActive']
        var activeShelf = " "

        if(active === true) {
          activeShelf = "&emsp; <i class='fa fa-check'></i>"
        }

        $('#availableShelfs').append(
          $('<li>').attr('id', id).addClass("list-group-item").html(
            name + "&emsp; <a><i class='fa fa-upload'></i></a> \
            &emsp; <a data-id='"+ id +"' onclick='removeShelf(this)'><i class='fa fa-times-rectangle'></i></a>" + activeShelf +"</div>"));
      });
    },
    error : function(xhdr, data, err) {
      showAlert(xhdr.responseText, "danger", alert_delay);
    }
  })
}

$(document).ready(function() {
  getVersion();
  getSettings();
  getShelfs();
})
