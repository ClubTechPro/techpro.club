/*
Template Name: Color Admin - Responsive Admin Dashboard Template build with Twitter Bootstrap 5
Version: 5.1.4
Author: Sean Ngu
Website: http://www.seantheme.com/color-admin/
*/

function calculateDivider() {
	var dividerValue = 4;
	if ($(this).width() <= 576) {
		dividerValue = 1;
	} else if ($(this).width() <= 992) {
		dividerValue = 2;
	} else if ($(this).width() <= 1200) {
		dividerValue = 3;
	}
	return dividerValue;
}
var handleIsotopesGallery = function() {
	"use strict";
	var container = $('#gallery');
	var dividerValue = calculateDivider();
	var containerWidth = $(container).width();
	var columnWidth = containerWidth / dividerValue;
	$(container).isotope({
		resizable: true,
		masonry: {
			columnWidth: columnWidth
		}
	});
	
	$(window).on('resize', function() {
		var dividerValue = calculateDivider();
		var containerWidth = $(container).width();
		var columnWidth = containerWidth / dividerValue;
		$(container).isotope({
			masonry: { 
				columnWidth: columnWidth 
			}
		});
	});
	
	var $optionSets = $('#options #filter'),
	$optionLinks = $optionSets.find('a');
	
	$optionLinks.click( function(){
		var $this = $(this);
		if ($this.hasClass('active')) {
			return false;
		}
		var $optionSet = $this.parents('#filter');
		$optionSet.find('.active').removeClass('active');
		$this.addClass('active');
	
		var options = {};
		var key = $optionSet.attr('data-option-key');
		var value = $this.attr('data-option-value');
			value = value === 'false' ? false : value;
			options[ key ] = value;
		$(container).isotope( options );
		return false;
	});
};


var Gallery = function () {
	"use strict";
	return {
		//main function
		init: function () {
			handleIsotopesGallery();
		}
	};
}();

$(document).ready(function() {
	Gallery.init();
});