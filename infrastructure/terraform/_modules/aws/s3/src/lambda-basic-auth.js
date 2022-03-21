"use strict";
function handler(event) {
  var request = event.request;
  var headers = request.headers;

  // Configure authentication
  var authUser = `${username}`;
  var authPass = `${password}`;

  var authString = authUser + ":" + authPass;
  var rawHeader =
    (headers["authorization"] && headers["authorization"].value) || "";
  var authHeader = _base64urlDecode(rawHeader.substr(6)); //Remove 'Basic ' prefix
  console.log("rawHeader", rawHeader);
  if (authHeader != authString) {
    // Require Basic authentication
    var response = {
      statusCode: 401,
      statusDescription: "Unauthorized",
      headers: {
        "www-authenticate": { value: "Basic" },
      },
    };
    return response;
  }

  // Continue request processing if authentication passed
  return request;
}

function _base64urlDecode(str) {
  return String.bytesFrom(str, "base64url");
}
