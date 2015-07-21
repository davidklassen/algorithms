function qsort(arr) {
  if (!arr.length) return [];

  var pivot = arr.splice(Math.random() * arr.length | 0, 1)[0];
  var left = arr.filter(function (el) { return el <= pivot; });
  var right = arr.filter(function (el) { return el > pivot; });

  return qsort(left).concat(pivot).concat(qsort(right));
}

// input handling

var data = '';
process.stdin.resume();
process.stdin.on('data', function (chunk) { data += chunk; });

process.stdin.on('end', function () {
  var arr = data
    .split('\n')
    .slice(0, -1)
    .map(function (el) { return parseInt(el, 10); });

  qsort(arr).forEach(function (el) { console.log(el); });
});
