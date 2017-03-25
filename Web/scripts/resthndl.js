var url_host = "http://localhost:1234/"

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
  getShelfs();
})
