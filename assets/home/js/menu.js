// Hamburger menu on small screen which shows page navigation
let menu = document.querySelector('#menu-bar');
let sidebar = document.querySelector('.sidebar-menu');

menu.addEventListener('click', () => {
  sidebar.classList.toggle('active');
});

sidebar.addEventListener('click', () => {
  sidebar.classList.toggle('active');
});
