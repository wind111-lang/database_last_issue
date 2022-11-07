        window.onload = function() {
            var conn;
            var msg = document.getElementById("msg");
            var log = document.getElementById("log");
            
            if(window.location.protocol == 'http:'){
              window.location.protocol = 'https:';
            }

            function appendLog(item) {
                var doScroll = log.scrollTop > log.scrollHeight - log.clientHeight - 1;
                log.appendChild(item);
                if (doScroll) {
                    log.scrollTop = log.scrollHeight - log.clientHeight;
                }
            }

            document.getElementById("form").onsubmit = function() {
                var obj = '{"user": "' + userName + '","message": "' + msg.value + '"}'
                var json = JSON.stringify(obj)
                var json_obj = JSON.parse(json)

                if (!conn) {
                    return false;
                }
                if (!json_obj) {
                    return false;
                }
                conn.send(json_obj);
                msg.value = "";
                return false;
            };

            if (window["WebSocket"]) {
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
                user.innerHTML = userName + ' でログインしています';

                conn.onmessage = function(evt) {
                    var res = evt.data;
                    var data = JSON.parse(res);

                    var name_data = data.name;
                    var text_data = data.text;

                    var fullmsg = name_data + ": " + text_data;
                    var item = document.createElement("div");

                    item.innerText = fullmsg;
                    appendLog(item);
                };
            } else {
                var item = document.createElement("div");
                item.innerHTML = "<b>Your browser does not support WebSockets.</b>";
                appendLog(item);
            }
        };
