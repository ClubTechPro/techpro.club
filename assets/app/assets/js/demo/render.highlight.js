/*
Template Name: Color Admin - Responsive Admin Dashboard Template build with Twitter Bootstrap 5
Version: 5.1.4
Author: Sean Ngu
Website: http://www.seantheme.com/color-admin/
*/

var handleRenderHighlight = function() { 
	$('.hljs-wrapper pre code').each(function(i, block) {
		hljs.highlightBlock(block);
	});
};

var Highlight = function () {
	"use strict";
	return {
		//main function
		init: function () {
			handleRenderHighlight();
		}
	};
}();

$(document).ready(function() {
	Highlight.init();
});