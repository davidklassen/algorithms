// input handling start

var data = '';
process.stdin.resume();
process.stdin.on('data', function (chunk) { data += chunk; });

process.stdin.on('end', function () {
  main(data);
});

// input handling end

function main(data) {
  data = data.split('\n');
  var pair = function (src) {
    return src.split(' ').map((i) => parseInt(i, 10));
  }

  var size = pair(data[0]);
  var init = pair(data[1]);
  var target = pair(data[2]);

  console.log(findPath(init, target, size));
}

function getAvailableMoves(board, positions) {
  var sizeY = board.length;
  var sizeX = board[0].length;
  var moves = [];
  var possible = [
    [1, 2],
    [2, 1],
    [2, -1],
    [1, -2],
    [-1, -2],
    [-2, -1],
    [-2, 1],
    [-1, 2]
  ];

  for (var i = 0; i < positions.length; i++) {
    moves = moves.concat(possible.map(function (p) {
      return [p[0] + positions[i][0], p[1] + positions[i][1]];
    }).filter(function (m) {
      return m[0] >= 0 && m[0] < sizeX &&
             m[1] >= 0 && m[1] < sizeY &&
             board[m[1]][m[0]] === null;
    }));
  }

  return moves;
}

function mark(board, positions, depth) {
  for (var i = 0; i < positions.length; i++) {
    board[positions[i][1]][positions[i][0]] = depth;
  }

  return board;
}

function unmarked(board) {
  return !!board
    .reduce((m, i) => m.concat(i), [])
    .filter((i) => i === null).length;
}

function findPath(init, target, size) {
  var board = [];

  for (var i = 0; i < size[1]; i++) {
    board[i] = [];
    for (var j = 0; j < size[0]; j++) {
      board[i][j] = null;
    }
  }

  var depth = 0;
  var available = [init];
  board = mark(board, available, depth);

  while (unmarked(board)) {
    available = getAvailableMoves(board, available);
    board = mark(board, available, ++depth);
  }

  return board[target[1]][target[0]];
}
