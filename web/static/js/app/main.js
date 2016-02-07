function log(s) {
    console.log(s);
}

var board_model = null;

var BOARD_ORIENTATION = {
    BLACK_ON_TOP: 0,
    WHITE_ON_TOP: 1,
};

var images = [
    "/static/images/white-pawn.png",
    "/static/images/white-knight.png",
    "/static/images/white-bishop.png",
    "/static/images/white-rook.png",
    "/static/images/white-queen.png",
    "/static/images/white-king.png",
    "/static/images/black-pawn.png",
    "/static/images/black-knight.png",
    "/static/images/black-bishop.png",
    "/static/images/black-rook.png",
    "/static/images/black-queen.png",
    "/static/images/black-king.png",
];

function create_board(orientation) {
    var create_cell = function (row, col, color, r, c) {
        var td = $("<td />")
            .attr("id", row + col)
            // .text("r" + r + ":c" + c)
            .data("row", r)
            .data("col", c)
            .addClass("cell");
        if (color == 0) {
            td.addClass("black-cell");
        }
        return td;
    };
    var board = $("#board");

    _.map(_.range(8,0,-1), function (col) {
        var tr = $("<tr />");
        _.map(_.range(1, 9), function (row) {
            if (orientation === BOARD_ORIENTATION.WHITE_ON_TOP) {
                col = 9 - col;
                row = 9 - row;
            }
            var row_char = String.fromCharCode(96 + row);
            tr.append(create_cell(row_char, col, ((row % 2) + col) % 2, 8 - col, row - 1));
            if (orientation === BOARD_ORIENTATION.WHITE_ON_TOP) {
                col = 9 - col;
                row = 9 - row;
            }
        });
        board.append(tr);
    });
    return board;
}

function create_piece(piece) {
    var img = $("<img />").attr("src", images[piece]).addClass("piece");
    return img;
}

function move(event, ui) {
    var cancel_move = function () {
        $(ui.draggable).animate($(ui.draggable).data('start-position'), 300);
        log("dropped on the same cell");
    }
    if (ui.draggable[0].parentElement === this) {
        cancel_move();
        return;
    }
    //cancel_move();
    var from_square = $(ui.draggable).data('origin-square');
    var to_square = $(this).attr("id");
    $(".move-list").append($("<div />").text(from_square + ' ' + to_square));
    $from = $("#" + from_square);
    ws.send(JSON.stringify({name: "move", data: {
        fromCol: $from.data("col"),
        fromRow: $from.data("row"),
        toCol: $(this).data("col"),
        toRow: $(this).data("row"),
    }}));
    try {
        //board_model.make_move(from_square, to_square);
    } catch (e) {
        cancel_move();
        return;
    }
    //var piece = ui.draggable[0];
    //var cell = $(this);
    //cell.children().remove();
    //$(piece).detach().css({top: 0, left: 0}).appendTo(cell);
    var audio = $("#audio").get(0);
    audio.currentTime = 0;
    audio.play();
}

function sync_with_model() {
    _.map(_.range(1, 9), function (col) {
        _.map(_.range(1, 9), function (row) {
            var row_char = String.fromCharCode(96 + row);
            var square = row_char + col;
            var piece = board_model.get_square(square);
            $('#' + square).empty();
            if (piece != EMPTY_SQUARE) {
                $('#' + square).append(create_piece(piece));
            }
        });
    });
    $(".piece").draggable({
        revert: "invalid",
        start: function (event, ui) {
            $(this).data('start-position', $(this).position());
            $(this).data('origin-square', $(this).parents("td").attr("id"));
            $("td.cell").css("z-index", 99);
            $(this).parents("td").css("z-index", 100);
        },
    });
}

var audio = null;
var ws = null;

$(function() {

    var board_view = create_board(BOARD_ORIENTATION.BLACK_ON_TOP);
    board_model = new BoardModel();
    board_model.reset_pieces();
    sync_with_model();
    $("td.cell").droppable({
        accept: ".piece",
        drop: move,
        hoverClass: "hoverCell",
    });
    $(".piece").draggable({
        revert: "invalid",
        start: function (event, ui) {
            $(this).data('start-position', $(this).position());
            $(this).data('origin-square', $(this).parents("td").attr("id"));
            $("td.cell").css("z-index", 99);
            $(this).parents("td").css("z-index", 100);
        },
    });
    $(".piece").on("click", function() {
        _.each(board_model.get_moves_of_square($(this).parents("td").attr("id")), function (element, index, list) {
            $("#" + element).addClass("movableCell");
        });
        setTimeout(function() {
            $("td.cell").removeClass("movableCell");
        }, 1000);
    });

    ws = new WebSocket("ws://localhost:7777/ws");

    ws.onopen = function() {
        $("div.status").text("Connection established...");
      ws.send(JSON.stringify({name: "connected", data: ""}));
    };

    ws.onmessage = function (evt) {
      var received_msg = evt.data;
      msg = JSON.parse(received_msg);
      if (msg.name == "move") {
      $(".move-list").append($("<div />").text(msg.data));
      board_model.make_move(msg.data.split(" ")[0], msg.data.split(" ")[1]);
      sync_with_model();
      var audio = $("#audio").get(0);
      audio.currentTime = 0;
      audio.play();
      } else if (msg.name == "status") {
        $("div.status").text(msg.data);
      }
    };

    ws.onclose = function() { 
        $("div.status").text("Connection closed...");
      console.log("Connection closed...");
    };
});
