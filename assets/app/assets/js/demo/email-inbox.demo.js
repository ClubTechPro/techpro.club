/*
Template Name: Color Admin - Responsive Admin Dashboard Template build with Twitter Bootstrap 5
Version: 5.1.4
Author: Sean Ngu
Website: http://www.seantheme.com/color-admin/
*/

var handleEmailActionButtonStatus = function() {
	if ($('[data-checked=email-checkbox]:checked').length !== 0) {
		$('[data-email-action]').removeClass('hide');
	} else {
		$('[data-email-action]').addClass('hide');
	}
};

var handleEmailCheckboxChecked = function() {
	$(document).on('change', '[data-checked=email-checkbox]', function() {
		var targetLabel = $(this).closest('label');
		var targetEmailList = $(this).closest('li');
		if ($(this).prop('checked')) {
			$(targetLabel).addClass('active');
			$(targetEmailList).addClass('selected');
		} else {
			$(targetLabel).removeClass('active');
			$(targetEmailList).removeClass('selected');
		}
		handleEmailActionButtonStatus();
	});
};

var handleEmailAction = function() {
	$(document).on('click', '[data-email-action]', function() {
		var targetEmailList = '[data-checked="email-checkbox"]:checked';
		if ($(targetEmailList).length !== 0) {
			$(targetEmailList).closest('li').slideToggle(function() {
				$(this).remove();
				handleEmailActionButtonStatus();
				if ($('.list-email > li').length === 0) {
					$('[data-email-action]').addClass('hide');
					$('[data-checked=email-checkbox]').prop('checked', false);
					$('.list-email').html('<li class="p-3 text-center bg-white mb-n1"><div class="p-3"><i class="fa fa-trash fa-5x text-silver"></i></div> This folder is empty</li>');
				}
			});
		}
	});
};

var handleEmailSelectAll = function () {
	"use strict";
	$(document).on('change', '[data-change=email-select-all]', function() {
		if (!$(this).is(':checked')) {
			$('.list-email .email-checkbox input[type="checkbox"]').prop('checked', false);
		} else {
			$('.list-email .email-checkbox input[type="checkbox"]').prop('checked', true);
		}
		$('.list-email .email-checkbox input[type="checkbox"]').trigger('change');
	});
};

var EmailInbox = function () {
	"use strict";
	return {
		//main function
		init: function () {
			handleEmailCheckboxChecked();
			handleEmailAction();
			handleEmailSelectAll();
		}
	};
}();

$(document).ready(function() {
	EmailInbox.init();
});