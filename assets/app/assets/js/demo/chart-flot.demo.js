/*
Template Name: Color Admin - Responsive Admin Dashboard Template build with Twitter Bootstrap 5
Version: 5.1.4
Author: Sean Ngu
Website: http://www.seantheme.com/color-admin/
*/

function showFlotTooltip(x, y, contents) {
	$('<div id="tooltip" class="flot-tooltip">' + contents + '</div>').css({
		top: y,
		left: x + 35,
		opacity: 0.80
	}).appendTo('body').fadeIn(200);
}
    
var handleBasicChart = function () {
	'use strict';
	
	$('#basic-chart').empty();
	var d1 = [], d2 = [], d3 = [];
	for (var x = 0; x < Math.PI * 2; x += 0.25) {
		d1.push([x, Math.sin(x)]);
		d2.push([x, Math.cos(x)]);
	}
	for (var z = 0; z < Math.PI * 2; z += 0.1) {
		d3.push([z, Math.tan(z)]);
	}
	if ($('#basic-chart').length !== 0) {
		$.plot($('#basic-chart'), [
			{ label: 'data 1',  data: d1, color: app.color.blue, shadowSize: 0 },
			{ label: 'data 2',  data: d2, color: app.color.success, shadowSize: 0 }
		], {
			series: {
				lines: { show: true },
				points: { show: false }
			},
			xaxis: {
				min: 0,
				max: 6,
				tickColor: 'rgba('+ app.color.darkRgb + ', .3)'
			},
			yaxis: {
				min: -2,
				max: 2,
				tickColor: 'rgba('+ app.color.darkRgb + ', .3)'
			},
			grid: {
				borderColor: 'rgba('+ app.color.darkRgb + ', .15)',
				borderWidth: 1,
				backgroundColor: 'rgba('+ app.color.darkRgb + ', .035)',
				tickColor: 'rgba('+ app.color.darkRgb + ', .15)'
			}
		});
	}
};

var handleStackedChart = function () {
	'use strict';
	
	$('#stacked-chart').empty();
	var d1 = [];
	for (var a = 0; a <= 5; a += 1) {
		d1.push([a, parseInt(Math.random() * 5)]);
	}
	var d2 = [];
	for (var b = 0; b <= 5; b += 1) {
		d2.push([b, parseInt(Math.random() * 5 + 5)]);
	}
	var d3 = [];
	for (var c = 0; c <= 5; c += 1) {
		d3.push([c, parseInt(Math.random() * 5 + 5)]);
	}
	var d4 = [];
	for (var d = 0; d <= 5; d += 1) {
		d4.push([d, parseInt(Math.random() * 5 + 5)]);
	}
	var d5 = [];
	for (var e = 0; e <= 5; e += 1) {
		d5.push([e, parseInt(Math.random() * 5 + 5)]);
	}
	var d6 = [];
	for (var f = 0; f <= 5; f += 1) {
		d6.push([f, parseInt(Math.random() * 5 + 5)]);
	}
    
	var xData = [{
		data:d1,
		color: app.color.gray500,
		label: 'China',
		bars: { fillColor: app.color.gray500 }
	}, {
		data:d2,
		color: app.color.gray700,
		label: 'Russia',
		bars: { fillColor: app.color.gray700 }
	}, {
		data:d3,
		color: app.color.success,
		label: 'Canada',
		bars: { fillColor: app.color.success }
	}, {
		data:d4,
		color: app.color.purple,
		label: 'Japan',
		bars: { fillColor: app.color.purple }
	}, {
		data:d5,
		color: app.color.blue,
		label: 'USA',
		bars: { fillColor: app.color.blue }
	}, {
		data:d6,
		color: app.color.cyan,
		label: 'Others',
		bars: { fillColor: app.color.cyan }
	}];

	$.plot('#stacked-chart', xData, { 
		xaxis: {  
			tickColor: 'rgba('+ app.color.darkRgb + ', .15)',  
			ticks: [[0, 'MON'], [1, 'TUE'], [2, 'WED'], [3, 'THU'], [4, 'FRI'], [5, 'SAT']],
			autoscaleMargin: 0.05
		},
		yaxis: { tickColor: 'rgba('+ app.color.darkRgb + ', .15)', ticksLength: 5},
		grid: { 
			hoverable: true, 
			tickColor: 'rgba('+ app.color.darkRgb + ', .15)',
			borderWidth: 1,
			borderColor: 'rgba('+ app.color.darkRgb + ', .15)',
			backgroundColor: 'rgba('+ app.color.darkRgb + ', .035)'
		},
		series: {
			stack: true,
			lines: { show: false, fill: false, steps: false },
			bars: { show: true, barWidth: 0.6, align: 'center', fillColor: null },
			highlightColor: 'rgba('+ app.color.darkRgb + ', .5)',
		},
		legend: {
			show: true,
			position: 'ne',
			noColumns: 1
		}
	});

	var previousXValue = null;
	var previousYValue = null;

	$('#stacked-chart').bind('plothover', function (event, pos, item) {
		if (item) {
			var y = item.datapoint[1] - item.datapoint[2];

			if (previousXValue != item.series.label || y != previousYValue) {
				previousXValue = item.series.label;
				previousYValue = y;
				$('#tooltip').remove();

				showFlotTooltip(item.pageX, item.pageY, y + ' ' + item.series.label);
			}
		} else {
			$('#tooltip').remove();
			previousXValue = null;
			previousYValue = null;       
		}
	});
};

var handleTrackingChart = function () {
	'use strict';
	
	$('#tracking-chart').empty();
	var sin = [], cos = [];
	for (var i = 0; i < 14; i += 0.1) {
		sin.push([i, Math.sin(i)]);
		cos.push([i, Math.cos(i)]);
	}
    
	function updateLegend() {
		updateLegendTimeout = null;

		var pos = latestPosition;
		var axes = plot.getAxes();
		if (pos.x < axes.xaxis.min || pos.x > axes.xaxis.max || pos.y < axes.yaxis.min || pos.y > axes.yaxis.max) {
			return;
		}
		var i, j, dataset = plot.getData();
		for (i = 0; i < dataset.length; ++i) {
			var series = dataset[i];

			for (j = 0; j < series.data.length; ++j) {
				if (series.data[j][0] > pos.x) {
					break;
				}
			}
			var y, p1 = series.data[j - 1], p2 = series.data[j];
			if (p1 === null) {
				y = p2[1];
			} else if (p2 === null) {
				y = p1[1];
			} else {
				y = p1[1];
			}
			legends.eq(i).text(series.label.replace(/=.*/, '= ' + y.toFixed(2)));
		}
	}
	if ($('#tracking-chart').length !== 0) {
		var plot = $.plot($('#tracking-chart'), [ 
			{ data: sin, label: 'Series1', color: app.color.gray500, shadowSize: 0},
			{ data: cos, label: 'Series2', color: app.color.blue, shadowSize: 0} 
		], {
			series: { 
				lines: { show: true } 
			},
			crosshair: {
				mode: 'x', 
				color: app.color.red 
			},
			grid: { 
				hoverable: true, 
				autoHighlight: false, 
				borderColor: 'rgba('+ app.color.darkRgb + ', .15)', 
				borderWidth: 1,
				backgroundColor: 'rgba('+ app.color.darkRgb + ', .035)',
				tickColor: 'rgba('+ app.color.darkRgb + ', .15)'
			},
			yaxis: { tickColor: 'rgba('+ app.color.darkRgb + ', .15)' },
			xaxis: {
				tickColor: 'rgba('+ app.color.darkRgb + ', .15)'
			},
			legend: { show: true }
		});

		var legends = $('#tracking-chart .legendLabel');
		legends.each(function () {
			$(this).css('width', $(this).width());
		});

		var updateLegendTimeout = null;
		var latestPosition = null;

		$('#tracking-chart').bind('plothover',  function (pos) {
			latestPosition = pos;
			if (!updateLegendTimeout) {
				updateLegendTimeout = setTimeout(updateLegend, 50);
			}
		});
	}
};

var handleBarChart = function () {
	'use strict';
	
	$('#bar-chart').empty();
	if ($('#bar-chart').length !== 0) {
		var data = [[0, 10], [1, 8], [2, 4], [3, 13], [4, 17], [5, 9]];
		var ticks = [[0, 'JAN'], [1, 'FEB'], [2, 'MAR'], [3, 'APR'], [4, 'MAY'], [5, 'JUN']];
		$.plot('#bar-chart', [{ label: 'Bounce Rate', data: data, color: app.color.componentColor }], {
			series: {
				bars: {
					show: true,
					barWidth: 0.6,
					align: 'center',
					fill: true,
					fillColor: 'rgba('+ app.color.componentColorRgb + ', .25)',
					zero: true
				}
			},
			xaxis: {
				tickColor: 'rgba('+ app.color.darkRgb + ', .15)',
				autoscaleMargin: 0.05,
				ticks: ticks
			},
			yaxis: {
				tickColor: 'rgba('+ app.color.darkRgb + ', .15)'
			},
			grid: {
				borderColor: 'rgba('+ app.color.darkRgb + ', .15)',
				borderWidth: 1,
				backgroundColor: 'rgba('+ app.color.darkRgb + ', .035)',
				tickColor: 'rgba('+ app.color.darkRgb + ', .15)'
			},
			legend: {
				noColumns: 0
			},
		});
	}
};

var handleInteractivePieChart = function () {
	'use strict';
	
	$('#interactive-pie-chart').empty();
	if ($('#interactive-pie-chart').length !== 0) {
		var data = [];
		var series = 3;
		var colorArray = [app.color.orange, 'rgba('+ app.color.componentColorRgb + ', .5)', 'rgba('+ app.color.componentColorRgb + ', .25)'];
		for( var i = 0; i < series; i++) {
			data[i] = { label: 'Series'+(i+1), data: Math.floor(Math.random()*100)+1, color: colorArray[i]};
		}
		$.plot($('#interactive-pie-chart'), data, {
			series: {
				pie: { 
					show: true,
					stroke: {
						color: app.color.componentBg
					}
				}
			},
			grid: {
				hoverable: true,
				clickable: true
			}
		});
		$('#interactive-pie-chart').bind('plotclick', function(event, pos, obj) {
			if (!obj) {
				return;
			}
			var percent = parseFloat(obj.series.percent).toFixed(2);
			alert(obj.series.label + ': ' + percent + '%');
		});
	}
};

var handleDonutChart = function () {
	'use strict';
	
	$('#donut-chart').empty();
	if ($('#donut-chart').length !== 0) {
		var data = [];
		var series = 3;
		var colorArray = [app.color.gray900, app.color.gray500, app.color.red];
		var nameArray = ['Unique Visitor', 'Bounce Rate', 'Total Page Views', 'Avg Time On Site', '% New Visits'];
		var dataArray = [20,14,12,31,23];
		
		for( var i = 0; i < series; i++) {
			data[i] = { label: nameArray[i], data: dataArray[i], color: colorArray[i] };
		}

		$.plot($('#donut-chart'), data, {
			series: {
				pie: { 
					innerRadius: 0.5,
					show: true,
					combine: {
						threshold: 0.1
					},
					stroke: {
						color: app.color.componentBg
					}
				}
			},
			grid:{borderWidth:0, hoverable: true, clickable: true},
			legend: {
				show: false
			}
		});
	}
};

var handleInteractiveChart = function () {
	'use strict';
	
	$('#interactive-chart').empty();
	if ($('#interactive-chart').length !== 0) {
		var d1 = [[0, 42], [1, 53], [2,66], [3, 60], [4, 68], [5, 66], [6,71],[7, 75], [8, 69], [9,70], [10, 68], [11, 72], [12, 78], [13, 86]];
		var d2 = [[0, 12], [1, 26], [2,13], [3, 18], [4, 35], [5, 23], [6, 18],[7, 35], [8, 24], [9,14], [10, 14], [11, 29], [12, 30], [13, 43]];

		$.plot($('#interactive-chart'), [{
			data: d1, 
			label: 'Page Views', 
			color: app.color.blue,
			lines: { show: true, fill:false, lineWidth: 2.5 },
			points: { show: true, radius: 5, fillColor: app.color.componentBg },
			shadowSize: 0
		}, {
			data: d2,
			label: 'Visitors',
			color: app.color.green,
			lines: { show: true, fill:false, lineWidth: 2.5, fillColor: '' },
			points: { show: true, radius: 5, fillColor: app.color.componentBg },
			shadowSize: 0
		}], {
			xaxis: {  tickColor: 'rgba('+ app.color.darkRgb + ', .3)',tickSize: 2 },
			yaxis: {  tickColor: 'rgba('+ app.color.darkRgb + ', .3)', tickSize: 20 },
			grid: { 
				hoverable: true, 
				clickable: true,
				tickColor: 'rgba('+ app.color.darkRgb + ', .15)',
				borderWidth: 1,
				borderColor: 'rgba('+ app.color.darkRgb + ', .15)',
				backgroundColor: 'rgba('+ app.color.darkRgb + ', .035)'
			},
			legend: {
				noColumns: 1,
				show: true
			}
		});
		
		var previousPoint = null;
		$('#interactive-chart').bind('plothover', function (event, pos, item) {
			console.log(pos);
			$('#x').text(pos.x.toFixed(2));
			$('#y').text(pos.y.toFixed(2));
			if (item) {
				if (previousPoint !== item.dataIndex) {
					previousPoint = item.dataIndex;
					$('#tooltip').remove();
					var y = item.datapoint[1].toFixed(2);

					var content = item.series.label + ' ' + y;
					showFlotTooltip(item.pageX, item.pageY, content);
				}
			} else {
				$('#tooltip').remove();
				previousPoint = null;            
			}
			event.preventDefault();
		});
	}
};

var handleLiveUpdatedChart = function () {
	'use strict';
	
  $('#live-updated-chart').empty();
	function update() {
		plot.setData([ getRandomData() ]);
		plot.draw();
		setTimeout(update, updateInterval);
	}
    
	function getRandomData() {
		if (data.length > 0) {
			data = data.slice(1);
		}
		while (data.length < totalPoints) {
			var prev = data.length > 0 ? data[data.length - 1] : 50;
			var y = prev + Math.random() * 10 - 5;
			if (y < 0) {
				y = 0;
			}
			if (y > 100) {
				y = 100;
			}
			data.push(y);
		}
		var res = [];
		for (var i = 0; i < data.length; ++i) {
			res.push([i, data[i]]);
		}
		return res;
	}
	if ($('#live-updated-chart').length !== 0) {
		var data = [], totalPoints = 150;
		var updateInterval = 1000;

		$('#updateInterval').val(updateInterval).change(function () {
			var v = $(this).val();
			if (v && !isNaN(+v)) {
				updateInterval = +v;
				if (updateInterval < 1) {
					updateInterval = 1;
				}
				if (updateInterval > 2000) {
					updateInterval = 2000;
				}
				$(this).val('' + updateInterval);
			}
		});

		var plot = $.plot($('#live-updated-chart'), [{ label: 'Server stats', data: getRandomData() }], {
			series: { 
				shadowSize: 0, 
				color: app.color.success, 
				lines: { 
					show: true, 
					fill:true 
				} 
			},
			yaxis: { 
				min: 0, 
				max: 100, 
				tickColor: 'rgba('+ app.color.darkRgb + ', .15)' 
			},
			xaxis: { 
				show: true, 
				tickColor: 'rgba('+ app.color.darkRgb + ', .15)' 
			},
			grid: {
				borderWidth: 1,
				borderColor: 'rgba('+ app.color.darkRgb + ', .15)',
				backgroundColor: 'rgba('+ app.color.darkRgb + ', .035)',
				tickColor: 'rgba('+ app.color.darkRgb + ', .15)' 
			},
			legend: {
				noColumns: 1,
				show: true
			}
		});

		update();
	}
};

var Chart = function () {
	'use strict';
	return {
		//main function
		init: function () {
			handleBasicChart();
			handleInteractiveChart();
			handleBarChart();
			handleLiveUpdatedChart();
			handleInteractivePieChart();
			handleStackedChart();
			handleTrackingChart();
			handleDonutChart();
		}
	};
}();

$(document).ready(function() {
	Chart.init();
	
	$(document).on('theme-change', function() {
		Chart.init();
	});
});