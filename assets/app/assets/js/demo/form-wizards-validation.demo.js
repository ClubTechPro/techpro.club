/*
Template Name: Color Admin - Responsive Admin Dashboard Template build with Twitter Bootstrap 5
Version: 5.1.4
Author: Sean Ngu
Website: http://www.seantheme.com/color-admin/
*/

var handleBootstrapWizardsValidation = function() {
	"use strict";
	$('#wizard').smartWizard({ 
		selected: 0, 
		theme: 'default',
		transitionEffect:'',
		transitionSpeed: 0,
		useURLhash: false,
		showStepURLhash: false,
		toolbarSettings: {
			toolbarPosition: 'bottom'
		}
	});
	$('#wizard').on('leaveStep', function(e, anchorObject, stepNumber, stepDirection) {
		var res = $('form[name="form-wizard"]').parsley().validate('step-' + (stepNumber + 1));
		return res;
	});
	
	$('#wizard').keypress(function( event ) {
		if (event.which == 13 ) {
			$('#wizard').smartWizard('next');
		}
	});
};

var FormWizardValidation = function () {
	"use strict";
	return {
		//main function
		init: function () {
			handleBootstrapWizardsValidation();
		}
	};
}();

$(document).ready(function() {
	FormWizardValidation.init();
});