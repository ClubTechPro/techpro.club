/*
Template Name: Color Admin - Responsive Admin Dashboard Template build with Twitter Bootstrap 5
Version: 5.1.4
Author: Sean Ngu
Website: http://www.seantheme.com/color-admin/
*/

var renderSwitcher = function() {
	if ($("[data-render='switchery']").length !== 0) {
		$("[data-render='switchery']").each(function() {
			var themeColor = app.color.success;
			if ($(this).attr("data-theme")) {
				switch ($(this).attr("data-theme")) {
					case "red":
						themeColor = app.color.red;
						break;
					case "blue":
						themeColor = app.color.blue;
						break;
					case "purple":
						themeColor = app.color.purple;
						break;
					case "orange":
						themeColor = app.color.orange;
						break;
					case "black":
						themeColor = app.color.black;
						break;
				}
			}
			var switchery = new Switchery(this, {
				color: themeColor
			});
		});
	}
};

var checkSwitcherState = function() {
	$(document).on("click", "[data-click='check-switchery-state']", function() {
		alert($("[data-id='switchery-state']").prop("checked"));
	});
	$(document).on("change", "[data-change='check-switchery-state-text']", function() {
		$("[data-id='switchery-state-text']").text($(this).prop("checked"));
	});
};

var renderPowerRangeSlider = function() {
	if ($("[data-render='powerange-slider']").length !== 0) {
		$("[data-render='powerange-slider']").each(function() {
			var option = {};
			option.decimal = ($(this).attr("data-decimal")) ? $(this).attr("data-decimal") : false;
			option.disable = ($(this).attr("data-disable")) ? $(this).attr("data-disable") : false;
			option.disableOpacity = ($(this).attr("data-disable-opacity")) ? parseFloat($(this).attr("data-disable-opacity")) : 0.5;
			option.hideRange = ($(this).attr("data-hide-range")) ? $(this).attr("data-hide-range") : false;
			option.klass = ($(this).attr("data-class")) ? $(this).attr("data-class") : "";
			option.min = ($(this).attr("data-min")) ? parseInt($(this).attr("data-min")) : 0;
			option.max = ($(this).attr("data-max")) ? parseInt($(this).attr("data-max")) : 100;
			option.start = ($(this).attr("data-start")) ? parseInt($(this).attr("data-start")) : null;
			option.step = ($(this).attr("data-step")) ? parseInt($(this).attr("data-step")) : null;
			option.vertical = ($(this).attr("data-vertical")) ? $(this).attr("data-vertical") : false;
			if ($(this).attr("data-height")) {
				$(this).closest(".powerange-wrapper").height($(this).attr("data-height"));
			}
			var powerange = new Powerange(this, option);
		});
	}
};

var checkPowerRangeState = function() {
	$(document).on("click", "[data-toggle='get-value-powerange']", function(e) {	
		e.preventDefault();
		alert($($(this).attr('data-target')).val());
	});
};

var FormSliderSwitcher = function () {
	"use strict";
	return {
		//main function
		init: function () {
			// switchery
			renderSwitcher();
			checkSwitcherState();

			// powerange slider
			renderPowerRangeSlider();
			checkPowerRangeState();
		}
	};
}();

$(document).ready(function() {
	FormSliderSwitcher.init();
});