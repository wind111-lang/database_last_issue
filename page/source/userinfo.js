window.onload = function() {
    if(window.location.protocol == 'http:'){
      window.location.protocol = 'https:';
    }
    conn = new WebSocket("wss://" + document.location.host + "/ws");
    var gotcookie = document.cookie;
    var cokkieItem = gotcookie.split(";");
    for (i = 0; i < cokkieItem.length; i++) {
        if (cokkieItem[i].indexOf("user=") != -1) {
            var userName = cokkieItem[i].split("=")[1];
            break;
        }
    }

    var user = document.getElementById("user");

    user.innerHTML = 'ユーザ名: ' + userName;
}
