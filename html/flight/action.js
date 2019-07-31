$(document).ready(function () {
    
    $("#add-flight-user").click(function () {
        var month = $("#add-month").val();
        var day = $("#add-date").val();
        var flight = $("#add-code").val().toUpperCase();
        var wechat = $("#add-wechat").val();
        var message = $("#add-message").val();
        
        if (month == "" || day == "") {
            setResult("错误", "请填写出发日期");
            return;
        }

        if (flight == "") {
            setResult("错误", "请填写航班号");
            return;
        }

        if (wechat == "") {
            setResult("错误", "请填写微信号");
            return;
        }

        var date = month + "-" + day;
        addFlightUser(date, flight, wechat, message);
    });

    $("#code-search-flight").click(function () {
        // get value
        var m = $("#code-month").val();
        var d = $("#code-date").val();

        if (m == "" || d == "") {
            setResult("错误", "请填写日期");
            return;
        }

        date = m + "-" + d;

        searchFlight(date, "");
    });

    $("#time-search").click(function () {
        var month = $("#time-month").val();
        var day = $("#time-date").val();
        var airport = $("#time-airport").val();
        var start = $("#time-start").val();
        var end = $("#time-end").val();
                
        if (month == "") {
            month="8";
        }

        if (day == "") {
            setResult("错误", "请填写出发日期");
            return;
        }

        if(airport == "不限") {
            airport = "";
        }

        if (start == "") {
            start = "0";
        }

        if (end == "") {
            end = "24";
        }

        var date = month + "-" + day;
        start = start +":00";
        end = end +":00";

        searchFlightTime(date, airport, start, end);
    });
})

function searchFlight(date, code) {
    // get
    $.get("/go/search-flight",
        {
            date: date
        },
        function (result) {
            var json = JSON.parse(result);

            html = "";
            var noFlight = true;
            for (const key in json.Flight) {
                if (json.Flight.hasOwnProperty(key)) {
                    const f = json.Flight[key];
                    noFlight = false;
                    html += "<label id='" + f.Code + "' class='btn btn-outline-primary mr-2 mt-2'><input type='radio' name='options' autocomplete='off'>" + f.Code + "</label>";
                }
            }

            if (noFlight) {
                html = "<p class='mt-2'>这一天还没有航班记录呢</p>";
            } else {
                html = "<p class='mt-2 mb-0'>选择自己的航班</p>" + html;
            }

            // modify flight result
            $("#code-flight-result").html(html);

            // add click event
            $("#code-flight-result").on("click", "label", function () {
                flight = $(this).text();
                searchFlightUser(date, flight);
            });

            // trigger button
            if (code != "") {
                $("#" + code).click();
            }

            $("#code-flight-result-box").show(500);
        });
}

function searchFlightUser(date, code) {
    // get
    $.get("/go/search-flight-user",
        {
            date: date,
            code: code
        },
        function (result) {
            var json = JSON.parse(result);

            html = "";
            var noUser = true;
            for (const key in json.User) {
                if (json.User.hasOwnProperty(key)) {
                    const u = json.User[key];
                    noUser = false;
                    html += "<p class='media-body pb-2 mb-0 lh-125'> \
                        <strong class='d-block text-gray-dark'>" + u.Wechat + "</strong> \
                        " + u.Message + "</p>";
                }
            }

            if (noUser) {
                html = "<p class='mt-2'>这一天还没有人飞这一班航班呢</p>";
            }

            // show flight result
            $("#code-flight-user-result").html(html);
            $("#code-flight-user-result-box").show(500);
        });
}

function addFlightUser(date, code, wechat, message) {
    $("#add-flight-user").attr("disabled", true);

    // post
    $.post("/go/add-flight-user",
        {
            code: code,
            date: date,
            wechat: wechat,
            message: message
        },
        function (data, status) {
            // request not success
            if (status != "success") {
                setResult("错误", "服务器错误");
                $("#add-flight-user").attr("disabled", false);
                return;
            }

            // parse data to json object
            var json = JSON.parse(data);

            if (json.Result != "Success") {
                setResult("错误", "服务器错误");
                $("#add-flight-user").attr("disabled", false);
                return;
            }

            setResult("成功", "成功提交信息");
            $("#add-flight-user").attr("disabled", false);
            return;
        });
}

function searchFlightTime(date, airport, start, end) {
    $("#time-search").attr("disabled", true);

    // post
    $.post("/go/search-flight-time",
        {
            date: date,
            airport: airport,
            start: start,
            end: end
        },
        function (data, status) {
            // request not success
            if (status != "success") {
                setResult("错误", "服务器错误");
                $("#time-search").attr("disabled", false);
                return;
            }

            // parse data to json object
            var json = JSON.parse(data);

            html = "";
            var noUser = true;
            for (const key in json.User) {
                if (json.User.hasOwnProperty(key)) {
                    const u = json.User[key];
                    noUser = false;
                    var strong = u.Wechat;
                    var s = u.Code + " &emsp; " + u.Dep + "-" + u.Arr + " &emsp; " + u.DepTime + "-" + u.ArrTime + "<br>" + u.Message;
                    html += "<p class='media-body pb-2 mb-0 lh-125'> \
                        <strong class='d-block text-gray-dark'>" + strong + "</strong> \
                        " + s + "</p>";
                }
            }

            if (noUser) {
                html = "<p class='mt-2'>没有查询到结果</p>";
            }

            // show flight result
            $("#time-result").html(html);
            $("#time-result-box").show(500);

            $("#time-search").attr("disabled", false);
            return;
        });
}

/*******************************************************************
    func:   setResult
    brief:  set modal and show
*******************************************************************/
function setResult(title, message) {
    // set modal
    $("#modal-title").text(title);
    $("#modal-body").text(message);

    // show modal
    $("#result").modal('show');
}
