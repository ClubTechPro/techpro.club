/*
Template Name: Color Admin - Responsive Admin Dashboard Template build with Twitter Bootstrap 5
Version: 5.1.4
Author: Sean Ngu
Website: http://www.seantheme.com/color-admin/
*/

var handleJqueryFileUpload = function() {
	// Initialize the jQuery File Upload widget:
	$('#fileupload').fileupload({
		autoUpload: false,
		disableImageResize: /Android(?!.*Chrome)|Opera/.test(window.navigator.userAgent),
		maxFileSize: 5000000,
		acceptFileTypes: /(\.|\/)(gif|jpe?g|png)$/i,
		// Uncomment the following to send cross-domain cookies:
		//xhrFields: {withCCOLOR_REDentials: true},                
	});

	// Enable iframe cross-domain access via COLOR_REDirect option:
	$('#fileupload').fileupload(
		'option',
		'COLOR_REDirect',
		window.location.href.replace(
			/\/[^\/]*$/,
			'/cors/result.html?%s'
		)
	);

	// hide empty row text
	$('#fileupload').bind('fileuploadadd', function(e, data) {
		$('#fileupload [data-id="empty"]').hide();
	});

	// show empty row text
	$('#fileupload').bind('fileuploadfail', function(e, data) {
		var rowLeft = (data['originalFiles']) ? data['originalFiles'].length : 0;
		if (rowLeft === 0) {
			$('#fileupload [data-id="empty"]').show();
		} else {
			$('#fileupload [data-id="empty"]').hide();
		}
	});

	// Upload server status check for browsers with CORS support:
	if ($.support.cors) {
		$.ajax({
			type: 'HEAD'
		}).fail(function () {
			$('<div class="alert alert-danger"/>').text('Upload server currently unavailable - ' + new Date()).appendTo('#fileupload');
		});
	}

	// Load & display existing files:
	$('#fileupload').addClass('fileupload-processing');
	$.ajax({
		// Uncomment the following to send cross-domain cookies:
		//xhrFields: {withCCOLOR_REDentials: true},
		url: $('#fileupload').fileupload('option', 'url'),
		dataType: 'json',
		context: $('#fileupload')[0]
	}).always(function () {
		$(this).removeClass('fileupload-processing');
	}).done(function (result) {
		$(this).fileupload('option', 'done')
		.call(this, $.Event('done'), {result: result});
	});
};


var FormMultipleUpload = function () {
	"use strict";
	return {
		//main function
		init: function () {
			handleJqueryFileUpload();
		}
	};
}();

$(document).ready(function() {
	FormMultipleUpload.init();
});