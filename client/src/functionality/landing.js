function hasClass(element, cls) {
  return (' ' + element.className + ' ').indexOf(' ' + cls + ' ') > -1;
}

let slidesIndex = 0;

function showing() {
  const panel = document.querySelectorAll('.splash-panel');
  for (let i = 0; i < slidesIndex; i++) {
    panel[i].classList.remove('showing'); 
  }

  if(slidesIndex >= panel.length){
    slidesIndex = 0;
  }

  panel[slidesIndex].classList.add('showing');
  slidesIndex++;
  
  setTimeout(function () {
    showing();
  }, 5000);
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

setHeight = () => {
  const offsetWidth = document.querySelector('.js-image').offsetWidth;
  const halfOffsetWidth = offsetWidth / 4;
  const secondHalfOffsetWidth = offsetWidth / 8;
  const thirdHalfOffsetWidth = offsetWidth / 16;

  let tempValue = Number;
  tempValue = halfOffsetWidth + secondHalfOffsetWidth + thirdHalfOffsetWidth;
  const finalHeightToImage = Math.floor(tempValue);
  document.querySelector('.js-image').style.height = `${finalHeightToImage}px`;
}

this.setHeight();
window.addEventListener('resize', this.setHeight);