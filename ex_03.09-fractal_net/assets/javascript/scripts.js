window.onload = function(e) {
	removeClass(document.body, ' loading');
}

function beforeLoad() {
	document.body.className += ' loading';
}

function hasClass(elem, className) {
	return new RegExp(' ' + className + ' ').test(' ' + elem.className + ' ');
}

function removeClass(elem, className) {
	var newClass = ' ' + elem.className.replace( /[\t\r\n]/g, ' ') + ' ';
	if (hasClass(elem, className)) {
		while (newClass.indexOf(' ' + className + ' ') >= 0 ) {
			newClass = newClass.replace(' ' + className + ' ', ' ');
		}
	elem.className = newClass.replace(/^\s+|\s+$/g, '');
	}
}
