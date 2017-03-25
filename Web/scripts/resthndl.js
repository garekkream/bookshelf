var url_host = "http://localhost:1234/"

function getVersion() {
  console.log("getVersion")
  var urll = url_host + "version"

  $.ajax({
    type: "GET",
    url : urll,
    success : function(data) {
      var version = data['version'];
      console.log(version)
      $("#version").text(version);
    },
    failure : function(xhdr, data, err) {
      console.log(err);
    }
  })
}

$(document).ready(function() {
  console.log("Website loaded");

  getVersion();
})
