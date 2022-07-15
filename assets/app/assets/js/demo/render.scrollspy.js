/*
Template Name: Color Admin - Responsive Admin Dashboard Template build with Twitter Bootstrap 5
Version: 5.1.4
Author: Sean Ngu
Website: http://www.seantheme.com/color-admin/
*/

var handleRenderScrollSpy = function() { 
	var scrollSpyTarget = '#navContent';
	var scrollSpyOffset = $('#header').height();
	var scrollSpy = new bootstrap.ScrollSpy(document.body, {
		target: scrollSpyTarget,
		offset: scrollSpyOffset
	});
};

var handleScrollTo = function() {
	$(document).on('click', '[data-toggle="scroll-to"]', function(e) {
		e.preventDefault();
		
		var targetId = $(this).attr('href');
		
		$('html, body').animate({
			scrollTop: $(targetId).offset().top - $('#header').height() + 1
		}, 250);
	});
};

var ScrollSpy = function () {
	"use strict";
	return {
		//main function
		init: function () {
			handleRenderScrollSpy();
			handleScrollTo();
		}
	};
}();

$(document).ready(function() {
	ScrollSpy.init();
});