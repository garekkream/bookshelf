var url_host = "http://localhost:1234/"

function getVersion() {
  console.log("getVersion")
  var urll = url_host + "version"

  $.ajax({
    type: "GET",
    url : urll,
  })
}

$(document).ready(function() {
  console.log("Website loaded");

  getVersion();
})
