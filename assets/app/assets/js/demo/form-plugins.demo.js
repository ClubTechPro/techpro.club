/*
Template Name: Color Admin - Responsive Admin Dashboard Template build with Twitter Bootstrap 5
Version: 5.1.4
Author: Sean Ngu
Website: http://www.seantheme.com/color-admin/
*/

var handleRenderBootstrapTimePicker = function () {
	"use strict";
	$("#timepicker-default").timepicker();
};

var handleRenderClipboard = function() {
	var clipboard = new ClipboardJS("[data-toggle='clipboard']");
	
	clipboard.on("success", function(e) {
		$(e.trigger).tooltip({
			title: "Copied",
			placement: "top"
		});
		$(e.trigger).tooltip("show");
		setTimeout(function() {
			$(e.trigger).tooltip("dispose");
		}, 500);
	});
};

var handleRenderColorpicker = function() {
	$("#colorpicker-default").spectrum({
    showInput: true
	});
};

var handleRenderDatepicker = function() {
	$("#datepicker-default").datepicker({
		todayHighlight: true
	});
	$("#datepicker-inline").datepicker({
		todayHighlight: true
	});
	$(".input-daterange").datepicker({
		todayHighlight: true
	});
	$("#datepicker-disabled-past").datepicker({
		todayHighlight: true
	});
	$("#datepicker-autoClose").datepicker({
		todayHighlight: true,
		autoclose: true
	});
};

var handleRenderDateRangePicker = function() {
	$("#default-daterange").daterangepicker({
		opens: "right",
		format: "MM/DD/YYYY",
		separator: " to ",
		startDate: moment().subtract(29, "days"),
		endDate: moment(),
		minDate: "01/01/2012",
		maxDate: "12/31/2018",
	}, function (start, end) {
		$("#default-daterange input").val(start.format("MMMM D, YYYY") + " - " + end.format("MMMM D, YYYY"));
	});

	$("#advance-daterange span").html(moment().subtract(29, "days").format("MMMM D, YYYY") + " - " + moment().format("MMMM D, YYYY"));

	$("#advance-daterange").daterangepicker({
		format: "MM/DD/YYYY",
		startDate: moment().subtract(29, "days"),
		endDate: moment(),
		minDate: "01/01/2012",
		maxDate: "12/31/2015",
		dateLimit: { days: 60 },
		showDropdowns: true,
		showWeekNumbers: true,
		timePicker: false,
		timePickerIncrement: 1,
		timePicker12Hour: true,
		ranges: {
			"Today": [moment(), moment()],
			"Yesterday": [moment().subtract(1, "days"), moment().subtract(1, "days")],
			"Last 7 Days": [moment().subtract(6, "days"), moment()],
			"Last 30 Days": [moment().subtract(29, "days"), moment()],
			"This Month": [moment().startOf("month"), moment().endOf("month")],
			"Last Month": [moment().subtract(1, "month").startOf("month"), moment().subtract(1, "month").endOf("month")]
		},
		opens: "right",
		drops: "down",
		buttonClasses: ["btn", "btn-sm"],
		applyClass: "btn-primary",
		cancelClass: "btn-default",
		separator: " to ",
		locale: {
			applyLabel: "Submit",
			cancelLabel: "Cancel",
			fromLabel: "From",
			toLabel: "To",
			customRangeLabel: "Custom",
			daysOfWeek: ["Su", "Mo", "Tu", "We", "Th", "Fr","Sa"],
			monthNames: ["January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"],
			firstDay: 1
		}
	}, function(start, end, label) {
		$("#advance-daterange span").html(start.format("MMMM D, YYYY") + " - " + end.format("MMMM D, YYYY"));
	});
};

var handleRenderFormMaskedInput = function() {
	"use strict";
	$("#masked-input-date").mask("99/99/9999");
	$("#masked-input-phone").mask("(999) 999-9999");
	$("#masked-input-tid").mask("99-9999999");
	$("#masked-input-ssn").mask("999-99-9999");
	$("#masked-input-pno").mask("aaa-9999-a");
	$("#masked-input-pkey").mask("a*-999-a999");
};

var handleRenderIonRangeSlider = function() {
	$("#default_rangeSlider").ionRangeSlider({
		min: 0,
		max: 5000,
		type: "double",
		prefix: "$",
		maxPostfix: "+",
		prettify: false,
		hasGrid: true,
		skin: "big"
	});
	$("#customRange_rangeSlider").ionRangeSlider({
		min: 1000,
		max: 100000,
		from: 30000,
		to: 90000,
		type: "double",
		step: 500,
		postfix: " €",
		hasGrid: true,
		skin: "flat"
	});
	$("#customValue_rangeSlider").ionRangeSlider({
		values: [
			"January", "February", "March",
			"April", "May", "June",
			"July", "August", "September",
			"October", "November", "December"
		],
		type: "single",
		hasGrid: true
	});
};

var handleRenderJqueryAutocomplete = function() {
	var availableTags = [
		"ActionScript",
		"AppleScript",
		"Asp",
		"BASIC",
		"C",
		"C++",
		"Clojure",
		"COBOL",
		"ColdFusion",
		"Erlang",
		"Fortran",
		"Groovy",
		"Haskell",
		"Java",
		"JavaScript",
		"Lisp",
		"Perl",
		"PHP",
		"Python",
		"Ruby",
		"Scala",
		"Scheme"
	];
	$("#jquery-autocomplete").autocomplete({
		source: availableTags
	});
};

var handleRenderJqueryTagIt = function() {
	$("#jquery-tagIt-default").tagit({
		availableTags: ["c++", "java", "php", "javascript", "ruby", "python", "c"]
	});
	$("#jquery-tagIt-inverse").tagit({
		availableTags: ["c++", "java", "php", "javascript", "ruby", "python", "c"]
	});
	$("#jquery-tagIt-white").tagit({
		availableTags: ["c++", "java", "php", "javascript", "ruby", "python", "c"]
	});
	$("#jquery-tagIt-primary").tagit({
		availableTags: ["c++", "java", "php", "javascript", "ruby", "python", "c"]
	});
	$("#jquery-tagIt-info").tagit({
		availableTags: ["c++", "java", "php", "javascript", "ruby", "python", "c"]
	});
	$("#jquery-tagIt-success").tagit({
		availableTags: ["c++", "java", "php", "javascript", "ruby", "python", "c"]
	});
	$("#jquery-tagIt-warning").tagit({
		availableTags: ["c++", "java", "php", "javascript", "ruby", "python", "c"]
	});
	$("#jquery-tagIt-danger").tagit({
		availableTags: ["c++", "java", "php", "javascript", "ruby", "python", "c"]
	});
};

var handleRenderSelect2 = function() {
	$(".default-select2").select2();
	$(".multiple-select2").select2({ placeholder: "Select a state" });
};

var handleRenderSelectPicker = function() {
	$("#ex-basic").picker();
	$("#ex-multiselect").picker();
	$("#ex-search").picker({ search: true });
};

var handleRenderTimepicker = function() {
	$("#timepicker").timepicker();
};

var FormPlugins = function () {
	"use strict";
	return {
		//main function
		init: function () {
			handleRenderBootstrapTimePicker();
			handleRenderClipboard();
			handleRenderColorpicker();
			handleRenderDatepicker();
			handleRenderDateRangePicker();
			handleRenderFormMaskedInput();
			handleRenderIonRangeSlider();
			handleRenderJqueryAutocomplete();
			handleRenderJqueryTagIt();
			handleRenderSelect2();
			handleRenderSelectPicker();
			handleRenderTimepicker();
		}
	};
}();

$(document).ready(function() {
	FormPlugins.init();
	
	$(document).on('theme-change', function() {
		handleRenderColorpicker();
	});
});