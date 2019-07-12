
var date = "";
var flight = "";

$(document).ready(function () {

    $("#search-flight").click(function () {
        // get value
        var m = $("#month").val();
        var d = $("#date").val();
        date = m + "-" + d;

        searchFlight(date);
    });

    $("#add-flight").click(function () {
        var c = $("#code").val().toUpperCase();
        addFlight(c, date);
    });

    $("#add-flight-user").click(function () {
        var wechat = $("#wechat").val();
        var message = $("#message").val();

        addFlightUser(date, flight, wechat, message);
    });
})

function searchFlight(date) {
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
                    html += "<label class='btn btn-outline-primary mr-2'><input type='radio' name='options' autocomplete='off'>" + f.Code + "</label>";
                }
            }

            if (noFlight) {
                html = "<p class='mt-2'>这一天还没有航班记录呢</p>";
            } else {
                html = "<p class='mt-2'>选择自己的航班</p>" + html;
            }

            // modify flight result
            $("#flight-result").html(html);

            // add click event
            $("#flight-result").on("click", "label", function () {
                flight = $(this).text();
                searchFlightUser(date, flight);
            });

            $("#flight-result-box").show(500);
        });
}

function addFlight(code, date) {
    // post
    $.post("/go/add-flight",
        {
            code: code,
            date: date
        },
        function (data, status) {
            // request not success
            if (status != "success") {
                alert("服务器错误");
                return;
            }

            // parse data to json object
            var json = JSON.parse(data);

            if (json.Result != "Success") {
                alert("服务器错误");
                return;
            }

            searchFlight(date);
            return;
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
            $("#flight-user-result").html(html);
            $("#flight-user-result-box").show(500);
        });
}

function addFlightUser(date, code, wechat, message) {
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
                alert("服务器错误");
                return;
            }

            // parse data to json object
            var json = JSON.parse(data);

            if (json.Result != "Success") {
                alert("服务器错误");
                return;
            }

            searchFlightUser(date, code);
            return;
        });
}