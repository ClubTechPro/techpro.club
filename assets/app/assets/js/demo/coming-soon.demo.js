/*
Template Name: Color Admin - Responsive Admin Dashboard Template build with Twitter Bootstrap 5
Version: 5.1.4
Author: Sean Ngu
Website: http://www.seantheme.com/color-admin/
*/

var handleCountdownTimer = function() {
	var newYear = new Date();
	newYear = new Date(newYear.getFullYear() + 1, 1 - 1, 1);
	$('#timer').countdown({until: newYear});
};

var ComingSoon = function () {
	"use strict";
	return {
		//main function
		init: function () {
			handleCountdownTimer();
		}
	};
}();


$(document).ready(function() {
	ComingSoon.init();
});