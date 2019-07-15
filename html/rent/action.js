
var building = "";

$(document).ready(function () {

    searchBuilding();

    $("#add-building").click(function () {
        $("#building").val($("#building-name").val());

        //clear result
        $("#building-user-result").html("");
        $("#building-user-result-box").show(500);
    });

    $("#add-building-user").click(function () {
        var building = $("#building").val();
        var room = $("#room").val();
        var wechat = $("#wechat").val();
        var message = $("#message").val();

        if (wechat == "") {
            setResult("错误", "请填写微信号");
            return;
        }

        addBuildingUser(building, room, wechat, message);
    });
})

function searchBuilding() {
    // get
    $.get("/go/search-building",
        function (result) {
            var json = JSON.parse(result);

            html = "<p class='mt-2 mb-0'>选择自己的大楼</p>";
            for (const key in json.Building) {
                if (json.Building.hasOwnProperty(key)) {
                    const f = json.Building[key];
                    html += "<label class='btn btn-outline-primary mr-2 mt-2'><input type='radio' name='options' autocomplete='off'>" + f.Name + "</label>";
                }
            }

            // modify building result
            $("#building-result").html(html);

            // add click event
            $("#building-result").on("click", "label", function () {
                building = $(this).text();
                $("#building").val(building);
                searchBuildingUser(building);
            });

            $("#building-result-box").show(500);
        });
}

function searchBuildingUser(building) {
    // get
    $.get("/go/search-building-user",
        {
            building: building,
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
                        <strong class='d-block text-gray-dark'>" + u.Wechat;

                    if (u.Room != "") {
                        html += "  (" + u.Room + ")";
                    }

                    html += "</strong>" + u.Message + "</p>";
                }
            }

            if (noUser) {
                html = "<p class='mt-2'>还没有人提交这个大楼的信息呢</p>";
            }

            html += "<hr class='my-3'>";

            // show building result
            $("#building-user-result").html(html);
            $("#building-user-result-box").show(500);
        });
}

function addBuildingUser(building, room, wechat, message) {
    $("#add-building-user").attr("disabled", true);

    // post
    $.post("/go/add-building-user",
        {
            building: building,
            room: room,
            wechat: wechat,
            message: message
        },
        function (data, status) {
            // request not success
            if (status != "success") {
                setResult("错误", "服务器错误");
                $("#add-building-user").attr("disabled", false);
                return;
            }

            // parse data to json object
            var json = JSON.parse(data);

            if (json.Result != "Success") {
                setResult("错误", "服务器错误");
                $("#add-building-user").attr("disabled", false);
                return;
            }

            searchBuildingUser(building);
            searchBuilding();
            $("#add-building-user").attr("disabled", false);
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
