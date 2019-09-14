function hasClass(element, cls) {
    return (' ' + element.className + ' ').indexOf(' ' + cls + ' ') > -1;
}

function showing() {
    const panel = document.querySelectorAll('.splash-panel');
    for (let i = 0; i < panel.length; i++) {
      panel[i].classList.toggle('showing');
    }
    setTimeout(function () {
      showing();
    }, 3000);
}

showing();