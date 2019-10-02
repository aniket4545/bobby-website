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

const loginPageElement = document.querySelector('.js-login-form');
document.querySelector('.js-login-page').addEventListener('click', () => {
  loginPageElement.style.display = 'block';
});

document.querySelector('.js-cancel-login').addEventListener('click', () => {
  loginPageElement.style.display = 'none';
});

const mobileHeaderElement = document.querySelector('.js-mobile-actions');
document.querySelector('.js-menu-icon').addEventListener('click', () => {
  mobileHeaderElement.style.display = 'block';
});

window.addEventListener('mouseup', (event) => {
  if (event.target != mobileHeaderElement && event.target.parentNode != mobileHeaderElement) {
    mobileHeaderElement.style.display = 'none';
  }
});