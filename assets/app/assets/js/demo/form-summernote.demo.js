/*
Template Name: Color Admin - Responsive Admin Dashboard Template build with Twitter Bootstrap 5
Version: 5.1.4
Author: Sean Ngu
Website: http://www.seantheme.com/color-admin/
*/

var handleSummernote = function() {
	$(".summernote").summernote({
		placeholder: "Hi, this is summernote. Please, write text here! Super simple WYSIWYG editor on Bootstrap",
		height: $(window).height() * 0.5
	});
};

var FormSummernote = function () {
	"use strict";
	return {
		//main function
		init: function () {
			handleSummernote();
		}
	};
}();

$(document).ready(function() {
	FormSummernote.init();
});