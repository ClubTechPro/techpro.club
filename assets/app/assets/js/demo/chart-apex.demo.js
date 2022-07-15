/*
Template Name: Color Admin - Responsive Admin Dashboard Template build with Twitter Bootstrap 5
Version: 5.1.4
Author: Sean Ngu
Website: http://www.seantheme.com/color-admin/
*/

var handleLineChart = function() {
	"use strict";
	
	$('#apex-line-chart').empty();
	var options = {
		chart: {
			height: 350,
			type: 'line',
			toolbar: {
				show: false
			}
		},
		title: {
			text: 'Average High & Low Temperature',
			align: 'center'
		},
		colors: [app.color.blue, app.color.success],
		dataLabels: {
      enabled: true,
      background: {
        borderWidth: 0
      },
    },
		stroke: {
			curve: 'smooth',
			width: 3
		},
		series: [{
			name: 'High - 2021',
			data: [28, 29, 33, 36, 32, 32, 33]
		}, {
			name: 'Low - 2021',
			data: [12, 11, 14, 18, 17, 13, 13]
		}],
		grid: {
			row: {
				colors: ['rgba('+ app.color.componentColorRgb + ', .1)', 'transparent'], // takes an array which will be repeated on columns
				opacity: 0.5
			},
		},
		markers: {
			size: 4
		},
		xaxis: {
			categories: ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul']
		},
		yaxis: {
			min: 5,
			max: 40
		},
		legend: {
			show: true,
			position: 'top',
			offsetY: -10,
			horizontalAlign: 'right',
      floating: true,
		}
	};

	var chart = new ApexCharts(
		document.querySelector('#apex-line-chart'),
		options
	);

	chart.render();
};

var handleColumnChart = function() {
	"use strict";
	
	$('#apex-column-chart').empty();
	var options = {
		chart: {
			height: 350,
			type: 'bar'
		},
		title: {
			text: 'Profit & Margin Chart',
			align: 'center'
		},
		plotOptions: {
			bar: {
				horizontal: false,
				columnWidth: '55%',
				endingShape: 'rounded'	
			},
		},
		dataLabels: {
			enabled: false
		},
		stroke: {
			show: true,
			width: 2,
			colors: ['transparent']
		},
		colors: ['rgba('+ app.color.componentColorRgb + ', .5)', app.color.indigo, 'rgba('+ app.color.componentColorRgb + ', .25)'],
		series: [{
			name: 'Net Profit',
			data: [44, 55, 57, 56, 61, 58, 63, 60, 66]
		}, {
			name: 'Revenue',
			data: [76, 85, 101, 98, 87, 105, 91, 114, 94]
		}, {
			name: 'Free Cash Flow',
			data: [35, 41, 36, 26, 45, 48, 52, 53, 41]
		}],
		xaxis: {
			categories: ['Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct']
		},
		yaxis: {
			title: {
				text: '$ (thousands)'
			}
		},
		fill: {
			opacity: 1
		},
		tooltip: {
			y: {
				formatter: function (val) {
					return "$ " + val + " thousands"
				}
			}
		}
	};
	
	var chart = new ApexCharts(
		document.querySelector('#apex-column-chart'),
		options
	);

	chart.render();
};

var handleAreaChart = function() {
	"use strict";
	
	$('#apex-area-chart').empty();
	var options = {
		chart: {
			height: 350,
			type: 'area',
		},
		dataLabels: {
			enabled: false
		},
		stroke: {
			curve: 'smooth',
			width: 3
		},
		colors: [app.color.pink, app.color.dark],
		series: [{
			name: 'series1',
			data: [31, 40, 28, 51, 42, 109, 100]
		}, {
			name: 'series2',
			data: [11, 32, 45, 32, 34, 52, 41]
		}],

		xaxis: {
			type: 'datetime',
			categories: ['2019-09-19T00:00:00', '2019-09-19T01:30:00', '2019-09-19T02:30:00', '2019-09-19T03:30:00', '2019-09-19T04:30:00', '2019-09-19T05:30:00', '2019-09-19T06:30:00'],
		},
		tooltip: {
			x: {
				format: 'dd/MM/yy HH:mm'
			},
		}
	};
	
	var chart = new ApexCharts(
		document.querySelector('#apex-area-chart'),
		options
	);

	chart.render();
};

var handleBarChart = function() {
	"use strict";
	
	$('#apex-bar-chart').empty();
	var options = {
		chart: {
			height: 350,
			type: 'bar',
		},
		plotOptions: {
			bar: {
				horizontal: true,
				dataLabels: {
					position: 'top',
				},
			}  
		},
		dataLabels: {
			enabled: true,
			offsetX: -6
		},
		colors: [app.color.orange, 'rgba('+ app.color.componentColorRgb + ', .5)'],
		stroke: {
			show: true,
			width: 1,
			colors: [app.color.componentBg]
		},
		series: [{
			data: [44, 55, 41, 64, 22, 43, 21]
			},{
			data: [53, 32, 33, 52, 13, 44, 32]
		}],
		xaxis: {
			categories: [2013, 2014, 2015, 2016, 2017, 2018, 2019]
		}
	};
	
	var chart = new ApexCharts(
		document.querySelector('#apex-bar-chart'),
		options
	);

	chart.render();
};

var handleMixedChart = function() {
	"use strict";
	
	$('#apex-mixed-chart').empty();
	var options = {
		chart: {
			height: 350,
			type: 'line',
			stacked: false
		},
		dataLabels: {
			enabled: false
		},
		series: [{
			name: 'Income',
			type: 'column',
			data: [1.4, 2, 2.5, 1.5, 2.5, 2.8, 3.8, 4.6]
		}, {
			name: 'Cashflow',
			type: 'column',
			data: [1.1, 3, 3.1, 4, 4.1, 4.9, 6.5, 8.5]
		}, {
			name: 'Revenue',
			type: 'line',
			data: [20, 29, 37, 36, 44, 45, 50, 58]
		}],
		stroke: {
			width: [0, 0, 3]
		},
		colors: [app.color.blue, app.color.success, app.color.orange],
		title: {
			text: 'XYZ - Stock Analysis (2013 - 2021)',
			align: 'left',
			offsetX: 110
		},
		xaxis: {
			categories: [2013, 2014, 2015, 2016, 2017, 2018, 2019, 2020, 2021]
		},
		yaxis: [{
			axisTicks: {
				show: true,
				color: 'rgba('+ app.color.componentColorRgb + ', .15)'
			},
			axisBorder: {
				show: true,
				color: 'rgba('+ app.color.componentColorRgb + ', .15)'
			},
			title: {
				text: "Income (thousand crores)",
				style: {
					color: app.color.componentColor
				}
			},
			tooltip: {
				enabled: true
			}
		},{
			seriesName: 'Income',
			opposite: true,
			axisTicks: {
				show: true,
				color: 'rgba('+ app.color.componentColorRgb + ', .15)'
			},
			axisBorder: {
				show: true,
				color: 'rgba('+ app.color.componentColorRgb + ', .15)'
			},
			title: {
				text: "Operating Cashflow (thousand crores)",
				style: {
					color: app.color.componentColor
				}
			},
		}, {
			seriesName: 'Revenue',
			opposite: true,
			axisTicks: {
				show: true,
			},
			axisBorder: {
				show: true,
				color: app.color.orange
			},
			title: {
				text: "Revenue (thousand crores)",
				style: {
					color: app.color.orange
				}
			}
		}],
		tooltip: {
			fixed: {
				enabled: true,
				position: 'topLeft', // topRight, topLeft, bottomRight, bottomLeft
				offsetY: 30,
				offsetX: 60
			},
		},
		legend: {
			horizontalAlign: 'left',
			offsetX: 40
		}
	};

	var chart = new ApexCharts(
		document.querySelector('#apex-mixed-chart'),
		options
	);

	chart.render();
};

var handleCandlestickChart = function() {
	"use strict";
	
	$('#apex-candelstick-chart').empty();
	var options = {
		chart: {
			height: 350,
			type: 'candlestick'
		},
		series: [{
			data: [{
				x: new Date(1538778600000),
				y: [6629.81, 6650.5, 6623.04, 6633.33]
			},
			{
				x: new Date(1538780400000),
				y: [6632.01, 6643.59, 6620, 6630.11]
			},
			{
				x: new Date(1538782200000),
				y: [6630.71, 6648.95, 6623.34, 6635.65]
			},
			{
				x: new Date(1538784000000),
				y: [6635.65, 6651, 6629.67, 6638.24]
			},
			{
				x: new Date(1538785800000),
				y: [6638.24, 6640, 6620, 6624.47]
			},
			{
				x: new Date(1538787600000),
				y: [6624.53, 6636.03, 6621.68, 6624.31]
			},
			{
				x: new Date(1538789400000),
				y: [6624.61, 6632.2, 6617, 6626.02]
			},
			{
				x: new Date(1538791200000),
				y: [6627, 6627.62, 6584.22, 6603.02]
			},
			{
				x: new Date(1538793000000),
				y: [6605, 6608.03, 6598.95, 6604.01]
			},
			{
				x: new Date(1538794800000),
				y: [6604.5, 6614.4, 6602.26, 6608.02]
			},
			{
				x: new Date(1538796600000),
				y: [6608.02, 6610.68, 6601.99, 6608.91]
			},
			{
				x: new Date(1538798400000),
				y: [6608.91, 6618.99, 6608.01, 6612]
			},
			{
				x: new Date(1538800200000),
				y: [6612, 6615.13, 6605.09, 6612]
			},
			{
				x: new Date(1538802000000),
				y: [6612, 6624.12, 6608.43, 6622.95]
			},
			{
				x: new Date(1538803800000),
				y: [6623.91, 6623.91, 6615, 6615.67]
			},
			{
				x: new Date(1538805600000),
				y: [6618.69, 6618.74, 6610, 6610.4]
			},
			{
				x: new Date(1538807400000),
				y: [6611, 6622.78, 6610.4, 6614.9]
			},
			{
				x: new Date(1538809200000),
				y: [6614.9, 6626.2, 6613.33, 6623.45]
			},
			{
				x: new Date(1538811000000),
				y: [6623.48, 6627, 6618.38, 6620.35]
			},
			{
				x: new Date(1538812800000),
				y: [6619.43, 6620.35, 6610.05, 6615.53]
			},
			{
				x: new Date(1538814600000),
				y: [6615.53, 6617.93, 6610, 6615.19]
			},
			{
				x: new Date(1538816400000),
				y: [6615.19, 6621.6, 6608.2, 6620]
			},
			{
				x: new Date(1538818200000),
				y: [6619.54, 6625.17, 6614.15, 6620]
			},
			{
				x: new Date(1538820000000),
				y: [6620.33, 6634.15, 6617.24, 6624.61]
			},
			{
				x: new Date(1538821800000),
				y: [6625.95, 6626, 6611.66, 6617.58]
			},
			{
				x: new Date(1538823600000),
				y: [6619, 6625.97, 6595.27, 6598.86]
			},
			{
				x: new Date(1538825400000),
				y: [6598.86, 6598.88, 6570, 6587.16]
			},
			{
				x: new Date(1538827200000),
				y: [6588.86, 6600, 6580, 6593.4]
			},
			{
				x: new Date(1538829000000),
				y: [6593.99, 6598.89, 6585, 6587.81]
			},
			{
				x: new Date(1538830800000),
				y: [6587.81, 6592.73, 6567.14, 6578]
			},
			{
				x: new Date(1538832600000),
				y: [6578.35, 6581.72, 6567.39, 6579]
			},
			{
				x: new Date(1538834400000),
				y: [6579.38, 6580.92, 6566.77, 6575.96]
			},
			{
				x: new Date(1538836200000),
				y: [6575.96, 6589, 6571.77, 6588.92]
			},
			{
				x: new Date(1538838000000),
				y: [6588.92, 6594, 6577.55, 6589.22]
			},
			{
				x: new Date(1538839800000),
				y: [6589.3, 6598.89, 6589.1, 6596.08]
			},
			{
				x: new Date(1538841600000),
				y: [6597.5, 6600, 6588.39, 6596.25]
			},
			{
				x: new Date(1538843400000),
				y: [6598.03, 6600, 6588.73, 6595.97]
			},
			{
				x: new Date(1538845200000),
				y: [6595.97, 6602.01, 6588.17, 6602]
			},
			{
				x: new Date(1538847000000),
				y: [6602, 6607, 6596.51, 6599.95]
			},
			{
				x: new Date(1538848800000),
				y: [6600.63, 6601.21, 6590.39, 6591.02]
			},
			{
				x: new Date(1538850600000),
				y: [6591.02, 6603.08, 6591, 6591]
			},
			{
				x: new Date(1538852400000),
				y: [6591, 6601.32, 6585, 6592]
			},
			{
				x: new Date(1538854200000),
				y: [6593.13, 6596.01, 6590, 6593.34]
			},
			{
				x: new Date(1538856000000),
				y: [6593.34, 6604.76, 6582.63, 6593.86]
			},
			{
				x: new Date(1538857800000),
				y: [6593.86, 6604.28, 6586.57, 6600.01]
			},
			{
				x: new Date(1538859600000),
				y: [6601.81, 6603.21, 6592.78, 6596.25]
			},
			{
				x: new Date(1538861400000),
				y: [6596.25, 6604.2, 6590, 6602.99]
			},
			{
				x: new Date(1538863200000),
				y: [6602.99, 6606, 6584.99, 6587.81]
			},
			{
				x: new Date(1538865000000),
				y: [6587.81, 6595, 6583.27, 6591.96]
			},
			{
				x: new Date(1538866800000),
				y: [6591.97, 6596.07, 6585, 6588.39]
			},
			{
				x: new Date(1538868600000),
				y: [6587.6, 6598.21, 6587.6, 6594.27]
			},
			{
				x: new Date(1538870400000),
				y: [6596.44, 6601, 6590, 6596.55]
			},
			{
				x: new Date(1538872200000),
				y: [6598.91, 6605, 6596.61, 6600.02]
			},
			{
				x: new Date(1538874000000),
				y: [6600.55, 6605, 6589.14, 6593.01]
			},
			{
				x: new Date(1538875800000),
				y: [6593.15, 6605, 6592, 6603.06]
			},
			{
				x: new Date(1538877600000),
				y: [6603.07, 6604.5, 6599.09, 6603.89]
			},
			{
				x: new Date(1538879400000),
				y: [6604.44, 6604.44, 6600, 6603.5]
			},
			{
				x: new Date(1538881200000),
				y: [6603.5, 6603.99, 6597.5, 6603.86]
			},
			{
				x: new Date(1538883000000),
				y: [6603.85, 6605, 6600, 6604.07]
			},
			{
				x: new Date(1538884800000),
				y: [6604.98, 6606, 6604.07, 6606]
			}]
		}],
		title: {
			text: 'CandleStick Chart',
			align: 'left'
		},
		xaxis: {
			type: 'datetime'
		},
		yaxis: {
			tooltip: {
				enabled: true
			}
		},
		plotOptions: {
			candlestick: {
				colors: {
					upward: app.color.success,
					downward: app.color.red
				}
			}
		}
	}

	var chart = new ApexCharts(
		document.querySelector('#apex-candelstick-chart'),
		options
	);

	chart.render();
};

var handleBubbleChart = function() {
	"use strict";
	
	$('#apex-bubble-chart').empty();
	function generateData(baseval, count, yrange) {
		var i = 0;
		var series = [];
		while (i < count) {
			var x = Math.floor(Math.random() * (750 - 1 + 1)) + 1;;
			var y = Math.floor(Math.random() * (yrange.max - yrange.min + 1)) + yrange.min;
			var z = Math.floor(Math.random() * (75 - 15 + 1)) + 15;

			series.push([x, y, z]);
			baseval += 86400000;
			i++;
		}
		return series;
	}
	
	var options = {
		chart: {
			height: 350,
			type: 'bubble',
		},
		dataLabels: {
			enabled: false
		},
		colors: [app.color.blue, app.color.orange, app.color.success, app.color.pink],
		series: [{
				name: 'Bubble1',
				data: generateData(new Date('11 Feb 2017 GMT').getTime(), 20, {
					min: 10,
					max: 60
				})
			},
			{
				name: 'Bubble2',
				data: generateData(new Date('11 Feb 2017 GMT').getTime(), 20, {
					min: 10,
					max: 60
				})
			},
			{
				name: 'Bubble3',
				data: generateData(new Date('11 Feb 2017 GMT').getTime(), 20, {
					min: 10,
					max: 60
				})
			},
			{
				name: 'Bubble4',
				data: generateData(new Date('11 Feb 2017 GMT').getTime(), 20, {
					min: 10,
					max: 60
				})
			}
		],
		fill: {
			opacity: 0.8
		},
		title: {
			text: 'Simple Bubble Chart'
		},
		xaxis: {
			tickAmount: 12,
			type: 'category'
		},
		yaxis: {
			max: 70
		}
	}

	var chart = new ApexCharts(
		document.querySelector('#apex-bubble-chart'),
		options
	);

	chart.render();
};

var handleScatterChart = function() {
	"use strict";
	
	$('#apex-scatter-chart').empty();
	var options = {
		chart: {
			height: 350,
			type: 'scatter',
			zoom: {
				enabled: true,
				type: 'xy'
			}
		},
		colors: [app.color.blue, app.color.orange, app.color.success],
		series: [{
			name: "SAMPLE A",
			data: [
				[16.4, 5.4],
				[21.7, 2],
				[25.4, 3],
				[19, 2],
				[10.9, 1],
				[13.6, 3.2],
				[10.9, 7.4],
				[10.9, 0],
				[10.9, 8.2],
				[16.4, 0],
				[16.4, 1.8],
				[13.6, 0.3],
				[13.6, 0],
				[29.9, 0],
				[27.1, 2.3],
				[16.4, 0],
				[13.6, 3.7],
				[10.9, 5.2],
				[16.4, 6.5],
				[10.9, 0],
				[24.5, 7.1],
				[10.9, 0],
				[8.1, 4.7],
				[19, 0],
				[21.7, 1.8],
				[27.1, 0],
				[24.5, 0],
				[27.1, 0],
				[29.9, 1.5],
				[27.1, 0.8],
				[22.1, 2]
			]
		}, {
			name: "SAMPLE B",
			data: [
				[36.4, 13.4],
				[1.7, 11],
				[5.4, 8],
				[9, 17],
				[1.9, 4],
				[3.6, 12.2],
				[1.9, 14.4],
				[1.9, 9],
				[1.9, 13.2],
				[1.4, 7],
				[6.4, 8.8],
				[3.6, 4.3],
				[1.6, 10],
				[9.9, 2],
				[7.1, 15],
				[1.4, 0],
				[3.6, 13.7],
				[1.9, 15.2],
				[6.4, 16.5],
				[0.9, 10],
				[4.5, 17.1],
				[10.9, 10],
				[0.1, 14.7],
				[9, 10],
				[12.7, 11.8],
				[2.1, 10],
				[2.5, 10],
				[27.1, 10],
				[2.9, 11.5],
				[7.1, 10.8],
				[2.1, 12]
			]
		}, {
			name: "SAMPLE C",
			data: [
				[21.7, 3],
				[23.6, 3.5],
				[24.6, 3],
				[29.9, 3],
				[21.7, 20],
				[23, 2],
				[10.9, 3],
				[28, 4],
				[27.1, 0.3],
				[16.4, 4],
				[13.6, 0],
				[19, 5],
				[22.4, 3],
				[24.5, 3],
				[32.6, 3],
				[27.1, 4],
				[29.6, 6],
				[31.6, 8],
				[21.6, 5],
				[20.9, 4],
				[22.4, 0],
				[32.6, 10.3],
				[29.7, 20.8],
				[24.5, 0.8],
				[21.4, 0],
				[21.7, 6.9],
				[28.6, 7.7],
				[15.4, 0],
				[18.1, 0],
				[33.4, 0],
				[16.4, 0]
			]
		}],
		xaxis: {
			tickAmount: 10,
			labels: {
				formatter: function(val) {
					return parseFloat(val).toFixed(1)
				}
			}
		},
		yaxis: {
			tickAmount: 7
		}
	}

	var chart = new ApexCharts(
		document.querySelector('#apex-scatter-chart'),
		options
	);

	chart.render();
};

var handleHeatmapChart = function() {
	"use strict";
	
	$('#apex-heatmap-chart').empty();
	function generateData(count, yrange) {
		var i = 0;
		var series = [];
		while (i < count) {
			var x = 'w' + (i + 1).toString();
			var y = Math.floor(Math.random() * (yrange.max - yrange.min + 1)) + yrange.min;

			series.push({
				x: x,
				y: y
			});
			i++;
		}
		return series;
	}

	var options = {
		chart: {
			height: 350,
			type: 'heatmap',
		},
		dataLabels: {
			enabled: false
		},
		stroke: {
    	show: true,
			colors: [app.color.componentBg],
			width: 2,
			dashArray: 0
		},
		colors: [app.color.blue],
		series: [{
				name: 'Metric1',
				data: generateData(18, {
					min: 0,
					max: 90
				})
			},
			{
				name: 'Metric2',
				data: generateData(18, {
					min: 0,
					max: 90
				})
			},
			{
				name: 'Metric3',
				data: generateData(18, {
					min: 0,
					max: 90
				})
			},
			{
				name: 'Metric4',
				data: generateData(18, {
					min: 0,
					max: 90
				})
			},
			{
				name: 'Metric5',
				data: generateData(18, {
					min: 0,
					max: 90
				})
			},
			{
				name: 'Metric6',
				data: generateData(18, {
					min: 0,
					max: 90
				})
			},
			{
				name: 'Metric7',
				data: generateData(18, {
					min: 0,
					max: 90
				})
			},
			{
				name: 'Metric8',
				data: generateData(18, {
					min: 0,
					max: 90
				})
			},
			{
				name: 'Metric9',
				data: generateData(18, {
					min: 0,
					max: 90
				})
			}
		],
		title: {
			text: 'HeatMap Chart (Single color)'
		}
	}

	var chart = new ApexCharts(
		document.querySelector('#apex-heatmap-chart'),
		options
	);

	chart.render();
};

var handlePieChart = function() {
	"use strict";
	
	$('#apex-pie-chart').empty();
	var options = {
		chart: {
			height: 365,
			type: 'pie',
		},
		stroke: {
    	show: true,
			colors: [app.color.componentBg],
			width: 2,
			dashArray: 0
		},
		dataLabels: {
			dropShadow: {
				enabled: false,
				top: 1,
				left: 1,
				blur: 1,
				opacity: 0.45
			}
		},
		colors: [app.color.pink, app.color.orange, app.color.blue, app.color.success, app.color.indigo],
		labels: ['Team A', 'Team B', 'Team C', 'Team D', 'Team E'],
		series: [44, 55, 13, 43, 22],
		title: {
			text: 'HeatMap Chart (Single color)'
		}
	};

	var chart = new ApexCharts(
		document.querySelector('#apex-pie-chart'),
		options
	);

	chart.render();
};

var handleRadialBarChart = function() {
	"use strict";
	
	$('#apex-radial-bar-chart').empty();
	var options = {
		chart: {
			height: 350,
			type: 'radialBar',
		},
		plotOptions: {
			radialBar: {
				offsetY: -10,
				startAngle: 0,
				endAngle: 270,
				hollow: {
					margin: 5,
					size: '30%',
					background: 'transparent',
					image: undefined,
				},
				dataLabels: {
					name: {
						show: false,

					},
					value: {
						show: false,
					}
				}
			}
		},
		colors: ['#1ab7ea', '#0084ff', '#39539E', '#0077B5'],
		series: [76, 67, 61, 90],
		labels: ['Vimeo', 'Messenger', 'Facebook', 'LinkedIn'],
		legend: {
			show: true,
			floating: true,
			position: 'left',
			offsetX: 140,
			offsetY: 15,
			labels: {
				useSeriesColors: true,
			},
			markers: {
				size: 0
			},
			formatter: function(seriesName, opts) {
				return seriesName + ":  " + opts.w.globals.series[opts.seriesIndex]
			},
			itemMargin: {
				horizontal: 1,
			}
		}
	}

	var chart = new ApexCharts(
		document.querySelector('#apex-radial-bar-chart'),
		options
	);

	chart.render();
};

var handleRadarChart = function() {
	"use strict";
	
	$('#apex-radar-chart').empty();
	var options = {
		chart: {
			height: 320,
			type: 'radar',
		},
		series: [{
			name: 'Series 1',
			data: [20, 100, 40, 30, 50, 80, 33],
		}],
		labels: ['Sunday', 'Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday'],
		plotOptions: {
			radar: {
				size: 140,
				polygons: {
					strokeColor: 'rgba('+ app.color.componentColorRgb + ', .15)',
					fill: {
						colors: ['rgba('+ app.color.componentColorRgb + ', .05)', 'rgba('+ app.color.componentColorRgb + ', .15)']
					}
				}
			}
		},
		title: {
			text: 'Radar with Polygon Fill'
		},
		colors: [app.color.blue],
		markers: {
			size: 4,
			colors: [app.color.componentBg],
			strokeColor: app.color.blue,
			strokeWidth: 2,
		},
		tooltip: {
			y: {
				formatter: function(val) {
					return val
				}
			}
		},
		yaxis: {
			tickAmount: 7,
			labels: {
				formatter: function(val, i) {
					if (i % 2 === 0) {
						return val
					} else {
						return ''
					}
				}
			}
		}
	}

	var chart = new ApexCharts(
		document.querySelector('#apex-radar-chart'),
		options
	);

	chart.render();
};


var ChartApex = function () {
	"use strict";
	return {
		//main function
		init: function () {
			Apex = {
				grid: {
					borderColor: 'rgba('+ app.color.componentColorRgb + ', .15)'
				},
				title: {
					style: {
						color: app.color.componentColor
					}
				},
				legend: {
					labels: {
						colors: app.color.componentColor
					}
				},
				xaxis: {
					axisBorder: {
						show: true,
						color: 'rgba('+ app.color.componentColorRgb + ', .25)',
						height: 1,
						width: '100%',
						offsetX: 0,
						offsetY: -1
					},
					axisTicks: {
						show: true,
						borderType: 'solid',
						color: 'rgba('+ app.color.componentColorRgb + ', .25)',
						height: 6,
						offsetX: 0,
						offsetY: 0
					},
					labels: {
						style: {
							colors: app.color.componentColor,
							fontSize: app.font.size,
							fontFamily: app.font.family,
							fontWeight: 400,
							cssClass: 'apexcharts-xaxis-label',
						}
					}
				},
				yaxis: {
					labels: {
						style: {
							colors: app.color.componentColor,
							fontSize: app.font.size,
							fontFamily: app.font.family,
							fontWeight: 400,
							cssClass: 'apexcharts-xaxis-label',
						}
					}
				}
			}
			
			handleLineChart();
			handleColumnChart();
			handleAreaChart();
			handleBarChart();
			handleMixedChart();
			handleCandlestickChart();
			handleBubbleChart();
			handleScatterChart();
			handleHeatmapChart();
			handlePieChart();
			handleRadialBarChart();
			handleRadarChart();
		}
	};
}();

$(document).ready(function() {
	ChartApex.init();
	
	$(document).on('theme-change', function() {
		ChartApex.init();
	});
});