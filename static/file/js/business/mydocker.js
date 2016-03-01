$(document).ready(function() {
  GetImages();
})

/// 获取所有镜像
function GetImages() {
  $.ajax({
    url: "/api/docker/getimages",
    type: 'GET',
    success: function(jsondata) {
      console.log(jsondata);

      if (jsondata.message == "ok") {

        var udata = JSON.parse(jsondata.data);

        $("#tb_dockerimagelist tbody").html("");

        if (jsondata.data != "null") {

          for (var i = 0, l = udata.length; i < udata.length; i++) {

            var shortid = i + 1;


            var tr = "<tr class=\"left aligned\">";
            tr += "<td><div class=\"ui ribbon label\">" + shortid + "</div></td>";
            tr += "<td>" + udata[i].RepoTags + "</td>";
            tr += "<td>" + bytesToSize(udata[i].VirtualSize) + "</td>";
            tr += "<td>" + getLocalTime(udata[i].Created) + "</td>";

            tr += "<td><div class=\"ui buttons\"><button class=\"ui positive button\">建容</button><div class=\"or\"></div><button class=\"ui  button\">删除</button><div class=\"or\"></div><button class=\"ui blue button\">标签</button></div></td>";

            tr += "</tr>";
            $("#tb_dockerimagelist tbody").append(tr);
          }

        }

      }
    }
  });
}


function bytesToSize(bytes) {
  if (bytes === 0) return '0 B';

  var k = 1024;

  sizes = ['B', 'KB', 'MB', 'GB', 'TB', 'PB', 'EB', 'ZB', 'YB'];

  i = Math.floor(Math.log(bytes) / Math.log(k));
  numbyte = (bytes / Math.pow(k, i));

  return numbyte.toPrecision(3) + ' ' + sizes[i];
  //toPrecision(3) 后面保留一位小数，如1.0GB                                                                                                                  //return (bytes / Math.pow(k, i)).toPrecision(3) + ' ' + sizes[i];
}

function getLocalTime(nS) {
  return new Date(parseInt(nS) * 1000).toLocaleString().
  replace("/", "-").
  replace("/", "-").
  replace(/上/g, " ").
  replace(/下/g, " ").
  replace(/午/g, " ").
  replace(/年|月/g, "-").
  replace(/日/g, " ");
}
