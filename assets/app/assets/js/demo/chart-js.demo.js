/*
Template Name: Color Admin - Responsive Admin Dashboard Template build with Twitter Bootstrap 5
Version: 5.1.4
Author: Sean Ngu
Website: http://www.seantheme.com/color-admin/
*/

if (document.querySelector('.app-cover')) {
	Chart.defaults.color = 'rgba(255,255,255, .75)';
	Chart.defaults.borderColor = 'rgba(255,255,255, .15)';
	Chart.defaults.scale.ticks.color = 'rgba(255,255,255, .75)';
}
var randomScalingFactor = function() { 
	return Math.round(Math.random()*100)
};

var lineChart, barChart, radarChart, polarAreaChart, pieChart, doughnutChart;

var handleChartJs = function() {
	Chart.defaults.color = 'rgba('+ app.color.componentColorRgb + ', .65)';
	Chart.defaults.font.family = app.font.family;
	Chart.defaults.font.weight = 600;
	Chart.defaults.scale.grid.color = 'rgba('+ app.color.componentColorRgb + ', .15)';
	Chart.defaults.scale.ticks.backdropColor = 'rgba('+ app.color.componentColorRgb + ', 0)';

	var lineChartData = {
		labels: ['January', 'February', 'March', 'April', 'May', 'June', 'July'],
		datasets: [{
			label: 'Dataset 1',
			borderColor: app.color.blue,
			pointBackgroundColor: app.color.componentBg,
			pointRadius: 4,
			borderWidth: 2,
			backgroundColor: 'rgba('+ app.color.blueRgb + ', .3)',
			data: [randomScalingFactor(), randomScalingFactor(), randomScalingFactor(), randomScalingFactor(), randomScalingFactor(), randomScalingFactor(), randomScalingFactor()]
		}, {
			label: 'Dataset 2',
			borderColor: 'rgba('+ app.color.componentColorRgb + ', .85)',
			pointBackgroundColor: app.color.componentBg,
			pointRadius: 4,
			borderWidth: 2,
			backgroundColor: 'rgba('+ app.color.componentColorRgb + ', .5)',
			data: [randomScalingFactor(), randomScalingFactor(), randomScalingFactor(), randomScalingFactor(), randomScalingFactor(), randomScalingFactor(), randomScalingFactor()]
		}]
	};

	var barChartData = {
		labels: ['January', 'February', 'March', 'April', 'May', 'June', 'July'],
		datasets: [{
			label: 'Dataset 1',
			borderWidth: 1,
			borderColor: app.color.indigo,
			backgroundColor: 'rgba('+ app.color.indigoRgb + ', .3)',
			data: [randomScalingFactor(), randomScalingFactor(), randomScalingFactor(), randomScalingFactor(), randomScalingFactor(), randomScalingFactor(), randomScalingFactor()]
		}, {
			label: 'Dataset 2',
			borderWidth: 1,
			borderColor: 'rgba('+ app.color.componentColorRgb + ', .85)',
			backgroundColor: 'rgba('+ app.color.componentColorRgb + ', .3)',
			data: [randomScalingFactor(), randomScalingFactor(), randomScalingFactor(), randomScalingFactor(), randomScalingFactor(), randomScalingFactor(), randomScalingFactor()]
		}]
	};

	var radarChartData = {
		labels: ['Eating', 'Drinking', 'Sleeping', 'Designing', 'Coding', 'Cycling', 'Running'],
		datasets: [{
			label: 'Dataset 1',
			borderWidth: 2,
			borderColor: app.color.red,
			pointBackgroundColor: app.color.red,
			pointRadius: 2,
			backgroundColor: 'rgba('+ app.color.redRgb + ', .2)',
			data: [65,59,90,81,56,55,40]
		}, {
			label: 'Dataset 2',
			borderWidth: 2,
			borderColor: app.color.componentColor,
			pointBackgroundColor: app.color.componentColor,
			pointRadius: 2,
			backgroundColor: 'rgba('+ app.color.componentColorRgb + ', .2)',
			data: [28,48,40,19,96,27,100]
		}]
	};

	var polarAreaData = {
		labels: ['Dataset 1', 'Dataset 2', 'Dataset 3', 'Dataset 4', 'Dataset 5'],
		datasets: [{
			data: [300, 160, 100, 200, 120],
			backgroundColor: ['rgba('+ app.color.indigoRgb + ',.7)', 'rgba('+ app.color.blueRgb + ',.7)', 'rgba('+ app.color.successRgb + ',.7)', 'rgba('+ app.color.gray300Rgb + ',.7)', 'rgba('+ app.color.gray900Rgb + ',.7)'],
			borderColor: [app.color.indigo, app.color.blue, app.color.success, app.color.gray300, app.color.gray900],
			borderWidth: 2,
			label: 'My dataset'
		}]
	};

	var pieChartData = {
		labels: ['Dataset 1', 'Dataset 2', 'Dataset 3', 'Dataset 4', 'Dataset 5'],
		datasets: [{
			data: [300, 50, 100, 40, 120],
			backgroundColor: ['rgba('+ app.color.redRgb + ',.7)', 'rgba('+ app.color.orangeRgb + ',.7)', 'rgba('+ app.color.gray500Rgb + ',.7)', 'rgba('+ app.color.gray300Rgb + ',.7)', 'rgba('+ app.color.gray900Rgb + ',.7)'],
			borderColor: [app.color.red, app.color.orange, app.color.gray500, app.color.gray300, app.color.gray900],
			borderWidth: 2,
			label: 'My dataset'
		}]
	};

	var doughnutChartData = {
		labels: ['Dataset 1', 'Dataset 2', 'Dataset 3', 'Dataset 4', 'Dataset 5'],
		datasets: [{
			data: [300, 50, 100, 40, 120],
			backgroundColor: ['rgba('+ app.color.indigoRgb + ',.7)', 'rgba('+ app.color.blueRgb + ',.7)', 'rgba('+ app.color.successRgb + ',.7)', 'rgba('+ app.color.gray300Rgb + ',.7)', 'rgba('+ app.color.gray900Rgb + ',.7)'],
			borderColor: [app.color.indigo, app.color.blue, app.color.success, app.color.gray300, app.color.gray900],
			borderWidth: 2,
			label: 'My dataset'
		}]
	};

	var ctx = document.getElementById('line-chart').getContext('2d');
	lineChart = new Chart(ctx, {
		type: 'line',
		data: lineChartData
	});
	
	var ctx2 = document.getElementById('bar-chart').getContext('2d');
	barChart = new Chart(ctx2, {
		type: 'bar',
		data: barChartData
	});

	var ctx3 = document.getElementById('radar-chart').getContext('2d');
	radarChart = new Chart(ctx3, {
		type: 'radar',
		data: radarChartData
	});

	var ctx4 = document.getElementById('polar-area-chart').getContext('2d');
	polarAreaChart = new Chart(ctx4, {
		type: 'polarArea',
		data: polarAreaData
	});

	var ctx5 = document.getElementById('pie-chart').getContext('2d');
	pieChart = new Chart(ctx5, {
		type: 'pie',
		data: pieChartData
	});

	var ctx6 = document.getElementById('doughnut-chart').getContext('2d');
	doughnutChart = new Chart(ctx6, {
		type: 'doughnut',
		data: doughnutChartData
	});
};

var ChartJs = function () {
	"use strict";
	return {
		//main function
		init: function () {
			handleChartJs();
		}
	};
}();

$(document).ready(function() {
	ChartJs.init();
	
	$(document).on('theme-change', function() {
		lineChart.destroy();
		barChart.destroy();
		radarChart.destroy();
		polarAreaChart.destroy();
		pieChart.destroy();
		doughnutChart.destroy();
		
		handleChartJs();
	});
});